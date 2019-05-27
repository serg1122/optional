package immutable

import "github.com/serg1122/optional"

type OptionalUint struct {
	isPresent bool
	value     uint
}

func OptionalUintCreate() *OptionalUint {
	return &OptionalUint{
		isPresent: false,
	}
}

func (o *OptionalUint) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalUint) ValueGet() (uint, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint) ValueSet(value uint) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}
