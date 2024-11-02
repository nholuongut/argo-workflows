# IonholuongutEventsV1alpha1HTTPTrigger


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**basic_auth** | [**IonholuongutEventsV1alpha1BasicAuth**](IonholuongutEventsV1alpha1BasicAuth.md) |  | [optional] 
**headers** | **{str: (str,)}** |  | [optional] 
**method** | **str** |  | [optional] 
**parameters** | [**[IonholuongutEventsV1alpha1TriggerParameter]**](IonholuongutEventsV1alpha1TriggerParameter.md) | Parameters is the list of key-value extracted from event&#39;s payload that are applied to the HTTP trigger resource. | [optional] 
**payload** | [**[IonholuongutEventsV1alpha1TriggerParameter]**](IonholuongutEventsV1alpha1TriggerParameter.md) |  | [optional] 
**secure_headers** | [**[IonholuongutEventsV1alpha1SecureHeader]**](IonholuongutEventsV1alpha1SecureHeader.md) |  | [optional] 
**timeout** | **str** |  | [optional] 
**tls** | [**IonholuongutEventsV1alpha1TLSConfig**](IonholuongutEventsV1alpha1TLSConfig.md) |  | [optional] 
**url** | **str** | URL refers to the URL to send HTTP request to. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


