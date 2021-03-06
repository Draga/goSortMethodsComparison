package sortMethods

func SelectSort(list []uint32) (checkCount uint64, swapCount uint64) {
	var minIndex int
	for i := 0; i < len(list) - 1; i++ {
		minIndex = i;
		for j := i + 1; j < len(list); j++ {
			checkCount++;
			if list[j] < list[minIndex] {
				minIndex = j;
			}
		}
		swapCount++;
		swap(&list[i], &list[minIndex])
	}

	return
}