package main

import (
"fmt"
"strings"
)


func main() {
	/*s := "Python is an easy to learn powerful programming language It has efficient high-level data structures and a simple but effective approach to objectoriented programming Python elegant syntax and dynamic typing"
	t := "Python is an easy to learn powerful programming language"
	fmt.Println(missingWords(s, t))*/
	/*arr := make([]int32, 4)
	arr[0] = 1
	arr[1] = 5
	arr[2] = 10
	arr[3] = 3
	fmt.Println(getMostVisited(10, arr))*/
	fmt.Println(longestPalindromeSubseq("bandaba"))
	//fmt.Println(longestPalindrome(2, "ba"))

}


///

func longestPalindrome(n int32, s string) int32 {
    // Write your code here
    f := make([][]int32, n)

	for i := range f {
        f[i] = make([]int32, n)
    }
    return dpfunc2(s, 0, n-1, f)
}

func dpfunc2(s string, lindex, rindex int32, f [][]int32) int32 {
	fmt.Println(lindex, rindex)
    if lindex > rindex {
        return 0
    }
    if f[lindex][rindex] != 0 {
        return f[lindex][rindex]
    }
    if lindex == rindex {
        f[lindex][rindex] = 1
        return f[lindex][rindex]
    }
    if s[lindex] == s[rindex] {
        f[lindex][rindex] = 2 + dpfunc2(s, lindex-1, rindex-1, f)
    } else {
        l := dpfunc2(s, lindex, rindex-1, f)
        r := dpfunc2(s, lindex+1, rindex, f)
        if l>r {
            f[lindex][rindex] = l
        } else {
            f[lindex][rindex] = r
        }
    }
    return f[lindex][rindex]
}

//

func longestPalindromeSubseq(s string) int32 {
    dp := make([][]int32, len(s))
    for i := range dp {
        dp[i] = make([]int32, len(s))
    }
    return dpfunc(s, 0, int32(len(s)-1), dp)
}


func dpfunc(s string, lindex, rindex int32, f [][]int32) int32 {
    if lindex > rindex {
        return 0
    }
    if f[lindex][rindex] != 0 {
        return f[lindex][rindex]
    }
    if lindex == rindex {
        f[lindex][rindex] = 1
        return f[lindex][rindex]
    }
    if s[lindex] == s[rindex] {
        f[lindex][rindex] = 2 + dpfunc(s, lindex-1, rindex-1, f)
    } else {
        l := dpfunc(s, lindex, rindex-1, f)
        r := dpfunc(s, lindex+1, rindex, f)
        if l>r {
            f[lindex][rindex] = l
        } else {
            f[lindex][rindex] = r
        }
    }
    return f[lindex][rindex]
}

func helper(s string, lBound, rBound int, dp [][]int) int {
    if lBound > rBound {
        return 0
    }
    if dp[lBound][rBound] != 0 {
        return dp[lBound][rBound]
    }
    if lBound == rBound {
        dp[lBound][rBound] = 1
        return 1
    }
    if s[lBound] == s[rBound] {
        dp[lBound][rBound] = 2 + helper(s, lBound+1, rBound-1, dp)
    } else {
        lMax := helper(s, lBound, rBound-1, dp)
        rMax := helper(s, lBound+1, rBound, dp)
        if lMax > rMax {
            dp[lBound][rBound] = lMax
        } else {
            dp[lBound][rBound] = rMax
        }
    }
    return dp[lBound][rBound]
}

func missingWords(s string, t string) []string {
    // Write your code here
    // Write your code here
    result := make([]string, 0)
    if s==t {
        return result
    }
    sSlice := strings.Split(s, " ")
    tSlice := strings.Split(t, " ")
    for i:=0; i<len(sSlice); i++ {
        if len(tSlice)<=0 || sSlice[i] != tSlice[0] {
            result = append(result, sSlice[i])
            continue
        } else {
            tSlice = tSlice[1:]
        }
    }
    return result
}

func getMostVisited(n int32, sprints []int32) int32 {
    // Write your code here
    f := make([]int32, n+1)
	var first int32
    first = sprints[0]
    for i:=1 ;i<len(sprints); i++ {
        second := sprints[i]
		fmt.Println(first, second)

        if first < second {
            for j:=first; j<=second; j++ {
                f[j]++
            }
        } else if first > second {
            for j:=first; j>=second; j-- {
                f[j]++
            }
        }
        first = second
    }
	var result int32
	fmt.Println(f)
    for i:=1; i<len(f); i++ {
		fmt.Println(result)
		fmt.Println(f[i])

        if f[i] > result {
            result = int32(i)
        }
    }

    return result
}
