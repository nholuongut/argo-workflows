

# IonholuongutWorkflowV1alpha1Template

Template is a reusable and composable unit of execution in a workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**activeDeadlineSeconds** | **String** |  |  [optional]
**affinity** | [**io.kubernetes.client.openapi.models.V1Affinity**](io.kubernetes.client.openapi.models.V1Affinity.md) |  |  [optional]
**archiveLocation** | [**IonholuongutWorkflowV1alpha1ArtifactLocation**](IonholuongutWorkflowV1alpha1ArtifactLocation.md) |  |  [optional]
**automountServiceAccountToken** | **Boolean** | AutomountServiceAccountToken indicates whether a service account token should be automatically mounted in pods. ServiceAccountName of ExecutorConfig must be specified if this value is false. |  [optional]
**container** | [**io.kubernetes.client.openapi.models.V1Container**](io.kubernetes.client.openapi.models.V1Container.md) |  |  [optional]
**containerSet** | [**IonholuongutWorkflowV1alpha1ContainerSetTemplate**](IonholuongutWorkflowV1alpha1ContainerSetTemplate.md) |  |  [optional]
**daemon** | **Boolean** | Deamon will allow a workflow to proceed to the next step so long as the container reaches readiness |  [optional]
**dag** | [**IonholuongutWorkflowV1alpha1DAGTemplate**](IonholuongutWorkflowV1alpha1DAGTemplate.md) |  |  [optional]
**data** | [**IonholuongutWorkflowV1alpha1Data**](IonholuongutWorkflowV1alpha1Data.md) |  |  [optional]
**executor** | [**IonholuongutWorkflowV1alpha1ExecutorConfig**](IonholuongutWorkflowV1alpha1ExecutorConfig.md) |  |  [optional]
**failFast** | **Boolean** | FailFast, if specified, will fail this template if any of its child pods has failed. This is useful for when this template is expanded with &#x60;withItems&#x60;, etc. |  [optional]
**hostAliases** | [**List&lt;io.kubernetes.client.openapi.models.V1HostAlias&gt;**](io.kubernetes.client.openapi.models.V1HostAlias.md) | HostAliases is an optional list of hosts and IPs that will be injected into the pod spec |  [optional]
**http** | [**IonholuongutWorkflowV1alpha1HTTP**](IonholuongutWorkflowV1alpha1HTTP.md) |  |  [optional]
**initContainers** | [**List&lt;IonholuongutWorkflowV1alpha1UserContainer&gt;**](IonholuongutWorkflowV1alpha1UserContainer.md) | InitContainers is a list of containers which run before the main container. |  [optional]
**inputs** | [**IonholuongutWorkflowV1alpha1Inputs**](IonholuongutWorkflowV1alpha1Inputs.md) |  |  [optional]
**memoize** | [**IonholuongutWorkflowV1alpha1Memoize**](IonholuongutWorkflowV1alpha1Memoize.md) |  |  [optional]
**metadata** | [**IonholuongutWorkflowV1alpha1Metadata**](IonholuongutWorkflowV1alpha1Metadata.md) |  |  [optional]
**metrics** | [**IonholuongutWorkflowV1alpha1Metrics**](IonholuongutWorkflowV1alpha1Metrics.md) |  |  [optional]
**name** | **String** | Name is the name of the template |  [optional]
**nodeSelector** | **Map&lt;String, String&gt;** | NodeSelector is a selector to schedule this step of the workflow to be run on the selected node(s). Overrides the selector set at the workflow level. |  [optional]
**outputs** | [**IonholuongutWorkflowV1alpha1Outputs**](IonholuongutWorkflowV1alpha1Outputs.md) |  |  [optional]
**parallelism** | **Integer** | Parallelism limits the max total parallel pods that can execute at the same time within the boundaries of this template invocation. If additional steps/dag templates are invoked, the pods created by those templates will not be counted towards this total. |  [optional]
**podSpecPatch** | **String** | PodSpecPatch holds strategic merge patch to apply against the pod spec. Allows parameterization of container fields which are not strings (e.g. resource limits). |  [optional]
**priority** | **Integer** | Priority to apply to workflow pods. |  [optional]
**priorityClassName** | **String** | PriorityClassName to apply to workflow pods. |  [optional]
**resource** | [**IonholuongutWorkflowV1alpha1ResourceTemplate**](IonholuongutWorkflowV1alpha1ResourceTemplate.md) |  |  [optional]
**retryStrategy** | [**IonholuongutWorkflowV1alpha1RetryStrategy**](IonholuongutWorkflowV1alpha1RetryStrategy.md) |  |  [optional]
**schedulerName** | **String** | If specified, the pod will be dispatched by specified scheduler. Or it will be dispatched by workflow scope scheduler if specified. If neither specified, the pod will be dispatched by default scheduler. |  [optional]
**script** | [**IonholuongutWorkflowV1alpha1ScriptTemplate**](IonholuongutWorkflowV1alpha1ScriptTemplate.md) |  |  [optional]
**securityContext** | [**io.kubernetes.client.openapi.models.V1PodSecurityContext**](io.kubernetes.client.openapi.models.V1PodSecurityContext.md) |  |  [optional]
**serviceAccountName** | **String** | ServiceAccountName to apply to workflow pods |  [optional]
**sidecars** | [**List&lt;IonholuongutWorkflowV1alpha1UserContainer&gt;**](IonholuongutWorkflowV1alpha1UserContainer.md) | Sidecars is a list of containers which run alongside the main container Sidecars are automatically killed when the main container completes |  [optional]
**steps** | **List&lt;IonholuongutWorkflowV1alpha1ParallelSteps&gt;** | Steps define a series of sequential/parallel workflow steps |  [optional]
**suspend** | [**IonholuongutWorkflowV1alpha1SuspendTemplate**](IonholuongutWorkflowV1alpha1SuspendTemplate.md) |  |  [optional]
**synchronization** | [**IonholuongutWorkflowV1alpha1Synchronization**](IonholuongutWorkflowV1alpha1Synchronization.md) |  |  [optional]
**timeout** | **String** | Timeout allows to set the total node execution timeout duration counting from the node&#39;s start time. This duration also includes time in which the node spends in Pending state. This duration may not be applied to Step or DAG templates. |  [optional]
**tolerations** | [**List&lt;io.kubernetes.client.openapi.models.V1Toleration&gt;**](io.kubernetes.client.openapi.models.V1Toleration.md) | Tolerations to apply to workflow pods. |  [optional]
**volumes** | [**List&lt;io.kubernetes.client.openapi.models.V1Volume&gt;**](io.kubernetes.client.openapi.models.V1Volume.md) | Volumes is a list of volumes that can be mounted by containers in a template. |  [optional]



