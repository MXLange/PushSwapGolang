package ps

func SortFive(a, b *List) {

	if a.IsSorted() {
		return
	}

	for a.Len > 3 {
		sendMinToB(a, b)
	}
	SortThree(a)
	returnValuesToA(a, b)
}

func sendMinToB(a *List, b *List) {
	putMinOnTop(a)
	Pb(a, b)
}

func putMinOnTop(a *List) {
	medium := a.Len / 2

	if a.MinIndex <= medium {
		for a.MinIndex != 0 {
			Ra(a)
		}
	} else {
		for a.MinIndex != 0 {
			Rra(a)
		}
	}
}

func returnValuesToA(a *List, b *List) {
	for b.Len > 0 {
		Pa(a, b)
	}
}
