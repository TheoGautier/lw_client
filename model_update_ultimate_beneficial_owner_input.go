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

// UpdateUltimateBeneficialOwnerInput struct for UpdateUltimateBeneficialOwnerInput
type UpdateUltimateBeneficialOwnerInput struct {
	// Client First Name
	FirstName string `json:"firstName"`
	// Client Last Name
	LastName string `json:"lastName"`
	// Client Nationality (use ISO 3166-1 alpha-3 format)
	Nationality *string `json:"nationality,omitempty"`
	// Client Birthdate
	DateOfBirth string `json:"dateOfBirth"`
	// Client City of Birth
	CityOfBirth string `json:"cityOfBirth"`
	// Client Country of Birth (use ISO 3166-1 alpha-3 format)
	CountryOfBirth string `json:"countryOfBirth"`
	// Country of Residence (use ISO 3166-1 alpha-3 format)
	CountryOfResidence *string `json:"countryOfResidence,omitempty"`
	// Indicate if the UBO is active or not
	IsActive *bool `json:"isActive,omitempty"`
}

// NewUpdateUltimateBeneficialOwnerInput instantiates a new UpdateUltimateBeneficialOwnerInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUltimateBeneficialOwnerInput(firstName string, lastName string, dateOfBirth string, cityOfBirth string, countryOfBirth string) *UpdateUltimateBeneficialOwnerInput {
	this := UpdateUltimateBeneficialOwnerInput{}
	this.FirstName = firstName
	this.LastName = lastName
	this.DateOfBirth = dateOfBirth
	this.CityOfBirth = cityOfBirth
	this.CountryOfBirth = countryOfBirth
	return &this
}

// NewUpdateUltimateBeneficialOwnerInputWithDefaults instantiates a new UpdateUltimateBeneficialOwnerInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUltimateBeneficialOwnerInputWithDefaults() *UpdateUltimateBeneficialOwnerInput {
	this := UpdateUltimateBeneficialOwnerInput{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *UpdateUltimateBeneficialOwnerInput) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UpdateUltimateBeneficialOwnerInput) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UpdateUltimateBeneficialOwnerInput) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UpdateUltimateBeneficialOwnerInput) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UpdateUltimateBeneficialOwnerInput) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UpdateUltimateBeneficialOwnerInput) SetLastName(v string) {
	o.LastName = v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *UpdateUltimateBeneficialOwnerInput) GetNationality() string {
	if o == nil || o.Nationality == nil {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUltimateBeneficialOwnerInput) GetNationalityOk() (*string, bool) {
	if o == nil || o.Nationality == nil {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *UpdateUltimateBeneficialOwnerInput) HasNationality() bool {
	if o != nil && o.Nationality != nil {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *UpdateUltimateBeneficialOwnerInput) SetNationality(v string) {
	o.Nationality = &v
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *UpdateUltimateBeneficialOwnerInput) GetDateOfBirth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *UpdateUltimateBeneficialOwnerInput) GetDateOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *UpdateUltimateBeneficialOwnerInput) SetDateOfBirth(v string) {
	o.DateOfBirth = v
}

// GetCityOfBirth returns the CityOfBirth field value
func (o *UpdateUltimateBeneficialOwnerInput) GetCityOfBirth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CityOfBirth
}

// GetCityOfBirthOk returns a tuple with the CityOfBirth field value
// and a boolean to check if the value has been set.
func (o *UpdateUltimateBeneficialOwnerInput) GetCityOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CityOfBirth, true
}

// SetCityOfBirth sets field value
func (o *UpdateUltimateBeneficialOwnerInput) SetCityOfBirth(v string) {
	o.CityOfBirth = v
}

// GetCountryOfBirth returns the CountryOfBirth field value
func (o *UpdateUltimateBeneficialOwnerInput) GetCountryOfBirth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryOfBirth
}

// GetCountryOfBirthOk returns a tuple with the CountryOfBirth field value
// and a boolean to check if the value has been set.
func (o *UpdateUltimateBeneficialOwnerInput) GetCountryOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryOfBirth, true
}

// SetCountryOfBirth sets field value
func (o *UpdateUltimateBeneficialOwnerInput) SetCountryOfBirth(v string) {
	o.CountryOfBirth = v
}

// GetCountryOfResidence returns the CountryOfResidence field value if set, zero value otherwise.
func (o *UpdateUltimateBeneficialOwnerInput) GetCountryOfResidence() string {
	if o == nil || o.CountryOfResidence == nil {
		var ret string
		return ret
	}
	return *o.CountryOfResidence
}

// GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUltimateBeneficialOwnerInput) GetCountryOfResidenceOk() (*string, bool) {
	if o == nil || o.CountryOfResidence == nil {
		return nil, false
	}
	return o.CountryOfResidence, true
}

// HasCountryOfResidence returns a boolean if a field has been set.
func (o *UpdateUltimateBeneficialOwnerInput) HasCountryOfResidence() bool {
	if o != nil && o.CountryOfResidence != nil {
		return true
	}

	return false
}

// SetCountryOfResidence gets a reference to the given string and assigns it to the CountryOfResidence field.
func (o *UpdateUltimateBeneficialOwnerInput) SetCountryOfResidence(v string) {
	o.CountryOfResidence = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UpdateUltimateBeneficialOwnerInput) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUltimateBeneficialOwnerInput) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateUltimateBeneficialOwnerInput) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UpdateUltimateBeneficialOwnerInput) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o UpdateUltimateBeneficialOwnerInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if o.Nationality != nil {
		toSerialize["nationality"] = o.Nationality
	}
	if true {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if true {
		toSerialize["cityOfBirth"] = o.CityOfBirth
	}
	if true {
		toSerialize["countryOfBirth"] = o.CountryOfBirth
	}
	if o.CountryOfResidence != nil {
		toSerialize["countryOfResidence"] = o.CountryOfResidence
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateUltimateBeneficialOwnerInput struct {
	value *UpdateUltimateBeneficialOwnerInput
	isSet bool
}

func (v NullableUpdateUltimateBeneficialOwnerInput) Get() *UpdateUltimateBeneficialOwnerInput {
	return v.value
}

func (v *NullableUpdateUltimateBeneficialOwnerInput) Set(val *UpdateUltimateBeneficialOwnerInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUltimateBeneficialOwnerInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUltimateBeneficialOwnerInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUltimateBeneficialOwnerInput(val *UpdateUltimateBeneficialOwnerInput) *NullableUpdateUltimateBeneficialOwnerInput {
	return &NullableUpdateUltimateBeneficialOwnerInput{value: val, isSet: true}
}

func (v NullableUpdateUltimateBeneficialOwnerInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUltimateBeneficialOwnerInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


