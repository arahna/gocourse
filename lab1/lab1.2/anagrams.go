package main

import (
	"strings"
	"fmt"
	"sort"
)

type sortBytes []byte
type sortRunes []rune

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

func sortStringAsBytes(s string) string {
	r := []byte(s)
	sort.Sort(sortBytes(r))
	return string(r)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortStringsAsRunes(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func isAnagram1(first, second string) bool {
	if first == "" || second == "" {
		return false
	}

	firstSorted := sortStringsAsRunes(strings.ToLower(first))
	secondSorted := sortStringsAsRunes(strings.ToLower(second))

	return strings.Compare(firstSorted, secondSorted) == 0
}

func isAnagram2(first, second string) bool {
	if first == "" || second == "" {
		return false
	}

	firstSorted := sortStringAsBytes(strings.ToLower(first))
	secondSorted := sortStringAsBytes(strings.ToLower(second))

	return strings.Compare(firstSorted, secondSorted) == 0
}

func main() {
	fmt.Printf("isAnagram1('Апельсин', 'Спаниель') == %v\n", isAnagram1("Апельсин", "Спаниель"))
	fmt.Printf("isAnagram1('Апельсины', 'Спаниель') == %v\n", isAnagram1("Апельсины", "Спаниель"))

	fmt.Printf("isAnagram2('Апельсин', 'Спаниель') == %v\n", isAnagram2("Апельсин", "Спаниель"))
	fmt.Printf("isAnagram2('Апельсины', 'Спаниель') == %v\n", isAnagram2("Апельсины", "Спаниель"))
}