// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregional

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	awstypes "github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_wafregional_web_acl_association", name="Web ACL Association")
func resourceWebACLAssociation() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceWebACLAssociationCreate,
		ReadWithoutTimeout:   resourceWebACLAssociationRead,
		DeleteWithoutTimeout: resourceWebACLAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			names.AttrResourceARN: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidARN,
			},
			"web_acl_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWebACLAssociationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).WAFRegionalClient(ctx)

	webACLID := d.Get("web_acl_id").(string)
	resourceARN := d.Get(names.AttrResourceARN).(string)
	id := webACLAssociationCreateResourceID(webACLID, resourceARN)
	input := &wafregional.AssociateWebACLInput{
		ResourceArn: aws.String(resourceARN),
		WebACLId:    aws.String(webACLID),
	}

	_, err := tfresource.RetryWhenIsA[*awstypes.WAFUnavailableEntityException](ctx, d.Timeout(schema.TimeoutCreate), func() (any, error) {
		return conn.AssociateWebACL(ctx, input)
	})

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating WAF Regional WebACL Association (%s): %s", id, err)
	}

	d.SetId(id)

	return diags
}

func resourceWebACLAssociationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).WAFRegionalClient(ctx)

	_, resourceARN, err := webACLAssociationParseResourceID(d.Id())
	if err != nil {
		return sdkdiag.AppendFromErr(diags, err)
	}

	webACL, err := findWebACLByResourceARN(ctx, conn, resourceARN)

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] WAF Regional WebACL Association (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading WAF Regional WebACL Association (%s): %s", d.Id(), err)
	}

	d.Set(names.AttrResourceARN, resourceARN)
	d.Set("web_acl_id", webACL.WebACLId)

	return diags
}

func resourceWebACLAssociationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).WAFRegionalClient(ctx)

	_, resourceARN, err := webACLAssociationParseResourceID(d.Id())
	if err != nil {
		return sdkdiag.AppendFromErr(diags, err)
	}

	input := wafregional.DisassociateWebACLInput{
		ResourceArn: aws.String(resourceARN),
	}
	_, err = conn.DisassociateWebACL(ctx, &input)

	if errs.IsA[*awstypes.WAFNonexistentItemException](err) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting WAF Regional Web ACL Association (%s): %s", d.Id(), err)
	}

	return diags
}

func findWebACLByResourceARN(ctx context.Context, conn *wafregional.Client, arn string) (*awstypes.WebACLSummary, error) {
	input := &wafregional.GetWebACLForResourceInput{
		ResourceArn: aws.String(arn),
	}

	output, err := conn.GetWebACLForResource(ctx, input)

	if errs.IsA[*awstypes.WAFNonexistentItemException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.WebACLSummary == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.WebACLSummary, nil
}

const webACLAssociationResourceIDSeparator = ":"

func webACLAssociationCreateResourceID(webACLID, resourceARN string) string {
	parts := []string{webACLID, resourceARN}
	id := strings.Join(parts, webACLAssociationResourceIDSeparator)

	return id
}

func webACLAssociationParseResourceID(id string) (string, string, error) { //nolint:unparam
	parts := strings.SplitN(id, webACLAssociationResourceIDSeparator, 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format for ID (%[1]s), expected WEB-ACL-ID%[2]sRESOURCE-ARN", id, webACLAssociationResourceIDSeparator)
	}

	return parts[0], parts[1], nil
}
