package ps

import (
	"fmt"
)

func Ra(l *List) {
	if l.Len > 2 {
		l.Rotate()
	}
	fmt.Println("ra")
}

func Rb(l *List) {
	if l.Len > 2 {
		l.Rotate()
	}
	fmt.Println("rb")
}

func Rr(a, b *List) {
	if a.Len > 2 {
		a.Rotate()
	}
	if b.Len > 2 {
		b.Rotate()
	}
	fmt.Println("rr")
}

func Sa(l *List) {
	if l.Len > 2 {
		l.Swap()
	}
	fmt.Println("sa")
}

func Sb(l *List) {
	if l.Len > 2 {
		l.Swap()
	}
	fmt.Println("sb")
}

func Ss(a, b *List) {
	if a.Len > 2 {
		a.Swap()
	}
	if b.Len > 2 {
		b.Swap()
	}
	fmt.Println("ss")
}

func Pa(a, b *List) {
	if b.Len > 0 {
		b.Push(a)
	}
	fmt.Println("pa")
}

func Pb(a, b *List) {
	if a.Len > 0 {
		a.Push(b)
	}
	fmt.Println("pb")
}

func Rra(l *List) {
	if l.Len > 2 {
		l.ReverseRotate()
	}
	fmt.Println("rra")
}

func Rrb(l *List) {
	if l.Len > 2 {
		l.ReverseRotate()
	}
	fmt.Println("rrb")
}

func Rrr(a, b *List) {
	if a.Len > 2 {
		a.ReverseRotate()
	}
	if b.Len > 2 {
		b.ReverseRotate()
	}
	fmt.Println("rrr")
}
