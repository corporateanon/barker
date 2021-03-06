package main

import (
	"os"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	converter := typescriptify.New().
		Add(types.Bot{}).
		Add(types.Campaign{}).
		Add(types.CampaignAggregatedStatistics{}).
		Add(types.User{}).
		Add(types.Delivery{}).
		Add(types.PaginatorRequest{}).
		Add(types.PaginatorResponse{}).
		AddEnum(types.AllDeliveryStates).
		Add(dao.DeliveryTakeResult{})

	converter.CreateInterface = true
	converter.BackupDir = os.TempDir()

	err := converter.ConvertToFile("ts/src/types.ts")
	if err != nil {
		panic(err.Error())
	}

}
