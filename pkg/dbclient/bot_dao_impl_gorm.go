package dbclient

import (
	"errors"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/pagination"
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type BotDaoImplGorm struct {
	db *gorm.DB
}

func NewBotDaoImplGorm(db *gorm.DB) dao.BotDao {
	return &BotDaoImplGorm{
		db: db,
	}
}

func (dao *BotDaoImplGorm) Create(bot *types.Bot) (*types.Bot, error) {
	botModel := &database.Bot{}
	botModel.FromEntity(bot)

	if err := dao.db.Create(botModel).Error; err != nil {
		return nil, err
	}
	resultingBot := &types.Bot{}
	botModel.ToEntity(resultingBot)
	return resultingBot, nil

}

func (dao *BotDaoImplGorm) Update(bot *types.Bot) (*types.Bot, error) {
	if bot.ID == 0 {
		return nil, errors.New("ID missing")
	}
	botModel := &database.Bot{}
	botModel.ID = bot.ID

	if err := dao.db.First(botModel).Error; err != nil {
		return nil, err
	}

	botModel.FromEntity(bot)

	if err := dao.db.Save(botModel).Error; err != nil {
		return nil, err
	}
	resultingBot := &types.Bot{}
	botModel.ToEntity(resultingBot)
	return resultingBot, nil
}

func (dao *BotDaoImplGorm) Get(ID int64) (*types.Bot, error) {
	botModel := &database.Bot{ID: ID}

	if err := dao.db.First(botModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	resultingBot := &types.Bot{ID: ID}
	botModel.ToEntity(resultingBot)
	return resultingBot, nil
}

func (dao *BotDaoImplGorm) List(pageRequest *types.PaginatorRequest) ([]types.Bot, *types.PaginatorResponse, error) {
	botModelsList := []database.Bot{}
	db := dao.db.Table("bots").Order("created_at DESC")
	resp := pagination.Paging(&pagination.Param{
		DB:    db,
		Page:  int(pageRequest.Page),
		Limit: int(pageRequest.Size),
	}, &botModelsList)

	if err := db.Error; err != nil {
		return nil, nil, err
	}

	botsList := make([]types.Bot, len(botModelsList))
	for i, model := range botModelsList {
		model.ToEntity(&botsList[i])
	}
	return botsList,
		&types.PaginatorResponse{
			Page:  resp.Page,
			Size:  resp.Limit,
			Total: resp.TotalPage,
		},
		nil
}
