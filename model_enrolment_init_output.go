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

// EnrolmentInitOutput struct for EnrolmentInitOutput
type EnrolmentInitOutput struct {
	// Redirect URL to Deutsche Post PostIdent
	Redirecturl *string `json:"redirecturl,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewEnrolmentInitOutput instantiates a new EnrolmentInitOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrolmentInitOutput() *EnrolmentInitOutput {
	this := EnrolmentInitOutput{}
	return &this
}

// NewEnrolmentInitOutputWithDefaults instantiates a new EnrolmentInitOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrolmentInitOutputWithDefaults() *EnrolmentInitOutput {
	this := EnrolmentInitOutput{}
	return &this
}

// GetRedirecturl returns the Redirecturl field value if set, zero value otherwise.
func (o *EnrolmentInitOutput) GetRedirecturl() string {
	if o == nil || o.Redirecturl == nil {
		var ret string
		return ret
	}
	return *o.Redirecturl
}

// GetRedirecturlOk returns a tuple with the Redirecturl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolmentInitOutput) GetRedirecturlOk() (*string, bool) {
	if o == nil || o.Redirecturl == nil {
		return nil, false
	}
	return o.Redirecturl, true
}

// HasRedirecturl returns a boolean if a field has been set.
func (o *EnrolmentInitOutput) HasRedirecturl() bool {
	if o != nil && o.Redirecturl != nil {
		return true
	}

	return false
}

// SetRedirecturl gets a reference to the given string and assigns it to the Redirecturl field.
func (o *EnrolmentInitOutput) SetRedirecturl(v string) {
	o.Redirecturl = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *EnrolmentInitOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolmentInitOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *EnrolmentInitOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *EnrolmentInitOutput) SetError(v Error) {
	o.Error = &v
}

func (o EnrolmentInitOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Redirecturl != nil {
		toSerialize["redirecturl"] = o.Redirecturl
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableEnrolmentInitOutput struct {
	value *EnrolmentInitOutput
	isSet bool
}

func (v NullableEnrolmentInitOutput) Get() *EnrolmentInitOutput {
	return v.value
}

func (v *NullableEnrolmentInitOutput) Set(val *EnrolmentInitOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrolmentInitOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrolmentInitOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrolmentInitOutput(val *EnrolmentInitOutput) *NullableEnrolmentInitOutput {
	return &NullableEnrolmentInitOutput{value: val, isSet: true}
}

func (v NullableEnrolmentInitOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrolmentInitOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


