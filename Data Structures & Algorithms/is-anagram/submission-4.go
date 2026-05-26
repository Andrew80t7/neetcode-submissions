func isAnagram(s string, t string) bool {

	bytes_s := []byte(s)
	bytes_t := []byte(t)

	sort.Slice(bytes_s, func(i, j int) bool { return bytes_s[i] < bytes_s[j] })

	sort.Slice(bytes_t, func(i, j int) bool { return bytes_t[i] < bytes_t[j] })

	if bytes.Equal(bytes_s, bytes_t) {
		return true
	}
	return false

}