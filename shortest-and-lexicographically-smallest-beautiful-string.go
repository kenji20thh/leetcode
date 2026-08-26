package main

import "fmt"

func shortestBeautifulSubstring(s string, k int) string {
	best := ""
	for i := 0; i < len(s); i++ {
		count := 0

		for j := i; j < len(s); j++ {
			if s[j] == '1' {
				count++
			}
			if count == k {
				if best == "" || len(s[i:j+1]) < len(best) || (len(s[i:j+1]) == len(best) && s[i:j+1] < best) {
					best = s[i : j+1]
				}
				break
			}

		}
	}
	return best
}

func main() {
	s := "100011001"
	k := 3
	fmt.Println(shortestBeautifulSubstring(s, k))
}
