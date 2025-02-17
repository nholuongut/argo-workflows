package commands

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"

	"github.com/nholuongut/argo-workflows/v3"
	"github.com/nholuongut/argo-workflows/v3/util/logs"
	"github.com/nholuongut/argo-workflows/v3/workflow/common"
	"github.com/nholuongut/argo-workflows/v3/workflow/executor"
)

func NewAgentCommand() *cobra.Command {
	return &cobra.Command{
		Use:          "agent",
		SilenceUsage: true, // this prevents confusing usage message being printed on error
		RunE: func(cmd *cobra.Command, args []string) error {
			return initAgentExecutor().Agent(context.Background())
		},
	}
}

func initAgentExecutor() *executor.AgentExecutor {
	version := argo.GetVersion()
	log.WithFields(log.Fields{"version": version.Version}).Info("Starting Workflow Executor")
	config, err := clientConfig.ClientConfig()
	checkErr(err)

	config = restclient.AddUserAgent(config, fmt.Sprintf("argo-workflows/%s argo-executor/%s", version.Version, "agent Executor"))

	logs.AddK8SLogTransportWrapper(config) // lets log all request as we should typically do < 5 per pod, so this is will show up problems

	namespace, _, err := clientConfig.Namespace()
	checkErr(err)

	clientSet, err := kubernetes.NewForConfig(config)
	checkErr(err)

	restClient := clientSet.RESTClient()

	workflowName, ok := os.LookupEnv(common.EnvVarWorkflowName)
	if !ok {
		log.Fatalf("Unable to determine workflow name from environment variable %s", common.EnvVarWorkflowName)
	}
	return executor.NewAgentExecutor(clientSet, restClient, config, namespace, workflowName)
}
