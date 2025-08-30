package shortner

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// generate a hash
func Sha256of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

// encode byts
func Base58Encoding(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateShortUrl(url, userId string) string{
	hashBytes := Sha256of(url + userId)
	generatedNumber := new(big.Int).SetBytes(hashBytes).Uint64()
	finalString := Base58Encoding([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
