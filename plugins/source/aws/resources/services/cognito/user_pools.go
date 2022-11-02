// Code generated by codegen; DO NOT EDIT.

package cognito

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserPools() *schema.Table {
	return &schema.Table{
		Name:                "aws_cognito_user_pools",
		Description:         "https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolType.html",
		Resolver:            fetchCognitoUserPools,
		PreResourceResolver: getUserPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer("cognito-identity"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "account_recovery_setting",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccountRecoverySetting"),
			},
			{
				Name:     "admin_create_user_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdminCreateUserConfig"),
			},
			{
				Name:     "alias_attributes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AliasAttributes"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "auto_verified_attributes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AutoVerifiedAttributes"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "custom_domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomDomain"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
			{
				Name:     "device_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeviceConfiguration"),
			},
			{
				Name:     "domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Domain"),
			},
			{
				Name:     "email_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EmailConfiguration"),
			},
			{
				Name:     "email_configuration_failure",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EmailConfigurationFailure"),
			},
			{
				Name:     "email_verification_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EmailVerificationMessage"),
			},
			{
				Name:     "email_verification_subject",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EmailVerificationSubject"),
			},
			{
				Name:     "estimated_number_of_users",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EstimatedNumberOfUsers"),
			},
			{
				Name:     "lambda_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LambdaConfig"),
			},
			{
				Name:     "last_modified_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedDate"),
			},
			{
				Name:     "mfa_configuration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MfaConfiguration"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policies"),
			},
			{
				Name:     "schema_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SchemaAttributes"),
			},
			{
				Name:     "sms_authentication_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SmsAuthenticationMessage"),
			},
			{
				Name:     "sms_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SmsConfiguration"),
			},
			{
				Name:     "sms_configuration_failure",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SmsConfigurationFailure"),
			},
			{
				Name:     "sms_verification_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SmsVerificationMessage"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "user_attribute_update_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserAttributeUpdateSettings"),
			},
			{
				Name:     "user_pool_add_ons",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserPoolAddOns"),
			},
			{
				Name:     "user_pool_tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserPoolTags"),
			},
			{
				Name:     "username_attributes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("UsernameAttributes"),
			},
			{
				Name:     "username_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UsernameConfiguration"),
			},
			{
				Name:     "verification_message_template",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VerificationMessageTemplate"),
			},
		},

		Relations: []*schema.Table{
			UserPoolIdentityProviders(),
		},
	}
}
