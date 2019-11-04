/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
func myAtoi(str string) int {
	idx, l, flag, result, max := 0, len(str), 1, 0, 0

	for idx<l && str[idx]==' ' {
		idx++
	}

	if idx >= l {
		return 0
	}

	if str[idx] == '-' {
		flag = -1
		idx++
	} else if str[idx] == '+' {
		idx++
	}

	if flag > 0 {
		max = math.MaxInt32
	} else {
		max = -math.MinInt32
	}

	tenth := max / 10
	n := 0

	for ; idx<l && str[idx] >= '0' && str[idx] <= '9'; idx++ {
		n = int(str[idx]) - '0'
		if result > tenth || result * 10 + n > max {
			return flag * max;
		}
		result = result * 10 + n
	}
	return flag * result
}

