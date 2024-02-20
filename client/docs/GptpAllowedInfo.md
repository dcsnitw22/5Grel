# GptpAllowedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**GptpAllowed** | **bool** |  | 
**CoverageArea** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**UuTimeSyncErrBdgt** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TempVals** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) |  | [optional] 

## Methods

### NewGptpAllowedInfo

`func NewGptpAllowedInfo(gptpAllowed bool, ) *GptpAllowedInfo`

NewGptpAllowedInfo instantiates a new GptpAllowedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGptpAllowedInfoWithDefaults

`func NewGptpAllowedInfoWithDefaults() *GptpAllowedInfo`

NewGptpAllowedInfoWithDefaults instantiates a new GptpAllowedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *GptpAllowedInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *GptpAllowedInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *GptpAllowedInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *GptpAllowedInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSNssai

`func (o *GptpAllowedInfo) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *GptpAllowedInfo) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *GptpAllowedInfo) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *GptpAllowedInfo) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetGptpAllowed

`func (o *GptpAllowedInfo) GetGptpAllowed() bool`

GetGptpAllowed returns the GptpAllowed field if non-nil, zero value otherwise.

### GetGptpAllowedOk

`func (o *GptpAllowedInfo) GetGptpAllowedOk() (*bool, bool)`

GetGptpAllowedOk returns a tuple with the GptpAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGptpAllowed

`func (o *GptpAllowedInfo) SetGptpAllowed(v bool)`

SetGptpAllowed sets GptpAllowed field to given value.


### GetCoverageArea

`func (o *GptpAllowedInfo) GetCoverageArea() []Tai`

GetCoverageArea returns the CoverageArea field if non-nil, zero value otherwise.

### GetCoverageAreaOk

`func (o *GptpAllowedInfo) GetCoverageAreaOk() (*[]Tai, bool)`

GetCoverageAreaOk returns a tuple with the CoverageArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageArea

`func (o *GptpAllowedInfo) SetCoverageArea(v []Tai)`

SetCoverageArea sets CoverageArea field to given value.

### HasCoverageArea

`func (o *GptpAllowedInfo) HasCoverageArea() bool`

HasCoverageArea returns a boolean if a field has been set.

### GetUuTimeSyncErrBdgt

`func (o *GptpAllowedInfo) GetUuTimeSyncErrBdgt() int32`

GetUuTimeSyncErrBdgt returns the UuTimeSyncErrBdgt field if non-nil, zero value otherwise.

### GetUuTimeSyncErrBdgtOk

`func (o *GptpAllowedInfo) GetUuTimeSyncErrBdgtOk() (*int32, bool)`

GetUuTimeSyncErrBdgtOk returns a tuple with the UuTimeSyncErrBdgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuTimeSyncErrBdgt

`func (o *GptpAllowedInfo) SetUuTimeSyncErrBdgt(v int32)`

SetUuTimeSyncErrBdgt sets UuTimeSyncErrBdgt field to given value.

### HasUuTimeSyncErrBdgt

`func (o *GptpAllowedInfo) HasUuTimeSyncErrBdgt() bool`

HasUuTimeSyncErrBdgt returns a boolean if a field has been set.

### GetTempVals

`func (o *GptpAllowedInfo) GetTempVals() []TemporalValidity`

GetTempVals returns the TempVals field if non-nil, zero value otherwise.

### GetTempValsOk

`func (o *GptpAllowedInfo) GetTempValsOk() (*[]TemporalValidity, bool)`

GetTempValsOk returns a tuple with the TempVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVals

`func (o *GptpAllowedInfo) SetTempVals(v []TemporalValidity)`

SetTempVals sets TempVals field to given value.

### HasTempVals

`func (o *GptpAllowedInfo) HasTempVals() bool`

HasTempVals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


