package dao

import (
	"errors"

	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BotDaoImplGorm struct {
	db *gorm.DB
}

func NewBotDaoImplGorm(db *gorm.DB) BotDao {
	return &BotDaoImplGorm{
		db: db,
	}
}

func (dao *BotDaoImplGorm) Create(bot *types.Bot) (*types.Bot, error) {
	botModel := &database.Bot{}
	copier.Copy(botModel, bot)
	if err := dao.db.Create(botModel).Error; err != nil {
		return nil, err
	}
	resultingBot := &types.Bot{}
	copier.Copy(resultingBot, botModel.Model)
	copier.Copy(resultingBot, botModel)
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

	copier.Copy(botModel, bot)

	if err := dao.db.Save(botModel).Error; err != nil {
		return nil, err
	}
	resultingBot := &types.Bot{}
	copier.Copy(resultingBot, botModel.Model)
	copier.Copy(resultingBot, botModel)
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
	copier.Copy(resultingBot, botModel)
	return resultingBot, nil

}

func (dao *BotDaoImplGorm) List() ([]types.Bot, error) {
	panic("not implemented")
}
