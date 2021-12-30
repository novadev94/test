package multicall

import (
	"context"
	"fmt"
	"github.com/chikob3/test/libs/contracts"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) AggregateCalls(addr common.Address, _abi abi.ABI, methodNames []string, inputs [][]interface{}, outputs []interface{}) error {
	methods := []abi.Method{}
	for _, name := range methodNames {
		method, exist := _abi.Methods[name]
		if !exist {
			return fmt.Errorf("method %s is not supported by abi", name)
		}
		methods = append(methods, method)
	}

	calls, err := PackCalls(addr, methods, inputs)
	if err != nil {
		return err
	}

	data, err := c.MakeCalls(calls)
	if err != nil {
		return err
	}

	return UnpackCallResults(data, methods, outputs)
}

func PackCalls(contractAddr common.Address, methods []abi.Method, inputs [][]interface{}) ([]contracts.Multicall2Call, error) {
	var calls []contracts.Multicall2Call
	for i, method := range methods {
		callData, err := method.Inputs.Pack(inputs[i]...)
		if err != nil {
			return nil, err
		}
		calls = append(calls, contracts.Multicall2Call{
			Target:   contractAddr,
			CallData: append(method.ID, callData...),
		})
	}
	return calls, nil
}

func (c *Client) MakeCalls(calls []contracts.Multicall2Call) ([][]byte, error) {
	callData, err := c.multicallAbi.Pack("aggregate", calls)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{
		To:   &c.address,
		Data: callData,
	}
	data, err := c.ethClient.CallContract(context.TODO(), msg, nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		BlockNumber *big.Int
		ReturnData  [][]byte
	}
	err = c.multicallAbi.Unpack(&result, "aggregate", data)
	if err != nil {
		return nil, err
	}

	return result.ReturnData, nil
}

func UnpackCallResults(data [][]byte, methods []abi.Method, outputs []interface{}) error {
	for i, method := range methods {
		err := method.Outputs.Unpack(outputs[i], data[i])
		if err != nil {
			return err
		}
	}
	return nil
}
