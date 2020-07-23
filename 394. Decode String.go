func decodeString(s string) string {
    if len(s) == 0 {
        return s
    }
    stack := make([]byte, 0)
    for i:=0; i<len(s); i++ {
        switch s[i] {
            case ']':
            temp := make([]byte, 0)
            for len(stack)>0 && stack[len(stack)-1]!='[' {
                temp = append(temp, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            
            stack = stack[:len(stack)-1]
            
            index := 1
            for len(stack)>=index && stack[len(stack)-index]>='0' && stack[len(stack)-index]<='9'{
                index++
            }
            
            numberString := stack[len(stack)-index+1:]
            stack = stack[:len(stack)-index+1]
            number,_ := strconv.Atoi(string(numberString))
            
            for j:=0; j<number; j++ {
                for k:=len(temp)-1; k>=0; k-- {
                stack = append(stack, temp[k])
                }
            }
            
            default:
            stack = append(stack, s[i])
        }
    }
    return string(stack)
}