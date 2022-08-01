package series

import (
	"reflect"
	"strconv"
	"strings"
)

var intType = reflect.TypeOf(int64(0))

type intElement struct {
	e   int64
	nan bool
}

type intElements []intElement

type intElementFactory struct{}

// force intElement struct to implement Element interface
var _ Element = (*intElement)(nil)

var localIntElementFactory intElementFactory

func init() {
	intVar := []interface{}{int8(0), int16(0), int32(0), int64(0), int(0), uint8(0), uint16(0), uint32(0),
		uint64(0)}
	for _, item := range intVar {
		typTmp := reflect.TypeOf(item)
		localElementsFactory[typTmp] = localIntElementFactory
	}
}

func (e *intElement) Get() interface{} {
	return e.e
}

func (e intElement) Clone() Element {
	return &intElement{e.e, e.nan}
}

func (e intElement) Compare(element Element) CompareResult {
	in, ok := element.(*intElement)
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

func (e intElement) Null() bool {
	return e.nan
}

func (e intElement) Type() reflect.Type {
	return intType
}

func (e intElement) String() string {
	if e.nan {
		return Null
	}
	return strconv.FormatInt(e.e, 10)
}

func (e *intElement) Set(value interface{}) {
	if value == nil {
		e.e, e.nan = 0, true
		return
	}
	e.nan = false
	switch valTmp := value.(type) {
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64:
		val := reflect.ValueOf(value).Convert(intType)
		reflect.ValueOf(&e.e).Elem().Set(val)
	case string:
		tmpNum, err := strconv.ParseInt(valTmp, 10, 64)
		if err != nil {
			e.e, e.nan = 0, true
			return
		}
		e.e = tmpNum
	case Element:
		e.nan = valTmp.Null()
		if !e.nan {
			e.Set(valTmp.Get())
		} else {
			e.nan = true
		}
	}
}

func (e intElements) Len() int {
	return len(e)
}

func (e intElements) Elem(idx int) Element {
	return &e[idx]
}

func (e intElements) Type() reflect.Type {
	return intType
}

func (e intElements) String() string {
	var ret []string
	for idx, item := range e {
		if idx == 0 {
			ret = append(ret, "["+item.String())
			continue
		}
		ret = append(ret, ", "+item.String())
	}
	ret = append(ret, "]")
	return strings.Join(ret, "")
}

func (f intElementFactory) new() Element {
	return &intElement{0, true}
}

func (f intElementFactory) newElements(num int) Elements {
	tmp := make(intElements, num)
	for idx := range tmp {
		tmp[idx].nan = true
	}
	return tmp
}
