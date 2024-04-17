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

// MoneyInValidateInput struct for MoneyInValidateInput
type MoneyInValidateInput struct {
	// Amount to Debit  Amounts are displayed as an integer in cents (Euro)
	TotalAmount int32 `json:"totalAmount"`
	// Your fee  Amounts are displayed as an integer in cents
	CommissionAmount *int32 `json:"commissionAmount,omitempty"`
	// Comment Regarding the Transaction
	Comment *string `json:"comment,omitempty"`
	// Unique ID of the call, generated by your server. This ID can be used as a search field when looking for operation details
	Reference *string `json:"reference,omitempty"`
	// Leave empty
	SpecialConfig *string `json:"specialConfig,omitempty"`
}

// NewMoneyInValidateInput instantiates a new MoneyInValidateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInValidateInput(totalAmount int32) *MoneyInValidateInput {
	this := MoneyInValidateInput{}
	this.TotalAmount = totalAmount
	return &this
}

// NewMoneyInValidateInputWithDefaults instantiates a new MoneyInValidateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInValidateInputWithDefaults() *MoneyInValidateInput {
	this := MoneyInValidateInput{}
	return &this
}

// GetTotalAmount returns the TotalAmount field value
func (o *MoneyInValidateInput) GetTotalAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *MoneyInValidateInput) GetTotalAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *MoneyInValidateInput) SetTotalAmount(v int32) {
	o.TotalAmount = v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *MoneyInValidateInput) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInValidateInput) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *MoneyInValidateInput) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *MoneyInValidateInput) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MoneyInValidateInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInValidateInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MoneyInValidateInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MoneyInValidateInput) SetComment(v string) {
	o.Comment = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MoneyInValidateInput) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInValidateInput) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MoneyInValidateInput) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MoneyInValidateInput) SetReference(v string) {
	o.Reference = &v
}

// GetSpecialConfig returns the SpecialConfig field value if set, zero value otherwise.
func (o *MoneyInValidateInput) GetSpecialConfig() string {
	if o == nil || o.SpecialConfig == nil {
		var ret string
		return ret
	}
	return *o.SpecialConfig
}

// GetSpecialConfigOk returns a tuple with the SpecialConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInValidateInput) GetSpecialConfigOk() (*string, bool) {
	if o == nil || o.SpecialConfig == nil {
		return nil, false
	}
	return o.SpecialConfig, true
}

// HasSpecialConfig returns a boolean if a field has been set.
func (o *MoneyInValidateInput) HasSpecialConfig() bool {
	if o != nil && o.SpecialConfig != nil {
		return true
	}

	return false
}

// SetSpecialConfig gets a reference to the given string and assigns it to the SpecialConfig field.
func (o *MoneyInValidateInput) SetSpecialConfig(v string) {
	o.SpecialConfig = &v
}

func (o MoneyInValidateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if o.CommissionAmount != nil {
		toSerialize["commissionAmount"] = o.CommissionAmount
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.SpecialConfig != nil {
		toSerialize["specialConfig"] = o.SpecialConfig
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyInValidateInput struct {
	value *MoneyInValidateInput
	isSet bool
}

func (v NullableMoneyInValidateInput) Get() *MoneyInValidateInput {
	return v.value
}

func (v *NullableMoneyInValidateInput) Set(val *MoneyInValidateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInValidateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInValidateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInValidateInput(val *MoneyInValidateInput) *NullableMoneyInValidateInput {
	return &NullableMoneyInValidateInput{value: val, isSet: true}
}

func (v NullableMoneyInValidateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInValidateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


