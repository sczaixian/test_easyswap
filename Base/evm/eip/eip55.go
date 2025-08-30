package eip

import (
	"encoding/hex"
	"errors"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"
)

func ToCheckSumAddress(address string) (string, error) {
	if address == "" {
		return "", errors.New("address is empty")
	}

	if strings.HasPrefix(address, "0x") {
		address = address[2:]
	}
	bytes, err := hex.DecodeString(address)
	if err != nil {
		return "", err
	}
	hash := calculateKeccak256(bytes)
	result := "0x"
	for i, b := range bytes {
		result += checksumByte(b>>4, hash[i]>>4)
		result += checksumByte(b&0xf, hash[i]&0xf)
	}
	return result, nil
}

func calculateKeccak256(addr []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(addr)
	return hash.Sum(nil)
}

func checksumByte(addr byte, hash byte) string {
	result := strconv.FormatUint(uint64(hash), 16)
	if hash >= 8 {
		return strings.ToUpper(result)
	} else {
		return result
	}
}
