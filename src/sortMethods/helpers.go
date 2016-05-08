package sortMethods

type SortFunc func ([]uint32) (uint, uint)

func swap(a *uint32, b *uint32) {
	tmp := *a;
	*a = *b;
	*b = tmp;
}
