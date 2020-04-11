package immutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalInt32 struct {
	isPresent bool
	value     int32
}

func CreateOptionalInt32() *OptionalInt32 {
	return &OptionalInt32{
		isPresent: false,
	}
}

func (o OptionalInt32) IsPresent() bool {
	return o.isPresent
}

func (o OptionalInt32) GetValue() (int32, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return int32(0), optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalInt32) SetValue(value int32) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o OptionalInt32) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalInt32) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value int32
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.SetValue(value); err2 != nil {
		return err2
	}
	return nil
}
