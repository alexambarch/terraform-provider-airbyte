// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceLeverHiring struct {
	Credentials *SourceLeverHiringAuthenticationMechanism `tfsdk:"credentials"`
	Environment types.String                              `tfsdk:"environment"`
	StartDate   types.String                              `tfsdk:"start_date"`
}
