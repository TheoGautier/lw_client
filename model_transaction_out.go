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

// TransactionOut struct for TransactionOut
type TransactionOut struct {
	// Debited Account
	SenderAccountId *string `json:"senderAccountId,omitempty"`
	// Debited Amount  Amounts are represented in cents (Euros)
	DebitAmount *int32 `json:"debitAmount,omitempty"`
	// IBAN ID
	IbanId *int64 `json:"IbanId,omitempty"`
	MaskedLabel *string `json:"maskedLabel,omitempty"`
	BankStatus *string `json:"bankStatus,omitempty"`
	PSP *PSP `json:"PSP,omitempty"`
	// Money-In - ID responsible for the chargeback
	OriginId *int64 `json:"originId,omitempty"`
	BuyNowPayLaterInfo *BnplInfo `json:"buyNowPayLaterInfo,omitempty"`
	// Transaction ID
	Id *int64 `json:"id,omitempty"`
	// Payment Method    0 Card      1 Bank transfer (MoneyIn)       3 Bank transfer (MoneyOut)       4 P2P      13 iDEAL      14 SEPA DirectDebit      15 Cheque      17 Sofort      19 Multibanco      21 MBWAY      23 Pagare        30 BNPL    35 PayPal     ---   **Important:** The following services have been discontinued.      16 Neosurf    18 PFS Physical Card    20 Payshop    22 Polish Instant Transfer     24 MobilePay  25 Paytrail  26 WeChat    27 P24    28 MoneyIn by TPE    29 Trustly
	Method *int32 `json:"method,omitempty"`
	// Payment Method Details  0 Standard  1 Pay By Bank<br/>0 = STANDARD.<br/>1 = PAY_BY_BANK.<br/>
	MethodDetails *int32 `json:"methodDetails,omitempty"`
	// Transaction initialization date, UTC Unix timestamp
	Date *int32 `json:"date,omitempty"`
	// Your fee  Amounts are given as integer numbers in cents
	CommissionAmount *int32 `json:"commissionAmount,omitempty"`
	// Comment
	Comment *string `json:"comment,omitempty"`
	// **Money-In and Money-Out**         0: Success    3: Lemonway Error    4: Pending    6: PSP error    7: Cancelled    16: Validation pending      **Important:**  For /v2/moneyins/{transactionid}/validate` use the status codes:      0: Waiting for finalization      3: Success      4: Lemonway Error      6: PSP Error    7: Cancelled    16: Validation Pending                **P2P**    0: Pending    3: Success
	Status *int32 `json:"status,omitempty"`
	// Transaction execution date
	ExecutionDate *int32 `json:"executionDate,omitempty"`
	LemonWayCommission *LemonWayCommission `json:"lemonWayCommission,omitempty"`
	// Unique ID generated by your server
	Reference *string `json:"reference,omitempty"`
}

// NewTransactionOut instantiates a new TransactionOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionOut() *TransactionOut {
	this := TransactionOut{}
	return &this
}

// NewTransactionOutWithDefaults instantiates a new TransactionOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionOutWithDefaults() *TransactionOut {
	this := TransactionOut{}
	return &this
}

// GetSenderAccountId returns the SenderAccountId field value if set, zero value otherwise.
func (o *TransactionOut) GetSenderAccountId() string {
	if o == nil || o.SenderAccountId == nil {
		var ret string
		return ret
	}
	return *o.SenderAccountId
}

// GetSenderAccountIdOk returns a tuple with the SenderAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetSenderAccountIdOk() (*string, bool) {
	if o == nil || o.SenderAccountId == nil {
		return nil, false
	}
	return o.SenderAccountId, true
}

// HasSenderAccountId returns a boolean if a field has been set.
func (o *TransactionOut) HasSenderAccountId() bool {
	if o != nil && o.SenderAccountId != nil {
		return true
	}

	return false
}

// SetSenderAccountId gets a reference to the given string and assigns it to the SenderAccountId field.
func (o *TransactionOut) SetSenderAccountId(v string) {
	o.SenderAccountId = &v
}

// GetDebitAmount returns the DebitAmount field value if set, zero value otherwise.
func (o *TransactionOut) GetDebitAmount() int32 {
	if o == nil || o.DebitAmount == nil {
		var ret int32
		return ret
	}
	return *o.DebitAmount
}

// GetDebitAmountOk returns a tuple with the DebitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetDebitAmountOk() (*int32, bool) {
	if o == nil || o.DebitAmount == nil {
		return nil, false
	}
	return o.DebitAmount, true
}

// HasDebitAmount returns a boolean if a field has been set.
func (o *TransactionOut) HasDebitAmount() bool {
	if o != nil && o.DebitAmount != nil {
		return true
	}

	return false
}

// SetDebitAmount gets a reference to the given int32 and assigns it to the DebitAmount field.
func (o *TransactionOut) SetDebitAmount(v int32) {
	o.DebitAmount = &v
}

// GetIbanId returns the IbanId field value if set, zero value otherwise.
func (o *TransactionOut) GetIbanId() int64 {
	if o == nil || o.IbanId == nil {
		var ret int64
		return ret
	}
	return *o.IbanId
}

// GetIbanIdOk returns a tuple with the IbanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetIbanIdOk() (*int64, bool) {
	if o == nil || o.IbanId == nil {
		return nil, false
	}
	return o.IbanId, true
}

// HasIbanId returns a boolean if a field has been set.
func (o *TransactionOut) HasIbanId() bool {
	if o != nil && o.IbanId != nil {
		return true
	}

	return false
}

// SetIbanId gets a reference to the given int64 and assigns it to the IbanId field.
func (o *TransactionOut) SetIbanId(v int64) {
	o.IbanId = &v
}

// GetMaskedLabel returns the MaskedLabel field value if set, zero value otherwise.
func (o *TransactionOut) GetMaskedLabel() string {
	if o == nil || o.MaskedLabel == nil {
		var ret string
		return ret
	}
	return *o.MaskedLabel
}

// GetMaskedLabelOk returns a tuple with the MaskedLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetMaskedLabelOk() (*string, bool) {
	if o == nil || o.MaskedLabel == nil {
		return nil, false
	}
	return o.MaskedLabel, true
}

// HasMaskedLabel returns a boolean if a field has been set.
func (o *TransactionOut) HasMaskedLabel() bool {
	if o != nil && o.MaskedLabel != nil {
		return true
	}

	return false
}

// SetMaskedLabel gets a reference to the given string and assigns it to the MaskedLabel field.
func (o *TransactionOut) SetMaskedLabel(v string) {
	o.MaskedLabel = &v
}

// GetBankStatus returns the BankStatus field value if set, zero value otherwise.
func (o *TransactionOut) GetBankStatus() string {
	if o == nil || o.BankStatus == nil {
		var ret string
		return ret
	}
	return *o.BankStatus
}

// GetBankStatusOk returns a tuple with the BankStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetBankStatusOk() (*string, bool) {
	if o == nil || o.BankStatus == nil {
		return nil, false
	}
	return o.BankStatus, true
}

// HasBankStatus returns a boolean if a field has been set.
func (o *TransactionOut) HasBankStatus() bool {
	if o != nil && o.BankStatus != nil {
		return true
	}

	return false
}

// SetBankStatus gets a reference to the given string and assigns it to the BankStatus field.
func (o *TransactionOut) SetBankStatus(v string) {
	o.BankStatus = &v
}

// GetPSP returns the PSP field value if set, zero value otherwise.
func (o *TransactionOut) GetPSP() PSP {
	if o == nil || o.PSP == nil {
		var ret PSP
		return ret
	}
	return *o.PSP
}

// GetPSPOk returns a tuple with the PSP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetPSPOk() (*PSP, bool) {
	if o == nil || o.PSP == nil {
		return nil, false
	}
	return o.PSP, true
}

// HasPSP returns a boolean if a field has been set.
func (o *TransactionOut) HasPSP() bool {
	if o != nil && o.PSP != nil {
		return true
	}

	return false
}

// SetPSP gets a reference to the given PSP and assigns it to the PSP field.
func (o *TransactionOut) SetPSP(v PSP) {
	o.PSP = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *TransactionOut) GetOriginId() int64 {
	if o == nil || o.OriginId == nil {
		var ret int64
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetOriginIdOk() (*int64, bool) {
	if o == nil || o.OriginId == nil {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *TransactionOut) HasOriginId() bool {
	if o != nil && o.OriginId != nil {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given int64 and assigns it to the OriginId field.
func (o *TransactionOut) SetOriginId(v int64) {
	o.OriginId = &v
}

// GetBuyNowPayLaterInfo returns the BuyNowPayLaterInfo field value if set, zero value otherwise.
func (o *TransactionOut) GetBuyNowPayLaterInfo() BnplInfo {
	if o == nil || o.BuyNowPayLaterInfo == nil {
		var ret BnplInfo
		return ret
	}
	return *o.BuyNowPayLaterInfo
}

// GetBuyNowPayLaterInfoOk returns a tuple with the BuyNowPayLaterInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetBuyNowPayLaterInfoOk() (*BnplInfo, bool) {
	if o == nil || o.BuyNowPayLaterInfo == nil {
		return nil, false
	}
	return o.BuyNowPayLaterInfo, true
}

// HasBuyNowPayLaterInfo returns a boolean if a field has been set.
func (o *TransactionOut) HasBuyNowPayLaterInfo() bool {
	if o != nil && o.BuyNowPayLaterInfo != nil {
		return true
	}

	return false
}

// SetBuyNowPayLaterInfo gets a reference to the given BnplInfo and assigns it to the BuyNowPayLaterInfo field.
func (o *TransactionOut) SetBuyNowPayLaterInfo(v BnplInfo) {
	o.BuyNowPayLaterInfo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionOut) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionOut) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TransactionOut) SetId(v int64) {
	o.Id = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *TransactionOut) GetMethod() int32 {
	if o == nil || o.Method == nil {
		var ret int32
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetMethodOk() (*int32, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *TransactionOut) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given int32 and assigns it to the Method field.
func (o *TransactionOut) SetMethod(v int32) {
	o.Method = &v
}

// GetMethodDetails returns the MethodDetails field value if set, zero value otherwise.
func (o *TransactionOut) GetMethodDetails() int32 {
	if o == nil || o.MethodDetails == nil {
		var ret int32
		return ret
	}
	return *o.MethodDetails
}

// GetMethodDetailsOk returns a tuple with the MethodDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetMethodDetailsOk() (*int32, bool) {
	if o == nil || o.MethodDetails == nil {
		return nil, false
	}
	return o.MethodDetails, true
}

// HasMethodDetails returns a boolean if a field has been set.
func (o *TransactionOut) HasMethodDetails() bool {
	if o != nil && o.MethodDetails != nil {
		return true
	}

	return false
}

// SetMethodDetails gets a reference to the given int32 and assigns it to the MethodDetails field.
func (o *TransactionOut) SetMethodDetails(v int32) {
	o.MethodDetails = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TransactionOut) GetDate() int32 {
	if o == nil || o.Date == nil {
		var ret int32
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetDateOk() (*int32, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TransactionOut) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given int32 and assigns it to the Date field.
func (o *TransactionOut) SetDate(v int32) {
	o.Date = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *TransactionOut) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *TransactionOut) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *TransactionOut) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TransactionOut) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TransactionOut) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TransactionOut) SetComment(v string) {
	o.Comment = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransactionOut) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransactionOut) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *TransactionOut) SetStatus(v int32) {
	o.Status = &v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *TransactionOut) GetExecutionDate() int32 {
	if o == nil || o.ExecutionDate == nil {
		var ret int32
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetExecutionDateOk() (*int32, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *TransactionOut) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given int32 and assigns it to the ExecutionDate field.
func (o *TransactionOut) SetExecutionDate(v int32) {
	o.ExecutionDate = &v
}

// GetLemonWayCommission returns the LemonWayCommission field value if set, zero value otherwise.
func (o *TransactionOut) GetLemonWayCommission() LemonWayCommission {
	if o == nil || o.LemonWayCommission == nil {
		var ret LemonWayCommission
		return ret
	}
	return *o.LemonWayCommission
}

// GetLemonWayCommissionOk returns a tuple with the LemonWayCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetLemonWayCommissionOk() (*LemonWayCommission, bool) {
	if o == nil || o.LemonWayCommission == nil {
		return nil, false
	}
	return o.LemonWayCommission, true
}

// HasLemonWayCommission returns a boolean if a field has been set.
func (o *TransactionOut) HasLemonWayCommission() bool {
	if o != nil && o.LemonWayCommission != nil {
		return true
	}

	return false
}

// SetLemonWayCommission gets a reference to the given LemonWayCommission and assigns it to the LemonWayCommission field.
func (o *TransactionOut) SetLemonWayCommission(v LemonWayCommission) {
	o.LemonWayCommission = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *TransactionOut) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOut) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *TransactionOut) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *TransactionOut) SetReference(v string) {
	o.Reference = &v
}

func (o TransactionOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SenderAccountId != nil {
		toSerialize["senderAccountId"] = o.SenderAccountId
	}
	if o.DebitAmount != nil {
		toSerialize["debitAmount"] = o.DebitAmount
	}
	if o.IbanId != nil {
		toSerialize["IbanId"] = o.IbanId
	}
	if o.MaskedLabel != nil {
		toSerialize["maskedLabel"] = o.MaskedLabel
	}
	if o.BankStatus != nil {
		toSerialize["bankStatus"] = o.BankStatus
	}
	if o.PSP != nil {
		toSerialize["PSP"] = o.PSP
	}
	if o.OriginId != nil {
		toSerialize["originId"] = o.OriginId
	}
	if o.BuyNowPayLaterInfo != nil {
		toSerialize["buyNowPayLaterInfo"] = o.BuyNowPayLaterInfo
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.MethodDetails != nil {
		toSerialize["methodDetails"] = o.MethodDetails
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.CommissionAmount != nil {
		toSerialize["commissionAmount"] = o.CommissionAmount
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ExecutionDate != nil {
		toSerialize["executionDate"] = o.ExecutionDate
	}
	if o.LemonWayCommission != nil {
		toSerialize["lemonWayCommission"] = o.LemonWayCommission
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionOut struct {
	value *TransactionOut
	isSet bool
}

func (v NullableTransactionOut) Get() *TransactionOut {
	return v.value
}

func (v *NullableTransactionOut) Set(val *TransactionOut) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionOut) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionOut(val *TransactionOut) *NullableTransactionOut {
	return &NullableTransactionOut{value: val, isSet: true}
}

func (v NullableTransactionOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

