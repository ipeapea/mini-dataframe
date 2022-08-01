package series

import (
	"reflect"
	"sort"
)

type Comparable interface {
	Compare(comparable Comparable) int
}

type Aggregation interface {
	Init(element Element)
	Reduce(element Element)
	Collect() interface{}
}

func CheckObjComparable(obj interface{}) bool {
	return reflect.TypeOf(obj).Implements(reflect.TypeOf((*Comparable)(nil)).Elem())
}

// Order returns the indexes for sorting a Series. NaN elements are pushed to the
// end by order of appearance.
func (s Series) Order(reverse bool) []int {
	var ie indexedElements
	var nasIdx []int
	for i := 0; i < s.Len(); i++ {
		e := s.elements.Elem(i)
		if e.Null() {
			nasIdx = append(nasIdx, i)
		} else {
			ie = append(ie, indexedElement{i, e})
		}
	}
	var srt sort.Interface
	srt = ie
	if reverse {
		srt = sort.Reverse(srt)
	}
	sort.Stable(srt)
	var ret []int
	for _, e := range ie {
		ret = append(ret, e.index)
	}
	return append(ret, nasIdx...)
}

type indexedElement struct {
	index   int
	element Element
}

type indexedElements []indexedElement

func (e indexedElements) Len() int           { return len(e) }
func (e indexedElements) Less(i, j int) bool { return false }
func (e indexedElements) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
