package ps

func SortThree(a *List) {

	if a.IsSorted() {
		return
	}

	if (a.Arr[0] > a.Arr[1]) && (a.Arr[1] < a.Arr[2]) && (a.Arr[0] < a.Arr[2]) {
		Sa(a)
	} else if (a.Arr[0] > a.Arr[1]) && (a.Arr[1] < a.Arr[2]) && (a.Arr[2] < a.Arr[0]) {
		Ra(a)
	} else if (a.Arr[0] < a.Arr[1]) && (a.Arr[1] > a.Arr[2]) && (a.Arr[2] < a.Arr[0]) {
		Rra(a)
	} else if (a.Arr[0] > a.Arr[1]) && (a.Arr[1] > a.Arr[2]) && (a.Arr[2] < a.Arr[0]) {
		Sa(a)
		Rra(a)
	} else if (a.Arr[0] < a.Arr[1]) && (a.Arr[1] > a.Arr[2]) && (a.Arr[2] > a.Arr[0]) {
		Sa(a)
		Ra(a)
	}
}
