package bchutil

import (
	"errors"

	"github.com/btcsuite/btcd/address/v2"
	"github.com/btcsuite/btcd/txscript/v2"
)

func PayToAddrScript(addr address.Address) ([]byte, error) {
	var script []byte
	var err error
	script, err = txscript.PayToAddrScript(addr)
	if err == nil {
		return script, nil
	}
	script, err = cashPayToAddrScript(addr)
	if err == nil {
		return script, nil
	}
	script, err = bitpayPayToAddrScript(addr)
	if err == nil {
		return script, nil
	}
	return script, errors.New("Unrecognized address format")
}
