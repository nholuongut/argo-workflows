# IonholuongutEventsV1alpha1TriggerTemplate

TriggerTemplate is the template that describes trigger specification.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**argo_workflow** | [**IonholuongutEventsV1alpha1ArgoWorkflowTrigger**](IonholuongutEventsV1alpha1ArgoWorkflowTrigger.md) |  | [optional] 
**aws_lambda** | [**IonholuongutEventsV1alpha1AWSLambdaTrigger**](IonholuongutEventsV1alpha1AWSLambdaTrigger.md) |  | [optional] 
**azure_event_hubs** | [**IonholuongutEventsV1alpha1AzureEventHubsTrigger**](IonholuongutEventsV1alpha1AzureEventHubsTrigger.md) |  | [optional] 
**conditions** | **str** |  | [optional] 
**custom** | [**IonholuongutEventsV1alpha1CustomTrigger**](IonholuongutEventsV1alpha1CustomTrigger.md) |  | [optional] 
**http** | [**IonholuongutEventsV1alpha1HTTPTrigger**](IonholuongutEventsV1alpha1HTTPTrigger.md) |  | [optional] 
**k8s** | [**IonholuongutEventsV1alpha1StandardK8STrigger**](IonholuongutEventsV1alpha1StandardK8STrigger.md) |  | [optional] 
**kafka** | [**IonholuongutEventsV1alpha1KafkaTrigger**](IonholuongutEventsV1alpha1KafkaTrigger.md) |  | [optional] 
**log** | [**IonholuongutEventsV1alpha1LogTrigger**](IonholuongutEventsV1alpha1LogTrigger.md) |  | [optional] 
**name** | **str** | Name is a unique name of the action to take. | [optional] 
**nats** | [**IonholuongutEventsV1alpha1NATSTrigger**](IonholuongutEventsV1alpha1NATSTrigger.md) |  | [optional] 
**open_whisk** | [**IonholuongutEventsV1alpha1OpenWhiskTrigger**](IonholuongutEventsV1alpha1OpenWhiskTrigger.md) |  | [optional] 
**pulsar** | [**IonholuongutEventsV1alpha1PulsarTrigger**](IonholuongutEventsV1alpha1PulsarTrigger.md) |  | [optional] 
**slack** | [**IonholuongutEventsV1alpha1SlackTrigger**](IonholuongutEventsV1alpha1SlackTrigger.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


