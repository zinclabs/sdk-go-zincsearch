/*
Zinc Search engine API

Zinc Search engine API documents https://docs.zincsearch.com

API version: 0.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MetaRegexpQuery struct for MetaRegexpQuery
type MetaRegexpQuery struct {
	Boost *float32 `json:"boost,omitempty"`
	Flags *string `json:"flags,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewMetaRegexpQuery instantiates a new MetaRegexpQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaRegexpQuery() *MetaRegexpQuery {
	this := MetaRegexpQuery{}
	return &this
}

// NewMetaRegexpQueryWithDefaults instantiates a new MetaRegexpQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaRegexpQueryWithDefaults() *MetaRegexpQuery {
	this := MetaRegexpQuery{}
	return &this
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *MetaRegexpQuery) GetBoost() float32 {
	if o == nil || o.Boost == nil {
		var ret float32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRegexpQuery) GetBoostOk() (*float32, bool) {
	if o == nil || o.Boost == nil {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *MetaRegexpQuery) HasBoost() bool {
	if o != nil && o.Boost != nil {
		return true
	}

	return false
}

// SetBoost gets a reference to the given float32 and assigns it to the Boost field.
func (o *MetaRegexpQuery) SetBoost(v float32) {
	o.Boost = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *MetaRegexpQuery) GetFlags() string {
	if o == nil || o.Flags == nil {
		var ret string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRegexpQuery) GetFlagsOk() (*string, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *MetaRegexpQuery) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given string and assigns it to the Flags field.
func (o *MetaRegexpQuery) SetFlags(v string) {
	o.Flags = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetaRegexpQuery) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRegexpQuery) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetaRegexpQuery) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MetaRegexpQuery) SetValue(v string) {
	o.Value = &v
}

func (o MetaRegexpQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Boost != nil {
		toSerialize["boost"] = o.Boost
	}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMetaRegexpQuery struct {
	value *MetaRegexpQuery
	isSet bool
}

func (v NullableMetaRegexpQuery) Get() *MetaRegexpQuery {
	return v.value
}

func (v *NullableMetaRegexpQuery) Set(val *MetaRegexpQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaRegexpQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaRegexpQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaRegexpQuery(val *MetaRegexpQuery) *NullableMetaRegexpQuery {
	return &NullableMetaRegexpQuery{value: val, isSet: true}
}

func (v NullableMetaRegexpQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaRegexpQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

