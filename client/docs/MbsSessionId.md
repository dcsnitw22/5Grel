# MbsSessionId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tmgi** | Pointer to [**Tmgi**](Tmgi.md) |  | [optional] 
**Ssm** | Pointer to [**Ssm**](Ssm.md) |  | [optional] 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 

## Methods

### NewMbsSessionId

`func NewMbsSessionId() *MbsSessionId`

NewMbsSessionId instantiates a new MbsSessionId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessionIdWithDefaults

`func NewMbsSessionIdWithDefaults() *MbsSessionId`

NewMbsSessionIdWithDefaults instantiates a new MbsSessionId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTmgi

`func (o *MbsSessionId) GetTmgi() Tmgi`

GetTmgi returns the Tmgi field if non-nil, zero value otherwise.

### GetTmgiOk

`func (o *MbsSessionId) GetTmgiOk() (*Tmgi, bool)`

GetTmgiOk returns a tuple with the Tmgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgi

`func (o *MbsSessionId) SetTmgi(v Tmgi)`

SetTmgi sets Tmgi field to given value.

### HasTmgi

`func (o *MbsSessionId) HasTmgi() bool`

HasTmgi returns a boolean if a field has been set.

### GetSsm

`func (o *MbsSessionId) GetSsm() Ssm`

GetSsm returns the Ssm field if non-nil, zero value otherwise.

### GetSsmOk

`func (o *MbsSessionId) GetSsmOk() (*Ssm, bool)`

GetSsmOk returns a tuple with the Ssm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsm

`func (o *MbsSessionId) SetSsm(v Ssm)`

SetSsm sets Ssm field to given value.

### HasSsm

`func (o *MbsSessionId) HasSsm() bool`

HasSsm returns a boolean if a field has been set.

### GetNid

`func (o *MbsSessionId) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *MbsSessionId) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *MbsSessionId) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *MbsSessionId) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


