package token

import (
	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/chikob3/test/common"
)

type Wrapper interface {
	Wrap(address etherCommon.Address) etherCommon.Address
	Unwrap(address etherCommon.Address) etherCommon.Address
}

type SimpleWrapper struct {
	nativeAddress etherCommon.Address
	wrapAddress   etherCommon.Address
}

func NewSimpleWrapper(native, wrap etherCommon.Address) *SimpleWrapper {
	return &SimpleWrapper{
		nativeAddress: native,
		wrapAddress:   wrap,
	}
}

func (s SimpleWrapper) Wrap(address etherCommon.Address) etherCommon.Address {
	if address == s.nativeAddress || address == common.ETHAddress {
		return s.wrapAddress
	}

	return address
}

func (s SimpleWrapper) Unwrap(address etherCommon.Address) etherCommon.Address {
	if address == s.wrapAddress {
		return s.nativeAddress
	}

	return address
}
