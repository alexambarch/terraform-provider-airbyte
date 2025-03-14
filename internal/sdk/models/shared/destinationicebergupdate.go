// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType string

const (
	DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogTypeGlue DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType = "Glue"
)

func (e DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType) ToPointer() *DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType {
	return &e
}
func (e *DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Glue":
		*e = DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType: %v", v)
	}
}

// DestinationIcebergUpdateGlueCatalog - The GlueCatalog connects to a AWS Glue Catalog
type DestinationIcebergUpdateGlueCatalog struct {
	CatalogType *DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType `default:"Glue" json:"catalog_type"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Database *string `default:"public" json:"database"`
}

func (d DestinationIcebergUpdateGlueCatalog) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIcebergUpdateGlueCatalog) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIcebergUpdateGlueCatalog) GetCatalogType() *DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfig5CatalogType {
	if o == nil {
		return nil
	}
	return o.CatalogType
}

func (o *DestinationIcebergUpdateGlueCatalog) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

type DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType string

const (
	DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogTypeRest DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType = "Rest"
)

func (e DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType) ToPointer() *DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType {
	return &e
}
func (e *DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Rest":
		*e = DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType: %v", v)
	}
}

// DestinationIcebergUpdateRESTCatalog - The RESTCatalog connects to a REST server at the specified URI
type DestinationIcebergUpdateRESTCatalog struct {
	CatalogType    *DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType `default:"Rest" json:"catalog_type"`
	RestURI        string                                                                       `json:"rest_uri"`
	RestCredential *string                                                                      `json:"rest_credential,omitempty"`
	RestToken      *string                                                                      `json:"rest_token,omitempty"`
}

func (d DestinationIcebergUpdateRESTCatalog) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIcebergUpdateRESTCatalog) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIcebergUpdateRESTCatalog) GetCatalogType() *DestinationIcebergUpdateSchemasCatalogConfigIcebergCatalogConfigCatalogType {
	if o == nil {
		return nil
	}
	return o.CatalogType
}

func (o *DestinationIcebergUpdateRESTCatalog) GetRestURI() string {
	if o == nil {
		return ""
	}
	return o.RestURI
}

func (o *DestinationIcebergUpdateRESTCatalog) GetRestCredential() *string {
	if o == nil {
		return nil
	}
	return o.RestCredential
}

func (o *DestinationIcebergUpdateRESTCatalog) GetRestToken() *string {
	if o == nil {
		return nil
	}
	return o.RestToken
}

type DestinationIcebergUpdateSchemasCatalogConfigCatalogType string

const (
	DestinationIcebergUpdateSchemasCatalogConfigCatalogTypeJdbc DestinationIcebergUpdateSchemasCatalogConfigCatalogType = "Jdbc"
)

func (e DestinationIcebergUpdateSchemasCatalogConfigCatalogType) ToPointer() *DestinationIcebergUpdateSchemasCatalogConfigCatalogType {
	return &e
}
func (e *DestinationIcebergUpdateSchemasCatalogConfigCatalogType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Jdbc":
		*e = DestinationIcebergUpdateSchemasCatalogConfigCatalogType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateSchemasCatalogConfigCatalogType: %v", v)
	}
}

// DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase - Using a table in a relational database to manage Iceberg tables through JDBC. Read more <a href="https://iceberg.apache.org/docs/latest/jdbc/">here</a>. Supporting: PostgreSQL
type DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase struct {
	CatalogType *DestinationIcebergUpdateSchemasCatalogConfigCatalogType `default:"Jdbc" json:"catalog_type"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Database *string `default:"public" json:"database"`
	JdbcURL  *string `json:"jdbc_url,omitempty"`
	// Username to use to access the database.
	Username *string `json:"username,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Encrypt data using SSL. When activating SSL, please select one of the connection modes.
	Ssl *bool `default:"false" json:"ssl"`
	// Iceberg catalog metadata tables are written to catalog schema. The usual value for this field is "public".
	CatalogSchema *string `default:"public" json:"catalog_schema"`
}

func (d DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) GetCatalogType() *DestinationIcebergUpdateSchemasCatalogConfigCatalogType {
	if o == nil {
		return nil
	}
	return o.CatalogType
}

func (o *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) GetJdbcURL() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURL
}

func (o *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) GetCatalogSchema() *string {
	if o == nil {
		return nil
	}
	return o.CatalogSchema
}

type DestinationIcebergUpdateSchemasCatalogType string

const (
	DestinationIcebergUpdateSchemasCatalogTypeHadoop DestinationIcebergUpdateSchemasCatalogType = "Hadoop"
)

func (e DestinationIcebergUpdateSchemasCatalogType) ToPointer() *DestinationIcebergUpdateSchemasCatalogType {
	return &e
}
func (e *DestinationIcebergUpdateSchemasCatalogType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Hadoop":
		*e = DestinationIcebergUpdateSchemasCatalogType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateSchemasCatalogType: %v", v)
	}
}

// DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig - A Hadoop catalog doesn’t need to connect to a Hive MetaStore, but can only be used with HDFS or similar file systems that support atomic rename.
type DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig struct {
	CatalogType *DestinationIcebergUpdateSchemasCatalogType `default:"Hadoop" json:"catalog_type"`
	// The default database tables are written to if the source does not specify a namespace. The usual value for this field is "default".
	Database *string `default:"default" json:"database"`
}

func (d DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig) GetCatalogType() *DestinationIcebergUpdateSchemasCatalogType {
	if o == nil {
		return nil
	}
	return o.CatalogType
}

func (o *DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

type DestinationIcebergUpdateCatalogType string

const (
	DestinationIcebergUpdateCatalogTypeHive DestinationIcebergUpdateCatalogType = "Hive"
)

func (e DestinationIcebergUpdateCatalogType) ToPointer() *DestinationIcebergUpdateCatalogType {
	return &e
}
func (e *DestinationIcebergUpdateCatalogType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Hive":
		*e = DestinationIcebergUpdateCatalogType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateCatalogType: %v", v)
	}
}

type DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore struct {
	CatalogType *DestinationIcebergUpdateCatalogType `default:"Hive" json:"catalog_type"`
	// Hive MetaStore thrift server uri of iceberg catalog.
	HiveThriftURI string `json:"hive_thrift_uri"`
	// The default database tables are written to if the source does not specify a namespace. The usual value for this field is "default".
	Database *string `default:"default" json:"database"`
}

func (d DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore) GetCatalogType() *DestinationIcebergUpdateCatalogType {
	if o == nil {
		return nil
	}
	return o.CatalogType
}

func (o *DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore) GetHiveThriftURI() string {
	if o == nil {
		return ""
	}
	return o.HiveThriftURI
}

func (o *DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

type DestinationIcebergUpdateIcebergCatalogConfigType string

const (
	DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore                            DestinationIcebergUpdateIcebergCatalogConfigType = "destination-iceberg-update_HiveCatalog: Use Apache Hive MetaStore"
	DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig DestinationIcebergUpdateIcebergCatalogConfigType = "destination-iceberg-update_HadoopCatalog: Use hierarchical file systems as same as storage config"
	DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateJdbcCatalogUseRelationalDatabase                             DestinationIcebergUpdateIcebergCatalogConfigType = "destination-iceberg-update_JdbcCatalog: Use relational database"
	DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateRESTCatalog                                                  DestinationIcebergUpdateIcebergCatalogConfigType = "destination-iceberg-update_RESTCatalog"
	DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateGlueCatalog                                                  DestinationIcebergUpdateIcebergCatalogConfigType = "destination-iceberg-update_GlueCatalog"
)

// DestinationIcebergUpdateIcebergCatalogConfig - Catalog config of Iceberg.
type DestinationIcebergUpdateIcebergCatalogConfig struct {
	DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore                            *DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore                            `queryParam:"inline"`
	DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig *DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig `queryParam:"inline"`
	DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase                             *DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase                             `queryParam:"inline"`
	DestinationIcebergUpdateRESTCatalog                                                  *DestinationIcebergUpdateRESTCatalog                                                  `queryParam:"inline"`
	DestinationIcebergUpdateGlueCatalog                                                  *DestinationIcebergUpdateGlueCatalog                                                  `queryParam:"inline"`

	Type DestinationIcebergUpdateIcebergCatalogConfigType
}

func CreateDestinationIcebergUpdateIcebergCatalogConfigDestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore(destinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore) DestinationIcebergUpdateIcebergCatalogConfig {
	typ := DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore

	return DestinationIcebergUpdateIcebergCatalogConfig{
		DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore: &destinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore,
		Type: typ,
	}
}

func CreateDestinationIcebergUpdateIcebergCatalogConfigDestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig(destinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig) DestinationIcebergUpdateIcebergCatalogConfig {
	typ := DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig

	return DestinationIcebergUpdateIcebergCatalogConfig{
		DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig: &destinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig,
		Type: typ,
	}
}

func CreateDestinationIcebergUpdateIcebergCatalogConfigDestinationIcebergUpdateJdbcCatalogUseRelationalDatabase(destinationIcebergUpdateJdbcCatalogUseRelationalDatabase DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase) DestinationIcebergUpdateIcebergCatalogConfig {
	typ := DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateJdbcCatalogUseRelationalDatabase

	return DestinationIcebergUpdateIcebergCatalogConfig{
		DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase: &destinationIcebergUpdateJdbcCatalogUseRelationalDatabase,
		Type: typ,
	}
}

func CreateDestinationIcebergUpdateIcebergCatalogConfigDestinationIcebergUpdateRESTCatalog(destinationIcebergUpdateRESTCatalog DestinationIcebergUpdateRESTCatalog) DestinationIcebergUpdateIcebergCatalogConfig {
	typ := DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateRESTCatalog

	return DestinationIcebergUpdateIcebergCatalogConfig{
		DestinationIcebergUpdateRESTCatalog: &destinationIcebergUpdateRESTCatalog,
		Type:                                typ,
	}
}

func CreateDestinationIcebergUpdateIcebergCatalogConfigDestinationIcebergUpdateGlueCatalog(destinationIcebergUpdateGlueCatalog DestinationIcebergUpdateGlueCatalog) DestinationIcebergUpdateIcebergCatalogConfig {
	typ := DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateGlueCatalog

	return DestinationIcebergUpdateIcebergCatalogConfig{
		DestinationIcebergUpdateGlueCatalog: &destinationIcebergUpdateGlueCatalog,
		Type:                                typ,
	}
}

func (u *DestinationIcebergUpdateIcebergCatalogConfig) UnmarshalJSON(data []byte) error {

	var destinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig = DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig{}
	if err := utils.UnmarshalJSON(data, &destinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig, "", true, true); err == nil {
		u.DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig = &destinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig
		u.Type = DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig
		return nil
	}

	var destinationIcebergUpdateGlueCatalog DestinationIcebergUpdateGlueCatalog = DestinationIcebergUpdateGlueCatalog{}
	if err := utils.UnmarshalJSON(data, &destinationIcebergUpdateGlueCatalog, "", true, true); err == nil {
		u.DestinationIcebergUpdateGlueCatalog = &destinationIcebergUpdateGlueCatalog
		u.Type = DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateGlueCatalog
		return nil
	}

	var destinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore = DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore{}
	if err := utils.UnmarshalJSON(data, &destinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore, "", true, true); err == nil {
		u.DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore = &destinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore
		u.Type = DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore
		return nil
	}

	var destinationIcebergUpdateRESTCatalog DestinationIcebergUpdateRESTCatalog = DestinationIcebergUpdateRESTCatalog{}
	if err := utils.UnmarshalJSON(data, &destinationIcebergUpdateRESTCatalog, "", true, true); err == nil {
		u.DestinationIcebergUpdateRESTCatalog = &destinationIcebergUpdateRESTCatalog
		u.Type = DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateRESTCatalog
		return nil
	}

	var destinationIcebergUpdateJdbcCatalogUseRelationalDatabase DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase = DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase{}
	if err := utils.UnmarshalJSON(data, &destinationIcebergUpdateJdbcCatalogUseRelationalDatabase, "", true, true); err == nil {
		u.DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase = &destinationIcebergUpdateJdbcCatalogUseRelationalDatabase
		u.Type = DestinationIcebergUpdateIcebergCatalogConfigTypeDestinationIcebergUpdateJdbcCatalogUseRelationalDatabase
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationIcebergUpdateIcebergCatalogConfig", string(data))
}

func (u DestinationIcebergUpdateIcebergCatalogConfig) MarshalJSON() ([]byte, error) {
	if u.DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore != nil {
		return utils.MarshalJSON(u.DestinationIcebergUpdateHiveCatalogUseApacheHiveMetaStore, "", true)
	}

	if u.DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig != nil {
		return utils.MarshalJSON(u.DestinationIcebergUpdateHadoopCatalogUseHierarchicalFileSystemsAsSameAsStorageConfig, "", true)
	}

	if u.DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase != nil {
		return utils.MarshalJSON(u.DestinationIcebergUpdateJdbcCatalogUseRelationalDatabase, "", true)
	}

	if u.DestinationIcebergUpdateRESTCatalog != nil {
		return utils.MarshalJSON(u.DestinationIcebergUpdateRESTCatalog, "", true)
	}

	if u.DestinationIcebergUpdateGlueCatalog != nil {
		return utils.MarshalJSON(u.DestinationIcebergUpdateGlueCatalog, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationIcebergUpdateIcebergCatalogConfig: all fields are null")
}

type DestinationIcebergUpdateSchemasStorageType string

const (
	DestinationIcebergUpdateSchemasStorageTypeManaged DestinationIcebergUpdateSchemasStorageType = "MANAGED"
)

func (e DestinationIcebergUpdateSchemasStorageType) ToPointer() *DestinationIcebergUpdateSchemasStorageType {
	return &e
}
func (e *DestinationIcebergUpdateSchemasStorageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANAGED":
		*e = DestinationIcebergUpdateSchemasStorageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateSchemasStorageType: %v", v)
	}
}

// DestinationIcebergUpdateServerManaged - Server-managed object storage
type DestinationIcebergUpdateServerManaged struct {
	StorageType *DestinationIcebergUpdateSchemasStorageType `default:"MANAGED" json:"storage_type"`
	// The name of the managed warehouse
	ManagedWarehouseName string `json:"managed_warehouse_name"`
}

func (d DestinationIcebergUpdateServerManaged) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIcebergUpdateServerManaged) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIcebergUpdateServerManaged) GetStorageType() *DestinationIcebergUpdateSchemasStorageType {
	if o == nil {
		return nil
	}
	return o.StorageType
}

func (o *DestinationIcebergUpdateServerManaged) GetManagedWarehouseName() string {
	if o == nil {
		return ""
	}
	return o.ManagedWarehouseName
}

type DestinationIcebergUpdateStorageType string

const (
	DestinationIcebergUpdateStorageTypeS3 DestinationIcebergUpdateStorageType = "S3"
)

func (e DestinationIcebergUpdateStorageType) ToPointer() *DestinationIcebergUpdateStorageType {
	return &e
}
func (e *DestinationIcebergUpdateStorageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3":
		*e = DestinationIcebergUpdateStorageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateStorageType: %v", v)
	}
}

// DestinationIcebergUpdateS3BucketRegion - The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
type DestinationIcebergUpdateS3BucketRegion string

const (
	DestinationIcebergUpdateS3BucketRegionUnknown      DestinationIcebergUpdateS3BucketRegion = ""
	DestinationIcebergUpdateS3BucketRegionAfSouth1     DestinationIcebergUpdateS3BucketRegion = "af-south-1"
	DestinationIcebergUpdateS3BucketRegionApEast1      DestinationIcebergUpdateS3BucketRegion = "ap-east-1"
	DestinationIcebergUpdateS3BucketRegionApNortheast1 DestinationIcebergUpdateS3BucketRegion = "ap-northeast-1"
	DestinationIcebergUpdateS3BucketRegionApNortheast2 DestinationIcebergUpdateS3BucketRegion = "ap-northeast-2"
	DestinationIcebergUpdateS3BucketRegionApNortheast3 DestinationIcebergUpdateS3BucketRegion = "ap-northeast-3"
	DestinationIcebergUpdateS3BucketRegionApSouth1     DestinationIcebergUpdateS3BucketRegion = "ap-south-1"
	DestinationIcebergUpdateS3BucketRegionApSouth2     DestinationIcebergUpdateS3BucketRegion = "ap-south-2"
	DestinationIcebergUpdateS3BucketRegionApSoutheast1 DestinationIcebergUpdateS3BucketRegion = "ap-southeast-1"
	DestinationIcebergUpdateS3BucketRegionApSoutheast2 DestinationIcebergUpdateS3BucketRegion = "ap-southeast-2"
	DestinationIcebergUpdateS3BucketRegionApSoutheast3 DestinationIcebergUpdateS3BucketRegion = "ap-southeast-3"
	DestinationIcebergUpdateS3BucketRegionApSoutheast4 DestinationIcebergUpdateS3BucketRegion = "ap-southeast-4"
	DestinationIcebergUpdateS3BucketRegionCaCentral1   DestinationIcebergUpdateS3BucketRegion = "ca-central-1"
	DestinationIcebergUpdateS3BucketRegionCaWest1      DestinationIcebergUpdateS3BucketRegion = "ca-west-1"
	DestinationIcebergUpdateS3BucketRegionCnNorth1     DestinationIcebergUpdateS3BucketRegion = "cn-north-1"
	DestinationIcebergUpdateS3BucketRegionCnNorthwest1 DestinationIcebergUpdateS3BucketRegion = "cn-northwest-1"
	DestinationIcebergUpdateS3BucketRegionEuCentral1   DestinationIcebergUpdateS3BucketRegion = "eu-central-1"
	DestinationIcebergUpdateS3BucketRegionEuCentral2   DestinationIcebergUpdateS3BucketRegion = "eu-central-2"
	DestinationIcebergUpdateS3BucketRegionEuNorth1     DestinationIcebergUpdateS3BucketRegion = "eu-north-1"
	DestinationIcebergUpdateS3BucketRegionEuSouth1     DestinationIcebergUpdateS3BucketRegion = "eu-south-1"
	DestinationIcebergUpdateS3BucketRegionEuSouth2     DestinationIcebergUpdateS3BucketRegion = "eu-south-2"
	DestinationIcebergUpdateS3BucketRegionEuWest1      DestinationIcebergUpdateS3BucketRegion = "eu-west-1"
	DestinationIcebergUpdateS3BucketRegionEuWest2      DestinationIcebergUpdateS3BucketRegion = "eu-west-2"
	DestinationIcebergUpdateS3BucketRegionEuWest3      DestinationIcebergUpdateS3BucketRegion = "eu-west-3"
	DestinationIcebergUpdateS3BucketRegionIlCentral1   DestinationIcebergUpdateS3BucketRegion = "il-central-1"
	DestinationIcebergUpdateS3BucketRegionMeCentral1   DestinationIcebergUpdateS3BucketRegion = "me-central-1"
	DestinationIcebergUpdateS3BucketRegionMeSouth1     DestinationIcebergUpdateS3BucketRegion = "me-south-1"
	DestinationIcebergUpdateS3BucketRegionSaEast1      DestinationIcebergUpdateS3BucketRegion = "sa-east-1"
	DestinationIcebergUpdateS3BucketRegionUsEast1      DestinationIcebergUpdateS3BucketRegion = "us-east-1"
	DestinationIcebergUpdateS3BucketRegionUsEast2      DestinationIcebergUpdateS3BucketRegion = "us-east-2"
	DestinationIcebergUpdateS3BucketRegionUsGovEast1   DestinationIcebergUpdateS3BucketRegion = "us-gov-east-1"
	DestinationIcebergUpdateS3BucketRegionUsGovWest1   DestinationIcebergUpdateS3BucketRegion = "us-gov-west-1"
	DestinationIcebergUpdateS3BucketRegionUsWest1      DestinationIcebergUpdateS3BucketRegion = "us-west-1"
	DestinationIcebergUpdateS3BucketRegionUsWest2      DestinationIcebergUpdateS3BucketRegion = "us-west-2"
)

func (e DestinationIcebergUpdateS3BucketRegion) ToPointer() *DestinationIcebergUpdateS3BucketRegion {
	return &e
}
func (e *DestinationIcebergUpdateS3BucketRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-south-2":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ap-southeast-3":
		fallthrough
	case "ap-southeast-4":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "ca-west-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-central-2":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-south-2":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "il-central-1":
		fallthrough
	case "me-central-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		*e = DestinationIcebergUpdateS3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateS3BucketRegion: %v", v)
	}
}

// DestinationIcebergUpdateS3 - S3 object storage
type DestinationIcebergUpdateS3 struct {
	StorageType *DestinationIcebergUpdateStorageType `default:"S3" json:"storage_type"`
	// The access key ID to access the S3 bucket. Airbyte requires Read and Write permissions to the given bucket. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>.
	AccessKeyID string `json:"access_key_id"`
	// The corresponding secret to the access key ID. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>
	SecretAccessKey string `json:"secret_access_key"`
	// The Warehouse Uri for Iceberg
	S3WarehouseURI string `json:"s3_warehouse_uri"`
	// The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
	S3BucketRegion *DestinationIcebergUpdateS3BucketRegion `default:"" json:"s3_bucket_region"`
	// Your S3 endpoint url. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/s3.html#:~:text=Service%20endpoints-,Amazon%20S3%20endpoints,-When%20you%20use">here</a>
	S3Endpoint *string `default:"" json:"s3_endpoint"`
	// Use path style access
	S3PathStyleAccess *bool `default:"true" json:"s3_path_style_access"`
}

func (d DestinationIcebergUpdateS3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIcebergUpdateS3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIcebergUpdateS3) GetStorageType() *DestinationIcebergUpdateStorageType {
	if o == nil {
		return nil
	}
	return o.StorageType
}

func (o *DestinationIcebergUpdateS3) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *DestinationIcebergUpdateS3) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

func (o *DestinationIcebergUpdateS3) GetS3WarehouseURI() string {
	if o == nil {
		return ""
	}
	return o.S3WarehouseURI
}

func (o *DestinationIcebergUpdateS3) GetS3BucketRegion() *DestinationIcebergUpdateS3BucketRegion {
	if o == nil {
		return nil
	}
	return o.S3BucketRegion
}

func (o *DestinationIcebergUpdateS3) GetS3Endpoint() *string {
	if o == nil {
		return nil
	}
	return o.S3Endpoint
}

func (o *DestinationIcebergUpdateS3) GetS3PathStyleAccess() *bool {
	if o == nil {
		return nil
	}
	return o.S3PathStyleAccess
}

type DestinationIcebergUpdateStorageConfigType string

const (
	DestinationIcebergUpdateStorageConfigTypeDestinationIcebergUpdateS3            DestinationIcebergUpdateStorageConfigType = "destination-iceberg-update_S3"
	DestinationIcebergUpdateStorageConfigTypeDestinationIcebergUpdateServerManaged DestinationIcebergUpdateStorageConfigType = "destination-iceberg-update_Server-managed"
)

// DestinationIcebergUpdateStorageConfig - Storage config of Iceberg.
type DestinationIcebergUpdateStorageConfig struct {
	DestinationIcebergUpdateS3            *DestinationIcebergUpdateS3            `queryParam:"inline"`
	DestinationIcebergUpdateServerManaged *DestinationIcebergUpdateServerManaged `queryParam:"inline"`

	Type DestinationIcebergUpdateStorageConfigType
}

func CreateDestinationIcebergUpdateStorageConfigDestinationIcebergUpdateS3(destinationIcebergUpdateS3 DestinationIcebergUpdateS3) DestinationIcebergUpdateStorageConfig {
	typ := DestinationIcebergUpdateStorageConfigTypeDestinationIcebergUpdateS3

	return DestinationIcebergUpdateStorageConfig{
		DestinationIcebergUpdateS3: &destinationIcebergUpdateS3,
		Type:                       typ,
	}
}

func CreateDestinationIcebergUpdateStorageConfigDestinationIcebergUpdateServerManaged(destinationIcebergUpdateServerManaged DestinationIcebergUpdateServerManaged) DestinationIcebergUpdateStorageConfig {
	typ := DestinationIcebergUpdateStorageConfigTypeDestinationIcebergUpdateServerManaged

	return DestinationIcebergUpdateStorageConfig{
		DestinationIcebergUpdateServerManaged: &destinationIcebergUpdateServerManaged,
		Type:                                  typ,
	}
}

func (u *DestinationIcebergUpdateStorageConfig) UnmarshalJSON(data []byte) error {

	var destinationIcebergUpdateServerManaged DestinationIcebergUpdateServerManaged = DestinationIcebergUpdateServerManaged{}
	if err := utils.UnmarshalJSON(data, &destinationIcebergUpdateServerManaged, "", true, true); err == nil {
		u.DestinationIcebergUpdateServerManaged = &destinationIcebergUpdateServerManaged
		u.Type = DestinationIcebergUpdateStorageConfigTypeDestinationIcebergUpdateServerManaged
		return nil
	}

	var destinationIcebergUpdateS3 DestinationIcebergUpdateS3 = DestinationIcebergUpdateS3{}
	if err := utils.UnmarshalJSON(data, &destinationIcebergUpdateS3, "", true, true); err == nil {
		u.DestinationIcebergUpdateS3 = &destinationIcebergUpdateS3
		u.Type = DestinationIcebergUpdateStorageConfigTypeDestinationIcebergUpdateS3
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationIcebergUpdateStorageConfig", string(data))
}

func (u DestinationIcebergUpdateStorageConfig) MarshalJSON() ([]byte, error) {
	if u.DestinationIcebergUpdateS3 != nil {
		return utils.MarshalJSON(u.DestinationIcebergUpdateS3, "", true)
	}

	if u.DestinationIcebergUpdateServerManaged != nil {
		return utils.MarshalJSON(u.DestinationIcebergUpdateServerManaged, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationIcebergUpdateStorageConfig: all fields are null")
}

type DestinationIcebergUpdateFileStorageFormat string

const (
	DestinationIcebergUpdateFileStorageFormatParquet DestinationIcebergUpdateFileStorageFormat = "Parquet"
	DestinationIcebergUpdateFileStorageFormatAvro    DestinationIcebergUpdateFileStorageFormat = "Avro"
)

func (e DestinationIcebergUpdateFileStorageFormat) ToPointer() *DestinationIcebergUpdateFileStorageFormat {
	return &e
}
func (e *DestinationIcebergUpdateFileStorageFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Parquet":
		fallthrough
	case "Avro":
		*e = DestinationIcebergUpdateFileStorageFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergUpdateFileStorageFormat: %v", v)
	}
}

// DestinationIcebergUpdateFileFormat - File format of Iceberg storage.
type DestinationIcebergUpdateFileFormat struct {
	Format *DestinationIcebergUpdateFileStorageFormat `default:"Parquet" json:"format"`
	// Iceberg data file flush batch size. Incoming rows write to cache firstly; When cache size reaches this 'batch size', flush into real Iceberg data file.
	FlushBatchSize *int64 `default:"10000" json:"flush_batch_size"`
	// Auto compact data files when stream close
	AutoCompact *bool `default:"false" json:"auto_compact"`
	// Specify the target size of Iceberg data file when performing a compaction action.
	CompactTargetFileSizeInMb *int64 `default:"100" json:"compact_target_file_size_in_mb"`
}

func (d DestinationIcebergUpdateFileFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIcebergUpdateFileFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIcebergUpdateFileFormat) GetFormat() *DestinationIcebergUpdateFileStorageFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *DestinationIcebergUpdateFileFormat) GetFlushBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.FlushBatchSize
}

func (o *DestinationIcebergUpdateFileFormat) GetAutoCompact() *bool {
	if o == nil {
		return nil
	}
	return o.AutoCompact
}

func (o *DestinationIcebergUpdateFileFormat) GetCompactTargetFileSizeInMb() *int64 {
	if o == nil {
		return nil
	}
	return o.CompactTargetFileSizeInMb
}

type DestinationIcebergUpdate struct {
	// Catalog config of Iceberg.
	CatalogConfig DestinationIcebergUpdateIcebergCatalogConfig `json:"catalog_config"`
	// Storage config of Iceberg.
	StorageConfig DestinationIcebergUpdateStorageConfig `json:"storage_config"`
	// File format of Iceberg storage.
	FormatConfig DestinationIcebergUpdateFileFormat `json:"format_config"`
}

func (o *DestinationIcebergUpdate) GetCatalogConfig() DestinationIcebergUpdateIcebergCatalogConfig {
	if o == nil {
		return DestinationIcebergUpdateIcebergCatalogConfig{}
	}
	return o.CatalogConfig
}

func (o *DestinationIcebergUpdate) GetStorageConfig() DestinationIcebergUpdateStorageConfig {
	if o == nil {
		return DestinationIcebergUpdateStorageConfig{}
	}
	return o.StorageConfig
}

func (o *DestinationIcebergUpdate) GetFormatConfig() DestinationIcebergUpdateFileFormat {
	if o == nil {
		return DestinationIcebergUpdateFileFormat{}
	}
	return o.FormatConfig
}
