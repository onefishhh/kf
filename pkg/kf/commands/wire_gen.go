// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject
// +build !wireinject

package commands

import (
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/kf/v2/pkg/client/kf/clientset/versioned/typed/kf/v1alpha1"
	"github.com/google/kf/v2/pkg/kf/apps"
	"github.com/google/kf/v2/pkg/kf/buildpacks"
	"github.com/google/kf/v2/pkg/kf/builds"
	apps2 "github.com/google/kf/v2/pkg/kf/commands/apps"
	"github.com/google/kf/v2/pkg/kf/commands/autoscaling"
	buildpacks2 "github.com/google/kf/v2/pkg/kf/commands/buildpacks"
	builds2 "github.com/google/kf/v2/pkg/kf/commands/builds"
	cluster2 "github.com/google/kf/v2/pkg/kf/commands/cluster"
	"github.com/google/kf/v2/pkg/kf/commands/config"
	"github.com/google/kf/v2/pkg/kf/commands/dependencies"
	"github.com/google/kf/v2/pkg/kf/commands/exporttok8s"
	logs2 "github.com/google/kf/v2/pkg/kf/commands/logs"
	"github.com/google/kf/v2/pkg/kf/commands/networkpolicies"
	"github.com/google/kf/v2/pkg/kf/commands/routes"
	"github.com/google/kf/v2/pkg/kf/commands/service-bindings"
	"github.com/google/kf/v2/pkg/kf/commands/service-brokers"
	"github.com/google/kf/v2/pkg/kf/commands/services"
	"github.com/google/kf/v2/pkg/kf/commands/spaces"
	tasks2 "github.com/google/kf/v2/pkg/kf/commands/tasks"
	"github.com/google/kf/v2/pkg/kf/commands/taskschedules"
	"github.com/google/kf/v2/pkg/kf/configmaps"
	"github.com/google/kf/v2/pkg/kf/logs"
	"github.com/google/kf/v2/pkg/kf/marketplace"
	routes2 "github.com/google/kf/v2/pkg/kf/routes"
	"github.com/google/kf/v2/pkg/kf/secrets"
	"github.com/google/kf/v2/pkg/kf/service-brokers/cluster"
	"github.com/google/kf/v2/pkg/kf/service-brokers/namespaced"
	"github.com/google/kf/v2/pkg/kf/serviceinstancebindings"
	"github.com/google/kf/v2/pkg/kf/serviceinstances"
	"github.com/google/kf/v2/pkg/kf/sourcepackages"
	spaces2 "github.com/google/kf/v2/pkg/kf/spaces"
	"github.com/google/kf/v2/pkg/kf/tasks"
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

import (
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

// Injectors from wire_injector.go:

func InjectPush(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	serviceInstanceBindingsGetter := provideServiceInstanceBindingsGetter(kfV1alpha1Interface)
	serviceinstancebindingsClient := serviceinstancebindings.NewClient(serviceInstanceBindingsGetter)
	secretsGetter := provideSecretsGetter(kubernetesInterface)
	secretsClient := secrets.NewClient(secretsGetter)
	sourcePackagesGetter := provideSourcePackagesGetter(kfV1alpha1Interface)
	poster := sourcepackages.NewPoster(kubernetesInterface)
	sourcepackagesClient := sourcepackages.NewClient(sourcePackagesGetter, poster)
	pusher := apps.NewPusher(appsClient, serviceinstancebindingsClient, secretsClient, sourcepackagesClient, poster)
	srcImageBuilder := provideSrcImageBuilder()
	command := apps2.NewPushCommand(p, pusher, srcImageBuilder)
	return command
}

func InjectDelete(p *config.KfParams) *cobra.Command {
	command := apps2.NewDeleteCommand(p)
	return command
}

func InjectApps(p *config.KfParams) *cobra.Command {
	command := apps2.NewAppsCommand(p)
	return command
}

func InjectGetApp(p *config.KfParams) *cobra.Command {
	command := apps2.NewGetAppCommand(p)
	return command
}

func InjectScale(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewScaleCommand(p, appsClient)
	return command
}

func InjectCreateAutoscalingRule(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := autoscaling.NewCreateAutoscalingRule(p, appsClient)
	return command
}

func InjectDeleteAutoscalingRules(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := autoscaling.NewDeleteAutoscalingRules(p, appsClient)
	return command
}

func InjectUpdateAutoscalingLimits(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := autoscaling.NewUpdateAutoscalingLimits(p, appsClient)
	return command
}

func InjectEnableAutoscale(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := autoscaling.NewEnableAutoscaling(p, appsClient)
	return command
}

func InjectDisableAutoscale(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := autoscaling.NewDisableAutoscaling(p, appsClient)
	return command
}

func InjectStart(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewStartCommand(p, appsClient)
	return command
}

func InjectStop(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewStopCommand(p, appsClient)
	return command
}

func InjectRestart(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewRestartCommand(p, appsClient)
	return command
}

func InjectRestage(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewRestageCommand(p, appsClient)
	return command
}

func InjectProxy(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewProxyCommand(p, appsClient)
	return command
}

func InjectLogs(p *config.KfParams) *cobra.Command {
	kubernetesInterface := config.GetKubernetes(p)
	tailer := logs.NewTailer(kubernetesInterface)
	command := logs2.NewLogsCommand(p, tailer)
	return command
}

func InjectSSH(p *config.KfParams) *cobra.Command {
	command := apps2.NewSSHCommand(p)
	return command
}

func InjectEnv(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewEnvCommand(p, appsClient)
	return command
}

func InjectSetEnv(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewSetEnvCommand(p, appsClient)
	return command
}

func InjectUnsetEnv(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := apps2.NewUnsetEnvCommand(p, appsClient)
	return command
}

func InjectCreateService(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	serviceInstancesGetter := provideServiceInstancesGetter(kfV1alpha1Interface)
	client := serviceinstances.NewClient(serviceInstancesGetter)
	kubernetesInterface := config.GetKubernetes(p)
	secretsGetter := provideSecretsGetter(kubernetesInterface)
	secretsClient := secrets.NewClient(secretsGetter)
	clientInterface := marketplace.NewClient(kfV1alpha1Interface)
	command := services.NewCreateServiceCommand(p, client, secretsClient, clientInterface)
	return command
}

func InjectCreateUserProvidedService(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	serviceInstancesGetter := provideServiceInstancesGetter(kfV1alpha1Interface)
	client := serviceinstances.NewClient(serviceInstancesGetter)
	kubernetesInterface := config.GetKubernetes(p)
	secretsGetter := provideSecretsGetter(kubernetesInterface)
	secretsClient := secrets.NewClient(secretsGetter)
	command := services.NewCreateUserProvidedServiceCommand(p, client, secretsClient)
	return command
}

func InjectUpdateUserProvidedService(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	serviceInstancesGetter := provideServiceInstancesGetter(kfV1alpha1Interface)
	client := serviceinstances.NewClient(serviceInstancesGetter)
	kubernetesInterface := config.GetKubernetes(p)
	secretsGetter := provideSecretsGetter(kubernetesInterface)
	secretsClient := secrets.NewClient(secretsGetter)
	command := services.NewUpdateUserProvidedServiceCommand(p, client, secretsClient)
	return command
}

func InjectDeleteService(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	serviceInstancesGetter := provideServiceInstancesGetter(kfV1alpha1Interface)
	client := serviceinstances.NewClient(serviceInstancesGetter)
	command := services.NewDeleteServiceCommand(p, client)
	return command
}

func InjectGetService(p *config.KfParams) *cobra.Command {
	command := services.NewGetServiceCommand(p)
	return command
}

func InjectListServices(p *config.KfParams) *cobra.Command {
	command := services.NewListServicesCommand(p)
	return command
}

func InjectMarketplace(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	clientInterface := marketplace.NewClient(kfV1alpha1Interface)
	command := services.NewMarketplaceCommand(p, clientInterface)
	return command
}

func InjectBindService(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	serviceInstanceBindingsGetter := provideServiceInstanceBindingsGetter(kfV1alpha1Interface)
	client := serviceinstancebindings.NewClient(serviceInstanceBindingsGetter)
	kubernetesInterface := config.GetKubernetes(p)
	secretsGetter := provideSecretsGetter(kubernetesInterface)
	secretsClient := secrets.NewClient(secretsGetter)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	buildsClient := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, buildsClient, tailer)
	command := servicebindings.NewBindServiceCommand(p, client, secretsClient, appsClient)
	return command
}

func InjectBindRouteService(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	serviceInstanceBindingsGetter := provideServiceInstanceBindingsGetter(kfV1alpha1Interface)
	client := serviceinstancebindings.NewClient(serviceInstanceBindingsGetter)
	kubernetesInterface := config.GetKubernetes(p)
	secretsGetter := provideSecretsGetter(kubernetesInterface)
	secretsClient := secrets.NewClient(secretsGetter)
	command := servicebindings.NewBindRouteServiceCommand(p, client, secretsClient)
	return command
}

func InjectListBindings(p *config.KfParams) *cobra.Command {
	command := servicebindings.NewListBindingsCommand(p)
	return command
}

func InjectUnbindService(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	serviceInstanceBindingsGetter := provideServiceInstanceBindingsGetter(kfV1alpha1Interface)
	client := serviceinstancebindings.NewClient(serviceInstanceBindingsGetter)
	command := servicebindings.NewUnbindServiceCommand(p, client)
	return command
}

func InjectUnbindRouteService(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	serviceInstanceBindingsGetter := provideServiceInstanceBindingsGetter(kfV1alpha1Interface)
	client := serviceinstancebindings.NewClient(serviceInstanceBindingsGetter)
	command := servicebindings.NewUnbindRouteServiceCommand(p, client)
	return command
}

func InjectVcapServices(p *config.KfParams) *cobra.Command {
	kubernetesInterface := config.GetKubernetes(p)
	command := servicebindings.NewVcapServicesCommand(p, kubernetesInterface)
	return command
}

func InjectCreateServiceBroker(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := cluster.NewClient(kfV1alpha1Interface)
	namespacedClient := namespaced.NewClient(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	secretsGetter := provideSecretsGetter(kubernetesInterface)
	secretsClient := secrets.NewClient(secretsGetter)
	command := servicebrokers.NewCreateServiceBrokerCommand(p, client, namespacedClient, secretsClient)
	return command
}

func InjectDeleteServiceBroker(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := cluster.NewClient(kfV1alpha1Interface)
	namespacedClient := namespaced.NewClient(kfV1alpha1Interface)
	command := servicebrokers.NewDeleteServiceBrokerCommand(p, client, namespacedClient)
	return command
}

func InjectBuildpacksClient(p *config.KfParams) buildpacks.Client {
	remoteImageFetcher := provideRemoteImageFetcher()
	client := buildpacks.NewClient(remoteImageFetcher)
	return client
}

func InjectWrapV2Buildpack(p *config.KfParams) *cobra.Command {
	command := buildpacks2.NewWrapV2BuildpackCommand()
	return command
}

func InjectBuildpacks(p *config.KfParams) *cobra.Command {
	client := InjectBuildpacksClient(p)
	command := buildpacks2.NewBuildpacksCommand(p, client)
	return command
}

func InjectStacks(p *config.KfParams) *cobra.Command {
	command := buildpacks2.NewStacksCommand(p)
	return command
}

func InjectSpaces(p *config.KfParams) *cobra.Command {
	command := spaces.NewListSpacesCommand(p)
	return command
}

func InjectSpace(p *config.KfParams) *cobra.Command {
	command := spaces.NewGetSpaceCommand(p)
	return command
}

func InjectCreateSpace(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces2.NewClient(spacesGetter)
	command := spaces.NewCreateSpaceCommand(p, client)
	return command
}

func InjectDeleteSpace(p *config.KfParams) *cobra.Command {
	command := spaces.NewDeleteSpaceCommand(p)
	return command
}

func InjectConfigSpace(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces2.NewClient(spacesGetter)
	command := spaces.NewConfigSpaceCommand(p, client)
	return command
}

func InjectSetSpaceRole(p *config.KfParams) *cobra.Command {
	kubernetesInterface := config.GetKubernetes(p)
	command := spaces.NewSetSpaceRoleCommand(p, kubernetesInterface)
	return command
}

func InjectSpaceUsers(p *config.KfParams) *cobra.Command {
	kubernetesInterface := config.GetKubernetes(p)
	command := spaces.NewSpaceUsersCommand(p, kubernetesInterface)
	return command
}

func InjectUnsetSpaceRole(p *config.KfParams) *cobra.Command {
	kubernetesInterface := config.GetKubernetes(p)
	command := spaces.NewUnsetSpaceRoleCommand(p, kubernetesInterface)
	return command
}

func InjectDomains(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces2.NewClient(spacesGetter)
	command := spaces.NewDomainsCommand(p, client)
	return command
}

func InjectTarget(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces2.NewClient(spacesGetter)
	command := NewTargetCommand(p, client)
	return command
}

func InjectConfigCluster(p *config.KfParams) *cobra.Command {
	kubernetesInterface := config.GetKubernetes(p)
	configMapsGetter := provideConfigMapsGetter(kubernetesInterface)
	client := configmaps.NewClient(configMapsGetter)
	command := cluster2.NewConfigClusterCommand(p, client)
	return command
}

func InjectRoutes(p *config.KfParams) *cobra.Command {
	command := routes.NewRoutesCommand(p)
	return command
}

func InjectCreateRoute(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := routes2.NewClient(kfV1alpha1Interface)
	command := routes.NewCreateRouteCommand(p, client)
	return command
}

func InjectDeleteRoute(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := routes2.NewClient(kfV1alpha1Interface)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	buildsClient := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, buildsClient, tailer)
	command := routes.NewDeleteRouteCommand(p, client, appsClient)
	return command
}

func InjectDeleteOrphanedRoutes(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := routes2.NewClient(kfV1alpha1Interface)
	command := routes.NewDeleteOrphanedRoutesCommand(p, client)
	return command
}

func InjectMapRoute(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := routes.NewMapRouteCommand(p, appsClient)
	return command
}

func InjectUnmapRoute(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, client, tailer)
	command := routes.NewUnmapRouteCommand(p, appsClient)
	return command
}

func InjectProxyRoute(p *config.KfParams) *cobra.Command {
	command := routes.NewProxyRouteCommand(p)
	return command
}

func InjectBuilds(p *config.KfParams) *cobra.Command {
	command := builds2.NewBuildsCommand(p)
	return command
}

func InjectBuild(p *config.KfParams) *cobra.Command {
	command := builds2.NewGetBuildCommand(p)
	return command
}

func InjectBuildLogs(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	client := builds.NewClient(p, buildsGetter, buildTailer)
	command := builds2.NewBuildLogsCommand(p, client)
	return command
}

func InjectRunTask(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	tasksGetter := provideTasksGetter(kfV1alpha1Interface)
	client := tasks.NewClient(tasksGetter)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	buildsGetter := provideKfBuilds(kfV1alpha1Interface)
	kubernetesInterface := config.GetKubernetes(p)
	buildTailer := builds.TektonLoggingShim(kubernetesInterface)
	buildsClient := builds.NewClient(p, buildsGetter, buildTailer)
	tailer := logs.NewTailer(kubernetesInterface)
	appsClient := apps.NewClient(appsGetter, buildsClient, tailer)
	command := tasks2.NewRunTaskCommand(p, client, appsClient)
	return command
}

func InjectTerminateTask(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	tasksGetter := provideTasksGetter(kfV1alpha1Interface)
	client := tasks.NewClient(tasksGetter)
	command := tasks2.NewTerminateTaskCommand(p, client, kfV1alpha1Interface)
	return command
}

func InjectTasks(p *config.KfParams) *cobra.Command {
	command := tasks2.NewTasksCommand(p)
	return command
}

func InjectCreateJob(p *config.KfParams) *cobra.Command {
	command := taskschedules.NewCreateJobCommand(p)
	return command
}

func InjectRunJob(p *config.KfParams) *cobra.Command {
	command := taskschedules.NewRunJobCommand(p)
	return command
}

func InjectScheduleJob(p *config.KfParams) *cobra.Command {
	command := taskschedules.NewScheduleJobCommand(p)
	return command
}

func InjectListJobs(p *config.KfParams) *cobra.Command {
	command := taskschedules.NewListJobsCommand(p)
	return command
}

func InjectListJobSchedules(p *config.KfParams) *cobra.Command {
	command := taskschedules.NewListJobSchedulesCommand(p)
	return command
}

func InjectJobHistory(p *config.KfParams) *cobra.Command {
	command := taskschedules.NewJobHistoryCommand(p)
	return command
}

func InjectDeleteJob(p *config.KfParams) *cobra.Command {
	command := taskschedules.NewDeleteJobCommand(p)
	return command
}

func InjectDeleteJobSchedule(p *config.KfParams) *cobra.Command {
	command := taskschedules.NewDeleteJobScheduleCommand(p)
	return command
}

func InjectDependencyCommand(p *config.KfParams) *cobra.Command {
	command := dependencies.NewDependencyCommand()
	return command
}

func InjectExportToK8sCommand(p *config.KfParams) *cobra.Command {
	command := exporttok8s.NewExportToK8s(p)
	return command
}

func InjectNetworkPolicies(p *config.KfParams) *cobra.Command {
	command := networkpolicies.NewListCommand(p)
	return command
}

func InjectDeleteNetworkPolicies(p *config.KfParams) *cobra.Command {
	command := networkpolicies.NewDeleteCommand(p)
	return command
}

func InjectDescribeNetworkPolicy(p *config.KfParams) *cobra.Command {
	command := networkpolicies.NewDescribeCommand(p)
	return command
}

// wire_injector.go:

func provideSrcImageBuilder() apps2.SrcImageBuilder {
	return apps2.SrcImageBuilderFunc(apps2.DefaultSrcImageBuilder)
}

var AppsSet = wire.NewSet(
	BuildsSet,
	provideAppsGetter, apps.NewClient, apps.NewPusher, logs.NewTailer,
)

func provideAppsGetter(ki v1alpha1.KfV1alpha1Interface) v1alpha1.AppsGetter {
	return ki
}

func provideSourcePackagesGetter(ki v1alpha1.KfV1alpha1Interface) v1alpha1.SourcePackagesGetter {
	return ki
}

func provideServiceInstancesGetter(ki v1alpha1.KfV1alpha1Interface) v1alpha1.ServiceInstancesGetter {
	return ki
}

func provideSecretsGetter(ki kubernetes.Interface) v1.SecretsGetter {
	return ki.CoreV1()
}

var ServicesSet = wire.NewSet(
	provideSecretsGetter, config.GetKubernetes, config.GetKfClient, provideServiceInstancesGetter, marketplace.NewClient, secrets.NewClient, serviceinstances.NewClient,
)

// /////////////////////
// Service Bindings //
// ///////////////////
func provideServiceInstanceBindingsGetter(ki v1alpha1.KfV1alpha1Interface) v1alpha1.ServiceInstanceBindingsGetter {
	return ki
}

var ServiceBindingsSet = wire.NewSet(
	AppsSet,
	provideSecretsGetter,
	provideServiceInstanceBindingsGetter, secrets.NewClient, serviceinstancebindings.NewClient,
)

var serviceBrokerSet = wire.NewSet(
	provideSecretsGetter, config.GetKubernetes, cluster.NewClient, namespaced.NewClient, config.GetKfClient, secrets.NewClient,
)

// ///////////////
// Buildpacks //
// /////////////
func provideRemoteImageFetcher() buildpacks.RemoteImageFetcher {
	return remote.Image
}

var SpacesSet = wire.NewSet(config.GetKfClient, config.GetKubernetes, provideKfSpaces, spaces2.NewClient)

func provideKfSpaces(ki v1alpha1.KfV1alpha1Interface) v1alpha1.SpacesGetter {
	return ki
}

var ConfigMapsSet = wire.NewSet(config.GetKubernetes, provideConfigMapsGetter, configmaps.NewClient)

func provideConfigMapsGetter(ki kubernetes.Interface) v1.ConfigMapsGetter {
	return ki.CoreV1()
}

var BuildsSet = wire.NewSet(config.GetKfClient, builds.TektonLoggingShim, provideKfBuilds, builds.NewClient, config.GetKubernetes)

func provideKfBuilds(ki v1alpha1.KfV1alpha1Interface) v1alpha1.BuildsGetter {
	return ki
}

func provideTasksGetter(ki v1alpha1.KfV1alpha1Interface) v1alpha1.TasksGetter {
	return ki
}

var TasksSet = wire.NewSet(
	AppsSet,
	provideTasksGetter, tasks.NewClient,
)
