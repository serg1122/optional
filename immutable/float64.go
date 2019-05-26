package immutable

import "github.com/serg1122/optional"

type OptionalFloat64 struct {
	isPresent bool
	value     float64
}

func OptionalFloat64Create() *OptionalFloat64 {
	return &OptionalFloat64{
		isPresent: false,
	}
}

func (o *OptionalFloat64) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalFloat64) ValueGet() (float64, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0.0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalFloat64) ValueSet(value float64) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}
