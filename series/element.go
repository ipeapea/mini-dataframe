package series

import (
	"fmt"
	"reflect"
	"sync"
)

type CompareResult int

var (
	CompareResultTypeErr   CompareResult = -4
	CompareResultSelfNull  CompareResult = -3
	CompareResultParamNull CompareResult = -2
	CompareResultLess      CompareResult = -1
	CompareResultEqual     CompareResult = 0
	CompareResultGreater   CompareResult = 1
)

// Element is the basic interface which series built on
type Element interface {
	Set(interface{})
	Get() interface{}
	Clone() Element
	Null() bool
	Type() reflect.Type
	String() string

	// Compare -1: litter; 0: equal; 1: greater; -2: element self is null; -3: param element is null
	Compare(element Element) CompareResult
}

// Elements is the interface that represents the array of elements contained on a Series.
type Elements interface {
	Elem(int) Element
	Len() int
	Type() reflect.Type
	String() string
}

// CustomElementProvider provide element implemented by users themselves when the built-in type does not meet the requirements
type CustomElementProvider interface {
	New() Element
}

var (
	// Null empty null representation
	Null                      = "Null"
	localElementsFactoryMutex = sync.Mutex{}
	localElementsFactory      = map[reflect.Type]elementsFactory{}
)

// RegisterCustomElement register custom element before using it
func RegisterCustomElement(typ reflect.Type, factory CustomElementProvider) error {
	localElementsFactoryMutex.Lock()
	defer localElementsFactoryMutex.Unlock()
	if _, ok := localElementsFactory[typ]; !ok {
		localElementsFactory[typ] = newCustomElementsFactory(typ, factory)
	} else {
		return fmt.Errorf("type [%s] has already registered", typ.String())
	}
	return nil
}

type elementsFactory interface {
	new() Element
	newElements(num int) Elements
	//Append(elements Elements, ele ...Element) Elements
	//SubSet(elements Elements, start, end int) Elements
}

func getElementsFactory(typ reflect.Type) elementsFactory {
	localElementsFactoryMutex.Lock()
	defer localElementsFactoryMutex.Unlock()
	f, ok := localElementsFactory[typ]
	if ok {
		return f
	}
	f = newCustomElementsFactory(typ, commonElementProvider{typ: typ})
	localElementsFactory[typ] = f
	return f
}
