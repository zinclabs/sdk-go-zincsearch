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

// V1AggregationDateRange struct for V1AggregationDateRange
type V1AggregationDateRange struct {
	From *string `json:"from,omitempty"`
	To *string `json:"to,omitempty"`
}

// NewV1AggregationDateRange instantiates a new V1AggregationDateRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AggregationDateRange() *V1AggregationDateRange {
	this := V1AggregationDateRange{}
	return &this
}

// NewV1AggregationDateRangeWithDefaults instantiates a new V1AggregationDateRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AggregationDateRangeWithDefaults() *V1AggregationDateRange {
	this := V1AggregationDateRange{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *V1AggregationDateRange) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AggregationDateRange) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *V1AggregationDateRange) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *V1AggregationDateRange) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *V1AggregationDateRange) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AggregationDateRange) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *V1AggregationDateRange) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *V1AggregationDateRange) SetTo(v string) {
	o.To = &v
}

func (o V1AggregationDateRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableV1AggregationDateRange struct {
	value *V1AggregationDateRange
	isSet bool
}

func (v NullableV1AggregationDateRange) Get() *V1AggregationDateRange {
	return v.value
}

func (v *NullableV1AggregationDateRange) Set(val *V1AggregationDateRange) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AggregationDateRange) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AggregationDateRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AggregationDateRange(val *V1AggregationDateRange) *NullableV1AggregationDateRange {
	return &NullableV1AggregationDateRange{value: val, isSet: true}
}

func (v NullableV1AggregationDateRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AggregationDateRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

