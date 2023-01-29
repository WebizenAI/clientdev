package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	var lookup = []string{}
	var revLookup = []string{}
	var code = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	for i := 0; i < len(code); i++ {
		lookup = append(lookup, string(code[i]))
		revLookup = append(revLookup, string(code[i]))
	}
	revLookup[45] = "62"
	revLookup[95] = "63"

	b64 := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	fmt.Println(getLens(b64))
}

func getLens(b64 string) []int {
	var len = len(b64)
	var validLen = strings.Index(b64, "=")
	if validLen == -1 {
		validLen = len
	}
	var placeHoldersLen = 0
	if validLen == len {
		placeHoldersLen = 0
	} else {
		placeHoldersLen = 4 - (validLen % 4)
	}
	return []int{validLen, placeHoldersLen}
}

func _byteLength(b64 string, validLen int, placeHoldersLen int) int {
	return ((validLen + placeHoldersLen) * 3 / 4) - placeHoldersLen
}

func toByteArray(b64 string) [256]uint8 {
	var tmp int
	var lens [2]int
	var validLen int
	var placeHoldersLen int
	var arr [256]uint8
	var curByte int
	var len int
	var i int

	lens = getLens(b64)
	validLen = lens[0]
	placeHoldersLen = lens[1]

	arr = new Arr(_byteLength(b64, validLen, placeHoldersLen))

	curByte = 0

	len = placeHoldersLen > 0
		? validLen - 4
		: validLen

	for i = 0; i < len; i += 4 {
		tmp =
			(revLookup[b64.charCodeAt(i)] << 18) |
				(revLookup[b64.charCodeAt(i+1)] << 12) |
				(revLookup[b64.charCodeAt(i+2)] << 6) |
				revLookup[b64.charCodeAt(i+3)]
		arr[curByte++] = (tmp >> 16) & 0xFF
		arr[curByte++] = (tmp >> 8) & 0xFF
		arr[curByte++] = tmp & 0xFF
	}

	if placeHoldersLen == 2 {
		tmp =
			(revLookup[b64.charCodeAt(i)] << 2) |
				(revLookup[b64.charCodeAt(i+1)] >> 4)
		arr[curByte++] = tmp & 0xFF
	}

	if placeHoldersLen == 1 {
		tmp =
			(revLookup[b64.charCodeAt(i)] << 10) |
				(revLookup[b64.charCodeAt(i+1)] << 4) |
				(revLookup[b64.charCodeAt(i+2)] >> 2)
		arr[curByte++] = (tmp >> 8) & 0xFF
		arr[curByte++] = tmp & 0xFF
	}

	return arr
}

func tripletToBase64(num int) string {
	return lookup[num>>18&0x3F] +
		lookup[num>>12&0x3F] +
		lookup[num>>6&0x3F] +
		lookup[num&0x3F]
}

func encodeChunk(uint8 [256]uint8, start int, end int) string {
	var tmp int
	var output [256]string
	var i int
	for i = start; i < end; i += 3 {
		tmp =
			((uint8[i] << 16) & 0xFF0000) +
				((uint8[i+1] << 8) & 0xFF00) +
				(uint8[i+2] & 0xFF)
		output = append(output, tripletToBase64(tmp))
	}
	return strings.Join(output[:], "")
}

func fromByteArray(uint8 [256]uint8) string {
	var tmp int
	var len int
	var extraBytes int
	var parts [256]string
	var maxChunkLength int
	var i int
	var len2 int

	len = len(uint8)
	extraBytes = len % 3
	maxChunkLength = 16383