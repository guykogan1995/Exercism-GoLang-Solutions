// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var sSlice []string
	if len(rhyme) == 0 {
		return sSlice
	} else if len(rhyme) == 1 {
		return append(sSlice, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	} else {
		for i, j := 0, 1; i < len(rhyme); i, j = i+1, j+1 {
			if i != len(rhyme)-1 {
				sSlice = append(sSlice, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[j]))
			} else {
				sSlice = append(sSlice, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			}
		}
		return sSlice
	}
}
