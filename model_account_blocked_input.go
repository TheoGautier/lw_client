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

// AccountBlockedInput struct for AccountBlockedInput
type AccountBlockedInput struct {
	IsBlocked bool `json:"isBlocked"`
	Comment *string `json:"comment,omitempty"`
}

// NewAccountBlockedInput instantiates a new AccountBlockedInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBlockedInput(isBlocked bool) *AccountBlockedInput {
	this := AccountBlockedInput{}
	this.IsBlocked = isBlocked
	return &this
}

// NewAccountBlockedInputWithDefaults instantiates a new AccountBlockedInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBlockedInputWithDefaults() *AccountBlockedInput {
	this := AccountBlockedInput{}
	return &this
}

// GetIsBlocked returns the IsBlocked field value
func (o *AccountBlockedInput) GetIsBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBlocked
}

// GetIsBlockedOk returns a tuple with the IsBlocked field value
// and a boolean to check if the value has been set.
func (o *AccountBlockedInput) GetIsBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBlocked, true
}

// SetIsBlocked sets field value
func (o *AccountBlockedInput) SetIsBlocked(v bool) {
	o.IsBlocked = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AccountBlockedInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockedInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AccountBlockedInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AccountBlockedInput) SetComment(v string) {
	o.Comment = &v
}

func (o AccountBlockedInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isBlocked"] = o.IsBlocked
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableAccountBlockedInput struct {
	value *AccountBlockedInput
	isSet bool
}

func (v NullableAccountBlockedInput) Get() *AccountBlockedInput {
	return v.value
}

func (v *NullableAccountBlockedInput) Set(val *AccountBlockedInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBlockedInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBlockedInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBlockedInput(val *AccountBlockedInput) *NullableAccountBlockedInput {
	return &NullableAccountBlockedInput{value: val, isSet: true}
}

func (v NullableAccountBlockedInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBlockedInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


