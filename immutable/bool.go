package immutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalBool struct {
	isPresent bool
	value     bool
}

func CreateOptionalBool() *OptionalBool {
	return &OptionalBool{
		isPresent: false,
	}
}

func (o OptionalBool) IsPresent() bool {
	return o.isPresent
}

func (o OptionalBool) GetValue() (bool, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return false, optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalBool) SetValue(value bool) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o OptionalBool) MarshalJSON() ([]byte, error) {

	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalBool) UnmarshalJSON(data []byte) error {

	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value bool
	if errorUnmarshal := json.Unmarshal(data, &value); errorUnmarshal != nil {
		return errorUnmarshal
	}
	if errorSetValue := o.SetValue(value); errorSetValue != nil {
		return errorSetValue
	}
	return nil
}
