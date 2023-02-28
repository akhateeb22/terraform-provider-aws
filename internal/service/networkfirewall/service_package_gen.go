// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package networkfirewall

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
		"aws_networkfirewall_firewall":        DataSourceFirewall,
		"aws_networkfirewall_firewall_policy": DataSourceFirewallPolicy,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_networkfirewall_firewall":              ResourceFirewall,
		"aws_networkfirewall_firewall_policy":       ResourceFirewallPolicy,
		"aws_networkfirewall_logging_configuration": ResourceLoggingConfiguration,
		"aws_networkfirewall_resource_policy":       ResourceResourcePolicy,
		"aws_networkfirewall_rule_group":            ResourceRuleGroup,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.NetworkFirewall
}

var ServicePackage = &servicePackage{}
