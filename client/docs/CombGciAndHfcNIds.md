# CombGciAndHfcNIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalCableId** | Pointer to **string** | Global Cable Identifier uniquely identifying the connection between the 5G-CRG or FN-CRG to the 5GS. See clause 28.15.4 of 3GPP TS 23.003. This shall be encoded as a string per clause 28.15.4 of 3GPP TS 23.003, and compliant with the syntax specified  in clause 2.2  of IETF RFC 7542 for the username part of a NAI. The GCI value is specified in CableLabs WR-TR-5WWC-ARCH.  | [optional] 
**HfcNId** | Pointer to **string** | This IE represents the identifier of the HFC node Id as specified in CableLabs WR-TR-5WWC-ARCH. It is provisioned by the wireline operator as part of wireline operations and may contain up to six characters.  | [optional] 

## Methods

### NewCombGciAndHfcNIds

`func NewCombGciAndHfcNIds() *CombGciAndHfcNIds`

NewCombGciAndHfcNIds instantiates a new CombGciAndHfcNIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombGciAndHfcNIdsWithDefaults

`func NewCombGciAndHfcNIdsWithDefaults() *CombGciAndHfcNIds`

NewCombGciAndHfcNIdsWithDefaults instantiates a new CombGciAndHfcNIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalCableId

`func (o *CombGciAndHfcNIds) GetGlobalCableId() string`

GetGlobalCableId returns the GlobalCableId field if non-nil, zero value otherwise.

### GetGlobalCableIdOk

`func (o *CombGciAndHfcNIds) GetGlobalCableIdOk() (*string, bool)`

GetGlobalCableIdOk returns a tuple with the GlobalCableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCableId

`func (o *CombGciAndHfcNIds) SetGlobalCableId(v string)`

SetGlobalCableId sets GlobalCableId field to given value.

### HasGlobalCableId

`func (o *CombGciAndHfcNIds) HasGlobalCableId() bool`

HasGlobalCableId returns a boolean if a field has been set.

### GetHfcNId

`func (o *CombGciAndHfcNIds) GetHfcNId() string`

GetHfcNId returns the HfcNId field if non-nil, zero value otherwise.

### GetHfcNIdOk

`func (o *CombGciAndHfcNIds) GetHfcNIdOk() (*string, bool)`

GetHfcNIdOk returns a tuple with the HfcNId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHfcNId

`func (o *CombGciAndHfcNIds) SetHfcNId(v string)`

SetHfcNId sets HfcNId field to given value.

### HasHfcNId

`func (o *CombGciAndHfcNIds) HasHfcNId() bool`

HasHfcNId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


