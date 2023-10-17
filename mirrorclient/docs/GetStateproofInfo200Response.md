# GetStateproofInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressBooks** | **[]string** | The network address book valid at the time of the transaction | 
**RecordFile** | **string** | The content of the record file the transaction belongs to | 
**SignatureFiles** | **map[string]string** | The nodes&#39; signature files for the record file | 
**Version** | **int32** | The record file format version, either 5 or 6 | 

## Methods

### NewGetStateproofInfo200Response

`func NewGetStateproofInfo200Response(addressBooks []string, recordFile string, signatureFiles map[string]string, version int32, ) *GetStateproofInfo200Response`

NewGetStateproofInfo200Response instantiates a new GetStateproofInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStateproofInfo200ResponseWithDefaults

`func NewGetStateproofInfo200ResponseWithDefaults() *GetStateproofInfo200Response`

NewGetStateproofInfo200ResponseWithDefaults instantiates a new GetStateproofInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressBooks

`func (o *GetStateproofInfo200Response) GetAddressBooks() []string`

GetAddressBooks returns the AddressBooks field if non-nil, zero value otherwise.

### GetAddressBooksOk

`func (o *GetStateproofInfo200Response) GetAddressBooksOk() (*[]string, bool)`

GetAddressBooksOk returns a tuple with the AddressBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBooks

`func (o *GetStateproofInfo200Response) SetAddressBooks(v []string)`

SetAddressBooks sets AddressBooks field to given value.


### GetRecordFile

`func (o *GetStateproofInfo200Response) GetRecordFile() string`

GetRecordFile returns the RecordFile field if non-nil, zero value otherwise.

### GetRecordFileOk

`func (o *GetStateproofInfo200Response) GetRecordFileOk() (*string, bool)`

GetRecordFileOk returns a tuple with the RecordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFile

`func (o *GetStateproofInfo200Response) SetRecordFile(v string)`

SetRecordFile sets RecordFile field to given value.


### GetSignatureFiles

`func (o *GetStateproofInfo200Response) GetSignatureFiles() map[string]string`

GetSignatureFiles returns the SignatureFiles field if non-nil, zero value otherwise.

### GetSignatureFilesOk

`func (o *GetStateproofInfo200Response) GetSignatureFilesOk() (*map[string]string, bool)`

GetSignatureFilesOk returns a tuple with the SignatureFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureFiles

`func (o *GetStateproofInfo200Response) SetSignatureFiles(v map[string]string)`

SetSignatureFiles sets SignatureFiles field to given value.


### GetVersion

`func (o *GetStateproofInfo200Response) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetStateproofInfo200Response) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetStateproofInfo200Response) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


