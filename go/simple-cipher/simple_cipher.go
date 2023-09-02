package cipher

import (
	"regexp"
	"strings"
)

type shift struct {
	distance int
}

type vigenere struct {
	distance []int
}

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
	a := ""
	for _, v := range strings.ToLower(input) {
		if int(v) < 97 || int(v) > 122 {
			continue
		}
		c := int(v) + c.distance
		if c > 122 {
			c -= 26
		}
		if c < 97 {
			c += 26
		}
		a += string(rune(c))
	}
	return a
}

func (c shift) Decode(input string) string {
	a := ""
	for _, v := range strings.ToLower(input) {
		c := int(v) - c.distance
		if c < 97 {
			c += 26
		}
		if c > 122 {
			c -= 26
		}
		a += string(rune(c))
	}
	return a
}

func NewVigenere(key string) Cipher {
	count := 0
	for _, x := range key {
		if x == 'a' {
			count++
		}
	}
	r := regexp.MustCompile("([a-z]+)")
	if r.FindString(key) == "" || strings.ContainsRune(key, ',') || strings.ContainsRune(key, ' ') || count == len(key) {
		return nil
	}
	a := make([]int, len(key))
	for i, x := range key {
		a[i] = int(x - 97)
	}
	return vigenere{distance: a}
}

func (v vigenere) Encode(input string) string {
	a := ""
	input = strings.ToLower(input)
	count := 0
	r := regexp.MustCompile("([a-z]+)")
	input2 := r.FindAllString(input, -1)
	input = ""
	for i := 0; i < len(input2); i++ {
		input += input2[i]
	}
	for i := 0; i < len(input); i++ {
		if i > len(v.distance)-1 {
			i = 0
		}
		c := rune(int(input[count]) + v.distance[i])
		if c > 122 {
			c -= 26
		}
		count++
		a += string(c)
		if count == len(input) {
			break
		}
	}
	return a
}

func (v vigenere) Decode(input string) string {
	a := ""
	input = strings.ToLower(input)
	count := 0
	for i := 0; i < len(input); i++ {
		if i > len(v.distance)-1 {
			i = 0
		}
		c := rune(int(input[count]) - v.distance[i])
		if c < 97 {
			c += 26
		}
		count++
		a += string(c)
		if len(a) == len(input) {
			break
		}
	}
	return a
}
