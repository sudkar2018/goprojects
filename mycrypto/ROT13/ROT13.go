package myCrypto

import (
	"bytes"
	"unicode"
)

var charMap map[string]string

func GetInstance() {
	charMap = make(map[string]string)

	charMap["a"] = "n"
	charMap["b"] = "o"
	charMap["c"] = "p"
	charMap["d"] = "q"
	charMap["e"] = "r"
	charMap["f"] = "s"
	charMap["g"] = "t"
	charMap["h"] = "u"
	charMap["i"] = "v"
	charMap["j"] = "w"
	charMap["k"] = "x"
	charMap["l"] = "y"
	charMap["m"] = "z"
	charMap["n"] = "a"
	charMap["o"] = "b"
	charMap["p"] = "c"
	charMap["q"] = "d"
	charMap["r"] = "e"
	charMap["s"] = "f"
	charMap["t"] = "g"
	charMap["u"] = "h"
	charMap["v"] = "i"
	charMap["w"] = "j"
	charMap["x"] = "k"
	charMap["y"] = "l"
	charMap["z"] = "m"
	charMap["A"] = "N"
	charMap["B"] = "O"
	charMap["C"] = "P"
	charMap["D"] = "Q"
	charMap["E"] = "R"
	charMap["F"] = "S"
	charMap["G"] = "T"
	charMap["H"] = "U"
	charMap["I"] = "V"
	charMap["J"] = "W"
	charMap["K"] = "X"
	charMap["L"] = "Y"
	charMap["M"] = "Z"
	charMap["N"] = "A"
	charMap["O"] = "B"
	charMap["P"] = "C"
	charMap["Q"] = "D"
	charMap["R"] = "E"
	charMap["S"] = "F"
	charMap["T"] = "G"
	charMap["U"] = "H"
	charMap["V"] = "I"
	charMap["W"] = "J"
	charMap["X"] = "K"
	charMap["Y"] = "L"
	charMap["Z"] = "M"
}
func EncryptDecrypt(s string) string {
	GetInstance()
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		k := unicode.IsNumber(rune(s[i]))
		if k {
			buf.WriteByte(s[i])
		} else {
			buf.WriteString(charMap[string(s[i])])

		}

	}

	return buf.String()
}
