var	letterMap = []string{
	" ",    //0
	"",     //1
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}

// solution 1
// DFS

func letterCombinations(digits string) []string {
    result := []string{}
    if len(digits) == 0 {
        return result
    }
    findCombinations(&digits, &result, 0, "")
    return result
}

func findCombinations(digits *string, result *[]string, index int, s string) {
    if index == len(*digits) {
        *result = append(*result, s)
        return
    }
    number := (*digits)[index]
    letter := letterMap[number - '0']
    for i:=0; i<len(letter); i++ {
        findCombinations(digits, result, index+1, s+string(letter[i]))
    }
    return
}

