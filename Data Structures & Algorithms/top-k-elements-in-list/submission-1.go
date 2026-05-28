func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)

	for _, num := range nums {
		hashmap[num]++
	}

	arr := []int{}
	for key, _ := range hashmap {
		arr = append(arr, key)
	}
	sort.Slice(arr, func(i, j int) bool {
		return hashmap[arr[i]] < hashmap[arr[j]]
	})

	result := []int{}

	for i := len(arr) - 1; i >= len(arr)-k; i-- {
		result = append(result, arr[i])
	}

	return result
}