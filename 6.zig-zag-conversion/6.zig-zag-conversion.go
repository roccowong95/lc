/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */
func convert(s string, numRows int) string {
	// numRows = 3
	// i:   0 1 2 3 4 5 6 7 8 9
	// idx: 0 1 2 1 0 1 2 1 0 1
	if numRows == 1 {
		return s
	}

	caches := make([][]byte, numRows)
	flag := -1
	idx := 0
	for i := range s {
		caches[idx] = append(caches[idx], s[i])
		if idx == 0 || idx == numRows-1 {
			flag=-flag
		}
		idx+=flag
	}
	var ret []byte
	for i := range caches {
		ret = append(ret, caches[i]...)
	}
	return string(ret)
}

