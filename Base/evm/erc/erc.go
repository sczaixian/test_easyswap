package erc

type NftErc struct {
	Endpoint string `toml:"endpoint" json:"endpoint"`
	Standard string `toml:"standard" json:"standard"`
}

type Erc interface {
	GetItemOwner(address string, tokenId string) (string, error)
}
