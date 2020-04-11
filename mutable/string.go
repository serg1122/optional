package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalString struct {
	isPresent bool
	value     string
}

func CreateOptionalString() *OptionalString {
	return &OptionalString{
		isPresent: false,
	}
}

func (o OptionalString) IsPresent() bool {
	return o.isPresent
}

func (o OptionalString) GetValue() (string, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return "", optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalString) SetValue(value string) {
	o.value = value
	o.isPresent = true
}

func (o OptionalString) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value string
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.SetValue(value)
	return nil
}
