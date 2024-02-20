# Nssai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**DefaultSingleNssais** | [**[]Snssai**](Snssai.md) |  | 
**SingleNssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**ProvisioningTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AdditionalSnssaiData** | Pointer to [**map[string]AdditionalSnssaiData**](AdditionalSnssaiData.md) | A map(list of key-value pairs) where singleNssai serves as key of AdditionalSnssaiData | [optional] 
**SuppressNssrgInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewNssai

`func NewNssai(defaultSingleNssais []Snssai, ) *Nssai`

NewNssai instantiates a new Nssai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNssaiWithDefaults

`func NewNssaiWithDefaults() *Nssai`

NewNssaiWithDefaults instantiates a new Nssai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *Nssai) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *Nssai) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *Nssai) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *Nssai) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetDefaultSingleNssais

`func (o *Nssai) GetDefaultSingleNssais() []Snssai`

GetDefaultSingleNssais returns the DefaultSingleNssais field if non-nil, zero value otherwise.

### GetDefaultSingleNssaisOk

`func (o *Nssai) GetDefaultSingleNssaisOk() (*[]Snssai, bool)`

GetDefaultSingleNssaisOk returns a tuple with the DefaultSingleNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSingleNssais

`func (o *Nssai) SetDefaultSingleNssais(v []Snssai)`

SetDefaultSingleNssais sets DefaultSingleNssais field to given value.


### GetSingleNssais

`func (o *Nssai) GetSingleNssais() []Snssai`

GetSingleNssais returns the SingleNssais field if non-nil, zero value otherwise.

### GetSingleNssaisOk

`func (o *Nssai) GetSingleNssaisOk() (*[]Snssai, bool)`

GetSingleNssaisOk returns a tuple with the SingleNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssais

`func (o *Nssai) SetSingleNssais(v []Snssai)`

SetSingleNssais sets SingleNssais field to given value.

### HasSingleNssais

`func (o *Nssai) HasSingleNssais() bool`

HasSingleNssais returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *Nssai) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *Nssai) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *Nssai) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.

### HasProvisioningTime

`func (o *Nssai) HasProvisioningTime() bool`

HasProvisioningTime returns a boolean if a field has been set.

### GetAdditionalSnssaiData

`func (o *Nssai) GetAdditionalSnssaiData() map[string]AdditionalSnssaiData`

GetAdditionalSnssaiData returns the AdditionalSnssaiData field if non-nil, zero value otherwise.

### GetAdditionalSnssaiDataOk

`func (o *Nssai) GetAdditionalSnssaiDataOk() (*map[string]AdditionalSnssaiData, bool)`

GetAdditionalSnssaiDataOk returns a tuple with the AdditionalSnssaiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSnssaiData

`func (o *Nssai) SetAdditionalSnssaiData(v map[string]AdditionalSnssaiData)`

SetAdditionalSnssaiData sets AdditionalSnssaiData field to given value.

### HasAdditionalSnssaiData

`func (o *Nssai) HasAdditionalSnssaiData() bool`

HasAdditionalSnssaiData returns a boolean if a field has been set.

### GetSuppressNssrgInd

`func (o *Nssai) GetSuppressNssrgInd() bool`

GetSuppressNssrgInd returns the SuppressNssrgInd field if non-nil, zero value otherwise.

### GetSuppressNssrgIndOk

`func (o *Nssai) GetSuppressNssrgIndOk() (*bool, bool)`

GetSuppressNssrgIndOk returns a tuple with the SuppressNssrgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressNssrgInd

`func (o *Nssai) SetSuppressNssrgInd(v bool)`

SetSuppressNssrgInd sets SuppressNssrgInd field to given value.

### HasSuppressNssrgInd

`func (o *Nssai) HasSuppressNssrgInd() bool`

HasSuppressNssrgInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


