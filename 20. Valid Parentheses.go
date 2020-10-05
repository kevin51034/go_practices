func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    stack := make([]byte, 0)
    for i:=0; i<len(s); i++ {
        stack = append(stack, s[i])
        if s[i] == ')' || s[i] == '}' || s[i] == ']' {
            if len(stack) <= 1 {
                return false
            }
            switch c := stack[len(stack)-1]; c {
                case ')':
                    if stack[len(stack)-2]!='(' {
                        return false
                    }
                case '}':
                    if stack[len(stack)-2]!='{' {
                        return false
                    }
                case ']':
                    if stack[len(stack)-2]!='[' {
                        return false
                    }
            }
            stack = stack[:len(stack)-2]
            continue
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}