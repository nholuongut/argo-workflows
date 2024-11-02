# IonholuongutWorkflowV1alpha1WorkflowStatus

WorkflowStatus contains overall status information about a workflow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**artifact_repository_ref** | [**IonholuongutWorkflowV1alpha1ArtifactRepositoryRefStatus**](IonholuongutWorkflowV1alpha1ArtifactRepositoryRefStatus.md) |  | [optional] 
**compressed_nodes** | **str** | Compressed and base64 decoded Nodes map | [optional] 
**conditions** | [**[IonholuongutWorkflowV1alpha1Condition]**](IonholuongutWorkflowV1alpha1Condition.md) | Conditions is a list of conditions the Workflow may have | [optional] 
**estimated_duration** | **int** | EstimatedDuration in seconds. | [optional] 
**finished_at** | **datetime** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**message** | **str** | A human readable message indicating details about why the workflow is in this condition. | [optional] 
**nodes** | [**{str: (IonholuongutWorkflowV1alpha1NodeStatus,)}**](IonholuongutWorkflowV1alpha1NodeStatus.md) | Nodes is a mapping between a node ID and the node&#39;s status. | [optional] 
**offload_node_status_version** | **str** | Whether on not node status has been offloaded to a database. If exists, then Nodes and CompressedNodes will be empty. This will actually be populated with a hash of the offloaded data. | [optional] 
**outputs** | [**IonholuongutWorkflowV1alpha1Outputs**](IonholuongutWorkflowV1alpha1Outputs.md) |  | [optional] 
**persistent_volume_claims** | [**[Volume]**](Volume.md) | PersistentVolumeClaims tracks all PVCs that were created as part of the io.nholuongut.workflow.v1alpha1. The contents of this list are drained at the end of the workflow. | [optional] 
**phase** | **str** | Phase a simple, high-level summary of where the workflow is in its lifecycle. | [optional] 
**progress** | **str** | Progress to completion | [optional] 
**resources_duration** | **{str: (int,)}** | ResourcesDuration is the total for the workflow | [optional] 
**started_at** | **datetime** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**stored_templates** | [**{str: (IonholuongutWorkflowV1alpha1Template,)}**](IonholuongutWorkflowV1alpha1Template.md) | StoredTemplates is a mapping between a template ref and the node&#39;s status. | [optional] 
**stored_workflow_template_spec** | [**IonholuongutWorkflowV1alpha1WorkflowSpec**](IonholuongutWorkflowV1alpha1WorkflowSpec.md) |  | [optional] 
**synchronization** | [**IonholuongutWorkflowV1alpha1SynchronizationStatus**](IonholuongutWorkflowV1alpha1SynchronizationStatus.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


