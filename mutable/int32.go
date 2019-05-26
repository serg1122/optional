package mutable

import "github.com/serg1122/optional"

type OptionalInt32 struct {
	isPresent bool
	value     int32
}

func OptionalInt32Create() *OptionalInt32 {
	return &OptionalInt32{
		isPresent: false,
	}
}

func (o *OptionalInt32) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalInt32) ValueGet() (int32, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return int32(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt32) ValueSet(value int32) {
	o.value = value
	o.isPresent = true
}
