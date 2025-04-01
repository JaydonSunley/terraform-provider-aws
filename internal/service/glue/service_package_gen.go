// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceRegistry,
			TypeName: "aws_glue_registry",
			Name:     "Registry",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceCatalogTableOptimizer,
			TypeName: "aws_glue_catalog_table_optimizer",
			Name:     "Catalog Table Optimizer",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCatalogTable,
			TypeName: "aws_glue_catalog_table",
			Name:     "Catalog Table",
		},
		{
			Factory:  dataSourceConnection,
			TypeName: "aws_glue_connection",
			Name:     "Connection",
		},
		{
			Factory:  dataSourceDataCatalogEncryptionSettings,
			TypeName: "aws_glue_data_catalog_encryption_settings",
			Name:     "Data Catalog Encryption Settings",
		},
		{
			Factory:  DataSourceScript,
			TypeName: "aws_glue_script",
			Name:     "Script",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCatalogDatabase,
			TypeName: "aws_glue_catalog_database",
			Name:     "Database",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCatalogTable,
			TypeName: "aws_glue_catalog_table",
			Name:     "Catalog Table",
		},
		{
			Factory:  resourceClassifier,
			TypeName: "aws_glue_classifier",
			Name:     "Classifier",
		},
		{
			Factory:  resourceConnection,
			TypeName: "aws_glue_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCrawler,
			TypeName: "aws_glue_crawler",
			Name:     "Crawler",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDataCatalogEncryptionSettings,
			TypeName: "aws_glue_data_catalog_encryption_settings",
			Name:     "Data Catalog Encryption Settings",
		},
		{
			Factory:  resourceDataQualityRuleset,
			TypeName: "aws_glue_data_quality_ruleset",
			Name:     "Data Quality Ruleset",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceDevEndpoint,
			TypeName: "aws_glue_dev_endpoint",
			Name:     "Dev Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceJob,
			TypeName: "aws_glue_job",
			Name:     "Job",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceMLTransform,
			TypeName: "aws_glue_ml_transform",
			Name:     "ML Transform",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourcePartition,
			TypeName: "aws_glue_partition",
			Name:     "Partition",
		},
		{
			Factory:  ResourcePartitionIndex,
			TypeName: "aws_glue_partition_index",
			Name:     "Partition Index",
		},
		{
			Factory:  ResourceRegistry,
			TypeName: "aws_glue_registry",
			Name:     "Registry",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_glue_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  ResourceSchema,
			TypeName: "aws_glue_schema",
			Name:     "Schema",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceSecurityConfiguration,
			TypeName: "aws_glue_security_configuration",
			Name:     "Security Configuration",
		},
		{
			Factory:  ResourceTrigger,
			TypeName: "aws_glue_trigger",
			Name:     "Trigger",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceUserDefinedFunction,
			TypeName: "aws_glue_user_defined_function",
			Name:     "User Defined Function",
		},
		{
			Factory:  ResourceWorkflow,
			TypeName: "aws_glue_workflow",
			Name:     "Workflow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Glue
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*glue.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*glue.Options){
		glue.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return glue.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*glue.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*glue.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *glue.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*glue.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
