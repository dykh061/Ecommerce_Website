package productbusiness

import (
	"OpenMarket/common"
	"fmt"
)

func BuildSKUWithUID(
	sellerId int,
	productId int,
) string {
	uid := common.NewUID(
		common.GenLocalID(), // localID ngẫu nhiên
		common.DbTypeVariant,
		1, // shard
	)

	return fmt.Sprintf(
		"OM-S%d-P%d-%s",
		sellerId,
		productId,
		uid.String(),
	)
}
