package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	etherCommon "github.com/ethereum/go-ethereum/common"
)

func IsInList(list []string, item string) bool {
	for _, itemInList := range list {
		if itemInList == item {
			return true
		}
	}
	return false
}

func MakeGetRequest(url string, query map[string]string, timeout time.Duration, response interface{}) error {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body := res.Body
	defer body.Close()

	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return errors.New(string(bytes))
	}

	return json.Unmarshal(bytes, response)
}

func IsEthAddress(address etherCommon.Address) bool {
	return address == ETHAddress
}

func FormatSigningMessage(rawMessage string) string {
	bytes := []byte(rawMessage)
	return fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(bytes), bytes)
}

func EncodeData(method abi.Method, args ...interface{}) ([]byte, error) {
	var result []byte
	inputs, err := method.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}

	result = append(result, method.ID...)
	result = append(result, inputs...)

	return result, nil
}

func GetEnv(keys ...string) string {
	for _, key := range keys {
		if value, ok := os.LookupEnv(key); ok {
			return value
		}
	}
	return ""
}

func GetEnvWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func PrettyPrint(obj interface{}) string {
	json, _ := json.MarshalIndent(obj, "", "\t")
	return string(json)
}

func IndexedArgs(input abi.Arguments) abi.Arguments {
	var indexed abi.Arguments
	for _, arg := range input {
		if !arg.Indexed {
			continue
		}

		indexed = append(indexed, arg)
	}
	return indexed
}

func ConvertIpfs(url string) string {
	if strings.HasPrefix(url, "ipfs://") {
		return strings.ReplaceAll(url, "ipfs://", "https://cloudflare-ipfs.com/ipfs/")
	}
	return url
}

func WaitFor(f func(chan error), duration time.Duration) {
	done := make(chan error, 1)
	if duration < 0 {
		// Duration < 0 means blocking function
		go f(done)
		<-done
	} else {
		timeout := time.After(duration)
		go f(done)
		// Either timeout or function done
		select {
		case <-timeout:
		case <-done:
		}
	}
}

func RoundDown(x float64, decimal int) float64 {
	n := math.Pow(10, float64(decimal))
	return math.Floor(x*n) / n
}

func MakePostRequest(url string, headers map[string]string, query map[string]string, timeout time.Duration, response interface{}) error {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body := res.Body
	defer body.Close()

	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return GetRequestHttpError{
			StatusCode: res.StatusCode,
			Body:       bytes,
		}
	}
	return json.Unmarshal(bytes, response)
}

type GetRequestHttpError struct {
	StatusCode int
	Body       []byte
}

func (a GetRequestHttpError) Error() string {
	return fmt.Sprintf("status: %d, body: %s", a.StatusCode, string(a.Body))
}
