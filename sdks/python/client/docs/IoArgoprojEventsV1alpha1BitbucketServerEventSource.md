# IonholuongutEventsV1alpha1BitbucketServerEventSource


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_token** | [**SecretKeySelector**](SecretKeySelector.md) |  | [optional] 
**bitbucketserver_base_url** | **str** |  | [optional] 
**delete_hook_on_finish** | **bool** |  | [optional] 
**events** | **[str]** |  | [optional] 
**metadata** | **{str: (str,)}** |  | [optional] 
**project_key** | **str** |  | [optional] 
**repository_slug** | **str** |  | [optional] 
**webhook** | [**IonholuongutEventsV1alpha1WebhookContext**](IonholuongutEventsV1alpha1WebhookContext.md) |  | [optional] 
**webhook_secret** | [**SecretKeySelector**](SecretKeySelector.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


