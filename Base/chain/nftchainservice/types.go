package nftchainservice

type JsonMetadata struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Image       string                  `json:"image"`
	Decimals    *int                    `json:"decimals"`
	Attributes  []*OpenseaMetadataProps `json:"attributes"`
}

type OpenseaMetadataProps struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}
