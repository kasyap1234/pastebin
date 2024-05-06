package handlers 

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)
func GenerateShortURL(longURL string)