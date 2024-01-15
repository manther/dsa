package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Return a reversed string using a queue
// This version uses a stack
// The time complexity is O(n)
func reverseString(s string) string {
	stck := stack.New()
	for _, c := range s {
		stck.Push(c)
	}
	rt := []rune{}
	for i := 0; i < len(s); i++ {
		res := stck.Pop()
		val, ok := res.(rune)
		if !ok {
			return ""
		}
		rt = append(rt, val)
	}
	return string(rt)
}

// Return a new reversed string
// This version uses a slice of runes
// The time complexity is O(n)
func reverseString2(s string) string {
	stck := []rune{}
	for _, c := range s {
		stck = append([]rune{c}, stck...)
	}

	return string(stck)
}

// Return a new reversed byte slice
// This version uses a slice of bytes
// The time complexity is O(n)
func reverseString3(s []byte) []byte {
	stck := []byte{}
	for _, c := range s {
		stck = append([]byte{c}, stck...)
	}

	return stck
}

// Reverse a byte slice in place
// This version uses a single loop
// It swaps the first and last characters
// Then itterates inwards
// The time complexity is O(n)
func reverseString4(s []byte) {
	for i, c := range s {
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = c
	}
}

// Return true if the string is a palandrome
// This version uses a slice of runes
// The slice acts like a stack
// The time complexity is O(n)
func palandrome(s string) bool {
	stck := []rune{}
	for _, c := range s {
		stck = append([]rune{c}, stck...)
	}

	return s == string(stck)
}

// Return true if the string is a palandrome
// This version ignores case and non-alphanumeric characters
// and uses a single loop that exits as early as possible
func palandrome2(s string) bool {
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
	s = strings.ToLower(s)
	for i, _ := range s {
		if s[i] != s[len(s)-i-1] {
			return false
		}
		if len(s)/2 == i || len(s)-(i+1) == i {
			break
		}
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Return true if the linked list holds a palandrome
// Converts the linked list to a slice of ints
// and uses a single loop that exits as early as possible
func palandromeLLInt(head *ListNode) bool {
	tmp := []int{}
	if head == nil {
		return false
	}
	for {
		tmp = append(tmp, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	for i, _ := range tmp {
		if tmp[i] != tmp[len(tmp)-i-1] {
			return false
		}
		if len(tmp)/2 == i || len(tmp)-(i+1) == i {
			break
		}
	}

	return true
}

// Return reversed linked list
// Return true if the linked list holds a palandrome
// This version creates a second head
// and builds the reversed linked list from that head
// The time complexity is O(n)
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var revHead *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = revHead
		revHead = head
		head = tmp
	}
	return revHead
}

// Return the number of common words between two slices of strings
// only count a word is it is used exactly once in each slice
// This version uses two maps, as frequency maps
// The time complexity is O(n + m)
// because it loops through each slice once
// and then loops through the smaller map
// to find the common words
// The space complexity is O(n)
// The space complexity is proportional to the size of both slices.
func countWords(words1 []string, words2 []string) int {
	m, v := map[string]int{}, map[string]int{}
	for _, w := range words1 {
		m[w]++
	}

	for _, w := range words2 {
		v[w]++
	}

	var count int
	for k := range m {
		if _, ok := v[k]; ok {
			if v[k] == 1 && m[k] == 1 {
				count++
			}
		}
	}

	return count
}

// Reverse a number slice in place
// If the number is negative, retain the negative sign
// This version uses a pointer to the int
// as it will modify the value of the int in place
func reverseInt(n *int) {
	isNeg := math.Signbit(float64(*n))
	s := strconv.Itoa(*n)
	s = strings.TrimLeft(s, "-")

	stck := []rune{}
	for _, c := range s {
		stck = append([]rune{c}, stck...)
	}
	if isNeg {
		stck = append([]rune{'-'}, stck...)
	}
	pi, err := strconv.ParseInt(string(stck), 10, 32)
	if err != nil {
		println(err.Error())
	}
	*n = int(pi)
}

// Reverse a number slice in place
// This version uses a pointer to the int
func reverseInt2(n *int) {
	isNeg := math.Signbit(float64(*n))
	s := strconv.Itoa(*n)
	s = strings.TrimLeft(s, "-")

	stck := []rune{}
	for _, c := range s {
		stck = append([]rune{c}, stck...)
	}

	pi, err := strconv.ParseInt(string(stck), 10, 32)
	if err != nil {
		println(err.Error())
		return
	}

	*n = int(pi)
	if isNeg {
		*n = *n * -1
	}
}

// Return the maximum repeated character in a string
// If there is a tie, return the first character
// If the string is empty, return the empty string
func maxChar(s string) rune {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	var max int
	var maxChar rune
	for k, v := range m {
		if v > max {
			max = v
			maxChar = k
		}
	}
	return maxChar
}

// Return a slice of string represented numbers up to n
// If the number is divisible by 3, replace with "Fizz"
// If the number is divisible by 5, replace with "Buzz"
// If the number is divisible by 3 and 5, replace with "FizzBuzz"
func fizzBuzz(n int) []string {
	arrStrn := make([]string, 0, n)
	// arrStrn := []string{}
	for i := 0; i <= n; i++ {
		switch true {
		case i == 0:
		case i%3+i%5 == 0:
			arrStrn = append(arrStrn, "FizzBuzz")
		case i%3 == 0:
			arrStrn = append(arrStrn, "Fizz")
		case i%5 == 0:
			arrStrn = append(arrStrn, "Buzz")
		default:
			arrStrn = append(arrStrn, strconv.Itoa(i))
		}
	}
	return arrStrn
}

// Return a slice of slices of ints
// Each slice (chunk) of ints is of size n
// The last slice of ints may be smaller than n
func chunkSlice(s []int, size int) [][]int {
	chunked := [][]int{}

	for i := 0; i < len(s); i += size {
		if i%size == 0 { //Don't need this, are incrementing by "size" already
			chunk := []int{}
			if i+size > len(s) {
				chunk = append(chunk, s[i:]...)
			} else {
				chunk = append(chunk, s[i:i+size]...)
			}
			chunked = append(chunked, chunk)
		}
	}

	return chunked
}

// Better version of chunkSlice
// Don't need to keep track of the index modulo
// Don't need to expand the appended slice
// Just assign the sliced slice directly to the chucked slice
func chunkSlice2(s []int, size int) [][]int {
	chunked := [][]int{}

	for i := 0; i < len(s); i += size {
		if i+size > len(s) {
			chunked = append(chunked, s[i:])
		} else {
			chunked = append(chunked, s[i:i+size])
		}
	}
	return chunked
}

// Return true if the string is an anagram of the other string
// Ignore case and non-alphanumeric characters
// The time complexity is O(n + m)
func anagrams(a string, b string) bool {
	am, bm := map[rune]int{}, map[rune]int{}
	a = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(a, ""))
	b = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(b, ""))
	for _, c := range a {
		am[c]++
	}

	for _, c := range b {
		bm[c]++
	}

	for k, v := range am {
		if bm[k] != v {
			return false
		}
	}
	return true
}

// Return true if the string is an anagram of the other string
// Ignore case and non-alphanumeric characters
// Uses sort.Strings to sort the strings
// Then compares the sorted strings
// The time complexity is O(n log n + m log m)
func bananagrams(a string, b string) bool {
	a = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(a, ""))
	b = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(b, ""))

	as := strings.Split(a, "")
	bs := strings.Split(b, "")
	sort.Strings(as)
	sort.Strings(bs)
	return strings.Join(as, "") == strings.Join(bs, "")
}

// Return a string with the first letter of each word capitalized
// This version is big O(n)
func capitalize(s string) string {
	sl := strings.Split(s, " ")
	for i, w := range sl {
		sw := strings.Split(w, "")
		sw[0] = strings.ToUpper(sw[0])
		sl[i] = strings.Join(sw, "")
	}

	return strings.Join(sl, " ")
}

// Return a string with the first letter of each word capitalized
// This version uses the golang.org/x/text/cases package
// This package is not in the standard library
// This version is big O(n)
func capitalize2(s string) string {
	sl := []string{}
	for _, w := range strings.Split(s, " ") {
		sl = append(sl, cases.Title(language.English).String(w))
	}

	return strings.Join(sl, " ")
}

// Write a function that accepts a positive number N. 
// The function should console log a step shape with N levels using the # character.  
// Make sure the step has spaces on the right hand side!
func steps(n int) {
	for i := 1; i <= n; i++ {
		s := strings.Repeat(" ", n-i)
		t := strings.Repeat("#", i)
		fmt.Printf("'%s%s'\n",t, s)
	}
}