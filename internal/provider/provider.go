package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tahacodes/terraform-provider-stalwart/internal/client"
)

var _ provider.Provider = (*stalwartProvider)(nil)

type stalwartProvider struct {
	version string
}

type providerModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Token    types.String `tfsdk:"token"`
	Insecure types.Bool   `tfsdk:"insecure"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &stalwartProvider{version: version}
	}
}

func (p *stalwartProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "stalwart"
	resp.Version = p.version
}

func (p *stalwartProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages configuration of a Stalwart mail and collaboration server through its JMAP management API.",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Base URL of the Stalwart server, for example `https://mail.example.com`. Falls back to the `STALWART_URL` environment variable.",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for basic authentication. Falls back to the `STALWART_USER` environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for basic authentication. Falls back to the `STALWART_PASSWORD` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "Bearer token, mutually exclusive with `username` and `password`. Falls back to the `STALWART_TOKEN` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Skip TLS certificate verification. Defaults to `false`.",
				Optional:            true,
			},
		},
	}
}

func (p *stalwartProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config providerModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := stringOrEnv(config.Endpoint, "STALWART_URL")
	username := stringOrEnv(config.Username, "STALWART_USER")
	password := stringOrEnv(config.Password, "STALWART_PASSWORD")
	token := stringOrEnv(config.Token, "STALWART_TOKEN")

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Missing Stalwart endpoint",
			"Set the provider `endpoint` attribute or the STALWART_URL environment variable.",
		)
	}

	if token == "" && (username == "" || password == "") {
		resp.Diagnostics.AddError(
			"Missing Stalwart credentials",
			"Set either `token` (STALWART_TOKEN) or both `username` and `password` (STALWART_USER, STALWART_PASSWORD).",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	c, err := client.New(ctx, client.Config{
		Endpoint: endpoint,
		Username: username,
		Password: password,
		Token:    token,
		Insecure: config.Insecure.ValueBool(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to Stalwart", err.Error())
		return
	}

	resp.ResourceData = c
	resp.DataSourceData = c
}

func (p *stalwartProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewHTTPResource,
	}
}

func (p *stalwartProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func stringOrEnv(v types.String, key string) string {
	if !v.IsNull() && v.ValueString() != "" {
		return v.ValueString()
	}

	return os.Getenv(key)
}
