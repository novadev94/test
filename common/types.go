package common

import (
	"encoding/json"
	"errors"
	"math/big"
	"reflect"
)

type ErrResponse struct {
	Message string `json:"error,omitempty"`
	Success bool   `json:"success"`
}

func (e *ErrResponse) String() string {
	b, _ := json.Marshal(e)
	return string(b[:])
}

type ApiParamsInfo struct {
	Name         string
	Required     bool
	DefaultValue interface{}
	Type         reflect.Type
}

type Number big.Int

func (a *Number) MarshalJSON() ([]byte, error) {
	str := (*big.Int)(a).String()
	return json.Marshal(str)
}

func (a *Number) UnmarshalJSON(bytes []byte) error {
	var valueStr string
	err := json.Unmarshal(bytes, &valueStr)
	if err != nil {
		return err
	}
	_, ok := (*big.Int)(a).SetString(valueStr, 0)
	if !ok {
		return errors.New("invalid numeric string")
	}

	return nil
}

func (a *Number) Int() *big.Int {
	return (*big.Int)(a)
}

type HexNumber big.Int

func (a *HexNumber) MarshalJSON() ([]byte, error) {
	str := (*big.Int)(a).Text(16)
	return json.Marshal("0x" + str)
}

func (a *HexNumber) UnmarshalJSON(bytes []byte) error {
	var valueStr string
	err := json.Unmarshal(bytes, &valueStr)
	if err != nil {
		return err
	}
	_, ok := (*big.Int)(a).SetString(valueStr, 0)
	if !ok {
		return errors.New("invalid numeric string")
	}
	return nil
}

func (a *HexNumber) Int() *big.Int {
	return (*big.Int)(a)
}
