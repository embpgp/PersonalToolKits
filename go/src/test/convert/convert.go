package convert

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	//一个循环是2r-2个元素，比如3行，4个；4行，6个
	//直接输出

	/* i 从0 -> n/(2r-2)
	0                4 i*(2r-2)              8
	1  3 r-i     5  i     7 3r-i             9    11
	2 (r-1)           6                      10

	*/
	ret := make([]byte, 0)
	r := numRows
	n := len(s)
	L := 2*r - 2
	for i := 0; i < r; i++ {
		for j := 0; j+i < n; j += L {
			ret = append(ret, s[j+i]) //当前元素为j+i，那么根据对称性 另一个为j+L-i，默认i和L-i是一对
			if i > 0 && i < r-1 && j+L-i < n {
				ret = append(ret, s[j+L-i])
			}
		}
	}
	return string(ret)
}
