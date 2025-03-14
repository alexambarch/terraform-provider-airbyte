// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSftpBulk struct {
	Credentials    SourceSftpBulkAuthentication          `tfsdk:"credentials"`
	DeliveryMethod *SourceSftpBulkDeliveryMethod         `tfsdk:"delivery_method"`
	FolderPath     types.String                          `tfsdk:"folder_path"`
	Host           types.String                          `tfsdk:"host"`
	Port           types.Int64                           `tfsdk:"port"`
	StartDate      types.String                          `tfsdk:"start_date"`
	Streams        []SourceSftpBulkFileBasedStreamConfig `tfsdk:"streams"`
	Username       types.String                          `tfsdk:"username"`
}
