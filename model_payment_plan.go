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

// PaymentPlan struct for PaymentPlan
type PaymentPlan struct {
	// Payment plan ID.
	Id *int64 `json:"id,omitempty"`
	// PSP name.
	Provider *string `json:"provider,omitempty"`
	// Payment Plan Type.  ex: Installments or deferrerd.
	Type *string `json:"type,omitempty"`
	// Number of deferred days.
	DeferredDays *int32 `json:"deferredDays,omitempty"`
	// Number of installments.
	Installments *int32 `json:"installments,omitempty"`
	// Payment plan description.  ex: The marketplace bears the fees.
	Description *string `json:"description,omitempty"`
}

// NewPaymentPlan instantiates a new PaymentPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentPlan() *PaymentPlan {
	this := PaymentPlan{}
	return &this
}

// NewPaymentPlanWithDefaults instantiates a new PaymentPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentPlanWithDefaults() *PaymentPlan {
	this := PaymentPlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentPlan) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPlan) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentPlan) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PaymentPlan) SetId(v int64) {
	o.Id = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PaymentPlan) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPlan) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PaymentPlan) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *PaymentPlan) SetProvider(v string) {
	o.Provider = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentPlan) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPlan) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentPlan) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentPlan) SetType(v string) {
	o.Type = &v
}

// GetDeferredDays returns the DeferredDays field value if set, zero value otherwise.
func (o *PaymentPlan) GetDeferredDays() int32 {
	if o == nil || o.DeferredDays == nil {
		var ret int32
		return ret
	}
	return *o.DeferredDays
}

// GetDeferredDaysOk returns a tuple with the DeferredDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPlan) GetDeferredDaysOk() (*int32, bool) {
	if o == nil || o.DeferredDays == nil {
		return nil, false
	}
	return o.DeferredDays, true
}

// HasDeferredDays returns a boolean if a field has been set.
func (o *PaymentPlan) HasDeferredDays() bool {
	if o != nil && o.DeferredDays != nil {
		return true
	}

	return false
}

// SetDeferredDays gets a reference to the given int32 and assigns it to the DeferredDays field.
func (o *PaymentPlan) SetDeferredDays(v int32) {
	o.DeferredDays = &v
}

// GetInstallments returns the Installments field value if set, zero value otherwise.
func (o *PaymentPlan) GetInstallments() int32 {
	if o == nil || o.Installments == nil {
		var ret int32
		return ret
	}
	return *o.Installments
}

// GetInstallmentsOk returns a tuple with the Installments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPlan) GetInstallmentsOk() (*int32, bool) {
	if o == nil || o.Installments == nil {
		return nil, false
	}
	return o.Installments, true
}

// HasInstallments returns a boolean if a field has been set.
func (o *PaymentPlan) HasInstallments() bool {
	if o != nil && o.Installments != nil {
		return true
	}

	return false
}

// SetInstallments gets a reference to the given int32 and assigns it to the Installments field.
func (o *PaymentPlan) SetInstallments(v int32) {
	o.Installments = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentPlan) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPlan) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentPlan) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentPlan) SetDescription(v string) {
	o.Description = &v
}

func (o PaymentPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DeferredDays != nil {
		toSerialize["deferredDays"] = o.DeferredDays
	}
	if o.Installments != nil {
		toSerialize["installments"] = o.Installments
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentPlan struct {
	value *PaymentPlan
	isSet bool
}

func (v NullablePaymentPlan) Get() *PaymentPlan {
	return v.value
}

func (v *NullablePaymentPlan) Set(val *PaymentPlan) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentPlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentPlan(val *PaymentPlan) *NullablePaymentPlan {
	return &NullablePaymentPlan{value: val, isSet: true}
}

func (v NullablePaymentPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


