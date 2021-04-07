// solution 1
// iterative

func reverseString(s []byte)  {
    if len(s)<=1 {
        return
    }
    for i:=0; i<len(s)/2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
}

// solution 2
// recursive

func reverseString(s []byte)  {
    result := make([]byte, 0)
    reverse(s, 0, &result)
    for i:=0; i<len(s); i++ {
        s[i] = result[i]
    }
}

func reverse(s[]byte, start int, result *[]byte) {
    if start == len(s) {
        return
    }
    reverse(s, start+1, result)
    *result = append(*result, s[start])
    return
}


// solution 2
// recursive but input only string without index
func reverseString(s []byte)  {
    if len(s) <= 1 {
        return
    }
    s[0], s[len(s)-1] = s[len(s)-1], s[0]
    reverseString(s[1:len(s)-1])
}