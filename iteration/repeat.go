package iteration

func Repeat(str string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += str
	}

	return
}
