package internal 
import (
	"crypto/rand"
	"encoding/base64"
)
func CreateLongURL {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return base64.URLEncoding.EncodeToString(randomBytes)
}
