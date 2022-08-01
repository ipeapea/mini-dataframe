package series

import (
	"fmt"
	"reflect"
	"strings"
)

var stringTyp = reflect.TypeOf("")

type stringElement struct {
	e   string
	nan bool
}

type stringElements []stringElement

type stringElementFactory struct{}

// force stringElement struct to implement Element interface
var _ Element = (*stringElement)(nil)

var localStringElementFactory stringElementFactory

func init() {
	localElementsFactory[stringTyp] = localStringElementFactory
}

// Set
func (e *stringElement) Set(value interface{}) {
	if value == nil {
		e.e, e.nan = "", true
		return
	}
	if tmpVal, ok := value.(string); ok {
		e.e, e.nan = tmpVal, false
		return
	}
	e.nan = false
	switch val := value.(type) {
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		e.e = fmt.Sprintf("%v", value)
	case Element:
		e.nan = val.Null()
		if !e.nan {
			e.Set(val.Get())
		} else {
			e.nan = true
		}
	}
}

func (e stringElement) Get() interface{} {
	return e.e
}

func (e stringElement) Clone() Element {
	return &stringElement{e.e, e.nan}
}

func (e stringElement) Null() bool {
	return e.nan
}

func (e stringElement) Type() reflect.Type {
	return stringTyp
}

func (e stringElement) Compare(element Element) CompareResult {
	in, ok := element.(*stringElement)
	if !ok {
		return CompareResultTypeErr
	}
	if e.nan {
		return CompareResultSelfNull
	}
	if in.nan {
		return CompareResultParamNull
	}
	if e.e == in.e {
		return CompareResultEqual
	}
	if e.e > in.e {
		return CompareResultGreater
	}
	return CompareResultLess
}

func (e stringElement) String() string {
	if e.nan {
		return Null
	}
	return e.e
}

func (e stringElements) Len() int {
	return len(e)
}

func (e stringElements) Elem(idx int) Element {
	return &e[idx]
}

func (e stringElements) Type() reflect.Type {
	return stringTyp
}

func (e stringElements) String() string {
	var ret []string
	ret = append(ret, "[")
	for _, item := range e {
		ret = append(ret, item.String())
	}
	return strings.Join(ret, ",")
}

func (f stringElementFactory) new() Element {
	return &stringElement{"", true}
}

func (f stringElementFactory) newElements(num int) Elements {
	tmp := make(stringElements, num)
	for idx := range tmp {
		tmp[idx].nan = true
	}
	return tmp
}
func (f stringElementFactory) append(elements Elements, ele ...Element) Elements {
	elementsTmp := elements.(stringElements)
	appendEle := make([]stringElement, 0, len(ele))
	for _, item := range ele {
		eleTmp := item.(*stringElement)
		appendEle = append(appendEle, *eleTmp)
	}
	elementsTmp = append(elementsTmp, appendEle...)
	return elementsTmp
}
func (f stringElementFactory) subSet(elements Elements, start, end int) Elements {
	lementsTmp := elements.(stringElements)
	return lementsTmp[start:end]
}
