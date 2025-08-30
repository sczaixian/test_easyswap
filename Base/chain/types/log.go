package types

import "math/big"

type FilterQuery struct {
	BlockHash string
	FromBlock *big.Int
	ToBlock   *big.Int
	Addresses []string
	Topics    [][]string // 条件过滤
}
