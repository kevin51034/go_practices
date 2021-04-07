func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    stack := make([]byte, 0)
    for i:=0; i<len(s); i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack = append(stack, s[i])
            continue
        }
        if len(stack) < 1 {
            return false
        }
        if s[i] == ')' {
            if stack[len(stack)-1] != '(' {
                return false
            }
        } else if s[i] == ']' {
            if stack[len(stack)-1] != '[' {
                return false
            }
        } else if s[i] == '}' {
            if stack[len(stack)-1] != '{' {
                return false
            }
        }
        stack = stack[:len(stack)-1]
    }
    return len(stack) == 0
}