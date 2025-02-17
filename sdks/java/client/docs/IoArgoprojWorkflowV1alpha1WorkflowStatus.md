

# IonholuongutWorkflowV1alpha1WorkflowStatus

WorkflowStatus contains overall status information about a workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**artifactRepositoryRef** | [**IonholuongutWorkflowV1alpha1ArtifactRepositoryRefStatus**](IonholuongutWorkflowV1alpha1ArtifactRepositoryRefStatus.md) |  |  [optional]
**compressedNodes** | **String** | Compressed and base64 decoded Nodes map |  [optional]
**conditions** | [**List&lt;IonholuongutWorkflowV1alpha1Condition&gt;**](IonholuongutWorkflowV1alpha1Condition.md) | Conditions is a list of conditions the Workflow may have |  [optional]
**estimatedDuration** | **Integer** | EstimatedDuration in seconds. |  [optional]
**finishedAt** | **java.time.Instant** |  |  [optional]
**message** | **String** | A human readable message indicating details about why the workflow is in this condition. |  [optional]
**nodes** | [**Map&lt;String, IonholuongutWorkflowV1alpha1NodeStatus&gt;**](IonholuongutWorkflowV1alpha1NodeStatus.md) | Nodes is a mapping between a node ID and the node&#39;s status. |  [optional]
**offloadNodeStatusVersion** | **String** | Whether on not node status has been offloaded to a database. If exists, then Nodes and CompressedNodes will be empty. This will actually be populated with a hash of the offloaded data. |  [optional]
**outputs** | [**IonholuongutWorkflowV1alpha1Outputs**](IonholuongutWorkflowV1alpha1Outputs.md) |  |  [optional]
**persistentVolumeClaims** | [**List&lt;io.kubernetes.client.openapi.models.V1Volume&gt;**](io.kubernetes.client.openapi.models.V1Volume.md) | PersistentVolumeClaims tracks all PVCs that were created as part of the io.nholuongut.workflow.v1alpha1. The contents of this list are drained at the end of the workflow. |  [optional]
**phase** | **String** | Phase a simple, high-level summary of where the workflow is in its lifecycle. |  [optional]
**progress** | **String** | Progress to completion |  [optional]
**resourcesDuration** | **Map&lt;String, Long&gt;** | ResourcesDuration is the total for the workflow |  [optional]
**startedAt** | **java.time.Instant** |  |  [optional]
**storedTemplates** | [**Map&lt;String, IonholuongutWorkflowV1alpha1Template&gt;**](IonholuongutWorkflowV1alpha1Template.md) | StoredTemplates is a mapping between a template ref and the node&#39;s status. |  [optional]
**storedWorkflowTemplateSpec** | [**IonholuongutWorkflowV1alpha1WorkflowSpec**](IonholuongutWorkflowV1alpha1WorkflowSpec.md) |  |  [optional]
**synchronization** | [**IonholuongutWorkflowV1alpha1SynchronizationStatus**](IonholuongutWorkflowV1alpha1SynchronizationStatus.md) |  |  [optional]



