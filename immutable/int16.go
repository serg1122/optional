package immutable

import "github.com/serg1122/optional"

type OptionalInt16 struct {
	isPresent bool
	value     int16
}

func OptionalInt16Create() *OptionalInt16 {
	return &OptionalInt16{
		isPresent: false,
	}
}

func (o *OptionalInt16) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalInt16) ValueGet() (int16, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return int16(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt16) ValueSet(value int16) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}
