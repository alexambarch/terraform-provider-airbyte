// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceClickupAPI struct {
	APIToken           types.String `tfsdk:"api_token"`
	IncludeClosedTasks types.Bool   `tfsdk:"include_closed_tasks"`
}
