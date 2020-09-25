type NumArray struct {
    sum []int
}


func Constructor(nums []int) NumArray {
    sum := make([]int, len(nums))
    for i:=0; i<len(nums); i++ {
        if i == 0 {
            sum[i] = nums[i]
        } else {
            sum[i] = sum[i-1] + nums[i]
        }
    }
    return NumArray{sum}
}


func (this *NumArray) SumRange(i int, j int) int {
    if i == 0 {
        return this.sum[j]
    }
    return this.sum[j] - this.sum[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */