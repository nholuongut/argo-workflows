

# IonholuongutWorkflowV1alpha1RetryStrategy

RetryStrategy provides controls on how to retry a workflow step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affinity** | [**IonholuongutWorkflowV1alpha1RetryAffinity**](IonholuongutWorkflowV1alpha1RetryAffinity.md) |  |  [optional]
**backoff** | [**IonholuongutWorkflowV1alpha1Backoff**](IonholuongutWorkflowV1alpha1Backoff.md) |  |  [optional]
**expression** | **String** | Expression is a condition expression for when a node will be retried. If it evaluates to false, the node will not be retried and the retry strategy will be ignored |  [optional]
**limit** | **String** |  |  [optional]
**retryPolicy** | **String** | RetryPolicy is a policy of NodePhase statuses that will be retried |  [optional]



