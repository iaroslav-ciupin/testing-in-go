package code 

func answer() (int, error) {
	a := 100
	a -= 58
	_ = make([]int, 0, a)
	if a > 0 {
		return a, nil	
	}
	return 0, nil
}

func f(i int) int {
	return i + 1
}