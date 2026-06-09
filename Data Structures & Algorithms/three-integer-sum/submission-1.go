func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i, j := k+1, len(nums)-1

		for i < j {
			sum := nums[k] + nums[i] + nums[j]

			if sum < 0 {
				i++
			} else if sum > 0 {
				j--
			} else {
				result = append(result, []int{nums[k], nums[i], nums[j]})
				i++
				j--

				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}

	return result
}