package main

import (
	"fmt"
)

func main() {
	fmt.Println("DCP-29  Run the unit tests!")

}

// RLEncode will perform run length encoding
// ex : s=AAABBCDEFFF, result: 3A2BCDE3F
func RLEncode(s string) string {
	ret := ""
	if s == "" {
		return s
	}

	var lastChar byte
	cnt := 0

	for x := 0; x < len(s); x++ {
 		writeIt := false

		if s[x] == lastChar {
			cnt++
		} else {
			cnt = 1
		}

		// if we are at the end write it out
		// if we are not at the end and the next character is different, write it out
		//  and reset lastChar and cnt

		if x == len(s)-1 { // if we are at the last character, we are doing a write
			writeIt = true
		} else {
			if s[x+1] != s[x] { // next character is different from this character, we are going to write
				writeIt = true
			}
		}

		if writeIt {
			if cnt > 1 {
				ret += fmt.Sprintf("%d%v", cnt, string(s[x]))
			} else {
				ret += fmt.Sprintf("%v", string(s[x]))
			}
			cnt = 0
		}
		lastChar = s[x]
	}

	return ret
}
