// AoC 2024, day 02

package main

import (
	"regexp"
    "fmt"
)

func test1() {
    var re1 = regexp.MustCompile(`(?m)mul\([0-9]*,[0-9]*\)`)
    var re2 = regexp.MustCompile(`\([0-9]*`)
    var re3 = regexp.MustCompile(`(?m),[0-9]*`)

    var str = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
    
    for i, match1 := range re1.FindAllString(str, -1) {
        fmt.Println(match1, "found at index", i)
		match2 := re2.FindAllString(match1, 1)
		match3 := re3.FindAllString(match1, 1)
		var opa []string = match2[1:]
		var opb []string = match3[1:]
        fmt.Println(match2)
        fmt.Println(match3)
        fmt.Println(opa, opb)
    }
}


	
func main() {
	fmt.Println("AoC 2024, day 03")
	fmt.Println("----------------")
	test1()
//	problem1("day02_input.txt")
//	problem2("day02_example1.txt")
}

