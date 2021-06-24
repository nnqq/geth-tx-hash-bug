package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func main() {
	// both real ropsten tx https://ropsten.etherscan.io/tx/0xb4848204c8432070136a41792003caf8dea08f9eb284eb4240845bf64a66a068
	legacyTx()
	accessListTx()
}

func legacyTx() {
	// real testnet tx https://ropsten.etherscan.io/tx/0xb4848204c8432070136a41792003caf8dea08f9eb284eb4240845bf64a66a068
	nonce, err := hexutil.DecodeUint64("0x1216")
	if err != nil {
		panic(err)
	}
	gasPrice, err := hexutil.DecodeBig("0x2bd0875aed")
	if err != nil {
		panic(err)
	}
	gas, err := hexutil.DecodeUint64("0x5208")
	if err != nil {
		panic(err)
	}
	to := common.HexToAddress("0x2f14582947e292a2ecd20c430b46f2d27cfe213c")
	value, err := hexutil.DecodeBig("0x2386f26fc10000")
	if err != nil {
		panic(err)
	}
	data := common.Hex2Bytes("0x")
	v, err := hexutil.DecodeBig("0x1")
	if err != nil {
		panic(err)
	}
	r, err := hexutil.DecodeBig("0x56b5bf9222ce26c3239492173249696740bc7c28cd159ad083a0f4940baf6d03")
	if err != nil {
		panic(err)
	}
	s, err := hexutil.DecodeBig("0x5fcd608b3b638950d3fe007b19ca8c4ead37237eaf89a8426777a594fd245c2a")
	if err != nil {
		panic(err)
	}

	newLegacyTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gas,
		To:       &to,
		Value:    value,
		Data:     data,
		V:        v,
		R:        r,
		S:        s,
	})
	fmt.Println("LegacyTx expected hash     => 0xb4848204c8432070136a41792003caf8dea08f9eb284eb4240845bf64a66a068")
	fmt.Println("LegacyTx actual hash       =>", newLegacyTx.Hash().String())
}

func accessListTx() {
	nonce, err := hexutil.DecodeUint64("0x1216")
	if err != nil {
		panic(err)
	}
	gasPrice, err := hexutil.DecodeBig("0x2bd0875aed")
	if err != nil {
		panic(err)
	}
	gas, err := hexutil.DecodeUint64("0x5208")
	if err != nil {
		panic(err)
	}
	to := common.HexToAddress("0x2f14582947e292a2ecd20c430b46f2d27cfe213c")
	value, err := hexutil.DecodeBig("0x2386f26fc10000")
	if err != nil {
		panic(err)
	}
	data := common.Hex2Bytes("0x")
	v, err := hexutil.DecodeBig("0x1")
	if err != nil {
		panic(err)
	}
	r, err := hexutil.DecodeBig("0x56b5bf9222ce26c3239492173249696740bc7c28cd159ad083a0f4940baf6d03")
	if err != nil {
		panic(err)
	}
	s, err := hexutil.DecodeBig("0x5fcd608b3b638950d3fe007b19ca8c4ead37237eaf89a8426777a594fd245c2a")
	if err != nil {
		panic(err)
	}
	chainID := new(big.Int).SetUint64(3)
	accessList := make(types.AccessList, 0)

	newAccessListTx := types.NewTx(&types.AccessListTx{
		Nonce:      nonce,
		GasPrice:   gasPrice,
		Gas:        gas,
		To:         &to,
		Value:      value,
		Data:       data,
		V:          v,
		R:          r,
		S:          s,
		ChainID:    chainID,
		AccessList: accessList,
	})
	fmt.Println("AccessListTx expected hash => 0xb4848204c8432070136a41792003caf8dea08f9eb284eb4240845bf64a66a068")
	fmt.Println("AccessListTx actual hash   =>", newAccessListTx.Hash().String())
}
