// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PermissionsResponse - List/Array of multiple permissions
type PermissionsResponse struct {
	Data []PermissionResponseRead `json:"data"`
}

func (o *PermissionsResponse) GetData() []PermissionResponseRead {
	if o == nil {
		return []PermissionResponseRead{}
	}
	return o.Data
}
