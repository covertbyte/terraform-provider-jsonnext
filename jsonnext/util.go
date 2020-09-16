package jsonnext

import (
	"crypto/sha256"
	"encoding/hex"
	"foxygo.at/jsonnext"
)

type config struct {
	importer jsonnext.Importer
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}
