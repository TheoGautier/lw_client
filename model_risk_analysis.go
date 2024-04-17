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

// RiskAnalysis struct for RiskAnalysis
type RiskAnalysis struct {
	BillingAddress *BillingAddress `json:"billingAddress,omitempty"`
	Holder *Holder `json:"holder,omitempty"`
	DeliveryAddress *DeliveryAddress `json:"deliveryAddress,omitempty"`
	DeliveryAdditionalInfo *DeliveryAdditionalInfo `json:"deliveryAdditionalInfo,omitempty"`
	CustomerAccountInfo *CustomerAccountInfo `json:"customerAccountInfo,omitempty"`
	Authentication *Authentication `json:"authentication,omitempty"`
}

// NewRiskAnalysis instantiates a new RiskAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskAnalysis() *RiskAnalysis {
	this := RiskAnalysis{}
	return &this
}

// NewRiskAnalysisWithDefaults instantiates a new RiskAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskAnalysisWithDefaults() *RiskAnalysis {
	this := RiskAnalysis{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *RiskAnalysis) GetBillingAddress() BillingAddress {
	if o == nil || o.BillingAddress == nil {
		var ret BillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAnalysis) GetBillingAddressOk() (*BillingAddress, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *RiskAnalysis) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given BillingAddress and assigns it to the BillingAddress field.
func (o *RiskAnalysis) SetBillingAddress(v BillingAddress) {
	o.BillingAddress = &v
}

// GetHolder returns the Holder field value if set, zero value otherwise.
func (o *RiskAnalysis) GetHolder() Holder {
	if o == nil || o.Holder == nil {
		var ret Holder
		return ret
	}
	return *o.Holder
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAnalysis) GetHolderOk() (*Holder, bool) {
	if o == nil || o.Holder == nil {
		return nil, false
	}
	return o.Holder, true
}

// HasHolder returns a boolean if a field has been set.
func (o *RiskAnalysis) HasHolder() bool {
	if o != nil && o.Holder != nil {
		return true
	}

	return false
}

// SetHolder gets a reference to the given Holder and assigns it to the Holder field.
func (o *RiskAnalysis) SetHolder(v Holder) {
	o.Holder = &v
}

// GetDeliveryAddress returns the DeliveryAddress field value if set, zero value otherwise.
func (o *RiskAnalysis) GetDeliveryAddress() DeliveryAddress {
	if o == nil || o.DeliveryAddress == nil {
		var ret DeliveryAddress
		return ret
	}
	return *o.DeliveryAddress
}

// GetDeliveryAddressOk returns a tuple with the DeliveryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAnalysis) GetDeliveryAddressOk() (*DeliveryAddress, bool) {
	if o == nil || o.DeliveryAddress == nil {
		return nil, false
	}
	return o.DeliveryAddress, true
}

// HasDeliveryAddress returns a boolean if a field has been set.
func (o *RiskAnalysis) HasDeliveryAddress() bool {
	if o != nil && o.DeliveryAddress != nil {
		return true
	}

	return false
}

// SetDeliveryAddress gets a reference to the given DeliveryAddress and assigns it to the DeliveryAddress field.
func (o *RiskAnalysis) SetDeliveryAddress(v DeliveryAddress) {
	o.DeliveryAddress = &v
}

// GetDeliveryAdditionalInfo returns the DeliveryAdditionalInfo field value if set, zero value otherwise.
func (o *RiskAnalysis) GetDeliveryAdditionalInfo() DeliveryAdditionalInfo {
	if o == nil || o.DeliveryAdditionalInfo == nil {
		var ret DeliveryAdditionalInfo
		return ret
	}
	return *o.DeliveryAdditionalInfo
}

// GetDeliveryAdditionalInfoOk returns a tuple with the DeliveryAdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAnalysis) GetDeliveryAdditionalInfoOk() (*DeliveryAdditionalInfo, bool) {
	if o == nil || o.DeliveryAdditionalInfo == nil {
		return nil, false
	}
	return o.DeliveryAdditionalInfo, true
}

// HasDeliveryAdditionalInfo returns a boolean if a field has been set.
func (o *RiskAnalysis) HasDeliveryAdditionalInfo() bool {
	if o != nil && o.DeliveryAdditionalInfo != nil {
		return true
	}

	return false
}

// SetDeliveryAdditionalInfo gets a reference to the given DeliveryAdditionalInfo and assigns it to the DeliveryAdditionalInfo field.
func (o *RiskAnalysis) SetDeliveryAdditionalInfo(v DeliveryAdditionalInfo) {
	o.DeliveryAdditionalInfo = &v
}

// GetCustomerAccountInfo returns the CustomerAccountInfo field value if set, zero value otherwise.
func (o *RiskAnalysis) GetCustomerAccountInfo() CustomerAccountInfo {
	if o == nil || o.CustomerAccountInfo == nil {
		var ret CustomerAccountInfo
		return ret
	}
	return *o.CustomerAccountInfo
}

// GetCustomerAccountInfoOk returns a tuple with the CustomerAccountInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAnalysis) GetCustomerAccountInfoOk() (*CustomerAccountInfo, bool) {
	if o == nil || o.CustomerAccountInfo == nil {
		return nil, false
	}
	return o.CustomerAccountInfo, true
}

// HasCustomerAccountInfo returns a boolean if a field has been set.
func (o *RiskAnalysis) HasCustomerAccountInfo() bool {
	if o != nil && o.CustomerAccountInfo != nil {
		return true
	}

	return false
}

// SetCustomerAccountInfo gets a reference to the given CustomerAccountInfo and assigns it to the CustomerAccountInfo field.
func (o *RiskAnalysis) SetCustomerAccountInfo(v CustomerAccountInfo) {
	o.CustomerAccountInfo = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *RiskAnalysis) GetAuthentication() Authentication {
	if o == nil || o.Authentication == nil {
		var ret Authentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAnalysis) GetAuthenticationOk() (*Authentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *RiskAnalysis) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given Authentication and assigns it to the Authentication field.
func (o *RiskAnalysis) SetAuthentication(v Authentication) {
	o.Authentication = &v
}

func (o RiskAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingAddress != nil {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if o.Holder != nil {
		toSerialize["holder"] = o.Holder
	}
	if o.DeliveryAddress != nil {
		toSerialize["deliveryAddress"] = o.DeliveryAddress
	}
	if o.DeliveryAdditionalInfo != nil {
		toSerialize["deliveryAdditionalInfo"] = o.DeliveryAdditionalInfo
	}
	if o.CustomerAccountInfo != nil {
		toSerialize["customerAccountInfo"] = o.CustomerAccountInfo
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableRiskAnalysis struct {
	value *RiskAnalysis
	isSet bool
}

func (v NullableRiskAnalysis) Get() *RiskAnalysis {
	return v.value
}

func (v *NullableRiskAnalysis) Set(val *RiskAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskAnalysis(val *RiskAnalysis) *NullableRiskAnalysis {
	return &NullableRiskAnalysis{value: val, isSet: true}
}

func (v NullableRiskAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


