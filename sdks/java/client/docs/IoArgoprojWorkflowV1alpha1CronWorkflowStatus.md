

# IonholuongutWorkflowV1alpha1CronWorkflowStatus

CronWorkflowStatus is the status of a CronWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active** | [**List&lt;io.kubernetes.client.openapi.models.V1ObjectReference&gt;**](io.kubernetes.client.openapi.models.V1ObjectReference.md) | Active is a list of active workflows stemming from this CronWorkflow | 
**conditions** | [**List&lt;IonholuongutWorkflowV1alpha1Condition&gt;**](IonholuongutWorkflowV1alpha1Condition.md) | Conditions is a list of conditions the CronWorkflow may have | 
**lastScheduledTime** | **java.time.Instant** |  | 



