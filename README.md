# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/TheoGautier/lemonway_client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsAdminApi* | [**AccountsAccountSingleGet**](docs/AccountsAdminApi.md#accountsaccountsingleget) | **Get** /v2/accounts/{accountid} | Get Detailed Payment Account Data
*AccountsAdminApi* | [**AccountsBalancesGet**](docs/AccountsAdminApi.md#accountsbalancesget) | **Get** /v2/accounts/balances | Get all Payment Account Balances
*AccountsAdminApi* | [**AccountsBalancesHistoryGet**](docs/AccountsAdminApi.md#accountsbalanceshistoryget) | **Get** /v2/accounts/{accountId}/balances/history | Get Payment Account Balance History
*AccountsAdminApi* | [**AccountsBlockedPut**](docs/AccountsAdminApi.md#accountsblockedput) | **Put** /v2/accounts/{accountid}/blocked | Block or Unblock an Account
*AccountsAdminApi* | [**AccountsDocumentGet**](docs/AccountsAdminApi.md#accountsdocumentget) | **Get** /v2/accounts/{accountid}/documents | Get Documents Associated with a Payment Account
*AccountsAdminApi* | [**AccountsDocumentsSignInitPost**](docs/AccountsAdminApi.md#accountsdocumentssigninitpost) | **Post** /v2/accounts/{accountid}/documents/{documentid}/signinit | Generate an Electronic Signature of a Document
*AccountsAdminApi* | [**AccountsEnrolmentInit**](docs/AccountsAdminApi.md#accountsenrolmentinit) | **Post** /v2/accounts/{accountid}/enrolment/init | Initialize a Deutsche Post POSTIDENT Identification
*AccountsAdminApi* | [**AccountsRetrievePost**](docs/AccountsAdminApi.md#accountsretrievepost) | **Post** /v2/accounts/retrieve | Get Detailed Payments Accounts Data
*AccountsAdminApi* | [**AccountsTransactionsGet**](docs/AccountsAdminApi.md#accountstransactionsget) | **Get** /v2/accounts/{accountId}/transactions | Get list of all Payment Account transactions
*AccountsCreateKYCApi* | [**AccountsAddUltimateBeneficialOwner**](docs/AccountsCreateKYCApi.md#accountsaddultimatebeneficialowner) | **Post** /v2/accounts/{accountId}/ultimateBeneficialOwner | Create an Ultimate Beneficial Owner
*AccountsCreateKYCApi* | [**AccountsDocumentsUploadPost**](docs/AccountsCreateKYCApi.md#accountsdocumentsuploadpost) | **Post** /v2/accounts/{accountid}/documents/upload | Upload Documents for KYC (Know Your Customers)
*AccountsCreateKYCApi* | [**AccountsGetUltimateBeneficialOwner**](docs/AccountsCreateKYCApi.md#accountsgetultimatebeneficialowner) | **Get** /v2/accounts/{accountId}ultimateBeneficialOwner | Get all Ultimate Beneficial Owners associated to a payment account.
*AccountsCreateKYCApi* | [**AccountsIndividualPost**](docs/AccountsCreateKYCApi.md#accountsindividualpost) | **Post** /v2/accounts/individual | Create a New Individual Account
*AccountsCreateKYCApi* | [**AccountsKycStatusGet**](docs/AccountsCreateKYCApi.md#accountskycstatusget) | **Get** /v2/accounts/kycstatus | Find a user, document or an IBAN that has been modified since an entry date
*AccountsCreateKYCApi* | [**AccountsLegalPost**](docs/AccountsCreateKYCApi.md#accountslegalpost) | **Post** /v2/accounts/legal | Create a New Legal Account
*AccountsCreateKYCApi* | [**AccountsUpdateUltimateBeneficialOwner**](docs/AccountsCreateKYCApi.md#accountsupdateultimatebeneficialowner) | **Put** /v2/accounts/{accountId}/ultimateBeneficialOwner/{UltimateBeneficialOwnerId} | Update Ultimate Beneficial Owner data
*AccountsUpdateApi* | [**AccountsIndividualPut**](docs/AccountsUpdateApi.md#accountsindividualput) | **Put** /v2/accounts/individual/{accountid} | Update Individual Payment Account Data
*AccountsUpdateApi* | [**AccountsKycstatusPut**](docs/AccountsUpdateApi.md#accountskycstatusput) | **Put** /v2/accounts/kycstatus/{accountid} | Update Payment Account Status
*AccountsUpdateApi* | [**AccountsLegalSinglePut**](docs/AccountsUpdateApi.md#accountslegalsingleput) | **Put** /v2/accounts/legal/{accountid} | Update Legal Payment Account Data
*DisputesApi* | [**DisputesDisputesGet**](docs/DisputesApi.md#disputesdisputesget) | **Get** /v2/disputes | Get list of disputes from a given date
*MoneyInsBNPLApi* | [**MoneyInsCreateBnplPayment**](docs/MoneyInsBNPLApi.md#moneyinscreatebnplpayment) | **Post** /v2/moneyins/buynowpaylater/init | Create New Pending Payment
*MoneyInsBNPLApi* | [**MoneyInsGetBnplPaymentPlans**](docs/MoneyInsBNPLApi.md#moneyinsgetbnplpaymentplans) | **Get** /v2/moneyins/buynowpaylater/plans | Get all Actived Payment Plans
*MoneyInsCardsApi* | [**MoneyInsCancelPut**](docs/MoneyInsCardsApi.md#moneyinscancelput) | **Put** /v2/moneyins/{transactionid}/cancel | Cancel a Money-In
*MoneyInsCardsApi* | [**MoneyInsCardCreatePost**](docs/MoneyInsCardsApi.md#moneyinscardcreatepost) | **Post** /v2/moneyins/card/create | Credit an Account with Money-In with Card without PSP process
*MoneyInsCardsApi* | [**MoneyInsCardDirect3DAuthenticatePost**](docs/MoneyInsCardsApi.md#moneyinscarddirect3dauthenticatepost) | **Post** /v2/moneyins/card/direct/{transactionid}/3dauthenticate | Check Money-In 3D-Secure Status (PCI-DSS compliant only)
*MoneyInsCardsApi* | [**MoneyInsCardDirect3DConfirmPut**](docs/MoneyInsCardsApi.md#moneyinscarddirect3dconfirmput) | **Put** /v2/moneyins/card/direct/{transactionid}/3dconfirm | Finalize a Direct Payment (PCI-DSS compliant only)
*MoneyInsCardsApi* | [**MoneyInsCardDirectPost**](docs/MoneyInsCardsApi.md#moneyinscarddirectpost) | **Post** /v2/moneyins/card/direct | (Deprecated) Credit an Account with a non 3-D Secure Card Payment (PCI-DSS compliant only)
*MoneyInsCardsApi* | [**MoneyInsCardGet**](docs/MoneyInsCardsApi.md#moneyinscardget) | **Get** /v2/moneyins/card/{cardId} | Get Card Information
*MoneyInsCardsApi* | [**MoneyInsCardGet_0**](docs/MoneyInsCardsApi.md#moneyinscardget_0) | **Get** /v2/moneyins/{accountid}/card | Get the Card Associated to a Payment Account
*MoneyInsCardsApi* | [**MoneyInsCardRebill**](docs/MoneyInsCardsApi.md#moneyinscardrebill) | **Post** /v2/moneyins/card/{cardid}/rebill | Charge a Registered Card
*MoneyInsCardsApi* | [**MoneyInsCardRegisterPost**](docs/MoneyInsCardsApi.md#moneyinscardregisterpost) | **Post** /v2/moneyins/card/register | (Deprecated) Register a Card for Direct Payments (PCI-DSS compliant only)
*MoneyInsCardsApi* | [**MoneyInsCardSubscriptionPost**](docs/MoneyInsCardsApi.md#moneyinscardsubscriptionpost) | **Post** /v2/moneyins/card/{cardid}/subscription | Initiate Monthly Recurring Payments
*MoneyInsCardsApi* | [**MoneyInsCardUnregisterPut**](docs/MoneyInsCardsApi.md#moneyinscardunregisterput) | **Put** /v2/moneyins/card/{cardid}/unregister | Unregister a Card Token
*MoneyInsCardsApi* | [**MoneyInsCardWebInitPost**](docs/MoneyInsCardsApi.md#moneyinscardwebinitpost) | **Post** /v2/moneyins/card/webinit | Initiate a Web Payment
*MoneyInsCardsApi* | [**MoneyInsDirect3DInitPost**](docs/MoneyInsCardsApi.md#moneyinsdirect3dinitpost) | **Post** /v2/moneyins/card/direct/3dinit | Initiate a Direct Payment (PCI-DSS compliant only)
*MoneyInsCardsApi* | [**MoneyInsMoneyInGet**](docs/MoneyInsCardsApi.md#moneyinsmoneyinget) | **Get** /v2/moneyins | Retrieve Payment Details
*MoneyInsCardsApi* | [**MoneyInsValidatePut**](docs/MoneyInsCardsApi.md#moneyinsvalidateput) | **Put** /v2/moneyins/{transactionid}/validate | Capture a Deferred Payment
*MoneyInsChequeApi* | [**MoneyInsChequeGet**](docs/MoneyInsChequeApi.md#moneyinschequeget) | **Get** /v2/moneyins/cheque | Search for Cheque Money-In by Date or by Token
*MoneyInsChequeApi* | [**MoneyInsChequeInitPost**](docs/MoneyInsChequeApi.md#moneyinschequeinitpost) | **Post** /v2/moneyins/cheque/init | Register a Cheque
*MoneyInsDirectDebitsApi* | [**MoneyInsCancelPut**](docs/MoneyInsDirectDebitsApi.md#moneyinscancelput) | **Put** /v2/moneyins/{transactionid}/cancel | Cancel a Money-In
*MoneyInsDirectDebitsApi* | [**MoneyInsMandateGet**](docs/MoneyInsDirectDebitsApi.md#moneyinsmandateget) | **Get** /v2/moneyins/{accountid}/mandate | Get Mandate Associated to a Payment Account
*MoneyInsDirectDebitsApi* | [**MoneyInsMandateGetDocument**](docs/MoneyInsDirectDebitsApi.md#moneyinsmandategetdocument) | **Get** /v2/moneyins/{accountid}/mandate/{mandateid}/document | Get Mandate Document
*MoneyInsDirectDebitsApi* | [**MoneyInsSddGet**](docs/MoneyInsDirectDebitsApi.md#moneyinssddget) | **Get** /v2/moneyins/sdd | List of Money-In by SEPA Direct Debit (SDD)
*MoneyInsDirectDebitsApi* | [**MoneyInsSddInitPost**](docs/MoneyInsDirectDebitsApi.md#moneyinssddinitpost) | **Post** /v2/moneyins/sdd/init | Request a SEPA Direct Debit (SDD)
*MoneyInsDirectDebitsApi* | [**MoneyInsSddMandatePost**](docs/MoneyInsDirectDebitsApi.md#moneyinssddmandatepost) | **Post** /v2/moneyins/sdd/mandate | Register a SDD Mandate
*MoneyInsDirectDebitsApi* | [**MoneyInsSddMandateUnregisterPut**](docs/MoneyInsDirectDebitsApi.md#moneyinssddmandateunregisterput) | **Put** /v2/moneyins/sdd/mandate/{mandatid}/unregister | Deactivate a Mandate
*MoneyInsLocalPaymentsApi* | [**MoneyInsIDealConfirmPut**](docs/MoneyInsLocalPaymentsApi.md#moneyinsidealconfirmput) | **Put** /v2/moneyins/ideal/{transactionid}/confirm | Finalize an iDeal Payment
*MoneyInsLocalPaymentsApi* | [**MoneyInsIDealInitPost**](docs/MoneyInsLocalPaymentsApi.md#moneyinsidealinitpost) | **Post** /v2/moneyins/ideal/init | Initialize iDeal Payment
*MoneyInsLocalPaymentsApi* | [**MoneyInsMbwayInitPost**](docs/MoneyInsLocalPaymentsApi.md#moneyinsmbwayinitpost) | **Post** /v2/moneyins/mbway/init | Initialize MB WAY Payment
*MoneyInsLocalPaymentsApi* | [**MoneyInsMobilePayInitPost**](docs/MoneyInsLocalPaymentsApi.md#moneyinsmobilepayinitpost) | **Post** /v2/moneyins/mobilePay/init | Initialize MobilePay Payment
*MoneyInsLocalPaymentsApi* | [**MoneyInsMultibancoInitPost**](docs/MoneyInsLocalPaymentsApi.md#moneyinsmultibancoinitpost) | **Post** /v2/moneyins/multibanco/init | Initialize Multibanco Payment
*MoneyInsLocalPaymentsApi* | [**MoneyInsPayTrailInitPost**](docs/MoneyInsLocalPaymentsApi.md#moneyinspaytrailinitpost) | **Post** /v2/moneyins/paytrail/init | Initialize PayTrail Payment (Deprecated)
*MoneyInsLocalPaymentsApi* | [**MoneyInsPayshopInitPost**](docs/MoneyInsLocalPaymentsApi.md#moneyinspayshopinitpost) | **Post** /v2/moneyins/payshop/init | Initialize Payshop Payment
*MoneyInsLocalPaymentsApi* | [**MoneyInsSofortInitPost**](docs/MoneyInsLocalPaymentsApi.md#moneyinssofortinitpost) | **Post** /v2/moneyins/sofort/init | Initialize Sofort Payment
*MoneyInsLocalPaymentsApi* | [**MoneyInsTrustlyInitPost**](docs/MoneyInsLocalPaymentsApi.md#moneyinstrustlyinitpost) | **Post** /v2/moneyins/trustly/init | Initialize Trustly Payment
*MoneyInsPayByBankApi* | [**MoneyInsGetMoneyInBanks**](docs/MoneyInsPayByBankApi.md#moneyinsgetmoneyinbanks) | **Get** /v2/moneyins/paybybank/transfer/banks | Get Pay by Bank List
*MoneyInsPayByBankApi* | [**MoneyInsMoneyInTransferInit**](docs/MoneyInsPayByBankApi.md#moneyinsmoneyintransferinit) | **Post** /v2/moneyins/paybybank/transfer/init | Initiate Pay by Bank
*MoneyInsPayByLinkApi* | [**MoneyInsCardPaymentFormPost**](docs/MoneyInsPayByLinkApi.md#moneyinscardpaymentformpost) | **Post** /v2/moneyins/card/paymentform | Create Payment Form (Pay by Link)
*MoneyInsPayByLinkApi* | [**MoneyInsPaymentFormCompletedGet**](docs/MoneyInsPayByLinkApi.md#moneyinspaymentformcompletedget) | **Get** /v2/moneyins/paymentform/{formid}/completed | Get Details of a Completed Payment Form
*MoneyInsPayByLinkApi* | [**MoneyInsPaymentFormDisablePut**](docs/MoneyInsPayByLinkApi.md#moneyinspaymentformdisableput) | **Put** /v2/moneyins/paymentform/{formid}/disable | Disable a Payment Form
*MoneyInsPayPalApi* | [**MoneyInsMoneyInPayPalInit**](docs/MoneyInsPayPalApi.md#moneyinsmoneyinpaypalinit) | **Post** /v2/moneyins/paypal/init | Initate Pay by PayPal
*MoneyInsPayPalApi* | [**MoneyInsPayPalTransactionResume**](docs/MoneyInsPayPalApi.md#moneyinspaypaltransactionresume) | **Post** /v2/moneyins/paypal/{transactionId}/resume | PayPal Resume
*MoneyInsTransfersInApi* | [**MoneyInsBankwireGet**](docs/MoneyInsTransfersInApi.md#moneyinsbankwireget) | **Get** /v2/moneyins/bankwire | Search for a Money-In by Fund Transfer
*MoneyInsVirtualIBANApi* | [**MoneyInsBankwireIbanCreatePost**](docs/MoneyInsVirtualIBANApi.md#moneyinsbankwireibancreatepost) | **Post** /v2/moneyins/bankwire/iban/create | Create Dedicated Virtual IBANs
*MoneyInsVirtualIBANApi* | [**MoneyInsBankwireIbanDisablePost**](docs/MoneyInsVirtualIBANApi.md#moneyinsbankwireibandisablepost) | **Post** /v2/moneyins/bankwire/iban/{ibanid}/disable | Disable a Dedicated Virtual IBAN
*MoneyOutsApi* | [**MoneyOutsCancelPut**](docs/MoneyOutsApi.md#moneyoutscancelput) | **Put** /v2/moneyouts/{transactionid}/cancel | Money-Out Cancellation
*MoneyOutsApi* | [**MoneyOutsIbanExtendedPost**](docs/MoneyOutsApi.md#moneyoutsibanextendedpost) | **Post** /v2/moneyouts/iban/extended | Add Bank Account to a Payment Account for Money-Outs
*MoneyOutsApi* | [**MoneyOutsIbanGet**](docs/MoneyOutsApi.md#moneyoutsibanget) | **Get** /v2/moneyouts/{accountid}/iban | Get the IBAN Associated with a Payment Account
*MoneyOutsApi* | [**MoneyOutsIbanPost**](docs/MoneyOutsApi.md#moneyoutsibanpost) | **Post** /v2/moneyouts/iban | Add an IBAN to a Payment Account for Money-Outs
*MoneyOutsApi* | [**MoneyOutsIbanUnregisterPut**](docs/MoneyOutsApi.md#moneyoutsibanunregisterput) | **Put** /v2/moneyouts/iban/{IbanId}/unregister | Disable Bank Information (IBAN) from a Payment Account
*MoneyOutsApi* | [**MoneyOutsMoneyOutGet**](docs/MoneyOutsApi.md#moneyoutsmoneyoutget) | **Get** /v2/moneyouts | Search for a Money-Out
*MoneyOutsApi* | [**MoneyOutsMoneyOutPost**](docs/MoneyOutsApi.md#moneyoutsmoneyoutpost) | **Post** /v2/moneyouts | External Fund Transfer from a Payment Account to a Bank Account
*P2PsApi* | [**P2PsP2pGet**](docs/P2PsApi.md#p2psp2pget) | **Get** /v2/p2p | Search for Transactions between Payments Accounts
*P2PsApi* | [**P2PsP2pPost**](docs/P2PsApi.md#p2psp2ppost) | **Post** /v2/p2p | Payment between Payment Accounts (P2P)
*RefundsApi* | [**RefundsRefundCreatePut**](docs/RefundsApi.md#refundsrefundcreateput) | **Put** /v2/refundcreate/{transactionId} | Refund a transaction without PSP process
*RefundsApi* | [**RefundsRefundPut**](docs/RefundsApi.md#refundsrefundput) | **Put** /v2/refund/{transactionid} | Refund a Money-In


## Documentation For Models

 - [ACS](docs/ACS.md)
 - [AccountBalance](docs/AccountBalance.md)
 - [AccountBalanceOutput](docs/AccountBalanceOutput.md)
 - [AccountBalancesInput](docs/AccountBalancesInput.md)
 - [AccountBlock](docs/AccountBlock.md)
 - [AccountBlockedInput](docs/AccountBlockedInput.md)
 - [AccountBlockedOutput](docs/AccountBlockedOutput.md)
 - [AccountCardsOutput](docs/AccountCardsOutput.md)
 - [AccountDetails](docs/AccountDetails.md)
 - [AccountDetailsBatchInput](docs/AccountDetailsBatchInput.md)
 - [AccountDetailsBatchOutput](docs/AccountDetailsBatchOutput.md)
 - [AccountDetailsOutput](docs/AccountDetailsOutput.md)
 - [AccountDocumentsOutput](docs/AccountDocumentsOutput.md)
 - [AccountIbansOutput](docs/AccountIbansOutput.md)
 - [AccountKycStatus](docs/AccountKycStatus.md)
 - [AccountKycStatusInput](docs/AccountKycStatusInput.md)
 - [AccountKycStatusOutput](docs/AccountKycStatusOutput.md)
 - [AccountMandatsOutput](docs/AccountMandatsOutput.md)
 - [AddUltimateBeneficialOwnerInput](docs/AddUltimateBeneficialOwnerInput.md)
 - [AddUltimateBeneficialOwnerOutput](docs/AddUltimateBeneficialOwnerOutput.md)
 - [Address](docs/Address.md)
 - [AmountBreakdown](docs/AmountBreakdown.md)
 - [Authentication](docs/Authentication.md)
 - [BalanceHistoryInput](docs/BalanceHistoryInput.md)
 - [BalanceHistoryOutput](docs/BalanceHistoryOutput.md)
 - [Bank](docs/Bank.md)
 - [BankBranchAddress](docs/BankBranchAddress.md)
 - [BillingAddress](docs/BillingAddress.md)
 - [Birth](docs/Birth.md)
 - [BnplAddress](docs/BnplAddress.md)
 - [BnplDelivery](docs/BnplDelivery.md)
 - [BnplInfo](docs/BnplInfo.md)
 - [BnplPurchaseItem](docs/BnplPurchaseItem.md)
 - [CancelMoneyInInput](docs/CancelMoneyInInput.md)
 - [CancelMoneyInOutput](docs/CancelMoneyInOutput.md)
 - [CancelMoneyOutInput](docs/CancelMoneyOutInput.md)
 - [CancelMoneyOutOutput](docs/CancelMoneyOutOutput.md)
 - [CancellationToken](docs/CancellationToken.md)
 - [Card](docs/Card.md)
 - [CardInfo](docs/CardInfo.md)
 - [CardInfoExtended](docs/CardInfoExtended.md)
 - [Company](docs/Company.md)
 - [Contact](docs/Contact.md)
 - [CreateBnplPaymentInput](docs/CreateBnplPaymentInput.md)
 - [CreateBnplPaymentOutput](docs/CreateBnplPaymentOutput.md)
 - [CreateIBANInput](docs/CreateIBANInput.md)
 - [CreateIBANOutput](docs/CreateIBANOutput.md)
 - [CreatePaymentFormInput](docs/CreatePaymentFormInput.md)
 - [CreatePaymentFormOutput](docs/CreatePaymentFormOutput.md)
 - [Customer](docs/Customer.md)
 - [CustomerAccountInfo](docs/CustomerAccountInfo.md)
 - [Delivery](docs/Delivery.md)
 - [DeliveryAdditionalInfo](docs/DeliveryAdditionalInfo.md)
 - [DeliveryAddress](docs/DeliveryAddress.md)
 - [DisableIBANInput](docs/DisableIBANInput.md)
 - [DisableIBANOutput](docs/DisableIBANOutput.md)
 - [DisablePaymentFormOutput](docs/DisablePaymentFormOutput.md)
 - [Document](docs/Document.md)
 - [EnrolmentInitInput](docs/EnrolmentInitInput.md)
 - [EnrolmentInitOutput](docs/EnrolmentInitOutput.md)
 - [Error](docs/Error.md)
 - [EuPagoInit](docs/EuPagoInit.md)
 - [GetCardOutput](docs/GetCardOutput.md)
 - [GetChargebacksInput](docs/GetChargebacksInput.md)
 - [GetChargebacksOutput](docs/GetChargebacksOutput.md)
 - [GetCompletedPaymentFormOutput](docs/GetCompletedPaymentFormOutput.md)
 - [GetMoneyInBanksInput](docs/GetMoneyInBanksInput.md)
 - [GetMoneyInBanksOutput](docs/GetMoneyInBanksOutput.md)
 - [GetMoneyInChequeDetailsInput](docs/GetMoneyInChequeDetailsInput.md)
 - [GetMoneyInChequeDetailsOutput](docs/GetMoneyInChequeDetailsOutput.md)
 - [GetMoneyInIBANDetailsInput](docs/GetMoneyInIBANDetailsInput.md)
 - [GetMoneyInIBANDetailsOutput](docs/GetMoneyInIBANDetailsOutput.md)
 - [GetMoneyInSddInput](docs/GetMoneyInSddInput.md)
 - [GetMoneyInSddOutput](docs/GetMoneyInSddOutput.md)
 - [GetMoneyInTransDetailsInput](docs/GetMoneyInTransDetailsInput.md)
 - [GetMoneyInTransDetailsOutput](docs/GetMoneyInTransDetailsOutput.md)
 - [GetMoneyOutTransDetailsInput](docs/GetMoneyOutTransDetailsInput.md)
 - [GetMoneyOutTransDetailsOutput](docs/GetMoneyOutTransDetailsOutput.md)
 - [GetPaymentDetailsInput](docs/GetPaymentDetailsInput.md)
 - [GetPaymentDetailsOutput](docs/GetPaymentDetailsOutput.md)
 - [GetPaymentPlansOutput](docs/GetPaymentPlansOutput.md)
 - [GetUltimateBeneficialOwnerOutput](docs/GetUltimateBeneficialOwnerOutput.md)
 - [Holder](docs/Holder.md)
 - [Iban](docs/Iban.md)
 - [IndividualAccount](docs/IndividualAccount.md)
 - [InitPayPalTransactionInput](docs/InitPayPalTransactionInput.md)
 - [InitPayPalTransactionOutput](docs/InitPayPalTransactionOutput.md)
 - [KycStatusOutput](docs/KycStatusOutput.md)
 - [LegalAccount](docs/LegalAccount.md)
 - [LemonWayCommission](docs/LemonWayCommission.md)
 - [Limits](docs/Limits.md)
 - [Link](docs/Link.md)
 - [Links](docs/Links.md)
 - [MandateGetDocumentOutput](docs/MandateGetDocumentOutput.md)
 - [MoneyIn3DAuthenticateInput](docs/MoneyIn3DAuthenticateInput.md)
 - [MoneyIn3DAuthenticateOutput](docs/MoneyIn3DAuthenticateOutput.md)
 - [MoneyIn3DConfirmInput](docs/MoneyIn3DConfirmInput.md)
 - [MoneyIn3DConfirmOutput](docs/MoneyIn3DConfirmOutput.md)
 - [MoneyIn3DInitInput](docs/MoneyIn3DInitInput.md)
 - [MoneyIn3DInitOutput](docs/MoneyIn3DInitOutput.md)
 - [MoneyInChequeInitInput](docs/MoneyInChequeInitInput.md)
 - [MoneyInChequeInitOutput](docs/MoneyInChequeInitOutput.md)
 - [MoneyInCreateInput](docs/MoneyInCreateInput.md)
 - [MoneyInIDealConfirmOutput](docs/MoneyInIDealConfirmOutput.md)
 - [MoneyInIDealInitInput](docs/MoneyInIDealInitInput.md)
 - [MoneyInIDealInitOutput](docs/MoneyInIDealInitOutput.md)
 - [MoneyInInput](docs/MoneyInInput.md)
 - [MoneyInMbwayInitInput](docs/MoneyInMbwayInitInput.md)
 - [MoneyInMbwayInitOutput](docs/MoneyInMbwayInitOutput.md)
 - [MoneyInMobilePayInitInput](docs/MoneyInMobilePayInitInput.md)
 - [MoneyInMobilePayInitOutput](docs/MoneyInMobilePayInitOutput.md)
 - [MoneyInMultibancoInitInput](docs/MoneyInMultibancoInitInput.md)
 - [MoneyInMultibancoInitOutput](docs/MoneyInMultibancoInitOutput.md)
 - [MoneyInOutput](docs/MoneyInOutput.md)
 - [MoneyInPayTrailInitInput](docs/MoneyInPayTrailInitInput.md)
 - [MoneyInPayTrailInitOutput](docs/MoneyInPayTrailInitOutput.md)
 - [MoneyInPayshopInitInput](docs/MoneyInPayshopInitInput.md)
 - [MoneyInPayshopInitOutput](docs/MoneyInPayshopInitOutput.md)
 - [MoneyInSddInitInput](docs/MoneyInSddInitInput.md)
 - [MoneyInSddInitOutput](docs/MoneyInSddInitOutput.md)
 - [MoneyInSofortInitInput](docs/MoneyInSofortInitInput.md)
 - [MoneyInSofortInitOutput](docs/MoneyInSofortInitOutput.md)
 - [MoneyInSubscriptionInitInput](docs/MoneyInSubscriptionInitInput.md)
 - [MoneyInSubscriptionInitOutput](docs/MoneyInSubscriptionInitOutput.md)
 - [MoneyInTransferInitInput](docs/MoneyInTransferInitInput.md)
 - [MoneyInTransferInitOutput](docs/MoneyInTransferInitOutput.md)
 - [MoneyInTrustlyInitInput](docs/MoneyInTrustlyInitInput.md)
 - [MoneyInTrustlyInitOutput](docs/MoneyInTrustlyInitOutput.md)
 - [MoneyInValidateInput](docs/MoneyInValidateInput.md)
 - [MoneyInValidateOutput](docs/MoneyInValidateOutput.md)
 - [MoneyInWebInitInput](docs/MoneyInWebInitInput.md)
 - [MoneyInWebInitOutput](docs/MoneyInWebInitOutput.md)
 - [MoneyInWithCardIdInput](docs/MoneyInWithCardIdInput.md)
 - [MoneyInWithCardIdOutput](docs/MoneyInWithCardIdOutput.md)
 - [MoneyOutInput](docs/MoneyOutInput.md)
 - [MoneyOutOutput](docs/MoneyOutOutput.md)
 - [PSP](docs/PSP.md)
 - [Pagination](docs/Pagination.md)
 - [PaymentForm](docs/PaymentForm.md)
 - [PaymentFormDetails](docs/PaymentFormDetails.md)
 - [PaymentPlan](docs/PaymentPlan.md)
 - [PaypalDeliveryAddress](docs/PaypalDeliveryAddress.md)
 - [PrivateData](docs/PrivateData.md)
 - [PurchaseItem](docs/PurchaseItem.md)
 - [Receiver](docs/Receiver.md)
 - [Redirections](docs/Redirections.md)
 - [RefundMoneyInCreateInput](docs/RefundMoneyInCreateInput.md)
 - [RefundMoneyInInput](docs/RefundMoneyInInput.md)
 - [RefundMoneyInOutput](docs/RefundMoneyInOutput.md)
 - [RegisterCardInput](docs/RegisterCardInput.md)
 - [RegisterCardOutput](docs/RegisterCardOutput.md)
 - [RegisterIBANExtendedInput](docs/RegisterIBANExtendedInput.md)
 - [RegisterIBANExtendedOutput](docs/RegisterIBANExtendedOutput.md)
 - [RegisterIBANInput](docs/RegisterIBANInput.md)
 - [RegisterIBANOutput](docs/RegisterIBANOutput.md)
 - [RegisterIndividualAccountInput](docs/RegisterIndividualAccountInput.md)
 - [RegisterIndividualAccountOutput](docs/RegisterIndividualAccountOutput.md)
 - [RegisterLegalAccountInput](docs/RegisterLegalAccountInput.md)
 - [RegisterLegalAccountOutput](docs/RegisterLegalAccountOutput.md)
 - [RegisterSddMandateInput](docs/RegisterSddMandateInput.md)
 - [RegisterSddMandateOutput](docs/RegisterSddMandateOutput.md)
 - [RiskAnalysis](docs/RiskAnalysis.md)
 - [SafeWaitHandle](docs/SafeWaitHandle.md)
 - [SddMandate](docs/SddMandate.md)
 - [SendPaymentInput](docs/SendPaymentInput.md)
 - [SendPaymentOutput](docs/SendPaymentOutput.md)
 - [SignDocumentInitInput](docs/SignDocumentInitInput.md)
 - [SignDocumentInitOutput](docs/SignDocumentInitOutput.md)
 - [ThreeDs](docs/ThreeDs.md)
 - [ThreeDs](docs/ThreeDs.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionAccount](docs/TransactionAccount.md)
 - [TransactionAuthorizedCapture](docs/TransactionAuthorizedCapture.md)
 - [TransactionCaptured](docs/TransactionCaptured.md)
 - [TransactionCapturedAuthorization](docs/TransactionCapturedAuthorization.md)
 - [TransactionIn](docs/TransactionIn.md)
 - [TransactionInDetails](docs/TransactionInDetails.md)
 - [TransactionInfo](docs/TransactionInfo.md)
 - [TransactionOut](docs/TransactionOut.md)
 - [TransactionP2P](docs/TransactionP2P.md)
 - [TransactionRefund](docs/TransactionRefund.md)
 - [TransactionsHistoryInput](docs/TransactionsHistoryInput.md)
 - [TransactionsTransactionAccount](docs/TransactionsTransactionAccount.md)
 - [TransactionsTransactionIn](docs/TransactionsTransactionIn.md)
 - [TransactionsTransactionInDetails](docs/TransactionsTransactionInDetails.md)
 - [TransactionsTransactionOut](docs/TransactionsTransactionOut.md)
 - [TransactionsTransactionP2P](docs/TransactionsTransactionP2P.md)
 - [UltimateBeneficialOwner](docs/UltimateBeneficialOwner.md)
 - [UnregisterCardInput](docs/UnregisterCardInput.md)
 - [UnregisterCardOutput](docs/UnregisterCardOutput.md)
 - [UnregisterIBANInput](docs/UnregisterIBANInput.md)
 - [UnregisterIBANOutput](docs/UnregisterIBANOutput.md)
 - [UnregisterSddMandateInput](docs/UnregisterSddMandateInput.md)
 - [UnregisterSddMandateOutput](docs/UnregisterSddMandateOutput.md)
 - [UpdateAccountStatus](docs/UpdateAccountStatus.md)
 - [UpdateAccountStatusInput](docs/UpdateAccountStatusInput.md)
 - [UpdateAccountStatusOutput](docs/UpdateAccountStatusOutput.md)
 - [UpdateIndividualAccountDetailsInput](docs/UpdateIndividualAccountDetailsInput.md)
 - [UpdateIndividualAccountDetailsOutput](docs/UpdateIndividualAccountDetailsOutput.md)
 - [UpdateLegalAccountDetailsInput](docs/UpdateLegalAccountDetailsInput.md)
 - [UpdateLegalAccountDetailsOutput](docs/UpdateLegalAccountDetailsOutput.md)
 - [UpdateUltimateBeneficialOwnerInput](docs/UpdateUltimateBeneficialOwnerInput.md)
 - [UpdateUltimateBeneficialOwnerOutput](docs/UpdateUltimateBeneficialOwnerOutput.md)
 - [UploadDocument](docs/UploadDocument.md)
 - [UploadDocumentInput](docs/UploadDocumentInput.md)
 - [UploadDocumentOutput](docs/UploadDocumentOutput.md)
 - [WaitHandle](docs/WaitHandle.md)
 - [WalletDetailsInput](docs/WalletDetailsInput.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



