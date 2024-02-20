# AstiAllowedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AstiAllowed** | **bool** |  | 
**CoverageArea** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**UuTimeSyncErrBdgt** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TempVals** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) |  | [optional] 

## Methods

### NewAstiAllowedInfo

`func NewAstiAllowedInfo(astiAllowed bool, ) *AstiAllowedInfo`

NewAstiAllowedInfo instantiates a new AstiAllowedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAstiAllowedInfoWithDefaults

`func NewAstiAllowedInfoWithDefaults() *AstiAllowedInfo`

NewAstiAllowedInfoWithDefaults instantiates a new AstiAllowedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAstiAllowed

`func (o *AstiAllowedInfo) GetAstiAllowed() bool`

GetAstiAllowed returns the AstiAllowed field if non-nil, zero value otherwise.

### GetAstiAllowedOk

`func (o *AstiAllowedInfo) GetAstiAllowedOk() (*bool, bool)`

GetAstiAllowedOk returns a tuple with the AstiAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAstiAllowed

`func (o *AstiAllowedInfo) SetAstiAllowed(v bool)`

SetAstiAllowed sets AstiAllowed field to given value.


### GetCoverageArea

`func (o *AstiAllowedInfo) GetCoverageArea() []Tai`

GetCoverageArea returns the CoverageArea field if non-nil, zero value otherwise.

### GetCoverageAreaOk

`func (o *AstiAllowedInfo) GetCoverageAreaOk() (*[]Tai, bool)`

GetCoverageAreaOk returns a tuple with the CoverageArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageArea

`func (o *AstiAllowedInfo) SetCoverageArea(v []Tai)`

SetCoverageArea sets CoverageArea field to given value.

### HasCoverageArea

`func (o *AstiAllowedInfo) HasCoverageArea() bool`

HasCoverageArea returns a boolean if a field has been set.

### GetUuTimeSyncErrBdgt

`func (o *AstiAllowedInfo) GetUuTimeSyncErrBdgt() int32`

GetUuTimeSyncErrBdgt returns the UuTimeSyncErrBdgt field if non-nil, zero value otherwise.

### GetUuTimeSyncErrBdgtOk

`func (o *AstiAllowedInfo) GetUuTimeSyncErrBdgtOk() (*int32, bool)`

GetUuTimeSyncErrBdgtOk returns a tuple with the UuTimeSyncErrBdgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuTimeSyncErrBdgt

`func (o *AstiAllowedInfo) SetUuTimeSyncErrBdgt(v int32)`

SetUuTimeSyncErrBdgt sets UuTimeSyncErrBdgt field to given value.

### HasUuTimeSyncErrBdgt

`func (o *AstiAllowedInfo) HasUuTimeSyncErrBdgt() bool`

HasUuTimeSyncErrBdgt returns a boolean if a field has been set.

### GetTempVals

`func (o *AstiAllowedInfo) GetTempVals() []TemporalValidity`

GetTempVals returns the TempVals field if non-nil, zero value otherwise.

### GetTempValsOk

`func (o *AstiAllowedInfo) GetTempValsOk() (*[]TemporalValidity, bool)`

GetTempValsOk returns a tuple with the TempVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVals

`func (o *AstiAllowedInfo) SetTempVals(v []TemporalValidity)`

SetTempVals sets TempVals field to given value.

### HasTempVals

`func (o *AstiAllowedInfo) HasTempVals() bool`

HasTempVals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


