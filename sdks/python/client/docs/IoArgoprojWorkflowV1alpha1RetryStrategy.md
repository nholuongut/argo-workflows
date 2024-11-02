# IonholuongutWorkflowV1alpha1RetryStrategy

RetryStrategy provides controls on how to retry a workflow step

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affinity** | [**IonholuongutWorkflowV1alpha1RetryAffinity**](IonholuongutWorkflowV1alpha1RetryAffinity.md) |  | [optional] 
**backoff** | [**IonholuongutWorkflowV1alpha1Backoff**](IonholuongutWorkflowV1alpha1Backoff.md) |  | [optional] 
**expression** | **str** | Expression is a condition expression for when a node will be retried. If it evaluates to false, the node will not be retried and the retry strategy will be ignored | [optional] 
**limit** | **str** |  | [optional] 
**retry_policy** | **str** | RetryPolicy is a policy of NodePhase statuses that will be retried | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


