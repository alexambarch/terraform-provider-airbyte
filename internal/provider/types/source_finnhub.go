// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceFinnhub struct {
	APIKey             types.String   `tfsdk:"api_key"`
	Exchange           types.String   `tfsdk:"exchange"`
	MarketNewsCategory types.String   `tfsdk:"market_news_category"`
	StartDate2         types.String   `tfsdk:"start_date_2"`
	Symbols            []types.String `tfsdk:"symbols"`
}
