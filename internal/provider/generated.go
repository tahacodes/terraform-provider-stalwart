package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var generatedResources = []resourceDescriptor{
	{
		Name:     "role",
		JMAPType: "x:Role",
		Schema: schema.Schema{
			MarkdownDescription: "Manages a Stalwart role (the `x:Role` object).",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Server-assigned identifier.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the role",
					Required:            true,
				},
				"disabled_permissions": schema.SetAttribute{
					MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
				"enabled_permissions": schema.SetAttribute{
					MarkdownDescription: "List of permissions that are explicitly enabled.",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this role belongs to",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"role_ids": schema.SetAttribute{
					MarkdownDescription: "List of roles this role extends",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "domain",
		JMAPType: "x:Domain",
		Schema: schema.Schema{
			MarkdownDescription: "Manages a Stalwart domain (the `x:Domain` object).",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Server-assigned identifier.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"aliases": schema.SetAttribute{
					MarkdownDescription: "List of additional domain names that are aliases of this domain",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
				"allow_relaying": schema.BoolAttribute{
					MarkdownDescription: "Whether to allow relaying for non-local recipients, useful in split delivery scenarios",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				},
				"catch_all_address": schema.StringAttribute{
					MarkdownDescription: "Catch-all email address that receives messages addressed to unknown local recipients",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"certificate_management": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether TLS certificates for this domain are managed manually or automatically by an ACME provider",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Manual", "Automatic")},
						},
						"acme_provider_id": schema.StringAttribute{
							MarkdownDescription: "Identifier for the ACME provider managing certificates for this domain",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"subject_alternative_names": schema.SetAttribute{
							MarkdownDescription: "Additional hostnames to include in the certificate as Subject Alternative Names (SANs).\nEnter hostnames only (e.g. `mta-sts`, `autoconfig`), the domain is appended automatically.\nTo include the apex domain, enter it in full (e.g. `example.org`).\nLeave empty to request a wildcard certificate when possible, or to use the default SANs.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
					},
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the domain",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the domain",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"directory_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the directory where accounts for this domain are stored, or null to use the internal directory",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"dkim_management": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether DKIM keys for this domain are managed manually or automatically by the server",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Automatic", "Manual")},
						},
						"algorithms": schema.SetAttribute{
							MarkdownDescription: "List of signing algorithms to use when generating new DKIM keys",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
						"delete_after": schema.Int64Attribute{
							MarkdownDescription: "How long to retain old DKIM keys on the server after rotation before deleting them permanently. Requires automatic DNS management.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"retire_after": schema.Int64Attribute{
							MarkdownDescription: "How long to keep the old key's DNS record published after rotation before removing it. Requires automatic DNS management.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"rotate_after": schema.Int64Attribute{
							MarkdownDescription: "How often to rotate DKIM keys. Requires automatic DNS management to be enabled for the domain.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"selector_template": schema.StringAttribute{
							MarkdownDescription: "Template for generating DKIM selectors during key rotation. Supported variables:\n- `{algorithm}`: signing algorithm in lowercase (`rsa`, `ed25519`)\n- `{hash}`: hash algorithm (`sha256`)\n- `{version}`: DKIM version number (`1`)\n- `{date-<fmt>}`: current UTC date formatted with chrono strftime (e.g. `{date-%Y%m%d}`)\n- `{epoch}`: current UTC unix timestamp\n- `{random}`: random 8-character alphanumeric string",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
					},
				},
				"dns_management": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether DNS records for this domain are managed manually or automatically by a DNS provider",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Manual", "Automatic")},
						},
						"dns_server_id": schema.StringAttribute{
							MarkdownDescription: "Identifier for the DNS server provider managing DNS records for this domain",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"origin": schema.StringAttribute{
							MarkdownDescription: "Origin domain used to determine the correct DNS zone for managing records. For example, if the domain is \"sub.example.com\" and DNS records should be managed in the \"example.com\" zone, set the origin to \"example.com\". Leave empty to use the domain name itself as the zone origin.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"publish_records": schema.SetAttribute{
							MarkdownDescription: "Which DNS record types should be automatically published and kept in sync",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
					},
				},
				"dns_zone_file": schema.StringAttribute{
					MarkdownDescription: "Current DNS zone data for the domain",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"is_enabled": schema.BoolAttribute{
					MarkdownDescription: "Whether this domain is enabled",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					Default:             booldefault.StaticBool(true),
				},
				"logo": schema.StringAttribute{
					MarkdownDescription: "URL or base64-encoded image representing the domain",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this domain belongs to",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Domain name",
					Required:            true,
				},
				"report_address_uri": schema.StringAttribute{
					MarkdownDescription: "Email address to receive DMARC, TLS-RPT and CAA reports for this domain, or null to not receive reports",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					Default:             stringdefault.StaticString("mailto:postmaster"),
				},
				"sub_addressing": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether sub-addressing (plus addressing) is enabled for the domain",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Enabled", "Custom", "Disabled")},
						},
					},
				},
			},
		},
	},
	{
		Name:     "mailing_list",
		JMAPType: "x:MailingList",
		Schema: schema.Schema{
			MarkdownDescription: "Manages a Stalwart mailing list (the `x:MailingList` object).",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Server-assigned identifier.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"aliases": schema.ListNestedAttribute{
					MarkdownDescription: "List of email aliases for the mailing list",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.List{listplanmodifier.UseStateForUnknown()},
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Description of the email alias",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"domain_id": schema.StringAttribute{
								MarkdownDescription: "Identifier for the domain of the email alias (the part after the @ symbol).",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Whether this email alias is enabled",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "The local part of the email alias (the part before the @ symbol)",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the mailing list",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this mailing list belongs to. This is used to determine the email address of the mailing list, which is formed as name@domain.",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"email_address": schema.StringAttribute{
					MarkdownDescription: "The email address of the mailing list, formed as name@domain.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this mailing list belongs to",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the mailing list, typically an email address local part.",
					Required:            true,
				},
				"recipients": schema.SetAttribute{
					MarkdownDescription: "List of email addresses that are members of the mailing list",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "tenant",
		JMAPType: "x:Tenant",
		Schema: schema.Schema{
			MarkdownDescription: "Manages a Stalwart tenant (the `x:Tenant` object).",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Server-assigned identifier.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the tenant",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"logo": schema.StringAttribute{
					MarkdownDescription: "URL or base64-encoded image representing the tenant",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the tenant",
					Required:            true,
				},
				"permissions": schema.SingleNestedAttribute{
					MarkdownDescription: "Permissions assigned to this tenant",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Inherit", "Merge", "Replace")},
						},
						"disabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
						"enabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly enabled.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
					},
				},
				"quotas": schema.MapAttribute{
					MarkdownDescription: "Quotas for different object types within this tenant",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
				"roles": schema.SingleNestedAttribute{
					MarkdownDescription: "Roles assigned to this tenant",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Default", "Custom")},
						},
						"role_ids": schema.SetAttribute{
							MarkdownDescription: "List of roles assigned to this principal.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
					},
				},
				"used_disk_quota": schema.Int64Attribute{
					MarkdownDescription: "Amount of disk space currently used by this tenant (bytes)",
					Computed:            true,
					PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				},
			},
		},
	},
	{
		Name:     "user",
		JMAPType: "x:Account",
		Variant:  "User",
		Schema: schema.Schema{
			MarkdownDescription: "Manages a Stalwart user (the `x:Account` object).",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Server-assigned identifier.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"aliases": schema.ListNestedAttribute{
					MarkdownDescription: "List of email aliases for the account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.List{listplanmodifier.UseStateForUnknown()},
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Description of the email alias",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"domain_id": schema.StringAttribute{
								MarkdownDescription: "Identifier for the domain of the email alias (the part after the @ symbol).",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Whether this email alias is enabled",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "The local part of the email alias (the part before the @ symbol)",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the account",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this account belongs to. This is used to determine the email address of the account, which is formed as name@domain.",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"email_address": schema.StringAttribute{
					MarkdownDescription: "Email address for the user account, formed as name@domain.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"encryption_at_rest": schema.SingleNestedAttribute{
					MarkdownDescription: "Encryption-at-rest settings for the account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Disabled", "Aes128", "Aes256", "Aes256Gcm", "ChaCha20Poly1305")},
						},
						"allow_spam_training": schema.BoolAttribute{
							MarkdownDescription: "Whether to allow training the spam classifier with plaintext emails before encryption",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
						},
						"encrypt_on_append": schema.BoolAttribute{
							MarkdownDescription: "Whether to encrypt emails when they are appended to mailboxes",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
						},
						"public_key": schema.StringAttribute{
							MarkdownDescription: "Public key used for encrypting emails",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
					},
				},
				"locale": schema.StringAttribute{
					MarkdownDescription: "Preferred locale for the account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					Validators:          []validator.String{stringvalidator.OneOf("POSIX", "aa_DJ", "aa_ER", "aa_ER@saaho", "aa_ET", "af_ZA", "agr_PE", "ak_GH", "am_ET", "an_ES", "anp_IN", "ar_AE", "ar_BH", "ar_DZ", "ar_EG", "ar_IN", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SD", "ar_SS", "ar_SY", "ar_TN", "ar_YE", "as_IN", "ast_ES", "ayc_PE", "az_AZ", "az_IR", "be_BY", "be_BY@latin", "bem_ZM", "ber_DZ", "ber_MA", "bg_BG", "bhb_IN", "bho_IN", "bho_NP", "bi_VU", "bn_BD", "bn_IN", "bo_CN", "bo_IN", "br_FR", "br_FR@euro", "brx_IN", "bs_BA", "byn_ER", "ca_AD", "ca_ES", "ca_ES@euro", "ca_ES@valencia", "ca_FR", "ca_IT", "ce_RU", "chr_US", "cmn_TW", "crh_UA", "cs_CZ", "csb_PL", "cv_RU", "cy_GB", "da_DK", "de_AT", "de_AT@euro", "de_BE", "de_BE@euro", "de_CH", "de_DE", "de_DE@euro", "de_IT", "de_LI", "de_LU", "de_LU@euro", "doi_IN", "dsb_DE", "dv_MV", "dz_BT", "el_CY", "el_GR", "el_GR@euro", "en_AG", "en_AU", "en_BW", "en_CA", "en_DK", "en_GB", "en_HK", "en_IE", "en_IE@euro", "en_IL", "en_IN", "en_NG", "en_NZ", "en_PH", "en_SC", "en_SG", "en_US", "en_ZA", "en_ZM", "en_ZW", "eo", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_CU", "es_DO", "es_EC", "es_ES", "es_ES@euro", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PE", "es_PR", "es_PY", "es_SV", "es_US", "es_UY", "es_VE", "et_EE", "eu_ES", "eu_ES@euro", "fa_IR", "ff_SN", "fi_FI", "fi_FI@euro", "fil_PH", "fo_FO", "fr_BE", "fr_BE@euro", "fr_CA", "fr_CH", "fr_FR", "fr_FR@euro", "fr_LU", "fr_LU@euro", "fur_IT", "fy_DE", "fy_NL", "ga_IE", "ga_IE@euro", "gd_GB", "gez_ER", "gez_ER@abegede", "gez_ET", "gez_ET@abegede", "gl_ES", "gl_ES@euro", "gu_IN", "gv_GB", "ha_NG", "hak_TW", "he_IL", "hi_IN", "hif_FJ", "hne_IN", "hr_HR", "hsb_DE", "ht_HT", "hu_HU", "hy_AM", "ia_FR", "id_ID", "ig_NG", "ik_CA", "is_IS", "it_CH", "it_IT", "it_IT@euro", "iu_CA", "ja_JP", "ka_GE", "kab_DZ", "kk_KZ", "kl_GL", "km_KH", "kn_IN", "ko_KR", "kok_IN", "ks_IN", "ks_IN@devanagari", "ku_TR", "kw_GB", "ky_KG", "lb_LU", "lg_UG", "li_BE", "li_NL", "lij_IT", "ln_CD", "lo_LA", "lt_LT", "lv_LV", "lzh_TW", "mag_IN", "mai_IN", "mai_NP", "mfe_MU", "mg_MG", "mhr_RU", "mi_NZ", "miq_NI", "mjw_IN", "mk_MK", "ml_IN", "mn_MN", "mni_IN", "mnw_MM", "mr_IN", "ms_MY", "mt_MT", "my_MM", "nan_TW", "nan_TW@latin", "nb_NO", "nds_DE", "nds_NL", "ne_NP", "nhn_MX", "niu_NU", "niu_NZ", "nl_AW", "nl_BE", "nl_BE@euro", "nl_NL", "nl_NL@euro", "nn_NO", "nr_ZA", "nso_ZA", "oc_FR", "om_ET", "om_KE", "or_IN", "os_RU", "pa_IN", "pa_PK", "pap_AW", "pap_CW", "pl_PL", "ps_AF", "pt_BR", "pt_PT", "pt_PT@euro", "quz_PE", "raj_IN", "ro_RO", "ru_RU", "ru_UA", "rw_RW", "sa_IN", "sah_RU", "sat_IN", "sc_IT", "sd_IN", "sd_IN@devanagari", "se_NO", "sgs_LT", "shn_MM", "shs_CA", "si_LK", "sid_ET", "sk_SK", "sl_SI", "sm_WS", "so_DJ", "so_ET", "so_KE", "so_SO", "sq_AL", "sq_MK", "sr_ME", "sr_RS", "sr_RS@latin", "ss_ZA", "st_ZA", "sv_FI", "sv_FI@euro", "sv_SE", "sw_KE", "sw_TZ", "szl_PL", "ta_IN", "ta_LK", "tcy_IN", "te_IN", "tg_TJ", "th_TH", "the_NP", "ti_ER", "ti_ET", "tig_ER", "tk_TM", "tl_PH", "tn_ZA", "to_TO", "tpi_PG", "tr_CY", "tr_TR", "ts_ZA", "tt_RU", "tt_RU@iqtelif", "ug_CN", "uk_UA", "unm_US", "ur_IN", "ur_PK", "uz_UZ", "uz_UZ@cyrillic", "ve_ZA", "vi_VN", "wa_BE", "wa_BE@euro", "wae_CH", "wal_ET", "wo_SN", "xh_ZA", "yi_US", "yo_NG", "yue_HK", "yuw_PG", "zh_CN", "zh_HK", "zh_SG", "zh_TW", "zu_ZA")},
				},
				"member_group_ids": schema.SetAttribute{
					MarkdownDescription: "List of groups that this account is a member of",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this account belongs to",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the account, typically an email address local part.",
					Required:            true,
				},
				"permissions": schema.SingleNestedAttribute{
					MarkdownDescription: "Permissions assigned to this account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Inherit", "Merge", "Replace")},
						},
						"disabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
						"enabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly enabled.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
					},
				},
				"quotas": schema.MapAttribute{
					MarkdownDescription: "Quotas for different object types within this account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
				"roles": schema.SingleNestedAttribute{
					MarkdownDescription: "Roles assigned to this user account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("User", "Admin", "Custom")},
						},
						"role_ids": schema.SetAttribute{
							MarkdownDescription: "List of roles assigned to this principal.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
					},
				},
				"time_zone": schema.StringAttribute{
					MarkdownDescription: "Preferred time zone for the account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					Validators:          []validator.String{stringvalidator.OneOf("Africa/Abidjan", "Africa/Accra", "Africa/Addis_Ababa", "Africa/Algiers", "Africa/Asmara", "Africa/Asmera", "Africa/Bamako", "Africa/Bangui", "Africa/Banjul", "Africa/Bissau", "Africa/Blantyre", "Africa/Brazzaville", "Africa/Bujumbura", "Africa/Cairo", "Africa/Casablanca", "Africa/Ceuta", "Africa/Conakry", "Africa/Dakar", "Africa/Dar_es_Salaam", "Africa/Djibouti", "Africa/Douala", "Africa/El_Aaiun", "Africa/Freetown", "Africa/Gaborone", "Africa/Harare", "Africa/Johannesburg", "Africa/Juba", "Africa/Kampala", "Africa/Khartoum", "Africa/Kigali", "Africa/Kinshasa", "Africa/Lagos", "Africa/Libreville", "Africa/Lome", "Africa/Luanda", "Africa/Lubumbashi", "Africa/Lusaka", "Africa/Malabo", "Africa/Maputo", "Africa/Maseru", "Africa/Mbabane", "Africa/Mogadishu", "Africa/Monrovia", "Africa/Nairobi", "Africa/Ndjamena", "Africa/Niamey", "Africa/Nouakchott", "Africa/Ouagadougou", "Africa/Porto-Novo", "Africa/Sao_Tome", "Africa/Timbuktu", "Africa/Tripoli", "Africa/Tunis", "Africa/Windhoek", "America/Adak", "America/Anchorage", "America/Anguilla", "America/Antigua", "America/Araguaina", "America/Argentina/Buenos_Aires", "America/Argentina/Catamarca", "America/Argentina/ComodRivadavia", "America/Argentina/Cordoba", "America/Argentina/Jujuy", "America/Argentina/La_Rioja", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos", "America/Argentina/Salta", "America/Argentina/San_Juan", "America/Argentina/San_Luis", "America/Argentina/Tucuman", "America/Argentina/Ushuaia", "America/Aruba", "America/Asuncion", "America/Atikokan", "America/Atka", "America/Bahia", "America/Bahia_Banderas", "America/Barbados", "America/Belem", "America/Belize", "America/Blanc-Sablon", "America/Boa_Vista", "America/Bogota", "America/Boise", "America/Buenos_Aires", "America/Cambridge_Bay", "America/Campo_Grande", "America/Cancun", "America/Caracas", "America/Catamarca", "America/Cayenne", "America/Cayman", "America/Chicago", "America/Chihuahua", "America/Ciudad_Juarez", "America/Coral_Harbour", "America/Cordoba", "America/Costa_Rica", "America/Coyhaique", "America/Creston", "America/Cuiaba", "America/Curacao", "America/Danmarkshavn", "America/Dawson", "America/Dawson_Creek", "America/Denver", "America/Detroit", "America/Dominica", "America/Edmonton", "America/Eirunepe", "America/El_Salvador", "America/Ensenada", "America/Fort_Nelson", "America/Fort_Wayne", "America/Fortaleza", "America/Glace_Bay", "America/Godthab", "America/Goose_Bay", "America/Grand_Turk", "America/Grenada", "America/Guadeloupe", "America/Guatemala", "America/Guayaquil", "America/Guyana", "America/Halifax", "America/Havana", "America/Hermosillo", "America/Indiana/Indianapolis", "America/Indiana/Knox", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Tell_City", "America/Indiana/Vevay", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Indianapolis", "America/Inuvik", "America/Iqaluit", "America/Jamaica", "America/Jujuy", "America/Juneau", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Knox_IN", "America/Kralendijk", "America/La_Paz", "America/Lima", "America/Los_Angeles", "America/Louisville", "America/Lower_Princes", "America/Maceio", "America/Managua", "America/Manaus", "America/Marigot", "America/Martinique", "America/Matamoros", "America/Mazatlan", "America/Mendoza", "America/Menominee", "America/Merida", "America/Metlakatla", "America/Mexico_City", "America/Miquelon", "America/Moncton", "America/Monterrey", "America/Montevideo", "America/Montreal", "America/Montserrat", "America/Nassau", "America/New_York", "America/Nipigon", "America/Nome", "America/Noronha", "America/North_Dakota/Beulah", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Nuuk", "America/Ojinaga", "America/Panama", "America/Pangnirtung", "America/Paramaribo", "America/Phoenix", "America/Port-au-Prince", "America/Port_of_Spain", "America/Porto_Acre", "America/Porto_Velho", "America/Puerto_Rico", "America/Punta_Arenas", "America/Rainy_River", "America/Rankin_Inlet", "America/Recife", "America/Regina", "America/Resolute", "America/Rio_Branco", "America/Rosario", "America/Santa_Isabel", "America/Santarem", "America/Santiago", "America/Santo_Domingo", "America/Sao_Paulo", "America/Scoresbysund", "America/Shiprock", "America/Sitka", "America/St_Barthelemy", "America/St_Johns", "America/St_Kitts", "America/St_Lucia", "America/St_Thomas", "America/St_Vincent", "America/Swift_Current", "America/Tegucigalpa", "America/Thule", "America/Thunder_Bay", "America/Tijuana", "America/Toronto", "America/Tortola", "America/Vancouver", "America/Virgin", "America/Whitehorse", "America/Winnipeg", "America/Yakutat", "America/Yellowknife", "Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Macquarie", "Antarctica/Mawson", "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/South_Pole", "Antarctica/Syowa", "Antarctica/Troll", "Antarctica/Vostok", "Arctic/Longyearbyen", "Asia/Aden", "Asia/Almaty", "Asia/Amman", "Asia/Anadyr", "Asia/Aqtau", "Asia/Aqtobe", "Asia/Ashgabat", "Asia/Ashkhabad", "Asia/Atyrau", "Asia/Baghdad", "Asia/Bahrain", "Asia/Baku", "Asia/Bangkok", "Asia/Barnaul", "Asia/Beirut", "Asia/Bishkek", "Asia/Brunei", "Asia/Calcutta", "Asia/Chita", "Asia/Choibalsan", "Asia/Chongqing", "Asia/Chungking", "Asia/Colombo", "Asia/Dacca", "Asia/Damascus", "Asia/Dhaka", "Asia/Dili", "Asia/Dubai", "Asia/Dushanbe", "Asia/Famagusta", "Asia/Gaza", "Asia/Harbin", "Asia/Hebron", "Asia/Ho_Chi_Minh", "Asia/Hong_Kong", "Asia/Hovd", "Asia/Irkutsk", "Asia/Istanbul", "Asia/Jakarta", "Asia/Jayapura", "Asia/Jerusalem", "Asia/Kabul", "Asia/Kamchatka", "Asia/Karachi", "Asia/Kashgar", "Asia/Kathmandu", "Asia/Katmandu", "Asia/Khandyga", "Asia/Kolkata", "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Kuching", "Asia/Kuwait", "Asia/Macao", "Asia/Macau", "Asia/Magadan", "Asia/Makassar", "Asia/Manila", "Asia/Muscat", "Asia/Nicosia", "Asia/Novokuznetsk", "Asia/Novosibirsk", "Asia/Omsk", "Asia/Oral", "Asia/Phnom_Penh", "Asia/Pontianak", "Asia/Pyongyang", "Asia/Qatar", "Asia/Qostanay", "Asia/Qyzylorda", "Asia/Rangoon", "Asia/Riyadh", "Asia/Saigon", "Asia/Sakhalin", "Asia/Samarkand", "Asia/Seoul", "Asia/Shanghai", "Asia/Singapore", "Asia/Srednekolymsk", "Asia/Taipei", "Asia/Tashkent", "Asia/Tbilisi", "Asia/Tehran", "Asia/Tel_Aviv", "Asia/Thimbu", "Asia/Thimphu", "Asia/Tokyo", "Asia/Tomsk", "Asia/Ujung_Pandang", "Asia/Ulaanbaatar", "Asia/Ulan_Bator", "Asia/Urumqi", "Asia/Ust-Nera", "Asia/Vientiane", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yangon", "Asia/Yekaterinburg", "Asia/Yerevan", "Atlantic/Azores", "Atlantic/Bermuda", "Atlantic/Canary", "Atlantic/Cape_Verde", "Atlantic/Faeroe", "Atlantic/Faroe", "Atlantic/Jan_Mayen", "Atlantic/Madeira", "Atlantic/Reykjavik", "Atlantic/South_Georgia", "Atlantic/St_Helena", "Atlantic/Stanley", "Australia/ACT", "Australia/Adelaide", "Australia/Brisbane", "Australia/Broken_Hill", "Australia/Canberra", "Australia/Currie", "Australia/Darwin", "Australia/Eucla", "Australia/Hobart", "Australia/LHI", "Australia/Lindeman", "Australia/Lord_Howe", "Australia/Melbourne", "Australia/NSW", "Australia/North", "Australia/Perth", "Australia/Queensland", "Australia/South", "Australia/Sydney", "Australia/Tasmania", "Australia/Victoria", "Australia/West", "Australia/Yancowinna", "Brazil/Acre", "Brazil/DeNoronha", "Brazil/East", "Brazil/West", "CET", "CST6CDT", "Canada/Atlantic", "Canada/Central", "Canada/Eastern", "Canada/Mountain", "Canada/Newfoundland", "Canada/Pacific", "Canada/Saskatchewan", "Canada/Yukon", "Chile/Continental", "Chile/EasterIsland", "Cuba", "EET", "EST", "EST5EDT", "Egypt", "Eire", "Etc/GMT", "Etc/GMT+0", "Etc/GMT+1", "Etc/GMT+10", "Etc/GMT+11", "Etc/GMT+12", "Etc/GMT+2", "Etc/GMT+3", "Etc/GMT+4", "Etc/GMT+5", "Etc/GMT+6", "Etc/GMT+7", "Etc/GMT+8", "Etc/GMT+9", "Etc/GMT-0", "Etc/GMT-1", "Etc/GMT-10", "Etc/GMT-11", "Etc/GMT-12", "Etc/GMT-13", "Etc/GMT-14", "Etc/GMT-2", "Etc/GMT-3", "Etc/GMT-4", "Etc/GMT-5", "Etc/GMT-6", "Etc/GMT-7", "Etc/GMT-8", "Etc/GMT-9", "Etc/GMT0", "Etc/Greenwich", "Etc/UCT", "Etc/UTC", "Etc/Universal", "Etc/Zulu", "Europe/Amsterdam", "Europe/Andorra", "Europe/Astrakhan", "Europe/Athens", "Europe/Belfast", "Europe/Belgrade", "Europe/Berlin", "Europe/Bratislava", "Europe/Brussels", "Europe/Bucharest", "Europe/Budapest", "Europe/Busingen", "Europe/Chisinau", "Europe/Copenhagen", "Europe/Dublin", "Europe/Gibraltar", "Europe/Guernsey", "Europe/Helsinki", "Europe/Isle_of_Man", "Europe/Istanbul", "Europe/Jersey", "Europe/Kaliningrad", "Europe/Kiev", "Europe/Kirov", "Europe/Kyiv", "Europe/Lisbon", "Europe/Ljubljana", "Europe/London", "Europe/Luxembourg", "Europe/Madrid", "Europe/Malta", "Europe/Mariehamn", "Europe/Minsk", "Europe/Monaco", "Europe/Moscow", "Europe/Nicosia", "Europe/Oslo", "Europe/Paris", "Europe/Podgorica", "Europe/Prague", "Europe/Riga", "Europe/Rome", "Europe/Samara", "Europe/San_Marino", "Europe/Sarajevo", "Europe/Saratov", "Europe/Simferopol", "Europe/Skopje", "Europe/Sofia", "Europe/Stockholm", "Europe/Tallinn", "Europe/Tirane", "Europe/Tiraspol", "Europe/Ulyanovsk", "Europe/Uzhgorod", "Europe/Vaduz", "Europe/Vatican", "Europe/Vienna", "Europe/Vilnius", "Europe/Volgograd", "Europe/Warsaw", "Europe/Zagreb", "Europe/Zaporozhye", "Europe/Zurich", "Factory", "GB", "GB-Eire", "GMT", "GMT+0", "GMT-0", "GMT0", "Greenwich", "HST", "Hongkong", "Iceland", "Indian/Antananarivo", "Indian/Chagos", "Indian/Christmas", "Indian/Cocos", "Indian/Comoro", "Indian/Kerguelen", "Indian/Mahe", "Indian/Maldives", "Indian/Mauritius", "Indian/Mayotte", "Indian/Reunion", "Iran", "Israel", "Jamaica", "Japan", "Kwajalein", "Libya", "MET", "MST", "MST7MDT", "Mexico/BajaNorte", "Mexico/BajaSur", "Mexico/General", "NZ", "NZ-CHAT", "Navajo", "PRC", "PST8PDT", "Pacific/Apia", "Pacific/Auckland", "Pacific/Bougainville", "Pacific/Chatham", "Pacific/Chuuk", "Pacific/Easter", "Pacific/Efate", "Pacific/Enderbury", "Pacific/Fakaofo", "Pacific/Fiji", "Pacific/Funafuti", "Pacific/Galapagos", "Pacific/Gambier", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Honolulu", "Pacific/Johnston", "Pacific/Kanton", "Pacific/Kiritimati", "Pacific/Kosrae", "Pacific/Kwajalein", "Pacific/Majuro", "Pacific/Marquesas", "Pacific/Midway", "Pacific/Nauru", "Pacific/Niue", "Pacific/Norfolk", "Pacific/Noumea", "Pacific/Pago_Pago", "Pacific/Palau", "Pacific/Pitcairn", "Pacific/Pohnpei", "Pacific/Ponape", "Pacific/Port_Moresby", "Pacific/Rarotonga", "Pacific/Saipan", "Pacific/Samoa", "Pacific/Tahiti", "Pacific/Tarawa", "Pacific/Tongatapu", "Pacific/Truk", "Pacific/Wake", "Pacific/Wallis", "Pacific/Yap", "Poland", "Portugal", "ROC", "ROK", "Singapore", "Turkey", "UCT", "US/Alaska", "US/Aleutian", "US/Arizona", "US/Central", "US/East-Indiana", "US/Eastern", "US/Hawaii", "US/Indiana-Starke", "US/Michigan", "US/Mountain", "US/Pacific", "US/Samoa", "UTC", "Universal", "W-SU", "WET", "Zulu")},
				},
				"used_disk_quota": schema.Int64Attribute{
					MarkdownDescription: "Amount of disk space currently used by this account (bytes)",
					Computed:            true,
					PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				},
			},
		},
	},
	{
		Name:     "group",
		JMAPType: "x:Account",
		Variant:  "Group",
		Schema: schema.Schema{
			MarkdownDescription: "Manages a Stalwart group (the `x:Account` object).",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Server-assigned identifier.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"aliases": schema.ListNestedAttribute{
					MarkdownDescription: "List of email aliases for the group",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.List{listplanmodifier.UseStateForUnknown()},
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Description of the email alias",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"domain_id": schema.StringAttribute{
								MarkdownDescription: "Identifier for the domain of the email alias (the part after the @ symbol).",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Whether this email alias is enabled",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "The local part of the email alias (the part before the @ symbol)",
								Optional:            true,
								Computed:            true,
								PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the account",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the group",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this group belongs to. This is used to determine the email address of the group, which is formed as name@domain.",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"email_address": schema.StringAttribute{
					MarkdownDescription: "Email address of the group, formed as name@domain.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"locale": schema.StringAttribute{
					MarkdownDescription: "Preferred locale for the group",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					Validators:          []validator.String{stringvalidator.OneOf("POSIX", "aa_DJ", "aa_ER", "aa_ER@saaho", "aa_ET", "af_ZA", "agr_PE", "ak_GH", "am_ET", "an_ES", "anp_IN", "ar_AE", "ar_BH", "ar_DZ", "ar_EG", "ar_IN", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SD", "ar_SS", "ar_SY", "ar_TN", "ar_YE", "as_IN", "ast_ES", "ayc_PE", "az_AZ", "az_IR", "be_BY", "be_BY@latin", "bem_ZM", "ber_DZ", "ber_MA", "bg_BG", "bhb_IN", "bho_IN", "bho_NP", "bi_VU", "bn_BD", "bn_IN", "bo_CN", "bo_IN", "br_FR", "br_FR@euro", "brx_IN", "bs_BA", "byn_ER", "ca_AD", "ca_ES", "ca_ES@euro", "ca_ES@valencia", "ca_FR", "ca_IT", "ce_RU", "chr_US", "cmn_TW", "crh_UA", "cs_CZ", "csb_PL", "cv_RU", "cy_GB", "da_DK", "de_AT", "de_AT@euro", "de_BE", "de_BE@euro", "de_CH", "de_DE", "de_DE@euro", "de_IT", "de_LI", "de_LU", "de_LU@euro", "doi_IN", "dsb_DE", "dv_MV", "dz_BT", "el_CY", "el_GR", "el_GR@euro", "en_AG", "en_AU", "en_BW", "en_CA", "en_DK", "en_GB", "en_HK", "en_IE", "en_IE@euro", "en_IL", "en_IN", "en_NG", "en_NZ", "en_PH", "en_SC", "en_SG", "en_US", "en_ZA", "en_ZM", "en_ZW", "eo", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_CU", "es_DO", "es_EC", "es_ES", "es_ES@euro", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PE", "es_PR", "es_PY", "es_SV", "es_US", "es_UY", "es_VE", "et_EE", "eu_ES", "eu_ES@euro", "fa_IR", "ff_SN", "fi_FI", "fi_FI@euro", "fil_PH", "fo_FO", "fr_BE", "fr_BE@euro", "fr_CA", "fr_CH", "fr_FR", "fr_FR@euro", "fr_LU", "fr_LU@euro", "fur_IT", "fy_DE", "fy_NL", "ga_IE", "ga_IE@euro", "gd_GB", "gez_ER", "gez_ER@abegede", "gez_ET", "gez_ET@abegede", "gl_ES", "gl_ES@euro", "gu_IN", "gv_GB", "ha_NG", "hak_TW", "he_IL", "hi_IN", "hif_FJ", "hne_IN", "hr_HR", "hsb_DE", "ht_HT", "hu_HU", "hy_AM", "ia_FR", "id_ID", "ig_NG", "ik_CA", "is_IS", "it_CH", "it_IT", "it_IT@euro", "iu_CA", "ja_JP", "ka_GE", "kab_DZ", "kk_KZ", "kl_GL", "km_KH", "kn_IN", "ko_KR", "kok_IN", "ks_IN", "ks_IN@devanagari", "ku_TR", "kw_GB", "ky_KG", "lb_LU", "lg_UG", "li_BE", "li_NL", "lij_IT", "ln_CD", "lo_LA", "lt_LT", "lv_LV", "lzh_TW", "mag_IN", "mai_IN", "mai_NP", "mfe_MU", "mg_MG", "mhr_RU", "mi_NZ", "miq_NI", "mjw_IN", "mk_MK", "ml_IN", "mn_MN", "mni_IN", "mnw_MM", "mr_IN", "ms_MY", "mt_MT", "my_MM", "nan_TW", "nan_TW@latin", "nb_NO", "nds_DE", "nds_NL", "ne_NP", "nhn_MX", "niu_NU", "niu_NZ", "nl_AW", "nl_BE", "nl_BE@euro", "nl_NL", "nl_NL@euro", "nn_NO", "nr_ZA", "nso_ZA", "oc_FR", "om_ET", "om_KE", "or_IN", "os_RU", "pa_IN", "pa_PK", "pap_AW", "pap_CW", "pl_PL", "ps_AF", "pt_BR", "pt_PT", "pt_PT@euro", "quz_PE", "raj_IN", "ro_RO", "ru_RU", "ru_UA", "rw_RW", "sa_IN", "sah_RU", "sat_IN", "sc_IT", "sd_IN", "sd_IN@devanagari", "se_NO", "sgs_LT", "shn_MM", "shs_CA", "si_LK", "sid_ET", "sk_SK", "sl_SI", "sm_WS", "so_DJ", "so_ET", "so_KE", "so_SO", "sq_AL", "sq_MK", "sr_ME", "sr_RS", "sr_RS@latin", "ss_ZA", "st_ZA", "sv_FI", "sv_FI@euro", "sv_SE", "sw_KE", "sw_TZ", "szl_PL", "ta_IN", "ta_LK", "tcy_IN", "te_IN", "tg_TJ", "th_TH", "the_NP", "ti_ER", "ti_ET", "tig_ER", "tk_TM", "tl_PH", "tn_ZA", "to_TO", "tpi_PG", "tr_CY", "tr_TR", "ts_ZA", "tt_RU", "tt_RU@iqtelif", "ug_CN", "uk_UA", "unm_US", "ur_IN", "ur_PK", "uz_UZ", "uz_UZ@cyrillic", "ve_ZA", "vi_VN", "wa_BE", "wa_BE@euro", "wae_CH", "wal_ET", "wo_SN", "xh_ZA", "yi_US", "yo_NG", "yue_HK", "yuw_PG", "zh_CN", "zh_HK", "zh_SG", "zh_TW", "zu_ZA")},
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this group belongs to",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the group, typically an email address local part.",
					Required:            true,
				},
				"permissions": schema.SingleNestedAttribute{
					MarkdownDescription: "Permissions assigned to this group",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Inherit", "Merge", "Replace")},
						},
						"disabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
						"enabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly enabled.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
					},
				},
				"quotas": schema.MapAttribute{
					MarkdownDescription: "Quotas for different object types within this group",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
					ElementType:         types.StringType,
				},
				"roles": schema.SingleNestedAttribute{
					MarkdownDescription: "Roles assigned to this group",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Default", "Custom")},
						},
						"role_ids": schema.SetAttribute{
							MarkdownDescription: "List of roles assigned to this principal.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
							ElementType:         types.StringType,
						},
					},
				},
				"time_zone": schema.StringAttribute{
					MarkdownDescription: "Preferred time zone for the account",
					Optional:            true,
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					Validators:          []validator.String{stringvalidator.OneOf("Africa/Abidjan", "Africa/Accra", "Africa/Addis_Ababa", "Africa/Algiers", "Africa/Asmara", "Africa/Asmera", "Africa/Bamako", "Africa/Bangui", "Africa/Banjul", "Africa/Bissau", "Africa/Blantyre", "Africa/Brazzaville", "Africa/Bujumbura", "Africa/Cairo", "Africa/Casablanca", "Africa/Ceuta", "Africa/Conakry", "Africa/Dakar", "Africa/Dar_es_Salaam", "Africa/Djibouti", "Africa/Douala", "Africa/El_Aaiun", "Africa/Freetown", "Africa/Gaborone", "Africa/Harare", "Africa/Johannesburg", "Africa/Juba", "Africa/Kampala", "Africa/Khartoum", "Africa/Kigali", "Africa/Kinshasa", "Africa/Lagos", "Africa/Libreville", "Africa/Lome", "Africa/Luanda", "Africa/Lubumbashi", "Africa/Lusaka", "Africa/Malabo", "Africa/Maputo", "Africa/Maseru", "Africa/Mbabane", "Africa/Mogadishu", "Africa/Monrovia", "Africa/Nairobi", "Africa/Ndjamena", "Africa/Niamey", "Africa/Nouakchott", "Africa/Ouagadougou", "Africa/Porto-Novo", "Africa/Sao_Tome", "Africa/Timbuktu", "Africa/Tripoli", "Africa/Tunis", "Africa/Windhoek", "America/Adak", "America/Anchorage", "America/Anguilla", "America/Antigua", "America/Araguaina", "America/Argentina/Buenos_Aires", "America/Argentina/Catamarca", "America/Argentina/ComodRivadavia", "America/Argentina/Cordoba", "America/Argentina/Jujuy", "America/Argentina/La_Rioja", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos", "America/Argentina/Salta", "America/Argentina/San_Juan", "America/Argentina/San_Luis", "America/Argentina/Tucuman", "America/Argentina/Ushuaia", "America/Aruba", "America/Asuncion", "America/Atikokan", "America/Atka", "America/Bahia", "America/Bahia_Banderas", "America/Barbados", "America/Belem", "America/Belize", "America/Blanc-Sablon", "America/Boa_Vista", "America/Bogota", "America/Boise", "America/Buenos_Aires", "America/Cambridge_Bay", "America/Campo_Grande", "America/Cancun", "America/Caracas", "America/Catamarca", "America/Cayenne", "America/Cayman", "America/Chicago", "America/Chihuahua", "America/Ciudad_Juarez", "America/Coral_Harbour", "America/Cordoba", "America/Costa_Rica", "America/Coyhaique", "America/Creston", "America/Cuiaba", "America/Curacao", "America/Danmarkshavn", "America/Dawson", "America/Dawson_Creek", "America/Denver", "America/Detroit", "America/Dominica", "America/Edmonton", "America/Eirunepe", "America/El_Salvador", "America/Ensenada", "America/Fort_Nelson", "America/Fort_Wayne", "America/Fortaleza", "America/Glace_Bay", "America/Godthab", "America/Goose_Bay", "America/Grand_Turk", "America/Grenada", "America/Guadeloupe", "America/Guatemala", "America/Guayaquil", "America/Guyana", "America/Halifax", "America/Havana", "America/Hermosillo", "America/Indiana/Indianapolis", "America/Indiana/Knox", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Tell_City", "America/Indiana/Vevay", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Indianapolis", "America/Inuvik", "America/Iqaluit", "America/Jamaica", "America/Jujuy", "America/Juneau", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Knox_IN", "America/Kralendijk", "America/La_Paz", "America/Lima", "America/Los_Angeles", "America/Louisville", "America/Lower_Princes", "America/Maceio", "America/Managua", "America/Manaus", "America/Marigot", "America/Martinique", "America/Matamoros", "America/Mazatlan", "America/Mendoza", "America/Menominee", "America/Merida", "America/Metlakatla", "America/Mexico_City", "America/Miquelon", "America/Moncton", "America/Monterrey", "America/Montevideo", "America/Montreal", "America/Montserrat", "America/Nassau", "America/New_York", "America/Nipigon", "America/Nome", "America/Noronha", "America/North_Dakota/Beulah", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Nuuk", "America/Ojinaga", "America/Panama", "America/Pangnirtung", "America/Paramaribo", "America/Phoenix", "America/Port-au-Prince", "America/Port_of_Spain", "America/Porto_Acre", "America/Porto_Velho", "America/Puerto_Rico", "America/Punta_Arenas", "America/Rainy_River", "America/Rankin_Inlet", "America/Recife", "America/Regina", "America/Resolute", "America/Rio_Branco", "America/Rosario", "America/Santa_Isabel", "America/Santarem", "America/Santiago", "America/Santo_Domingo", "America/Sao_Paulo", "America/Scoresbysund", "America/Shiprock", "America/Sitka", "America/St_Barthelemy", "America/St_Johns", "America/St_Kitts", "America/St_Lucia", "America/St_Thomas", "America/St_Vincent", "America/Swift_Current", "America/Tegucigalpa", "America/Thule", "America/Thunder_Bay", "America/Tijuana", "America/Toronto", "America/Tortola", "America/Vancouver", "America/Virgin", "America/Whitehorse", "America/Winnipeg", "America/Yakutat", "America/Yellowknife", "Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Macquarie", "Antarctica/Mawson", "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/South_Pole", "Antarctica/Syowa", "Antarctica/Troll", "Antarctica/Vostok", "Arctic/Longyearbyen", "Asia/Aden", "Asia/Almaty", "Asia/Amman", "Asia/Anadyr", "Asia/Aqtau", "Asia/Aqtobe", "Asia/Ashgabat", "Asia/Ashkhabad", "Asia/Atyrau", "Asia/Baghdad", "Asia/Bahrain", "Asia/Baku", "Asia/Bangkok", "Asia/Barnaul", "Asia/Beirut", "Asia/Bishkek", "Asia/Brunei", "Asia/Calcutta", "Asia/Chita", "Asia/Choibalsan", "Asia/Chongqing", "Asia/Chungking", "Asia/Colombo", "Asia/Dacca", "Asia/Damascus", "Asia/Dhaka", "Asia/Dili", "Asia/Dubai", "Asia/Dushanbe", "Asia/Famagusta", "Asia/Gaza", "Asia/Harbin", "Asia/Hebron", "Asia/Ho_Chi_Minh", "Asia/Hong_Kong", "Asia/Hovd", "Asia/Irkutsk", "Asia/Istanbul", "Asia/Jakarta", "Asia/Jayapura", "Asia/Jerusalem", "Asia/Kabul", "Asia/Kamchatka", "Asia/Karachi", "Asia/Kashgar", "Asia/Kathmandu", "Asia/Katmandu", "Asia/Khandyga", "Asia/Kolkata", "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Kuching", "Asia/Kuwait", "Asia/Macao", "Asia/Macau", "Asia/Magadan", "Asia/Makassar", "Asia/Manila", "Asia/Muscat", "Asia/Nicosia", "Asia/Novokuznetsk", "Asia/Novosibirsk", "Asia/Omsk", "Asia/Oral", "Asia/Phnom_Penh", "Asia/Pontianak", "Asia/Pyongyang", "Asia/Qatar", "Asia/Qostanay", "Asia/Qyzylorda", "Asia/Rangoon", "Asia/Riyadh", "Asia/Saigon", "Asia/Sakhalin", "Asia/Samarkand", "Asia/Seoul", "Asia/Shanghai", "Asia/Singapore", "Asia/Srednekolymsk", "Asia/Taipei", "Asia/Tashkent", "Asia/Tbilisi", "Asia/Tehran", "Asia/Tel_Aviv", "Asia/Thimbu", "Asia/Thimphu", "Asia/Tokyo", "Asia/Tomsk", "Asia/Ujung_Pandang", "Asia/Ulaanbaatar", "Asia/Ulan_Bator", "Asia/Urumqi", "Asia/Ust-Nera", "Asia/Vientiane", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yangon", "Asia/Yekaterinburg", "Asia/Yerevan", "Atlantic/Azores", "Atlantic/Bermuda", "Atlantic/Canary", "Atlantic/Cape_Verde", "Atlantic/Faeroe", "Atlantic/Faroe", "Atlantic/Jan_Mayen", "Atlantic/Madeira", "Atlantic/Reykjavik", "Atlantic/South_Georgia", "Atlantic/St_Helena", "Atlantic/Stanley", "Australia/ACT", "Australia/Adelaide", "Australia/Brisbane", "Australia/Broken_Hill", "Australia/Canberra", "Australia/Currie", "Australia/Darwin", "Australia/Eucla", "Australia/Hobart", "Australia/LHI", "Australia/Lindeman", "Australia/Lord_Howe", "Australia/Melbourne", "Australia/NSW", "Australia/North", "Australia/Perth", "Australia/Queensland", "Australia/South", "Australia/Sydney", "Australia/Tasmania", "Australia/Victoria", "Australia/West", "Australia/Yancowinna", "Brazil/Acre", "Brazil/DeNoronha", "Brazil/East", "Brazil/West", "CET", "CST6CDT", "Canada/Atlantic", "Canada/Central", "Canada/Eastern", "Canada/Mountain", "Canada/Newfoundland", "Canada/Pacific", "Canada/Saskatchewan", "Canada/Yukon", "Chile/Continental", "Chile/EasterIsland", "Cuba", "EET", "EST", "EST5EDT", "Egypt", "Eire", "Etc/GMT", "Etc/GMT+0", "Etc/GMT+1", "Etc/GMT+10", "Etc/GMT+11", "Etc/GMT+12", "Etc/GMT+2", "Etc/GMT+3", "Etc/GMT+4", "Etc/GMT+5", "Etc/GMT+6", "Etc/GMT+7", "Etc/GMT+8", "Etc/GMT+9", "Etc/GMT-0", "Etc/GMT-1", "Etc/GMT-10", "Etc/GMT-11", "Etc/GMT-12", "Etc/GMT-13", "Etc/GMT-14", "Etc/GMT-2", "Etc/GMT-3", "Etc/GMT-4", "Etc/GMT-5", "Etc/GMT-6", "Etc/GMT-7", "Etc/GMT-8", "Etc/GMT-9", "Etc/GMT0", "Etc/Greenwich", "Etc/UCT", "Etc/UTC", "Etc/Universal", "Etc/Zulu", "Europe/Amsterdam", "Europe/Andorra", "Europe/Astrakhan", "Europe/Athens", "Europe/Belfast", "Europe/Belgrade", "Europe/Berlin", "Europe/Bratislava", "Europe/Brussels", "Europe/Bucharest", "Europe/Budapest", "Europe/Busingen", "Europe/Chisinau", "Europe/Copenhagen", "Europe/Dublin", "Europe/Gibraltar", "Europe/Guernsey", "Europe/Helsinki", "Europe/Isle_of_Man", "Europe/Istanbul", "Europe/Jersey", "Europe/Kaliningrad", "Europe/Kiev", "Europe/Kirov", "Europe/Kyiv", "Europe/Lisbon", "Europe/Ljubljana", "Europe/London", "Europe/Luxembourg", "Europe/Madrid", "Europe/Malta", "Europe/Mariehamn", "Europe/Minsk", "Europe/Monaco", "Europe/Moscow", "Europe/Nicosia", "Europe/Oslo", "Europe/Paris", "Europe/Podgorica", "Europe/Prague", "Europe/Riga", "Europe/Rome", "Europe/Samara", "Europe/San_Marino", "Europe/Sarajevo", "Europe/Saratov", "Europe/Simferopol", "Europe/Skopje", "Europe/Sofia", "Europe/Stockholm", "Europe/Tallinn", "Europe/Tirane", "Europe/Tiraspol", "Europe/Ulyanovsk", "Europe/Uzhgorod", "Europe/Vaduz", "Europe/Vatican", "Europe/Vienna", "Europe/Vilnius", "Europe/Volgograd", "Europe/Warsaw", "Europe/Zagreb", "Europe/Zaporozhye", "Europe/Zurich", "Factory", "GB", "GB-Eire", "GMT", "GMT+0", "GMT-0", "GMT0", "Greenwich", "HST", "Hongkong", "Iceland", "Indian/Antananarivo", "Indian/Chagos", "Indian/Christmas", "Indian/Cocos", "Indian/Comoro", "Indian/Kerguelen", "Indian/Mahe", "Indian/Maldives", "Indian/Mauritius", "Indian/Mayotte", "Indian/Reunion", "Iran", "Israel", "Jamaica", "Japan", "Kwajalein", "Libya", "MET", "MST", "MST7MDT", "Mexico/BajaNorte", "Mexico/BajaSur", "Mexico/General", "NZ", "NZ-CHAT", "Navajo", "PRC", "PST8PDT", "Pacific/Apia", "Pacific/Auckland", "Pacific/Bougainville", "Pacific/Chatham", "Pacific/Chuuk", "Pacific/Easter", "Pacific/Efate", "Pacific/Enderbury", "Pacific/Fakaofo", "Pacific/Fiji", "Pacific/Funafuti", "Pacific/Galapagos", "Pacific/Gambier", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Honolulu", "Pacific/Johnston", "Pacific/Kanton", "Pacific/Kiritimati", "Pacific/Kosrae", "Pacific/Kwajalein", "Pacific/Majuro", "Pacific/Marquesas", "Pacific/Midway", "Pacific/Nauru", "Pacific/Niue", "Pacific/Norfolk", "Pacific/Noumea", "Pacific/Pago_Pago", "Pacific/Palau", "Pacific/Pitcairn", "Pacific/Pohnpei", "Pacific/Ponape", "Pacific/Port_Moresby", "Pacific/Rarotonga", "Pacific/Saipan", "Pacific/Samoa", "Pacific/Tahiti", "Pacific/Tarawa", "Pacific/Tongatapu", "Pacific/Truk", "Pacific/Wake", "Pacific/Wallis", "Pacific/Yap", "Poland", "Portugal", "ROC", "ROK", "Singapore", "Turkey", "UCT", "US/Alaska", "US/Aleutian", "US/Arizona", "US/Central", "US/East-Indiana", "US/Eastern", "US/Hawaii", "US/Indiana-Starke", "US/Michigan", "US/Mountain", "US/Pacific", "US/Samoa", "UTC", "Universal", "W-SU", "WET", "Zulu")},
				},
				"used_disk_quota": schema.Int64Attribute{
					MarkdownDescription: "Amount of disk space currently used by this account (bytes)",
					Computed:            true,
					PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				},
			},
		},
	},
}
