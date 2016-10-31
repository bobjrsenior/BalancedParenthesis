package main

import "fmt"

type node struct {
	Prev *node
	Value rune
}

func main() {
	result := checkBalance("}([)]")
	if result {
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}

func checkBalance(input string) (bool){
	var stack *node 
	pairs := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	if input == "" {
		return true
	}
	for _, c := range input {
		if stack == nil {
			stack = new (node)
			stack.Value = c
			
		}else if stack.Value == pairs[c] {
			stack = stack.Prev
		}else if c == '{' || c == '(' || c == '['{
			newNode := new (node)
			newNode.Prev = stack
			newNode.Value = c
			stack = newNode
		}else{
			return false
		}
	}
	
	return stack == nil
}
