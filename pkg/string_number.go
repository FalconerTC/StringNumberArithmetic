package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StringNumber struct {
	value string
}

// Used to add the numeric values to another StringNumber
// Treats non-numeric values as 0s
// Ex:
// 		sn := StringNumber{2}
//		sn.Add(2)
//		sn.value === 4
func (sn *StringNumber) Add(other StringNumber) {
	snArr := extractNums(sn.value)
	oArr := extractNums(other.value)

	if len(snArr) <= len(oArr) {
		sn.value = add(snArr, oArr)
	} else {
		sn.value = add(oArr, snArr)
	}
}

// Used to multiply the numeric values to another StringNumber
// Treats non-numeric values as 0s
// Ex:
// 		sn := StringNumber{2}
//		sn.Multiply(3)
//		sn.value === 6
func (sn *StringNumber) Multiply(other StringNumber) {
	snArr := extractNums(sn.value)
	oArr := extractNums(other.value)

	if len(snArr) <= len(oArr) {
		sn.value = multiply(snArr, oArr)
	} else {
		sn.value = multiply(oArr, snArr)
	}
}

func extractNums(str string) []uint8 {
	extracted := make([]uint8, 0)
	for _, char := range str {
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
	if len(short) == 0 && len(long) == 0 {
		return "0"
	}

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

func multiply(short, long []uint8) string {
	if len(short) == 0 || len(long) == 0 {
		return "0"
	}

	reverseInts(&short)
	reverseInts(&long)

	final := ""
	for i := 0; i < len(short); i++ {

		nextSum := make([]string, i)
		// Fill with preceding zeros
		for c := 0; c < i; c++ {
			nextSum = append(nextSum, "0")
		}

		carry := uint8(0)
		for j := 0; j < len(long); j++ {
			mult := (short[i] * long[j]) + carry
			if mult >= 10 {
				carry = uint8(mult / 10)
				nextSum = append([]string{strconv.Itoa(int(mult % 10))}, nextSum...)
			} else {
				carry = 0
				nextSum = append([]string{strconv.Itoa(int(mult))}, nextSum...)
			}
		}
		if carry != 0 {
			nextSum = append([]string{strconv.Itoa(int(carry))}, nextSum...)
		}
		final = add(extractNums(final), extractNums(strings.Join(nextSum, "")))
	}

	return final
}

// Used for testing
func main() {
	a := StringNumber{value: "356"}
	b := StringNumber{value: "21"}

	fmt.Println(a, b)
	a.Multiply(b)
	fmt.Println(a, b)
}
