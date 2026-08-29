package main

import "fmt"

func lexGreaterPermutation(s string, target string) string {
	result := ""
	index := 0
	found := false
	for len(s) != 0 {
		if !found {
			
		}
	}	

}

func findSmallest(s string) rune {
	var char rune
	for _, c := range s {
		if c < char || char == 0 {
			char = c
		}
	}
	return char
}

func main() {
	c := findSmallest("bomchoar")
	fmt.Println(c)
}

// result = ""
// s = "abc" target = "bba"
// position 0 has b
// we have b in s --> put b in result & remove b from s
// position 1 has b ! we do not have b , look for a greater character than b
// c is greater than b --> put c in result & remove c from s
// once we found a greater character , the rest we put the smallest possible

// result = ""
// found = false (when we're done with equals, and put a greater character, found = true)
// when found == true, we start ordering them by the smallest possible

// approach :
// index = 0
// while s != ""
// if !found --> we look for equals
