func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j, count:=0, 0, 0
    if len(nums2) == 0 || len(nums1) == 0{
        return
    } 
    for count<m && j<n {
        //fmt.Println(nums1)
        if nums1[i] <= nums2[j] {
            i++
            count++
        } else {
            copy(nums1[i+1:], nums1[i:])
            nums1[i] = nums2[j]
            i++
            j++
        }
    }
    for i<m+n && j<n {
        nums1[i] = nums2[j]
        i++
        j++
    }
}