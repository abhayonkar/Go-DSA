func nextPermutation(nums []int)  {
    index := -1
    n := len(nums)

    for i := n-2; i >= 0; i-- {
        if (nums[i] < nums[i+1]) {
            index = i
            break
        }
    }

    if index == -1 {
        // If index is -1, reverse the whole array
        for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
            nums[i], nums[j] = nums[j], nums[i]
        }
        return
    }

    if index != -1{
        for i:=n-1; i > index; i-- {
            if nums[i] > nums[index] {
                nums[i], nums[index] = nums[index], nums[i]
            break
            }
        }
    }

    for i, j := index + 1 , n-1; i < j; i, j = i+1, j-1 {
            nums[i], nums[j] = nums[j], nums[i]
        }

    return
}
