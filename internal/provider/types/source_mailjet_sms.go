// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMailjetSms struct {
	EndDate   types.Int64  `tfsdk:"end_date"`
	StartDate types.Int64  `tfsdk:"start_date"`
	Token     types.String `tfsdk:"token"`
}
