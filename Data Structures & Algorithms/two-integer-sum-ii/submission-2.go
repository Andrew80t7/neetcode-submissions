func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		result := numbers[i] + numbers[j]
		if result == target {
			return []int{i + 1, j + 1}
		}
		if result > target {
			j--
		}
		if result < target {
			i++
		}
	}
	return []int{}

}
