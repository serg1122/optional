package immutable

import "github.com/serg1122/optional"

type OptionalUint8 struct {
	isPresent bool
	value     uint8
}

func OptionalUint8Create() *OptionalUint8 {
	return &OptionalUint8{
		isPresent: false,
	}
}

func (o *OptionalUint8) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalUint8) ValueGet() (uint8, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint8(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint8) ValueSet(value uint8) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}
