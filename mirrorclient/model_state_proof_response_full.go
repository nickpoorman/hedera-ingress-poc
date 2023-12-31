/*
Hedera Mirror Node REST API

The Mirror Node REST API offers the ability to query cryptocurrency transactions and account information from a Hedera managed mirror node.  Base url: [/api/v1](/api/v1)  OpenAPI Spec: [/api/v1/docs/openapi.yml](/api/v1/docs/openapi.yml)

API version: 0.89.0
Contact: mirrornode@hedera.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mirrorclient

import (
	"encoding/json"
)

// checks if the StateProofResponseFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateProofResponseFull{}

// StateProofResponseFull struct for StateProofResponseFull
type StateProofResponseFull struct {
	// The network address book valid at the time of the transaction
	AddressBooks []string `json:"address_books"`
	// The content of the record file the transaction belongs to
	RecordFile string `json:"record_file"`
	// The nodes' signature files for the record file
	SignatureFiles map[string]string `json:"signature_files"`
}

// NewStateProofResponseFull instantiates a new StateProofResponseFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateProofResponseFull(addressBooks []string, recordFile string, signatureFiles map[string]string) *StateProofResponseFull {
	this := StateProofResponseFull{}
	this.AddressBooks = addressBooks
	this.RecordFile = recordFile
	this.SignatureFiles = signatureFiles
	return &this
}

// NewStateProofResponseFullWithDefaults instantiates a new StateProofResponseFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateProofResponseFullWithDefaults() *StateProofResponseFull {
	this := StateProofResponseFull{}
	return &this
}

// GetAddressBooks returns the AddressBooks field value
func (o *StateProofResponseFull) GetAddressBooks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AddressBooks
}

// GetAddressBooksOk returns a tuple with the AddressBooks field value
// and a boolean to check if the value has been set.
func (o *StateProofResponseFull) GetAddressBooksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressBooks, true
}

// SetAddressBooks sets field value
func (o *StateProofResponseFull) SetAddressBooks(v []string) {
	o.AddressBooks = v
}

// GetRecordFile returns the RecordFile field value
func (o *StateProofResponseFull) GetRecordFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordFile
}

// GetRecordFileOk returns a tuple with the RecordFile field value
// and a boolean to check if the value has been set.
func (o *StateProofResponseFull) GetRecordFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordFile, true
}

// SetRecordFile sets field value
func (o *StateProofResponseFull) SetRecordFile(v string) {
	o.RecordFile = v
}

// GetSignatureFiles returns the SignatureFiles field value
func (o *StateProofResponseFull) GetSignatureFiles() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.SignatureFiles
}

// GetSignatureFilesOk returns a tuple with the SignatureFiles field value
// and a boolean to check if the value has been set.
func (o *StateProofResponseFull) GetSignatureFilesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureFiles, true
}

// SetSignatureFiles sets field value
func (o *StateProofResponseFull) SetSignatureFiles(v map[string]string) {
	o.SignatureFiles = v
}

func (o StateProofResponseFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateProofResponseFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address_books"] = o.AddressBooks
	toSerialize["record_file"] = o.RecordFile
	toSerialize["signature_files"] = o.SignatureFiles
	return toSerialize, nil
}

type NullableStateProofResponseFull struct {
	value *StateProofResponseFull
	isSet bool
}

func (v NullableStateProofResponseFull) Get() *StateProofResponseFull {
	return v.value
}

func (v *NullableStateProofResponseFull) Set(val *StateProofResponseFull) {
	v.value = val
	v.isSet = true
}

func (v NullableStateProofResponseFull) IsSet() bool {
	return v.isSet
}

func (v *NullableStateProofResponseFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateProofResponseFull(val *StateProofResponseFull) *NullableStateProofResponseFull {
	return &NullableStateProofResponseFull{value: val, isSet: true}
}

func (v NullableStateProofResponseFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateProofResponseFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


