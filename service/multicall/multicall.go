package multicall

import (
	"context"
	"fmt"
	"github.com/chikob3/test/common"
	"github.com/chikob3/test/libs/contracts"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"
)

const (
	timeout    = 10 * time.Second
	rateLimit  = 4
	batchLimit = 200
)

type Client struct {
	sugar        *zap.SugaredLogger
	address      etherCommon.Address
	ethClient    *ethclient.Client
	multicallAbi abi.ABI
	rpcClient    *rpc.Client
}

func NewClient(sugar *zap.SugaredLogger, address etherCommon.Address, ethClient *ethclient.Client, rpcClient *rpc.Client) (*Client, error) {
	multicallAbi, err := abi.JSON(strings.NewReader(contracts.MulticallABI))
	if err != nil {
		return nil, fmt.Errorf("error loading multicall contract abi %w", err)
	}

	return &Client{
		sugar:        sugar,
		address:      address,
		ethClient:    ethClient,
		multicallAbi: multicallAbi,
		rpcClient:    rpcClient,
	}, nil
}

type multicallReturnData contracts.Multicall2Result

type multicallOutput struct {
	BlockNumber *big.Int
	Results     []multicallReturnData
}

type CallInput contracts.Multicall2Call

func (c *Client) Do(calls []CallInput) (multicallOutput, error) {
	if len(calls) == 0 {
		return multicallOutput{}, nil
	}

	type doPartialResult struct {
		output multicallOutput
		err    error
	}
	rateLimiter := make(chan interface{}, rateLimit)
	fn := func(ctx context.Context, callInputs []CallInput, ch chan doPartialResult) {
		rateLimiter <- nil
		defer func() {
			<-rateLimiter
		}()

		if ctx.Err() != nil {
			ch <- doPartialResult{
				output: multicallOutput{},
				err:    ctx.Err(),
			}
			return
		}
		out, err := c.doPartial(callInputs)
		ch <- doPartialResult{
			output: out,
			err:    err,
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numberOfReq := len(calls)/batchLimit + 1
	resultCh := make([]chan doPartialResult, numberOfReq)
	for i := 0; i < numberOfReq; i++ {
		start := i * batchLimit
		end := start + batchLimit
		if end > len(calls) {
			end = len(calls)
		}
		particalCalls := calls[start:end]

		ch := make(chan doPartialResult, 1)
		resultCh[i] = ch
		go fn(ctx, particalCalls, ch)
	}

	output := multicallOutput{
		Results: []multicallReturnData{},
	}
	for _, ch := range resultCh {
		result := <-ch
		if result.err != nil {
			return output, result.err
		}

		output.BlockNumber = result.output.BlockNumber
		output.Results = append(output.Results, result.output.Results...)
	}
	return output, nil
}

func Encode(target etherCommon.Address, method abi.Method, args ...interface{}) (CallInput, error) {
	data, err := common.EncodeData(method, args...)
	if err != nil {
		return CallInput{}, err
	}
	return CallInput{
		Target:   target,
		CallData: data,
	}, nil
}

func (c *Client) doPartial(calls []CallInput) (multicallOutput, error) {
	if len(calls) == 0 {
		return multicallOutput{}, nil
	}

	// Try using multicall contract first as it was better
	result, err := c.executeByContract(calls)
	if err == nil {
		return result, nil
	}

	if strings.Contains(err.Error(), "out of gas") {
		result, err = c.executeByBatch(calls)
		if err == nil {
			return result, nil
		}
		c.sugar.Debugw("error executing multicall by batch", "err", err, "inputs", len(calls))
	} else {
		c.sugar.Debugw("error executing multicall by contract", "err", err, "inputs", len(calls))
	}
	return multicallOutput{}, err
}

func (c *Client) executeByContract(calls []CallInput) (multicallOutput, error) {
	multicall, _ := common.EncodeData(c.multicallAbi.Methods["tryBlockAndAggregate"], false, calls)
	msg := ethereum.CallMsg{
		To:    &c.address,
		Value: big.NewInt(0),
		Data:  multicall,
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	bytes, err := c.ethClient.CallContract(ctx, msg, nil)
	if err != nil {
		return multicallOutput{}, fmt.Errorf("error executing multicall %w", err)
	}

	var output struct {
		BlockNumber *big.Int
		BlockHash   etherCommon.Hash
		ReturnData  []multicallReturnData
	}
	err = c.multicallAbi.Unpack(&output, "tryBlockAndAggregate", bytes)
	if err != nil {
		return multicallOutput{}, fmt.Errorf("error decoding multicall output, %w", err)
	}

	return multicallOutput{
		BlockNumber: output.BlockNumber,
		Results:     output.ReturnData,
	}, nil
}

func (c *Client) executeByBatch(calls []CallInput) (multicallOutput, error) {
	var batchElems []rpc.BatchElem
	for _, call := range calls {
		params := map[string]interface{}{
			"from": etherCommon.Address{}.Hex(),
			"to":   call.Target.Hex(),
			"data": etherCommon.ToHex(call.CallData),
		}

		var output string
		batchElems = append(batchElems, rpc.BatchElem{
			Method: "eth_call",
			Args: []interface{}{
				params,
				"latest",
			},
			Result: &output,
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := c.rpcClient.BatchCallContext(ctx, batchElems)
	if err != nil {
		return multicallOutput{}, err
	}

	var result multicallOutput
	for _, elem := range batchElems {
		returnData := multicallReturnData{
			Success:    false,
			ReturnData: []byte{},
		}
		if elem.Error == nil {
			out := elem.Result.(*string)
			returnData.Success = true
			returnData.ReturnData = etherCommon.FromHex(*out)
		}
		result.Results = append(result.Results, returnData)
	}

	return result, err
}
