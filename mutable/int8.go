package mutable

import "github.com/serg1122/optional"

type OptionalInt8 struct {
	isPresent bool
	value     int8
}

func OptionalInt8Create() *OptionalInt8 {
	return &OptionalInt8{
		isPresent: false,
	}
}

func (o *OptionalInt8) IsPresent() bool {
	return o.isPresent
}

func (o OptionalInt8) ValueGet() (int8, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return int8(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt8) ValueSet(value int8) {
	o.value = value
	o.isPresent = true
}
