// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSquare struct {
	Credentials           *SourceSquareAuthentication `tfsdk:"credentials"`
	IncludeDeletedObjects types.Bool                  `tfsdk:"include_deleted_objects"`
	IsSandbox             types.Bool                  `tfsdk:"is_sandbox"`
	StartDate             types.String                `tfsdk:"start_date"`
}
