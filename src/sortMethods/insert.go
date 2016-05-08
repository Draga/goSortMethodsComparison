package sortMethods

func InsertSort(list []uint32) (checkCount uint, swapCount uint) {
	var next uint32
	var j int

	for i := 0; i < len(list); i++ {
		next = list[i];
		for j = i - 1; j >= 0 && next < list[j]; j-- {
			swapCount++; checkCount++;
			list[j + 1] = list[j];
		}
		swapCount++;
		list[j + 1] = next;
	}

	return
}