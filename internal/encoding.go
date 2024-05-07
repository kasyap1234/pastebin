package internal 

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)
func GenerateShortURL(longURL string) (string,error){
	// generating md5 hash for the long url 

	
	hash :=md5.Sum([]byte(longURL))
	// converting the hash to hash string 
	hashString :=hex.EncodeToString(hash[:])
   // now converting the hashstring to base62 
   base62Str :=base62Encode(hashString); 
   shortURL :="http://pastebin.com/"+base62Str
   return shortURL, nil ; 

}
func base62Encode(input string) string {
	const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte

	num := new(big.Int)
	num.SetString(input, 16)

	base := big.NewInt(62)
	zero := big.NewInt(0)

	for num.Cmp(zero) > 0 {
		rem := new(big.Int)
		num.DivMod(num, base, rem)
		result = append([]byte{base62Chars[rem.Int64()]}, result...)
	}

	return string(result)
}
