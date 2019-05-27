package immutable

import (
	"encoding/json"

	"github.com/serg1122/optional"
)

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

func (o *OptionalBool) MarshalJSON() ([]byte, error) {

	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalBool) UnmarshalJSON(data []byte) error {

	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	if string(data) == "null" {
		return nil
	}
	var value bool
	if errorUnmarshal := json.Unmarshal(data, &value); errorUnmarshal != nil {
		return errorUnmarshal
	}
	if errorValueSet := o.ValueSet(value); errorValueSet != nil {
		return errorValueSet
	}
	return nil
}
