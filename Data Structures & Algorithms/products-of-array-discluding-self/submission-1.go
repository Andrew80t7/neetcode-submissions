func productExceptSelf(nums []int) []int {
	length := len(nums)
	output := make([]int, length)

	output[0] = 1
	for i := 1; i < length; i++ {
		output[i] = output[i-1] * nums[i-1]
	}
	right := 1
	for i := length - 1; i >= 0; i-- {
		output[i] = output[i] * right
		right = right * nums[i]
	}

	return output
}