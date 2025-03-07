// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import "strings"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// countURLsWithTLD() returns the number of URLs with the specified TLD
func countURLsWithTLD(urls []string, tld string) int {
    // Your code goes here.
    var count int = 0
    for _, url := range urls {
        if strings.HasSuffix(url, tld){
            count++
        }
    }
    return count
}
