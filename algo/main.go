package algo

func Exclude(arr1, arr2 []int64) (left, right []int64) {
	var min, max *[]int64

	if len(arr1) < len(arr2) {
		min = &arr1
		max = &arr2
	} else {
		min = &arr2
		max = &arr1
	}

	var t []int64

	for i := 0; i < len(*min); i++ {
		flag := false
		for j := 0; j < len(*max); j++ {
			if (*min)[i] == (*max)[j] {
				flag = true
				t = append(t, (*min)[i])
				break
			}
		}
		if !flag {
			left = append(left, (*min)[i])
		}
	}

	for i := 0; i < len(*max); i++ {
		flag := false
		for j := 0; j < len(t); j++ {
			if (*max)[i] == t[j] {
				flag = true
				break
			}
		}
		if !flag {
			right = append(right, (*max)[i])
		}
	}

	return
}
