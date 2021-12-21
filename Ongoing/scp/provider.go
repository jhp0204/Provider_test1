package scp

import (
	"fmt"
	"os"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ScpResources map[string]*schema.Resource
var ScpDataSources map[string]*schema.Resource

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:         schemaMap(),
		DataSourcesMap: DataSourcesMap(),
		ResourcesMap:   ResourcesMap(),
		ConfigureFunc:  providerConfigure,
	}
}

func schemaMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_key": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SCP_ACCESS_KEY", nil),
			Description: descriptions["access_key"],
		},
		"secret_key": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SCP_SECRET_KEY", nil),
			Description: descriptions["secret_key"],
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// Set client
	config := Config{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
	}
	return &providerConfig, nil
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"access_key":  "Access key of scp",
		"secret_key":  "Secret key of scp",
	}
}

//RegisterDataSource Register data sources terraform for SDS CLOUD PLATFORM.
func RegisterDataSource(name string, DataSourceSchema *schema.Resource) {
	if ScpDataSources == nil {
		ScpDataSources = make(map[string]*schema.Resource)
	}
	ScpDataSources[name] = DataSourceSchema
}

//RegisterResource Register resources terraform for SDS CLOUD PLATFORM.
func RegisterResource(name string, resourceSchema *schema.Resource) {
	if ScpResources == nil {
		ScpResources = make(map[string]*schema.Resource)
	}
	ScpResources[name] = resourceSchema
}

//DataSourcesMap This returns a map of all data sources to register with Terraform
func DataSourcesMap() map[string]*schema.Resource {
	return ScpDataSources
}

//ResourcesMap This returns a map of all resources to register with Terraform
func ResourcesMap() map[string]*schema.Resource {
	return ScpResources
}
