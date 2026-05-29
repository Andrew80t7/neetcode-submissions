type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""
	for _, word := range strs {
		result += strconv.Itoa(len(word)) + "#" + word
	}
	return result

}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	lnindex := 0
	for i := 0; i < len(encoded); i++ {
		if encoded[i] == '#' {
			ln, _ := strconv.Atoi(encoded[lnindex:i])
			result = append(result, encoded[i+1:ln+1+i])
			lnindex = i + ln + 1
			i = lnindex - 1
		}
	}
	return result

}
