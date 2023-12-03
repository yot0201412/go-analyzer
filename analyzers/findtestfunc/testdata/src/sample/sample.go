package sample

func TestSample(i int) string { // want HogeSample:`findTestFuncFact`

	if i%2 == 0 {
		return "even"
	}
	return "odd"
}

func Fuga(i int) string { //  OK

	if i%2 == 0 {
		return "even"
	}
	return "odd"
}
