# TransactionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderAccountId** | Pointer to **string** | Debited Account | [optional] 
**DebitAmount** | Pointer to **int32** | Debited Amount  Amounts are represented in cents (Euros) | [optional] 
**IbanId** | Pointer to **int64** | IBAN ID | [optional] 
**MaskedLabel** | Pointer to **string** |  | [optional] 
**BankStatus** | Pointer to **string** |  | [optional] 
**PSP** | Pointer to [**PSP**](PSP.md) |  | [optional] 
**OriginId** | Pointer to **int64** | Money-In - ID responsible for the chargeback | [optional] 
**BuyNowPayLaterInfo** | Pointer to [**BnplInfo**](BnplInfo.md) |  | [optional] 
**Id** | Pointer to **int64** | Transaction ID | [optional] 
**Method** | Pointer to **int32** | Payment Method    0 Card      1 Bank transfer (MoneyIn)       3 Bank transfer (MoneyOut)       4 P2P      13 iDEAL      14 SEPA DirectDebit      15 Cheque      17 Sofort      19 Multibanco      21 MBWAY      23 Pagare        30 BNPL    35 PayPal     ---   **Important:** The following services have been discontinued.      16 Neosurf    18 PFS Physical Card    20 Payshop    22 Polish Instant Transfer     24 MobilePay  25 Paytrail  26 WeChat    27 P24    28 MoneyIn by TPE    29 Trustly | [optional] 
**MethodDetails** | Pointer to **int32** | Payment Method Details  0 Standard  1 Pay By Bank&lt;br/&gt;0 &#x3D; STANDARD.&lt;br/&gt;1 &#x3D; PAY_BY_BANK.&lt;br/&gt; | [optional] 
**Date** | Pointer to **int32** | Transaction initialization date, UTC Unix timestamp | [optional] 
**CommissionAmount** | Pointer to **int32** | Your fee  Amounts are given as integer numbers in cents | [optional] 
**Comment** | Pointer to **string** | Comment | [optional] 
**Status** | Pointer to **int32** | **Money-In and Money-Out**         0: Success    3: Lemonway Error    4: Pending    6: PSP error    7: Cancelled    16: Validation pending      **Important:**  For /v2/moneyins/{transactionid}/validate&#x60; use the status codes:      0: Waiting for finalization      3: Success      4: Lemonway Error      6: PSP Error    7: Cancelled    16: Validation Pending                **P2P**    0: Pending    3: Success | [optional] 
**ExecutionDate** | Pointer to **int32** | Transaction execution date | [optional] 
**LemonWayCommission** | Pointer to [**LemonWayCommission**](LemonWayCommission.md) |  | [optional] 
**Reference** | Pointer to **string** | Unique ID generated by your server | [optional] 

## Methods

### NewTransactionOut

`func NewTransactionOut() *TransactionOut`

NewTransactionOut instantiates a new TransactionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionOutWithDefaults

`func NewTransactionOutWithDefaults() *TransactionOut`

NewTransactionOutWithDefaults instantiates a new TransactionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderAccountId

`func (o *TransactionOut) GetSenderAccountId() string`

GetSenderAccountId returns the SenderAccountId field if non-nil, zero value otherwise.

### GetSenderAccountIdOk

`func (o *TransactionOut) GetSenderAccountIdOk() (*string, bool)`

GetSenderAccountIdOk returns a tuple with the SenderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccountId

`func (o *TransactionOut) SetSenderAccountId(v string)`

SetSenderAccountId sets SenderAccountId field to given value.

### HasSenderAccountId

`func (o *TransactionOut) HasSenderAccountId() bool`

HasSenderAccountId returns a boolean if a field has been set.

### GetDebitAmount

`func (o *TransactionOut) GetDebitAmount() int32`

GetDebitAmount returns the DebitAmount field if non-nil, zero value otherwise.

### GetDebitAmountOk

`func (o *TransactionOut) GetDebitAmountOk() (*int32, bool)`

GetDebitAmountOk returns a tuple with the DebitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAmount

`func (o *TransactionOut) SetDebitAmount(v int32)`

SetDebitAmount sets DebitAmount field to given value.

### HasDebitAmount

`func (o *TransactionOut) HasDebitAmount() bool`

HasDebitAmount returns a boolean if a field has been set.

### GetIbanId

`func (o *TransactionOut) GetIbanId() int64`

GetIbanId returns the IbanId field if non-nil, zero value otherwise.

### GetIbanIdOk

`func (o *TransactionOut) GetIbanIdOk() (*int64, bool)`

GetIbanIdOk returns a tuple with the IbanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanId

`func (o *TransactionOut) SetIbanId(v int64)`

SetIbanId sets IbanId field to given value.

### HasIbanId

`func (o *TransactionOut) HasIbanId() bool`

HasIbanId returns a boolean if a field has been set.

### GetMaskedLabel

`func (o *TransactionOut) GetMaskedLabel() string`

GetMaskedLabel returns the MaskedLabel field if non-nil, zero value otherwise.

### GetMaskedLabelOk

`func (o *TransactionOut) GetMaskedLabelOk() (*string, bool)`

GetMaskedLabelOk returns a tuple with the MaskedLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedLabel

`func (o *TransactionOut) SetMaskedLabel(v string)`

SetMaskedLabel sets MaskedLabel field to given value.

### HasMaskedLabel

`func (o *TransactionOut) HasMaskedLabel() bool`

HasMaskedLabel returns a boolean if a field has been set.

### GetBankStatus

`func (o *TransactionOut) GetBankStatus() string`

GetBankStatus returns the BankStatus field if non-nil, zero value otherwise.

### GetBankStatusOk

`func (o *TransactionOut) GetBankStatusOk() (*string, bool)`

GetBankStatusOk returns a tuple with the BankStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankStatus

`func (o *TransactionOut) SetBankStatus(v string)`

SetBankStatus sets BankStatus field to given value.

### HasBankStatus

`func (o *TransactionOut) HasBankStatus() bool`

HasBankStatus returns a boolean if a field has been set.

### GetPSP

`func (o *TransactionOut) GetPSP() PSP`

GetPSP returns the PSP field if non-nil, zero value otherwise.

### GetPSPOk

`func (o *TransactionOut) GetPSPOk() (*PSP, bool)`

GetPSPOk returns a tuple with the PSP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSP

`func (o *TransactionOut) SetPSP(v PSP)`

SetPSP sets PSP field to given value.

### HasPSP

`func (o *TransactionOut) HasPSP() bool`

HasPSP returns a boolean if a field has been set.

### GetOriginId

`func (o *TransactionOut) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *TransactionOut) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *TransactionOut) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *TransactionOut) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetBuyNowPayLaterInfo

`func (o *TransactionOut) GetBuyNowPayLaterInfo() BnplInfo`

GetBuyNowPayLaterInfo returns the BuyNowPayLaterInfo field if non-nil, zero value otherwise.

### GetBuyNowPayLaterInfoOk

`func (o *TransactionOut) GetBuyNowPayLaterInfoOk() (*BnplInfo, bool)`

GetBuyNowPayLaterInfoOk returns a tuple with the BuyNowPayLaterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyNowPayLaterInfo

`func (o *TransactionOut) SetBuyNowPayLaterInfo(v BnplInfo)`

SetBuyNowPayLaterInfo sets BuyNowPayLaterInfo field to given value.

### HasBuyNowPayLaterInfo

`func (o *TransactionOut) HasBuyNowPayLaterInfo() bool`

HasBuyNowPayLaterInfo returns a boolean if a field has been set.

### GetId

`func (o *TransactionOut) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionOut) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionOut) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionOut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *TransactionOut) GetMethod() int32`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionOut) GetMethodOk() (*int32, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionOut) SetMethod(v int32)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TransactionOut) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodDetails

`func (o *TransactionOut) GetMethodDetails() int32`

GetMethodDetails returns the MethodDetails field if non-nil, zero value otherwise.

### GetMethodDetailsOk

`func (o *TransactionOut) GetMethodDetailsOk() (*int32, bool)`

GetMethodDetailsOk returns a tuple with the MethodDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodDetails

`func (o *TransactionOut) SetMethodDetails(v int32)`

SetMethodDetails sets MethodDetails field to given value.

### HasMethodDetails

`func (o *TransactionOut) HasMethodDetails() bool`

HasMethodDetails returns a boolean if a field has been set.

### GetDate

`func (o *TransactionOut) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionOut) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionOut) SetDate(v int32)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionOut) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCommissionAmount

`func (o *TransactionOut) GetCommissionAmount() int32`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *TransactionOut) GetCommissionAmountOk() (*int32, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *TransactionOut) SetCommissionAmount(v int32)`

SetCommissionAmount sets CommissionAmount field to given value.

### HasCommissionAmount

`func (o *TransactionOut) HasCommissionAmount() bool`

HasCommissionAmount returns a boolean if a field has been set.

### GetComment

`func (o *TransactionOut) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TransactionOut) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TransactionOut) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TransactionOut) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionOut) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionOut) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionOut) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExecutionDate

`func (o *TransactionOut) GetExecutionDate() int32`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *TransactionOut) GetExecutionDateOk() (*int32, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *TransactionOut) SetExecutionDate(v int32)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *TransactionOut) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetLemonWayCommission

`func (o *TransactionOut) GetLemonWayCommission() LemonWayCommission`

GetLemonWayCommission returns the LemonWayCommission field if non-nil, zero value otherwise.

### GetLemonWayCommissionOk

`func (o *TransactionOut) GetLemonWayCommissionOk() (*LemonWayCommission, bool)`

GetLemonWayCommissionOk returns a tuple with the LemonWayCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLemonWayCommission

`func (o *TransactionOut) SetLemonWayCommission(v LemonWayCommission)`

SetLemonWayCommission sets LemonWayCommission field to given value.

### HasLemonWayCommission

`func (o *TransactionOut) HasLemonWayCommission() bool`

HasLemonWayCommission returns a boolean if a field has been set.

### GetReference

`func (o *TransactionOut) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TransactionOut) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TransactionOut) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TransactionOut) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

