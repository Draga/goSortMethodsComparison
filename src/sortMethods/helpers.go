package sortMethods

func swap(a *uint32, b *uint32) {
	tmp := *a;
	*a = *b;
	*b = tmp;
}
