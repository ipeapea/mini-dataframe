package series

import (
	"fmt"
	"reflect"
	"strings"
)

type commonElement struct {
	value      interface{}
	nan        bool
	typ        reflect.Type
	comparable bool
}

type commonElements struct {
	elements []commonElement
	typ      reflect.Type
}

type commonElementProvider struct {
	typ reflect.Type
}

// force intElement struct to implement Element interface
var _ Element = (*commonElement)(nil)

func (e commonElement) Get() interface{} {
	if e.nan {
		return nil
	}
	return e.value
}
func (e commonElement) Clone() Element {
	newObj := reflect.New(e.typ)
	newObj.Elem().Set(reflect.ValueOf(e.value))
	return &commonElement{
		value:      newObj.Elem().Interface(),
		typ:        e.typ,
		comparable: e.comparable,
	}
}

func (e *commonElement) Set(val interface{}) {
	if val == nil {
		e.value, e.nan = nil, true
		return
	}
	e.value, e.nan = val, false
}

func (e commonElement) Null() bool {
	return e.nan
}
func (e commonElement) Type() reflect.Type {
	return e.typ
}

func (e commonElement) Compare(element Element) CompareResult {
	if !e.comparable {
		return CompareResultTypeErr
	}
	in, ok := element.(*commonElement)
	if !ok {
		return CompareResultTypeErr
	}
	if e.nan {
		return CompareResultSelfNull
	}
	if in.nan {
		return CompareResultParamNull
	}
	self := e.value.(Comparable)
	param := e.value.(Comparable)
	ret := self.Compare(param)
	if ret < 0 {
		return CompareResultLess
	}
	if ret == 0 {
		return CompareResultEqual
	}
	return CompareResultGreater
}

func (e commonElement) String() string {
	if e.nan {
		return Null
	}
	return fmt.Sprintf("%v", e.value)
}

func (e commonElements) Elem(idx int) Element {
	return &e.elements[idx]
}
func (e commonElements) Len() int {
	return len(e.elements)
}

func (e commonElements) Type() reflect.Type {
	return e.typ
}

func (e commonElements) String() string {
	var ret []string
	ret = append(ret, "[")
	for _, item := range e.elements {
		//var stringer fmt.Stringer = fmt.Stringer(nil)
		//ok := reflect.TypeOf(item).Implements(reflect.TypeOf(stringer))
		//if ok {
		//
		//}
		//itemStr := fmt.Sprintf("%v", item.value)
		ret = append(ret, fmt.Sprintf("%v", item.value))
	}
	return strings.Join(ret, ",")
}

func (f commonElementProvider) New() Element {
	newObj := reflect.New(f.typ)
	return &commonElement{
		value:      newObj.Elem().Interface(),
		typ:        f.typ,
		comparable: CheckObjComparable(newObj.Elem().Interface()),
	}
}
