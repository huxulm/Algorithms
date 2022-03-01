package kmp

type KMP interface {
	Search(s string) int
}

type kmp struct {
	R   int     // the radix
	m   int     // the length of pattern
	dfa [][]int // the KMP automoton
}

func Build(pat string) KMP {
	R, m := 256, len(pat)

	// build dfa from pattern
	dfa := make([][]int, R)
	for i := range dfa {
		dfa[i] = make([]int, m)
	}

	// base case
	dfa[pat[0]][0] = 1
	// x current restart state
	// j current state
	for x, j := 0, 1; j < m; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x] // Copy mismatch cases.
		}
		dfa[pat[j]][j] = j + 1 // Set match case
		x = dfa[pat[j]][x]     // Update restart state
	}

	return &kmp{R: R, m: m, dfa: dfa}
}

func (k *kmp) Search(s string) int {
	n := len(s)
	var i, j int

	for ; i < n && j < k.m; i++ {
		j = k.dfa[s[i]][j]
	}
	if j == k.m { // Found
		return i - j
	}
	// Not found
	return -1
}

func Search(s, p string) int {
	n, m := len(s), len(p)
	pi := make([]int, m)
	// build pi
	for j, i := 0, 1; i < m; i++ {
		for j > 0 && p[i] != p[j] {
			j = pi[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		pi[i] = j
	}

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && s[i] != p[j] {
			j = pi[j-1]
		}
		if s[i] == p[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
