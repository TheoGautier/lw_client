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

// UnregisterSddMandateOutput struct for UnregisterSddMandateOutput
type UnregisterSddMandateOutput struct {
	SDDMandate *SddMandate `json:"SDDMandate,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewUnregisterSddMandateOutput instantiates a new UnregisterSddMandateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnregisterSddMandateOutput() *UnregisterSddMandateOutput {
	this := UnregisterSddMandateOutput{}
	return &this
}

// NewUnregisterSddMandateOutputWithDefaults instantiates a new UnregisterSddMandateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnregisterSddMandateOutputWithDefaults() *UnregisterSddMandateOutput {
	this := UnregisterSddMandateOutput{}
	return &this
}

// GetSDDMandate returns the SDDMandate field value if set, zero value otherwise.
func (o *UnregisterSddMandateOutput) GetSDDMandate() SddMandate {
	if o == nil || o.SDDMandate == nil {
		var ret SddMandate
		return ret
	}
	return *o.SDDMandate
}

// GetSDDMandateOk returns a tuple with the SDDMandate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnregisterSddMandateOutput) GetSDDMandateOk() (*SddMandate, bool) {
	if o == nil || o.SDDMandate == nil {
		return nil, false
	}
	return o.SDDMandate, true
}

// HasSDDMandate returns a boolean if a field has been set.
func (o *UnregisterSddMandateOutput) HasSDDMandate() bool {
	if o != nil && o.SDDMandate != nil {
		return true
	}

	return false
}

// SetSDDMandate gets a reference to the given SddMandate and assigns it to the SDDMandate field.
func (o *UnregisterSddMandateOutput) SetSDDMandate(v SddMandate) {
	o.SDDMandate = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UnregisterSddMandateOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnregisterSddMandateOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UnregisterSddMandateOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *UnregisterSddMandateOutput) SetError(v Error) {
	o.Error = &v
}

func (o UnregisterSddMandateOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SDDMandate != nil {
		toSerialize["SDDMandate"] = o.SDDMandate
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableUnregisterSddMandateOutput struct {
	value *UnregisterSddMandateOutput
	isSet bool
}

func (v NullableUnregisterSddMandateOutput) Get() *UnregisterSddMandateOutput {
	return v.value
}

func (v *NullableUnregisterSddMandateOutput) Set(val *UnregisterSddMandateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUnregisterSddMandateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUnregisterSddMandateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnregisterSddMandateOutput(val *UnregisterSddMandateOutput) *NullableUnregisterSddMandateOutput {
	return &NullableUnregisterSddMandateOutput{value: val, isSet: true}
}

func (v NullableUnregisterSddMandateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnregisterSddMandateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


