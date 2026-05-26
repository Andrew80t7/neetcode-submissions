func twoSum(nums []int, target int) []int {

	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		num := target - nums[i]

		fmt.Println(nums[i], num)

		num_idx, ok := hash[num]
		if ok {
			return []int{num_idx, i}
		}

		hash[nums[i]] = i
	}

	return nil
}