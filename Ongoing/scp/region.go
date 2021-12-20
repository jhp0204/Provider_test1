//NCP의 경우, Region 관련 내용은 server sdk에서 가져오기 떄문에 vserver,server sdk import 
//무수정, ctrl+c, ctrl+v
//region code 확인 필요 
package scp

import (
	"fmt"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vserver"
	"os"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/server"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Region struct {
	RegionNo   *string `json:"regionNo,omitempty"`
	RegionCode *string `json:"regionCode,omitempty"`
	RegionName *string `json:"regionName,omitempty"`
}

var regionCacheByCode = make(map[string]Region)

func parseRegionNoParameter(d *schema.ResourceData) (*string, error) {
	if regionCode, regionCodeOk := d.GetOk("region"); regionCodeOk {
		regionNo := getRegionNoByCode(regionCode.(string))
		if regionNo == nil {
			return nil, fmt.Errorf("no region data for region_code `%s`. please change region_code and try again", regionCode.(string))
		}
		return regionNo, nil
	}

	// provider region
	if regionCode := os.Getenv("NCLOUD_REGION"); regionCode != "" {
		regionNo := getRegionNoByCode(regionCode)
		if regionNo == nil {
			return nil, fmt.Errorf("no region data for region_code `%s`. please change region_code and try again", regionCode)
		}
		return regionNo, nil
	}

	return nil, nil
}

func parseRegionCodeParameter(client *NcloudAPIClient, d *schema.ResourceData) (*string, error) {
	if regionCode, regionCodeOk := d.GetOk("region"); regionCodeOk {
		region, err := getRegionByCode(client, regionCode.(string))
		if region == nil || err != nil {
			return nil, fmt.Errorf("no region data for region_code `%s`. please change region_code and try again", regionCode.(string))
		}
		return region.RegionCode, nil
	}

	// provider region
	if regionCode := os.Getenv("NCLOUD_REGION"); regionCode != "" {
		region, err := getRegionByCode(client, regionCode)
		if region == nil || err != nil {
			return nil, fmt.Errorf("no region data for region_code `%s`. please change region_code and try again", regionCode)
		}
		return region.RegionCode, nil
	}

	return nil, nil
}

func getRegionNoByCode(code string) *string {
	if region, ok := regionCacheByCode[code]; ok {
		return region.RegionNo
	}
	return nil
}

func getRegionByCode(client *NcloudAPIClient, code string) (*server.Region, error) {
	resp, err := client.server.V2Api.GetRegionList(&server.GetRegionListRequest{})
	if err != nil {
		return nil, err
	}
	regionList := resp.RegionList

	var filteredRegion *server.Region
	for _, region := range regionList {
		if code == *region.RegionCode {
			filteredRegion = region
			break
		}
	}

	return filteredRegion, nil
}

func setRegionCache(client *NcloudAPIClient, supportVPC bool) error {
	var regionList []*Region
	var err error
	if supportVPC {
		regionList, err = getVpcRegionList(client)
	} else {
		regionList, err = getClassicRegionList(client)
	}

	if err != nil {
		return err
	}

	for _, r := range regionList {
		region := Region{
			RegionCode: r.RegionCode,
			RegionName: r.RegionName,
		}
		if !supportVPC {
			region.RegionNo = r.RegionNo
		}

		regionCacheByCode[*region.RegionCode] = region
	}

	return nil
}

func getClassicRegionList(client *NcloudAPIClient) ([]*Region, error) {
	resp, err := client.server.V2Api.GetRegionList(&server.GetRegionListRequest{})
	if err != nil {
		return nil, err
	}

	var regionList []*Region
	for _, r := range resp.RegionList {
		region := &Region{
			RegionNo:   r.RegionNo,
			RegionCode: r.RegionCode,
			RegionName: r.RegionName,
		}
		regionList = append(regionList, region)
	}

	return regionList, nil
}

func getVpcRegionList(client *NcloudAPIClient) ([]*Region, error) {
	resp, err := client.vserver.V2Api.GetRegionList(&vserver.GetRegionListRequest{})
	if err != nil {
		return nil, err
	}
	var regionList []*Region
	for _, r := range resp.RegionList {
		region := &Region{
			RegionCode: r.RegionCode,
			RegionName: r.RegionName,
		}
		regionList = append(regionList, region)
	}

	return regionList, nil
}

func isValidRegionCode(code string) bool {
	_, ok := regionCacheByCode[code]
	return ok
}