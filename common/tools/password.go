package tools

import (
	"github.com/duke-git/lancet/v2/cryptor"
)

func SHA256(str string) string {
	return cryptor.Sha256WithBase64(str)
}
