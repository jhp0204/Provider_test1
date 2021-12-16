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
		"region": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SCP_REGION", nil),
			Description: descriptions["region"],
		},
		"site": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("SCP_SITE", nil),
			Description: descriptions["site"],
		},
// support_vpc의 역할 및 SCP 상품 생성 간 해당 argument가 필요한지 확인 필요
    "support_vpc": {
			Type:        schema.TypeBool,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("SCP_SUPPORT_VPC", nil),
			Description: descriptions["support_vpc"],
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	providerConfig := ProviderConfig{
		SupportVPC: d.Get("support_vpc").(bool),
	}

	// Set site, 각 환경 별 API GW 확인 및 반영하기 (API가 GW를 통해서 들어가는 경우 한정) 
	if site, ok := d.GetOk("site"); ok {
		providerConfig.Site = site.(string)

		switch site {
		case "dev":
			os.Setenv("SCP_API_GW", "개발환경 SCP API GW 입력")
		case "stage":
			os.Setenv("SCP_API_GW", "검증환경 SCP API GW 입력")
		}
	}

	// Fin only supports VPC, SCP에서는 불필요 예상, 확인필요 
	if providerConfig.Site == "fin" {
		providerConfig.SupportVPC = true
	}


	// Set client
	config := Config{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
	}

	if client, err := config.Client(); err != nil {
		return nil, err
	} else {
		providerConfig.Client = client
	}

	// Set region
	if err := setRegionCache(providerConfig.Client, providerConfig.SupportVPC); err != nil {
		return nil, err
	}

	if region, ok := d.GetOk("region"); ok && isValidRegionCode(region.(string)) {
		os.Setenv("SCP_REGION", region.(string))
		providerConfig.RegionCode = region.(string)
		if !providerConfig.SupportVPC {
			providerConfig.RegionNo = *regionCacheByCode[region.(string)].RegionNo
		}
    // Region을 String이 아닌, no로 변경하여 사용하는 것으로 추정된다. SCP에서의 REGION이 단순 String인 경우, if 구문을 삭제하는 것도 고려하기
	} else {
		return nil, fmt.Errorf("no region data for region_code `%s`. please change region_code and try again", region)
	}

	return &providerConfig, nil
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"access_key":  "Access key of scp",
		"secret_key":  "Secret key of scp",
		"region":      "Region of scp",
		"site":        "Site of scp (dev / stg)",
		"support_vpc": "Support VPC platform", // 해당 argument의 정체가 무엇인지 모른다. 
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
