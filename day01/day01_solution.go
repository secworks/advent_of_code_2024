package main

import (
	"bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)


// Read a file and return a slice with the lines.
func ReadLines(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}


// Read the given file  and return as two sorted slices.
func GetSortSlices(filename string) ([]int, []int) {
	// Load the input into an array.
    lines, _ := ReadLines(filename)

	// Extract fields, convert to ints and append to slices.
	var a_ops []int
	var b_ops []int
	for i := 0 ; i < len(lines) ; i += 1 {
		strops := strings.Fields(lines[i])
		a, _ := strconv.Atoi(strops[0])
		a_ops = append(a_ops, a)
		b, _ := strconv.Atoi(strops[1])
		b_ops = append(b_ops, b)
	}	

	// Sort the slices.
	sort.Ints(a_ops)
	sort.Ints(b_ops)

	return a_ops, b_ops
}


// Calculate the absolute diffeerence between
// the given values.
func Abs(x int, y int) int {
	if x < y {
		return y - x
	} else {
	return x - y
	}
}


// Get sum of differences between the sorted columns.
func problem1(fname string) {
	var a_ops []int
	var b_ops []int

	a_ops, b_ops = GetSortSlices(fname)
	
	// Loop over the slices and add diff.
	diffsum := 0
	for i := 0 ; i < len(a_ops) ; i += 1 {
		diffsum += Abs(a_ops[i], b_ops[i])
	}
	fmt.Println("Total distance:", diffsum)
}


// Get similarity score.
func problem2(fname string) {
	var a_ops []int
	var b_ops []int

	a_ops, b_ops = GetSortSlices(fname)

	score := 0
	// Loop over all elements in a_ops.
	// Check number of occurences of the element ai.
	for i := 0 ; i < len(a_ops) ; i += 1 {
		var ai = a_ops[i]
		occ := 0
		for j := 0 ; j < len(b_ops) ; j += 1 {
			if b_ops[j] == ai {
				occ += 1
			}
		}
		score += ai * occ
	}
	fmt.Println("Total similarity score:", score)
}


func main() {
	fmt.Println("AoC 2024, day 01")
	fmt.Println("----------------")
	problem1("day01_input.txt")
	problem2("day01_input.txt")
}

