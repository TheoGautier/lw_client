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

// GetMoneyInSddInput struct for GetMoneyInSddInput
type GetMoneyInSddInput struct {
	// Unique ID Generated by your Server
	Reference *string `json:"reference,omitempty"`
	// UTC Unix Timestamp
	UpdateDate *string `json:"updateDate,omitempty"`
	// End Date UTC Unix Timestamp
	UpdateEndDate *string `json:"UpdateEndDate,omitempty"`
}

// NewGetMoneyInSddInput instantiates a new GetMoneyInSddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMoneyInSddInput() *GetMoneyInSddInput {
	this := GetMoneyInSddInput{}
	return &this
}

// NewGetMoneyInSddInputWithDefaults instantiates a new GetMoneyInSddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMoneyInSddInputWithDefaults() *GetMoneyInSddInput {
	this := GetMoneyInSddInput{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GetMoneyInSddInput) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyInSddInput) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GetMoneyInSddInput) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GetMoneyInSddInput) SetReference(v string) {
	o.Reference = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *GetMoneyInSddInput) GetUpdateDate() string {
	if o == nil || o.UpdateDate == nil {
		var ret string
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyInSddInput) GetUpdateDateOk() (*string, bool) {
	if o == nil || o.UpdateDate == nil {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *GetMoneyInSddInput) HasUpdateDate() bool {
	if o != nil && o.UpdateDate != nil {
		return true
	}

	return false
}

// SetUpdateDate gets a reference to the given string and assigns it to the UpdateDate field.
func (o *GetMoneyInSddInput) SetUpdateDate(v string) {
	o.UpdateDate = &v
}

// GetUpdateEndDate returns the UpdateEndDate field value if set, zero value otherwise.
func (o *GetMoneyInSddInput) GetUpdateEndDate() string {
	if o == nil || o.UpdateEndDate == nil {
		var ret string
		return ret
	}
	return *o.UpdateEndDate
}

// GetUpdateEndDateOk returns a tuple with the UpdateEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyInSddInput) GetUpdateEndDateOk() (*string, bool) {
	if o == nil || o.UpdateEndDate == nil {
		return nil, false
	}
	return o.UpdateEndDate, true
}

// HasUpdateEndDate returns a boolean if a field has been set.
func (o *GetMoneyInSddInput) HasUpdateEndDate() bool {
	if o != nil && o.UpdateEndDate != nil {
		return true
	}

	return false
}

// SetUpdateEndDate gets a reference to the given string and assigns it to the UpdateEndDate field.
func (o *GetMoneyInSddInput) SetUpdateEndDate(v string) {
	o.UpdateEndDate = &v
}

func (o GetMoneyInSddInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.UpdateDate != nil {
		toSerialize["updateDate"] = o.UpdateDate
	}
	if o.UpdateEndDate != nil {
		toSerialize["UpdateEndDate"] = o.UpdateEndDate
	}
	return json.Marshal(toSerialize)
}

type NullableGetMoneyInSddInput struct {
	value *GetMoneyInSddInput
	isSet bool
}

func (v NullableGetMoneyInSddInput) Get() *GetMoneyInSddInput {
	return v.value
}

func (v *NullableGetMoneyInSddInput) Set(val *GetMoneyInSddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMoneyInSddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMoneyInSddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMoneyInSddInput(val *GetMoneyInSddInput) *NullableGetMoneyInSddInput {
	return &NullableGetMoneyInSddInput{value: val, isSet: true}
}

func (v NullableGetMoneyInSddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMoneyInSddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


