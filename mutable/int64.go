package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalInt64 struct {
	isPresent bool
	value     int64
}

func CreateOptionalInt64() *OptionalInt64 {
	return &OptionalInt64{
		isPresent: false,
	}
}

func (o OptionalInt64) IsPresent() bool {
	return o.isPresent
}

func (o OptionalInt64) GetValue() (int64, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0, optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalInt64) SetValue(value int64) {
	o.value = value
	o.isPresent = true
}

func (o OptionalInt64) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalInt64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value int64
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.SetValue(value)
	return nil
}
