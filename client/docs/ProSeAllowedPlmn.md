# ProSeAllowedPlmn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitedPlmn** | [**PlmnId**](PlmnId.md) |  | 
**ProseDirectAllowed** | Pointer to [**[]ProseDirectAllowed**](ProseDirectAllowed.md) |  | [optional] 

## Methods

### NewProSeAllowedPlmn

`func NewProSeAllowedPlmn(visitedPlmn PlmnId, ) *ProSeAllowedPlmn`

NewProSeAllowedPlmn instantiates a new ProSeAllowedPlmn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProSeAllowedPlmnWithDefaults

`func NewProSeAllowedPlmnWithDefaults() *ProSeAllowedPlmn`

NewProSeAllowedPlmnWithDefaults instantiates a new ProSeAllowedPlmn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisitedPlmn

`func (o *ProSeAllowedPlmn) GetVisitedPlmn() PlmnId`

GetVisitedPlmn returns the VisitedPlmn field if non-nil, zero value otherwise.

### GetVisitedPlmnOk

`func (o *ProSeAllowedPlmn) GetVisitedPlmnOk() (*PlmnId, bool)`

GetVisitedPlmnOk returns a tuple with the VisitedPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedPlmn

`func (o *ProSeAllowedPlmn) SetVisitedPlmn(v PlmnId)`

SetVisitedPlmn sets VisitedPlmn field to given value.


### GetProseDirectAllowed

`func (o *ProSeAllowedPlmn) GetProseDirectAllowed() []ProseDirectAllowed`

GetProseDirectAllowed returns the ProseDirectAllowed field if non-nil, zero value otherwise.

### GetProseDirectAllowedOk

`func (o *ProSeAllowedPlmn) GetProseDirectAllowedOk() (*[]ProseDirectAllowed, bool)`

GetProseDirectAllowedOk returns a tuple with the ProseDirectAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseDirectAllowed

`func (o *ProSeAllowedPlmn) SetProseDirectAllowed(v []ProseDirectAllowed)`

SetProseDirectAllowed sets ProseDirectAllowed field to given value.

### HasProseDirectAllowed

`func (o *ProSeAllowedPlmn) HasProseDirectAllowed() bool`

HasProseDirectAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


