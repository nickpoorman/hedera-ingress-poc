# AccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]AccountInfo**](AccountInfo.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewAccountsResponse

`func NewAccountsResponse(accounts []AccountInfo, links Links, ) *AccountsResponse`

NewAccountsResponse instantiates a new AccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsResponseWithDefaults

`func NewAccountsResponseWithDefaults() *AccountsResponse`

NewAccountsResponseWithDefaults instantiates a new AccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountsResponse) GetAccounts() []AccountInfo`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountsResponse) GetAccountsOk() (*[]AccountInfo, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountsResponse) SetAccounts(v []AccountInfo)`

SetAccounts sets Accounts field to given value.


### GetLinks

`func (o *AccountsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccountsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccountsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


