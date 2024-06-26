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

// MoneyInSddInitInput struct for MoneyInSddInitInput
type MoneyInSddInitInput struct {
	// Mandate ID
	SddMandateId int64 `json:"sddMandateId"`
	// Debit date of the bank account, must be later than the default date.  Leave empty in order to use the default date: current date + 1 working days before 10:30 AM or current date + 2 working days after.
	CollectionDate *string `json:"collectionDate,omitempty"`
	// Unique ID of the call, generated by your server. This ID can be used as a search field when looking for operation details
	Reference *string `json:"reference,omitempty"`
	// Payment Account ID to Credit
	AccountId string `json:"accountId"`
	// Amount to Debit  Amounts are given as integer numbers in cents
	TotalAmount *int32 `json:"totalAmount,omitempty"`
	// Your Fee  Amounts are given as integer numbers in cents
	CommissionAmount *int32 `json:"commissionAmount,omitempty"`
	// Comment Regarding the Transaction
	Comment *string `json:"comment,omitempty"`
	// If true:  1. [amountCom] will be ignored and will be replaced with Lemonway's fee  2. You will not receive any fee
	AutoCommission *bool `json:"autoCommission,omitempty"`
}

// NewMoneyInSddInitInput instantiates a new MoneyInSddInitInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInSddInitInput(sddMandateId int64, accountId string) *MoneyInSddInitInput {
	this := MoneyInSddInitInput{}
	this.SddMandateId = sddMandateId
	this.AccountId = accountId
	return &this
}

// NewMoneyInSddInitInputWithDefaults instantiates a new MoneyInSddInitInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInSddInitInputWithDefaults() *MoneyInSddInitInput {
	this := MoneyInSddInitInput{}
	return &this
}

// GetSddMandateId returns the SddMandateId field value
func (o *MoneyInSddInitInput) GetSddMandateId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SddMandateId
}

// GetSddMandateIdOk returns a tuple with the SddMandateId field value
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitInput) GetSddMandateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SddMandateId, true
}

// SetSddMandateId sets field value
func (o *MoneyInSddInitInput) SetSddMandateId(v int64) {
	o.SddMandateId = v
}

// GetCollectionDate returns the CollectionDate field value if set, zero value otherwise.
func (o *MoneyInSddInitInput) GetCollectionDate() string {
	if o == nil || o.CollectionDate == nil {
		var ret string
		return ret
	}
	return *o.CollectionDate
}

// GetCollectionDateOk returns a tuple with the CollectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitInput) GetCollectionDateOk() (*string, bool) {
	if o == nil || o.CollectionDate == nil {
		return nil, false
	}
	return o.CollectionDate, true
}

// HasCollectionDate returns a boolean if a field has been set.
func (o *MoneyInSddInitInput) HasCollectionDate() bool {
	if o != nil && o.CollectionDate != nil {
		return true
	}

	return false
}

// SetCollectionDate gets a reference to the given string and assigns it to the CollectionDate field.
func (o *MoneyInSddInitInput) SetCollectionDate(v string) {
	o.CollectionDate = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MoneyInSddInitInput) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitInput) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MoneyInSddInitInput) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MoneyInSddInitInput) SetReference(v string) {
	o.Reference = &v
}

// GetAccountId returns the AccountId field value
func (o *MoneyInSddInitInput) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitInput) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MoneyInSddInitInput) SetAccountId(v string) {
	o.AccountId = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MoneyInSddInitInput) GetTotalAmount() int32 {
	if o == nil || o.TotalAmount == nil {
		var ret int32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitInput) GetTotalAmountOk() (*int32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MoneyInSddInitInput) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int32 and assigns it to the TotalAmount field.
func (o *MoneyInSddInitInput) SetTotalAmount(v int32) {
	o.TotalAmount = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *MoneyInSddInitInput) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitInput) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *MoneyInSddInitInput) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *MoneyInSddInitInput) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MoneyInSddInitInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MoneyInSddInitInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MoneyInSddInitInput) SetComment(v string) {
	o.Comment = &v
}

// GetAutoCommission returns the AutoCommission field value if set, zero value otherwise.
func (o *MoneyInSddInitInput) GetAutoCommission() bool {
	if o == nil || o.AutoCommission == nil {
		var ret bool
		return ret
	}
	return *o.AutoCommission
}

// GetAutoCommissionOk returns a tuple with the AutoCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitInput) GetAutoCommissionOk() (*bool, bool) {
	if o == nil || o.AutoCommission == nil {
		return nil, false
	}
	return o.AutoCommission, true
}

// HasAutoCommission returns a boolean if a field has been set.
func (o *MoneyInSddInitInput) HasAutoCommission() bool {
	if o != nil && o.AutoCommission != nil {
		return true
	}

	return false
}

// SetAutoCommission gets a reference to the given bool and assigns it to the AutoCommission field.
func (o *MoneyInSddInitInput) SetAutoCommission(v bool) {
	o.AutoCommission = &v
}

func (o MoneyInSddInitInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sddMandateId"] = o.SddMandateId
	}
	if o.CollectionDate != nil {
		toSerialize["collectionDate"] = o.CollectionDate
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["accountId"] = o.AccountId
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
	if o.AutoCommission != nil {
		toSerialize["autoCommission"] = o.AutoCommission
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyInSddInitInput struct {
	value *MoneyInSddInitInput
	isSet bool
}

func (v NullableMoneyInSddInitInput) Get() *MoneyInSddInitInput {
	return v.value
}

func (v *NullableMoneyInSddInitInput) Set(val *MoneyInSddInitInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInSddInitInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInSddInitInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInSddInitInput(val *MoneyInSddInitInput) *NullableMoneyInSddInitInput {
	return &NullableMoneyInSddInitInput{value: val, isSet: true}
}

func (v NullableMoneyInSddInitInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInSddInitInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


