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

func removeDuplicates1(nums []int) int {
    slow := 1
    for fast:=1;fast<len(nums);fast++ {
        if nums[fast]!=nums[fast-1] {
		nums[slow] =  nums[fast]
		slow++		
	}
    }
    return slow
}


func removeDuplicates2(nums []int,keep int) int {
    n := len(nums)
    if n <= keep {
        return n
    }
    slow := keep
    for fast := keep;fast<len(nums);fast++{
        if nums[slow-keep]!=nums[fast] {
            nums[slow] = nums[fast]
            slow++
        }
    }
    return slow
}
