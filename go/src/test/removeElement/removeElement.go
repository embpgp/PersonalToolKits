package removeElement


func removeElement(nums []int, val int) int {
    left,right := 0,0
    for ;right<len(nums);right++ {
        if nums[right] == val {
            continue
        }
        nums[left] = nums[right]
        left++
    }
    return left

}

func bak(nums []int, val int) int {
    cnt := 0
    m := make(map[int]int)
    for i := range nums {
        if nums[i] == val {
            continue
        }
        m[nums[i]]++
    }
    idx:=0
    for k,v:=range m {
        for i:=0;i<v;i++ {
            nums[idx] = k
            cnt++
            idx++
        }
    }
    return cnt
}
