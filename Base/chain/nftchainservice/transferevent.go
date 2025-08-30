package nftchainservice

import "github.com/ethereum/go-ethereum/common"

type TransferLog struct {
	Address         string        `json:"address" gencodec:"required"`
	TransactionHash string        `json:"transaction_hash" gencodec:"required"`
	BlockNumber     uint64        `json:"block_number"`
	BlockTime       uint64        `json:"block_time"`
	BlockHash       string        `json:"block_hash"`
	Data            []byte        `json:"data" gencodec:"required"`
	Topics          []common.Hash `json:"topics" gencodec:"required"`
	Topic0          string        `json:"topic0"`
	From            string        `json:"topic1"`
	To              string        `json:"topic2"`
	TokenID         string        `json:"topic3"`
	TxIndex         uint          `json:"transactionIndex"`
	Index           uint          `json:"logIndex"`
	Removed         bool          `json:"removed"`
}
