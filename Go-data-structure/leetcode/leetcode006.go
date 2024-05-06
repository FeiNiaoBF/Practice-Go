package leetcode

/*
Useing array to simulate stack
*/

func push(s []string, c string) []string {
	return append(s, c)
}

func pop(s []string) []string {
	return s[:len(s)-1]
}

func peek(s []string) string {
	return s[len(s)-1]
}

func isValid(s string) bool {
	stack := new([]string)
	count := 0
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			*stack = push(*stack, string(c))
			count++
		} else if c == ')' || c == ']' || c == '}' {
			if count == 0 {
				return false
			}
			if !(c == ')' && peek(*stack) == "(") && !(c == ']' && peek(*stack) == "[") && !(c == '}' && peek(*stack) == "{") {
				return false
			}

			*stack = pop(*stack)
			count--
		}
	}
	return count == 0
}
