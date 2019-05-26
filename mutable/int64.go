package mutable

import "github.com/serg1122/optional"

type OptionalInt64 struct {
	isPresent bool
	value     int64
}

func OptionalInt64Create() *OptionalInt64 {
	return &OptionalInt64{
		isPresent: false,
	}
}

func (o *OptionalInt64) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalInt64) ValueGet() (int64, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt64) ValueSet(value int64) {
	o.value = value
	o.isPresent = true
}
