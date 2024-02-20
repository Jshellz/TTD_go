package integers

const repeatCount = 5

func Add(x, y int) int {
	return x + y
}

func Repeat(word string) string {

	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += word
	}

	return repeated
}
