package immutable

import "github.com/serg1122/optional"

type OptionalString struct {
	isPresent bool
	value     string
}

func OptionalStringCreate() *OptionalString {
	return &OptionalString{
		isPresent: false,
	}
}

func (o *OptionalString) IsPresnt() bool {
	return o.isPresent
}

func (o *OptionalString) ValueGet() (string, *optional.ErrorValueIsNotPresent) {
	if o.IsPresnt() {
		return o.value, nil
	}
	return "", optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalString) ValueSet(value string) *optional.ErrorValueIsPresent {
	if o.IsPresnt() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}
