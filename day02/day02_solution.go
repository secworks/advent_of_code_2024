// AoC 2024, day 02

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


func check_values_problem1(a, b int, dir bool) int {
	if (b <= a) && dir {
		return 1
	}

	if (b >= a) && !dir {
		return 1
	}

	if Abs(a, b) > 3 {
		return 1
	}

	return 0
}


// Get sum of differences between the sorted columns.
func problem1(filename string) {
    lines, _ := ReadLines(filename)

	safe_lines := 0
	for i := 0 ; i < len(lines) ; i += 1 {
		numbers := strings.Fields(lines[i])

		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[len(numbers) - 1])
		errors := 0

		var increasing bool
		if (b > a) {
			increasing = true
		} else if (b < a) {
			increasing = false
		} else {
			increasing = false
			errors += 1
		}

		for j := 0 ; j < len(numbers) - 1 ; j += 1 {
			a, _ := strconv.Atoi(numbers[j])
			b, _ := strconv.Atoi(numbers[j + 1])

			errors += check_values_problem1(a, b, increasing)
		}

		if errors == 0 {
			safe_lines += 1
		}
	}
	
	fmt.Println("Total safe lines:", safe_lines)
}


func check_values_problem2(a, b int, dir bool) int {
	if (a > b) && dir {
		return 1
	}
	
	if (a < b) && !dir {
		return 1
	}

	if (a == b) {
		return 1
	}

	if Abs(a, b) > 3 {
		return 1
	}

	return 0
}


// Get sum of differences between the sorted columns.
func problem2(filename string) {
    lines, _ := ReadLines(filename)
	safe_lines := 0

	for i := 0 ; i < len(lines) ; i += 1 {
		numbers := strings.Fields(lines[i])

		// No errors at the beginning of a line.
		errors := 0

		// Figure out the general direction.
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])
		var increasing bool
		if (b > a) {
			increasing = true
		} else {
			increasing = false
		}
		
		for j := 0 ; j < len(numbers) - 1 ; j += 1 {
			a, _ := strconv.Atoi(numbers[j])
			b, _ := strconv.Atoi(numbers[j + 1])

			errors += check_values_problem2(a, b, increasing)
		}
		
		if errors < 2 {
			safe_lines += 1
		}
	}
	
	fmt.Println("Total safe lines:", safe_lines)
}


func main() {
	fmt.Println("AoC 2024, day 02")
	fmt.Println("----------------")
	problem1("day02_input.txt")
	problem2("day02_example1.txt")
}

