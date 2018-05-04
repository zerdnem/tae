package utils

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type HashAlgorithm string

const (
	HashAlgorithmSHA256 HashAlgorithm = "sha256"
	HashAlgorithmSHA1   HashAlgorithm = "sha1"
	HashAlgorithmMD5    HashAlgorithm = "md5"
)

type Hash struct {
	Algorithm HashAlgorithm
	HashValue []byte
}

func (ha HashAlgorithm) FromString(s string) (*Hash, error) {
	l := -1
	switch ha {
	case HashAlgorithmMD5:
		l = 32
	case HashAlgorithmSHA1:
		l = 40
	case HashAlgorithmSHA256:
		l = 64
	default:
		return nil, fmt.Errorf("unknown hash algorithm: %q", ha)
	}

	if len(s) != l {
		return nil, fmt.Errorf("invalid %q hash - unexpected length %d", ha, len(s))
	}

	hashValue, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("invalid hash %q - not hex", s)
	}
	return &Hash{Algorithm: ha, HashValue: hashValue}, nil
}

func FromString(s string) (*Hash, error) {
	for _, ha := range []HashAlgorithm{HashAlgorithmMD5, HashAlgorithmSHA1, HashAlgorithmSHA256} {
		prefix := fmt.Sprintf("%s:", ha)
		if strings.HasPrefix(s, prefix) {
			return ha.FromString(s[len(prefix):])
		}
	}

	var ha HashAlgorithm
	switch len(s) {
	case 32:
		ha = HashAlgorithmMD5
	case 40:
		ha = HashAlgorithmSHA1
	case 64:
		ha = HashAlgorithmSHA256
	default:
		return nil, fmt.Errorf("Cannot determine algorithm for hash length: %d", len(s))
	}

	return ha.FromString(s)
}
