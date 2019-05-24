package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StringNumber struct {
	value string
}

func (sn *StringNumber) Add(other StringNumber) {
	snArr := extractNums(*sn)
	oArr := extractNums(other)

	if len(snArr) <= len(oArr) {
		sn.value = add(snArr, oArr)
	} else {
		sn.value = add(oArr, snArr)
	}
	fmt.Println("final", sn.value)
}

func (sn *StringNumber) Multiply(other StringNumber) {

}

func extractNums(sn StringNumber) []uint8 {
	extracted := make([]uint8, 0)
	for _, char := range sn.value {
		// Treat non-numeric values as 0
		if char < 48 || char > 57 {
			extracted = append(extracted, 0)
		} else {
			extracted = append(extracted, uint8(char-48))
		}
	}
	return extracted
}

func reverseInts(input *[]uint8) {
	for i, j := 0, len(*input)-1; i < j; i, j = i+1, j-1 {
		(*input)[i], (*input)[j] = (*input)[j], (*input)[i]
	}
}

func add(short, long []uint8) string {
	reverseInts(&short)
	reverseInts(&long)

	str := make([]string, 0)
	carry := uint8(0)
	for i := 0; i < len(long); i++ {
		var sum uint8
		if i >= len(short) {
			sum = long[i] + carry
		} else {
			sum = long[i] + short[i] + carry
		}

		if sum >= 10 {
			carry = 1
			str = append([]string{strconv.Itoa(int(sum % 10))}, str...)
		} else {
			carry = 0
			str = append([]string{strconv.Itoa(int(sum))}, str...)
		}
	}
	return strings.Join(str, "")
}

func multiply() {

}

func main() {
	a := StringNumber{value: "11"}
	b := StringNumber{value: "10"}

	fmt.Println(a, b)
	a.Add(b)
	fmt.Println(a, b)
}
