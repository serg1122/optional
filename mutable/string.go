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

func OptionalStringCreate() *OptionalString {
	return &OptionalString{
		isPresent: false,
	}
}

func (o OptionalString) IsPresnt() bool {
	return o.isPresent
}

func (o OptionalString) ValueGet() (string, *optional.ErrorValueIsNotPresent) {
	if o.IsPresnt() {
		return o.value, nil
	}
	return "", optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalString) ValueSet(value string) {
	o.value = value
	o.isPresent = true
}

func (o OptionalString) MarshalJSON() ([]byte, error) {
	if o.IsPresnt() {
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
	o.ValueSet(value)
	return nil
}
