package main

import (
    "fmt"
    "strings"
)

// function to take list of prefixes as slice , and input string 

// I need to iterate over the list prefixes using for loop to check if the input string starts with the prefix 

//iterates over the prefixes and checks if the input string starts with the prefix   ---
// and if the length of the prefix is greater than the length of the current longest prefix found. 
//If both conditions are true, it updates the longestPrefix variable to be the current prefix.

func findLongestPrefix(prefixes []string, input string) string {
    longestPrefix := ""
    for _, prefix := range prefixes {
 
		 
        if strings.HasPrefix(input, prefix) && len(prefix) > len(longestPrefix) {
            longestPrefix = prefix
        }
    }
    return longestPrefix
}

func main() {

	// pass the array of data 
    prefixes := []string{"tamil","tamilnadu", "nadu",  "allindiaradio", "india", "radio"}
	input := "allindiaradio  state largest beach in asia"
    longestPrefix := findLongestPrefix(prefixes, input)
    fmt.Println(longestPrefix)
}