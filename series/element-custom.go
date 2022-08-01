package series

import "reflect"

type customElements struct {
	elements []Element
	typ      reflect.Type
}

func (e customElements) Elem(idx int) Element {
	return e.elements[idx]
}

func (e customElements) Len() int {
	return len(e.elements)
}

func (e customElements) Type() reflect.Type {
	return e.typ
}

func (e customElements) String() string {
	return ""
}

type customElementsFactory struct {
	typ     reflect.Type
	factory CustomElementProvider
}

func newCustomElementsFactory(typ reflect.Type, factory CustomElementProvider) elementsFactory {
	return customElementsFactory{typ: typ, factory: factory}
}

func (f customElementsFactory) new() Element {
	return f.factory.New()
}

func (f customElementsFactory) newElements(num int) Elements {
	underlyingElements := make([]Element, 0, num)
	for i := 0; i < num; i++ {
		underlyingElements = append(underlyingElements, f.new())
	}
	return customElements{
		elements: underlyingElements,
		typ:      f.typ,
	}
}
