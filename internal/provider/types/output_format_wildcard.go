// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type OutputFormatWildcard struct {
	JSONLinesNewlineDelimitedJSON *JSONLinesNewlineDelimitedJSON `queryParam:"inline" tfsdk:"json_lines_newline_delimited_json" tfPlanOnly:"true"`
	ParquetColumnarStorage        *ParquetColumnarStorage        `queryParam:"inline" tfsdk:"parquet_columnar_storage" tfPlanOnly:"true"`
}
