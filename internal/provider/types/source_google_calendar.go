// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleCalendar struct {
	Calendarid          types.String `tfsdk:"calendarid"`
	ClientID            types.String `tfsdk:"client_id"`
	ClientRefreshToken2 types.String `tfsdk:"client_refresh_token_2"`
	ClientSecret        types.String `tfsdk:"client_secret"`
}
