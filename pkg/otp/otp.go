package otp

import (
	"strconv"
	"strings"

	"math/rand/v2"
)

const codeLength = 6

func GenerateCode() string {
	var code strings.Builder

	code.Grow(codeLength)

	for i := 0; i < codeLength; i++ {
		code.WriteString(strconv.Itoa(rand.IntN(10)))
	}

	return code.String()
}
