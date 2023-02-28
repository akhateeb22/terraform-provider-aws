// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package pinpoint

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
	return map[string]func() *schema.Resource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_pinpoint_adm_channel":               ResourceADMChannel,
		"aws_pinpoint_apns_channel":              ResourceAPNSChannel,
		"aws_pinpoint_apns_sandbox_channel":      ResourceAPNSSandboxChannel,
		"aws_pinpoint_apns_voip_channel":         ResourceAPNSVoIPChannel,
		"aws_pinpoint_apns_voip_sandbox_channel": ResourceAPNSVoIPSandboxChannel,
		"aws_pinpoint_app":                       ResourceApp,
		"aws_pinpoint_baidu_channel":             ResourceBaiduChannel,
		"aws_pinpoint_email_channel":             ResourceEmailChannel,
		"aws_pinpoint_event_stream":              ResourceEventStream,
		"aws_pinpoint_gcm_channel":               ResourceGCMChannel,
		"aws_pinpoint_sms_channel":               ResourceSMSChannel,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Pinpoint
}

var ServicePackage = &servicePackage{}
