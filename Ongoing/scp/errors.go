// errors.go 상세내용이 각각 필요한지 확인 필요 >> Case만 반영시키고 잘 모르는 부분은 그대로 남겨두어도 될 것 같다.
package scp

import "fmt"

//NotSupportClassic return error for not support classic
func NotSupportClassic(name string) error {
	return fmt.Errorf("%s doesn't support classic", name)
}

//NotSupportVpc return error for not support vpc
//func NotSupportVpc(name string) error {
//	return fmt.Errorf("%s doesn't support vpc", name)
//}

//ErrorRequiredArgOnVpc return error for required on vpc
//func ErrorRequiredArgOnVpc(name string) error {
//	return fmt.Errorf("missing required argument: The argument \"%s\" is required on vpc", name)
//}

//ErrorRequiredArgOnClassic return error for required on classic
//func ErrorRequiredArgOnClassic(name string) error {
//	return fmt.Errorf("missing required argument: The argument \"%s\" is required on classic", name)
//}
