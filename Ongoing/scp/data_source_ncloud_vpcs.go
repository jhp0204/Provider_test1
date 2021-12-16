// datasource vpc와 vpcs의 차이가 무엇인지 확인 필요 (구분 목적?)
package scp

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	RegisterDataSource("scp_vpcs", dataSourceScpVpcs())
}

func dataSourceScpVpcs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceScpVpcsRead,

		Schema: map[string]*schema.Schema{
			"vpc_no": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": dataSourceFiltersSchema(),
			"vpcs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(resourceScpVpc()),
			},
		},
	}
}

func dataSourceScpVpcsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*ProviderConfig)

	if !config.SupportVPC {
		return NotSupportClassic("data source `scp_vpcs`")
	}

	resources, err := getVpcListFiltered(d, config)

	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	if err := d.Set("vpcs", resources); err != nil {
		return fmt.Errorf("Error setting vpcs: %s", err)
	}

	return nil
}
