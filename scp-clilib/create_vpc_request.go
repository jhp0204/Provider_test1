// VPC Network 관련 API<br/>https://openapi.samsungsdscloud.com/vpc/v1

package vpc

type CreateVpcRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// IPv4 CIDR블록
Ipv4CidrBlock *string `json:"ipv4CidrBlock"`

	// VPC이름
VpcName *string `json:"vpcName,omitempty"`
}
