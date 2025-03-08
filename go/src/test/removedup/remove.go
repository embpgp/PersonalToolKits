package removeElement


func removeDuplicates(nums []int) int {
    l,r := 0, 0
    m := make(map[int]struct{})
    for ;r<len(nums);r++{
        if _,ok := m[nums[r]];ok{
            continue
        }
        nums[l] = nums[r]
        m[nums[l]] = struct{}{}
        l++
    }
    return l
}


func removeDuplicates2(nums []int) int {
    n := len(nums)
    if n <= 2 {
        return n
    }
    slow := 2
    for fast := 2;fast<len(nums);fast++{
        if nums[slow-2]!=nums[fast] {
            nums[slow] = nums[fast]
            slow++
        }
    }
    return slow
}
