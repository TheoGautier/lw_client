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

// CreatePaymentFormInput struct for CreatePaymentFormInput
type CreatePaymentFormInput struct {
	// To cross-reference a buyer with a payment, use `optId` to help associate them with one another.     **Note:** Do not use special characters in this field, for example: #, @, and !.
	OptId *string `json:"optId,omitempty"`
	// Associated with the **Payer Wallet**.    If this field is filled then the money reaches this wallet before arriving at the beneficiary wallet (via a transfer wallet to wallet).  **Important:** Don't put the <b>SC Wallet</b> here, it won't work. You cannot debit the <b>SC Wallet</b> with a credit card.
	WalletPayer *string `json:"walletPayer,omitempty"`
	// Beneficiary Wallet. If this field is not complete, then the end-user must to fill its value on the payment form.   <b>We recommend you complete the form instead of your end-user</b>  Don't put the <b>SC Wallet</b> here, it won't work. You cannot credit the <b>SC Wallet</b> with a credit card.
	WalletReceiver *string `json:"walletReceiver,omitempty"`
	// Amount or a range of the amount to be debited.  1. If this field is configured with an interval (eg, 15.30-500.26) then the final customer will have to enter an appropriate amount in the form.  2. If this field is not filled then the end-user can enter any amount to the form.  3. If this field is filled with a precise value (eg 15.60), then the end-user has no choice in the amount field of the form.  Amounts are given as integer numbers in cents.  Note: If you enter a fixed amount
	TotalAmount *string `json:"totalAmount,omitempty"`
	// Your Fee  Amounts are given as integer numbers in cents
	CommissionAmount *int32 `json:"commissionAmount,omitempty"`
	// Comment Regarding the Transaction
	Comment *string `json:"comment,omitempty"`
	// URL redirection after the payment procedure is successfully finished
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// URL redirection after the payment procedure is cancelled
	CancelUrl *string `json:"cancelUrl,omitempty"`
	// URL edirection after the payment procedure is failed
	ErrorUrl *string `json:"errorUrl,omitempty"`
	// Payer's First Name  If this field is not filled then the end-user have to fill it in the payment form
	FirstNamePayer *string `json:"firstNamePayer,omitempty"`
	// Payer's Last Name  If this field is not filled then the end-user have to fill it in the payment form
	LastNamePayer *string `json:"lastNamePayer,omitempty"`
	// Payer's Email  If this field is not filled then the end-user have to fill it in the payment form
	EmailPayer *string `json:"emailPayer,omitempty"`
	// Link to a custom CSS Stylesheet  The stylesheet should be publicly accessible via **HTTPS*
	Style *string `json:"style,omitempty"`
	AtosStyle *string `json:"atosStyle,omitempty"`
	// At the end of the payment procedure, an HTTP POST message containing the payment status (PAID, ERROR, CANCEL) is sent to this address.   It is possible that the same notification might be sent several times.
	NotifUrl *string `json:"notifUrl,omitempty"`
	// Reserved for a future version
	Options *string `json:"options,omitempty"`
}

// NewCreatePaymentFormInput instantiates a new CreatePaymentFormInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentFormInput() *CreatePaymentFormInput {
	this := CreatePaymentFormInput{}
	return &this
}

// NewCreatePaymentFormInputWithDefaults instantiates a new CreatePaymentFormInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentFormInputWithDefaults() *CreatePaymentFormInput {
	this := CreatePaymentFormInput{}
	return &this
}

// GetOptId returns the OptId field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetOptId() string {
	if o == nil || o.OptId == nil {
		var ret string
		return ret
	}
	return *o.OptId
}

// GetOptIdOk returns a tuple with the OptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetOptIdOk() (*string, bool) {
	if o == nil || o.OptId == nil {
		return nil, false
	}
	return o.OptId, true
}

// HasOptId returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasOptId() bool {
	if o != nil && o.OptId != nil {
		return true
	}

	return false
}

// SetOptId gets a reference to the given string and assigns it to the OptId field.
func (o *CreatePaymentFormInput) SetOptId(v string) {
	o.OptId = &v
}

// GetWalletPayer returns the WalletPayer field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetWalletPayer() string {
	if o == nil || o.WalletPayer == nil {
		var ret string
		return ret
	}
	return *o.WalletPayer
}

// GetWalletPayerOk returns a tuple with the WalletPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetWalletPayerOk() (*string, bool) {
	if o == nil || o.WalletPayer == nil {
		return nil, false
	}
	return o.WalletPayer, true
}

// HasWalletPayer returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasWalletPayer() bool {
	if o != nil && o.WalletPayer != nil {
		return true
	}

	return false
}

// SetWalletPayer gets a reference to the given string and assigns it to the WalletPayer field.
func (o *CreatePaymentFormInput) SetWalletPayer(v string) {
	o.WalletPayer = &v
}

// GetWalletReceiver returns the WalletReceiver field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetWalletReceiver() string {
	if o == nil || o.WalletReceiver == nil {
		var ret string
		return ret
	}
	return *o.WalletReceiver
}

// GetWalletReceiverOk returns a tuple with the WalletReceiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetWalletReceiverOk() (*string, bool) {
	if o == nil || o.WalletReceiver == nil {
		return nil, false
	}
	return o.WalletReceiver, true
}

// HasWalletReceiver returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasWalletReceiver() bool {
	if o != nil && o.WalletReceiver != nil {
		return true
	}

	return false
}

// SetWalletReceiver gets a reference to the given string and assigns it to the WalletReceiver field.
func (o *CreatePaymentFormInput) SetWalletReceiver(v string) {
	o.WalletReceiver = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetTotalAmount() string {
	if o == nil || o.TotalAmount == nil {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetTotalAmountOk() (*string, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *CreatePaymentFormInput) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *CreatePaymentFormInput) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreatePaymentFormInput) SetComment(v string) {
	o.Comment = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetReturnUrl() string {
	if o == nil || o.ReturnUrl == nil {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetReturnUrlOk() (*string, bool) {
	if o == nil || o.ReturnUrl == nil {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasReturnUrl() bool {
	if o != nil && o.ReturnUrl != nil {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *CreatePaymentFormInput) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetCancelUrl() string {
	if o == nil || o.CancelUrl == nil {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetCancelUrlOk() (*string, bool) {
	if o == nil || o.CancelUrl == nil {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasCancelUrl() bool {
	if o != nil && o.CancelUrl != nil {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *CreatePaymentFormInput) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetErrorUrl returns the ErrorUrl field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetErrorUrl() string {
	if o == nil || o.ErrorUrl == nil {
		var ret string
		return ret
	}
	return *o.ErrorUrl
}

// GetErrorUrlOk returns a tuple with the ErrorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetErrorUrlOk() (*string, bool) {
	if o == nil || o.ErrorUrl == nil {
		return nil, false
	}
	return o.ErrorUrl, true
}

// HasErrorUrl returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasErrorUrl() bool {
	if o != nil && o.ErrorUrl != nil {
		return true
	}

	return false
}

// SetErrorUrl gets a reference to the given string and assigns it to the ErrorUrl field.
func (o *CreatePaymentFormInput) SetErrorUrl(v string) {
	o.ErrorUrl = &v
}

// GetFirstNamePayer returns the FirstNamePayer field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetFirstNamePayer() string {
	if o == nil || o.FirstNamePayer == nil {
		var ret string
		return ret
	}
	return *o.FirstNamePayer
}

// GetFirstNamePayerOk returns a tuple with the FirstNamePayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetFirstNamePayerOk() (*string, bool) {
	if o == nil || o.FirstNamePayer == nil {
		return nil, false
	}
	return o.FirstNamePayer, true
}

// HasFirstNamePayer returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasFirstNamePayer() bool {
	if o != nil && o.FirstNamePayer != nil {
		return true
	}

	return false
}

// SetFirstNamePayer gets a reference to the given string and assigns it to the FirstNamePayer field.
func (o *CreatePaymentFormInput) SetFirstNamePayer(v string) {
	o.FirstNamePayer = &v
}

// GetLastNamePayer returns the LastNamePayer field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetLastNamePayer() string {
	if o == nil || o.LastNamePayer == nil {
		var ret string
		return ret
	}
	return *o.LastNamePayer
}

// GetLastNamePayerOk returns a tuple with the LastNamePayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetLastNamePayerOk() (*string, bool) {
	if o == nil || o.LastNamePayer == nil {
		return nil, false
	}
	return o.LastNamePayer, true
}

// HasLastNamePayer returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasLastNamePayer() bool {
	if o != nil && o.LastNamePayer != nil {
		return true
	}

	return false
}

// SetLastNamePayer gets a reference to the given string and assigns it to the LastNamePayer field.
func (o *CreatePaymentFormInput) SetLastNamePayer(v string) {
	o.LastNamePayer = &v
}

// GetEmailPayer returns the EmailPayer field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetEmailPayer() string {
	if o == nil || o.EmailPayer == nil {
		var ret string
		return ret
	}
	return *o.EmailPayer
}

// GetEmailPayerOk returns a tuple with the EmailPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetEmailPayerOk() (*string, bool) {
	if o == nil || o.EmailPayer == nil {
		return nil, false
	}
	return o.EmailPayer, true
}

// HasEmailPayer returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasEmailPayer() bool {
	if o != nil && o.EmailPayer != nil {
		return true
	}

	return false
}

// SetEmailPayer gets a reference to the given string and assigns it to the EmailPayer field.
func (o *CreatePaymentFormInput) SetEmailPayer(v string) {
	o.EmailPayer = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetStyle() string {
	if o == nil || o.Style == nil {
		var ret string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetStyleOk() (*string, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given string and assigns it to the Style field.
func (o *CreatePaymentFormInput) SetStyle(v string) {
	o.Style = &v
}

// GetAtosStyle returns the AtosStyle field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetAtosStyle() string {
	if o == nil || o.AtosStyle == nil {
		var ret string
		return ret
	}
	return *o.AtosStyle
}

// GetAtosStyleOk returns a tuple with the AtosStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetAtosStyleOk() (*string, bool) {
	if o == nil || o.AtosStyle == nil {
		return nil, false
	}
	return o.AtosStyle, true
}

// HasAtosStyle returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasAtosStyle() bool {
	if o != nil && o.AtosStyle != nil {
		return true
	}

	return false
}

// SetAtosStyle gets a reference to the given string and assigns it to the AtosStyle field.
func (o *CreatePaymentFormInput) SetAtosStyle(v string) {
	o.AtosStyle = &v
}

// GetNotifUrl returns the NotifUrl field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetNotifUrl() string {
	if o == nil || o.NotifUrl == nil {
		var ret string
		return ret
	}
	return *o.NotifUrl
}

// GetNotifUrlOk returns a tuple with the NotifUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetNotifUrlOk() (*string, bool) {
	if o == nil || o.NotifUrl == nil {
		return nil, false
	}
	return o.NotifUrl, true
}

// HasNotifUrl returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasNotifUrl() bool {
	if o != nil && o.NotifUrl != nil {
		return true
	}

	return false
}

// SetNotifUrl gets a reference to the given string and assigns it to the NotifUrl field.
func (o *CreatePaymentFormInput) SetNotifUrl(v string) {
	o.NotifUrl = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreatePaymentFormInput) GetOptions() string {
	if o == nil || o.Options == nil {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentFormInput) GetOptionsOk() (*string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreatePaymentFormInput) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *CreatePaymentFormInput) SetOptions(v string) {
	o.Options = &v
}

func (o CreatePaymentFormInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OptId != nil {
		toSerialize["optId"] = o.OptId
	}
	if o.WalletPayer != nil {
		toSerialize["walletPayer"] = o.WalletPayer
	}
	if o.WalletReceiver != nil {
		toSerialize["walletReceiver"] = o.WalletReceiver
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
	if o.ReturnUrl != nil {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if o.CancelUrl != nil {
		toSerialize["cancelUrl"] = o.CancelUrl
	}
	if o.ErrorUrl != nil {
		toSerialize["errorUrl"] = o.ErrorUrl
	}
	if o.FirstNamePayer != nil {
		toSerialize["firstNamePayer"] = o.FirstNamePayer
	}
	if o.LastNamePayer != nil {
		toSerialize["lastNamePayer"] = o.LastNamePayer
	}
	if o.EmailPayer != nil {
		toSerialize["emailPayer"] = o.EmailPayer
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	if o.AtosStyle != nil {
		toSerialize["atosStyle"] = o.AtosStyle
	}
	if o.NotifUrl != nil {
		toSerialize["notifUrl"] = o.NotifUrl
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePaymentFormInput struct {
	value *CreatePaymentFormInput
	isSet bool
}

func (v NullableCreatePaymentFormInput) Get() *CreatePaymentFormInput {
	return v.value
}

func (v *NullableCreatePaymentFormInput) Set(val *CreatePaymentFormInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentFormInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentFormInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentFormInput(val *CreatePaymentFormInput) *NullableCreatePaymentFormInput {
	return &NullableCreatePaymentFormInput{value: val, isSet: true}
}

func (v NullableCreatePaymentFormInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentFormInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

