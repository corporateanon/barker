package client

import (
	"strconv"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/go-resty/resty/v2"
)

type CampaignDaoImplResty struct {
	resty *resty.Client
}

func NewCampaignDaoImplResty(resty *resty.Client) dao.CampaignDao {
	return &CampaignDaoImplResty{
		resty: resty,
	}
}

func (dao *CampaignDaoImplResty) Create(campaign *types.Campaign) (*types.Campaign, error) {
	resultWrapper := &struct{ Data *types.Campaign }{Data: &types.Campaign{}}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetBody(campaign).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"BotID": strconv.FormatInt(campaign.BotID, 10),
		}).
		Post("/bot/{BotID}/campaign")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}

func (dao *CampaignDaoImplResty) Update(campaign *types.Campaign) (*types.Campaign, error) {
	resultWrapper := &struct{ Data *types.Campaign }{Data: &types.Campaign{}}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetBody(campaign).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"BotID":      strconv.FormatInt(campaign.BotID, 10),
			"CampaignID": strconv.FormatInt(campaign.ID, 10),
		}).
		Put("/bot/{BotID}/campaign/{CampaignID}")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}

func (dao *CampaignDaoImplResty) Get(botID int64, ID int64) (*types.Campaign, error) {
	resultWrapper := &struct{ Data *types.Campaign }{Data: &types.Campaign{}}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"BotID":      strconv.FormatInt(botID, 10),
			"CampaignID": strconv.FormatInt(ID, 10),
		}).
		Get("/bot/{BotID}/campaign/{CampaignID}")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}

func (dao *CampaignDaoImplResty) List(botID int64, pageRequest *types.PaginatorRequest) ([]types.Campaign, *types.PaginatorResponse, error) {
	resultWrapper := &struct {
		Data   []types.Campaign
		Paging *types.PaginatorResponse
	}{}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		SetQueryParams(pageRequest.ToMap()).
		SetPathParams(map[string]string{
			"BotID": strconv.FormatInt(botID, 10),
		}).
		Get("/bot/{BotID}/campaign")
	if err != nil {
		return nil, nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, resultWrapper.Paging, nil
}

func (dao *CampaignDaoImplResty) GetAggregatedStatistics(botID int64, campaignID int64) (*types.CampaignAggregatedStatistics, error) {
	resultWrapper := &struct {
		Data *types.CampaignAggregatedStatistics
	}{Data: &types.CampaignAggregatedStatistics{}}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"BotID":      strconv.FormatInt(botID, 10),
			"CampaignID": strconv.FormatInt(campaignID, 10),
		}).
		Get("/bot/{BotID}/campaign/{CampaignID}/aggregatedStatistics")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil

}
