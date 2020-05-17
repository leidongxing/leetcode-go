package stack

func isValid(s string) bool {
	if s == "" {
		return true
	}

	var stack []uint8

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		}

		if len(stack) == 0 {
			return false
		}
		if s[i] == '}' {
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if s[i] == ']' {
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if s[i] == ')' {
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
