package main

import (
	"testing"
)

func TestRLEncode(t *testing.T) {

	testCases := [][2]string{
		{"AABBCC", "2A2B2C"},
		{"TTEST", "2TEST"},
		{"AAABBBCCCCDDDEEFGHII", "3A3B4C3D2EFGH2I"},
		{"ABC", "ABC"},
		{"", ""},
		{"BLTT", "BL2T"},
	}

	for _, a := range testCases {
		got := RLEncode(a[0])
		if got != a[1] {
			t.Errorf("RLEncode(%v) = >>%v<<; want >>%v<<", a[0], got, a[1])
		}
	}

}
