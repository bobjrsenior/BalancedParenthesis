package main

import "fmt"

// Stack Stuct to hold a stack of opening braces
type node struct {
	Prev *node
	Value rune
}

func main() {
	result := checkBalance("[](){}[[[]]]{(()){}}")
	if result {
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}

func checkBalance(input string) (bool){
	// Map closing braces to their corresponding open braces
	pairs := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	
	// Declare start of the stack
	var stack *node 

	for _, c := range input {
		
		if stack == nil { // If the stack is empty, make it a new node
			stack = new (node)
			stack.Value = c
		}else if stack.Value == pairs[c] { // If the current value is a closing brace matching the top value on the stack
			// Pop the stack
			stack = stack.Prev
		}else if c == '{' || c == '(' || c == '['{ // If it is a valid open brace
			// Push it to the stack
			newNode := new (node)
			newNode.Prev = stack
			newNode.Value = c
			stack = newNode
		}else{ // If it isn't an open brace and doesn't match the top of the stack
			// It's an invalid string
			return false
		}
	}
	
	return stack == nil
}
