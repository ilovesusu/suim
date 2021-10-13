package pkg

import (
	"crypto/rand"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
)

// RandNumber 创建一个在10,000,000,000内的随机数
func RandNumber(logger *log.Helper) (string, error) {
	var (
		result *big.Int
		err    error
	)
	for i := 1; i <= 3; i++ {
		result, err = rand.Int(rand.Reader, big.NewInt(10000000000))
		if err != nil {
			logger.Log(log.LevelError, "Error", err)
			continue
		}
		break
	}
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
