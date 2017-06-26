package levenshtein

func Distance(s1 []byte, s2 []byte) int {

	mat := make([][]int, len(s1) + 1)
	for i := range mat {
		mat[i] = make([]int, len(s2) + 1)
	}

	for i := range mat {
		mat[i][0] = i
	}

	for j := range mat[0] {
		mat[0][j] = j
	}

	for j := 1; j <= len(s2); j++ {
		for i := 1; i <= len(s1); i++ {
			if s1[i-1] == s2[j-1] {
				mat[i][j] = mat[i-1][j-1]
			} else {
				min := mat[i-1][j]
				if mat[i][j-1] < min {
					min = mat[i][j-1]
				}
				if mat[i-1][j-1] < min {
					min = mat[i-1][j-1]
				}
				mat[i][j] = min + 1
			}
		}

	}
	return mat[len(s1)][len(s2)]
}
