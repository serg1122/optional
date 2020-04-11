package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalInt16 struct {
	isPresent bool
	value     int16
}

func CreateOptionalInt16() *OptionalInt16 {
	return &OptionalInt16{
		isPresent: false,
	}
}

func (o OptionalInt16) IsPresent() bool {
	return o.isPresent
}

func (o OptionalInt16) GetValue() (int16, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return int16(0), optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalInt16) SetValue(value int16) {
	o.value = value
	o.isPresent = true
}

func (o OptionalInt16) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalInt16) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value int16
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.SetValue(value)
	return nil
}
