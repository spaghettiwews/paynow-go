package paynow

import (
	"crypto/sha512"
	"encoding/hex"
	"sort"
	"strings"
)

// MakeHash ...
func MakeHash(_args map[string]string, integrationKey string) string {
	var stringer string
	args := _args

	var keys []string
	for key := range args {
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

// VerifyHash ...
func VerifyHash() bool {
	return false
}