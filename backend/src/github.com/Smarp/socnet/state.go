package socnet

import (
	"crypto/rand"
	"encoding/base64"
)

func State() string {
	stateByte := make([]byte, 3)
	rand.Read(stateByte)
	return base64.StdEncoding.EncodeToString(stateByte)
}
