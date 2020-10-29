package dao

import "github.com/corporateanon/barker/pkg/types"

type BotDao interface {
	Create(bot *types.Bot) (*types.Bot, error)
	Update(bot *types.Bot) (*types.Bot, error)
	Get(ID int64) (*types.Bot, error)
	List(pageRequest *types.PaginatorRequest) ([]types.Bot, *types.PaginatorResponse, error)
	RRTake() (*types.Bot, error)
}
