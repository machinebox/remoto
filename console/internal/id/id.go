// Package id is a pretty string to int64 converter for IDs in URLs.
package id

import (
	"math/big"
	"strconv"
)

var chars = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// Itoa encodes an int64 into a string.
func Itoa(val int64) string {
	var result []byte
	var index int
	var strVal string
	base := big.NewInt(int64(len(chars)))
	a := big.NewInt(0)
	b := big.NewInt(0)
	c := big.NewInt(0)
	d := big.NewInt(0)
	exponent := 1
	remaining := big.NewInt(val)
	for remaining.Cmp(big.NewInt(0)) != 0 {
		a.Exp(base, big.NewInt(int64(exponent)), nil) //16^1 = 16
		b = b.Mod(remaining, a)                       //119 % 16 = 7 | 112 % 256 = 112
		c = c.Exp(base, big.NewInt(int64(exponent-1)), nil)
		d = d.Div(b, c)
		strVal = d.String()
		index, _ = strconv.Atoi(strVal)
		result = append(result, chars[index])
		remaining = remaining.Sub(remaining, b) //119 - 7 = 112 | 112 - 112 = 0
		exponent = exponent + 1
	}
	return string(reverse(result))
}

// Atoi decodes a string into an int64.
func Atoi(s string) int64 {
	chars2 := reverse([]byte(s))
	dictMap := make(map[byte]*big.Int)
	j := 0
	for _, val := range chars {
		dictMap[val] = big.NewInt(int64(j))
		j = j + 1
	}
	bi := big.NewInt(0)
	base := big.NewInt(int64(len(chars)))
	exponent := 0
	b := big.NewInt(0)
	intermed := big.NewInt(0)
	for _, c := range chars2 {
		a := dictMap[c]
		intermed = intermed.Exp(base, big.NewInt(int64(exponent)), nil)
		b = b.Mul(intermed, a)
		bi = bi.Add(bi, b)
		exponent = exponent + 1
	}
	return bi.Int64()
}

func reverse(bs []byte) []byte {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return bs
}
