package util

import (
	"github.com/expert-pancake/service/purchasing/impl/client"
)

func WarehouseApiToMap(data []client.Warehouse) map[string]client.Warehouse {
	dataMap := map[string]client.Warehouse{}
	for _, v := range data {
		dataMap[v.WarehouseId] = v
	}

	return dataMap
}
