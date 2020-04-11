package immutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalUint64 struct {
	isPresent bool
	value     uint64
}

func OptionalUint64Create() *OptionalUint64 {
	return &OptionalUint64{
		isPresent: false,
	}
}

func (o OptionalUint64) IsPresent() bool {
	return o.isPresent
}

func (o OptionalUint64) GetValue() (uint64, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint64(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint64) ValueSet(value uint64) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o OptionalUint64) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalUint64) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value uint64
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.ValueSet(value); err2 != nil {
		return err2
	}
	return nil
}
