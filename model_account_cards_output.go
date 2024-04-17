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

// AccountCardsOutput struct for AccountCardsOutput
type AccountCardsOutput struct {
	// Dispalays a list of documents that have changed since original entry date.
	Cards []Card `json:"cards,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewAccountCardsOutput instantiates a new AccountCardsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCardsOutput() *AccountCardsOutput {
	this := AccountCardsOutput{}
	return &this
}

// NewAccountCardsOutputWithDefaults instantiates a new AccountCardsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCardsOutputWithDefaults() *AccountCardsOutput {
	this := AccountCardsOutput{}
	return &this
}

// GetCards returns the Cards field value if set, zero value otherwise.
func (o *AccountCardsOutput) GetCards() []Card {
	if o == nil || o.Cards == nil {
		var ret []Card
		return ret
	}
	return o.Cards
}

// GetCardsOk returns a tuple with the Cards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCardsOutput) GetCardsOk() ([]Card, bool) {
	if o == nil || o.Cards == nil {
		return nil, false
	}
	return o.Cards, true
}

// HasCards returns a boolean if a field has been set.
func (o *AccountCardsOutput) HasCards() bool {
	if o != nil && o.Cards != nil {
		return true
	}

	return false
}

// SetCards gets a reference to the given []Card and assigns it to the Cards field.
func (o *AccountCardsOutput) SetCards(v []Card) {
	o.Cards = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AccountCardsOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCardsOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AccountCardsOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *AccountCardsOutput) SetError(v Error) {
	o.Error = &v
}

func (o AccountCardsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cards != nil {
		toSerialize["cards"] = o.Cards
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCardsOutput struct {
	value *AccountCardsOutput
	isSet bool
}

func (v NullableAccountCardsOutput) Get() *AccountCardsOutput {
	return v.value
}

func (v *NullableAccountCardsOutput) Set(val *AccountCardsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCardsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCardsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCardsOutput(val *AccountCardsOutput) *NullableAccountCardsOutput {
	return &NullableAccountCardsOutput{value: val, isSet: true}
}

func (v NullableAccountCardsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCardsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


