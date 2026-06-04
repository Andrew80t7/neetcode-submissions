func longestConsecutive(nums []int) int {
	box := make(map[int]bool)
	for _, num := range nums {
		box[num] = true
	}
	mx := 0

	for _, num := range nums {

		if !box[num-1] {
			currentNum := num
			currentLen := 1

			for box[currentNum+1] {
				currentLen++
				currentNum++
			}

			if currentLen > mx {
				mx = currentLen
			}

		}
	}
	return mx
}