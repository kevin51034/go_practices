func evalRPN(tokens []string) int {
    if len(tokens) == 0 {
        return 0
    }   
    stack := make([]int, 0)
    for i:=0; i<len(tokens); i++ {
        switch tokens[i] {
            case "+", "-", "*", "/":
            if len(stack)<2 {
                return -1
            }
            n2 := stack[len(stack) -1]
            n1 := stack[len(stack) -2]
            stack = stack[:len(stack)-2]
            var result int
            switch tokens[i] {
                case "+":
                result = n1 + n2
                case "-":
                result = n1 - n2
                case "*":
                result = n1 * n2
                case "/":
                result = n1 / n2
            }
            stack = append(stack, result)
            
            default :
            number,_ := strconv.Atoi(tokens[i])
            stack = append(stack, number)
        }
    }
    return stack[0]
}