package series

import (
	"fmt"
	"reflect"
	"strings"
)

// Series is a data structure designed for operating on arrays of elements that
// should comply with a certain type structure. They are flexible enough that can
// be transformed to other Series types and account for missing or non valid
// elements. Most of the power of Series resides on the ability to compare and
// subset Series of different types.
type Series struct {
	Name     string       // The name of the series
	elements Elements     // The values of the elements
	t        reflect.Type // The type of the series
	err      error
}

// New is the generic Series constructor
func New(values interface{}, name string) Series {
	ret := Series{
		Name: name,
	}
	if values == nil {
		ret.err = fmt.Errorf("values is nil ")
		return ret
	}
	typ, val := reflect.TypeOf(values), reflect.ValueOf(values)
	if typ.Kind() != reflect.Slice {
		ret.err = fmt.Errorf("values is not slice ")
		return ret
	}
	if err := checkSliceTyp(values); err != nil {
		ret.err = fmt.Errorf("values is not same type")
		return ret
	}
	eleTyp := val.Index(0).Type()
	f := getElementsFactory(eleTyp)
	ret.elements = f.newElements(val.Len())
	for i := 0; i < val.Len(); i++ {
		ret.elements.Elem(i).Set(val.Index(i).Interface())
	}
	ret.t = eleTyp
	return ret
}

func checkSliceTyp(values interface{}) error {
	return nil
}

// Type returns the type of a given series
func (s Series) Type() reflect.Type {
	return s.t
}

// Len returns the length of a given Series
func (s Series) Len() int {
	return s.elements.Len()
}

// Str prints some extra information about a given series
func (s Series) Str() string {
	var ret []string
	// If name exists print name
	if s.Name != "" {
		ret = append(ret, "Name: "+s.Name)
	}
	ret = append(ret, "Type: "+s.t.String())
	ret = append(ret, "Length: "+fmt.Sprint(s.Len()))
	if s.Len() != 0 {
		ret = append(ret, "Values: "+s.elements.String())
	}
	return strings.Join(ret, "\n")
}
