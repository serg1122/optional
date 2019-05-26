package immutable

import "github.com/serg1122/optional"

type OptionalInt struct {
	isPresent bool
	value     int
}

func OptionalIntCreate() *OptionalInt {
	return &OptionalInt{
		isPresent: false,
	}
}

func (o *OptionalInt) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalInt) ValueGet() (int, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt) ValueSet(value int) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}
