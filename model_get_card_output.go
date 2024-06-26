/*
Lemonway DirectKit API 2.0

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetCardOutput struct for GetCardOutput
type GetCardOutput struct {
	Card *Card `json:"card,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewGetCardOutput instantiates a new GetCardOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCardOutput() *GetCardOutput {
	this := GetCardOutput{}
	return &this
}

// NewGetCardOutputWithDefaults instantiates a new GetCardOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCardOutputWithDefaults() *GetCardOutput {
	this := GetCardOutput{}
	return &this
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *GetCardOutput) GetCard() Card {
	if o == nil || o.Card == nil {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCardOutput) GetCardOk() (*Card, bool) {
	if o == nil || o.Card == nil {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *GetCardOutput) HasCard() bool {
	if o != nil && o.Card != nil {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *GetCardOutput) SetCard(v Card) {
	o.Card = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetCardOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCardOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetCardOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *GetCardOutput) SetError(v Error) {
	o.Error = &v
}

func (o GetCardOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Card != nil {
		toSerialize["card"] = o.Card
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableGetCardOutput struct {
	value *GetCardOutput
	isSet bool
}

func (v NullableGetCardOutput) Get() *GetCardOutput {
	return v.value
}

func (v *NullableGetCardOutput) Set(val *GetCardOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCardOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCardOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCardOutput(val *GetCardOutput) *NullableGetCardOutput {
	return &NullableGetCardOutput{value: val, isSet: true}
}

func (v NullableGetCardOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCardOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


