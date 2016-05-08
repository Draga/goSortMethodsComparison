package sortMethods

func QuickSort(list []uint32) (checkCount uint, swapCount uint) {
	var start uint
	end := uint(len(list))

	quickSort(list, start, end, &checkCount, &swapCount)

	return
}

func quickSort(list []uint32, start uint, end uint, checkCount *uint, swapCount *uint) {
	var pivot uint

	*checkCount++;
	if (end == start) {
		return
	}
	pivot = getPivot(&list, start, end, swapCount);
	*checkCount++;
	if (start < pivot) {
		quickSort(list, start, pivot - 1, checkCount, swapCount)
	}
	*checkCount++;
	if (pivot < end) {
		quickSort(list, pivot + 1, end, checkCount, swapCount)
	}

	return
}

func getPivot(list *[]uint32, start uint, end uint, swapCount *uint) uint {
	i := start
	j := end
	pivot := (*list)[start]

	for {
		for {
			i++;
			if i >= j || (*list)[i] >= pivot {
				break
			}
		}

		for {
			j--;
			if j <= start || (*list)[j] <= pivot {
				break
			}
		}

		if (i < j) {
			*swapCount++;
			swap(&(*list)[i], &(*list)[j])
		}

		if i >= j {
			break
		}
	}
	*swapCount++;
	swap(&(*list)[start], &(*list)[j]);

	return j;
}