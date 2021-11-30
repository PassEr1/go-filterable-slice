package filterable

import (
	"reflect"
)

type FilterableContainer struct {
	InnerSlice interface{}
}

type FilterLambda func(interface{}) bool

type Filterable interface {
	Filter(FilterLambda) Filterable
	Get() []interface{}
}

func (container FilterableContainer) Filter(lambda FilterLambda) FilterableContainer {
	new_slice := make([]interface{}, 0)

	switch reflect.TypeOf(container.InnerSlice).Kind() {
	case reflect.Slice:
		extracted_slice := reflect.ValueOf(container.InnerSlice)
		for index := 0; index < extracted_slice.Len(); index++ {
			element := extracted_slice.Index(index).Interface()
			if lambda(element) != true {
				continue
			}

			new_slice = append(new_slice, element)
		}
	}
	return FilterableContainer{new_slice}
}

func (container FilterableContainer) Get() interface{} {
	return container.InnerSlice
}

func New(values interface{}) *FilterableContainer {
	as_array := values
	return &FilterableContainer{
		InnerSlice: as_array,
	}
}
