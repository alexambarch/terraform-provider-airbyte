// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type DestinationMssqlSSHTunnelMethod struct {
	NoTunnel               *DestinationMssqlNoTunnel               `queryParam:"inline" tfsdk:"no_tunnel" tfPlanOnly:"true"`
	PasswordAuthentication *DestinationMssqlPasswordAuthentication `queryParam:"inline" tfsdk:"password_authentication" tfPlanOnly:"true"`
	SSHKeyAuthentication   *DestinationMssqlSSHKeyAuthentication   `queryParam:"inline" tfsdk:"ssh_key_authentication" tfPlanOnly:"true"`
}
