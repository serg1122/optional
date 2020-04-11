package immutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalUint8 struct {
	isPresent bool
	value     uint8
}

func CreateOptionalUint8() *OptionalUint8 {
	return &OptionalUint8{
		isPresent: false,
	}
}

func (o OptionalUint8) IsPresent() bool {
	return o.isPresent
}

func (o OptionalUint8) GetValue() (uint8, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint8(0), optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalUint8) SetValue(value uint8) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o OptionalUint8) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalUint8) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value uint8
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.SetValue(value); err2 != nil {
		return err2
	}
	return nil
}
