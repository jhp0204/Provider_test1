// !!! ncloud sdk 내용과 관련된 내용 추후 수정 필요!!!
package scp

// 대상 csp의 sdk를 import한다. < 이후 scp 향으로 수정 필요 
import (
	"encoding/json"
	"fmt"
	"log"

	"strings"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
)

// 정체를 정확히 모르고, 굳이 수정할 필요가 없는 것 같아서 일단은 유지
const (
	ApiErrorAuthorityParameter = "800"
	ApiErrorUnknown            = "1300"

	ApiErrorObjectInOperation                            = "25013"
	ApiErrorPortForwardingObjectInOperation              = "25033"
	ApiErrorServerObjectInOperation                      = "23006" // Unable to request server termination and creation simultaneously
	ApiErrorServerObjectInOperation2                     = "25017"
	ApiErrorPreviousServersHaveNotBeenEntirelyTerminated = "23003"

	ApiErrorDetachingMountedStorage = "24002"

	ApiErrorAcgCantChangeSameTime           = "1007009"
	ApiErrorNetworkAclCantAccessaApropriate = "1011002"
	ApiErrorNetworkAclRuleChangeIngRules    = "1012005"

	ApiErrorASGIsUsingPolicyOrLaunchConfiguration      = "50150" // This is returned when you cannot delete a launch configuration, scaling policy, or auto scaling group because it is being used.
	ApiErrorASGScalingIsActive                         = "50160" // You cannot request actions while there are scaling activities in progress for that group.
	ApiErrorASGIsUsingPolicyOrLaunchConfigurationOnVpc = "1250700"
)

const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

// ★추가할 Response 확인 >> SCP 기준 ProjectId, requestId, resourceId로 차이 有  >> 
type CommonResponse struct {
	RequestId     *string `json:"requestId,omitempty"`
	projectId    *string `json:"projectId,omitempty"`
	resourceId *string `json:"resourceId,omitempty"`
}

//정체 ? 
type CommonCode struct {
	Code     *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}

//CommonError response error body  >> SCP 기준 Error 시, response Error body 확인 필요
type CommonError struct {
	ReturnCode    string
	ReturnMessage string
}

// Response 내용 확인 필요
func logErrorResponse(tag string, err error, args interface{}) {
	param, _ := json.Marshal(args)
	log.Printf("[ERROR] %s error params=%s, err=%s", tag, param, err)
}

func logCommonRequest(tag string, args interface{}) {
	param, _ := json.Marshal(args)
	log.Printf("[INFO] %s params=%s", tag, param)
}

func logResponse(tag string, args interface{}) {
	resp, _ := json.Marshal(args)
	log.Printf("[INFO] %s response=%s", tag, resp)
}

func logCommonResponse(tag string, commonResponse *CommonResponse, logs ...string) {
	result := fmt.Sprintf("RequestID: %s, ReturnCode: %s, ReturnMessage: %s", ncloud.StringValue(commonResponse.RequestId), ncloud.StringValue(commonResponse.ReturnCode), ncloud.StringValue(commonResponse.ReturnMessage))
	log.Printf("[INFO] %s success response=%s %s", tag, result, strings.Join(logs, " "))
}

func isRetryableErr(commResp *CommonResponse, code []string) bool {
	for _, c := range code {
		if commResp != nil && commResp.ReturnCode != nil && ncloud.StringValue(commResp.ReturnCode) == c {
			return true
		}
	}

	return false
}

func containsInStringList(str string, s []string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
