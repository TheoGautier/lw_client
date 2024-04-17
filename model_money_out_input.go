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

// MoneyOutInput struct for MoneyOutInput
type MoneyOutInput struct {
	// Payment Account ID to be debited
	AccountId string `json:"accountId"`
	// IBAN ID  If no IBAN is specified, the last verified(validated) IBAN will be used.
	IbanId *int64 `json:"ibanId,omitempty"`
	// Total amount to debit from the Wallet  The client will receive on his bank account[totalAmount] minus[commissionAmount].  Amounts are given as integer numbers in cents
	TotalAmount *int32 `json:"totalAmount,omitempty"`
	// Your fee  Amounts are given as integer numbers in cents
	CommissionAmount *int32 `json:"commissionAmount,omitempty"`
	// Payment Comment
	Comment *string `json:"comment,omitempty"`
	// This should be set to No (0) for most sites  If true:  1. [amountCom] will be ignored and will be replaced with Lemonway's fee.  2. You will not receive any fee.
	AutoCommission bool `json:"autoCommission"`
	// Unique ID of the call, generated by your server. This ID can be used as a search field when looking for operation details
	Reference *string `json:"reference,omitempty"`
}

// NewMoneyOutInput instantiates a new MoneyOutInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyOutInput(accountId string, autoCommission bool) *MoneyOutInput {
	this := MoneyOutInput{}
	this.AccountId = accountId
	this.AutoCommission = autoCommission
	return &this
}

// NewMoneyOutInputWithDefaults instantiates a new MoneyOutInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyOutInputWithDefaults() *MoneyOutInput {
	this := MoneyOutInput{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *MoneyOutInput) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MoneyOutInput) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MoneyOutInput) SetAccountId(v string) {
	o.AccountId = v
}

// GetIbanId returns the IbanId field value if set, zero value otherwise.
func (o *MoneyOutInput) GetIbanId() int64 {
	if o == nil || o.IbanId == nil {
		var ret int64
		return ret
	}
	return *o.IbanId
}

// GetIbanIdOk returns a tuple with the IbanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyOutInput) GetIbanIdOk() (*int64, bool) {
	if o == nil || o.IbanId == nil {
		return nil, false
	}
	return o.IbanId, true
}

// HasIbanId returns a boolean if a field has been set.
func (o *MoneyOutInput) HasIbanId() bool {
	if o != nil && o.IbanId != nil {
		return true
	}

	return false
}

// SetIbanId gets a reference to the given int64 and assigns it to the IbanId field.
func (o *MoneyOutInput) SetIbanId(v int64) {
	o.IbanId = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MoneyOutInput) GetTotalAmount() int32 {
	if o == nil || o.TotalAmount == nil {
		var ret int32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyOutInput) GetTotalAmountOk() (*int32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MoneyOutInput) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int32 and assigns it to the TotalAmount field.
func (o *MoneyOutInput) SetTotalAmount(v int32) {
	o.TotalAmount = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *MoneyOutInput) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyOutInput) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *MoneyOutInput) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *MoneyOutInput) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MoneyOutInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyOutInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MoneyOutInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MoneyOutInput) SetComment(v string) {
	o.Comment = &v
}

// GetAutoCommission returns the AutoCommission field value
func (o *MoneyOutInput) GetAutoCommission() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCommission
}

// GetAutoCommissionOk returns a tuple with the AutoCommission field value
// and a boolean to check if the value has been set.
func (o *MoneyOutInput) GetAutoCommissionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCommission, true
}

// SetAutoCommission sets field value
func (o *MoneyOutInput) SetAutoCommission(v bool) {
	o.AutoCommission = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MoneyOutInput) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyOutInput) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MoneyOutInput) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MoneyOutInput) SetReference(v string) {
	o.Reference = &v
}

func (o MoneyOutInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if o.IbanId != nil {
		toSerialize["ibanId"] = o.IbanId
	}
	if o.TotalAmount != nil {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if o.CommissionAmount != nil {
		toSerialize["commissionAmount"] = o.CommissionAmount
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["autoCommission"] = o.AutoCommission
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyOutInput struct {
	value *MoneyOutInput
	isSet bool
}

func (v NullableMoneyOutInput) Get() *MoneyOutInput {
	return v.value
}

func (v *NullableMoneyOutInput) Set(val *MoneyOutInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyOutInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyOutInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyOutInput(val *MoneyOutInput) *NullableMoneyOutInput {
	return &NullableMoneyOutInput{value: val, isSet: true}
}

func (v NullableMoneyOutInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyOutInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

