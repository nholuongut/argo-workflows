from pprint import pprint

import argo_workflows
from argo_workflows.api import workflow_service_api
from argo_workflows.model.io_nholuongut_workflow_v1alpha1_workflow_create_request import \
    IonholuongutWorkflowV1alpha1WorkflowCreateRequest
import requests
import yaml

configuration = argo_workflows.Configuration(host="https://127.0.0.1:2746")
configuration.verify_ssl = False

resp = requests.get('https://raw.githubusercontent.com/nholuongut/argo-workflows/master/examples/hello-world.yaml')
manifest = yaml.safe_load(resp.text)

api_client = argo_workflows.ApiClient(configuration)
api_instance = workflow_service_api.WorkflowServiceApi(api_client)
api_response = api_instance.create_workflow(
    namespace='argo',
    body=IonholuongutWorkflowV1alpha1WorkflowCreateRequest(workflow=manifest, _check_return_type=False, _check_type=False))
pprint(api_response)
