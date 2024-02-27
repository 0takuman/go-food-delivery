package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hasher struct{}

func NewMd5Hasher() *md5Hasher {
	return &md5Hasher{}
}

func (md5Hasher) Hash(data string) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
