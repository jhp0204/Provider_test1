//여기서는 생성된 vpc관련 작업으로 확인(datasource, get api)
// vpc create만 진행할거면 data_source를 굳이 만들 필요가 없지 않은지..
package scp

//scp sdk 개발 후, 해당 서비스로 package 수정 필요, 내용은 선 수정 진행 

import (
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	RegisterDataSource("scp_vpc", dataSourceNcloudVpc())
}

//생성된 vpc가 가지는 field가 무엇인지 확인 필요 (ncloud 기준: id, name)
func dataSourceScpVpc() *schema.Resource {
	fieldMap := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"filter": dataSourceFiltersSchema(),
	}

	return GetSingularDataSourceItemSchema(resourceScpVpc(), fieldMap, dataSourceScpVpcRead)
}

func dataSourceScpVpcRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*ProviderConfig)

	if !config.SupportVPC {
		return NotSupportClassic("data source `scp_vpc`")
	}

	resources, err := getVpcListFiltered(d, config)

	if err != nil {
		return err
	}

	if err := validateOneResult(len(resources)); err != nil {
		return err
	}

	SetSingularResourceDataFromMap(d, resources[0])

	return nil
}

func getVpcListFiltered(d *schema.ResourceData, config *ProviderConfig) ([]map[string]interface{}, error) {
	reqParams := &vpc.GetVpcListRequest{
		RegionCode: &config.RegionCode,
	}

	if v, ok := d.GetOk("name"); ok {
		reqParams.VpcName = scp.String(v.(string))
	}

	if v, ok := d.GetOk("id"); ok {
		reqParams.VpcNoList = []*string{scp.String(v.(string))}
	}

	logCommonRequest("GetVpcList", reqParams)
	resp, err := config.Client.vpc.V2Api.GetVpcList(reqParams)

	if err != nil {
		logErrorResponse("GetVpcList", err, reqParams)
		return nil, err
	}
	logResponse("GetVpcList", resp)

	resources := []map[string]interface{}{}

	for _, r := range resp.VpcList {
		instance := map[string]interface{}{
			"id":              *r.VpcNo,
			"vpc_no":          *r.VpcNo,
			"name":            *r.VpcName,
			"ipv4_cidr_block": *r.Ipv4CidrBlock,
		}

		resources = append(resources, instance)
	}

	if f, ok := d.GetOk("filter"); ok {
		resources = ApplyFilters(f.(*schema.Set), resources, resourceNcloudVpc().Schema)
	}

	return resources, nil
}
