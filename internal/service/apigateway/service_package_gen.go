// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_api_gateway_api_key":     DataSourceAPIKey,
		"aws_api_gateway_domain_name": DataSourceDomainName,
		"aws_api_gateway_export":      DataSourceExport,
		"aws_api_gateway_resource":    DataSourceResource,
		"aws_api_gateway_rest_api":    DataSourceRestAPI,
		"aws_api_gateway_sdk":         DataSourceSdk,
		"aws_api_gateway_vpc_link":    DataSourceVPCLink,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_api_gateway_account":               ResourceAccount,
		"aws_api_gateway_api_key":               ResourceAPIKey,
		"aws_api_gateway_authorizer":            ResourceAuthorizer,
		"aws_api_gateway_base_path_mapping":     ResourceBasePathMapping,
		"aws_api_gateway_client_certificate":    ResourceClientCertificate,
		"aws_api_gateway_deployment":            ResourceDeployment,
		"aws_api_gateway_documentation_part":    ResourceDocumentationPart,
		"aws_api_gateway_documentation_version": ResourceDocumentationVersion,
		"aws_api_gateway_domain_name":           ResourceDomainName,
		"aws_api_gateway_gateway_response":      ResourceGatewayResponse,
		"aws_api_gateway_integration":           ResourceIntegration,
		"aws_api_gateway_integration_response":  ResourceIntegrationResponse,
		"aws_api_gateway_method":                ResourceMethod,
		"aws_api_gateway_method_response":       ResourceMethodResponse,
		"aws_api_gateway_method_settings":       ResourceMethodSettings,
		"aws_api_gateway_model":                 ResourceModel,
		"aws_api_gateway_request_validator":     ResourceRequestValidator,
		"aws_api_gateway_resource":              ResourceResource,
		"aws_api_gateway_rest_api":              ResourceRestAPI,
		"aws_api_gateway_rest_api_policy":       ResourceRestAPIPolicy,
		"aws_api_gateway_stage":                 ResourceStage,
		"aws_api_gateway_usage_plan":            ResourceUsagePlan,
		"aws_api_gateway_usage_plan_key":        ResourceUsagePlanKey,
		"aws_api_gateway_vpc_link":              ResourceVPCLink,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.APIGateway
}

var ServicePackage = &servicePackage{}
