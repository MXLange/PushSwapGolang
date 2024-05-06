package ps

import (
	"math"
	"strconv"
	"strings"
)

type List struct {
	Len      int
	Arr      []int
	Max      int
	Min      int
	MaxIndex int
	MinIndex int
}

func Run(args []string) error {

	var numsArr []string

	if len(args) == 1 {
		numsArr = strings.Split(args[0], " ")
	} else {
		numsArr = args
	}

	aList := NewList(numsArr)
	bList := &List{
		Len: 0,
		Arr: make([]int, 0),
	}

	if aList.IsSorted() || aList.Len == 1 {
		return nil
	}
	if aList.Len == 2 {
		Sa(aList)
	} else if aList.Len == 3 {
		SortThree(aList)
	} else if aList.Len == 5 {
		SortFive(aList, bList)
	} else {
		RadixSort(aList, bList)
	}
	return nil
}

func NewList(args []string) *List {
	list := &List{
		Len:      len(args),
		Arr:      make([]int, len(args)),
		Min:      math.MaxInt64,
		Max:      math.MinInt64,
		MaxIndex: 0,
		MinIndex: 0,
	}
	for i, v := range args {
		num, _ := strconv.Atoi(v)
		list.Arr[i] = num
		if num > list.Max {
			list.Max = num
			list.MaxIndex = i
		}
		if num < list.Min {
			list.Min = num
			list.MinIndex = i
		}

	}
	return list
}
