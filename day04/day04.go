package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func Run(secret string, prefix string) int {
	attempt := 0
	for {
		sum := md5.Sum([]byte(secret + strconv.Itoa(attempt)))
		result := hex.EncodeToString(sum[:])

		if strings.HasPrefix(result, prefix) {
			return attempt
		}

		attempt++
	}
}

func main() {
	fmt.Println(Run("bgvyzdsv", "00000"))
	fmt.Println(Run("bgvyzdsv", "000000"))
}
