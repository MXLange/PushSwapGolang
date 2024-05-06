package ps

func (l *List) GetMaxBits() int {
	max := l.Max
	bits := 0
	for max>>bits != 0 {
		bits++
	}
	return bits
}

func RadixSort(a, b *List) {
	i := 0
	maxBits := a.GetMaxBits()
	for i < maxBits && !a.IsSorted() {
		j := 0
		for j < a.Len {
			if a.Arr[0]>>i&1 == 0 {
				Pb(a, b)
			} else {
				Ra(a)
			}
			j++
		}
		for b.Len > 0 {
			Pa(a, b)
		}
		i++
	}
}
