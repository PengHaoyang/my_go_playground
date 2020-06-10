package main

func trailingZeros(n int64) int64 {
	// write your code here, try to do it without arithmetic operators.
	cf, ct := do_acc_count_of_five_and_two_in_num(n, 0, 0)
	if cf > ct {
		return ct
	}
	return cf
}

func do_acc_count_of_five_and_two_in_num(n, pcf, pct int64) (int64, int64) {
	if n == 0 {
		return pcf, pct
	}
	tempn := n
	for n%5 == 0 {
		pcf++
		n /= 5
	}
	for n%2 == 0 {
		pct++
		n /= 2
	}
	return do_acc_count_of_five_and_two_in_num(tempn-1, pcf, pct)
}
