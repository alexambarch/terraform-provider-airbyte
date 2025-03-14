// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SourceOracleEncryption struct {
	NativeNetworkEncryptionNNE    *NativeNetworkEncryptionNNE    `queryParam:"inline" tfsdk:"native_network_encryption_nne" tfPlanOnly:"true"`
	TLSEncryptedVerifyCertificate *TLSEncryptedVerifyCertificate `queryParam:"inline" tfsdk:"tls_encrypted_verify_certificate" tfPlanOnly:"true"`
	Unencrypted                   *SourceOracleUnencrypted       `queryParam:"inline" tfsdk:"unencrypted" tfPlanOnly:"true"`
}
