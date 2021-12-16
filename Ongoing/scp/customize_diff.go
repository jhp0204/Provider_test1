//사용 시점 확인 필요, scp향으로 수정은 완료 
package scp
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func scpVpcCommonCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	if diff.HasChange("name") {
		old, new := diff.GetChange("name")
		if len(old.(string)) > 0 {
			return fmt.Errorf("Change 'name' is not support, Please set `name` as a old value = [%s -> %s]", new, old)
		}
	}

	return nil
}
