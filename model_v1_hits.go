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

// V1Hits struct for V1Hits
type V1Hits struct {
	Hits []V1Hit `json:"hits,omitempty"`
	MaxScore *float32 `json:"max_score,omitempty"`
	Total *V1Total `json:"total,omitempty"`
}

// NewV1Hits instantiates a new V1Hits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Hits() *V1Hits {
	this := V1Hits{}
	return &this
}

// NewV1HitsWithDefaults instantiates a new V1Hits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1HitsWithDefaults() *V1Hits {
	this := V1Hits{}
	return &this
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *V1Hits) GetHits() []V1Hit {
	if o == nil || o.Hits == nil {
		var ret []V1Hit
		return ret
	}
	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Hits) GetHitsOk() ([]V1Hit, bool) {
	if o == nil || o.Hits == nil {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *V1Hits) HasHits() bool {
	if o != nil && o.Hits != nil {
		return true
	}

	return false
}

// SetHits gets a reference to the given []V1Hit and assigns it to the Hits field.
func (o *V1Hits) SetHits(v []V1Hit) {
	o.Hits = v
}

// GetMaxScore returns the MaxScore field value if set, zero value otherwise.
func (o *V1Hits) GetMaxScore() float32 {
	if o == nil || o.MaxScore == nil {
		var ret float32
		return ret
	}
	return *o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Hits) GetMaxScoreOk() (*float32, bool) {
	if o == nil || o.MaxScore == nil {
		return nil, false
	}
	return o.MaxScore, true
}

// HasMaxScore returns a boolean if a field has been set.
func (o *V1Hits) HasMaxScore() bool {
	if o != nil && o.MaxScore != nil {
		return true
	}

	return false
}

// SetMaxScore gets a reference to the given float32 and assigns it to the MaxScore field.
func (o *V1Hits) SetMaxScore(v float32) {
	o.MaxScore = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V1Hits) GetTotal() V1Total {
	if o == nil || o.Total == nil {
		var ret V1Total
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Hits) GetTotalOk() (*V1Total, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V1Hits) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given V1Total and assigns it to the Total field.
func (o *V1Hits) SetTotal(v V1Total) {
	o.Total = &v
}

func (o V1Hits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hits != nil {
		toSerialize["hits"] = o.Hits
	}
	if o.MaxScore != nil {
		toSerialize["max_score"] = o.MaxScore
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableV1Hits struct {
	value *V1Hits
	isSet bool
}

func (v NullableV1Hits) Get() *V1Hits {
	return v.value
}

func (v *NullableV1Hits) Set(val *V1Hits) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Hits) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Hits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Hits(val *V1Hits) *NullableV1Hits {
	return &NullableV1Hits{value: val, isSet: true}
}

func (v NullableV1Hits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Hits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

