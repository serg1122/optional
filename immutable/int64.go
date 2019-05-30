package immutable

import (
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalInt64 struct {
	isPresent bool
	value     int64
}

func OptionalInt64Create() *OptionalInt64 {
	return &OptionalInt64{
		isPresent: false,
	}
}

func (o *OptionalInt64) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalInt64) ValueGet() (int64, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt64) ValueSet(value int64) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o *OptionalInt64) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalInt64) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	if string(data) == "null" {
		return nil
	}
	var value int64
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.ValueSet(value); err2 != nil {
		return err2
	}
	return nil
}
