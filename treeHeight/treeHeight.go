package treeHeight

func TreeHeight(n int, nums []int) int {
	m := make(map[int][]int, 0)
	for i := 0; i < n; i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}

	quenue := make([]int, 0)
	quenue = append(quenue, m[-1]...)
	res := 0
	for len(quenue) > 0 {
		l := len(quenue)
		for i := 0; i < l; i++ {
			if value, ok := m[quenue[i]]; ok {
				quenue = append(quenue, value...)
			}
		}

		res++
		quenue = quenue[l:]
	}

	return res

}
