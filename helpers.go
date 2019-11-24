package paynow-go

import (
	"crypto/sha512"
	"encoding/hex"
	"sort"
	"strings"
)

func MakeHash(_args map[string]string, int_key string) string {
	var stringer string
	args := _args

	var keys []string
	for key, _ := range args {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		if strings.ToUpper(key) != "HASH" {
			stringer = stringer + args[key]
		}
	}
	
	stringer = stringer + int_key
	alg := sha512.New()
	alg.Write([]byte(stringer))
	hash := hex.EncodeToString(alg.Sum(nil))
	return strings.ToUpper(hash)
}

func VerifyHash() bool {
	return false
}