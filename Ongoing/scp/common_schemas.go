// ncloud의 zone과 관련된 schema의 용어들을 선언, 정의하는 문서로 판단, ncloud의 sdk를 사용하지는 않는다. 추후 scp zone, region 관련정보 확인 후 수정 필요
package scp

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// zone관련하여 정확하게 어떤 내용들이 필요한지, 특히 region_no 등이 scp에서도 사용되는 argument인지를 확인, 이후 수정 필요 
var zoneSchemaResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"zone_no": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"zone_code": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"zone_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"zone_description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"region_no": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"region_code": {
			Type:     schema.TypeString,
			Computed: true,
		},
	},
}

var regionSchemaResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"region_no": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"region_code": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"region_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	},
}
