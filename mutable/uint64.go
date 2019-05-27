package mutable

import "github.com/serg1122/optional"

type OptionalUint64 struct {
	isPresent bool
	value     uint64
}

func OptionalUint64Create() *OptionalUint64 {
	return &OptionalUint64{
		isPresent: false,
	}
}

func (o *OptionalUint64) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalUint64) ValueGet() (uint64, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint64(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint64) ValueSet(value uint64) {
	o.value = value
	o.isPresent = true
}
