package main

import "fmt"

// GO - Generate Valid Combinations of n Parentheses pairs
// Given an integer, generate an array containing all pairs of valid parentheses.
// Parentheses are valid if (()()(()))() each pair opens and closes properly. For example, )( would NOT be valid! (Neither would ()))
// generateParens(1) => ["()"]
// generateParens(2) => ["()()","(())"]
// generateParens(3) => ["()()()", "()(())", "(())()", "(()())", "((()))"]

// Strings - one at a time
func recursiveParens(n int, current string, open int, close int) {
	if close == n {
		fmt.Println(current)
		return
	} else {
		if open > close {
			newCurrent := current + ")"
			recursiveParens(n, newCurrent, open, close+1)
		}
		if open < n {
			newCurrent := current + "("
			recursiveParens(n, newCurrent, open+1, close)
		}
	}
	return
}

// Array of strings
func recursiveParensArray(n int, current string, results *[]string, open int, close int) *[]string {
	if close == n {
		fmt.Println(current)
		*results = append(*results, current)


		return results
	} else {
		if open > close {
			newCurrent := current + ")"
			recursiveParensArray(n, newCurrent, results, open, close+1)
		}
		if open < n {
			newCurrent := current + "("
			recursiveParensArray(n, newCurrent, results, open+1, close)
		}
	}
	fmt.Println(len(*results))
	return results
}

func main() {
	var myResults = &[]string{}
	fmt.Println("\n-----print parens-----\n")
	recursiveParens(2, "", 0, 0)
	fmt.Println("\n-----print array of parens-----\n")
	fmt.Println(recursiveParensArray(3, "", myResults, 0, 0 ))
}

