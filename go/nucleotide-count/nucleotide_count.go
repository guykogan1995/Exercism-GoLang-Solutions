package dna

import (
	"errors"
	"regexp"
)

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	var m Histogram
	m = map[rune]int{
		'A': 0,
		'C': 0,
		'T': 0,
		'G': 0,
	}
	r, _ := regexp.Compile(`^[ACTG]*$`)
	if !r.MatchString(string(d)) {
		return m, errors.New("non actg characters exist")
	}
	for _, v := range d {
		m[v]++
	}
	return m, nil
}
