package main

import (
	"fmt"
)

var m = "0123456789bcdefghjkmnpqrstuvwxyz"

func decode(s string) (lng, lat float64) {
	lng = 0
	lat = 0
	decodeMap := make([]byte, 256)
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	for i := 0; i < len(m); i++ {
		decodeMap[m[i]] = byte(i)
	}
	odd := make([]byte, 0)
	even := make([]byte, 0)
	isOdd := false
	for i := 0; i < len(s); i++ {
		bit := byte(0x10)
		for j := 0; j < 5; j++ {
			d := (bit & decodeMap[s[i]]) >> uint(4-j)
			bit = bit >> 1
			if isOdd {
				odd = append(odd, d)
			} else {
				even = append(even, d)
			}
			isOdd = !isOdd
		}
	}
	lat = calc(even, -180, 180)
	lng = calc(odd, -90, 90)
	return
}

func encode(pos float64) string {
	return ""
}

func calc(arr []byte, min, max float64) float64 {
	var res float64
	res = 0.0
	for i := 0; i < len(arr); i++ {
		res = (min + max) / 2.0
		if arr[i] > 0 {
			min = res
		} else {
			max = res
		}
	}
	return res
}

func main() {
	a, b := decode("6gkzwgjzn820")
	fmt.Println(a, b)
}
