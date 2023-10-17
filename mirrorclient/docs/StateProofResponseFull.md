# StateProofResponseFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressBooks** | **[]string** | The network address book valid at the time of the transaction | 
**RecordFile** | **string** | The content of the record file the transaction belongs to | 
**SignatureFiles** | **map[string]string** | The nodes&#39; signature files for the record file | 

## Methods

### NewStateProofResponseFull

`func NewStateProofResponseFull(addressBooks []string, recordFile string, signatureFiles map[string]string, ) *StateProofResponseFull`

NewStateProofResponseFull instantiates a new StateProofResponseFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateProofResponseFullWithDefaults

`func NewStateProofResponseFullWithDefaults() *StateProofResponseFull`

NewStateProofResponseFullWithDefaults instantiates a new StateProofResponseFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressBooks

`func (o *StateProofResponseFull) GetAddressBooks() []string`

GetAddressBooks returns the AddressBooks field if non-nil, zero value otherwise.

### GetAddressBooksOk

`func (o *StateProofResponseFull) GetAddressBooksOk() (*[]string, bool)`

GetAddressBooksOk returns a tuple with the AddressBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBooks

`func (o *StateProofResponseFull) SetAddressBooks(v []string)`

SetAddressBooks sets AddressBooks field to given value.


### GetRecordFile

`func (o *StateProofResponseFull) GetRecordFile() string`

GetRecordFile returns the RecordFile field if non-nil, zero value otherwise.

### GetRecordFileOk

`func (o *StateProofResponseFull) GetRecordFileOk() (*string, bool)`

GetRecordFileOk returns a tuple with the RecordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFile

`func (o *StateProofResponseFull) SetRecordFile(v string)`

SetRecordFile sets RecordFile field to given value.


### GetSignatureFiles

`func (o *StateProofResponseFull) GetSignatureFiles() map[string]string`

GetSignatureFiles returns the SignatureFiles field if non-nil, zero value otherwise.

### GetSignatureFilesOk

`func (o *StateProofResponseFull) GetSignatureFilesOk() (*map[string]string, bool)`

GetSignatureFilesOk returns a tuple with the SignatureFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureFiles

`func (o *StateProofResponseFull) SetSignatureFiles(v map[string]string)`

SetSignatureFiles sets SignatureFiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


