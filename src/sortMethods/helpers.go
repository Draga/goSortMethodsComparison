package sortMethods

type SortFunc func (list []uint32) (checkCount uint64, swapCount uint64)

func swap(a *uint32, b *uint32) {
	tmp := *a;
	*a = *b;
	*b = tmp;
}
