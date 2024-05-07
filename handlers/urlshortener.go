package handlers 

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
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