package chainclient

import (
	"context"
	"math/big"

	"Base/chain"
	"Base/chain/chainclient/evmclient"
	logTypes "Base/chain/types"

	"github.com/ethereum/go-ethereum"
	"github.com/pkg/errors"
)

type ChainClient interface {
	FilterLogs(ctx context.Context, q logTypes.FilterQuery) ([]interface{}, error)
	BlockTimeByNumber(ctx context.Context, int2 *big.Int) (uint64, error)
	Client() interface{}
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	CallContractByChain(ctx context.Context, param logTypes.CallParam) (interface{}, error)
	BlockNumber() (uint64, error)
	BlockWithTxs(ctx context.Context, blockNumber uint64) (interface{}, error)
}

func New(chainID int, nodeUrl string) (ChainClient, error) {
	switch chainID {
	case chain.EthChainID, chain.OptimismChainID, chain.SepoliaChainID:
		return evmclient.New(nodeUrl)
	default:
		return nil, errors.New("unsupported chain id")
	}
}
