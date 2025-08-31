package chain

import (
	"github.com/ProjectsTask/Base/evm/eip"
	"github.com/pkg/errors"
	"strings"
)

const (
	Eth      = "eth"
	Optimism = "optimism"
	Sepolia  = "sepolia"
)

const (
	EthChainID      = 1
	OptimismChainID = 10
	SepoliaChainID  = 11155111
)

func UniformAddress(chainName string, address string) (string, error) {
	address, err := eip.ToCheckSumAddress(address)
	if err != nil {
		return "", errors.Wrap(err, "failed on uniform evm chain address")
	}
	return strings.ToLower(address), nil
}
