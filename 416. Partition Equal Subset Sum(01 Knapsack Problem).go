func canPartition(nums []int) bool {    

    sum := 0
    for _,v := range nums {
        sum += v
    }
    if sum%2 != 0 {
        return false
    }
    halfSum := sum/2
    f := make([]bool, halfSum+1)
    f[0] = true
    for i:=0; i<len(nums); i++ {
        for j:=halfSum; j>=nums[i]; j-- {
            f[j] = f[j] || f[j-nums[i]]
        }
        fmt.Println(f)
    }
    return f[halfSum]
}