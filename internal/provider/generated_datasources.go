package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var generatedDataSources = []dataSourceDescriptor{
	{
		Name:      "account_settings",
		JMAPType:  "x:AccountSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures default account settings for locale and encryption. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the account",
					Computed:            true,
				},
				"encryption_at_rest": schema.SingleNestedAttribute{
					MarkdownDescription: "Encryption-at-rest settings for the account",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"allow_spam_training": schema.BoolAttribute{
							MarkdownDescription: "Whether to allow training the spam classifier with plaintext emails before encryption",
							Computed:            true,
						},
						"encrypt_on_append": schema.BoolAttribute{
							MarkdownDescription: "Whether to encrypt emails when they are appended to mailboxes",
							Computed:            true,
						},
						"public_key": schema.StringAttribute{
							MarkdownDescription: "Public key used for encrypting emails",
							Computed:            true,
						},
					},
				},
				"locale": schema.StringAttribute{
					MarkdownDescription: "Preferred locale for the account",
					Computed:            true,
				},
				"time_zone": schema.StringAttribute{
					MarkdownDescription: "Preferred time zone for the account",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "accounts",
		JMAPType: "x:Account",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Account` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "acme_provider",
		JMAPType: "x:AcmeProvider",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an ACME provider for automatic TLS certificate management.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"account_key": schema.StringAttribute{
					MarkdownDescription: "The account key used to authenticate with the ACME provider.",
					Computed:            true,
					Sensitive:           true,
				},
				"account_uri": schema.StringAttribute{
					MarkdownDescription: "The account URI returned by the ACME server after registration. Used for CAA record accounturi binding.",
					Computed:            true,
				},
				"challenge_type": schema.StringAttribute{
					MarkdownDescription: "The ACME challenge type used to validate domain ownership",
					Computed:            true,
				},
				"contact": schema.SetAttribute{
					MarkdownDescription: "Contact email address, which is used for important communications regarding your ACME account and certificates",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Descriptive label for this provider, built from the directory URL and the contact address",
					Computed:            true,
				},
				"directory": schema.StringAttribute{
					MarkdownDescription: "The URL of the ACME directory endpoint",
					Computed:            true,
				},
				"eab_hmac_key": schema.StringAttribute{
					MarkdownDescription: "The External Account Binding (EAB) HMAC key",
					Computed:            true,
					Sensitive:           true,
				},
				"eab_key_id": schema.StringAttribute{
					MarkdownDescription: "The External Account Binding (EAB) key ID",
					Computed:            true,
				},
				"max_retries": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of retry attempts for failed challenges",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this ACME provider belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"preferred_chain": schema.StringAttribute{
					MarkdownDescription: "Preferred certificate chain to use when multiple chains are available",
					Computed:            true,
				},
				"renew_before": schema.StringAttribute{
					MarkdownDescription: "How long before expiration the certificate should be renewed",
					Computed:            true,
				},
				"reuse_key": schema.BoolAttribute{
					MarkdownDescription: "Whether to reuse the existing private key when renewing a certificate",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "acme_providers",
		JMAPType: "x:AcmeProvider",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:AcmeProvider` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "address_book",
		JMAPType:  "x:AddressBook",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures address book and contact storage settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"default_display_name": schema.StringAttribute{
					MarkdownDescription: "Specifies the default display name for a contact when it is created",
					Computed:            true,
				},
				"default_href_name": schema.StringAttribute{
					MarkdownDescription: "Specifies the default href name for a contact when it is created",
					Computed:            true,
				},
				"max_address_books": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of address books a user can create",
					Computed:            true,
				},
				"max_contacts": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of contact cards a user can create",
					Computed:            true,
				},
				"max_v_card_size": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum size of a vCard file that can be uploaded to the server",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "ai_model",
		JMAPType: "x:AiModel",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines an AI model endpoint for LLM-based features. Requires an Enterprise license.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Whether Stalwart should connect to an endpoint that has an invalid TLS certificate",
					Computed:            true,
				},
				"http_auth": schema.SingleNestedAttribute{
					MarkdownDescription: "The type of HTTP authentication to use",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"bearer_token": schema.SingleNestedAttribute{
							MarkdownDescription: "Bearer token for HTTP Bearer Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password for HTTP Basic Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Username for HTTP Basic Authentication",
							Computed:            true,
						},
					},
				},
				"http_headers": schema.MapAttribute{
					MarkdownDescription: "Additional headers to include in HTTP requests",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"model": schema.StringAttribute{
					MarkdownDescription: "The name of the AI model to use.",
					Computed:            true,
				},
				"model_type": schema.StringAttribute{
					MarkdownDescription: "API type",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short name for the AI Model",
					Optional:            true,
					Computed:            true,
				},
				"temperature": schema.Float64Attribute{
					MarkdownDescription: "The temperature of the AI model, which controls the randomness of the output. A higher temperature will produce more random output.",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that Stalwart will wait for a response from this endpoint",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "URL of the OpenAI compatible endpoint",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "ai_models",
		JMAPType: "x:AiModel",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:AiModel` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "alert",
		JMAPType: "x:Alert",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an alert rule triggered by metric conditions. Requires an Enterprise license.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"condition": schema.SingleNestedAttribute{
					MarkdownDescription: "The condition that triggers the alert.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"email_alert": schema.SingleNestedAttribute{
					MarkdownDescription: "Email notification settings",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"body": schema.StringAttribute{
							MarkdownDescription: "The body of the email",
							Computed:            true,
						},
						"from_address": schema.StringAttribute{
							MarkdownDescription: "The email address of the sender",
							Computed:            true,
						},
						"from_name": schema.StringAttribute{
							MarkdownDescription: "The name of the sender",
							Computed:            true,
						},
						"subject": schema.StringAttribute{
							MarkdownDescription: "The subject of the email",
							Computed:            true,
						},
						"to": schema.SetAttribute{
							MarkdownDescription: "The email address of the recipient(s)",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable or disable the alert",
					Computed:            true,
				},
				"event_alert": schema.SingleNestedAttribute{
					MarkdownDescription: "Event notification settings",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"event_message": schema.StringAttribute{
							MarkdownDescription: "The message of the event to trigger",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:     "alerts",
		JMAPType: "x:Alert",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Alert` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "allowed_ip",
		JMAPType: "x:AllowedIp",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an allowed IP address or network range.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"address": schema.StringAttribute{
					MarkdownDescription: "The IP address or mask to allow",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "The date and time when this IP address was allowed",
					Computed:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "The date and time when this IP address allowance will expire",
					Computed:            true,
				},
				"reason": schema.StringAttribute{
					MarkdownDescription: "The reason for allowing this IP address",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "allowed_ips",
		JMAPType: "x:AllowedIp",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:AllowedIp` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "api_key",
		JMAPType: "x:ApiKey",
		Schema: schema.Schema{
			MarkdownDescription: "API key credential for programmatic access.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"allowed_ips": schema.SetAttribute{
					MarkdownDescription: "List of allowed IP addresses or CIDR ranges for this credential",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the credential",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the credential",
					Computed:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "Expiration date of the credential",
					Computed:            true,
				},
				"permissions": schema.SingleNestedAttribute{
					MarkdownDescription: "List of permissions assigned to this credential",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions to assign.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"secret": schema.StringAttribute{
					MarkdownDescription: "Secret value of the credential",
					Computed:            true,
					Sensitive:           true,
				},
			},
		},
	},
	{
		Name:     "api_keys",
		JMAPType: "x:ApiKey",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:ApiKey` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "app_password",
		JMAPType: "x:AppPassword",
		Schema: schema.Schema{
			MarkdownDescription: "App password credential for programmatic access.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"allowed_ips": schema.SetAttribute{
					MarkdownDescription: "List of allowed IP addresses or CIDR ranges for this credential",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the credential",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the credential",
					Computed:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "Expiration date of the credential",
					Computed:            true,
				},
				"permissions": schema.SingleNestedAttribute{
					MarkdownDescription: "List of permissions assigned to this credential",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions to assign.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"secret": schema.StringAttribute{
					MarkdownDescription: "Secret value of the credential",
					Computed:            true,
					Sensitive:           true,
				},
			},
		},
	},
	{
		Name:     "app_passwords",
		JMAPType: "x:AppPassword",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:AppPassword` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "application",
		JMAPType: "x:Application",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a web application served by the server.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auto_update_frequency": schema.Int64Attribute{
					MarkdownDescription: "Frequency to check for application updates",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "A short description of the web application",
					Computed:            true,
				},
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Whether the application is enabled and should be served by the server",
					Computed:            true,
				},
				"resource_url": schema.StringAttribute{
					MarkdownDescription: "Override the URL to download application updates from.",
					Computed:            true,
				},
				"unpack_directory": schema.StringAttribute{
					MarkdownDescription: "The local path to unpack the application bundle to. If left empty, the application will be unpacked to /tmp.",
					Computed:            true,
				},
				"url_prefix": schema.SetAttribute{
					MarkdownDescription: "The URL prefixes to serve the application on. For example, if set to \"/admin\", the application will be accessible at http://server/admin.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "applications",
		JMAPType: "x:Application",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Application` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "archived_item",
		JMAPType: "x:ArchivedItem",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Represents an archived item that can be restored.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"account_id": schema.StringAttribute{
					MarkdownDescription: "The account to which the archived item belongs",
					Computed:            true,
				},
				"archived_at": schema.StringAttribute{
					MarkdownDescription: "Timestamp when the item was archived",
					Computed:            true,
				},
				"archived_until": schema.StringAttribute{
					MarkdownDescription: "Timestamp until which the archived item will be deleted permanently if not restored",
					Computed:            true,
				},
				"blob_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the archived blob",
					Computed:            true,
				},
				"content": schema.StringAttribute{
					MarkdownDescription: "Content of the archived sieve script",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the archived file",
					Computed:            true,
				},
				"from": schema.StringAttribute{
					MarkdownDescription: "Sender of the archived email message",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the archived file",
					Optional:            true,
					Computed:            true,
				},
				"received_at": schema.StringAttribute{
					MarkdownDescription: "Received date of the archived email message",
					Computed:            true,
				},
				"size": schema.Int64Attribute{
					MarkdownDescription: "Size of the archived email message",
					Computed:            true,
				},
				"start_time": schema.StringAttribute{
					MarkdownDescription: "Start time of the archived calendar event",
					Computed:            true,
				},
				"status": schema.StringAttribute{
					MarkdownDescription: "Current status of the archived item",
					Computed:            true,
				},
				"subject": schema.StringAttribute{
					MarkdownDescription: "Subject of the archived email message",
					Computed:            true,
				},
				"title": schema.StringAttribute{
					MarkdownDescription: "Title of the archived calendar event",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "archived_items",
		JMAPType: "x:ArchivedItem",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:ArchivedItem` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "arf_external_report",
		JMAPType: "x:ArfExternalReport",
		Schema: schema.Schema{
			MarkdownDescription: "Stores an ARF feedback report received from an external source.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "When the report is scheduled to be deleted",
					Computed:            true,
				},
				"from": schema.StringAttribute{
					MarkdownDescription: "Email address of the report sender",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this report belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"received_at": schema.StringAttribute{
					MarkdownDescription: "When the report email was received",
					Computed:            true,
				},
				"report": schema.SingleNestedAttribute{
					MarkdownDescription: "Parsed ARF feedback report content",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"arrival_date": schema.StringAttribute{
							MarkdownDescription: "When the original message arrived",
							Computed:            true,
						},
						"auth_failure": schema.StringAttribute{
							MarkdownDescription: "Type of authentication failure (for auth-failure reports)",
							Computed:            true,
						},
						"authentication_results": schema.SetAttribute{
							MarkdownDescription: "Authentication-Results header values from the original message",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"delivery_result": schema.StringAttribute{
							MarkdownDescription: "What happened to the original message",
							Computed:            true,
						},
						"dkim_adsp_dns": schema.StringAttribute{
							MarkdownDescription: "DKIM ADSP DNS record content",
							Computed:            true,
						},
						"dkim_canonicalized_body": schema.StringAttribute{
							MarkdownDescription: "Message body after DKIM canonicalization",
							Computed:            true,
						},
						"dkim_canonicalized_header": schema.StringAttribute{
							MarkdownDescription: "Message headers after DKIM canonicalization",
							Computed:            true,
						},
						"dkim_domain": schema.StringAttribute{
							MarkdownDescription: "Domain from the DKIM signature",
							Computed:            true,
						},
						"dkim_identity": schema.StringAttribute{
							MarkdownDescription: "Identity from the DKIM signature (i= tag)",
							Computed:            true,
						},
						"dkim_selector": schema.StringAttribute{
							MarkdownDescription: "Selector from the DKIM signature",
							Computed:            true,
						},
						"dkim_selector_dns": schema.StringAttribute{
							MarkdownDescription: "DKIM selector DNS record content",
							Computed:            true,
						},
						"feedback_type": schema.StringAttribute{
							MarkdownDescription: "Type of feedback being reported",
							Computed:            true,
						},
						"headers": schema.StringAttribute{
							MarkdownDescription: "Original message headers that triggered the report",
							Computed:            true,
						},
						"identity_alignment": schema.StringAttribute{
							MarkdownDescription: "Which identities were aligned",
							Computed:            true,
						},
						"incidents": schema.Int64Attribute{
							MarkdownDescription: "Number of incidents represented by this report",
							Computed:            true,
						},
						"message": schema.StringAttribute{
							MarkdownDescription: "Original message content that triggered the report",
							Computed:            true,
						},
						"original_envelope_id": schema.StringAttribute{
							MarkdownDescription: "Original SMTP envelope ID (ENVID)",
							Computed:            true,
						},
						"original_mail_from": schema.StringAttribute{
							MarkdownDescription: "Original envelope sender address (MAIL FROM)",
							Computed:            true,
						},
						"original_rcpt_to": schema.StringAttribute{
							MarkdownDescription: "Original envelope recipient address (RCPT TO)",
							Computed:            true,
						},
						"reported_domains": schema.SetAttribute{
							MarkdownDescription: "Domains being reported",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"reported_uris": schema.SetAttribute{
							MarkdownDescription: "URIs being reported",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"reporting_mta": schema.StringAttribute{
							MarkdownDescription: "Hostname of the MTA generating this report",
							Computed:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: "IP address of the original message source",
							Computed:            true,
						},
						"source_port": schema.Int64Attribute{
							MarkdownDescription: "Port of the original message source",
							Computed:            true,
						},
						"spf_dns": schema.StringAttribute{
							MarkdownDescription: "SPF DNS record content",
							Computed:            true,
						},
						"user_agent": schema.StringAttribute{
							MarkdownDescription: "Software that generated this report",
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "ARF format version",
							Computed:            true,
						},
					},
				},
				"subject": schema.StringAttribute{
					MarkdownDescription: "Subject line of the report email",
					Computed:            true,
				},
				"to": schema.SetAttribute{
					MarkdownDescription: "List of recipient email addresses",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "arf_external_reports",
		JMAPType: "x:ArfExternalReport",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:ArfExternalReport` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "asn",
		JMAPType:  "x:Asn",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures ASN and geolocation data sources for IP address lookups. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"asn_urls": schema.SetAttribute{
					MarkdownDescription: "URLs to fetch CSV file containing the IP to ASN mappings.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"expires": schema.Int64Attribute{
					MarkdownDescription: "How often to refresh the ASN/Geo data.",
					Computed:            true,
				},
				"geo_urls": schema.SetAttribute{
					MarkdownDescription: "URLs to fetch CSV file containing the IP to country code mappings.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"http_auth": schema.SingleNestedAttribute{
					MarkdownDescription: "The type of HTTP authentication to use",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"bearer_token": schema.SingleNestedAttribute{
							MarkdownDescription: "Bearer token for HTTP Bearer Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password for HTTP Basic Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Username for HTTP Basic Authentication",
							Computed:            true,
						},
					},
				},
				"http_headers": schema.MapAttribute{
					MarkdownDescription: "Additional headers to include in HTTP requests",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"index_asn": schema.Int64Attribute{
					MarkdownDescription: "The position of the ASN in the DNS TXT record.",
					Computed:            true,
				},
				"index_asn_name": schema.Int64Attribute{
					MarkdownDescription: "The position of the ASN Name in the DNS TXT record.",
					Computed:            true,
				},
				"index_country": schema.Int64Attribute{
					MarkdownDescription: "The position of the country code in the DNS TXT record.",
					Computed:            true,
				},
				"max_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the ASN/Geo data file.",
					Computed:            true,
				},
				"separator": schema.StringAttribute{
					MarkdownDescription: "The separator character used in the DNS TXT record.",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Time after which the ASN/Geo resource fetch is considered failed.",
					Computed:            true,
				},
				"zone_ip_v4": schema.StringAttribute{
					MarkdownDescription: "The DNS zone to query for IPv4 ASN and geolocation data.",
					Computed:            true,
				},
				"zone_ip_v6": schema.StringAttribute{
					MarkdownDescription: "The DNS zone to query for IPv6 ASN and geolocation data.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "authentication",
		JMAPType:  "x:Authentication",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures authentication settings including password policies and default roles. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"default_admin_role_ids": schema.SetAttribute{
					MarkdownDescription: "Default roles to assign for administrators.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"default_group_role_ids": schema.SetAttribute{
					MarkdownDescription: "Default roles to assign for groups.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"default_tenant_role_ids": schema.SetAttribute{
					MarkdownDescription: "Default roles to assign for tenants in multi-tenant environments. Requires an Enterprise license.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"default_user_role_ids": schema.SetAttribute{
					MarkdownDescription: "Default roles to assign for accounts.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"directory_id": schema.StringAttribute{
					MarkdownDescription: "External directory used for authentication, or null to use the internal directory",
					Computed:            true,
				},
				"max_api_keys": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of API keys a user can create",
					Computed:            true,
				},
				"max_app_passwords": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of app passwords a user can create",
					Computed:            true,
				},
				"password_default_expiry": schema.Int64Attribute{
					MarkdownDescription: "Default expiration time for user passwords in the internal directory, after which the user will be required to change their password.",
					Computed:            true,
				},
				"password_hash_algorithm": schema.StringAttribute{
					MarkdownDescription: "Password hashing algorithm to use for storing user passwords in the internal directory.",
					Computed:            true,
				},
				"password_max_length": schema.Int64Attribute{
					MarkdownDescription: "Maximum length for user passwords in the internal directory.",
					Computed:            true,
				},
				"password_min_length": schema.Int64Attribute{
					MarkdownDescription: "Minimum length for user passwords in the internal directory.",
					Computed:            true,
				},
				"password_min_strength": schema.StringAttribute{
					MarkdownDescription: "Minimum strength for user passwords in the internal directory, calculated using the zxcvbn algorithm.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "blob_store",
		JMAPType:  "x:BlobStore",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the blob storage backend for messages and files. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"access_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Identifies the S3 account",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "Value",
							Computed:            true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Allow invalid TLS certificates when connecting to the S3 service",
					Computed:            true,
				},
				"auth_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the store",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"auth_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the store",
					Computed:            true,
				},
				"bucket": schema.StringAttribute{
					MarkdownDescription: "The S3 bucket where blobs (e-mail messages, Sieve scripts, etc.) will be stored",
					Computed:            true,
				},
				"cluster_file": schema.StringAttribute{
					MarkdownDescription: "Path to the cluster file for the FoundationDB cluster",
					Computed:            true,
				},
				"container": schema.StringAttribute{
					MarkdownDescription: "The name of the container in the Storage Account",
					Computed:            true,
				},
				"database": schema.StringAttribute{
					MarkdownDescription: "Name of the database",
					Computed:            true,
				},
				"datacenter_id": schema.StringAttribute{
					MarkdownDescription: "Data center ID (optional)",
					Computed:            true,
				},
				"depth": schema.Int64Attribute{
					MarkdownDescription: "Maximum depth of nested directories",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "Hostname of the database server",
					Computed:            true,
				},
				"key_prefix": schema.StringAttribute{
					MarkdownDescription: "A prefix that will be added to the keys of all objects stored in the blob store",
					Computed:            true,
				},
				"machine_id": schema.StringAttribute{
					MarkdownDescription: "Machine ID in the FoundationDB cluster (optional)",
					Computed:            true,
				},
				"max_allowed_packet": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a packet in bytes",
					Computed:            true,
				},
				"max_retries": schema.Int64Attribute{
					MarkdownDescription: "The maximum number of times to retry failed requests. Set to 0 to disable retries",
					Computed:            true,
				},
				"options": schema.StringAttribute{
					MarkdownDescription: "Additional connection options",
					Computed:            true,
				},
				"path": schema.StringAttribute{
					MarkdownDescription: "Where to store the data in the server's filesystem",
					Computed:            true,
				},
				"pool_max_connections": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of connections to the store",
					Computed:            true,
				},
				"pool_min_connections": schema.Int64Attribute{
					MarkdownDescription: "Minimum number of connections to the store",
					Computed:            true,
				},
				"pool_recycling_method": schema.StringAttribute{
					MarkdownDescription: "Method to use when recycling connections in the pool",
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "Port of the database server",
					Computed:            true,
				},
				"profile": schema.StringAttribute{
					MarkdownDescription: "Used when retrieving credentials from a shared credentials file. If specified, the server will use the access key ID, secret access key, and session token (if available) associated with the given profile",
					Computed:            true,
				},
				"read_replicas": schema.ListNestedAttribute{
					MarkdownDescription: "List of read replicas for the store Requires an Enterprise license.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"auth_secret": schema.SingleNestedAttribute{
								MarkdownDescription: "Password to connect to the store",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"auth_username": schema.StringAttribute{
								MarkdownDescription: "Username to connect to the store",
								Computed:            true,
							},
							"database": schema.StringAttribute{
								MarkdownDescription: "Name of the database",
								Computed:            true,
							},
							"host": schema.StringAttribute{
								MarkdownDescription: "Hostname of the database server",
								Computed:            true,
							},
							"options": schema.StringAttribute{
								MarkdownDescription: "Additional connection options",
								Computed:            true,
							},
							"port": schema.Int64Attribute{
								MarkdownDescription: "Port of the database server",
								Computed:            true,
							},
						},
					},
				},
				"region": schema.SingleNestedAttribute{
					MarkdownDescription: "The S3 region where the bucket resides",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"custom_endpoint": schema.StringAttribute{
							MarkdownDescription: "Endpoint URL",
							Computed:            true,
						},
						"custom_region": schema.StringAttribute{
							MarkdownDescription: "Region name",
							Computed:            true,
						},
					},
				},
				"sas_token": schema.SingleNestedAttribute{
					MarkdownDescription: "SAS Token, when not using accessKey based authentication",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"secret_key": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret key for the S3 account",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"security_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Security token for temporary credentials",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"session_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Temporary session token for the S3 account",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"storage_account": schema.StringAttribute{
					MarkdownDescription: "The Azure Storage Account where blobs (e-mail messages, Sieve scripts, etc.) will be stored",
					Computed:            true,
				},
				"stores": schema.ListNestedAttribute{
					MarkdownDescription: "Stores to use for sharding",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								MarkdownDescription: "Variant discriminator.",
								Computed:            true,
							},
							"access_key": schema.SingleNestedAttribute{
								MarkdownDescription: "Identifies the S3 account",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Value",
										Computed:            true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"allow_invalid_certs": schema.BoolAttribute{
								MarkdownDescription: "Allow invalid TLS certificates when connecting to the S3 service",
								Computed:            true,
							},
							"auth_secret": schema.SingleNestedAttribute{
								MarkdownDescription: "Password to connect to the store",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"auth_username": schema.StringAttribute{
								MarkdownDescription: "Username to connect to the store",
								Computed:            true,
							},
							"bucket": schema.StringAttribute{
								MarkdownDescription: "The S3 bucket where blobs (e-mail messages, Sieve scripts, etc.) will be stored",
								Computed:            true,
							},
							"cluster_file": schema.StringAttribute{
								MarkdownDescription: "Path to the cluster file for the FoundationDB cluster",
								Computed:            true,
							},
							"container": schema.StringAttribute{
								MarkdownDescription: "The name of the container in the Storage Account",
								Computed:            true,
							},
							"database": schema.StringAttribute{
								MarkdownDescription: "Name of the database",
								Computed:            true,
							},
							"datacenter_id": schema.StringAttribute{
								MarkdownDescription: "Data center ID (optional)",
								Computed:            true,
							},
							"depth": schema.Int64Attribute{
								MarkdownDescription: "Maximum depth of nested directories",
								Computed:            true,
							},
							"host": schema.StringAttribute{
								MarkdownDescription: "Hostname of the database server",
								Computed:            true,
							},
							"key_prefix": schema.StringAttribute{
								MarkdownDescription: "A prefix that will be added to the keys of all objects stored in the blob store",
								Computed:            true,
							},
							"machine_id": schema.StringAttribute{
								MarkdownDescription: "Machine ID in the FoundationDB cluster (optional)",
								Computed:            true,
							},
							"max_allowed_packet": schema.Int64Attribute{
								MarkdownDescription: "Maximum size of a packet in bytes",
								Computed:            true,
							},
							"max_retries": schema.Int64Attribute{
								MarkdownDescription: "The maximum number of times to retry failed requests. Set to 0 to disable retries",
								Computed:            true,
							},
							"options": schema.StringAttribute{
								MarkdownDescription: "Additional connection options",
								Computed:            true,
							},
							"path": schema.StringAttribute{
								MarkdownDescription: "Where to store the data in the server's filesystem",
								Computed:            true,
							},
							"pool_max_connections": schema.Int64Attribute{
								MarkdownDescription: "Maximum number of connections to the store",
								Computed:            true,
							},
							"pool_min_connections": schema.Int64Attribute{
								MarkdownDescription: "Minimum number of connections to the store",
								Computed:            true,
							},
							"pool_recycling_method": schema.StringAttribute{
								MarkdownDescription: "Method to use when recycling connections in the pool",
								Computed:            true,
							},
							"port": schema.Int64Attribute{
								MarkdownDescription: "Port of the database server",
								Computed:            true,
							},
							"profile": schema.StringAttribute{
								MarkdownDescription: "Used when retrieving credentials from a shared credentials file. If specified, the server will use the access key ID, secret access key, and session token (if available) associated with the given profile",
								Computed:            true,
							},
							"read_replicas": schema.ListNestedAttribute{
								MarkdownDescription: "List of read replicas for the store Requires an Enterprise license.",
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"auth_secret": schema.SingleNestedAttribute{
											MarkdownDescription: "Password to connect to the store",
											Computed:            true,
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "Variant discriminator.",
													Computed:            true,
												},
												"file_path": schema.StringAttribute{
													MarkdownDescription: "File path to read the secret from",
													Computed:            true,
												},
												"secret": schema.StringAttribute{
													MarkdownDescription: "Password or secret value",
													Computed:            true,
													Sensitive:           true,
												},
												"variable_name": schema.StringAttribute{
													MarkdownDescription: "Environment variable name to read the secret from",
													Computed:            true,
												},
											},
										},
										"auth_username": schema.StringAttribute{
											MarkdownDescription: "Username to connect to the store",
											Computed:            true,
										},
										"database": schema.StringAttribute{
											MarkdownDescription: "Name of the database",
											Computed:            true,
										},
										"host": schema.StringAttribute{
											MarkdownDescription: "Hostname of the database server",
											Computed:            true,
										},
										"options": schema.StringAttribute{
											MarkdownDescription: "Additional connection options",
											Computed:            true,
										},
										"port": schema.Int64Attribute{
											MarkdownDescription: "Port of the database server",
											Computed:            true,
										},
									},
								},
							},
							"region": schema.SingleNestedAttribute{
								MarkdownDescription: "The S3 region where the bucket resides",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"custom_endpoint": schema.StringAttribute{
										MarkdownDescription: "Endpoint URL",
										Computed:            true,
									},
									"custom_region": schema.StringAttribute{
										MarkdownDescription: "Region name",
										Computed:            true,
									},
								},
							},
							"sas_token": schema.SingleNestedAttribute{
								MarkdownDescription: "SAS Token, when not using accessKey based authentication",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"secret_key": schema.SingleNestedAttribute{
								MarkdownDescription: "The secret key for the S3 account",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"security_token": schema.SingleNestedAttribute{
								MarkdownDescription: "Security token for temporary credentials",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"session_token": schema.SingleNestedAttribute{
								MarkdownDescription: "Temporary session token for the S3 account",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"storage_account": schema.StringAttribute{
								MarkdownDescription: "The Azure Storage Account where blobs (e-mail messages, Sieve scripts, etc.) will be stored",
								Computed:            true,
							},
							"timeout": schema.Int64Attribute{
								MarkdownDescription: "Connection timeout to the S3 service",
								Computed:            true,
							},
							"transaction_retry_delay": schema.Int64Attribute{
								MarkdownDescription: "Transaction maximum retry delay",
								Computed:            true,
							},
							"transaction_retry_limit": schema.Int64Attribute{
								MarkdownDescription: "Transaction retry limit",
								Computed:            true,
							},
							"transaction_timeout": schema.Int64Attribute{
								MarkdownDescription: "Transaction timeout",
								Computed:            true,
							},
							"use_tls": schema.BoolAttribute{
								MarkdownDescription: "Use TLS to connect to the store",
								Computed:            true,
							},
							"verify_after_write": schema.BoolAttribute{
								MarkdownDescription: "After each successful write, verify the object is readable on the backend. Defends against the rare case where an S3-compatible backend returns success but does not actually persist the data. Adds one extra request per write.",
								Computed:            true,
							},
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Connection timeout to the S3 service",
					Computed:            true,
				},
				"transaction_retry_delay": schema.Int64Attribute{
					MarkdownDescription: "Transaction maximum retry delay",
					Computed:            true,
				},
				"transaction_retry_limit": schema.Int64Attribute{
					MarkdownDescription: "Transaction retry limit",
					Computed:            true,
				},
				"transaction_timeout": schema.Int64Attribute{
					MarkdownDescription: "Transaction timeout",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Use TLS to connect to the store",
					Computed:            true,
				},
				"verify_after_write": schema.BoolAttribute{
					MarkdownDescription: "After each successful write, verify the object is readable on the backend. Defends against the rare case where an S3-compatible backend returns success but does not actually persist the data. Adds one extra request per write.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "blocked_ip",
		JMAPType: "x:BlockedIp",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a blocked IP address or network range.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"address": schema.StringAttribute{
					MarkdownDescription: "The IP address or mask to block",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "The date and time when this IP address was blocked",
					Computed:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "The date and time when this IP address block will expire",
					Computed:            true,
				},
				"reason": schema.StringAttribute{
					MarkdownDescription: "The reason for blocking this IP address",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "blocked_ips",
		JMAPType: "x:BlockedIp",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:BlockedIp` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "cache",
		JMAPType:  "x:Cache",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures in-memory cache sizes for data, DNS records, and authorization tokens. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"access_tokens": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the access tokens cache",
					Computed:            true,
				},
				"accounts": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the accounts cache",
					Computed:            true,
				},
				"contacts": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the address books and contacts cache",
					Computed:            true,
				},
				"dkim_signatures": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the DKIM signatures cache",
					Computed:            true,
				},
				"dns_ipv4": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the IPv4 record cache",
					Computed:            true,
				},
				"dns_ipv6": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the IPv6 record cache",
					Computed:            true,
				},
				"dns_mta_sts": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the MTA-STS record cache",
					Computed:            true,
				},
				"dns_mx": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the MX record cache",
					Computed:            true,
				},
				"dns_ptr": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the PTR record cache",
					Computed:            true,
				},
				"dns_rbl": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the DNSBl record cache",
					Computed:            true,
				},
				"dns_tlsa": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the TLSA record cache",
					Computed:            true,
				},
				"dns_txt": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the TXT record cache",
					Computed:            true,
				},
				"domain_names": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the domain name lookup cache",
					Computed:            true,
				},
				"domain_names_negative": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the domain name lookup negative cache",
					Computed:            true,
				},
				"domains": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the domains cache",
					Computed:            true,
				},
				"email_addresses": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the email addresses lookup cache",
					Computed:            true,
				},
				"email_addresses_negative": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the email addresses lookup negative cache",
					Computed:            true,
				},
				"events": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the calendar and events cache",
					Computed:            true,
				},
				"files": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the file storage data cache",
					Computed:            true,
				},
				"http_auth": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the HTTP authorization headers cache",
					Computed:            true,
				},
				"mailing_lists": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the mailing lists cache",
					Computed:            true,
				},
				"messages": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the e-mail data cache",
					Computed:            true,
				},
				"negative_ttl": schema.Int64Attribute{
					MarkdownDescription: "Time-to-live for domain and account name lookup negative cache entries",
					Computed:            true,
				},
				"roles": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the roles cache",
					Computed:            true,
				},
				"scheduling": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the scheduling cache",
					Computed:            true,
				},
				"tenants": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the tenants cache",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "calendar",
		JMAPType:  "x:Calendar",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures calendar settings including iCalendar limits and default names. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"default_display_name": schema.StringAttribute{
					MarkdownDescription: "Specifies the default display name for a calendar when it is created",
					Computed:            true,
				},
				"default_href_name": schema.StringAttribute{
					MarkdownDescription: "Specifies the default href name for a calendar when it is created",
					Computed:            true,
				},
				"max_attendees": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum number of attendees that can be included in a single iCalendar instance",
					Computed:            true,
				},
				"max_calendars": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of calendars a user can create",
					Computed:            true,
				},
				"max_event_notifications": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of notifications a user can create for calendar events",
					Computed:            true,
				},
				"max_events": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of calendar events a user can create",
					Computed:            true,
				},
				"max_i_calendar_size": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum size of an iCalendar file that can be uploaded to the server",
					Computed:            true,
				},
				"max_participant_identities": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of participant identities a user can create",
					Computed:            true,
				},
				"max_recurrence_expansions": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum number of instances that can be generated from a recurring iCalendar event",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "calendar_alarm",
		JMAPType:  "x:CalendarAlarm",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures calendar alarm email notifications. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"allow_external_rcpts": schema.BoolAttribute{
					MarkdownDescription: "Allows calendar alarms to be sent to external recipients, enabling notifications to users outside the server",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enables the calendar alarms feature, allowing users to set alarms for events and receive notifications via e-mail",
					Computed:            true,
				},
				"from_email": schema.StringAttribute{
					MarkdownDescription: "Specifies the e-mail address that will appear in the 'From' field of calendar alarm e-mails, ensuring that users can reply to or contact the sender",
					Computed:            true,
				},
				"from_name": schema.StringAttribute{
					MarkdownDescription: "Specifies the name that will appear in the 'From' field of calendar alarm e-mails, providing a recognizable sender name for users",
					Computed:            true,
				},
				"min_trigger_interval": schema.Int64Attribute{
					MarkdownDescription: "Specifies the minimum interval for calendar alarms, ensuring that alarms are not triggered too frequently",
					Computed:            true,
				},
				"template": schema.StringAttribute{
					MarkdownDescription: "Specifies the HTML template used for rendering calendar alarm e-mails, allowing customization of the alarm notification format Requires an Enterprise license.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "calendar_scheduling",
		JMAPType:  "x:CalendarScheduling",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures calendar scheduling, iTIP messaging, and HTTP RSVP settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"auto_add_invitations": schema.BoolAttribute{
					MarkdownDescription: "Automatically adds incoming invitations to the user's calendar.",
					Computed:            true,
				},
				"email_template": schema.StringAttribute{
					MarkdownDescription: "Specifies the HTML template used for rendering iMIP invitations. Requires an Enterprise license.",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enables the scheduling features for calendar events, allowing users to send and receive invitations",
					Computed:            true,
				},
				"http_rsvp_enable": schema.BoolAttribute{
					MarkdownDescription: "Enables the HTTP RSVP feature for calendar invitations, allowing users to respond via a web interface.",
					Computed:            true,
				},
				"http_rsvp_link_expiry": schema.Int64Attribute{
					MarkdownDescription: "Sets the expiration duration for HTTP RSVP links, after which they will no longer be valid.",
					Computed:            true,
				},
				"http_rsvp_template": schema.StringAttribute{
					MarkdownDescription: "Specifies the HTML template used for rendering HTTP RSVP confirmations. Requires an Enterprise license.",
					Computed:            true,
				},
				"http_rsvp_url": schema.StringAttribute{
					MarkdownDescription: "Specifies a custom URL for the HTTP RSVP endpoint, where users can respond to calendar invitations.",
					Computed:            true,
				},
				"itip_max_size": schema.Int64Attribute{
					MarkdownDescription: "Sets the maximum iCalendar object size for incoming iTIP messages.",
					Computed:            true,
				},
				"max_recipients": schema.Int64Attribute{
					MarkdownDescription: "Sets the maximum number of recipients for outbound iTIP messages.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "certificate",
		JMAPType: "x:Certificate",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a TLS certificate and its associated private key.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"certificate": schema.SingleNestedAttribute{
					MarkdownDescription: "TLS certificate in PEM format",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "Text value",
							Computed:            true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"issuer": schema.StringAttribute{
					MarkdownDescription: "Certificate issuer",
					Computed:            true,
				},
				"not_valid_after": schema.StringAttribute{
					MarkdownDescription: "Expiration date of the certificate",
					Computed:            true,
				},
				"not_valid_before": schema.StringAttribute{
					MarkdownDescription: "Issuance date of the certificate",
					Computed:            true,
				},
				"private_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Private key in PEM format",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"subject_alternative_names": schema.SetAttribute{
					MarkdownDescription: "Subject Alternative Names (SAN) for the certificate",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "certificates",
		JMAPType: "x:Certificate",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Certificate` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "cluster_node",
		JMAPType: "x:ClusterNode",
		Schema: schema.Schema{
			MarkdownDescription: "Represents a node in the cluster",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"hostname": schema.StringAttribute{
					MarkdownDescription: "Hostname of the node",
					Computed:            true,
				},
				"last_renewal": schema.StringAttribute{
					MarkdownDescription: "Timestamp of the last lease renewal from this node, used to determine if the node is still active in the cluster",
					Computed:            true,
				},
				"node_id": schema.Int64Attribute{
					MarkdownDescription: "Unique identifier for the node in the cluster",
					Computed:            true,
				},
				"status": schema.StringAttribute{
					MarkdownDescription: "Current status of the node in the cluster",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "cluster_nodes",
		JMAPType: "x:ClusterNode",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:ClusterNode` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "cluster_role",
		JMAPType: "x:ClusterRole",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a cluster node role with enabled tasks and listeners.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the role",
					Computed:            true,
				},
				"listeners": schema.SingleNestedAttribute{
					MarkdownDescription: "Which network listeners are enabled for this cluster role",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"listener_ids": schema.SetAttribute{
							MarkdownDescription: "List of network listeners to enable or disable for this group",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Unique identifier for the role",
					Optional:            true,
					Computed:            true,
				},
				"tasks": schema.SingleNestedAttribute{
					MarkdownDescription: "Which tasks are enabled for this cluster role",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"task_types": schema.SetAttribute{
							MarkdownDescription: "Tasks to enable or disable for this group",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	},
	{
		Name:     "cluster_roles",
		JMAPType: "x:ClusterRole",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:ClusterRole` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "coordinator",
		JMAPType:  "x:Coordinator",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the cluster coordinator for inter-node communication. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"addresses": schema.SetAttribute{
					MarkdownDescription: "Address of the NATS server",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"auth_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the store",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"auth_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the store",
					Computed:            true,
				},
				"brokers": schema.SetAttribute{
					MarkdownDescription: "List of Kafka brokers",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"capacity_client": schema.Int64Attribute{
					MarkdownDescription: "By default, Client dispatches op's to the Client onto the channel with capacity of 2048. This option enables overriding it",
					Computed:            true,
				},
				"capacity_read_buffer": schema.Int64Attribute{
					MarkdownDescription: "Sets the initial capacity of the read buffer. Which is a buffer used to gather partial protocol messages.",
					Computed:            true,
				},
				"capacity_subscription": schema.Int64Attribute{
					MarkdownDescription: "Sets the capacity for Subscribers. Exceeding it will trigger slow consumer error callback and drop messages.",
					Computed:            true,
				},
				"config": schema.StringAttribute{
					MarkdownDescription: "Zenoh configuration string",
					Computed:            true,
				},
				"credentials": schema.SingleNestedAttribute{
					MarkdownDescription: "String containing the JWT credentials",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"group_id": schema.StringAttribute{
					MarkdownDescription: "Consumer group ID",
					Computed:            true,
				},
				"max_reconnects": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of times to attempt to reconnect to the server",
					Computed:            true,
				},
				"max_retries": schema.Int64Attribute{
					MarkdownDescription: "Number of retries to connect to the Redis cluster",
					Computed:            true,
				},
				"max_retry_wait": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait between retries",
					Computed:            true,
				},
				"min_retry_wait": schema.Int64Attribute{
					MarkdownDescription: "Minimum time to wait between retries",
					Computed:            true,
				},
				"no_echo": schema.BoolAttribute{
					MarkdownDescription: "Disables delivering messages that were published from the same connection.",
					Computed:            true,
				},
				"ping_interval": schema.Int64Attribute{
					MarkdownDescription: "Interval between pings to the server",
					Computed:            true,
				},
				"pool_max_connections": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of connections to the store",
					Computed:            true,
				},
				"pool_timeout_create": schema.Int64Attribute{
					MarkdownDescription: "Timeout for creating a new connection",
					Computed:            true,
				},
				"pool_timeout_recycle": schema.Int64Attribute{
					MarkdownDescription: "Timeout for recycling a connection",
					Computed:            true,
				},
				"pool_timeout_wait": schema.Int64Attribute{
					MarkdownDescription: "Timeout for waiting for a connection from the pool",
					Computed:            true,
				},
				"protocol_version": schema.StringAttribute{
					MarkdownDescription: "Protocol Version",
					Computed:            true,
				},
				"read_from_replicas": schema.BoolAttribute{
					MarkdownDescription: "Whether to read from replicas",
					Computed:            true,
				},
				"sentinel_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the Sentinel nodes",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"sentinel_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the Sentinel nodes",
					Computed:            true,
				},
				"service_name": schema.StringAttribute{
					MarkdownDescription: "Name of the monitored master (service) to query via the Sentinels",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Connection timeout to the database",
					Computed:            true,
				},
				"timeout_connection": schema.Int64Attribute{
					MarkdownDescription: "Timeout for establishing a connection to the server",
					Computed:            true,
				},
				"timeout_message": schema.Int64Attribute{
					MarkdownDescription: "Timeout for message processing",
					Computed:            true,
				},
				"timeout_request": schema.Int64Attribute{
					MarkdownDescription: "Timeout for requests to the server",
					Computed:            true,
				},
				"timeout_session": schema.Int64Attribute{
					MarkdownDescription: "Timeout for session",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "URL of the Redis server",
					Computed:            true,
				},
				"urls": schema.SetAttribute{
					MarkdownDescription: "URL(s) of the Redis server(s)",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Use TLS to connect to the store",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "data_retention",
		JMAPType:  "x:DataRetention",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures data retention policies, expunge schedules, and archival settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"archive_deleted_accounts_for": schema.Int64Attribute{
					MarkdownDescription: "How long to keep deleted accounts in the archive before they are permanently deleted. If null, deleted accounts will be permanently deleted immediately. Requires an Enterprise license.",
					Computed:            true,
				},
				"archive_deleted_items_for": schema.Int64Attribute{
					MarkdownDescription: "How long to keep deleted items in the archive before they are permanently deleted. If null, deleted items will not be archived and will be permanently deleted immediately. Requires an Enterprise license.",
					Computed:            true,
				},
				"blob_cleanup_schedule": schema.SingleNestedAttribute{
					MarkdownDescription: "How often to purge the data store. Expects a cron expression.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"day": schema.Int64Attribute{
							MarkdownDescription: "Day",
							Computed:            true,
						},
						"hour": schema.Int64Attribute{
							MarkdownDescription: "Hour",
							Computed:            true,
						},
						"minute": schema.Int64Attribute{
							MarkdownDescription: "Minute",
							Computed:            true,
						},
					},
				},
				"data_cleanup_schedule": schema.SingleNestedAttribute{
					MarkdownDescription: "How often to purge the data store. Expects a cron expression.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"day": schema.Int64Attribute{
							MarkdownDescription: "Day",
							Computed:            true,
						},
						"hour": schema.Int64Attribute{
							MarkdownDescription: "Hour",
							Computed:            true,
						},
						"minute": schema.Int64Attribute{
							MarkdownDescription: "Minute",
							Computed:            true,
						},
					},
				},
				"expunge_schedule": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies when to run the auto-expunge process",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"day": schema.Int64Attribute{
							MarkdownDescription: "Day",
							Computed:            true,
						},
						"hour": schema.Int64Attribute{
							MarkdownDescription: "Hour",
							Computed:            true,
						},
						"minute": schema.Int64Attribute{
							MarkdownDescription: "Minute",
							Computed:            true,
						},
					},
				},
				"expunge_scheduling_inbox_after": schema.Int64Attribute{
					MarkdownDescription: "Sets the duration after which the iTIP inbox will automatically expunge old messages.",
					Computed:            true,
				},
				"expunge_share_notify_after": schema.Int64Attribute{
					MarkdownDescription: "Specifies the duration for which the JMAP share notification history is retained before it is automatically purged.",
					Computed:            true,
				},
				"expunge_submissions_after": schema.Int64Attribute{
					MarkdownDescription: "How long to keep sent e-mail submissions before auto-expunging",
					Computed:            true,
				},
				"expunge_trash_after": schema.Int64Attribute{
					MarkdownDescription: "How long to keep messages in the Trash and Junk Mail folders before auto-expunging",
					Computed:            true,
				},
				"hold_metrics_for": schema.Int64Attribute{
					MarkdownDescription: "How long to keep metrics history before it is permanently deleted. Requires an Enterprise license.",
					Computed:            true,
				},
				"hold_mta_reports_for": schema.Int64Attribute{
					MarkdownDescription: "The duration for which MTA reports should be stored before being deleted, or None to disable storage",
					Computed:            true,
				},
				"hold_traces_for": schema.Int64Attribute{
					MarkdownDescription: "How long to keep message delivery history before it is permanently deleted. Requires an Enterprise license.",
					Computed:            true,
				},
				"max_changes_history": schema.Int64Attribute{
					MarkdownDescription: "How many changes to keep in the history for each account. This is used to determine the changes that have occurred since the last time the client requested changes.",
					Computed:            true,
				},
				"metrics_collection_interval": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies how often to collect metrics history. Requires an Enterprise license.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"day": schema.Int64Attribute{
							MarkdownDescription: "Day",
							Computed:            true,
						},
						"hour": schema.Int64Attribute{
							MarkdownDescription: "Hour",
							Computed:            true,
						},
						"minute": schema.Int64Attribute{
							MarkdownDescription: "Minute",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:      "data_store",
		JMAPType:  "x:DataStore",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the primary data store backend. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Allow invalid TLS certificates when connecting to the store",
					Computed:            true,
				},
				"auth_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the store",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"auth_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the store",
					Computed:            true,
				},
				"blob_size": schema.Int64Attribute{
					MarkdownDescription: "Minimum size of a blob to store in the blob store, smaller blobs are stored in the metadata store",
					Computed:            true,
				},
				"buffer_size": schema.Int64Attribute{
					MarkdownDescription: "Size of the write buffer in bytes, used to batch writes to the store",
					Computed:            true,
				},
				"cluster_file": schema.StringAttribute{
					MarkdownDescription: "Path to the cluster file for the FoundationDB cluster",
					Computed:            true,
				},
				"database": schema.StringAttribute{
					MarkdownDescription: "Name of the database",
					Computed:            true,
				},
				"datacenter_id": schema.StringAttribute{
					MarkdownDescription: "Data center ID (optional)",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "Hostname of the database server",
					Computed:            true,
				},
				"machine_id": schema.StringAttribute{
					MarkdownDescription: "Machine ID in the FoundationDB cluster (optional)",
					Computed:            true,
				},
				"max_allowed_packet": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a packet in bytes",
					Computed:            true,
				},
				"options": schema.StringAttribute{
					MarkdownDescription: "Additional connection options",
					Computed:            true,
				},
				"path": schema.StringAttribute{
					MarkdownDescription: "Path to the RocksDB data directory",
					Computed:            true,
				},
				"pool_max_connections": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of connections to the store",
					Computed:            true,
				},
				"pool_min_connections": schema.Int64Attribute{
					MarkdownDescription: "Minimum number of connections to the store",
					Computed:            true,
				},
				"pool_recycling_method": schema.StringAttribute{
					MarkdownDescription: "Method to use when recycling connections in the pool",
					Computed:            true,
				},
				"pool_workers": schema.Int64Attribute{
					MarkdownDescription: "Number of worker threads to use for the store, defaults to the number of cores",
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "Port of the database server",
					Computed:            true,
				},
				"read_replicas": schema.ListNestedAttribute{
					MarkdownDescription: "List of read replicas for the store Requires an Enterprise license.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"auth_secret": schema.SingleNestedAttribute{
								MarkdownDescription: "Password to connect to the store",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"auth_username": schema.StringAttribute{
								MarkdownDescription: "Username to connect to the store",
								Computed:            true,
							},
							"database": schema.StringAttribute{
								MarkdownDescription: "Name of the database",
								Computed:            true,
							},
							"host": schema.StringAttribute{
								MarkdownDescription: "Hostname of the database server",
								Computed:            true,
							},
							"options": schema.StringAttribute{
								MarkdownDescription: "Additional connection options",
								Computed:            true,
							},
							"port": schema.Int64Attribute{
								MarkdownDescription: "Port of the database server",
								Computed:            true,
							},
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Connection timeout to the database",
					Computed:            true,
				},
				"transaction_retry_delay": schema.Int64Attribute{
					MarkdownDescription: "Transaction maximum retry delay",
					Computed:            true,
				},
				"transaction_retry_limit": schema.Int64Attribute{
					MarkdownDescription: "Transaction retry limit",
					Computed:            true,
				},
				"transaction_timeout": schema.Int64Attribute{
					MarkdownDescription: "Transaction timeout",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Use TLS to connect to the store",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "directories",
		JMAPType: "x:Directory",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Directory` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "directory_ldap",
		JMAPType: "x:Directory",
		Variant:  "Ldap",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an external directory for account authentication and lookups. Reads the LDAP Directory variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Allow invalid TLS certificates when connecting to the server",
					Computed:            true,
				},
				"attr_class": schema.SetAttribute{
					MarkdownDescription: "LDAP attribute for the user's account type, if missing defaults to individual.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"attr_description": schema.SetAttribute{
					MarkdownDescription: "LDAP attributes used to store the user's description",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"attr_email": schema.SetAttribute{
					MarkdownDescription: "LDAP attribute for the user's primary email address",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"attr_email_alias": schema.SetAttribute{
					MarkdownDescription: "LDAP attribute for the user's email alias(es)",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"attr_member_of": schema.SetAttribute{
					MarkdownDescription: "LDAP attributes for the groups that a user belongs to. Used when filterMemberOf is not configured or when the group membership is also provided in the account entry.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"attr_secret": schema.SetAttribute{
					MarkdownDescription: "LDAP attribute for the user's password hash. This setting is required when binding as a service user. When using bind authentication, configure the secret-changed attribute instead.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"attr_secret_changed": schema.SetAttribute{
					MarkdownDescription: "LDAP attribute that provides a password change hash or a timestamp indicating when the password was last changed. When using bind authentication, this attribute is used to determine when to invalidate OAuth tokens.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"base_dn": schema.StringAttribute{
					MarkdownDescription: "The base distinguished name (DN) from where searches should begin",
					Computed:            true,
				},
				"bind_authentication": schema.BoolAttribute{
					MarkdownDescription: "Whether to use bind authentication. When enabled, the server will use the filterLogin to search for the user account and then attempt to bind as that account using the provided password. When disabled, the server will use the bind DN and secret to connect to the LDAP server and obtain the secret from the account entry using the attrSecret attribute.",
					Computed:            true,
				},
				"bind_dn": schema.StringAttribute{
					MarkdownDescription: "The distinguished name of the account that the server will bind as to connect to the LDAP directory",
					Computed:            true,
				},
				"bind_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The password or secret for the bind DN account",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this directory",
					Computed:            true,
				},
				"filter_login": schema.StringAttribute{
					MarkdownDescription: "Searches for user accounts by e-mail address during authentication",
					Computed:            true,
				},
				"filter_mailbox": schema.StringAttribute{
					MarkdownDescription: "Searches for users or groups matching a recipient e-mail address or alias",
					Computed:            true,
				},
				"filter_member_of": schema.StringAttribute{
					MarkdownDescription: "Searches for groups that an account is member of. Use when the group membership is not provided in the account entry. The ? in the filter will be replaced with the account DN.",
					Computed:            true,
				},
				"group_class": schema.StringAttribute{
					MarkdownDescription: "LDAP object class used to identify group entries",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this directory belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"pool_max_connections": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of connections that can be maintained simultaneously in the connection pool",
					Computed:            true,
				},
				"pool_timeout_create": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that the connection pool will wait for a new connection to be created",
					Computed:            true,
				},
				"pool_timeout_recycle": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that the connection pool manager will wait for a connection to be recycled",
					Computed:            true,
				},
				"pool_timeout_wait": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that the connection pool will wait for a connection to become available",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Connection timeout to the server",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "URL of the LDAP server",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Use TLS to connect to the remote server",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "directory_oidc",
		JMAPType: "x:Directory",
		Variant:  "Oidc",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an external directory for account authentication and lookups. Reads the OpenID Connect variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"claim_groups": schema.StringAttribute{
					MarkdownDescription: "The claim name used to retrieve the user's group memberships from the token or user info response. Common values are groups or roles depending on your provider's configuration. If not set, group information will not be populated. Note that some providers omit group claims from the token to keep its size small and only return them via the user info endpoint, if group information is missing, ensure your provider is configured to include it.",
					Computed:            true,
				},
				"claim_name": schema.StringAttribute{
					MarkdownDescription: "The claim name used to retrieve the user's display name from the token or user info response. Common values are name or display_name. If not set, the display name will not be populated.",
					Computed:            true,
				},
				"claim_username": schema.StringAttribute{
					MarkdownDescription: "The claim name used to retrieve the user's login name from the token or user info response. Common values are preferred_username, email, or sub depending on your provider's configuration. If the claim value is not an email address and usernameDomain is set, the domain will be appended automatically (e.g. john becomes john@example.com). If the claim value already contains an @, it is used as-is. If the claim value is not an email address and no usernameDomain is configured, Stalwart will fall back to the email claim. If neither yields a valid email address, authentication will be rejected.",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this directory",
					Computed:            true,
				},
				"issuer_url": schema.StringAttribute{
					MarkdownDescription: "The base URL of the OpenID Connect provider (e.g. https://sso.example.com/realms/myrealm). Stalwart will use this URL to automatically discover the provider's endpoints, including the token validation and user info endpoints.",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this directory belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"require_audience": schema.StringAttribute{
					MarkdownDescription: "If set, Stalwart will reject any token whose aud (audience) claim does not include this value. Set this to the client ID or resource identifier registered for Stalwart in your identity provider to ensure tokens issued for other applications are not accepted.",
					Computed:            true,
				},
				"require_scopes": schema.SetAttribute{
					MarkdownDescription: "If set, Stalwart will reject any token that does not include all of the specified scopes. Useful for ensuring that only tokens explicitly granted access to the mail server are accepted.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"username_domain": schema.StringAttribute{
					MarkdownDescription: "The domain name to append to the username when the value of claimUsername does not contain an @ symbol (e.g. setting this to example.com will turn john into john@example.com). If not set, Stalwart will fall back to the email claim when the username claim does not contain a valid email address.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "directory_sql",
		JMAPType: "x:Directory",
		Variant:  "Sql",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an external directory for account authentication and lookups. Reads the SQL Database variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"column_class": schema.StringAttribute{
					MarkdownDescription: "Column name for account type",
					Computed:            true,
				},
				"column_description": schema.StringAttribute{
					MarkdownDescription: "Column name for account full name or description",
					Computed:            true,
				},
				"column_email": schema.StringAttribute{
					MarkdownDescription: "Column name for e-mail address. Optional, you can use instead a query to obtain the account's addresses.",
					Computed:            true,
				},
				"column_secret": schema.StringAttribute{
					MarkdownDescription: "Column name for the account password.",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this directory",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this directory belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"query_email_aliases": schema.StringAttribute{
					MarkdownDescription: "Query to obtain the e-mail aliases of an account.",
					Computed:            true,
				},
				"query_login": schema.StringAttribute{
					MarkdownDescription: "Query to obtain the account details by login e-mail address.",
					Computed:            true,
				},
				"query_member_of": schema.StringAttribute{
					MarkdownDescription: "Query to obtain the groups an account is member of.",
					Computed:            true,
				},
				"query_recipient": schema.StringAttribute{
					MarkdownDescription: "Query to obtain the account details by recipient e-mail address or alias.",
					Computed:            true,
				},
				"store": schema.SingleNestedAttribute{
					MarkdownDescription: "Storage backend where accounts and groups are stored",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"allow_invalid_certs": schema.BoolAttribute{
							MarkdownDescription: "Allow invalid TLS certificates when connecting to the store",
							Computed:            true,
						},
						"auth_secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password to connect to the store",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"auth_username": schema.StringAttribute{
							MarkdownDescription: "Username to connect to the store",
							Computed:            true,
						},
						"database": schema.StringAttribute{
							MarkdownDescription: "Name of the database",
							Computed:            true,
						},
						"host": schema.StringAttribute{
							MarkdownDescription: "Hostname of the database server",
							Computed:            true,
						},
						"max_allowed_packet": schema.Int64Attribute{
							MarkdownDescription: "Maximum size of a packet in bytes",
							Computed:            true,
						},
						"options": schema.StringAttribute{
							MarkdownDescription: "Additional connection options",
							Computed:            true,
						},
						"path": schema.StringAttribute{
							MarkdownDescription: "Path to the SQLite data directory",
							Computed:            true,
						},
						"pool_max_connections": schema.Int64Attribute{
							MarkdownDescription: "Maximum number of connections to the store",
							Computed:            true,
						},
						"pool_min_connections": schema.Int64Attribute{
							MarkdownDescription: "Minimum number of connections to the store",
							Computed:            true,
						},
						"pool_recycling_method": schema.StringAttribute{
							MarkdownDescription: "Method to use when recycling connections in the pool",
							Computed:            true,
						},
						"pool_workers": schema.Int64Attribute{
							MarkdownDescription: "Number of worker threads to use for the store, defaults to the number of cores",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Port of the database server",
							Computed:            true,
						},
						"read_replicas": schema.ListNestedAttribute{
							MarkdownDescription: "List of read replicas for the store Requires an Enterprise license.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"auth_secret": schema.SingleNestedAttribute{
										MarkdownDescription: "Password to connect to the store",
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												MarkdownDescription: "Variant discriminator.",
												Computed:            true,
											},
											"file_path": schema.StringAttribute{
												MarkdownDescription: "File path to read the secret from",
												Computed:            true,
											},
											"secret": schema.StringAttribute{
												MarkdownDescription: "Password or secret value",
												Computed:            true,
												Sensitive:           true,
											},
											"variable_name": schema.StringAttribute{
												MarkdownDescription: "Environment variable name to read the secret from",
												Computed:            true,
											},
										},
									},
									"auth_username": schema.StringAttribute{
										MarkdownDescription: "Username to connect to the store",
										Computed:            true,
									},
									"database": schema.StringAttribute{
										MarkdownDescription: "Name of the database",
										Computed:            true,
									},
									"host": schema.StringAttribute{
										MarkdownDescription: "Hostname of the database server",
										Computed:            true,
									},
									"options": schema.StringAttribute{
										MarkdownDescription: "Additional connection options",
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: "Port of the database server",
										Computed:            true,
									},
								},
							},
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Connection timeout to the database",
							Computed:            true,
						},
						"use_tls": schema.BoolAttribute{
							MarkdownDescription: "Use TLS to connect to the store",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:      "dkim_report_settings",
		JMAPType:  "x:DkimReportSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures DKIM authentication failure report generation. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"dkim_sign_domain": schema.SingleNestedAttribute{
					MarkdownDescription: "Which domain's DKIM signatures to use when signing the DKIM report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"from_address": schema.SingleNestedAttribute{
					MarkdownDescription: "Email address that will be used in the From header of the DKIM report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"from_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Name that will be used in the From header of the DKIM report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"send_frequency": schema.SingleNestedAttribute{
					MarkdownDescription: "Rate at which DKIM reports will be sent to a given email address. When this rate is exceeded, no further DKIM failure reports will be sent to that address",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"subject": schema.SingleNestedAttribute{
					MarkdownDescription: "Subject name that will be used in the DKIM report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "dkim_signature_dkim1_ed25519_sha256",
		JMAPType: "x:DkimSignature",
		Variant:  "Dkim1Ed25519Sha256",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DKIM signature used to sign outgoing email messages. Reads the DKIM1 (Ed25519 SHA-256) variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auid": schema.StringAttribute{
					MarkdownDescription: "Agent or user identifier included in the DKIM signature header",
					Computed:            true,
				},
				"canonicalization": schema.StringAttribute{
					MarkdownDescription: "Canonicalization algorithm applied to the headers and body before signing",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the DKIM signature",
					Computed:            true,
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this DKIM signature is associated with",
					Computed:            true,
				},
				"expire": schema.Int64Attribute{
					MarkdownDescription: "Time after which this DKIM signature expires and should no longer be considered valid",
					Computed:            true,
				},
				"headers": schema.SetAttribute{
					MarkdownDescription: "List of message headers to include in the DKIM signature",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DKIM signature belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"next_transition_at": schema.StringAttribute{
					MarkdownDescription: "Date when this key will transition to the next rotation stage, or null if no transition is scheduled",
					Computed:            true,
				},
				"private_key": schema.SingleNestedAttribute{
					MarkdownDescription: "PEM-encoded private key used to sign outgoing messages",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"public_key": schema.StringAttribute{
					MarkdownDescription: "PEM-encoded public key used to verify signatures, derived from the private key",
					Computed:            true,
				},
				"report": schema.BoolAttribute{
					MarkdownDescription: "Whether to request failure reports when signature verification fails on the recipient side",
					Computed:            true,
				},
				"selector": schema.StringAttribute{
					MarkdownDescription: "Selector used to locate the DKIM public key in DNS",
					Computed:            true,
				},
				"stage": schema.StringAttribute{
					MarkdownDescription: "Current stage of the DKIM key in its rotation lifecycle",
					Computed:            true,
				},
				"third_party": schema.StringAttribute{
					MarkdownDescription: "Authorized third-party signature value, used when signing on behalf of another domain",
					Computed:            true,
				},
				"third_party_hash": schema.StringAttribute{
					MarkdownDescription: "Hashing algorithm used to verify the authorized third-party signature DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dkim_signature_dkim1_rsa_sha256",
		JMAPType: "x:DkimSignature",
		Variant:  "Dkim1RsaSha256",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DKIM signature used to sign outgoing email messages. Reads the DKIM1 (RSA SHA-256) variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auid": schema.StringAttribute{
					MarkdownDescription: "Agent or user identifier included in the DKIM signature header",
					Computed:            true,
				},
				"canonicalization": schema.StringAttribute{
					MarkdownDescription: "Canonicalization algorithm applied to the headers and body before signing",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the DKIM signature",
					Computed:            true,
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this DKIM signature is associated with",
					Computed:            true,
				},
				"expire": schema.Int64Attribute{
					MarkdownDescription: "Time after which this DKIM signature expires and should no longer be considered valid",
					Computed:            true,
				},
				"headers": schema.SetAttribute{
					MarkdownDescription: "List of message headers to include in the DKIM signature",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DKIM signature belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"next_transition_at": schema.StringAttribute{
					MarkdownDescription: "Date when this key will transition to the next rotation stage, or null if no transition is scheduled",
					Computed:            true,
				},
				"private_key": schema.SingleNestedAttribute{
					MarkdownDescription: "PEM-encoded private key used to sign outgoing messages",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"public_key": schema.StringAttribute{
					MarkdownDescription: "PEM-encoded public key used to verify signatures, derived from the private key",
					Computed:            true,
				},
				"report": schema.BoolAttribute{
					MarkdownDescription: "Whether to request failure reports when signature verification fails on the recipient side",
					Computed:            true,
				},
				"selector": schema.StringAttribute{
					MarkdownDescription: "Selector used to locate the DKIM public key in DNS",
					Computed:            true,
				},
				"stage": schema.StringAttribute{
					MarkdownDescription: "Current stage of the DKIM key in its rotation lifecycle",
					Computed:            true,
				},
				"third_party": schema.StringAttribute{
					MarkdownDescription: "Authorized third-party signature value, used when signing on behalf of another domain",
					Computed:            true,
				},
				"third_party_hash": schema.StringAttribute{
					MarkdownDescription: "Hashing algorithm used to verify the authorized third-party signature DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dkim_signature_dkim2_ed25519_sha256",
		JMAPType: "x:DkimSignature",
		Variant:  "Dkim2Ed25519Sha256",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DKIM signature used to sign outgoing email messages. Reads the DKIM2 (Ed25519 SHA-256) variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the DKIM signature",
					Computed:            true,
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this DKIM signature is associated with",
					Computed:            true,
				},
				"flags": schema.SetAttribute{
					MarkdownDescription: "Policy flags added to the signature, requesting downstream handlers to honor delivery constraints or provide feedback",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DKIM signature belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"next_transition_at": schema.StringAttribute{
					MarkdownDescription: "Date when this key will transition to the next rotation stage, or null if no transition is scheduled",
					Computed:            true,
				},
				"private_key": schema.SingleNestedAttribute{
					MarkdownDescription: "PEM-encoded private key used to sign outgoing messages",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"public_key": schema.StringAttribute{
					MarkdownDescription: "PEM-encoded public key used to verify signatures, derived from the private key",
					Computed:            true,
				},
				"selector": schema.StringAttribute{
					MarkdownDescription: "Selector used to locate the DKIM public key in DNS",
					Computed:            true,
				},
				"stage": schema.StringAttribute{
					MarkdownDescription: "Current stage of the DKIM key in its rotation lifecycle",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dkim_signature_dkim2_rsa_sha256",
		JMAPType: "x:DkimSignature",
		Variant:  "Dkim2RsaSha256",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DKIM signature used to sign outgoing email messages. Reads the DKIM2 (RSA SHA-256) variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the DKIM signature",
					Computed:            true,
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this DKIM signature is associated with",
					Computed:            true,
				},
				"flags": schema.SetAttribute{
					MarkdownDescription: "Policy flags added to the signature, requesting downstream handlers to honor delivery constraints or provide feedback",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DKIM signature belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"next_transition_at": schema.StringAttribute{
					MarkdownDescription: "Date when this key will transition to the next rotation stage, or null if no transition is scheduled",
					Computed:            true,
				},
				"private_key": schema.SingleNestedAttribute{
					MarkdownDescription: "PEM-encoded private key used to sign outgoing messages",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"public_key": schema.StringAttribute{
					MarkdownDescription: "PEM-encoded public key used to verify signatures, derived from the private key",
					Computed:            true,
				},
				"selector": schema.StringAttribute{
					MarkdownDescription: "Selector used to locate the DKIM public key in DNS",
					Computed:            true,
				},
				"stage": schema.StringAttribute{
					MarkdownDescription: "Current stage of the DKIM key in its rotation lifecycle",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dkim_signatures",
		JMAPType: "x:DkimSignature",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:DkimSignature` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "dmarc_external_report",
		JMAPType: "x:DmarcExternalReport",
		Schema: schema.Schema{
			MarkdownDescription: "Stores a DMARC aggregate report received from an external source.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "When the report is scheduled to be deleted",
					Computed:            true,
				},
				"from": schema.StringAttribute{
					MarkdownDescription: "Email address of the report sender",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this report belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"received_at": schema.StringAttribute{
					MarkdownDescription: "When the report email was received",
					Computed:            true,
				},
				"report": schema.SingleNestedAttribute{
					MarkdownDescription: "DMARC report content",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"date_range_begin": schema.StringAttribute{
							MarkdownDescription: "Start of the reporting period",
							Computed:            true,
						},
						"date_range_end": schema.StringAttribute{
							MarkdownDescription: "End of the reporting period",
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: "Contact email address of the reporting organization",
							Computed:            true,
						},
						"errors": schema.SetAttribute{
							MarkdownDescription: "Errors encountered during report generation",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"extensions": schema.ListNestedAttribute{
							MarkdownDescription: "Custom vendor-specific extensions to the report",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"definition": schema.StringAttribute{
										MarkdownDescription: "Extension content or value",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Extension identifier",
										Computed:            true,
									},
								},
							},
						},
						"extra_contact_info": schema.StringAttribute{
							MarkdownDescription: "Additional contact information for the reporting organization",
							Computed:            true,
						},
						"generator": schema.StringAttribute{
							MarkdownDescription: "Name and version of the software that generated the report",
							Computed:            true,
						},
						"org_name": schema.StringAttribute{
							MarkdownDescription: "Name of the organization that generated the report",
							Computed:            true,
						},
						"policy_adkim": schema.StringAttribute{
							MarkdownDescription: "DKIM alignment mode specified in the policy",
							Computed:            true,
						},
						"policy_aspf": schema.StringAttribute{
							MarkdownDescription: "SPF alignment mode specified in the policy",
							Computed:            true,
						},
						"policy_discovery_method": schema.StringAttribute{
							MarkdownDescription: "Method used to discover the DMARC policy record",
							Computed:            true,
						},
						"policy_disposition": schema.StringAttribute{
							MarkdownDescription: "Requested handling policy for failing messages",
							Computed:            true,
						},
						"policy_domain": schema.StringAttribute{
							MarkdownDescription: "Domain for which the DMARC policy is published",
							Computed:            true,
						},
						"policy_failure_reporting_options": schema.SetAttribute{
							MarkdownDescription: "Conditions under which failure reports should be generated",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"policy_np": schema.StringAttribute{
							MarkdownDescription: "Requested handling policy for failing messages from non-existent subdomains",
							Computed:            true,
						},
						"policy_subdomain_disposition": schema.StringAttribute{
							MarkdownDescription: "Requested handling policy for failing messages from subdomains",
							Computed:            true,
						},
						"policy_testing_mode": schema.BoolAttribute{
							MarkdownDescription: "Whether the policy is in testing mode",
							Computed:            true,
						},
						"policy_version": schema.StringAttribute{
							MarkdownDescription: "Version of the published DMARC policy",
							Computed:            true,
						},
						"records": schema.ListNestedAttribute{
							MarkdownDescription: "Aggregated authentication results grouped by source",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"count": schema.Int64Attribute{
										MarkdownDescription: "Number of messages from this source matching this result",
										Computed:            true,
									},
									"dkim_results": schema.ListNestedAttribute{
										MarkdownDescription: "DKIM authentication results for the messages",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"domain": schema.StringAttribute{
													MarkdownDescription: "Domain that signed the message",
													Computed:            true,
												},
												"human_result": schema.StringAttribute{
													MarkdownDescription: "Human-readable explanation of the result",
													Computed:            true,
												},
												"result": schema.StringAttribute{
													MarkdownDescription: "DKIM verification result",
													Computed:            true,
												},
												"selector": schema.StringAttribute{
													MarkdownDescription: "DKIM selector used for signing",
													Computed:            true,
												},
											},
										},
									},
									"envelope_from": schema.StringAttribute{
										MarkdownDescription: "Envelope sender domain (MAIL FROM)",
										Computed:            true,
									},
									"envelope_to": schema.StringAttribute{
										MarkdownDescription: "Envelope recipient domain",
										Computed:            true,
									},
									"evaluated_disposition": schema.StringAttribute{
										MarkdownDescription: "Action taken on the messages",
										Computed:            true,
									},
									"evaluated_dkim": schema.StringAttribute{
										MarkdownDescription: "DMARC result based on DKIM authentication",
										Computed:            true,
									},
									"evaluated_spf": schema.StringAttribute{
										MarkdownDescription: "DMARC result based on SPF authentication",
										Computed:            true,
									},
									"extensions": schema.ListNestedAttribute{
										MarkdownDescription: "Custom vendor-specific extensions to this record",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"definition": schema.StringAttribute{
													MarkdownDescription: "Extension content or value",
													Computed:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Extension identifier",
													Computed:            true,
												},
											},
										},
									},
									"header_from": schema.StringAttribute{
										MarkdownDescription: "Domain from the message From header",
										Computed:            true,
									},
									"policy_override_reasons": schema.ListNestedAttribute{
										MarkdownDescription: "Reasons why the evaluated disposition differs from the published policy",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"comment": schema.StringAttribute{
													MarkdownDescription: "Additional explanation for the override",
													Computed:            true,
												},
												"override_type": schema.StringAttribute{
													MarkdownDescription: "Type of policy override applied",
													Computed:            true,
												},
											},
										},
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: "IP address of the sending mail server",
										Computed:            true,
									},
									"spf_results": schema.ListNestedAttribute{
										MarkdownDescription: "SPF authentication results for the messages",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"domain": schema.StringAttribute{
													MarkdownDescription: "Domain checked for SPF",
													Computed:            true,
												},
												"human_result": schema.StringAttribute{
													MarkdownDescription: "Human-readable explanation of the result",
													Computed:            true,
												},
												"result": schema.StringAttribute{
													MarkdownDescription: "SPF verification result",
													Computed:            true,
												},
												"scope": schema.StringAttribute{
													MarkdownDescription: "Which identity was checked",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
						"report_id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for this report",
							Computed:            true,
						},
						"version": schema.Float64Attribute{
							MarkdownDescription: "DMARC report format version",
							Computed:            true,
						},
					},
				},
				"subject": schema.StringAttribute{
					MarkdownDescription: "Subject line of the report email",
					Computed:            true,
				},
				"to": schema.SetAttribute{
					MarkdownDescription: "List of recipient email addresses",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "dmarc_external_reports",
		JMAPType: "x:DmarcExternalReport",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:DmarcExternalReport` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "dmarc_internal_report",
		JMAPType: "x:DmarcInternalReport",
		Schema: schema.Schema{
			MarkdownDescription: "Stores an outbound DMARC aggregate report pending delivery.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "When the report was created",
					Computed:            true,
				},
				"deliver_at": schema.StringAttribute{
					MarkdownDescription: "When the report is scheduled to be delivered",
					Computed:            true,
				},
				"domain": schema.StringAttribute{
					MarkdownDescription: "Domain this report is associated with",
					Computed:            true,
				},
				"policy_identifier": schema.Int64Attribute{
					MarkdownDescription: "Identifier for the DMARC policy that generated this report",
					Computed:            true,
				},
				"report": schema.SingleNestedAttribute{
					MarkdownDescription: "DMARC report content",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"date_range_begin": schema.StringAttribute{
							MarkdownDescription: "Start of the reporting period",
							Computed:            true,
						},
						"date_range_end": schema.StringAttribute{
							MarkdownDescription: "End of the reporting period",
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: "Contact email address of the reporting organization",
							Computed:            true,
						},
						"errors": schema.SetAttribute{
							MarkdownDescription: "Errors encountered during report generation",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"extensions": schema.ListNestedAttribute{
							MarkdownDescription: "Custom vendor-specific extensions to the report",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"definition": schema.StringAttribute{
										MarkdownDescription: "Extension content or value",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Extension identifier",
										Computed:            true,
									},
								},
							},
						},
						"extra_contact_info": schema.StringAttribute{
							MarkdownDescription: "Additional contact information for the reporting organization",
							Computed:            true,
						},
						"generator": schema.StringAttribute{
							MarkdownDescription: "Name and version of the software that generated the report",
							Computed:            true,
						},
						"org_name": schema.StringAttribute{
							MarkdownDescription: "Name of the organization that generated the report",
							Computed:            true,
						},
						"policy_adkim": schema.StringAttribute{
							MarkdownDescription: "DKIM alignment mode specified in the policy",
							Computed:            true,
						},
						"policy_aspf": schema.StringAttribute{
							MarkdownDescription: "SPF alignment mode specified in the policy",
							Computed:            true,
						},
						"policy_discovery_method": schema.StringAttribute{
							MarkdownDescription: "Method used to discover the DMARC policy record",
							Computed:            true,
						},
						"policy_disposition": schema.StringAttribute{
							MarkdownDescription: "Requested handling policy for failing messages",
							Computed:            true,
						},
						"policy_domain": schema.StringAttribute{
							MarkdownDescription: "Domain for which the DMARC policy is published",
							Computed:            true,
						},
						"policy_failure_reporting_options": schema.SetAttribute{
							MarkdownDescription: "Conditions under which failure reports should be generated",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"policy_np": schema.StringAttribute{
							MarkdownDescription: "Requested handling policy for failing messages from non-existent subdomains",
							Computed:            true,
						},
						"policy_subdomain_disposition": schema.StringAttribute{
							MarkdownDescription: "Requested handling policy for failing messages from subdomains",
							Computed:            true,
						},
						"policy_testing_mode": schema.BoolAttribute{
							MarkdownDescription: "Whether the policy is in testing mode",
							Computed:            true,
						},
						"policy_version": schema.StringAttribute{
							MarkdownDescription: "Version of the published DMARC policy",
							Computed:            true,
						},
						"records": schema.ListNestedAttribute{
							MarkdownDescription: "Aggregated authentication results grouped by source",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"count": schema.Int64Attribute{
										MarkdownDescription: "Number of messages from this source matching this result",
										Computed:            true,
									},
									"dkim_results": schema.ListNestedAttribute{
										MarkdownDescription: "DKIM authentication results for the messages",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"domain": schema.StringAttribute{
													MarkdownDescription: "Domain that signed the message",
													Computed:            true,
												},
												"human_result": schema.StringAttribute{
													MarkdownDescription: "Human-readable explanation of the result",
													Computed:            true,
												},
												"result": schema.StringAttribute{
													MarkdownDescription: "DKIM verification result",
													Computed:            true,
												},
												"selector": schema.StringAttribute{
													MarkdownDescription: "DKIM selector used for signing",
													Computed:            true,
												},
											},
										},
									},
									"envelope_from": schema.StringAttribute{
										MarkdownDescription: "Envelope sender domain (MAIL FROM)",
										Computed:            true,
									},
									"envelope_to": schema.StringAttribute{
										MarkdownDescription: "Envelope recipient domain",
										Computed:            true,
									},
									"evaluated_disposition": schema.StringAttribute{
										MarkdownDescription: "Action taken on the messages",
										Computed:            true,
									},
									"evaluated_dkim": schema.StringAttribute{
										MarkdownDescription: "DMARC result based on DKIM authentication",
										Computed:            true,
									},
									"evaluated_spf": schema.StringAttribute{
										MarkdownDescription: "DMARC result based on SPF authentication",
										Computed:            true,
									},
									"extensions": schema.ListNestedAttribute{
										MarkdownDescription: "Custom vendor-specific extensions to this record",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"definition": schema.StringAttribute{
													MarkdownDescription: "Extension content or value",
													Computed:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Extension identifier",
													Computed:            true,
												},
											},
										},
									},
									"header_from": schema.StringAttribute{
										MarkdownDescription: "Domain from the message From header",
										Computed:            true,
									},
									"policy_override_reasons": schema.ListNestedAttribute{
										MarkdownDescription: "Reasons why the evaluated disposition differs from the published policy",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"comment": schema.StringAttribute{
													MarkdownDescription: "Additional explanation for the override",
													Computed:            true,
												},
												"override_type": schema.StringAttribute{
													MarkdownDescription: "Type of policy override applied",
													Computed:            true,
												},
											},
										},
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: "IP address of the sending mail server",
										Computed:            true,
									},
									"spf_results": schema.ListNestedAttribute{
										MarkdownDescription: "SPF authentication results for the messages",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"domain": schema.StringAttribute{
													MarkdownDescription: "Domain checked for SPF",
													Computed:            true,
												},
												"human_result": schema.StringAttribute{
													MarkdownDescription: "Human-readable explanation of the result",
													Computed:            true,
												},
												"result": schema.StringAttribute{
													MarkdownDescription: "SPF verification result",
													Computed:            true,
												},
												"scope": schema.StringAttribute{
													MarkdownDescription: "Which identity was checked",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
						"report_id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for this report",
							Computed:            true,
						},
						"version": schema.Float64Attribute{
							MarkdownDescription: "DMARC report format version",
							Computed:            true,
						},
					},
				},
				"rua": schema.SetAttribute{
					MarkdownDescription: "Reporting email addresses from the DMARC policy",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "dmarc_internal_reports",
		JMAPType: "x:DmarcInternalReport",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:DmarcInternalReport` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "dmarc_report_settings",
		JMAPType:  "x:DmarcReportSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures DMARC aggregate and failure report generation. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"aggregate_contact_info": schema.SingleNestedAttribute{
					MarkdownDescription: "Contact information to be included in the report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"aggregate_dkim_sign_domain": schema.SingleNestedAttribute{
					MarkdownDescription: "Which domain's DKIM signatures to use when signing the DMARC aggregate report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"aggregate_from_address": schema.SingleNestedAttribute{
					MarkdownDescription: "Email address that will be used in the From header of the DMARC aggregate report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"aggregate_from_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Name that will be used in the From header of the DMARC aggregate report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"aggregate_max_report_size": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum size of the DMARC aggregate report in bytes",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"aggregate_org_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Name of the organization to be included in the report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"aggregate_send_frequency": schema.SingleNestedAttribute{
					MarkdownDescription: "Frequency at which the DMARC aggregate reports will be sent. The options are hourly, daily, weekly, or disable to disable reporting",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"aggregate_subject": schema.SingleNestedAttribute{
					MarkdownDescription: "Subject name that will be used in the DMARC aggregate report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"failure_dkim_sign_domain": schema.SingleNestedAttribute{
					MarkdownDescription: "Which domain's DKIM signatures to use when signing the DMARC authentication failure report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"failure_from_address": schema.SingleNestedAttribute{
					MarkdownDescription: "Email address that will be used in the From header of the DMARC authentication failure report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"failure_from_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Name that will be used in the From header of the DMARC report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"failure_send_frequency": schema.SingleNestedAttribute{
					MarkdownDescription: "Rate at which DMARC reports will be sent to a given email address. When this rate is exceeded, no further DMARC failure reports will be sent to that address",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"failure_subject": schema.SingleNestedAttribute{
					MarkdownDescription: "Subject name that will be used in the DMARC authentication failure report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "dns_resolver",
		JMAPType:  "x:DnsResolver",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the DNS resolver used for domain lookups. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"attempts": schema.Int64Attribute{
					MarkdownDescription: "Number of times a resolution request will be retried before it is considered failed",
					Computed:            true,
				},
				"concurrency": schema.Int64Attribute{
					MarkdownDescription: "Number of concurrent resolution requests that can be made at the same time",
					Computed:            true,
				},
				"enable_edns": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable EDNS (Extension Mechanisms for DNS) support",
					Computed:            true,
				},
				"preserve_intermediates": schema.BoolAttribute{
					MarkdownDescription: "Whether to preserve the intermediate name servers in the DNS resolution results",
					Computed:            true,
				},
				"servers": schema.ListNestedAttribute{
					MarkdownDescription: "List of custom DNS server URLs to use for resolution",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								MarkdownDescription: "Address of the DNS server",
								Computed:            true,
							},
							"port": schema.Int64Attribute{
								MarkdownDescription: "Port of the DNS server",
								Computed:            true,
							},
							"protocol": schema.StringAttribute{
								MarkdownDescription: "Protocol to use for DNS queries",
								Computed:            true,
							},
						},
					},
				},
				"tcp_on_error": schema.BoolAttribute{
					MarkdownDescription: "Whether to try using TCP for resolution requests if an error occurs during a UDP resolution request",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Time after which a resolution request will be timed out if no response is received",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Whether to use TLS for DNS resolution",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_alidns",
		JMAPType: "x:DnsServer",
		Variant:  "Alidns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Alibaba Cloud DNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"access_key": schema.StringAttribute{
					MarkdownDescription: "The Alibaba Cloud access key ID",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"line": schema.StringAttribute{
					MarkdownDescription: "Optional ISP line identifier (used for split-resolution accounts)",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"region": schema.StringAttribute{
					MarkdownDescription: "Optional regional endpoint (defaults to the global endpoint)",
					Computed:            true,
				},
				"secret_key": schema.SingleNestedAttribute{
					MarkdownDescription: "The Alibaba Cloud access key secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"security_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Optional STS security token for temporary credentials",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_arvan_cloud",
		JMAPType: "x:DnsServer",
		Variant:  "ArvanCloud",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the ArvanCloud variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_autodns",
		JMAPType: "x:DnsServer",
		Variant:  "Autodns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the InterNetX AutoDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"context": schema.Int64Attribute{
					MarkdownDescription: "Optional account context identifier",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"password": schema.SingleNestedAttribute{
					MarkdownDescription: "AutoDNS account password",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "AutoDNS account username",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_azure_dns",
		JMAPType: "x:DnsServer",
		Variant:  "AzureDns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Microsoft Azure DNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"client_id": schema.StringAttribute{
					MarkdownDescription: "Application (client) ID",
					Computed:            true,
				},
				"client_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Application client secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"environment": schema.StringAttribute{
					MarkdownDescription: "Azure cloud environment",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"resource_group": schema.StringAttribute{
					MarkdownDescription: "Resource group that contains the DNS zone",
					Computed:            true,
				},
				"subscription_id": schema.StringAttribute{
					MarkdownDescription: "Azure subscription ID that owns the DNS zone",
					Computed:            true,
				},
				"tenant_id": schema.StringAttribute{
					MarkdownDescription: "Azure Active Directory tenant ID",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_baidu_cloud",
		JMAPType: "x:DnsServer",
		Variant:  "BaiduCloud",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Baidu Cloud DNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"access_key": schema.StringAttribute{
					MarkdownDescription: "Baidu Cloud access key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Baidu Cloud secret key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_bluecat_v2",
		JMAPType: "x:DnsServer",
		Variant:  "BluecatV2",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the BlueCat Address Manager variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"base_url": schema.StringAttribute{
					MarkdownDescription: "Base URL of the BlueCat Address Manager",
					Computed:            true,
				},
				"config_name": schema.StringAttribute{
					MarkdownDescription: "BlueCat configuration name",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"password": schema.SingleNestedAttribute{
					MarkdownDescription: "BlueCat account password",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"skip_deploy": schema.BoolAttribute{
					MarkdownDescription: "Skip deploying changes after applying them",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "BlueCat account username",
					Computed:            true,
				},
				"view_name": schema.StringAttribute{
					MarkdownDescription: "BlueCat DNS view name",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_bunny",
		JMAPType: "x:DnsServer",
		Variant:  "Bunny",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the BunnyDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_clou_dns",
		JMAPType: "x:DnsServer",
		Variant:  "ClouDns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the ClouDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auth_id": schema.StringAttribute{
					MarkdownDescription: "ClouDNS auth ID (use either auth-id or sub-auth-id)",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"password": schema.SingleNestedAttribute{
					MarkdownDescription: "ClouDNS auth password",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"sub_auth_id": schema.StringAttribute{
					MarkdownDescription: "ClouDNS sub-auth ID",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_cloudflare",
		JMAPType: "x:DnsServer",
		Variant:  "Cloudflare",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Cloudflare variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"email": schema.StringAttribute{
					MarkdownDescription: "Optional account email to authenticate with Cloudflare",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_constellix",
		JMAPType: "x:DnsServer",
		Variant:  "Constellix",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Constellix variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.StringAttribute{
					MarkdownDescription: "Constellix API key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Constellix secret key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_cpanel",
		JMAPType: "x:DnsServer",
		Variant:  "Cpanel",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the cPanel variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"base_url": schema.StringAttribute{
					MarkdownDescription: "Base URL of the cPanel server (e.g. https://host:2083)",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"token": schema.SingleNestedAttribute{
					MarkdownDescription: "cPanel API token",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "cPanel account username",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_ddnss",
		JMAPType: "x:DnsServer",
		Variant:  "Ddnss",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the DDNSS.de variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_de_sec",
		JMAPType: "x:DnsServer",
		Variant:  "DeSEC",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the DeSEC variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_digital_ocean",
		JMAPType: "x:DnsServer",
		Variant:  "DigitalOcean",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the DigitalOcean variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_dns_made_easy",
		JMAPType: "x:DnsServer",
		Variant:  "DnsMadeEasy",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the DNS Made Easy variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.StringAttribute{
					MarkdownDescription: "DNS Made Easy API key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "DNS Made Easy API secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_dnsimple",
		JMAPType: "x:DnsServer",
		Variant:  "Dnsimple",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the DNSimple variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"account_identifier": schema.StringAttribute{
					MarkdownDescription: "The account ID used to authenticate with DNSimple",
					Computed:            true,
				},
				"auth_token": schema.SingleNestedAttribute{
					MarkdownDescription: "The authentication token used to authenticate with DNSimple",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Deprecated secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_domeneshop",
		JMAPType: "x:DnsServer",
		Variant:  "Domeneshop",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Domeneshop variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auth_token": schema.StringAttribute{
					MarkdownDescription: "Domeneshop API token",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Domeneshop API secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_dreamhost",
		JMAPType: "x:DnsServer",
		Variant:  "Dreamhost",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Dreamhost variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_duck_dns",
		JMAPType: "x:DnsServer",
		Variant:  "DuckDns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the DuckDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_dynu",
		JMAPType: "x:DnsServer",
		Variant:  "Dynu",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Dynu variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_easy_dns",
		JMAPType: "x:DnsServer",
		Variant:  "EasyDns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the EasyDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"key": schema.SingleNestedAttribute{
					MarkdownDescription: "EasyDNS key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"token": schema.StringAttribute{
					MarkdownDescription: "EasyDNS token",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_edge_dns",
		JMAPType: "x:DnsServer",
		Variant:  "EdgeDns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Akamai EdgeDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"access_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Akamai access token",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"account_switch_key": schema.StringAttribute{
					MarkdownDescription: "Optional account switch key for managing multiple accounts",
					Computed:            true,
				},
				"client_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Akamai client secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"client_token": schema.StringAttribute{
					MarkdownDescription: "Akamai client token",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "Akamai API host",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_exoscale",
		JMAPType: "x:DnsServer",
		Variant:  "Exoscale",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Exoscale variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.StringAttribute{
					MarkdownDescription: "Exoscale API key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Exoscale API secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_free_my_ip",
		JMAPType: "x:DnsServer",
		Variant:  "FreeMyIp",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the freemyip.com variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_gandi_v5",
		JMAPType: "x:DnsServer",
		Variant:  "GandiV5",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Gandi LiveDNS v5 variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_gcore",
		JMAPType: "x:DnsServer",
		Variant:  "Gcore",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Gcore variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_glesys",
		JMAPType: "x:DnsServer",
		Variant:  "Glesys",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the GleSYS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.SingleNestedAttribute{
					MarkdownDescription: "GleSYS API key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"api_user": schema.StringAttribute{
					MarkdownDescription: "GleSYS API user",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_godaddy",
		JMAPType: "x:DnsServer",
		Variant:  "Godaddy",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the GoDaddy variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.StringAttribute{
					MarkdownDescription: "GoDaddy API key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "GoDaddy API secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_google_cloud_dns",
		JMAPType: "x:DnsServer",
		Variant:  "GoogleCloudDns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Google Cloud DNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"impersonate_service_account": schema.StringAttribute{
					MarkdownDescription: "Optional service account email to impersonate",
					Computed:            true,
				},
				"managed_zone": schema.StringAttribute{
					MarkdownDescription: "Managed zone name (resolved automatically by longest suffix match if not set)",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"private_zone": schema.BoolAttribute{
					MarkdownDescription: "Whether to restrict zone resolution to private zones only",
					Computed:            true,
				},
				"project_id": schema.StringAttribute{
					MarkdownDescription: "The Google Cloud project ID that owns the managed zone",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"service_account_json": schema.SingleNestedAttribute{
					MarkdownDescription: "Service account JSON credentials used to authenticate with Google Cloud",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_hetzner",
		JMAPType: "x:DnsServer",
		Variant:  "Hetzner",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Hetzner variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_hosting_de",
		JMAPType: "x:DnsServer",
		Variant:  "HostingDe",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the hosting.de variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_hostinger",
		JMAPType: "x:DnsServer",
		Variant:  "Hostinger",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Hostinger variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_huawei_cloud",
		JMAPType: "x:DnsServer",
		Variant:  "HuaweiCloud",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Huawei Cloud DNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"access_key": schema.StringAttribute{
					MarkdownDescription: "Huawei Cloud access key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"region": schema.StringAttribute{
					MarkdownDescription: "Huawei Cloud region",
					Computed:            true,
				},
				"secret_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Huawei Cloud secret key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_hurricane",
		JMAPType: "x:DnsServer",
		Variant:  "Hurricane",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Hurricane Electric variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"credentials": schema.ListNestedAttribute{
					MarkdownDescription: "Per-zone Hurricane Electric DDNS keys",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"secret": schema.SingleNestedAttribute{
								MarkdownDescription: "DDNS key for the zone",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"zone": schema.StringAttribute{
								MarkdownDescription: "DNS zone (origin) the credential applies to",
								Computed:            true,
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_ibm_cloud",
		JMAPType: "x:DnsServer",
		Variant:  "IbmCloud",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the IBM Cloud variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.SingleNestedAttribute{
					MarkdownDescription: "IBM Cloud API key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "IBM Cloud account username",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_infoblox",
		JMAPType: "x:DnsServer",
		Variant:  "Infoblox",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Infoblox NIOS WAPI variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"dns_view": schema.StringAttribute{
					MarkdownDescription: "DNS view name (defaults to External)",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "Infoblox grid master host",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"password": schema.SingleNestedAttribute{
					MarkdownDescription: "Infoblox account password",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"port": schema.StringAttribute{
					MarkdownDescription: "Optional port (defaults to 443)",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "Infoblox account username",
					Computed:            true,
				},
				"wapi_version": schema.StringAttribute{
					MarkdownDescription: "WAPI version to use (defaults to 2.11)",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_infomaniak",
		JMAPType: "x:DnsServer",
		Variant:  "Infomaniak",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Infomaniak variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_inwx",
		JMAPType: "x:DnsServer",
		Variant:  "Inwx",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the INWX variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"password": schema.SingleNestedAttribute{
					MarkdownDescription: "INWX account password",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"sandbox": schema.BoolAttribute{
					MarkdownDescription: "Use the INWX sandbox API instead of production",
					Computed:            true,
				},
				"shared_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Optional shared secret for TOTP-based two-factor authentication",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "INWX account username",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_ionos",
		JMAPType: "x:DnsServer",
		Variant:  "Ionos",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the IONOS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_ipv64",
		JMAPType: "x:DnsServer",
		Variant:  "Ipv64",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the IPv64 variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_joker",
		JMAPType: "x:DnsServer",
		Variant:  "Joker",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Joker variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auth": schema.SingleNestedAttribute{
					MarkdownDescription: "Joker authentication method",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"api_key": schema.SingleNestedAttribute{
							MarkdownDescription: "Joker DMAPI API key",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"password": schema.SingleNestedAttribute{
							MarkdownDescription: "Joker DMAPI account password",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Joker DMAPI account username",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_lightsail",
		JMAPType: "x:DnsServer",
		Variant:  "Lightsail",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the AWS Lightsail variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"access_key_id": schema.StringAttribute{
					MarkdownDescription: "AWS access key ID",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"domain": schema.StringAttribute{
					MarkdownDescription: "Optional Lightsail domain name to scope record operations to",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"region": schema.StringAttribute{
					MarkdownDescription: "AWS region (defaults to us-east-1)",
					Computed:            true,
				},
				"secret_access_key": schema.SingleNestedAttribute{
					MarkdownDescription: "AWS secret access key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"session_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Optional session token for temporary AWS credentials",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_linode",
		JMAPType: "x:DnsServer",
		Variant:  "Linode",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Linode variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_lua_dns",
		JMAPType: "x:DnsServer",
		Variant:  "LuaDns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the LuaDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auth_token": schema.SingleNestedAttribute{
					MarkdownDescription: "LuaDNS API token",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "LuaDNS account email or username",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_mythic_beasts",
		JMAPType: "x:DnsServer",
		Variant:  "MythicBeasts",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Mythic Beasts variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"password": schema.SingleNestedAttribute{
					MarkdownDescription: "Mythic Beasts API key secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "Mythic Beasts API key ID",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_name_dot_com",
		JMAPType: "x:DnsServer",
		Variant:  "NameDotCom",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Name.com variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auth_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Name.com API token",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "Name.com account username",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_name_silo",
		JMAPType: "x:DnsServer",
		Variant:  "NameSilo",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the NameSilo variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_namecheap",
		JMAPType: "x:DnsServer",
		Variant:  "Namecheap",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Namecheap variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Namecheap API key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"api_user": schema.StringAttribute{
					MarkdownDescription: "Namecheap API user",
					Computed:            true,
				},
				"client_ip": schema.StringAttribute{
					MarkdownDescription: "Whitelisted client IP address registered with Namecheap",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "Optional account username (defaults to the API user)",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_netcup",
		JMAPType: "x:DnsServer",
		Variant:  "Netcup",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Netcup variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.StringAttribute{
					MarkdownDescription: "Netcup API key",
					Computed:            true,
				},
				"customer_number": schema.StringAttribute{
					MarkdownDescription: "Netcup customer number",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"password": schema.SingleNestedAttribute{
					MarkdownDescription: "Netcup API password",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_netlify",
		JMAPType: "x:DnsServer",
		Variant:  "Netlify",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Netlify variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_nifcloud",
		JMAPType: "x:DnsServer",
		Variant:  "Nifcloud",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Nifcloud variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"access_key": schema.StringAttribute{
					MarkdownDescription: "Nifcloud access key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Nifcloud secret key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_ns1",
		JMAPType: "x:DnsServer",
		Variant:  "Ns1",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the NS1 variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_oracle_cloud",
		JMAPType: "x:DnsServer",
		Variant:  "OracleCloud",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Oracle Cloud variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"compartment_ocid": schema.StringAttribute{
					MarkdownDescription: "Compartment OCID that owns the DNS zone",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"fingerprint": schema.StringAttribute{
					MarkdownDescription: "API signing key fingerprint",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"private_key_password": schema.SingleNestedAttribute{
					MarkdownDescription: "Optional passphrase for the private key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"private_key_pem": schema.SingleNestedAttribute{
					MarkdownDescription: "API signing private key in PEM format",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"region": schema.StringAttribute{
					MarkdownDescription: "OCI region (e.g. us-ashburn-1)",
					Computed:            true,
				},
				"tenancy_ocid": schema.StringAttribute{
					MarkdownDescription: "Tenancy OCID",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"user_ocid": schema.StringAttribute{
					MarkdownDescription: "User OCID",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_ovh",
		JMAPType: "x:DnsServer",
		Variant:  "Ovh",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the OVH variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"application_key": schema.StringAttribute{
					MarkdownDescription: "The application key used to authenticate with the OVH DNS server",
					Computed:            true,
				},
				"application_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The application secret used to authenticate with the OVH DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"consumer_key": schema.SingleNestedAttribute{
					MarkdownDescription: "The consumer key used to authenticate with the OVH DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"ovh_endpoint": schema.StringAttribute{
					MarkdownDescription: "Which OVH endpoint to use",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_plesk",
		JMAPType: "x:DnsServer",
		Variant:  "Plesk",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Plesk variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Plesk API key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"base_url": schema.StringAttribute{
					MarkdownDescription: "Base URL of the Plesk server (e.g. https://host:8443)",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_porkbun",
		JMAPType: "x:DnsServer",
		Variant:  "Porkbun",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Porkbun variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.StringAttribute{
					MarkdownDescription: "The API key used to authenticate with Porkbun",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Deprecated secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"secret_api_key": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret API key used to authenticate with Porkbun",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_route53",
		JMAPType: "x:DnsServer",
		Variant:  "Route53",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the AWS Route53 variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"access_key_id": schema.StringAttribute{
					MarkdownDescription: "The AWS access key ID",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"hosted_zone_id": schema.StringAttribute{
					MarkdownDescription: "Hosted zone ID to use (resolved automatically by name if not set)",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"private_zone_only": schema.BoolAttribute{
					MarkdownDescription: "Whether to restrict zone resolution to private zones only",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"region": schema.StringAttribute{
					MarkdownDescription: "The AWS region",
					Computed:            true,
				},
				"secret_access_key": schema.SingleNestedAttribute{
					MarkdownDescription: "The AWS secret access key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"session_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Optional session token for temporary AWS credentials",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_safedns",
		JMAPType: "x:DnsServer",
		Variant:  "Safedns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the ANS SafeDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_scaleway",
		JMAPType: "x:DnsServer",
		Variant:  "Scaleway",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Scaleway variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_spaceship",
		JMAPType: "x:DnsServer",
		Variant:  "Spaceship",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Spaceship variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.StringAttribute{
					MarkdownDescription: "The API key used to authenticate with Spaceship",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_tencent_cloud",
		JMAPType: "x:DnsServer",
		Variant:  "TencentCloud",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Tencent Cloud DNSPod variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"region": schema.StringAttribute{
					MarkdownDescription: "Optional regional endpoint",
					Computed:            true,
				},
				"secret_id": schema.StringAttribute{
					MarkdownDescription: "Tencent Cloud secret ID",
					Computed:            true,
				},
				"secret_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Tencent Cloud secret key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"session_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Optional STS session token for temporary credentials",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_transip",
		JMAPType: "x:DnsServer",
		Variant:  "Transip",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the TransIP variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"private_key_pem": schema.SingleNestedAttribute{
					MarkdownDescription: "TransIP private key in PEM format",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "TransIP account login",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_tsig",
		JMAPType: "x:DnsServer",
		Variant:  "Tsig",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the RFC2136 (TSIG) variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "The IP address of the DNS server",
					Computed:            true,
				},
				"key": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"key_name": schema.StringAttribute{
					MarkdownDescription: "The key used to authenticate with the DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "The port used to communicate with the DNS server",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"protocol": schema.StringAttribute{
					MarkdownDescription: "The protocol used to communicate with the DNS server",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"tsig_algorithm": schema.StringAttribute{
					MarkdownDescription: "The TSIG algorithm used to authenticate with the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_ultra_dns",
		JMAPType: "x:DnsServer",
		Variant:  "UltraDns",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the UltraDNS variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"endpoint": schema.StringAttribute{
					MarkdownDescription: "Optional REST API endpoint override",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"password": schema.SingleNestedAttribute{
					MarkdownDescription: "UltraDNS account password",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
				"username": schema.StringAttribute{
					MarkdownDescription: "UltraDNS account username",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_vercel",
		JMAPType: "x:DnsServer",
		Variant:  "Vercel",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Vercel variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"auth_token": schema.SingleNestedAttribute{
					MarkdownDescription: "Vercel auth token",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"team_id": schema.StringAttribute{
					MarkdownDescription: "Optional team ID to scope API requests to",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_volcengine",
		JMAPType: "x:DnsServer",
		Variant:  "Volcengine",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Volcano Engine variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"access_key": schema.StringAttribute{
					MarkdownDescription: "Volcengine access key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "Optional API host override",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"region": schema.StringAttribute{
					MarkdownDescription: "Optional regional endpoint",
					Computed:            true,
				},
				"scheme": schema.StringAttribute{
					MarkdownDescription: "Optional URL scheme (http or https)",
					Computed:            true,
				},
				"secret_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Volcengine secret key",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_vultr",
		JMAPType: "x:DnsServer",
		Variant:  "Vultr",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Vultr variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret or token used to authenticate with the DNS server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_web_support",
		JMAPType: "x:DnsServer",
		Variant:  "WebSupport",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the WebSupport variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.StringAttribute{
					MarkdownDescription: "WebSupport API key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"secret": schema.SingleNestedAttribute{
					MarkdownDescription: "WebSupport API secret",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_server_yandex_cloud",
		JMAPType: "x:DnsServer",
		Variant:  "YandexCloud",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNS server for automatic record management. Reads the Yandex Cloud variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"api_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Base64-encoded IAM service account key JSON",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of this DNS server",
					Computed:            true,
				},
				"folder_id": schema.StringAttribute{
					MarkdownDescription: "Yandex Cloud folder ID that owns the DNS zone",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this DNS server belongs to",
					Computed:            true,
				},
				"polling_interval": schema.Int64Attribute{
					MarkdownDescription: "How often to check for DNS records to propagate",
					Computed:            true,
				},
				"propagation_delay": schema.Int64Attribute{
					MarkdownDescription: "Initial delay before first propagation check (useful for slow providers)",
					Computed:            true,
				},
				"propagation_timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for DNS records to propagate",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Request timeout for the DNS server",
					Computed:            true,
				},
				"ttl": schema.Int64Attribute{
					MarkdownDescription: "The TTL for new DNS record",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "dns_servers",
		JMAPType: "x:DnsServer",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:DnsServer` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "domain",
		JMAPType: "x:Domain",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines an email domain and its DNS, DKIM, and TLS certificate settings.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"aliases": schema.SetAttribute{
					MarkdownDescription: "List of additional domain names that are aliases of this domain",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"allow_relaying": schema.BoolAttribute{
					MarkdownDescription: "Whether to allow relaying for non-local recipients, useful in split delivery scenarios",
					Computed:            true,
				},
				"catch_all_address": schema.StringAttribute{
					MarkdownDescription: "Catch-all email address that receives messages addressed to unknown local recipients",
					Computed:            true,
				},
				"certificate_management": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether TLS certificates for this domain are managed manually or automatically by an ACME provider",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"acme_provider_id": schema.StringAttribute{
							MarkdownDescription: "Identifier for the ACME provider managing certificates for this domain",
							Computed:            true,
						},
						"subject_alternative_names": schema.SetAttribute{
							MarkdownDescription: "Additional hostnames to include in the certificate as Subject Alternative Names (SANs).\nEnter hostnames only (e.g. `mta-sts`, `autoconfig`), the domain is appended automatically.\nTo include the apex domain, enter it in full (e.g. `example.org`).\nLeave empty to request a wildcard certificate when possible, or to use the default SANs.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the domain",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the domain",
					Computed:            true,
				},
				"directory_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the directory where accounts for this domain are stored, or null to use the internal directory Requires an Enterprise license.",
					Computed:            true,
				},
				"dkim_management": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether DKIM keys for this domain are managed manually or automatically by the server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"algorithms": schema.SetAttribute{
							MarkdownDescription: "List of signing algorithms to use when generating new DKIM keys",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"delete_after": schema.Int64Attribute{
							MarkdownDescription: "How long to retain old DKIM keys on the server after rotation before deleting them permanently. Requires automatic DNS management.",
							Computed:            true,
						},
						"retire_after": schema.Int64Attribute{
							MarkdownDescription: "How long to keep the old key's DNS record published after rotation before removing it. Requires automatic DNS management.",
							Computed:            true,
						},
						"rotate_after": schema.Int64Attribute{
							MarkdownDescription: "How often to rotate DKIM keys. Requires automatic DNS management to be enabled for the domain.",
							Computed:            true,
						},
						"selector_template": schema.StringAttribute{
							MarkdownDescription: "Template for generating DKIM selectors during key rotation. Supported variables:\n- `{algorithm}`: signing algorithm in lowercase (`rsa`, `ed25519`)\n- `{hash}`: hash algorithm (`sha256`)\n- `{version}`: DKIM version number (`1`)\n- `{date-<fmt>}`: current UTC date formatted with chrono strftime (e.g. `{date-%Y%m%d}`)\n- `{epoch}`: current UTC unix timestamp\n- `{random}`: random 8-character alphanumeric string",
							Computed:            true,
						},
					},
				},
				"dns_management": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether DNS records for this domain are managed manually or automatically by a DNS provider",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"dns_server_id": schema.StringAttribute{
							MarkdownDescription: "Identifier for the DNS server provider managing DNS records for this domain",
							Computed:            true,
						},
						"origin": schema.StringAttribute{
							MarkdownDescription: "Origin domain used to determine the correct DNS zone for managing records. For example, if the domain is \"sub.example.com\" and DNS records should be managed in the \"example.com\" zone, set the origin to \"example.com\". Leave empty to use the domain name itself as the zone origin.",
							Computed:            true,
						},
						"publish_records": schema.SetAttribute{
							MarkdownDescription: "Which DNS record types should be automatically published and kept in sync",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"dns_zone_file": schema.StringAttribute{
					MarkdownDescription: "Current DNS zone data for the domain",
					Computed:            true,
				},
				"is_enabled": schema.BoolAttribute{
					MarkdownDescription: "Whether this domain is enabled",
					Computed:            true,
				},
				"logo": schema.StringAttribute{
					MarkdownDescription: "URL or base64-encoded image representing the domain Requires an Enterprise license.",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this domain belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Domain name",
					Optional:            true,
					Computed:            true,
				},
				"report_address_uri": schema.StringAttribute{
					MarkdownDescription: "Email address to receive DMARC, TLS-RPT and CAA reports for this domain, or null to not receive reports",
					Computed:            true,
				},
				"sub_addressing": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether sub-addressing (plus addressing) is enabled for the domain",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"custom_rule": schema.SingleNestedAttribute{
							MarkdownDescription: "Expression that defines custom sub-addressing rules for the domain",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"else": schema.StringAttribute{
									MarkdownDescription: "Else condition",
									Computed:            true,
								},
								"match": schema.ListNestedAttribute{
									MarkdownDescription: "List of conditions and their corresponding results",
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"if": schema.StringAttribute{
												MarkdownDescription: "If condition",
												Computed:            true,
											},
											"then": schema.StringAttribute{
												MarkdownDescription: "Then clause",
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "domains",
		JMAPType: "x:Domain",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Domain` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "dsn_report_settings",
		JMAPType:  "x:DsnReportSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures Delivery Status Notification (DSN) report generation. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"dkim_sign_domain": schema.SingleNestedAttribute{
					MarkdownDescription: "Which domain's DKIM signatures to use when signing the Delivery Status Notifications",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"from_address": schema.SingleNestedAttribute{
					MarkdownDescription: "Email address that will be used in the From header of Delivery Status Notifications (DSN) reports",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"from_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Name that will be used in the From header of Delivery Status Notifications (DSN) reports",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "email",
		JMAPType:  "x:Email",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures email message limits, encryption, compression, and default folder settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"compression_algorithm": schema.StringAttribute{
					MarkdownDescription: "Algorithm to use to compress e-mail data",
					Computed:            true,
				},
				"default_folders": schema.MapAttribute{
					MarkdownDescription: "Default special-use folders to create for new users",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"encrypt_at_rest": schema.BoolAttribute{
					MarkdownDescription: "Allow users to configure encryption at rest for their data",
					Computed:            true,
				},
				"encrypt_on_append": schema.BoolAttribute{
					MarkdownDescription: "Encrypt messages that are manually appended by the user using JMAP or IMAP",
					Computed:            true,
				},
				"max_attachment_size": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum size for an email attachment",
					Computed:            true,
				},
				"max_identities": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of identities a user can create",
					Computed:            true,
				},
				"max_mailbox_depth": schema.Int64Attribute{
					MarkdownDescription: "Restricts the maximum depth of nested mailboxes a user can create",
					Computed:            true,
				},
				"max_mailbox_name_length": schema.Int64Attribute{
					MarkdownDescription: "Establishes the maximum length of a mailbox name",
					Computed:            true,
				},
				"max_mailboxes": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of mailboxes a user can create",
					Computed:            true,
				},
				"max_masked_addresses": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of masked email addresses a user can create Requires an Enterprise license.",
					Computed:            true,
				},
				"max_message_size": schema.Int64Attribute{
					MarkdownDescription: "Determines the maximum size for an email message",
					Computed:            true,
				},
				"max_messages": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of emails a user can create",
					Computed:            true,
				},
				"max_public_keys": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of encryption-at-rest public keys a user can create",
					Computed:            true,
				},
				"max_submissions": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of email submissions a user can create",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "enterprise",
		JMAPType:  "x:Enterprise",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures enterprise licensing and branding settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"api_key": schema.SingleNestedAttribute{
					MarkdownDescription: "API key for license retrieval and automatic renewals. Obtain your API key at https://license.stalw.art.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"license_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Upgrade to the enterprise version of Stalwart by entering your license key here. Obtain your license at https://license.stalw.art/buy",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"logo_url": schema.StringAttribute{
					MarkdownDescription: "URL to the default logo to use in the Webadmin interface. Requires an Enterprise license.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "event_tracing_level",
		JMAPType: "x:EventTracingLevel",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a custom logging level override for a specific event type.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"event": schema.StringAttribute{
					MarkdownDescription: "Unique identifier of the event",
					Computed:            true,
				},
				"level": schema.StringAttribute{
					MarkdownDescription: "The logging level for this event",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "event_tracing_levels",
		JMAPType: "x:EventTracingLevel",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:EventTracingLevel` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "file_storage",
		JMAPType:  "x:FileStorage",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures file storage limits. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"max_files": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of files a user can create",
					Computed:            true,
				},
				"max_folders": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of file folders a user can create",
					Computed:            true,
				},
				"max_size": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum size of a file that can be uploaded to the server",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "group",
		JMAPType: "x:Account",
		Variant:  "Group",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a user or group account for authentication and email access. Reads the Group account variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"aliases": schema.ListNestedAttribute{
					MarkdownDescription: "List of email aliases for the group",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Description of the email alias",
								Computed:            true,
							},
							"domain_id": schema.StringAttribute{
								MarkdownDescription: "Identifier for the domain of the email alias (the part after the @ symbol).",
								Computed:            true,
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Whether this email alias is enabled",
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "The local part of the email alias (the part before the @ symbol)",
								Computed:            true,
							},
						},
					},
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the account",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the group",
					Computed:            true,
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this group belongs to. This is used to determine the email address of the group, which is formed as name@domain.",
					Computed:            true,
				},
				"email_address": schema.StringAttribute{
					MarkdownDescription: "Email address of the group, formed as name@domain.",
					Computed:            true,
				},
				"locale": schema.StringAttribute{
					MarkdownDescription: "Preferred locale for the group",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this group belongs to",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the group, typically an email address local part.",
					Optional:            true,
					Computed:            true,
				},
				"permissions": schema.SingleNestedAttribute{
					MarkdownDescription: "Permissions assigned to this group",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"disabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"enabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly enabled.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"quotas": schema.MapAttribute{
					MarkdownDescription: "Quotas for different object types within this group",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"roles": schema.SingleNestedAttribute{
					MarkdownDescription: "Roles assigned to this group",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"role_ids": schema.SetAttribute{
							MarkdownDescription: "List of roles assigned to this principal.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"time_zone": schema.StringAttribute{
					MarkdownDescription: "Preferred time zone for the account",
					Computed:            true,
				},
				"used_disk_quota": schema.Int64Attribute{
					MarkdownDescription: "Amount of disk space currently used by this account (bytes)",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "http",
		JMAPType:  "x:Http",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures HTTP server settings including rate limiting, CORS, and security headers. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"allowed_endpoints": schema.SingleNestedAttribute{
					MarkdownDescription: "An expression that determines whether access to an endpoint is allowed. The expression should an HTTP status code (200, 403, etc.)",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"enable_hsts": schema.BoolAttribute{
					MarkdownDescription: "Specifies whether to enable HTTP Strict Transport Security for the HTTP server.",
					Computed:            true,
				},
				"rate_limit_anonymous": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies the request rate limit for unauthenticated users",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
				"rate_limit_authenticated": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies the request rate limit for authenticated users",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
				"redirect_root": schema.StringAttribute{
					MarkdownDescription: "The URL to redirect users to when they access the root path of the HTTP server. If not set, the server will return a 404 Not Found response.",
					Computed:            true,
				},
				"response_headers": schema.MapAttribute{
					MarkdownDescription: "Additional headers to include in HTTP responses",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"use_permissive_cors": schema.BoolAttribute{
					MarkdownDescription: "Specifies whether to allow all origins in the CORS policy for the HTTP server",
					Computed:            true,
				},
				"use_x_forwarded": schema.BoolAttribute{
					MarkdownDescription: "Specifies whether to use the Forwarded or X-Forwarded-For header to determine the client's IP address",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "http_form",
		JMAPType:  "x:HttpForm",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the contact form submission endpoint. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"default_from_address": schema.StringAttribute{
					MarkdownDescription: "The default e-mail address to use when the sender does not provide one.",
					Computed:            true,
				},
				"default_name": schema.StringAttribute{
					MarkdownDescription: "The default name to use when the sender does not provide one.",
					Computed:            true,
				},
				"default_subject": schema.StringAttribute{
					MarkdownDescription: "The default subject to use when the sender does not provide one.",
					Computed:            true,
				},
				"deliver_to": schema.SetAttribute{
					MarkdownDescription: "List of local e-mail addresses to deliver the contact form to.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable contact form submissions.",
					Computed:            true,
				},
				"field_email": schema.StringAttribute{
					MarkdownDescription: "The name of the field in the contact form that contains the e-mail address of the sender.",
					Computed:            true,
				},
				"field_honey_pot": schema.StringAttribute{
					MarkdownDescription: "The name of the field in the contact form that is used as a honey pot to catch spam bots.",
					Computed:            true,
				},
				"field_name": schema.StringAttribute{
					MarkdownDescription: "The name of the field in the contact form that contains the name of the sender.",
					Computed:            true,
				},
				"field_subject": schema.StringAttribute{
					MarkdownDescription: "The name of the field in the contact form that contains the subject of the message.",
					Computed:            true,
				},
				"max_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the contact form submission in bytes.",
					Computed:            true,
				},
				"rate_limit": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum number of contact form submissions that can be made in a timeframe by a given IP address.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
				"validate_domain": schema.BoolAttribute{
					MarkdownDescription: "Whether to validate the domain of the sender's email address.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "http_lookup",
		JMAPType: "x:HttpLookup",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an HTTP-based lookup list.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this HTTP list",
					Computed:            true,
				},
				"format": schema.SingleNestedAttribute{
					MarkdownDescription: "Format of the list",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"index_key": schema.Int64Attribute{
							MarkdownDescription: "The position of the key field in the HTTP List.",
							Computed:            true,
						},
						"index_value": schema.Int64Attribute{
							MarkdownDescription: "The position of the value field in the HTTP List.",
							Computed:            true,
						},
						"separator": schema.StringAttribute{
							MarkdownDescription: "The separator character used to parse the HTTP list.",
							Computed:            true,
						},
						"skip_first": schema.BoolAttribute{
							MarkdownDescription: "Whether to skip the first line of the list",
							Computed:            true,
						},
					},
				},
				"is_gzipped": schema.BoolAttribute{
					MarkdownDescription: "Whether to use gzip compression when downloading the list",
					Computed:            true,
				},
				"max_entries": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of entries allowed in the list. The list is truncated if it exceeds this limit.",
					Computed:            true,
				},
				"max_entry_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum length of an entry in the list.",
					Computed:            true,
				},
				"max_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of the list. The list is truncated if it exceeds this size.",
					Computed:            true,
				},
				"namespace": schema.StringAttribute{
					MarkdownDescription: "Unique identifier for this store when used in lookups",
					Computed:            true,
				},
				"refresh": schema.Int64Attribute{
					MarkdownDescription: "How often to refresh the list",
					Computed:            true,
				},
				"retry": schema.Int64Attribute{
					MarkdownDescription: "How long to wait before retrying to download the list in case of failure.",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "How long to wait for the list to download before timing out",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "URL of the list",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "http_lookups",
		JMAPType: "x:HttpLookup",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:HttpLookup` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "imap",
		JMAPType:  "x:Imap",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures IMAP protocol settings including authentication, timeouts, and rate limits. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"allow_plain_text_auth": schema.BoolAttribute{
					MarkdownDescription: "Whether to allow plain text authentication on unencrypted connections",
					Computed:            true,
				},
				"max_auth_failures": schema.Int64Attribute{
					MarkdownDescription: "Number of authentication attempts a user can make before being disconnected by the server",
					Computed:            true,
				},
				"max_concurrent": schema.Int64Attribute{
					MarkdownDescription: "The maximum number of concurrent connections",
					Computed:            true,
				},
				"max_messages_per_command": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of messages processed by a single FETCH, SEARCH, STORE, MOVE or UID EXPUNGE command, announced to clients through the MESSAGELIMIT capability. Lowering this truncates results for clients that do not implement RFC 9738",
					Computed:            true,
				},
				"max_messages_per_save": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of messages accepted by a single COPY or APPEND command, announced to clients through the SAVELIMIT capability. These commands are atomic, so exceeding the limit stores nothing; keep it at or above the MESSAGELIMIT value",
					Computed:            true,
				},
				"max_request_rate": schema.SingleNestedAttribute{
					MarkdownDescription: "The maximum number of requests per minute",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
				"max_request_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of an IMAP request that the server will accept",
					Computed:            true,
				},
				"max_uid_batches": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of UID ranges returned by a single UIDBATCHES command; wider requests are rejected with a TOOMANY response code",
					Computed:            true,
				},
				"min_uid_batch_size": schema.Int64Attribute{
					MarkdownDescription: "Smallest batch size a client may ask for in a UIDBATCHES command; smaller requests are rejected with a TOOFEW response code. Cannot exceed 500, which is the batch size every server is required to support",
					Computed:            true,
				},
				"timeout_anonymous": schema.Int64Attribute{
					MarkdownDescription: "Time an unauthenticated session can stay inactive before being ended by the server",
					Computed:            true,
				},
				"timeout_authenticated": schema.Int64Attribute{
					MarkdownDescription: "Time an authenticated session can remain idle before the server terminates it",
					Computed:            true,
				},
				"timeout_idle": schema.Int64Attribute{
					MarkdownDescription: "Time a connection can stay idle in the IMAP IDLE state before the server breaks the connection",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "in_memory_store",
		JMAPType:  "x:InMemoryStore",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the in-memory cache and lookup store. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"auth_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the store",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"auth_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the store",
					Computed:            true,
				},
				"max_retries": schema.Int64Attribute{
					MarkdownDescription: "Number of retries to connect to the Redis cluster",
					Computed:            true,
				},
				"max_retry_wait": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait between retries",
					Computed:            true,
				},
				"min_retry_wait": schema.Int64Attribute{
					MarkdownDescription: "Minimum time to wait between retries",
					Computed:            true,
				},
				"pool_max_connections": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of connections to the store",
					Computed:            true,
				},
				"pool_timeout_create": schema.Int64Attribute{
					MarkdownDescription: "Timeout for creating a new connection",
					Computed:            true,
				},
				"pool_timeout_recycle": schema.Int64Attribute{
					MarkdownDescription: "Timeout for recycling a connection",
					Computed:            true,
				},
				"pool_timeout_wait": schema.Int64Attribute{
					MarkdownDescription: "Timeout for waiting for a connection from the pool",
					Computed:            true,
				},
				"protocol_version": schema.StringAttribute{
					MarkdownDescription: "Protocol Version",
					Computed:            true,
				},
				"read_from_replicas": schema.BoolAttribute{
					MarkdownDescription: "Whether to read from replicas",
					Computed:            true,
				},
				"sentinel_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the Sentinel nodes",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"sentinel_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the Sentinel nodes",
					Computed:            true,
				},
				"service_name": schema.StringAttribute{
					MarkdownDescription: "Name of the monitored master (service) to query via the Sentinels",
					Computed:            true,
				},
				"stores": schema.ListNestedAttribute{
					MarkdownDescription: "Stores to use for sharding",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								MarkdownDescription: "Variant discriminator.",
								Computed:            true,
							},
							"auth_secret": schema.SingleNestedAttribute{
								MarkdownDescription: "Password to connect to the store",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"auth_username": schema.StringAttribute{
								MarkdownDescription: "Username to connect to the store",
								Computed:            true,
							},
							"max_retries": schema.Int64Attribute{
								MarkdownDescription: "Number of retries to connect to the Redis cluster",
								Computed:            true,
							},
							"max_retry_wait": schema.Int64Attribute{
								MarkdownDescription: "Maximum time to wait between retries",
								Computed:            true,
							},
							"min_retry_wait": schema.Int64Attribute{
								MarkdownDescription: "Minimum time to wait between retries",
								Computed:            true,
							},
							"pool_max_connections": schema.Int64Attribute{
								MarkdownDescription: "Maximum number of connections to the store",
								Computed:            true,
							},
							"pool_timeout_create": schema.Int64Attribute{
								MarkdownDescription: "Timeout for creating a new connection",
								Computed:            true,
							},
							"pool_timeout_recycle": schema.Int64Attribute{
								MarkdownDescription: "Timeout for recycling a connection",
								Computed:            true,
							},
							"pool_timeout_wait": schema.Int64Attribute{
								MarkdownDescription: "Timeout for waiting for a connection from the pool",
								Computed:            true,
							},
							"protocol_version": schema.StringAttribute{
								MarkdownDescription: "Protocol Version",
								Computed:            true,
							},
							"read_from_replicas": schema.BoolAttribute{
								MarkdownDescription: "Whether to read from replicas",
								Computed:            true,
							},
							"sentinel_secret": schema.SingleNestedAttribute{
								MarkdownDescription: "Password to connect to the Sentinel nodes",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"sentinel_username": schema.StringAttribute{
								MarkdownDescription: "Username to connect to the Sentinel nodes",
								Computed:            true,
							},
							"service_name": schema.StringAttribute{
								MarkdownDescription: "Name of the monitored master (service) to query via the Sentinels",
								Computed:            true,
							},
							"timeout": schema.Int64Attribute{
								MarkdownDescription: "Connection timeout to the database",
								Computed:            true,
							},
							"url": schema.StringAttribute{
								MarkdownDescription: "URL of the Redis server",
								Computed:            true,
							},
							"urls": schema.SetAttribute{
								MarkdownDescription: "URL(s) of the Redis server(s)",
								Computed:            true,
								ElementType:         types.StringType,
							},
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Connection timeout to the database",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "URL of the Redis server",
					Computed:            true,
				},
				"urls": schema.SetAttribute{
					MarkdownDescription: "URL(s) of the Redis server(s)",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "jmap",
		JMAPType:  "x:Jmap",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures JMAP protocol limits for requests, uploads, and push notifications. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"changes_max_results": schema.Int64Attribute{
					MarkdownDescription: "Determines the maximum number of change objects that a Changes method can return",
					Computed:            true,
				},
				"event_source_throttle": schema.Int64Attribute{
					MarkdownDescription: "Specifies the minimum time between two event source notifications",
					Computed:            true,
				},
				"get_max_results": schema.Int64Attribute{
					MarkdownDescription: "Determines the maximum number of objects that can be fetched in a single method call",
					Computed:            true,
				},
				"max_concurrent_requests": schema.Int64Attribute{
					MarkdownDescription: "Restricts the number of concurrent requests a user can make to the JMAP server",
					Computed:            true,
				},
				"max_concurrent_uploads": schema.Int64Attribute{
					MarkdownDescription: "Restricts the number of concurrent file uploads a user can perform",
					Computed:            true,
				},
				"max_method_calls": schema.Int64Attribute{
					MarkdownDescription: "Limits the maximum number of method calls that can be included in a single request",
					Computed:            true,
				},
				"max_push_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size in bytes for a Web Push notification payload; EmailPush objects are truncated to fit within this limit",
					Computed:            true,
				},
				"max_request_size": schema.Int64Attribute{
					MarkdownDescription: "Defines the maximum size of a single request, in bytes, that the server will accept",
					Computed:            true,
				},
				"max_subscriptions": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of push subscriptions a user can create",
					Computed:            true,
				},
				"max_upload_count": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum number of files that a user can upload within a certain period",
					Computed:            true,
				},
				"max_upload_size": schema.Int64Attribute{
					MarkdownDescription: "Defines the maximum file size for file uploads to the server",
					Computed:            true,
				},
				"parse_limit_contact": schema.Int64Attribute{
					MarkdownDescription: "Limits the maximum number of vCard items that can be parsed in a single request",
					Computed:            true,
				},
				"parse_limit_email": schema.Int64Attribute{
					MarkdownDescription: "Limits the maximum number of e-mail message that can be parsed in a single request",
					Computed:            true,
				},
				"parse_limit_event": schema.Int64Attribute{
					MarkdownDescription: "Limits the maximum number of iCalendar items that can be parsed in a single request",
					Computed:            true,
				},
				"push_attempt_wait": schema.Int64Attribute{
					MarkdownDescription: "Time to wait between push attempts",
					Computed:            true,
				},
				"push_max_attempts": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of push attempts before a notification is discarded",
					Computed:            true,
				},
				"push_request_timeout": schema.Int64Attribute{
					MarkdownDescription: "Time before a connection with a push service URL times out",
					Computed:            true,
				},
				"push_retry_wait": schema.Int64Attribute{
					MarkdownDescription: "Time to wait between retry attempts",
					Computed:            true,
				},
				"push_shards_total": schema.Int64Attribute{
					MarkdownDescription: "Total number of shards for push notification processing across multiple nodes",
					Computed:            true,
				},
				"push_throttle": schema.Int64Attribute{
					MarkdownDescription: "Time to wait before sending a new request to the push service",
					Computed:            true,
				},
				"push_verify_timeout": schema.Int64Attribute{
					MarkdownDescription: "Time to wait for the push service to verify a subscription",
					Computed:            true,
				},
				"query_max_results": schema.Int64Attribute{
					MarkdownDescription: "Sets the maximum number of results that a Query method can return",
					Computed:            true,
				},
				"set_max_objects": schema.Int64Attribute{
					MarkdownDescription: "Establishes the maximum number of objects that can be modified in a single method call",
					Computed:            true,
				},
				"snippet_max_results": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of search snippets to return in a single request",
					Computed:            true,
				},
				"upload_quota": schema.Int64Attribute{
					MarkdownDescription: "Defines the total size of files that a user can upload within a certain period",
					Computed:            true,
				},
				"upload_ttl": schema.Int64Attribute{
					MarkdownDescription: "Specifies the Time-To-Live (TTL) for each uploaded file, after which the file is deleted from temporary storage",
					Computed:            true,
				},
				"web_push_contact": schema.StringAttribute{
					MarkdownDescription: "Optional contact URI (a mailto: or https: address) included as the sub claim of VAPID tokens so that push services can reach the server operator. Leave empty to omit the claim.",
					Computed:            true,
				},
				"web_push_key": schema.SingleNestedAttribute{
					MarkdownDescription: "ECDSA P-256 private key (PKCS#8 PEM) used to sign VAPID (RFC 8292) authentication tokens for Web Push.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"websocket_heartbeat": schema.Int64Attribute{
					MarkdownDescription: "Time to wait before sending a new heartbeat to the WebSocket client",
					Computed:            true,
				},
				"websocket_throttle": schema.Int64Attribute{
					MarkdownDescription: "Amount of time to wait before sending a batch of notifications to a WS client",
					Computed:            true,
				},
				"websocket_timeout": schema.Int64Attribute{
					MarkdownDescription: "Time before an inactive WebSocket connection times out",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "log",
		JMAPType: "x:Log",
		Schema: schema.Schema{
			MarkdownDescription: "Represents a server log entry.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"details": schema.StringAttribute{
					MarkdownDescription: "Log message",
					Computed:            true,
				},
				"event": schema.StringAttribute{
					MarkdownDescription: "Event type of the log entry",
					Computed:            true,
				},
				"level": schema.StringAttribute{
					MarkdownDescription: "Severity level of the log entry",
					Computed:            true,
				},
				"timestamp": schema.StringAttribute{
					MarkdownDescription: "Timestamp of the log entry",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "logs",
		JMAPType: "x:Log",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Log` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "mailing_list",
		JMAPType: "x:MailingList",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a mailing list that distributes messages to a group of recipients.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"aliases": schema.ListNestedAttribute{
					MarkdownDescription: "List of email aliases for the mailing list",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Description of the email alias",
								Computed:            true,
							},
							"domain_id": schema.StringAttribute{
								MarkdownDescription: "Identifier for the domain of the email alias (the part after the @ symbol).",
								Computed:            true,
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Whether this email alias is enabled",
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "The local part of the email alias (the part before the @ symbol)",
								Computed:            true,
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the mailing list",
					Computed:            true,
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this mailing list belongs to. This is used to determine the email address of the mailing list, which is formed as name@domain.",
					Computed:            true,
				},
				"email_address": schema.StringAttribute{
					MarkdownDescription: "The email address of the mailing list, formed as name@domain.",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this mailing list belongs to",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the mailing list, typically an email address local part.",
					Optional:            true,
					Computed:            true,
				},
				"recipients": schema.SetAttribute{
					MarkdownDescription: "List of email addresses that are members of the mailing list",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "mailing_lists",
		JMAPType: "x:MailingList",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MailingList` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "masked_email",
		JMAPType: "x:MaskedEmail",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a masked email address for privacy protection. Requires an Enterprise license.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"account_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the account this masked email address belongs to",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "The date-time the email address was created.",
					Computed:            true,
				},
				"created_by": schema.StringAttribute{
					MarkdownDescription: "The name of the client that created this email address. This will be set by the server automatically based on the credentials used to authenticate the request, e.g. \"ACME Password Manager\".",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the masked email address",
					Computed:            true,
				},
				"email": schema.StringAttribute{
					MarkdownDescription: "The masked email address",
					Computed:            true,
				},
				"email_domain": schema.StringAttribute{
					MarkdownDescription: "This is only used on create and otherwise ignored; if supplied, the server-assigned email will end with the given domain.",
					Computed:            true,
				},
				"email_prefix": schema.StringAttribute{
					MarkdownDescription: "This is only used on create and otherwise ignored; if supplied, the server-assigned email will start with the given prefix. The string MUST be <= 64 characters in length and MUST only contain characters a-z, 0-9 and _ (underscore).",
					Computed:            true,
				},
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Whether this masked email address is enabled",
					Computed:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "Expiration date of the email address",
					Computed:            true,
				},
				"for_domain": schema.StringAttribute{
					MarkdownDescription: "The domain name of the site this address was created for, e.g. \"https://example.com\". This is intended to be added automatically by password managers.",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "A URL pointing back to the integrator's use of this email address, e.g. a custom-uri to open \"ACME Password Manager\" at the appropriate entry.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "masked_emails",
		JMAPType: "x:MaskedEmail",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MaskedEmail` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "memory_lookup_key",
		JMAPType: "x:MemoryLookupKey",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an in-memory lookup key for fast data access.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"is_glob_pattern": schema.BoolAttribute{
					MarkdownDescription: "Whether the key is a glob pattern",
					Computed:            true,
				},
				"key": schema.StringAttribute{
					MarkdownDescription: "The key name",
					Computed:            true,
				},
				"namespace": schema.StringAttribute{
					MarkdownDescription: "The namespace of the key",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "memory_lookup_key_value",
		JMAPType: "x:MemoryLookupKeyValue",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an in-memory lookup key-value pair.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"is_glob_pattern": schema.BoolAttribute{
					MarkdownDescription: "Whether the key is a glob pattern",
					Computed:            true,
				},
				"key": schema.StringAttribute{
					MarkdownDescription: "The key name",
					Computed:            true,
				},
				"namespace": schema.StringAttribute{
					MarkdownDescription: "The namespace of the key",
					Computed:            true,
				},
				"value": schema.StringAttribute{
					MarkdownDescription: "The key value",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "memory_lookup_key_values",
		JMAPType: "x:MemoryLookupKeyValue",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MemoryLookupKeyValue` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "memory_lookup_keys",
		JMAPType: "x:MemoryLookupKey",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MemoryLookupKey` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "metric",
		JMAPType: "x:Metric",
		Schema: schema.Schema{
			MarkdownDescription: "Stores a collected server metric data point.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"count_value": schema.Int64Attribute{
					MarkdownDescription: "Count associated with the metric",
					Computed:            true,
				},
				"metric": schema.StringAttribute{
					MarkdownDescription: "Metric event type",
					Computed:            true,
				},
				"sum": schema.Int64Attribute{
					MarkdownDescription: "Sum associated with the metric",
					Computed:            true,
				},
				"timestamp": schema.StringAttribute{
					MarkdownDescription: "Timestamp of the metric entry",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "metric_data_points",
		JMAPType: "x:Metric",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Metric` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "metrics",
		JMAPType:  "x:Metrics",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures metrics collection and export via OpenTelemetry and Prometheus. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"metrics": schema.SetAttribute{
					MarkdownDescription: "List of metrics to include or exclude based on filter mode",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"metrics_policy": schema.StringAttribute{
					MarkdownDescription: "How to interpret the metrics list",
					Computed:            true,
				},
				"open_telemetry": schema.SingleNestedAttribute{
					MarkdownDescription: "OpenTelemetry metrics configuration",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"endpoint": schema.StringAttribute{
							MarkdownDescription: "The endpoint for Open Telemetry",
							Computed:            true,
						},
						"http_auth": schema.SingleNestedAttribute{
							MarkdownDescription: "The type of HTTP authentication to use",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"bearer_token": schema.SingleNestedAttribute{
									MarkdownDescription: "Bearer token for HTTP Bearer Authentication",
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											MarkdownDescription: "Variant discriminator.",
											Computed:            true,
										},
										"file_path": schema.StringAttribute{
											MarkdownDescription: "File path to read the secret from",
											Computed:            true,
										},
										"secret": schema.StringAttribute{
											MarkdownDescription: "Password or secret value",
											Computed:            true,
											Sensitive:           true,
										},
										"variable_name": schema.StringAttribute{
											MarkdownDescription: "Environment variable name to read the secret from",
											Computed:            true,
										},
									},
								},
								"secret": schema.SingleNestedAttribute{
									MarkdownDescription: "Password for HTTP Basic Authentication",
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											MarkdownDescription: "Variant discriminator.",
											Computed:            true,
										},
										"file_path": schema.StringAttribute{
											MarkdownDescription: "File path to read the secret from",
											Computed:            true,
										},
										"secret": schema.StringAttribute{
											MarkdownDescription: "Password or secret value",
											Computed:            true,
											Sensitive:           true,
										},
										"variable_name": schema.StringAttribute{
											MarkdownDescription: "Environment variable name to read the secret from",
											Computed:            true,
										},
									},
								},
								"username": schema.StringAttribute{
									MarkdownDescription: "Username for HTTP Basic Authentication",
									Computed:            true,
								},
							},
						},
						"http_headers": schema.MapAttribute{
							MarkdownDescription: "Additional headers to include in HTTP requests",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "The minimum amount of time that must pass between each push request to the OpenTelemetry endpoint",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Maximum amount of time that Stalwart will wait for a response from the OpenTelemetry endpoint",
							Computed:            true,
						},
					},
				},
				"prometheus": schema.SingleNestedAttribute{
					MarkdownDescription: "Prometheus metrics endpoint configuration",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"auth_secret": schema.SingleNestedAttribute{
							MarkdownDescription: "The Prometheus endpoint's secret for Basic authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"auth_username": schema.StringAttribute{
							MarkdownDescription: "The Prometheus endpoint's username for Basic authentication",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:      "metrics_store",
		JMAPType:  "x:MetricsStore",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the storage backend for metrics data. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state. Requires an Enterprise license.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Allow invalid TLS certificates when connecting to the store",
					Computed:            true,
				},
				"auth_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the store",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"auth_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the store",
					Computed:            true,
				},
				"cluster_file": schema.StringAttribute{
					MarkdownDescription: "Path to the cluster file for the FoundationDB cluster",
					Computed:            true,
				},
				"database": schema.StringAttribute{
					MarkdownDescription: "Name of the database",
					Computed:            true,
				},
				"datacenter_id": schema.StringAttribute{
					MarkdownDescription: "Data center ID (optional)",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "Hostname of the database server",
					Computed:            true,
				},
				"machine_id": schema.StringAttribute{
					MarkdownDescription: "Machine ID in the FoundationDB cluster (optional)",
					Computed:            true,
				},
				"max_allowed_packet": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a packet in bytes",
					Computed:            true,
				},
				"options": schema.StringAttribute{
					MarkdownDescription: "Additional connection options",
					Computed:            true,
				},
				"pool_max_connections": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of connections to the store",
					Computed:            true,
				},
				"pool_min_connections": schema.Int64Attribute{
					MarkdownDescription: "Minimum number of connections to the store",
					Computed:            true,
				},
				"pool_recycling_method": schema.StringAttribute{
					MarkdownDescription: "Method to use when recycling connections in the pool",
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "Port of the database server",
					Computed:            true,
				},
				"read_replicas": schema.ListNestedAttribute{
					MarkdownDescription: "List of read replicas for the store Requires an Enterprise license.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"auth_secret": schema.SingleNestedAttribute{
								MarkdownDescription: "Password to connect to the store",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"auth_username": schema.StringAttribute{
								MarkdownDescription: "Username to connect to the store",
								Computed:            true,
							},
							"database": schema.StringAttribute{
								MarkdownDescription: "Name of the database",
								Computed:            true,
							},
							"host": schema.StringAttribute{
								MarkdownDescription: "Hostname of the database server",
								Computed:            true,
							},
							"options": schema.StringAttribute{
								MarkdownDescription: "Additional connection options",
								Computed:            true,
							},
							"port": schema.Int64Attribute{
								MarkdownDescription: "Port of the database server",
								Computed:            true,
							},
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Connection timeout to the database",
					Computed:            true,
				},
				"transaction_retry_delay": schema.Int64Attribute{
					MarkdownDescription: "Transaction maximum retry delay",
					Computed:            true,
				},
				"transaction_retry_limit": schema.Int64Attribute{
					MarkdownDescription: "Transaction retry limit",
					Computed:            true,
				},
				"transaction_timeout": schema.Int64Attribute{
					MarkdownDescription: "Transaction timeout",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Use TLS to connect to the store",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_connection_strategies",
		JMAPType: "x:MtaConnectionStrategy",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaConnectionStrategy` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "mta_connection_strategy",
		JMAPType: "x:MtaConnectionStrategy",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a connection strategy for outbound message delivery.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"connect_timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait for the connection to be established",
					Computed:            true,
				},
				"data_timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait for the DATA command response",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description of the connection strategy",
					Computed:            true,
				},
				"ehlo_hostname": schema.StringAttribute{
					MarkdownDescription: "Overrides the EHLO hostname that will be used when connecting using this strategy",
					Computed:            true,
				},
				"ehlo_timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait for the EHLO command response",
					Computed:            true,
				},
				"greeting_timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait for the SMTP greeting message",
					Computed:            true,
				},
				"mail_from_timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait for the MAIL-FROM command response",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short identifier for the strategy",
					Optional:            true,
					Computed:            true,
				},
				"rcpt_to_timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait for the RCPT-TO command response",
					Computed:            true,
				},
				"source_ips": schema.ListNestedAttribute{
					MarkdownDescription: "List of local IPv4 and IPv6 addresses to use when delivering emails to remote SMTP servers",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"ehlo_hostname": schema.StringAttribute{
								MarkdownDescription: "Overrides the EHLO hostname that will be used when connecting from this IP address",
								Computed:            true,
							},
							"source_ip": schema.StringAttribute{
								MarkdownDescription: "Local IPv4 and IPv6 address to use when delivering emails to remote SMTP servers",
								Computed:            true,
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "mta_delivery_schedule",
		JMAPType: "x:MtaDeliverySchedule",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines retry and notification intervals for message delivery.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "A short description of the schedule, which can be used to identify it in the list of schedules",
					Computed:            true,
				},
				"expiry": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to expire messages after a number of delivery attempts or after certain time (TTL)",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"expire": schema.Int64Attribute{
							MarkdownDescription: "Time after which the message will be expired if it is not delivered",
							Computed:            true,
						},
						"max_attempts": schema.Int64Attribute{
							MarkdownDescription: "Maximum number of delivery attempts before the message is considered failed",
							Computed:            true,
						},
					},
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short identifier for the schedule",
					Optional:            true,
					Computed:            true,
				},
				"notify": schema.SingleNestedAttribute{
					MarkdownDescription: "List of delayed delivery DSN notification intervals",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"intervals": schema.ListNestedAttribute{
							MarkdownDescription: "List of intervals",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"duration": schema.Int64Attribute{
										MarkdownDescription: "Time interval for retries or notifications",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"queue_id": schema.StringAttribute{
					MarkdownDescription: "The name of the virtual queue to use for this schedule",
					Computed:            true,
				},
				"retry": schema.SingleNestedAttribute{
					MarkdownDescription: "List of retry intervals for message delivery",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"intervals": schema.ListNestedAttribute{
							MarkdownDescription: "List of intervals",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"duration": schema.Int64Attribute{
										MarkdownDescription: "Time interval for retries or notifications",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "mta_delivery_schedules",
		JMAPType: "x:MtaDeliverySchedule",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaDeliverySchedule` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "mta_extensions",
		JMAPType:  "x:MtaExtensions",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures SMTP protocol extensions offered to clients. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"chunking": schema.SingleNestedAttribute{
					MarkdownDescription: "Enables chunking (RFC 1830), an extension that allows large messages to be transferred in chunks which may reduce the load on the network and server.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"deliver_by": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies the maximum delivery time for a message using the DELIVERBY (RFC 2852) extension, which allows the sender to request a specific delivery time for a message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"dsn": schema.SingleNestedAttribute{
					MarkdownDescription: "Enables delivery status notifications (RFC 3461), which allows the sender to request a delivery status notification (DSN) from the recipient's mail server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"expn": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies whether to enable the EXPN command, which allows the sender to request the membership of a mailing list. It is recommended to disable this command to prevent spammers from harvesting email addresses",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"future_release": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies the maximum time that a message can be held for delivery using the FUTURERELEASE (RFC 4865) extension",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"mt_priority": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies the priority assignment policy to advertise on the MT-PRIORITY (RFC 6710) extension, which allows the sender to specify a priority for a message. Available policies are mixer, stanag4406 and nsep, or false to disable this extension",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"no_soliciting": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies the text to include in the NOSOLICITING (RFC 3865) message, which indicates that the server does not accept unsolicited commercial email (UCE or spam)",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"pipelining": schema.SingleNestedAttribute{
					MarkdownDescription: "Enables SMTP pipelining (RFC 2920), which enables multiple commands to be sent in a single request to speed up communication between the client and server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"require_tls": schema.SingleNestedAttribute{
					MarkdownDescription: "Enables require TLS (RFC 8689), an extension that allows clients to require TLS encryption for the SMTP session",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"vrfy": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies whether to enable the VRFY command, which allows the sender to verify the existence of a mailbox. It is recommended to disable this command to prevent spammers from harvesting email addresses",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "mta_hook",
		JMAPType: "x:MtaHook",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an MTA hook endpoint for message processing.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Whether Stalwart should connect to a hook server that has an invalid TLS certificate",
					Computed:            true,
				},
				"enable": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that determines whether to enable this hook",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"http_auth": schema.SingleNestedAttribute{
					MarkdownDescription: "The type of HTTP authentication to use",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"bearer_token": schema.SingleNestedAttribute{
							MarkdownDescription: "Bearer token for HTTP Bearer Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password for HTTP Basic Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Username for HTTP Basic Authentication",
							Computed:            true,
						},
					},
				},
				"http_headers": schema.MapAttribute{
					MarkdownDescription: "Additional headers to include in HTTP requests",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"max_response_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size, in bytes, of a response that Stalwart will accept from this MTA Hook server",
					Computed:            true,
				},
				"stages": schema.SetAttribute{
					MarkdownDescription: "Which SMTP stages to run this hook on",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"temp_fail_on_error": schema.BoolAttribute{
					MarkdownDescription: "Whether to respond with a temporary failure (typically a 4xx SMTP status code) when Stalwart encounters an error while communicating with this MTA Hook server",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that Stalwart will wait for a response from this hook server",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "URL of the hook endpoint",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_hooks",
		JMAPType: "x:MtaHook",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaHook` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "mta_inbound_session",
		JMAPType:  "x:MtaInboundSession",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures inbound SMTP session timeouts and transfer limits. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"max_duration": schema.SingleNestedAttribute{
					MarkdownDescription: "The maximum duration of a session",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"timeout": schema.SingleNestedAttribute{
					MarkdownDescription: "How long to wait for a client to send a command before timing out",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"transfer_limit": schema.SingleNestedAttribute{
					MarkdownDescription: "The maximum number of bytes that can be transferred per session",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "mta_inbound_throttle",
		JMAPType: "x:MtaInboundThrottle",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an inbound rate limit rule for SMTP connections.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description for the throttle",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this throttle",
					Computed:            true,
				},
				"key": schema.SetAttribute{
					MarkdownDescription: "Optional list of context variables that determine where this throttle should be applied",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"match": schema.SingleNestedAttribute{
					MarkdownDescription: "Enable the imposition of concurrency and rate limits only when a specific condition is met",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"rate": schema.SingleNestedAttribute{
					MarkdownDescription: "Number of incoming requests over a period of time that the rate limiter will allow",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:     "mta_inbound_throttles",
		JMAPType: "x:MtaInboundThrottle",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaInboundThrottle` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "mta_milter",
		JMAPType: "x:MtaMilter",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a Milter filter endpoint for message processing.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Whether Stalwart should connect to a Milter filter server that has an invalid TLS certificate",
					Computed:            true,
				},
				"enable": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that determines whether to enable this milter",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"flags_action": schema.Int64Attribute{
					MarkdownDescription: "Optional flags to set on the Milter connection. See the Milter protocol documentation for details on available flags.",
					Computed:            true,
				},
				"flags_protocol": schema.Int64Attribute{
					MarkdownDescription: "Optional protocol flags to set on the Milter connection. See the Milter protocol documentation for details on available protocol flags.",
					Computed:            true,
				},
				"hostname": schema.StringAttribute{
					MarkdownDescription: "Hostname or IP address of the server where the Milter filter is running",
					Computed:            true,
				},
				"max_response_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size, in bytes, of a response that Stalwart will accept from this Milter server",
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "Network port on the Milter filter host server",
					Computed:            true,
				},
				"protocol_version": schema.StringAttribute{
					MarkdownDescription: "Version of the Milter protocol that Stalwart should use when communicating with the Milter server",
					Computed:            true,
				},
				"stages": schema.SetAttribute{
					MarkdownDescription: "Which SMTP stages to run the milter on",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"temp_fail_on_error": schema.BoolAttribute{
					MarkdownDescription: "Whether to respond with a temporary failure (typically a 4xx SMTP status code) when Stalwart encounters an error while communicating with this Milter server",
					Computed:            true,
				},
				"timeout_command": schema.Int64Attribute{
					MarkdownDescription: "How long Stalwart will wait to send a command to the Milter server",
					Computed:            true,
				},
				"timeout_connect": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that Stalwart will wait to establish a connection with this Milter server",
					Computed:            true,
				},
				"timeout_data": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time Stalwart will wait for a response from the Milter server",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Whether to use Transport Layer Security (TLS) for the connection between Stalwart and the Milter filter",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_milters",
		JMAPType: "x:MtaMilter",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaMilter` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "mta_outbound_strategy",
		JMAPType:  "x:MtaOutboundStrategy",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures outbound message delivery routing, scheduling, and TLS strategies. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"connection_strategy": schema.SingleNestedAttribute{
					MarkdownDescription: "An expression that returns the connection strategy to use when delivering messages to remote SMTP servers",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"route": schema.SingleNestedAttribute{
					MarkdownDescription: "An expression that returns the route name to use when delivering queued messages",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"schedule": schema.SingleNestedAttribute{
					MarkdownDescription: "An expression that returns the scheduling strategy to use when queueing messages",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"tls": schema.SingleNestedAttribute{
					MarkdownDescription: "An expression that returns the TLS strategy to use when delivering messages to remote SMTP servers",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "mta_outbound_throttle",
		JMAPType: "x:MtaOutboundThrottle",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an outbound rate limit rule for message delivery.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description for the throttle",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this throttle",
					Computed:            true,
				},
				"key": schema.SetAttribute{
					MarkdownDescription: "Optional list of context variables that determine where this throttle should be applied",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"match": schema.SingleNestedAttribute{
					MarkdownDescription: "Enable the imposition of concurrency and rate limits only when a specific condition is met",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"rate": schema.SingleNestedAttribute{
					MarkdownDescription: "Number of incoming requests over a period of time that the rate limiter will allow",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:     "mta_outbound_throttles",
		JMAPType: "x:MtaOutboundThrottle",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaOutboundThrottle` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "mta_queue_quota",
		JMAPType: "x:MtaQueueQuota",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a quota rule for message queues.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Short description for the quota",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this quota",
					Computed:            true,
				},
				"key": schema.SetAttribute{
					MarkdownDescription: "Optional list of context variables that determine where this quota should be applied",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"match": schema.SingleNestedAttribute{
					MarkdownDescription: "Enable the imposition of concurrency and rate limits only when a specific condition is met",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"messages": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of messages in the queue that this quota will allow",
					Computed:            true,
				},
				"size": schema.Int64Attribute{
					MarkdownDescription: "Maximum total size of messages in the queue that this quota will allow",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_queue_quotas",
		JMAPType: "x:MtaQueueQuota",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaQueueQuota` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "mta_route_local",
		JMAPType: "x:MtaRoute",
		Variant:  "Local",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a routing rule for outbound message delivery. Reads the Local Delivery variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "A short description of the route, which can be used to identify it in the list of routes",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short identifier for the route",
					Optional:            true,
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_route_mx",
		JMAPType: "x:MtaRoute",
		Variant:  "Mx",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a routing rule for outbound message delivery. Reads the Remote Delivery (MX) variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "A short description of the route, which can be used to identify it in the list of routes",
					Computed:            true,
				},
				"ip_lookup_strategy": schema.StringAttribute{
					MarkdownDescription: "IP resolution strategy for MX hosts",
					Computed:            true,
				},
				"max_multihomed": schema.Int64Attribute{
					MarkdownDescription: "For multi-homed remote servers, it is the maximum number of IP addresses to try on each delivery attempt",
					Computed:            true,
				},
				"max_mx_hosts": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of MX hosts to try on each delivery attempt",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short identifier for the route",
					Optional:            true,
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_route_relay",
		JMAPType: "x:MtaRoute",
		Variant:  "Relay",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a routing rule for outbound message delivery. Reads the Relay Host variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"address": schema.StringAttribute{
					MarkdownDescription: "The address of the remote SMTP server, which can be an IP address or a domain name",
					Computed:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Whether to allow connections to servers with invalid TLS certificates",
					Computed:            true,
				},
				"auth_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "The secret to use when authenticating with the remote server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"auth_username": schema.StringAttribute{
					MarkdownDescription: "The username to use when authenticating with the remote server",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "A short description of the route, which can be used to identify it in the list of routes",
					Computed:            true,
				},
				"implicit_tls": schema.BoolAttribute{
					MarkdownDescription: "Whether to use TLS encryption for all connections to the remote server",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short identifier for the route",
					Optional:            true,
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "The port number of the remote server, which is typically 25 for SMTP and 11200 for LMTP",
					Computed:            true,
				},
				"protocol": schema.StringAttribute{
					MarkdownDescription: "The protocol to use when delivering messages to the remote server, which can be either SMTP or LMTP",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_routes",
		JMAPType: "x:MtaRoute",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaRoute` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "mta_stage_auth",
		JMAPType:  "x:MtaStageAuth",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures SMTP authentication requirements and error handling. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"max_failures": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum number of authentication errors allowed before the session is disconnected",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"must_match_sender": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies whether the authenticated user or any of their associated e-mail addresses must match the sender of the email message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"require": schema.SingleNestedAttribute{
					MarkdownDescription: "Specifies whether authentication is necessary to send email messages",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"sasl_mechanisms": schema.SingleNestedAttribute{
					MarkdownDescription: "A list of SASL authentication mechanisms offered to clients, or an empty list to disable authentication. Stalwart supports PLAIN, LOGIN, and OAUTHBEARER mechanisms",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"wait_on_fail": schema.SingleNestedAttribute{
					MarkdownDescription: "Time interval to wait after an authentication failure",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "mta_stage_connect",
		JMAPType:  "x:MtaStageConnect",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures SMTP connection greeting and hostname settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"hostname": schema.SingleNestedAttribute{
					MarkdownDescription: "The SMTP server hostname",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"script": schema.SingleNestedAttribute{
					MarkdownDescription: "Which Sieve script to run when a client connects",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"smtp_greeting": schema.SingleNestedAttribute{
					MarkdownDescription: "The greeting message sent by the SMTP/LMTP server",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "mta_stage_data",
		JMAPType:  "x:MtaStageData",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures message processing rules for the SMTP DATA stage. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"add_auth_results_header": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to add an Authentication-Results header to the message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"add_date_header": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to add a Date header to the message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"add_delivered_to_header": schema.BoolAttribute{
					MarkdownDescription: "Whether to add a Delivered-To header to the message",
					Computed:            true,
				},
				"add_message_id_header": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to add a Message-Id header to the message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"add_received_header": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to add a Received header to the message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"add_received_spf_header": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to add a Received-SPF header to the message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"add_return_path_header": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to add a Return-Path header to the message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"enable_spam_filter": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to enable the spam filter for incoming messages",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"max_message_size": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum size of a message in bytes",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"max_messages": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum number of messages that can be submitted per SMTP session",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"max_received_headers": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum limit on the number of Received headers, which helps to prevent message loops",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"script": schema.SingleNestedAttribute{
					MarkdownDescription: "Which Sieve script to run after the client sends a DATA command",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "mta_stage_ehlo",
		JMAPType:  "x:MtaStageEhlo",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures EHLO command requirements and validation. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"reject_non_fqdn": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to reject EHLO commands that do not include a fully-qualified domain name as a parameter",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"require": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether the remote client must send an EHLO command before starting an SMTP transaction",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"script": schema.SingleNestedAttribute{
					MarkdownDescription: "Which Sieve script to run after the client sends an EHLO command",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "mta_stage_mail",
		JMAPType:  "x:MtaStageMail",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures MAIL FROM stage processing and sender validation. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"is_sender_allowed": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns true when the sender is allowed to send",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"rewrite": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression to rewrite the sender address",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"script": schema.SingleNestedAttribute{
					MarkdownDescription: "Which Sieve script to run after the client sends a MAIL command",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "mta_stage_rcpt",
		JMAPType:  "x:MtaStageRcpt",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures RCPT TO stage processing and recipient validation. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"allow_relaying": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether to allow relaying for non-local domains",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"max_failures": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum number of recipient errors before the session is disconnected",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"max_recipients": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum number of recipients per message",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"rewrite": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression to rewrite the recipient address",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"script": schema.SingleNestedAttribute{
					MarkdownDescription: "Which Sieve script to run after the client sends a RCPT command",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"wait_on_fail": schema.SingleNestedAttribute{
					MarkdownDescription: "Amount of time to wait after a recipient error",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "mta_sts",
		JMAPType:  "x:MtaSts",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the MTA-STS policy for the server. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"max_age": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to cache the MTA-STS policy",
					Computed:            true,
				},
				"mode": schema.StringAttribute{
					MarkdownDescription: "Whether to enforce, test, or disable the MTA-STS policy",
					Computed:            true,
				},
				"mx_hosts": schema.SetAttribute{
					MarkdownDescription: "Override the allowed MX hosts for the MTA-STS policy domain. If empty, the MX hosts are determined from the system settings",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "mta_tls_strategies",
		JMAPType: "x:MtaTlsStrategy",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaTlsStrategy` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "mta_tls_strategy",
		JMAPType: "x:MtaTlsStrategy",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a TLS security strategy for outbound connections.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Whether to allow connections to servers with invalid TLS certificates",
					Computed:            true,
				},
				"dane": schema.StringAttribute{
					MarkdownDescription: "Whether DANE is required, optional, or disabled",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "A short description of the TLS strategy, which can be used to identify it in the list of strategies",
					Computed:            true,
				},
				"mta_sts": schema.StringAttribute{
					MarkdownDescription: "Whether MTA-STS is required, optional, or disabled",
					Computed:            true,
				},
				"mta_sts_timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait for the MTA-STS policy lookup to complete",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short identifier for the TLS strategy",
					Optional:            true,
					Computed:            true,
				},
				"start_tls": schema.StringAttribute{
					MarkdownDescription: "Whether TLS support is required, optional, or disabled",
					Computed:            true,
				},
				"tls_timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum time to wait for the TLS handshake to complete",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_virtual_queue",
		JMAPType: "x:MtaVirtualQueue",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a virtual queue for organizing outbound message delivery.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "A short description of the queue, which can be used to identify it in the list of queues",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Unique identifier for the queue, max 8 characters",
					Optional:            true,
					Computed:            true,
				},
				"threads_per_node": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of threads to use for delivery on each node in the cluster",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "mta_virtual_queues",
		JMAPType: "x:MtaVirtualQueue",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:MtaVirtualQueue` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "network_listener",
		JMAPType: "x:NetworkListener",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a network listener for accepting incoming connections.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"bind": schema.SetAttribute{
					MarkdownDescription: "The addresses the listener will bind to",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"max_connections": schema.Int64Attribute{
					MarkdownDescription: "The maximum number of concurrent connections the listener will accept",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Unique identifier for the listener",
					Optional:            true,
					Computed:            true,
				},
				"override_proxy_trusted_networks": schema.SetAttribute{
					MarkdownDescription: "Enable proxy protocol for connections from these networks",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"protocol": schema.StringAttribute{
					MarkdownDescription: "The protocol used by the listener",
					Computed:            true,
				},
				"socket_backlog": schema.Int64Attribute{
					MarkdownDescription: "The maximum number of incoming connections that can be pending in the backlog queue",
					Computed:            true,
				},
				"socket_no_delay": schema.BoolAttribute{
					MarkdownDescription: "Whether the Nagle algorithm should be disabled for the socket",
					Computed:            true,
				},
				"socket_receive_buffer_size": schema.Int64Attribute{
					MarkdownDescription: "The size of the buffer used for receiving data",
					Computed:            true,
				},
				"socket_reuse_address": schema.BoolAttribute{
					MarkdownDescription: "Whether the socket can be bound to an address that is already in use by another socket",
					Computed:            true,
				},
				"socket_reuse_port": schema.BoolAttribute{
					MarkdownDescription: "Whether multiple sockets can be bound to the same address and port",
					Computed:            true,
				},
				"socket_send_buffer_size": schema.Int64Attribute{
					MarkdownDescription: "The size of the buffer used for sending data",
					Computed:            true,
				},
				"socket_tos_v4": schema.Int64Attribute{
					MarkdownDescription: "The type of service (TOS) value for the socket, which determines the priority of the traffic sent through the socket",
					Computed:            true,
				},
				"socket_ttl": schema.Int64Attribute{
					MarkdownDescription: "Time-to-live (TTL) value for the socket, which determines how many hops a packet can make before it is discarded",
					Computed:            true,
				},
				"tls_disable_cipher_suites": schema.SetAttribute{
					MarkdownDescription: "Which cipher suites to disable",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"tls_disable_protocols": schema.SetAttribute{
					MarkdownDescription: "Which TLS protocols to disable",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"tls_ignore_client_order": schema.BoolAttribute{
					MarkdownDescription: "Whether to ignore the client's cipher order",
					Computed:            true,
				},
				"tls_implicit": schema.BoolAttribute{
					MarkdownDescription: "Whether to use implicit TLS",
					Computed:            true,
				},
				"tls_timeout": schema.Int64Attribute{
					MarkdownDescription: "TLS handshake timeout",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable TLS for this listener",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "network_listeners",
		JMAPType: "x:NetworkListener",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:NetworkListener` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "oauth_client",
		JMAPType: "x:OAuthClient",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a registered OAuth client application.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"client_id": schema.StringAttribute{
					MarkdownDescription: "Unique identifier of the OAuth client",
					Computed:            true,
				},
				"contacts": schema.SetAttribute{
					MarkdownDescription: "Contact email addresses for the OAuth client",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the OAuth client",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the OAuth client",
					Computed:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "Expiration date of the OAuth client",
					Computed:            true,
				},
				"logo": schema.StringAttribute{
					MarkdownDescription: "URL or base64-encoded image representing the OAuth client",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this OAuth client belongs to",
					Computed:            true,
				},
				"redirect_uris": schema.SetAttribute{
					MarkdownDescription: "List of redirect URIs for the OAuth client",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"secret": schema.StringAttribute{
					MarkdownDescription: "Secret value of the OAuth client",
					Computed:            true,
					Sensitive:           true,
				},
			},
		},
	},
	{
		Name:     "oauth_clients",
		JMAPType: "x:OAuthClient",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:OAuthClient` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "oidc_provider",
		JMAPType:  "x:OidcProvider",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the OAuth and OpenID Connect provider settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"access_token_expiry": schema.Int64Attribute{
					MarkdownDescription: "Expiration time of an OAuth access token",
					Computed:            true,
				},
				"anonymous_client_registration": schema.BoolAttribute{
					MarkdownDescription: "Whether to allow OAuth clients to register without authentication",
					Computed:            true,
				},
				"auth_code_expiry": schema.Int64Attribute{
					MarkdownDescription: "Expiration time of an authorization code issued by the authorization code flow",
					Computed:            true,
				},
				"auth_code_max_attempts": schema.Int64Attribute{
					MarkdownDescription: "Number of failed login attempts before an authorization code is invalidated",
					Computed:            true,
				},
				"encryption_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Encryption key to use for OAuth",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"id_token_expiry": schema.Int64Attribute{
					MarkdownDescription: "Expiration time of an OpenID Connect ID token",
					Computed:            true,
				},
				"refresh_token_expiry": schema.Int64Attribute{
					MarkdownDescription: "Expiration time of an OAuth refresh token",
					Computed:            true,
				},
				"refresh_token_renewal": schema.Int64Attribute{
					MarkdownDescription: "Remaining time in a refresh token before a new one is issued to the client",
					Computed:            true,
				},
				"require_client_registration": schema.BoolAttribute{
					MarkdownDescription: "Whether to require OAuth client_ids to be registered before they can be used",
					Computed:            true,
				},
				"signature_algorithm": schema.StringAttribute{
					MarkdownDescription: "JWT signature algorithm to use for OpenID Connect.",
					Computed:            true,
				},
				"signature_key": schema.SingleNestedAttribute{
					MarkdownDescription: "Contents of the private key PEM used to sign JWTs for OpenID Connect.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"user_code_expiry": schema.Int64Attribute{
					MarkdownDescription: "Expiration time of a user code issued by the device authentication flow",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "public_key",
		JMAPType: "x:PublicKey",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a public key for email encryption (OpenPGP or S/MIME).",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"account_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the account this public key belongs to",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the public key",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the public key",
					Computed:            true,
				},
				"email_addresses": schema.SetAttribute{
					MarkdownDescription: "Email addresses associated with the public key",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "Expiration date of the public key",
					Computed:            true,
				},
				"key": schema.StringAttribute{
					MarkdownDescription: "OpenPGP or S/MIME public key data",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "public_keys",
		JMAPType: "x:PublicKey",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:PublicKey` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "queued_message",
		JMAPType: "x:QueuedMessage",
		Schema: schema.Schema{
			MarkdownDescription: "Represents a queued email message pending delivery.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"blob_id": schema.StringAttribute{
					MarkdownDescription: "Reference to the stored message content",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "When the message was received and queued",
					Computed:            true,
				},
				"env_id": schema.StringAttribute{
					MarkdownDescription: "SMTP ENVID parameter for delivery status notifications",
					Computed:            true,
				},
				"flags": schema.SetAttribute{
					MarkdownDescription: "Classification flags for the message",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"next_notify": schema.StringAttribute{
					MarkdownDescription: "When the next DSN notification is scheduled",
					Computed:            true,
				},
				"next_retry": schema.StringAttribute{
					MarkdownDescription: "When the next delivery attempt is scheduled",
					Computed:            true,
				},
				"priority": schema.Int64Attribute{
					MarkdownDescription: "Message priority (lower values = higher priority)",
					Computed:            true,
				},
				"received_from_ip": schema.StringAttribute{
					MarkdownDescription: "IP address of the client that submitted the message",
					Computed:            true,
				},
				"received_via_port": schema.Int64Attribute{
					MarkdownDescription: "Local port on which the message was received",
					Computed:            true,
				},
				"recipients": schema.MapAttribute{
					MarkdownDescription: "List of envelope recipients and their delivery status",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"return_path": schema.StringAttribute{
					MarkdownDescription: "Envelope sender address (MAIL FROM)",
					Computed:            true,
				},
				"size": schema.Int64Attribute{
					MarkdownDescription: "Size of the message in bytes",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "queued_messages",
		JMAPType: "x:QueuedMessage",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:QueuedMessage` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "report_settings",
		JMAPType:  "x:ReportSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures inbound report analysis and outbound report settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"inbound_report_addresses": schema.SetAttribute{
					MarkdownDescription: "List of addresses (which may include wildcards) from which reports will be intercepted and analyzed",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"inbound_report_forwarding": schema.BoolAttribute{
					MarkdownDescription: "Whether reports should be forwarded to their final recipient after analysis",
					Computed:            true,
				},
				"outbound_report_domain": schema.StringAttribute{
					MarkdownDescription: "The default domain name used for DSNs and other reports. If left empty, the default domain will be used.",
					Computed:            true,
				},
				"outbound_report_submitter": schema.SingleNestedAttribute{
					MarkdownDescription: "Report submitter address or leave empty to use the default hostname",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "role",
		JMAPType: "x:Role",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a named set of permissions that can be assigned to accounts, groups, or tenants.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the role",
					Computed:            true,
				},
				"disabled_permissions": schema.SetAttribute{
					MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"enabled_permissions": schema.SetAttribute{
					MarkdownDescription: "List of permissions that are explicitly enabled.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this role belongs to",
					Computed:            true,
				},
				"role_ids": schema.SetAttribute{
					MarkdownDescription: "List of roles this role extends",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "roles",
		JMAPType: "x:Role",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Role` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "search",
		JMAPType:  "x:Search",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures full-text search indexing for emails, calendars, contacts, and tracing. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"default_language": schema.StringAttribute{
					MarkdownDescription: "Default language to use when language detection is not possible",
					Computed:            true,
				},
				"index_batch_size": schema.Int64Attribute{
					MarkdownDescription: "Number of items to process in each batch during indexing operations",
					Computed:            true,
				},
				"index_calendar": schema.BoolAttribute{
					MarkdownDescription: "Enable full-text search indexing for calendar data",
					Computed:            true,
				},
				"index_calendar_fields": schema.SetAttribute{
					MarkdownDescription: "List of calendar fields to index",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"index_contact_fields": schema.SetAttribute{
					MarkdownDescription: "List of contact fields to index",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"index_contacts": schema.BoolAttribute{
					MarkdownDescription: "Enable full-text search indexing for contacts data",
					Computed:            true,
				},
				"index_email": schema.BoolAttribute{
					MarkdownDescription: "Enable full-text search indexing for email content and metadata",
					Computed:            true,
				},
				"index_email_fields": schema.SetAttribute{
					MarkdownDescription: "List of email fields to index",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"index_telemetry": schema.BoolAttribute{
					MarkdownDescription: "Enable full-text search indexing for tracing data Requires an Enterprise license.",
					Computed:            true,
				},
				"index_tracing_fields": schema.SetAttribute{
					MarkdownDescription: "List of tracing fields to index Requires an Enterprise license.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"supported_languages": schema.SetAttribute{
					MarkdownDescription: "List of languages to enable for full-text search",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "search_store",
		JMAPType:  "x:SearchStore",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the full-text search backend. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Whether to allow invalid SSL certificates",
					Computed:            true,
				},
				"auth_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the store",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"auth_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the store",
					Computed:            true,
				},
				"cluster_file": schema.StringAttribute{
					MarkdownDescription: "Path to the cluster file for the FoundationDB cluster",
					Computed:            true,
				},
				"database": schema.StringAttribute{
					MarkdownDescription: "Name of the database",
					Computed:            true,
				},
				"datacenter_id": schema.StringAttribute{
					MarkdownDescription: "Data center ID (optional)",
					Computed:            true,
				},
				"fail_on_timeout": schema.BoolAttribute{
					MarkdownDescription: "Whether to fail the operation if the task does not complete within the polling retries",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "Hostname of the database server",
					Computed:            true,
				},
				"http_auth": schema.SingleNestedAttribute{
					MarkdownDescription: "The type of HTTP authentication to use",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"bearer_token": schema.SingleNestedAttribute{
							MarkdownDescription: "Bearer token for HTTP Bearer Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password for HTTP Basic Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Username for HTTP Basic Authentication",
							Computed:            true,
						},
					},
				},
				"http_headers": schema.MapAttribute{
					MarkdownDescription: "Additional headers to include in HTTP requests",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"include_source": schema.BoolAttribute{
					MarkdownDescription: "Whether to index the full source document",
					Computed:            true,
				},
				"machine_id": schema.StringAttribute{
					MarkdownDescription: "Machine ID in the FoundationDB cluster (optional)",
					Computed:            true,
				},
				"max_allowed_packet": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a packet in bytes",
					Computed:            true,
				},
				"max_retries": schema.Int64Attribute{
					MarkdownDescription: "Number of times to poll for task status before giving up",
					Computed:            true,
				},
				"num_replicas": schema.Int64Attribute{
					MarkdownDescription: "Number of replicas for the index",
					Computed:            true,
				},
				"num_shards": schema.Int64Attribute{
					MarkdownDescription: "Number of shards for the index",
					Computed:            true,
				},
				"options": schema.StringAttribute{
					MarkdownDescription: "Additional connection options",
					Computed:            true,
				},
				"poll_interval": schema.Int64Attribute{
					MarkdownDescription: "Interval between polling for task status",
					Computed:            true,
				},
				"pool_max_connections": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of connections to the store",
					Computed:            true,
				},
				"pool_min_connections": schema.Int64Attribute{
					MarkdownDescription: "Minimum number of connections to the store",
					Computed:            true,
				},
				"pool_recycling_method": schema.StringAttribute{
					MarkdownDescription: "Method to use when recycling connections in the pool",
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "Port of the database server",
					Computed:            true,
				},
				"read_replicas": schema.ListNestedAttribute{
					MarkdownDescription: "List of read replicas for the store Requires an Enterprise license.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"auth_secret": schema.SingleNestedAttribute{
								MarkdownDescription: "Password to connect to the store",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"auth_username": schema.StringAttribute{
								MarkdownDescription: "Username to connect to the store",
								Computed:            true,
							},
							"database": schema.StringAttribute{
								MarkdownDescription: "Name of the database",
								Computed:            true,
							},
							"host": schema.StringAttribute{
								MarkdownDescription: "Hostname of the database server",
								Computed:            true,
							},
							"options": schema.StringAttribute{
								MarkdownDescription: "Additional connection options",
								Computed:            true,
							},
							"port": schema.Int64Attribute{
								MarkdownDescription: "Port of the database server",
								Computed:            true,
							},
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Timeout for HTTP requests",
					Computed:            true,
				},
				"transaction_retry_delay": schema.Int64Attribute{
					MarkdownDescription: "Transaction maximum retry delay",
					Computed:            true,
				},
				"transaction_retry_limit": schema.Int64Attribute{
					MarkdownDescription: "Transaction retry limit",
					Computed:            true,
				},
				"transaction_timeout": schema.Int64Attribute{
					MarkdownDescription: "Transaction timeout",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "URL of the ElasticSearch server",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Use TLS to connect to the store",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "security",
		JMAPType:  "x:Security",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures automatic IP banning rules for abuse, authentication failures, and port scanning. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"abuse_ban_period": schema.Int64Attribute{
					MarkdownDescription: "The duration of the ban for abuse attempts",
					Computed:            true,
				},
				"abuse_ban_rate": schema.SingleNestedAttribute{
					MarkdownDescription: "The maximum number of abuse attempts (relaying or failed RCPT TO attempts) before the IP is banned",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
				"auth_ban_period": schema.Int64Attribute{
					MarkdownDescription: "The duration of the ban for failed login attempts",
					Computed:            true,
				},
				"auth_ban_rate": schema.SingleNestedAttribute{
					MarkdownDescription: "The maximum number of failed login attempts before the IP is banned",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
				"loiter_ban_period": schema.Int64Attribute{
					MarkdownDescription: "The duration of the ban for loitering connections.",
					Computed:            true,
				},
				"loiter_ban_rate": schema.SingleNestedAttribute{
					MarkdownDescription: "The maximum number of loitering disconnections before the IP is banned",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
				"scan_ban_paths": schema.SetAttribute{
					MarkdownDescription: "The paths that will trigger an immediate ban if accessed. Each path should be a glob expression",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"scan_ban_period": schema.Int64Attribute{
					MarkdownDescription: "The duration of the ban for port scanning attempts",
					Computed:            true,
				},
				"scan_ban_rate": schema.SingleNestedAttribute{
					MarkdownDescription: "The maximum number of port scanning attempts before the IP is banned",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"count": schema.Int64Attribute{
							MarkdownDescription: "Count",
							Computed:            true,
						},
						"period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:      "sender_auth",
		JMAPType:  "x:SenderAuth",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures sender authentication verification including DKIM, SPF, DMARC, and ARC. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"arc_verify": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether ARC verification is strict, relaxed or disabled.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"dkim_sign_domain": schema.SingleNestedAttribute{
					MarkdownDescription: "Domain to use for DKIM signing",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"dkim_strict": schema.BoolAttribute{
					MarkdownDescription: "Whether to ignore insecure DKIM signatures such as those containing a length parameter",
					Computed:            true,
				},
				"dkim_verify": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether DKIM verification is strict, relaxed or disabled",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"dmarc_verify": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether DMARC verification is strict, relaxed or disabled",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"reverse_ip_verify": schema.SingleNestedAttribute{
					MarkdownDescription: "How strict to be when verifying the reverse DNS of the client IP",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"spf_ehlo_verify": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether SPF EHLO verification is strict, relaxed or disabled",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"spf_from_verify": schema.SingleNestedAttribute{
					MarkdownDescription: "Whether SPF MAIL FROM verification is strict, relaxed or disabled",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:      "sharing",
		JMAPType:  "x:Sharing",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures sharing settings for calendars, address books, and files. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"allow_directory_queries": schema.BoolAttribute{
					MarkdownDescription: "Whether authenticated users can query the directory via WebDAV and JMAP",
					Computed:            true,
				},
				"max_shares": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum number of sharees that can be added to a single shared item (calendar, address book or file)",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "sieve_system_interpreter",
		JMAPType:  "x:SieveSystemInterpreter",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the system-level Sieve script interpreter settings and limits. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"default_from_address": schema.SingleNestedAttribute{
					MarkdownDescription: "Default email address to use for the from field in email notifications sent from a Sieve script",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"default_from_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Default name to use for the from field in email notifications sent from a Sieve script",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"default_return_path": schema.SingleNestedAttribute{
					MarkdownDescription: "Default return path to use in email notifications sent from a Sieve script",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"dkim_sign_domain": schema.SingleNestedAttribute{
					MarkdownDescription: "Which domain's DKIM signatures to use when signing the email notifications sent from a Sieve script",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"duplicate_expiry": schema.Int64Attribute{
					MarkdownDescription: "Default expiration time for IDs stored by the duplicate extension from trusted scripts",
					Computed:            true,
				},
				"max_cpu_cycles": schema.Int64Attribute{
					MarkdownDescription: "Maximum number CPU cycles a script can use",
					Computed:            true,
				},
				"max_nested_includes": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of nested includes",
					Computed:            true,
				},
				"max_out_messages": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of outgoing messages",
					Computed:            true,
				},
				"max_received_headers": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of received headers",
					Computed:            true,
				},
				"max_redirects": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of redirects",
					Computed:            true,
				},
				"max_var_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a variable",
					Computed:            true,
				},
				"message_id_hostname": schema.StringAttribute{
					MarkdownDescription: "Override the default local hostname to use when generating a Message-Id header",
					Computed:            true,
				},
				"no_capability_check": schema.BoolAttribute{
					MarkdownDescription: "If enabled, language extensions can be used without being explicitly declared using the require statement",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "sieve_system_script",
		JMAPType: "x:SieveSystemScript",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a system Sieve script executed by the server.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"contents": schema.StringAttribute{
					MarkdownDescription: "Content of the script",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the script",
					Computed:            true,
				},
				"is_active": schema.BoolAttribute{
					MarkdownDescription: "Whether the script is active",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the script",
					Optional:            true,
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "sieve_system_scripts",
		JMAPType: "x:SieveSystemScript",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:SieveSystemScript` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "sieve_user_interpreter",
		JMAPType:  "x:SieveUserInterpreter",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the user-level Sieve script interpreter settings and limits. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"allowed_notify_uris": schema.SetAttribute{
					MarkdownDescription: "List of allowed URIs for the notify extension",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"default_expiry_duplicate": schema.Int64Attribute{
					MarkdownDescription: "Default expiration time for IDs stored by the duplicate extension from user scripts",
					Computed:            true,
				},
				"default_expiry_vacation": schema.Int64Attribute{
					MarkdownDescription: "Default expiration time for IDs stored by the vacation extension",
					Computed:            true,
				},
				"default_subject": schema.StringAttribute{
					MarkdownDescription: "Default subject of vacation responses",
					Computed:            true,
				},
				"default_subject_prefix": schema.StringAttribute{
					MarkdownDescription: "Default subject prefix of vacation responses",
					Computed:            true,
				},
				"disable_capabilities": schema.SetAttribute{
					MarkdownDescription: "List of capabilities to disable in the user interpreter",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"max_cpu_cycles": schema.Int64Attribute{
					MarkdownDescription: "Maximum number CPU cycles a script can use",
					Computed:            true,
				},
				"max_header_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a header",
					Computed:            true,
				},
				"max_includes": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of includes",
					Computed:            true,
				},
				"max_local_vars": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of local variables",
					Computed:            true,
				},
				"max_match_vars": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of match variables",
					Computed:            true,
				},
				"max_nested_blocks": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of nested blocks",
					Computed:            true,
				},
				"max_nested_for_every": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of nested foreach blocks",
					Computed:            true,
				},
				"max_nested_includes": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of nested includes",
					Computed:            true,
				},
				"max_nested_tests": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of nested tests",
					Computed:            true,
				},
				"max_out_messages": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of outgoing messages",
					Computed:            true,
				},
				"max_received_headers": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of received headers",
					Computed:            true,
				},
				"max_redirects": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of redirects",
					Computed:            true,
				},
				"max_script_name_length": schema.Int64Attribute{
					MarkdownDescription: "Maximum length of a script name",
					Computed:            true,
				},
				"max_script_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a script",
					Computed:            true,
				},
				"max_scripts": schema.Int64Attribute{
					MarkdownDescription: "The default maximum number of sieve scripts a user can create",
					Computed:            true,
				},
				"max_string_length": schema.Int64Attribute{
					MarkdownDescription: "Maximum length of a string",
					Computed:            true,
				},
				"max_var_name_length": schema.Int64Attribute{
					MarkdownDescription: "Maximum length of a variable name",
					Computed:            true,
				},
				"max_var_size": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a variable",
					Computed:            true,
				},
				"protected_headers": schema.SetAttribute{
					MarkdownDescription: "List of headers that cannot be deleted or added using the editheader extension",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "sieve_user_script",
		JMAPType: "x:SieveUserScript",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a global Sieve script available for user imports.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"contents": schema.StringAttribute{
					MarkdownDescription: "Content of the script",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the script",
					Computed:            true,
				},
				"is_active": schema.BoolAttribute{
					MarkdownDescription: "Whether the script is active",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the script",
					Optional:            true,
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "sieve_user_scripts",
		JMAPType: "x:SieveUserScript",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:SieveUserScript` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "spam_classifier",
		JMAPType:  "x:SpamClassifier",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the spam classifier model, training parameters, and auto-learning settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"hold_samples_for": schema.Int64Attribute{
					MarkdownDescription: "Duration to hold training samples for",
					Computed:            true,
				},
				"learn_ham_from_card": schema.BoolAttribute{
					MarkdownDescription: "Whether to automatically learn ham messages from senders in the user's address book.",
					Computed:            true,
				},
				"learn_ham_from_reply": schema.BoolAttribute{
					MarkdownDescription: "Whether to automatically learn ham messages that are replies to messages sent by the recipient.",
					Computed:            true,
				},
				"learn_spam_from_rbl_hits": schema.Int64Attribute{
					MarkdownDescription: "Number of DNSBL servers that list the sender to auto-learn as spam",
					Computed:            true,
				},
				"learn_spam_from_traps": schema.BoolAttribute{
					MarkdownDescription: "Train as spam messages sent to spam trap addresses",
					Computed:            true,
				},
				"min_ham_samples": schema.Int64Attribute{
					MarkdownDescription: "Minimum number of ham samples required for training",
					Computed:            true,
				},
				"min_spam_samples": schema.Int64Attribute{
					MarkdownDescription: "Minimum number of spam samples required for training",
					Computed:            true,
				},
				"model": schema.SingleNestedAttribute{
					MarkdownDescription: "The spam classifier model to use.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"feature_l2_normalize": schema.BoolAttribute{
							MarkdownDescription: "Whether to L2-normalize feature values in the spam classifier",
							Computed:            true,
						},
						"feature_log_scale": schema.BoolAttribute{
							MarkdownDescription: "Whether to apply sublinear scaling to feature values in the spam classifier",
							Computed:            true,
						},
						"indicator_parameters": schema.SingleNestedAttribute{
							MarkdownDescription: "Hyperparameters for the indicator features in the FTRL-CCFH model",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"alpha": schema.Float64Attribute{
									MarkdownDescription: "The alpha parameter for the FTRL-Proximal algorithm",
									Computed:            true,
								},
								"beta": schema.Float64Attribute{
									MarkdownDescription: "The beta parameter for the FTRL-Proximal algorithm",
									Computed:            true,
								},
								"l1_ratio": schema.Float64Attribute{
									MarkdownDescription: "The L1 regularization parameter for the FTRL-Proximal algorithm",
									Computed:            true,
								},
								"l2_ratio": schema.Float64Attribute{
									MarkdownDescription: "The L2 regularization parameter for the FTRL-Proximal algorithm",
									Computed:            true,
								},
								"num_features": schema.StringAttribute{
									MarkdownDescription: "The number of parameters (2^n)",
									Computed:            true,
								},
							},
						},
						"parameters": schema.SingleNestedAttribute{
							MarkdownDescription: "Hyperparameters for the FTRL-FH model",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"alpha": schema.Float64Attribute{
									MarkdownDescription: "The alpha parameter for the FTRL-Proximal algorithm",
									Computed:            true,
								},
								"beta": schema.Float64Attribute{
									MarkdownDescription: "The beta parameter for the FTRL-Proximal algorithm",
									Computed:            true,
								},
								"l1_ratio": schema.Float64Attribute{
									MarkdownDescription: "The L1 regularization parameter for the FTRL-Proximal algorithm",
									Computed:            true,
								},
								"l2_ratio": schema.Float64Attribute{
									MarkdownDescription: "The L2 regularization parameter for the FTRL-Proximal algorithm",
									Computed:            true,
								},
								"num_features": schema.StringAttribute{
									MarkdownDescription: "The number of parameters (2^n)",
									Computed:            true,
								},
							},
						},
					},
				},
				"reservoir_capacity": schema.Int64Attribute{
					MarkdownDescription: "The capacity of the training sample reservoir",
					Computed:            true,
				},
				"train_frequency": schema.Int64Attribute{
					MarkdownDescription: "Frequency to train the spam classifier",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_dnsbl_server_any",
		JMAPType: "x:SpamDnsblServer",
		Variant:  "Any",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNSBL server used for spam filtering lookups. Reads the Any variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the DNSBL server",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this DNSBL server",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "A unique name for this DNSBL server configuration",
					Optional:            true,
					Computed:            true,
				},
				"tag": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"zone": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the DNS zone to query.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "spam_dnsbl_server_body",
		JMAPType: "x:SpamDnsblServer",
		Variant:  "Body",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNSBL server used for spam filtering lookups. Reads the Body variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the DNSBL server",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this DNSBL server",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "A unique name for this DNSBL server configuration",
					Optional:            true,
					Computed:            true,
				},
				"tag": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"zone": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the DNS zone to query.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "spam_dnsbl_server_domain",
		JMAPType: "x:SpamDnsblServer",
		Variant:  "Domain",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNSBL server used for spam filtering lookups. Reads the Domain variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the DNSBL server",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this DNSBL server",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "A unique name for this DNSBL server configuration",
					Optional:            true,
					Computed:            true,
				},
				"tag": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"zone": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the DNS zone to query.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "spam_dnsbl_server_email",
		JMAPType: "x:SpamDnsblServer",
		Variant:  "Email",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNSBL server used for spam filtering lookups. Reads the E-mail variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the DNSBL server",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this DNSBL server",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "A unique name for this DNSBL server configuration",
					Optional:            true,
					Computed:            true,
				},
				"tag": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"zone": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the DNS zone to query.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "spam_dnsbl_server_header",
		JMAPType: "x:SpamDnsblServer",
		Variant:  "Header",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNSBL server used for spam filtering lookups. Reads the Header variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the DNSBL server",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this DNSBL server",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "A unique name for this DNSBL server configuration",
					Optional:            true,
					Computed:            true,
				},
				"tag": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"zone": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the DNS zone to query.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "spam_dnsbl_server_ip",
		JMAPType: "x:SpamDnsblServer",
		Variant:  "Ip",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNSBL server used for spam filtering lookups. Reads the IP variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the DNSBL server",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this DNSBL server",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "A unique name for this DNSBL server configuration",
					Optional:            true,
					Computed:            true,
				},
				"tag": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"zone": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the DNS zone to query.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "spam_dnsbl_server_url",
		JMAPType: "x:SpamDnsblServer",
		Variant:  "Url",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a DNSBL server used for spam filtering lookups. Reads the URL variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the DNSBL server",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this DNSBL server",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "A unique name for this DNSBL server configuration",
					Optional:            true,
					Computed:            true,
				},
				"tag": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"zone": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the DNS zone to query.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "spam_dnsbl_servers",
		JMAPType: "x:SpamDnsblServer",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:SpamDnsblServer` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "spam_dnsbl_settings",
		JMAPType:  "x:SpamDnsblSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures DNSBL query limits for spam filtering. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"domain_limit": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of DNSBL checks for domain names",
					Computed:            true,
				},
				"email_limit": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of DNSBL checks for E-mail addresses",
					Computed:            true,
				},
				"ip_limit": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of DNSBL checks for IP addresses",
					Computed:            true,
				},
				"url_limit": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of DNSBL checks for URLs",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_file_extension",
		JMAPType: "x:SpamFileExtension",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a file extension classification rule for spam filtering.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"content_types": schema.SetAttribute{
					MarkdownDescription: "The MIME types associated with this file extension",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"extension": schema.StringAttribute{
					MarkdownDescription: "The file name extension",
					Computed:            true,
				},
				"is_archive": schema.BoolAttribute{
					MarkdownDescription: "Whether this file extension is considered an archive",
					Computed:            true,
				},
				"is_bad": schema.BoolAttribute{
					MarkdownDescription: "Whether this file extension is considered bad",
					Computed:            true,
				},
				"is_nz": schema.BoolAttribute{
					MarkdownDescription: "Whether this file extension is considered a NZ file",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_file_extensions",
		JMAPType: "x:SpamFileExtension",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:SpamFileExtension` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "spam_llm",
		JMAPType:  "x:SpamLlm",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the LLM-based spam classifier. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state. Requires an Enterprise license.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"categories": schema.SetAttribute{
					MarkdownDescription: "The expected categories in the LLM response",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"confidence": schema.SetAttribute{
					MarkdownDescription: "The expected confidence levels in the LLM response",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"model_id": schema.StringAttribute{
					MarkdownDescription: "The AI model to use for the LLM classifier",
					Computed:            true,
				},
				"prompt": schema.StringAttribute{
					MarkdownDescription: "The prompt to use for the LLM classifier",
					Computed:            true,
				},
				"response_pos_category": schema.Int64Attribute{
					MarkdownDescription: "The position of the category field in the LLM response.",
					Computed:            true,
				},
				"response_pos_confidence": schema.Int64Attribute{
					MarkdownDescription: "The position of the confidence field in the LLM response.",
					Computed:            true,
				},
				"response_pos_explanation": schema.Int64Attribute{
					MarkdownDescription: "The position of the explanation field in the LLM response.",
					Computed:            true,
				},
				"separator": schema.StringAttribute{
					MarkdownDescription: "The separator character used to parse the LLM response.",
					Computed:            true,
				},
				"temperature": schema.Float64Attribute{
					MarkdownDescription: "The temperature to use for the LLM classifier",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "spam_pyzor",
		JMAPType:  "x:SpamPyzor",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the Pyzor collaborative spam detection service. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"allow_count": schema.Int64Attribute{
					MarkdownDescription: "The number of times the hash appears in the Pyzor allowlist",
					Computed:            true,
				},
				"block_count": schema.Int64Attribute{
					MarkdownDescription: "The number of times the hash appears in the Pyzor blocklist",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable the Pyzor classifier. Pyzor is a collaborative, networked system to detect and report spam.",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "The hostname of the Pyzor server",
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "The port to connect to the Pyzor server",
					Computed:            true,
				},
				"ratio": schema.Float64Attribute{
					MarkdownDescription: "The ratio of the number of times the hash appears in the Pyzor allowlist to the blocklist",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "The timeout for the Pyzor server. If the server does not respond within this time, the check is considered failed.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_rule_any",
		JMAPType: "x:SpamRule",
		Variant:  "Any",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a spam filter rule for message classification. Reads the Any variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"condition": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the rule",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this rule",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short name for the rule",
					Optional:            true,
					Computed:            true,
				},
				"priority": schema.Int64Attribute{
					MarkdownDescription: "The priority of the rule",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_rule_body",
		JMAPType: "x:SpamRule",
		Variant:  "Body",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a spam filter rule for message classification. Reads the Body variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"condition": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the rule",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this rule",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short name for the rule",
					Optional:            true,
					Computed:            true,
				},
				"priority": schema.Int64Attribute{
					MarkdownDescription: "The priority of the rule",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_rule_domain",
		JMAPType: "x:SpamRule",
		Variant:  "Domain",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a spam filter rule for message classification. Reads the Domain variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"condition": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the rule",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this rule",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short name for the rule",
					Optional:            true,
					Computed:            true,
				},
				"priority": schema.Int64Attribute{
					MarkdownDescription: "The priority of the rule",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_rule_email",
		JMAPType: "x:SpamRule",
		Variant:  "Email",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a spam filter rule for message classification. Reads the E-mail variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"condition": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the rule",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this rule",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short name for the rule",
					Optional:            true,
					Computed:            true,
				},
				"priority": schema.Int64Attribute{
					MarkdownDescription: "The priority of the rule",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_rule_header",
		JMAPType: "x:SpamRule",
		Variant:  "Header",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a spam filter rule for message classification. Reads the Header variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"condition": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the rule",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this rule",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short name for the rule",
					Optional:            true,
					Computed:            true,
				},
				"priority": schema.Int64Attribute{
					MarkdownDescription: "The priority of the rule",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_rule_ip",
		JMAPType: "x:SpamRule",
		Variant:  "Ip",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a spam filter rule for message classification. Reads the IP variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"condition": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the rule",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this rule",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short name for the rule",
					Optional:            true,
					Computed:            true,
				},
				"priority": schema.Int64Attribute{
					MarkdownDescription: "The priority of the rule",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_rule_url",
		JMAPType: "x:SpamRule",
		Variant:  "Url",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a spam filter rule for message classification. Reads the URL variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"condition": schema.SingleNestedAttribute{
					MarkdownDescription: "Expression that returns the tag to assign to the message.",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description for the rule",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable this rule",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Short name for the rule",
					Optional:            true,
					Computed:            true,
				},
				"priority": schema.Int64Attribute{
					MarkdownDescription: "The priority of the rule",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_rules",
		JMAPType: "x:SpamRule",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:SpamRule` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "spam_settings",
		JMAPType:  "x:SpamSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures global spam filter thresholds, greylisting, and trust settings. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Whether to enable the spam filter",
					Computed:            true,
				},
				"greylist_for": schema.Int64Attribute{
					MarkdownDescription: "Time to keep an IP address in the grey list. The grey list is used to delay messages from unknown senders.",
					Computed:            true,
				},
				"score_discard": schema.Float64Attribute{
					MarkdownDescription: "Discard messages with a score above this threshold",
					Computed:            true,
				},
				"score_reject": schema.Float64Attribute{
					MarkdownDescription: "Reject messages with a score above this threshold",
					Computed:            true,
				},
				"score_spam": schema.Float64Attribute{
					MarkdownDescription: "Mark as Spam messages with a score above this threshold",
					Computed:            true,
				},
				"spam_filter_rules_url": schema.StringAttribute{
					MarkdownDescription: "URL to download spam filter rules from",
					Computed:            true,
				},
				"trust_contacts": schema.BoolAttribute{
					MarkdownDescription: "Never classify messages as spam if they are sent from addresses present in the user's address book.",
					Computed:            true,
				},
				"trust_replies": schema.BoolAttribute{
					MarkdownDescription: "Never classify messages as spam if they are replies to messages sent by the recipient.",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_tag_discard",
		JMAPType: "x:SpamTag",
		Variant:  "Discard",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a score or action assigned to a spam classification tag. Reads the Discard variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"tag": schema.StringAttribute{
					MarkdownDescription: "The spam tag name",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_tag_reject",
		JMAPType: "x:SpamTag",
		Variant:  "Reject",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a score or action assigned to a spam classification tag. Reads the Reject variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"tag": schema.StringAttribute{
					MarkdownDescription: "The spam tag name",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_tag_score",
		JMAPType: "x:SpamTag",
		Variant:  "Score",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a score or action assigned to a spam classification tag. Reads the Assign Score variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"score": schema.Float64Attribute{
					MarkdownDescription: "The score for the tag",
					Computed:            true,
				},
				"tag": schema.StringAttribute{
					MarkdownDescription: "The spam tag name",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_tags",
		JMAPType: "x:SpamTag",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:SpamTag` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "spam_training_sample",
		JMAPType: "x:SpamTrainingSample",
		Schema: schema.Schema{
			MarkdownDescription: "Stores an email sample used for spam classifier training.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"account_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the account associated with this training sample",
					Computed:            true,
				},
				"blob_id": schema.StringAttribute{
					MarkdownDescription: "Reference to the stored message content",
					Computed:            true,
				},
				"delete_after_use": schema.BoolAttribute{
					MarkdownDescription: "Indicates whether the training sample should be deleted after being used for training",
					Computed:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "Timestamp when the training sample is scheduled to expire",
					Computed:            true,
				},
				"from": schema.StringAttribute{
					MarkdownDescription: "Email address of the sender of the message associated with this training sample",
					Computed:            true,
				},
				"is_spam": schema.BoolAttribute{
					MarkdownDescription: "Indicates whether the sample is spam (true) or ham (false)",
					Computed:            true,
				},
				"subject": schema.StringAttribute{
					MarkdownDescription: "Subject of the message associated with this training sample",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "spam_training_samples",
		JMAPType: "x:SpamTrainingSample",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:SpamTrainingSample` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "spf_report_settings",
		JMAPType:  "x:SpfReportSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures SPF authentication failure report generation. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"dkim_sign_domain": schema.SingleNestedAttribute{
					MarkdownDescription: "Which domain's DKIM signatures to use when signing the SPF authentication failure report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"from_address": schema.SingleNestedAttribute{
					MarkdownDescription: "Email address that will be used in the From header of the SPF authentication failure report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"from_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Name that will be used in the From header of the SPF authentication failure report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"send_frequency": schema.SingleNestedAttribute{
					MarkdownDescription: "Rate at which SPF reports will be sent to a given email address. When this rate is exceeded, no further SPF failure reports will be sent to that address",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"subject": schema.SingleNestedAttribute{
					MarkdownDescription: "Subject name that will be used in the SPF authentication failure report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "store_lookup",
		JMAPType: "x:StoreLookup",
		Schema: schema.Schema{
			MarkdownDescription: "Defines an external store used for lookups.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"namespace": schema.StringAttribute{
					MarkdownDescription: "Unique identifier for this store when used in lookups",
					Computed:            true,
				},
				"store": schema.SingleNestedAttribute{
					MarkdownDescription: "Store to use for lookups",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"allow_invalid_certs": schema.BoolAttribute{
							MarkdownDescription: "Allow invalid TLS certificates when connecting to the store",
							Computed:            true,
						},
						"auth_secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password to connect to the store",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"auth_username": schema.StringAttribute{
							MarkdownDescription: "Username to connect to the store",
							Computed:            true,
						},
						"database": schema.StringAttribute{
							MarkdownDescription: "Name of the database",
							Computed:            true,
						},
						"host": schema.StringAttribute{
							MarkdownDescription: "Hostname of the database server",
							Computed:            true,
						},
						"max_allowed_packet": schema.Int64Attribute{
							MarkdownDescription: "Maximum size of a packet in bytes",
							Computed:            true,
						},
						"max_retries": schema.Int64Attribute{
							MarkdownDescription: "Number of retries to connect to the Redis cluster",
							Computed:            true,
						},
						"max_retry_wait": schema.Int64Attribute{
							MarkdownDescription: "Maximum time to wait between retries",
							Computed:            true,
						},
						"min_retry_wait": schema.Int64Attribute{
							MarkdownDescription: "Minimum time to wait between retries",
							Computed:            true,
						},
						"options": schema.StringAttribute{
							MarkdownDescription: "Additional connection options",
							Computed:            true,
						},
						"path": schema.StringAttribute{
							MarkdownDescription: "Path to the SQLite data directory",
							Computed:            true,
						},
						"pool_max_connections": schema.Int64Attribute{
							MarkdownDescription: "Maximum number of connections to the store",
							Computed:            true,
						},
						"pool_min_connections": schema.Int64Attribute{
							MarkdownDescription: "Minimum number of connections to the store",
							Computed:            true,
						},
						"pool_recycling_method": schema.StringAttribute{
							MarkdownDescription: "Method to use when recycling connections in the pool",
							Computed:            true,
						},
						"pool_timeout_create": schema.Int64Attribute{
							MarkdownDescription: "Timeout for creating a new connection",
							Computed:            true,
						},
						"pool_timeout_recycle": schema.Int64Attribute{
							MarkdownDescription: "Timeout for recycling a connection",
							Computed:            true,
						},
						"pool_timeout_wait": schema.Int64Attribute{
							MarkdownDescription: "Timeout for waiting for a connection from the pool",
							Computed:            true,
						},
						"pool_workers": schema.Int64Attribute{
							MarkdownDescription: "Number of worker threads to use for the store, defaults to the number of cores",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Port of the database server",
							Computed:            true,
						},
						"protocol_version": schema.StringAttribute{
							MarkdownDescription: "Protocol Version",
							Computed:            true,
						},
						"read_from_replicas": schema.BoolAttribute{
							MarkdownDescription: "Whether to read from replicas",
							Computed:            true,
						},
						"read_replicas": schema.ListNestedAttribute{
							MarkdownDescription: "List of read replicas for the store Requires an Enterprise license.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"auth_secret": schema.SingleNestedAttribute{
										MarkdownDescription: "Password to connect to the store",
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												MarkdownDescription: "Variant discriminator.",
												Computed:            true,
											},
											"file_path": schema.StringAttribute{
												MarkdownDescription: "File path to read the secret from",
												Computed:            true,
											},
											"secret": schema.StringAttribute{
												MarkdownDescription: "Password or secret value",
												Computed:            true,
												Sensitive:           true,
											},
											"variable_name": schema.StringAttribute{
												MarkdownDescription: "Environment variable name to read the secret from",
												Computed:            true,
											},
										},
									},
									"auth_username": schema.StringAttribute{
										MarkdownDescription: "Username to connect to the store",
										Computed:            true,
									},
									"database": schema.StringAttribute{
										MarkdownDescription: "Name of the database",
										Computed:            true,
									},
									"host": schema.StringAttribute{
										MarkdownDescription: "Hostname of the database server",
										Computed:            true,
									},
									"options": schema.StringAttribute{
										MarkdownDescription: "Additional connection options",
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: "Port of the database server",
										Computed:            true,
									},
								},
							},
						},
						"sentinel_secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password to connect to the Sentinel nodes",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"sentinel_username": schema.StringAttribute{
							MarkdownDescription: "Username to connect to the Sentinel nodes",
							Computed:            true,
						},
						"service_name": schema.StringAttribute{
							MarkdownDescription: "Name of the monitored master (service) to query via the Sentinels",
							Computed:            true,
						},
						"stores": schema.ListNestedAttribute{
							MarkdownDescription: "Stores to use for sharding",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"auth_secret": schema.SingleNestedAttribute{
										MarkdownDescription: "Password to connect to the store",
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												MarkdownDescription: "Variant discriminator.",
												Computed:            true,
											},
											"file_path": schema.StringAttribute{
												MarkdownDescription: "File path to read the secret from",
												Computed:            true,
											},
											"secret": schema.StringAttribute{
												MarkdownDescription: "Password or secret value",
												Computed:            true,
												Sensitive:           true,
											},
											"variable_name": schema.StringAttribute{
												MarkdownDescription: "Environment variable name to read the secret from",
												Computed:            true,
											},
										},
									},
									"auth_username": schema.StringAttribute{
										MarkdownDescription: "Username to connect to the store",
										Computed:            true,
									},
									"max_retries": schema.Int64Attribute{
										MarkdownDescription: "Number of retries to connect to the Redis cluster",
										Computed:            true,
									},
									"max_retry_wait": schema.Int64Attribute{
										MarkdownDescription: "Maximum time to wait between retries",
										Computed:            true,
									},
									"min_retry_wait": schema.Int64Attribute{
										MarkdownDescription: "Minimum time to wait between retries",
										Computed:            true,
									},
									"pool_max_connections": schema.Int64Attribute{
										MarkdownDescription: "Maximum number of connections to the store",
										Computed:            true,
									},
									"pool_timeout_create": schema.Int64Attribute{
										MarkdownDescription: "Timeout for creating a new connection",
										Computed:            true,
									},
									"pool_timeout_recycle": schema.Int64Attribute{
										MarkdownDescription: "Timeout for recycling a connection",
										Computed:            true,
									},
									"pool_timeout_wait": schema.Int64Attribute{
										MarkdownDescription: "Timeout for waiting for a connection from the pool",
										Computed:            true,
									},
									"protocol_version": schema.StringAttribute{
										MarkdownDescription: "Protocol Version",
										Computed:            true,
									},
									"read_from_replicas": schema.BoolAttribute{
										MarkdownDescription: "Whether to read from replicas",
										Computed:            true,
									},
									"sentinel_secret": schema.SingleNestedAttribute{
										MarkdownDescription: "Password to connect to the Sentinel nodes",
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												MarkdownDescription: "Variant discriminator.",
												Computed:            true,
											},
											"file_path": schema.StringAttribute{
												MarkdownDescription: "File path to read the secret from",
												Computed:            true,
											},
											"secret": schema.StringAttribute{
												MarkdownDescription: "Password or secret value",
												Computed:            true,
												Sensitive:           true,
											},
											"variable_name": schema.StringAttribute{
												MarkdownDescription: "Environment variable name to read the secret from",
												Computed:            true,
											},
										},
									},
									"sentinel_username": schema.StringAttribute{
										MarkdownDescription: "Username to connect to the Sentinel nodes",
										Computed:            true,
									},
									"service_name": schema.StringAttribute{
										MarkdownDescription: "Name of the monitored master (service) to query via the Sentinels",
										Computed:            true,
									},
									"timeout": schema.Int64Attribute{
										MarkdownDescription: "Connection timeout to the database",
										Computed:            true,
									},
									"url": schema.StringAttribute{
										MarkdownDescription: "URL of the Redis server",
										Computed:            true,
									},
									"urls": schema.SetAttribute{
										MarkdownDescription: "URL(s) of the Redis server(s)",
										Computed:            true,
										ElementType:         types.StringType,
									},
								},
							},
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Connection timeout to the database",
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "URL of the Redis server",
							Computed:            true,
						},
						"urls": schema.SetAttribute{
							MarkdownDescription: "URL(s) of the Redis server(s)",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"use_tls": schema.BoolAttribute{
							MarkdownDescription: "Use TLS to connect to the store",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:     "store_lookups",
		JMAPType: "x:StoreLookup",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:StoreLookup` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "system_settings",
		JMAPType:  "x:SystemSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures core server settings including hostname, thread pool, and network services. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"default_certificate_id": schema.StringAttribute{
					MarkdownDescription: "Default TLS certificate to use when no SNI is provided by the client",
					Computed:            true,
				},
				"default_domain_id": schema.StringAttribute{
					MarkdownDescription: "Default domain to use for authentication and reports.",
					Computed:            true,
				},
				"default_hostname": schema.StringAttribute{
					MarkdownDescription: "The default hostname to use in SMTP greetings, MTA reports and other places where a hostname is needed but not specified.",
					Computed:            true,
				},
				"mail_exchangers": schema.ListNestedAttribute{
					MarkdownDescription: "List of mail exchangers to publish in DNS MX records.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"hostname": schema.StringAttribute{
								MarkdownDescription: "The hostname of the mail exchanger, or null to use the default hostname",
								Computed:            true,
							},
							"priority": schema.Int64Attribute{
								MarkdownDescription: "The priority of the mail exchanger, lower values are preferred. Mail exchangers with the same priority will be selected randomly.",
								Computed:            true,
							},
						},
					},
				},
				"max_connections": schema.Int64Attribute{
					MarkdownDescription: "The maximum number of concurrent connections the server will accept",
					Computed:            true,
				},
				"provider_info": schema.MapAttribute{
					MarkdownDescription: "Information about the provider to advertise in auto configuration services.",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"proxy_trusted_networks": schema.SetAttribute{
					MarkdownDescription: "Enable proxy protocol for connections from these networks",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"services": schema.MapAttribute{
					MarkdownDescription: "List of services to advertise in DNS and auto configuration services",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"thread_pool_size": schema.Int64Attribute{
					MarkdownDescription: "The number of threads in the global thread pool for CPU intensive tasks. Defaults to the number of CPU cores",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "task",
		JMAPType: "x:Task",
		Schema: schema.Schema{
			MarkdownDescription: "Represents a background task scheduled for execution.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"account_domain_id": schema.StringAttribute{
					MarkdownDescription: "Domain identifier of the account to be destroyed, if applicable",
					Computed:            true,
				},
				"account_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the account associated with this task",
					Computed:            true,
				},
				"account_name": schema.StringAttribute{
					MarkdownDescription: "Name of the account to be destroyed",
					Computed:            true,
				},
				"account_type": schema.StringAttribute{
					MarkdownDescription: "Type of the deleted account (user or group)",
					Computed:            true,
				},
				"alarm_id": schema.Int64Attribute{
					MarkdownDescription: "Identifier of the calendar alarm associated with this task",
					Computed:            true,
				},
				"archived_item_type": schema.StringAttribute{
					MarkdownDescription: "Type of the archived item associated with the blob",
					Computed:            true,
				},
				"archived_until": schema.StringAttribute{
					MarkdownDescription: "Timestamp until which the archived item will be deleted permanently if not restored",
					Computed:            true,
				},
				"blob_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the archived blob to be restored",
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Timestamp when the item was originally created",
					Computed:            true,
				},
				"document_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the document associated with this task",
					Computed:            true,
				},
				"document_type": schema.StringAttribute{
					MarkdownDescription: "Type of document associated with the task",
					Computed:            true,
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the domain associated with this task",
					Computed:            true,
				},
				"due": schema.StringAttribute{
					MarkdownDescription: "Due date and time for the task",
					Computed:            true,
				},
				"event_end": schema.StringAttribute{
					MarkdownDescription: "End date and time of the calendar event",
					Computed:            true,
				},
				"event_end_tz": schema.Int64Attribute{
					MarkdownDescription: "Timezone identifier for the end date and time",
					Computed:            true,
				},
				"event_id": schema.Int64Attribute{
					MarkdownDescription: "Identifier of the calendar event associated with this task",
					Computed:            true,
				},
				"event_start": schema.StringAttribute{
					MarkdownDescription: "Start date and time of the calendar event",
					Computed:            true,
				},
				"event_start_tz": schema.Int64Attribute{
					MarkdownDescription: "Timezone identifier for the start date and time",
					Computed:            true,
				},
				"maintenance_type": schema.StringAttribute{
					MarkdownDescription: "Type of maintenance operation to perform on the account",
					Computed:            true,
				},
				"message_ids": schema.SetAttribute{
					MarkdownDescription: "Message-IDs of the email messages to be merged into the thread",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"messages": schema.ListNestedAttribute{
					MarkdownDescription: "List of iTIP messages associated with this task",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"from": schema.StringAttribute{
								MarkdownDescription: "Email address of the sender of the iTIP message",
								Computed:            true,
							},
							"i_calendar_data": schema.StringAttribute{
								MarkdownDescription: "iCalendar data associated with the iTIP message",
								Computed:            true,
							},
							"is_from_organizer": schema.BoolAttribute{
								MarkdownDescription: "Indicates whether the sender is the organizer of the calendar event",
								Computed:            true,
							},
							"summary": schema.StringAttribute{
								MarkdownDescription: "Summary of the calendar event associated with the iTIP message",
								Computed:            true,
							},
							"to": schema.SetAttribute{
								MarkdownDescription: "Email addresses of the recipients of the iTIP message",
								Computed:            true,
								ElementType:         types.StringType,
							},
						},
					},
				},
				"on_success_renew_certificate": schema.BoolAttribute{
					MarkdownDescription: "Whether to automatically renew the domain's TLS certificate using ACME after successfully updating DNS records",
					Computed:            true,
				},
				"recurrence_id": schema.Int64Attribute{
					MarkdownDescription: "Recurrence identifier for the alarm, if applicable",
					Computed:            true,
				},
				"report_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the DMARC aggregate report associated with this task",
					Computed:            true,
				},
				"shard_index": schema.Int64Attribute{
					MarkdownDescription: "Index of the shard to perform maintenance on, if applicable for the maintenance type",
					Computed:            true,
				},
				"status": schema.SingleNestedAttribute{
					MarkdownDescription: "Current status of the task",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"attempt_number": schema.Int64Attribute{
							MarkdownDescription: "Number of attempts made to complete the task",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Date and time when the task was created",
							Computed:            true,
						},
						"due": schema.StringAttribute{
							MarkdownDescription: "Due date and time for the task",
							Computed:            true,
						},
						"failed_at": schema.StringAttribute{
							MarkdownDescription: "Date and time when the task failed",
							Computed:            true,
						},
						"failed_attempt_number": schema.Int64Attribute{
							MarkdownDescription: "Number of attempts made before the task failed",
							Computed:            true,
						},
						"failure_reason": schema.StringAttribute{
							MarkdownDescription: "Reason for the last failure",
							Computed:            true,
						},
					},
				},
				"tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the tenant to be maintained",
					Computed:            true,
				},
				"thread_name": schema.StringAttribute{
					MarkdownDescription: "Name of the thread to be merged",
					Computed:            true,
				},
				"trace_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the trace associated with this task",
					Computed:            true,
				},
				"update_records": schema.SetAttribute{
					MarkdownDescription: "Which DNS records should be updated for the domain as part of this task",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "task_manager",
		JMAPType:  "x:TaskManager",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures task execution settings including retry strategies. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"max_attempts": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of attempts for retrying a task",
					Computed:            true,
				},
				"strategy": schema.SingleNestedAttribute{
					MarkdownDescription: "Strategy to use for retrying failed tasks",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"delay": schema.Int64Attribute{
							MarkdownDescription: "Fixed delay before retrying a failed task",
							Computed:            true,
						},
						"factor": schema.Float64Attribute{
							MarkdownDescription: "Backoff factor for calculating retry delays",
							Computed:            true,
						},
						"initial_delay": schema.Int64Attribute{
							MarkdownDescription: "Initial delay before retrying a failed task",
							Computed:            true,
						},
						"jitter": schema.BoolAttribute{
							MarkdownDescription: "Whether to apply jitter to the retry delay to avoid thundering herd problem",
							Computed:            true,
						},
						"max_delay": schema.Int64Attribute{
							MarkdownDescription: "Maximum delay between retry attempts",
							Computed:            true,
						},
					},
				},
				"total_deadline": schema.Int64Attribute{
					MarkdownDescription: "Total deadline for retrying a task before it is marked as failed",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "tasks",
		JMAPType: "x:Task",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Task` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "tenant",
		JMAPType: "x:Tenant",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a tenant for multi-tenant environments with isolated resources and quotas. Requires an Enterprise license.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the tenant",
					Computed:            true,
				},
				"logo": schema.StringAttribute{
					MarkdownDescription: "URL or base64-encoded image representing the tenant",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the tenant",
					Optional:            true,
					Computed:            true,
				},
				"permissions": schema.SingleNestedAttribute{
					MarkdownDescription: "Permissions assigned to this tenant",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"disabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"enabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly enabled.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"quotas": schema.MapAttribute{
					MarkdownDescription: "Quotas for different object types within this tenant",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"roles": schema.SingleNestedAttribute{
					MarkdownDescription: "Roles assigned to this tenant",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"role_ids": schema.SetAttribute{
							MarkdownDescription: "List of roles assigned to this principal.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"used_disk_quota": schema.Int64Attribute{
					MarkdownDescription: "Amount of disk space currently used by this tenant (bytes)",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "tenants",
		JMAPType: "x:Tenant",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Tenant` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "tls_external_report",
		JMAPType: "x:TlsExternalReport",
		Schema: schema.Schema{
			MarkdownDescription: "Stores a TLS aggregate report received from an external source.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"expires_at": schema.StringAttribute{
					MarkdownDescription: "When the report is scheduled to be deleted",
					Computed:            true,
				},
				"from": schema.StringAttribute{
					MarkdownDescription: "Email address of the report sender",
					Computed:            true,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this report belongs to Requires an Enterprise license.",
					Computed:            true,
				},
				"received_at": schema.StringAttribute{
					MarkdownDescription: "When the report email was received",
					Computed:            true,
				},
				"report": schema.SingleNestedAttribute{
					MarkdownDescription: "TLS report content",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"contact_info": schema.StringAttribute{
							MarkdownDescription: "Contact information for the reporting organization",
							Computed:            true,
						},
						"date_range_end": schema.StringAttribute{
							MarkdownDescription: "End of the reporting period",
							Computed:            true,
						},
						"date_range_start": schema.StringAttribute{
							MarkdownDescription: "Start of the reporting period",
							Computed:            true,
						},
						"organization_name": schema.StringAttribute{
							MarkdownDescription: "Name of the organization that generated the report",
							Computed:            true,
						},
						"policies": schema.ListNestedAttribute{
							MarkdownDescription: "Policy evaluation results for each domain",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"failure_details": schema.ListNestedAttribute{
										MarkdownDescription: "Details of TLS failures encountered",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"additional_information": schema.StringAttribute{
													MarkdownDescription: "Additional context about the failure",
													Computed:            true,
												},
												"failed_session_count": schema.Int64Attribute{
													MarkdownDescription: "Number of sessions that failed with this error",
													Computed:            true,
												},
												"failure_reason_code": schema.StringAttribute{
													MarkdownDescription: "Error code or reason string for the failure",
													Computed:            true,
												},
												"receiving_ip": schema.StringAttribute{
													MarkdownDescription: "IP address of the receiving mail server",
													Computed:            true,
												},
												"receiving_mx_helo": schema.StringAttribute{
													MarkdownDescription: "HELO/EHLO string of the receiving mail server",
													Computed:            true,
												},
												"receiving_mx_hostname": schema.StringAttribute{
													MarkdownDescription: "Hostname of the receiving mail server",
													Computed:            true,
												},
												"result_type": schema.StringAttribute{
													MarkdownDescription: "Type of failure encountered",
													Computed:            true,
												},
												"sending_mta_ip": schema.StringAttribute{
													MarkdownDescription: "IP address of the sending mail server",
													Computed:            true,
												},
											},
										},
									},
									"mx_hosts": schema.SetAttribute{
										MarkdownDescription: "MX hostnames covered by the policy",
										Computed:            true,
										ElementType:         types.StringType,
									},
									"policy_domain": schema.StringAttribute{
										MarkdownDescription: "Domain the policy applies to",
										Computed:            true,
									},
									"policy_strings": schema.SetAttribute{
										MarkdownDescription: "Raw policy strings as retrieved",
										Computed:            true,
										ElementType:         types.StringType,
									},
									"policy_type": schema.StringAttribute{
										MarkdownDescription: "Type of TLS policy that was evaluated",
										Computed:            true,
									},
									"total_failed_sessions": schema.Int64Attribute{
										MarkdownDescription: "Number of sessions that failed TLS establishment",
										Computed:            true,
									},
									"total_successful_sessions": schema.Int64Attribute{
										MarkdownDescription: "Number of sessions that successfully established TLS",
										Computed:            true,
									},
								},
							},
						},
						"report_id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for this report",
							Computed:            true,
						},
					},
				},
				"subject": schema.StringAttribute{
					MarkdownDescription: "Subject line of the report email",
					Computed:            true,
				},
				"to": schema.SetAttribute{
					MarkdownDescription: "List of recipient email addresses",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "tls_external_reports",
		JMAPType: "x:TlsExternalReport",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:TlsExternalReport` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:     "tls_internal_report",
		JMAPType: "x:TlsInternalReport",
		Schema: schema.Schema{
			MarkdownDescription: "Stores an outbound TLS aggregate report pending delivery.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "When the report was created",
					Computed:            true,
				},
				"deliver_at": schema.StringAttribute{
					MarkdownDescription: "When the report is scheduled to be delivered",
					Computed:            true,
				},
				"domain": schema.StringAttribute{
					MarkdownDescription: "Domain this report is associated with",
					Computed:            true,
				},
				"http_rua": schema.SetAttribute{
					MarkdownDescription: "Reporting URIs from the TLS policy",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"mail_rua": schema.SetAttribute{
					MarkdownDescription: "Reporting email addresses from the TLS policy",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"policy_identifiers": schema.SetAttribute{
					MarkdownDescription: "Identifiers for the TLS policies that generated this report",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"report": schema.SingleNestedAttribute{
					MarkdownDescription: "TLS report content",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"contact_info": schema.StringAttribute{
							MarkdownDescription: "Contact information for the reporting organization",
							Computed:            true,
						},
						"date_range_end": schema.StringAttribute{
							MarkdownDescription: "End of the reporting period",
							Computed:            true,
						},
						"date_range_start": schema.StringAttribute{
							MarkdownDescription: "Start of the reporting period",
							Computed:            true,
						},
						"organization_name": schema.StringAttribute{
							MarkdownDescription: "Name of the organization that generated the report",
							Computed:            true,
						},
						"policies": schema.ListNestedAttribute{
							MarkdownDescription: "Policy evaluation results for each domain",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"failure_details": schema.ListNestedAttribute{
										MarkdownDescription: "Details of TLS failures encountered",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"additional_information": schema.StringAttribute{
													MarkdownDescription: "Additional context about the failure",
													Computed:            true,
												},
												"failed_session_count": schema.Int64Attribute{
													MarkdownDescription: "Number of sessions that failed with this error",
													Computed:            true,
												},
												"failure_reason_code": schema.StringAttribute{
													MarkdownDescription: "Error code or reason string for the failure",
													Computed:            true,
												},
												"receiving_ip": schema.StringAttribute{
													MarkdownDescription: "IP address of the receiving mail server",
													Computed:            true,
												},
												"receiving_mx_helo": schema.StringAttribute{
													MarkdownDescription: "HELO/EHLO string of the receiving mail server",
													Computed:            true,
												},
												"receiving_mx_hostname": schema.StringAttribute{
													MarkdownDescription: "Hostname of the receiving mail server",
													Computed:            true,
												},
												"result_type": schema.StringAttribute{
													MarkdownDescription: "Type of failure encountered",
													Computed:            true,
												},
												"sending_mta_ip": schema.StringAttribute{
													MarkdownDescription: "IP address of the sending mail server",
													Computed:            true,
												},
											},
										},
									},
									"mx_hosts": schema.SetAttribute{
										MarkdownDescription: "MX hostnames covered by the policy",
										Computed:            true,
										ElementType:         types.StringType,
									},
									"policy_domain": schema.StringAttribute{
										MarkdownDescription: "Domain the policy applies to",
										Computed:            true,
									},
									"policy_strings": schema.SetAttribute{
										MarkdownDescription: "Raw policy strings as retrieved",
										Computed:            true,
										ElementType:         types.StringType,
									},
									"policy_type": schema.StringAttribute{
										MarkdownDescription: "Type of TLS policy that was evaluated",
										Computed:            true,
									},
									"total_failed_sessions": schema.Int64Attribute{
										MarkdownDescription: "Number of sessions that failed TLS establishment",
										Computed:            true,
									},
									"total_successful_sessions": schema.Int64Attribute{
										MarkdownDescription: "Number of sessions that successfully established TLS",
										Computed:            true,
									},
								},
							},
						},
						"report_id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for this report",
							Computed:            true,
						},
					},
				},
			},
		},
	},
	{
		Name:     "tls_internal_reports",
		JMAPType: "x:TlsInternalReport",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:TlsInternalReport` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "tls_report_settings",
		JMAPType:  "x:TlsReportSettings",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures TLS aggregate report generation. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"contact_info": schema.SingleNestedAttribute{
					MarkdownDescription: "Contact information to be included in the report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"dkim_sign_domain": schema.SingleNestedAttribute{
					MarkdownDescription: "Which domain's DKIM signatures to use when signing the TLS aggregate report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"from_address": schema.SingleNestedAttribute{
					MarkdownDescription: "Email address that will be used in the From header of the TLS aggregate report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"from_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Name that will be used in the From header of the TLS aggregate report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"max_report_size": schema.SingleNestedAttribute{
					MarkdownDescription: "Maximum size of the TLS aggregate report in bytes",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"org_name": schema.SingleNestedAttribute{
					MarkdownDescription: "Name of the organization to be included in the report",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"send_frequency": schema.SingleNestedAttribute{
					MarkdownDescription: "Frequency at which the TLS aggregate reports will be sent. The options are hourly, daily, weekly, or disable to disable reporting",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"subject": schema.SingleNestedAttribute{
					MarkdownDescription: "Subject name that will be used in the TLS aggregate report email",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"else": schema.StringAttribute{
							MarkdownDescription: "Else condition",
							Computed:            true,
						},
						"match": schema.ListNestedAttribute{
							MarkdownDescription: "List of conditions and their corresponding results",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"if": schema.StringAttribute{
										MarkdownDescription: "If condition",
										Computed:            true,
									},
									"then": schema.StringAttribute{
										MarkdownDescription: "Then clause",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "tracer_journal",
		JMAPType: "x:Tracer",
		Variant:  "Journal",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a logging and tracing output method. Reads the Systemd Journal variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable or disable the tracer",
					Computed:            true,
				},
				"events": schema.SetAttribute{
					MarkdownDescription: "List of events to include or exclude based on filter mode",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"events_policy": schema.StringAttribute{
					MarkdownDescription: "How to interpret the events list",
					Computed:            true,
				},
				"level": schema.StringAttribute{
					MarkdownDescription: "The logging level for this tracer",
					Computed:            true,
				},
				"lossy": schema.BoolAttribute{
					MarkdownDescription: "Whether to drop log entries if there is backlog",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "tracer_log",
		JMAPType: "x:Tracer",
		Variant:  "Log",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a logging and tracing output method. Reads the Log file variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"ansi": schema.BoolAttribute{
					MarkdownDescription: "Whether to use ANSI colors in logs",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable or disable the tracer",
					Computed:            true,
				},
				"events": schema.SetAttribute{
					MarkdownDescription: "List of events to include or exclude based on filter mode",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"events_policy": schema.StringAttribute{
					MarkdownDescription: "How to interpret the events list",
					Computed:            true,
				},
				"level": schema.StringAttribute{
					MarkdownDescription: "The logging level for this tracer",
					Computed:            true,
				},
				"lossy": schema.BoolAttribute{
					MarkdownDescription: "Whether to drop log entries if there is backlog",
					Computed:            true,
				},
				"multiline": schema.BoolAttribute{
					MarkdownDescription: "Whether to write log entries as a single line or multiline",
					Computed:            true,
				},
				"path": schema.StringAttribute{
					MarkdownDescription: "The path to the log file",
					Computed:            true,
				},
				"prefix": schema.StringAttribute{
					MarkdownDescription: "The prefix for the log file",
					Computed:            true,
				},
				"rotate": schema.StringAttribute{
					MarkdownDescription: "The frequency to rotate the log file",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "tracer_otel_grpc",
		JMAPType: "x:Tracer",
		Variant:  "OtelGrpc",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a logging and tracing output method. Reads the Open Telemetry (gRPC) variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable or disable the tracer",
					Computed:            true,
				},
				"enable_log_exporter": schema.BoolAttribute{
					MarkdownDescription: "Whether to export logs to OpenTelemetry",
					Computed:            true,
				},
				"enable_span_exporter": schema.BoolAttribute{
					MarkdownDescription: "Whether to export spans to OpenTelemetry",
					Computed:            true,
				},
				"endpoint": schema.StringAttribute{
					MarkdownDescription: "The endpoint for Open Telemetry",
					Computed:            true,
				},
				"events": schema.SetAttribute{
					MarkdownDescription: "List of events to include or exclude based on filter mode",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"events_policy": schema.StringAttribute{
					MarkdownDescription: "How to interpret the events list",
					Computed:            true,
				},
				"http_auth": schema.SingleNestedAttribute{
					MarkdownDescription: "The type of HTTP authentication to use",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"bearer_token": schema.SingleNestedAttribute{
							MarkdownDescription: "Bearer token for HTTP Bearer Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password for HTTP Basic Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Username for HTTP Basic Authentication",
							Computed:            true,
						},
					},
				},
				"http_headers": schema.MapAttribute{
					MarkdownDescription: "Additional headers to include in HTTP requests",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"level": schema.StringAttribute{
					MarkdownDescription: "The logging level for this tracer",
					Computed:            true,
				},
				"lossy": schema.BoolAttribute{
					MarkdownDescription: "Whether to drop log entries if there is backlog",
					Computed:            true,
				},
				"throttle": schema.Int64Attribute{
					MarkdownDescription: "The minimum amount of time that must pass between each request to the OpenTelemetry endpoint",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that Stalwart will wait for a response from the OpenTelemetry endpoint",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "tracer_otel_http",
		JMAPType: "x:Tracer",
		Variant:  "OtelHttp",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a logging and tracing output method. Reads the Open Telemetry (HTTP) variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable or disable the tracer",
					Computed:            true,
				},
				"enable_log_exporter": schema.BoolAttribute{
					MarkdownDescription: "Whether to export logs to OpenTelemetry",
					Computed:            true,
				},
				"enable_span_exporter": schema.BoolAttribute{
					MarkdownDescription: "Whether to export spans to OpenTelemetry",
					Computed:            true,
				},
				"endpoint": schema.StringAttribute{
					MarkdownDescription: "The endpoint for Open Telemetry",
					Computed:            true,
				},
				"events": schema.SetAttribute{
					MarkdownDescription: "List of events to include or exclude based on filter mode",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"events_policy": schema.StringAttribute{
					MarkdownDescription: "How to interpret the events list",
					Computed:            true,
				},
				"http_auth": schema.SingleNestedAttribute{
					MarkdownDescription: "The type of HTTP authentication to use",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"bearer_token": schema.SingleNestedAttribute{
							MarkdownDescription: "Bearer token for HTTP Bearer Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password for HTTP Basic Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Username for HTTP Basic Authentication",
							Computed:            true,
						},
					},
				},
				"http_headers": schema.MapAttribute{
					MarkdownDescription: "Additional headers to include in HTTP requests",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"level": schema.StringAttribute{
					MarkdownDescription: "The logging level for this tracer",
					Computed:            true,
				},
				"lossy": schema.BoolAttribute{
					MarkdownDescription: "Whether to drop log entries if there is backlog",
					Computed:            true,
				},
				"throttle": schema.Int64Attribute{
					MarkdownDescription: "The minimum amount of time that must pass between each request to the OpenTelemetry endpoint",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that Stalwart will wait for a response from the OpenTelemetry endpoint",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "tracer_stdout",
		JMAPType: "x:Tracer",
		Variant:  "Stdout",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a logging and tracing output method. Reads the Console variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"ansi": schema.BoolAttribute{
					MarkdownDescription: "Whether to use ANSI colors in logs",
					Computed:            true,
				},
				"buffered": schema.BoolAttribute{
					MarkdownDescription: "Whether to buffer log entries before writing to console",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable or disable the tracer",
					Computed:            true,
				},
				"events": schema.SetAttribute{
					MarkdownDescription: "List of events to include or exclude based on filter mode",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"events_policy": schema.StringAttribute{
					MarkdownDescription: "How to interpret the events list",
					Computed:            true,
				},
				"level": schema.StringAttribute{
					MarkdownDescription: "The logging level for this tracer",
					Computed:            true,
				},
				"lossy": schema.BoolAttribute{
					MarkdownDescription: "Whether to drop log entries if there is backlog",
					Computed:            true,
				},
				"multiline": schema.BoolAttribute{
					MarkdownDescription: "Whether to write log entries as a single line or multiline",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "tracers",
		JMAPType: "x:Tracer",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:Tracer` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
	{
		Name:      "tracing_store",
		JMAPType:  "x:TracingStore",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures the storage backend for tracing data. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state. Requires an Enterprise license.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Variant discriminator.",
					Computed:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Allow invalid TLS certificates when connecting to the store",
					Computed:            true,
				},
				"auth_secret": schema.SingleNestedAttribute{
					MarkdownDescription: "Password to connect to the store",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"auth_username": schema.StringAttribute{
					MarkdownDescription: "Username to connect to the store",
					Computed:            true,
				},
				"cluster_file": schema.StringAttribute{
					MarkdownDescription: "Path to the cluster file for the FoundationDB cluster",
					Computed:            true,
				},
				"database": schema.StringAttribute{
					MarkdownDescription: "Name of the database",
					Computed:            true,
				},
				"datacenter_id": schema.StringAttribute{
					MarkdownDescription: "Data center ID (optional)",
					Computed:            true,
				},
				"host": schema.StringAttribute{
					MarkdownDescription: "Hostname of the database server",
					Computed:            true,
				},
				"machine_id": schema.StringAttribute{
					MarkdownDescription: "Machine ID in the FoundationDB cluster (optional)",
					Computed:            true,
				},
				"max_allowed_packet": schema.Int64Attribute{
					MarkdownDescription: "Maximum size of a packet in bytes",
					Computed:            true,
				},
				"options": schema.StringAttribute{
					MarkdownDescription: "Additional connection options",
					Computed:            true,
				},
				"pool_max_connections": schema.Int64Attribute{
					MarkdownDescription: "Maximum number of connections to the store",
					Computed:            true,
				},
				"pool_min_connections": schema.Int64Attribute{
					MarkdownDescription: "Minimum number of connections to the store",
					Computed:            true,
				},
				"pool_recycling_method": schema.StringAttribute{
					MarkdownDescription: "Method to use when recycling connections in the pool",
					Computed:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "Port of the database server",
					Computed:            true,
				},
				"read_replicas": schema.ListNestedAttribute{
					MarkdownDescription: "List of read replicas for the store Requires an Enterprise license.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"auth_secret": schema.SingleNestedAttribute{
								MarkdownDescription: "Password to connect to the store",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"file_path": schema.StringAttribute{
										MarkdownDescription: "File path to read the secret from",
										Computed:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "Password or secret value",
										Computed:            true,
										Sensitive:           true,
									},
									"variable_name": schema.StringAttribute{
										MarkdownDescription: "Environment variable name to read the secret from",
										Computed:            true,
									},
								},
							},
							"auth_username": schema.StringAttribute{
								MarkdownDescription: "Username to connect to the store",
								Computed:            true,
							},
							"database": schema.StringAttribute{
								MarkdownDescription: "Name of the database",
								Computed:            true,
							},
							"host": schema.StringAttribute{
								MarkdownDescription: "Hostname of the database server",
								Computed:            true,
							},
							"options": schema.StringAttribute{
								MarkdownDescription: "Additional connection options",
								Computed:            true,
							},
							"port": schema.Int64Attribute{
								MarkdownDescription: "Port of the database server",
								Computed:            true,
							},
						},
					},
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Connection timeout to the database",
					Computed:            true,
				},
				"transaction_retry_delay": schema.Int64Attribute{
					MarkdownDescription: "Transaction maximum retry delay",
					Computed:            true,
				},
				"transaction_retry_limit": schema.Int64Attribute{
					MarkdownDescription: "Transaction retry limit",
					Computed:            true,
				},
				"transaction_timeout": schema.Int64Attribute{
					MarkdownDescription: "Transaction timeout",
					Computed:            true,
				},
				"use_tls": schema.BoolAttribute{
					MarkdownDescription: "Use TLS to connect to the store",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "user",
		JMAPType: "x:Account",
		Variant:  "User",
		HasName:  true,
		Schema: schema.Schema{
			MarkdownDescription: "Defines a user or group account for authentication and email access. Reads the User account variant.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object. Set either `id` or `name`.",
					Optional:            true,
					Computed:            true,
				},
				"aliases": schema.ListNestedAttribute{
					MarkdownDescription: "List of email aliases for the account",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Description of the email alias",
								Computed:            true,
							},
							"domain_id": schema.StringAttribute{
								MarkdownDescription: "Identifier for the domain of the email alias (the part after the @ symbol).",
								Computed:            true,
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Whether this email alias is enabled",
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "The local part of the email alias (the part before the @ symbol)",
								Computed:            true,
							},
						},
					},
				},
				"created_at": schema.StringAttribute{
					MarkdownDescription: "Creation date of the account",
					Computed:            true,
				},
				"credentials": schema.ListNestedAttribute{
					MarkdownDescription: "List of credential objects representing authentication methods for the account",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								MarkdownDescription: "Variant discriminator.",
								Computed:            true,
							},
							"allowed_ips": schema.SetAttribute{
								MarkdownDescription: "List of allowed IP addresses or CIDR ranges for this credential",
								Computed:            true,
								ElementType:         types.StringType,
							},
							"created_at": schema.StringAttribute{
								MarkdownDescription: "Creation date of the credential",
								Computed:            true,
							},
							"credential_id": schema.StringAttribute{
								MarkdownDescription: "Unique identifier for the credential",
								Computed:            true,
							},
							"description": schema.StringAttribute{
								MarkdownDescription: "Description of the credential",
								Computed:            true,
							},
							"expires_at": schema.StringAttribute{
								MarkdownDescription: "Expiration date of the credential",
								Computed:            true,
							},
							"otp_auth": schema.StringAttribute{
								MarkdownDescription: "OTP authentication URI for the account",
								Computed:            true,
							},
							"permissions": schema.SingleNestedAttribute{
								MarkdownDescription: "List of permissions assigned to this credential",
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Variant discriminator.",
										Computed:            true,
									},
									"permissions": schema.SetAttribute{
										MarkdownDescription: "List of permissions to assign.",
										Computed:            true,
										ElementType:         types.StringType,
									},
								},
							},
							"secret": schema.StringAttribute{
								MarkdownDescription: "Secret value of the account",
								Computed:            true,
								Sensitive:           true,
							},
						},
					},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the account",
					Computed:            true,
				},
				"domain_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the domain this account belongs to. This is used to determine the email address of the account, which is formed as name@domain.",
					Computed:            true,
				},
				"email_address": schema.StringAttribute{
					MarkdownDescription: "Email address for the user account, formed as name@domain.",
					Computed:            true,
				},
				"encryption_at_rest": schema.SingleNestedAttribute{
					MarkdownDescription: "Encryption-at-rest settings for the account",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"allow_spam_training": schema.BoolAttribute{
							MarkdownDescription: "Whether to allow training the spam classifier with plaintext emails before encryption",
							Computed:            true,
						},
						"encrypt_on_append": schema.BoolAttribute{
							MarkdownDescription: "Whether to encrypt emails when they are appended to mailboxes",
							Computed:            true,
						},
						"public_key": schema.StringAttribute{
							MarkdownDescription: "Public key used for encrypting emails",
							Computed:            true,
						},
					},
				},
				"locale": schema.StringAttribute{
					MarkdownDescription: "Preferred locale for the account",
					Computed:            true,
				},
				"member_group_ids": schema.SetAttribute{
					MarkdownDescription: "List of groups that this account is a member of",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this account belongs to",
					Computed:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Name of the account, typically an email address local part.",
					Optional:            true,
					Computed:            true,
				},
				"permissions": schema.SingleNestedAttribute{
					MarkdownDescription: "Permissions assigned to this account",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"disabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"enabled_permissions": schema.SetAttribute{
							MarkdownDescription: "List of permissions that are explicitly enabled.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"quotas": schema.MapAttribute{
					MarkdownDescription: "Quotas for different object types within this account",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"roles": schema.SingleNestedAttribute{
					MarkdownDescription: "Roles assigned to this user account",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"role_ids": schema.SetAttribute{
							MarkdownDescription: "List of roles assigned to this principal.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
				"time_zone": schema.StringAttribute{
					MarkdownDescription: "Preferred time zone for the account",
					Computed:            true,
				},
				"used_disk_quota": schema.Int64Attribute{
					MarkdownDescription: "Amount of disk space currently used by this account (bytes)",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:      "web_dav",
		JMAPType:  "x:WebDav",
		Singleton: true,
		Schema: schema.Schema{
			MarkdownDescription: "Configures WebDAV protocol settings including property limits and locking. This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Always `singleton`.",
					Computed:            true,
				},
				"dead_property_max_size": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum size of a WebDAV dead property value that the server will accept",
					Computed:            true,
				},
				"enable_assisted_discovery": schema.BoolAttribute{
					MarkdownDescription: "Enables assisted discovery of WebDAV shared collections by modifying PROPFIND requests to the root collection. Requests with depth 1 are automatically changed to depth 2, which may cause compatibility issues with some clients that expect the original behavior.",
					Computed:            true,
				},
				"live_property_max_size": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum size of a WebDAV live property value that the server will accept",
					Computed:            true,
				},
				"max_lock_timeout": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum duration for which a lock can be held on a resource",
					Computed:            true,
				},
				"max_locks": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum number of locks that a user can create on a resource",
					Computed:            true,
				},
				"max_results": schema.Int64Attribute{
					MarkdownDescription: "Specifies the maximum number of results that a WebDAV query can return",
					Computed:            true,
				},
				"request_max_size": schema.Int64Attribute{
					MarkdownDescription: "Determines the maximum XML size of a WebDAV request that the server will accept",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "web_hook",
		JMAPType: "x:WebHook",
		Schema: schema.Schema{
			MarkdownDescription: "Defines a webhook endpoint for event notifications.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the object.",
					Required:            true,
				},
				"allow_invalid_certs": schema.BoolAttribute{
					MarkdownDescription: "Whether Stalwart should connect to a webhook endpoint that has an invalid TLS certificate",
					Computed:            true,
				},
				"discard_after": schema.Int64Attribute{
					MarkdownDescription: "The duration after which the webhook will be discarded if it cannot be delivered",
					Computed:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable or disable the tracer",
					Computed:            true,
				},
				"events": schema.SetAttribute{
					MarkdownDescription: "List of events to include or exclude based on filter mode",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"events_policy": schema.StringAttribute{
					MarkdownDescription: "How to interpret the events list",
					Computed:            true,
				},
				"http_auth": schema.SingleNestedAttribute{
					MarkdownDescription: "The type of HTTP authentication to use",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"bearer_token": schema.SingleNestedAttribute{
							MarkdownDescription: "Bearer token for HTTP Bearer Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"secret": schema.SingleNestedAttribute{
							MarkdownDescription: "Password for HTTP Basic Authentication",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Variant discriminator.",
									Computed:            true,
								},
								"file_path": schema.StringAttribute{
									MarkdownDescription: "File path to read the secret from",
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: "Password or secret value",
									Computed:            true,
									Sensitive:           true,
								},
								"variable_name": schema.StringAttribute{
									MarkdownDescription: "Environment variable name to read the secret from",
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Username for HTTP Basic Authentication",
							Computed:            true,
						},
					},
				},
				"http_headers": schema.MapAttribute{
					MarkdownDescription: "Additional headers to include in HTTP requests",
					Computed:            true,
					ElementType:         types.StringType,
				},
				"level": schema.StringAttribute{
					MarkdownDescription: "The logging level for this tracer",
					Computed:            true,
				},
				"lossy": schema.BoolAttribute{
					MarkdownDescription: "Whether to drop log entries if there is backlog",
					Computed:            true,
				},
				"signature_key": schema.SingleNestedAttribute{
					MarkdownDescription: "The HMAC key used to sign the webhook request body to prevent tampering",
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Variant discriminator.",
							Computed:            true,
						},
						"file_path": schema.StringAttribute{
							MarkdownDescription: "File path to read the secret from",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Password or secret value",
							Computed:            true,
							Sensitive:           true,
						},
						"variable_name": schema.StringAttribute{
							MarkdownDescription: "Environment variable name to read the secret from",
							Computed:            true,
						},
					},
				},
				"throttle": schema.Int64Attribute{
					MarkdownDescription: "The minimum amount of time that must pass between each request to the webhook endpoint",
					Computed:            true,
				},
				"timeout": schema.Int64Attribute{
					MarkdownDescription: "Maximum amount of time that Stalwart will wait for a response from this webhook",
					Computed:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "URL of the webhook endpoint",
					Computed:            true,
				},
			},
		},
	},
	{
		Name:     "web_hooks",
		JMAPType: "x:WebHook",
		Plural:   true,
		Schema: schema.Schema{
			MarkdownDescription: "Lists the identifiers of all `x:WebHook` objects on the server.",
			Attributes: map[string]schema.Attribute{
				"ids": schema.SetAttribute{
					MarkdownDescription: "Identifiers of all objects of this type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
}
