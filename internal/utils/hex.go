package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Hex2int(hexStr string) uint64 {
	cleaned := strings.Replace(hexStr, "0x", "", 1)
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return result
}

func Int2hex(num uint64) string {
	return fmt.Sprintf("0x%x", num)
}
