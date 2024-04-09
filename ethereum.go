package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to an Ethereum node
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Construct a new block with dummy data
	block := types.NewBlock(&types.Header{
		ParentHash: common.HexToHash("0000000000000000000000000000000000000000000000000000000000000000"),
		Number:     big.NewInt(123456), // Block number
		Time:       uint64(time.Now().Unix()),
	}, nil, nil, nil)

	// Sign the block
	if err := block.SignBlock(123456, func(hash common.Hash) ([]byte, error) {
		// Dummy private key for signing (not real, do not use in production!)
		privateKey := common.FromHex("YOUR_PRIVATE_KEY")
		signature, err := types.SignTx(hash.Bytes(), types.HomesteadSigner{}, privateKey)
		if err != nil {
			return nil, err
		}
		return signature, nil
	}); err != nil {
		log.Fatal(err)
	}

	// Validate the block
	if err := block.Validate(); err != nil {
		log.Fatal(err)
	}

	// Print the block hash
	fmt.Println("Block Hash:", block.Hash().Hex())
}
