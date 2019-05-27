package immutable

import "github.com/serg1122/optional"

type OptionalUint32 struct {
	isPresent bool
	value     uint32
}

func OptinoalUint32Create() *OptionalUint32 {
	return &OptionalUint32{
		isPresent: false,
	}
}

func (o *OptionalUint32) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalUint32) ValueGet() (uint32, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint32(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint32) ValueSet(value uint32) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}
