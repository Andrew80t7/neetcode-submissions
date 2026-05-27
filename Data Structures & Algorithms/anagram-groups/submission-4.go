
func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)

	for _, word := range strs {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		hash[string(runes)] = append(hash[string(runes)], word)
	}

	result := make([][]string, 0, len(hash))
    
	for _, group := range hash {
		result = append(result, group)
	}

	return result
}
