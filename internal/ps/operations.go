package ps

import "math"

func (l *List) IsSorted() bool {
	for i := 0; i < l.Len-1; i++ {
		if l.Arr[i] > l.Arr[i+1] {
			return false
		}
	}
	return true
}

func (l *List) SetMaxMin() {

	if l.Len == 0 {
		l.Max = 0
		l.Min = 0
		l.MaxIndex = 0
		l.MinIndex = 0
		return
	}

	l.Max = math.MinInt64
	l.Min = math.MaxInt64
	for i, v := range l.Arr {
		if v > l.Max {
			l.Max = v
			l.MaxIndex = i
		}
		if v < l.Min {
			l.Min = v
			l.MinIndex = i
		}
	}
}

func (l *List) Push(destList *List) {
	if l.Len == 0 {
		return
	}
	destCopy := make([]int, destList.Len)
	copy(destCopy, destList.Arr)
	destList.Arr = append([]int{l.Arr[0]}, destCopy...)
	destList.Len++
	l.Arr = l.Arr[1:]
	l.Len--

	l.SetMaxMin()
	destList.SetMaxMin()
}

func (l *List) Swap() {
	if l.Len < 2 {
		return
	}
	l.Arr[0], l.Arr[1] = l.Arr[1], l.Arr[0]
	l.SetMaxMin()
}

func (l *List) Rotate() {
	if l.Len < 2 {
		return
	}
	l.Arr = append(l.Arr[1:], l.Arr[0])
	l.SetMaxMin()
}

func (l *List) ReverseRotate() {
	if l.Len < 2 {
		return
	}
	l.Arr = append([]int{l.Arr[l.Len-1]}, l.Arr[:l.Len-1]...)
	l.SetMaxMin()
}
