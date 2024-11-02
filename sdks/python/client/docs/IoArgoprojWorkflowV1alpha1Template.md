# IonholuongutWorkflowV1alpha1Template

Template is a reusable and composable unit of execution in a workflow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active_deadline_seconds** | **str** |  | [optional] 
**affinity** | [**Affinity**](Affinity.md) |  | [optional] 
**archive_location** | [**IonholuongutWorkflowV1alpha1ArtifactLocation**](IonholuongutWorkflowV1alpha1ArtifactLocation.md) |  | [optional] 
**automount_service_account_token** | **bool** | AutomountServiceAccountToken indicates whether a service account token should be automatically mounted in pods. ServiceAccountName of ExecutorConfig must be specified if this value is false. | [optional] 
**container** | [**Container**](Container.md) |  | [optional] 
**container_set** | [**IonholuongutWorkflowV1alpha1ContainerSetTemplate**](IonholuongutWorkflowV1alpha1ContainerSetTemplate.md) |  | [optional] 
**daemon** | **bool** | Deamon will allow a workflow to proceed to the next step so long as the container reaches readiness | [optional] 
**dag** | [**IonholuongutWorkflowV1alpha1DAGTemplate**](IonholuongutWorkflowV1alpha1DAGTemplate.md) |  | [optional] 
**data** | [**IonholuongutWorkflowV1alpha1Data**](IonholuongutWorkflowV1alpha1Data.md) |  | [optional] 
**executor** | [**IonholuongutWorkflowV1alpha1ExecutorConfig**](IonholuongutWorkflowV1alpha1ExecutorConfig.md) |  | [optional] 
**fail_fast** | **bool** | FailFast, if specified, will fail this template if any of its child pods has failed. This is useful for when this template is expanded with &#x60;withItems&#x60;, etc. | [optional] 
**host_aliases** | [**[HostAlias]**](HostAlias.md) | HostAliases is an optional list of hosts and IPs that will be injected into the pod spec | [optional] 
**http** | [**IonholuongutWorkflowV1alpha1HTTP**](IonholuongutWorkflowV1alpha1HTTP.md) |  | [optional] 
**init_containers** | [**[IonholuongutWorkflowV1alpha1UserContainer]**](IonholuongutWorkflowV1alpha1UserContainer.md) | InitContainers is a list of containers which run before the main container. | [optional] 
**inputs** | [**IonholuongutWorkflowV1alpha1Inputs**](IonholuongutWorkflowV1alpha1Inputs.md) |  | [optional] 
**memoize** | [**IonholuongutWorkflowV1alpha1Memoize**](IonholuongutWorkflowV1alpha1Memoize.md) |  | [optional] 
**metadata** | [**IonholuongutWorkflowV1alpha1Metadata**](IonholuongutWorkflowV1alpha1Metadata.md) |  | [optional] 
**metrics** | [**IonholuongutWorkflowV1alpha1Metrics**](IonholuongutWorkflowV1alpha1Metrics.md) |  | [optional] 
**name** | **str** | Name is the name of the template | [optional] 
**node_selector** | **{str: (str,)}** | NodeSelector is a selector to schedule this step of the workflow to be run on the selected node(s). Overrides the selector set at the workflow level. | [optional] 
**outputs** | [**IonholuongutWorkflowV1alpha1Outputs**](IonholuongutWorkflowV1alpha1Outputs.md) |  | [optional] 
**parallelism** | **int** | Parallelism limits the max total parallel pods that can execute at the same time within the boundaries of this template invocation. If additional steps/dag templates are invoked, the pods created by those templates will not be counted towards this total. | [optional] 
**pod_spec_patch** | **str** | PodSpecPatch holds strategic merge patch to apply against the pod spec. Allows parameterization of container fields which are not strings (e.g. resource limits). | [optional] 
**priority** | **int** | Priority to apply to workflow pods. | [optional] 
**priority_class_name** | **str** | PriorityClassName to apply to workflow pods. | [optional] 
**resource** | [**IonholuongutWorkflowV1alpha1ResourceTemplate**](IonholuongutWorkflowV1alpha1ResourceTemplate.md) |  | [optional] 
**retry_strategy** | [**IonholuongutWorkflowV1alpha1RetryStrategy**](IonholuongutWorkflowV1alpha1RetryStrategy.md) |  | [optional] 
**scheduler_name** | **str** | If specified, the pod will be dispatched by specified scheduler. Or it will be dispatched by workflow scope scheduler if specified. If neither specified, the pod will be dispatched by default scheduler. | [optional] 
**script** | [**IonholuongutWorkflowV1alpha1ScriptTemplate**](IonholuongutWorkflowV1alpha1ScriptTemplate.md) |  | [optional] 
**security_context** | [**PodSecurityContext**](PodSecurityContext.md) |  | [optional] 
**service_account_name** | **str** | ServiceAccountName to apply to workflow pods | [optional] 
**sidecars** | [**[IonholuongutWorkflowV1alpha1UserContainer]**](IonholuongutWorkflowV1alpha1UserContainer.md) | Sidecars is a list of containers which run alongside the main container Sidecars are automatically killed when the main container completes | [optional] 
**steps** | [**[IonholuongutWorkflowV1alpha1ParallelSteps]**](IonholuongutWorkflowV1alpha1ParallelSteps.md) | Steps define a series of sequential/parallel workflow steps | [optional] 
**suspend** | [**IonholuongutWorkflowV1alpha1SuspendTemplate**](IonholuongutWorkflowV1alpha1SuspendTemplate.md) |  | [optional] 
**synchronization** | [**IonholuongutWorkflowV1alpha1Synchronization**](IonholuongutWorkflowV1alpha1Synchronization.md) |  | [optional] 
**timeout** | **str** | Timeout allows to set the total node execution timeout duration counting from the node&#39;s start time. This duration also includes time in which the node spends in Pending state. This duration may not be applied to Step or DAG templates. | [optional] 
**tolerations** | [**[Toleration]**](Toleration.md) | Tolerations to apply to workflow pods. | [optional] 
**volumes** | [**[Volume]**](Volume.md) | Volumes is a list of volumes that can be mounted by containers in a template. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


