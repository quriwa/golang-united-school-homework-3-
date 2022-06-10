package homework

func average(input [15]float32) (result float32) {
    sum := 0
	for _, a := range input {
	    sum = sum + a
	}
	average = sum / len(input)
	return average
}
