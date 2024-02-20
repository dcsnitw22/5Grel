# ModificationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyItems** | [**[]NotifyItem**](NotifyItem.md) |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewModificationNotification

`func NewModificationNotification(notifyItems []NotifyItem, ) *ModificationNotification`

NewModificationNotification instantiates a new ModificationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModificationNotificationWithDefaults

`func NewModificationNotificationWithDefaults() *ModificationNotification`

NewModificationNotificationWithDefaults instantiates a new ModificationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyItems

`func (o *ModificationNotification) GetNotifyItems() []NotifyItem`

GetNotifyItems returns the NotifyItems field if non-nil, zero value otherwise.

### GetNotifyItemsOk

`func (o *ModificationNotification) GetNotifyItemsOk() (*[]NotifyItem, bool)`

GetNotifyItemsOk returns a tuple with the NotifyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyItems

`func (o *ModificationNotification) SetNotifyItems(v []NotifyItem)`

SetNotifyItems sets NotifyItems field to given value.


### GetSubscriptionId

`func (o *ModificationNotification) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ModificationNotification) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ModificationNotification) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ModificationNotification) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


