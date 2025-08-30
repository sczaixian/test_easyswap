package evmclient

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Service struct {
	client *ethclient.Client
}

func New(nodeUrl string) (*Service, error) {
	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		return nil, errors.Wrap(err, "Dial")
	}
	return &Service{client: client}, nil
}
