# SmfSelectionSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SubscribedSnssaiInfos** | Pointer to [**map[string]SnssaiInfo**](SnssaiInfo.md) | A map(list of key-value pairs) where singleNssai serves as key of SnssaiInfo | [optional] 
**SharedSnssaiInfosId** | Pointer to **string** |  | [optional] 
**HssGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 

## Methods

### NewSmfSelectionSubscriptionData

`func NewSmfSelectionSubscriptionData() *SmfSelectionSubscriptionData`

NewSmfSelectionSubscriptionData instantiates a new SmfSelectionSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfSelectionSubscriptionDataWithDefaults

`func NewSmfSelectionSubscriptionDataWithDefaults() *SmfSelectionSubscriptionData`

NewSmfSelectionSubscriptionDataWithDefaults instantiates a new SmfSelectionSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *SmfSelectionSubscriptionData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmfSelectionSubscriptionData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmfSelectionSubscriptionData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmfSelectionSubscriptionData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSubscribedSnssaiInfos

`func (o *SmfSelectionSubscriptionData) GetSubscribedSnssaiInfos() map[string]SnssaiInfo`

GetSubscribedSnssaiInfos returns the SubscribedSnssaiInfos field if non-nil, zero value otherwise.

### GetSubscribedSnssaiInfosOk

`func (o *SmfSelectionSubscriptionData) GetSubscribedSnssaiInfosOk() (*map[string]SnssaiInfo, bool)`

GetSubscribedSnssaiInfosOk returns a tuple with the SubscribedSnssaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedSnssaiInfos

`func (o *SmfSelectionSubscriptionData) SetSubscribedSnssaiInfos(v map[string]SnssaiInfo)`

SetSubscribedSnssaiInfos sets SubscribedSnssaiInfos field to given value.

### HasSubscribedSnssaiInfos

`func (o *SmfSelectionSubscriptionData) HasSubscribedSnssaiInfos() bool`

HasSubscribedSnssaiInfos returns a boolean if a field has been set.

### GetSharedSnssaiInfosId

`func (o *SmfSelectionSubscriptionData) GetSharedSnssaiInfosId() string`

GetSharedSnssaiInfosId returns the SharedSnssaiInfosId field if non-nil, zero value otherwise.

### GetSharedSnssaiInfosIdOk

`func (o *SmfSelectionSubscriptionData) GetSharedSnssaiInfosIdOk() (*string, bool)`

GetSharedSnssaiInfosIdOk returns a tuple with the SharedSnssaiInfosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSnssaiInfosId

`func (o *SmfSelectionSubscriptionData) SetSharedSnssaiInfosId(v string)`

SetSharedSnssaiInfosId sets SharedSnssaiInfosId field to given value.

### HasSharedSnssaiInfosId

`func (o *SmfSelectionSubscriptionData) HasSharedSnssaiInfosId() bool`

HasSharedSnssaiInfosId returns a boolean if a field has been set.

### GetHssGroupId

`func (o *SmfSelectionSubscriptionData) GetHssGroupId() string`

GetHssGroupId returns the HssGroupId field if non-nil, zero value otherwise.

### GetHssGroupIdOk

`func (o *SmfSelectionSubscriptionData) GetHssGroupIdOk() (*string, bool)`

GetHssGroupIdOk returns a tuple with the HssGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssGroupId

`func (o *SmfSelectionSubscriptionData) SetHssGroupId(v string)`

SetHssGroupId sets HssGroupId field to given value.

### HasHssGroupId

`func (o *SmfSelectionSubscriptionData) HasHssGroupId() bool`

HasHssGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


