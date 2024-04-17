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

// PurchaseItem Item
type PurchaseItem struct {
	// Merchant account id
	MerchantAccountId string `json:"merchantAccountId"`
	// Description
	Description string `json:"description"`
	// Quantity
	Quantity int32 `json:"quantity"`
	// Unit amount
	UnitAmount int32 `json:"unitAmount"`
	// Type<br/>1 = Physical.<br/>2 = Digital.<br/>
	Type int32 `json:"type"`
	// Tax amount
	TaxAmount *int32 `json:"taxAmount,omitempty"`
}

// NewPurchaseItem instantiates a new PurchaseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseItem(merchantAccountId string, description string, quantity int32, unitAmount int32, type_ int32) *PurchaseItem {
	this := PurchaseItem{}
	this.MerchantAccountId = merchantAccountId
	this.Description = description
	this.Quantity = quantity
	this.UnitAmount = unitAmount
	this.Type = type_
	return &this
}

// NewPurchaseItemWithDefaults instantiates a new PurchaseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseItemWithDefaults() *PurchaseItem {
	this := PurchaseItem{}
	return &this
}

// GetMerchantAccountId returns the MerchantAccountId field value
func (o *PurchaseItem) GetMerchantAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccountId
}

// GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field value
// and a boolean to check if the value has been set.
func (o *PurchaseItem) GetMerchantAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccountId, true
}

// SetMerchantAccountId sets field value
func (o *PurchaseItem) SetMerchantAccountId(v string) {
	o.MerchantAccountId = v
}

// GetDescription returns the Description field value
func (o *PurchaseItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PurchaseItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PurchaseItem) SetDescription(v string) {
	o.Description = v
}

// GetQuantity returns the Quantity field value
func (o *PurchaseItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PurchaseItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PurchaseItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetUnitAmount returns the UnitAmount field value
func (o *PurchaseItem) GetUnitAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value
// and a boolean to check if the value has been set.
func (o *PurchaseItem) GetUnitAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitAmount, true
}

// SetUnitAmount sets field value
func (o *PurchaseItem) SetUnitAmount(v int32) {
	o.UnitAmount = v
}

// GetType returns the Type field value
func (o *PurchaseItem) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PurchaseItem) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PurchaseItem) SetType(v int32) {
	o.Type = v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *PurchaseItem) GetTaxAmount() int32 {
	if o == nil || o.TaxAmount == nil {
		var ret int32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseItem) GetTaxAmountOk() (*int32, bool) {
	if o == nil || o.TaxAmount == nil {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *PurchaseItem) HasTaxAmount() bool {
	if o != nil && o.TaxAmount != nil {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given int32 and assigns it to the TaxAmount field.
func (o *PurchaseItem) SetTaxAmount(v int32) {
	o.TaxAmount = &v
}

func (o PurchaseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["merchantAccountId"] = o.MerchantAccountId
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["unitAmount"] = o.UnitAmount
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.TaxAmount != nil {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	return json.Marshal(toSerialize)
}

type NullablePurchaseItem struct {
	value *PurchaseItem
	isSet bool
}

func (v NullablePurchaseItem) Get() *PurchaseItem {
	return v.value
}

func (v *NullablePurchaseItem) Set(val *PurchaseItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseItem(val *PurchaseItem) *NullablePurchaseItem {
	return &NullablePurchaseItem{value: val, isSet: true}
}

func (v NullablePurchaseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


