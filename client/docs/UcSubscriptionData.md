# UcSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserConsentPerPurposeList** | Pointer to [**map[string]UserConsent**](UserConsent.md) | A map(list of key-value pairs) where user consent purpose serves as key of user consent | [optional] 

## Methods

### NewUcSubscriptionData

`func NewUcSubscriptionData() *UcSubscriptionData`

NewUcSubscriptionData instantiates a new UcSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUcSubscriptionDataWithDefaults

`func NewUcSubscriptionDataWithDefaults() *UcSubscriptionData`

NewUcSubscriptionDataWithDefaults instantiates a new UcSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserConsentPerPurposeList

`func (o *UcSubscriptionData) GetUserConsentPerPurposeList() map[string]UserConsent`

GetUserConsentPerPurposeList returns the UserConsentPerPurposeList field if non-nil, zero value otherwise.

### GetUserConsentPerPurposeListOk

`func (o *UcSubscriptionData) GetUserConsentPerPurposeListOk() (*map[string]UserConsent, bool)`

GetUserConsentPerPurposeListOk returns a tuple with the UserConsentPerPurposeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConsentPerPurposeList

`func (o *UcSubscriptionData) SetUserConsentPerPurposeList(v map[string]UserConsent)`

SetUserConsentPerPurposeList sets UserConsentPerPurposeList field to given value.

### HasUserConsentPerPurposeList

`func (o *UcSubscriptionData) HasUserConsentPerPurposeList() bool`

HasUserConsentPerPurposeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


