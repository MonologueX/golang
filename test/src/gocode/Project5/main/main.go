// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
    for i:=1; i<len(nums); i++ {
        temp := nums[i-1]
		if temp == nums[i] {
			nums = append(nums[:i],nums[i+1:]...)
            i--
        }
	}
	return len(nums)
}
