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

// V1AggregationResponse struct for V1AggregationResponse
type V1AggregationResponse struct {
	Buckets []V1AggregationBucket `json:"buckets,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewV1AggregationResponse instantiates a new V1AggregationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AggregationResponse() *V1AggregationResponse {
	this := V1AggregationResponse{}
	return &this
}

// NewV1AggregationResponseWithDefaults instantiates a new V1AggregationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AggregationResponseWithDefaults() *V1AggregationResponse {
	this := V1AggregationResponse{}
	return &this
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *V1AggregationResponse) GetBuckets() []V1AggregationBucket {
	if o == nil || o.Buckets == nil {
		var ret []V1AggregationBucket
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AggregationResponse) GetBucketsOk() ([]V1AggregationBucket, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *V1AggregationResponse) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []V1AggregationBucket and assigns it to the Buckets field.
func (o *V1AggregationResponse) SetBuckets(v []V1AggregationBucket) {
	o.Buckets = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V1AggregationResponse) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AggregationResponse) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V1AggregationResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *V1AggregationResponse) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o V1AggregationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableV1AggregationResponse struct {
	value *V1AggregationResponse
	isSet bool
}

func (v NullableV1AggregationResponse) Get() *V1AggregationResponse {
	return v.value
}

func (v *NullableV1AggregationResponse) Set(val *V1AggregationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AggregationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AggregationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AggregationResponse(val *V1AggregationResponse) *NullableV1AggregationResponse {
	return &NullableV1AggregationResponse{value: val, isSet: true}
}

func (v NullableV1AggregationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AggregationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

