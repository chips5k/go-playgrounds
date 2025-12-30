package main

func main() {
	isBalanced("(( ))")
}

func isBalanced(input string) bool {

	if len(input) == 0 {
		return true
	}

	// if len(input)%2 != 0 {
	// 	return false
	// }

	if input[0] == ')' {
		return false
	}
	open := 0
	for _, s := range input {

		if s == '(' {
			open++
		}

		if s == ')' && open == 0 {
			return false
		}

		if s == ')' {
			open--
		}
	}

	return open == 0
}
