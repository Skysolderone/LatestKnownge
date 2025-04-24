package main

func numRabbits(answers []int) int {
	count := 0
	data := make(map[int]int, 0)
	n := len(answers)
	for _, v := range answers {
		data[v]++
	}
	for key, val := range data {
		if val == n {
			count += key + 1
			continue
		}
		if key == 0 {
			count += val
			continue
		}
		if val == 1 {
			count += key + 1
			continue
		}
		if key%2 == 0 && val%2 != 0 {
			count += (key * (val - 1)) + key + 1
			continue
		}

		if val%2 != 0 {
			count += 2 * 2
			continue
		}

		if val > 1 {
			count += key + 1
		}
	}
	return count
}

func numRabbits2(answers []int) int {
	count := 0
	data := make(map[int]int)

	for _, v := range answers {
		data[v]++
	}

	for key, val := range data {
		groupSize := key + 1
		// 用整数除法代替 math.Ceil
		groupCount := (val + key) / groupSize
		count += groupCount * groupSize
	}

	return count
}
