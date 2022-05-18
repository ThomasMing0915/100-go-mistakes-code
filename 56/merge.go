package main

func merge(s []int, middle int) {
	slLen := len(s[:middle])
	srLen := len(s[middle:])
	sl := make([]int, slLen)
	sr := make([]int, srLen)

	copy(sl, s[:middle])
	copy(sr, s[middle:])

	var i, il, ir int
	for il < slLen && ir < srLen {
		if sl[il] <= sr[ir] {
			s[i] = sl[il]
			il++
		} else {
			s[i] = sr[ir]
			ir++
		}
		i++
	}

	for il < slLen {
		s[i] = sl[il]
		i++
		il++
	}

	for ir < srLen {
		s[i] = sr[ir]
		i++
		ir++
	}
}
