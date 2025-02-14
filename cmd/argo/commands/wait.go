package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/nholuongut/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nholuongut/argo-workflows/v3/cmd/argo/commands/client"
	workflowpkg "github.com/nholuongut/argo-workflows/v3/pkg/apiclient/workflow"
	wfv1 "github.com/nholuongut/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/nholuongut/argo-workflows/v3/util"
)

func NewWaitCommand() *cobra.Command {
	var ignoreNotFound bool
	command := &cobra.Command{
		Use:   "wait [WORKFLOW...]",
		Short: "waits for workflows to complete",
		Example: `# Wait on a workflow:

  argo wait my-wf

# Wait on the latest workflow:

  argo wait @latest
`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, apiClient := client.NewAPIClient(cmd.Context())
			serviceClient := apiClient.NewWorkflowServiceClient()
			namespace := client.Namespace()
			waitWorkflows(ctx, serviceClient, namespace, args, ignoreNotFound, false)
		},
	}
	command.Flags().BoolVar(&ignoreNotFound, "ignore-not-found", false, "Ignore the wait if the workflow is not found")
	return command
}

// waitWorkflows waits for the given workflowNames.
func waitWorkflows(ctx context.Context, serviceClient workflowpkg.WorkflowServiceClient, namespace string, workflowNames []string, ignoreNotFound, quiet bool) {
	var wg sync.WaitGroup
	wfSuccessStatus := true

	for _, name := range workflowNames {
		wg.Add(1)
		go func(name string) {
			if !waitOnOne(serviceClient, ctx, name, namespace, ignoreNotFound, quiet) {
				wfSuccessStatus = false
			}
			wg.Done()
		}(name)

	}
	wg.Wait()

	if !wfSuccessStatus {
		os.Exit(1)
	}
}

func waitOnOne(serviceClient workflowpkg.WorkflowServiceClient, ctx context.Context, wfName, namespace string, ignoreNotFound, quiet bool) bool {
	req := &workflowpkg.WatchWorkflowsRequest{
		Namespace: namespace,
		ListOptions: &metav1.ListOptions{
			FieldSelector:   util.GenerateFieldSelectorFromWorkflowName(wfName),
			ResourceVersion: "0",
		},
	}
	stream, err := serviceClient.WatchWorkflows(ctx, req)
	if err != nil {
		if status.Code(err) == codes.NotFound && ignoreNotFound {
			return true
		}
		errors.CheckError(err)
		return false
	}
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			log.Debug("Re-establishing workflow watch")
			stream, err = serviceClient.WatchWorkflows(ctx, req)
			errors.CheckError(err)
			continue
		}
		errors.CheckError(err)
		if event == nil {
			continue
		}
		wf := event.Object
		if !wf.Status.FinishedAt.IsZero() {
			if !quiet {
				fmt.Printf("%s %s at %v\n", wfName, wf.Status.Phase, wf.Status.FinishedAt)
			}
			if wf.Status.Phase == wfv1.WorkflowFailed || wf.Status.Phase == wfv1.WorkflowError {
				return false
			}
			return true
		}
	}
}
