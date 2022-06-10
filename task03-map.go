package homework

func sortMapValues(input map[int]string) (result []string) {
	keys := []int{}
	values := []string{}
	for key := range input{
	    keys = append(keys, key)
	}
	for i := 0; i < len(keys)-1; i++ {
		for j := 0; j < len(keys)-i-1; j++ {
			if keys[j] > keys[j+1] {
				keys[j], keys[j+1] = keys[j+1], keys[j]
			}
		}
	}
    for _, k := range keys {
		values = append(values, basket[k])
	}

	return values
}
