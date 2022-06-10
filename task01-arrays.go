package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	for _, a := range input {
		sum += a
	}
	return sum / float32(len(input))
}