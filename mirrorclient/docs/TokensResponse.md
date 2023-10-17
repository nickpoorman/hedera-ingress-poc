# TokensResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | Pointer to [**[]Token**](Token.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewTokensResponse

`func NewTokensResponse() *TokensResponse`

NewTokensResponse instantiates a new TokensResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensResponseWithDefaults

`func NewTokensResponseWithDefaults() *TokensResponse`

NewTokensResponseWithDefaults instantiates a new TokensResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *TokensResponse) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TokensResponse) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TokensResponse) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *TokensResponse) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetLinks

`func (o *TokensResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokensResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokensResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TokensResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


