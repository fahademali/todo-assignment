package client

import (
	"context"
	"todo_service/log"
	"user_service/config"

	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
)

func buildWorkflowServiceClient() workflowserviceclient.Interface {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(config.AppConfig.CADENCE_CLIENT_NAME))
	if err != nil {
		panic("Failed to setup tchannel")
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: config.AppConfig.CADENCE_CLIENT_NAME,
		Outbounds: yarpc.Outbounds{
			config.AppConfig.CADENCE_SERVICE: {Unary: ch.NewSingleOutbound(config.AppConfig.CADENCE_HOST_PORT)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(config.AppConfig.CADENCE_SERVICE))
}

func buildCadenceClient(domain string) client.Client {
	return client.NewClient(buildWorkflowServiceClient(), domain, &client.Options{})
}

func BeginWorkflows(ctx context.Context, beginWorkflowOpts client.StartWorkflowOptions, workflows ...interface{}) {
	cadenceClient := buildCadenceClient(config.AppConfig.CADENCE_DOMAIN)
	cadenceClient.StartWorkflow(ctx, beginWorkflowOpts, "worker/workflows.RemindUsersForDueDateWorkflow")
	// cadenceClient.StartWorkflow(ctx, beginWorkflowOpts)
	log.GetLog().Info("running BeginWorkflows.............")
	log.GetLog().Info(cadenceClient)
	log.GetLog().Info(workflows...)
}
