package immutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalInt8 struct {
	isPresent bool
	value     int8
}

func CreateOptionalInt8() *OptionalInt8 {
	return &OptionalInt8{
		isPresent: false,
	}
}

func (o OptionalInt8) IsPresent() bool {
	return o.isPresent
}

func (o OptionalInt8) GetValue() (int8, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return int8(0), optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalInt8) SetValue(value int8) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o *OptionalInt8) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalInt8) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value int8
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.SetValue(value); err2 != nil {
		return err2
	}
	return nil
}
