package immutable

import "github.com/serg1122/optional"

type OptionalBool struct {
	isPresent bool
	value     bool
}

func OptionalBoolCreate() *OptionalBool {
	return &OptionalBool{
		isPresent: false,
	}
}

func (o *OptionalBool) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalBool) ValueGet() (bool, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return false, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalBool) ValueSet(value bool) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}
