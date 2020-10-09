package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gotest.tools/assert"
)

func createIntegrationTestServerInvocation() fx.Option {
	return fx.Invoke(func(
		lc fx.Lifecycle,
		r *gin.Engine,
	) {
		lc.Append(fx.Hook{
			OnStart: func(c context.Context) error {
				go http.ListenAndServe(":3000", r)
				return nil
			},
		})
	})
}

func createIntegrationTestInvocation(t *testing.T) fx.Option {
	return fx.Invoke(
		func(
			botDao dao.BotDao,
			userDao dao.UserDao,
			campaignDao dao.CampaignDao,
			deliveryDao dao.DeliveryDao,
		) {

			// #region(collapsed) [create bots]
			t.Run("create bots", func(t *testing.T) {
				bot1Created, err := botDao.Create(&types.Bot{
					Title: "hello_bot",
					Token: "hello",
				})
				assert.NilError(t, err)

				bot2Created, err := botDao.Create(&types.Bot{
					Title: "world_bot",
					Token: "world",
				})
				assert.NilError(t, err)

				assert.DeepEqual(t, bot1Created, &types.Bot{
					ID:    1,
					Title: "hello_bot",
					Token: "hello",
				})

				assert.DeepEqual(t, bot2Created, &types.Bot{
					ID:    2,
					Title: "world_bot",
					Token: "world",
				})

				bot1, err := botDao.Get(1)
				assert.NilError(t, err)

				assert.DeepEqual(t, bot1, &types.Bot{
					ID:    1,
					Title: "hello_bot",
					Token: "hello",
				})

				bot2, err := botDao.Get(2)
				assert.NilError(t, err)

				assert.DeepEqual(t, bot2, &types.Bot{
					ID:    2,
					Title: "world_bot",
					Token: "world",
				})

			})
			// #endregion

			// #region(collapsed) [create users]
			t.Run("create users", func(t *testing.T) {
				user1, err := userDao.Put(&types.User{
					FirstName:  "User",
					LastName:   "One",
					TelegramID: 100,
					BotID:      1,
				})
				assert.NilError(t, err)
				user2, err := userDao.Put(&types.User{
					FirstName:  "User",
					LastName:   "Two",
					TelegramID: 200,
					BotID:      2,
				})
				assert.NilError(t, err)

				assert.DeepEqual(t, user1, &types.User{
					FirstName:  "User",
					LastName:   "One",
					TelegramID: 100,
					BotID:      1,
				})
				assert.DeepEqual(t, user2, &types.User{
					FirstName:  "User",
					LastName:   "Two",
					TelegramID: 200,
					BotID:      2,
				})
			})
			// #endregion

			// #region(collapsed) [update users]
			t.Run("update users", func(t *testing.T) {
				user1, err := userDao.Put(&types.User{
					LastName:   "Um",
					TelegramID: 100,
					BotID:      1,
				})
				assert.NilError(t, err)

				user2, err := userDao.Put(&types.User{
					LastName:   "Dois",
					TelegramID: 200,
					BotID:      2,
				})
				assert.NilError(t, err)

				user1, err = userDao.Get(1, 100)
				assert.NilError(t, err)

				user2, err = userDao.Get(2, 200)
				assert.NilError(t, err)
				assert.DeepEqual(t, user1, &types.User{
					FirstName:  "User",
					LastName:   "Um",
					TelegramID: 100,
					BotID:      1,
				})
				assert.DeepEqual(t, user2, &types.User{
					FirstName:  "User",
					LastName:   "Dois",
					TelegramID: 200,
					BotID:      2,
				})
			})
			// #endregion

			// #region(collapsed) [create campaigns]
			t.Run("create campaigns", func(t *testing.T) {
				campaign1Created, err := campaignDao.Create(&types.Campaign{
					BotID:   1,
					Active:  true,
					Title:   "hello world",
					Message: "hello, user",
				})
				assert.NilError(t, err)
				assert.DeepEqual(t, campaign1Created, &types.Campaign{
					ID:      1,
					BotID:   1,
					Active:  true,
					Title:   "hello world",
					Message: "hello, user",
				})
				campaign1, err := campaignDao.Get(1, 1)
				assert.DeepEqual(t, campaign1, &types.Campaign{
					ID:      1,
					BotID:   1,
					Active:  true,
					Title:   "hello world",
					Message: "hello, user",
				})

				campaign2Created, err := campaignDao.Create(&types.Campaign{
					BotID:   1,
					Active:  true,
					Title:   "foo",
					Message: "bar",
				})
				assert.NilError(t, err)
				assert.DeepEqual(t, campaign2Created, &types.Campaign{
					ID:      2,
					BotID:   1,
					Active:  true,
					Title:   "foo",
					Message: "bar",
				})
				campaign2, err := campaignDao.Get(1, 2)
				assert.DeepEqual(t, campaign2, &types.Campaign{
					ID:      2,
					BotID:   1,
					Active:  true,
					Title:   "foo",
					Message: "bar",
				})
			})
			// #endregion

			// #region(collapsed) [update campaigns]
			t.Run("update campaigns", func(t *testing.T) {
				campaign1Updated, errorWrongBotID := campaignDao.Update(&types.Campaign{
					ID:      1,
					BotID:   1,
					Active:  false,
					Message: "hello",
					Title:   "world",
				})
				assert.NilError(t, errorWrongBotID)

				campaign2Updated, errorWrongBotID := campaignDao.Update(&types.Campaign{
					ID:      2,
					BotID:   1,
					Active:  false,
					Message: "qwerty",
					Title:   "uiop",
				})
				assert.NilError(t, errorWrongBotID)

				assert.DeepEqual(t, campaign1Updated, &types.Campaign{
					ID:      1,
					BotID:   1,
					Active:  false,
					Message: "hello",
					Title:   "world",
				})
				assert.DeepEqual(t, campaign2Updated, &types.Campaign{
					ID:      2,
					BotID:   1,
					Active:  false,
					Message: "qwerty",
					Title:   "uiop",
				})

				_, errorWrongBotID = campaignDao.Update(&types.Campaign{
					ID:      1,
					BotID:   2,
					Active:  false,
					Message: "hello",
					Title:   "world",
				})
				assert.Error(t, errorWrongBotID, "record not found")

				campaign1, err := campaignDao.Get(1, 1)
				assert.NilError(t, err)
				campaign2, err := campaignDao.Get(1, 2)
				assert.NilError(t, err)

				assert.DeepEqual(t, campaign1, &types.Campaign{
					ID:      1,
					BotID:   1,
					Active:  false,
					Message: "hello",
					Title:   "world",
				})
				assert.DeepEqual(t, campaign2, &types.Campaign{
					ID:      2,
					BotID:   1,
					Active:  false,
					Message: "qwerty",
					Title:   "uiop",
				})
			})
			// #endregion

			// #region(collapsed) [take deliveries]
			t.Run("take deliveries", func(t *testing.T) {
				botA, err := botDao.Create(&types.Bot{
					Title: "Bot A",
					Token: "bot:a",
				})
				assert.NilError(t, err)
				botB, err := botDao.Create(&types.Bot{
					Title: "Bot B",
					Token: "bot:b",
				})
				assert.NilError(t, err)

				userA1, err := userDao.Put(&types.User{
					DisplayName: "User A 1",
					TelegramID:  1,
					BotID:       botA.ID,
				})

				assert.NilError(t, err)
				userA2, err := userDao.Put(&types.User{
					DisplayName: "User A 2",
					TelegramID:  2,
					BotID:       botA.ID,
				})

				assert.NilError(t, err)

				userB1, err := userDao.Put(&types.User{
					DisplayName: "User B 1",
					TelegramID:  11,
					BotID:       botB.ID,
				})
				assert.NilError(t, err)
				userB2, err := userDao.Put(&types.User{
					DisplayName: "User B 2",
					TelegramID:  22,
					BotID:       botB.ID,
				})

				assert.NilError(t, err)

				campaignA, err := campaignDao.Create(&types.Campaign{
					BotID:   botA.ID,
					Title:   "Campaign A",
					Message: "Campaign A",
				})
				assert.NilError(t, err)

				campaignB, err := campaignDao.Create(&types.Campaign{
					BotID:   botB.ID,
					Title:   "Campaign B",
					Message: "Campaign B",
				})
				assert.NilError(t, err)

				//--------

				deliveryA1, dUserA1, err := deliveryDao.Take(botA.ID, campaignA.ID)
				assert.NilError(t, err)
				assert.DeepEqual(t, deliveryA1, &types.Delivery{
					BotID:      botA.ID,
					CampaignID: campaignA.ID,
					State:      types.DeliveryStateProgress,
					TelegramID: userA1.TelegramID,
				})
				assert.DeepEqual(t, dUserA1, userA1)

				deliveryA2, dUserA2, err := deliveryDao.Take(botA.ID, campaignA.ID)
				assert.NilError(t, err)
				assert.DeepEqual(t, deliveryA2, &types.Delivery{
					BotID:      botA.ID,
					CampaignID: campaignA.ID,
					State:      types.DeliveryStateProgress,
					TelegramID: userA2.TelegramID,
				})
				assert.DeepEqual(t, dUserA2, userA2)

				deliveryA3, dUserA3, err := deliveryDao.Take(botA.ID, campaignA.ID)
				assert.NilError(t, err)
				assert.Assert(t, deliveryA3 == nil)
				assert.Assert(t, dUserA3 == nil)

				//--------

				deliveryB1, dUserB1, err := deliveryDao.Take(botB.ID, campaignB.ID)
				assert.NilError(t, err)
				assert.DeepEqual(t, deliveryB1, &types.Delivery{
					BotID:      botB.ID,
					CampaignID: campaignB.ID,
					State:      types.DeliveryStateProgress,
					TelegramID: userB1.TelegramID,
				})
				assert.DeepEqual(t, dUserB1, userB1)

				deliveryB2, dUserB2, err := deliveryDao.Take(botB.ID, campaignB.ID)
				assert.NilError(t, err)
				assert.DeepEqual(t, deliveryB2, &types.Delivery{
					BotID:      botB.ID,
					CampaignID: campaignB.ID,
					State:      types.DeliveryStateProgress,
					TelegramID: userB2.TelegramID,
				})
				assert.DeepEqual(t, dUserB2, userB2)

				deliveryB3, dUserB3, err := deliveryDao.Take(botB.ID, campaignB.ID)
				assert.NilError(t, err)
				assert.Assert(t, deliveryB3 == nil)
				assert.Assert(t, dUserB3 == nil)

				t.Run("campaign does not belong to a bot", func(t *testing.T) {
					wrongDelivery, wrongUser, _ := deliveryDao.Take(botA.ID, campaignB.ID)
					//Error depends on an implementation.
					//Gorm implementation does not return error.
					//Resty implementation returns it, because campaign is checked against bot in HTTP request middleware.
					assert.Assert(t, wrongDelivery == nil)
					assert.Assert(t, wrongUser == nil)
				})

				t.Run("update deliveries", func(t *testing.T) {
					err := deliveryDao.SetState(deliveryA1, types.DeliveryStateSuccess)
					assert.NilError(t, err)

					deliveryA1UpdatedState, err := deliveryDao.GetState(deliveryA1)
					assert.Assert(t, deliveryA1UpdatedState == types.DeliveryStateSuccess)

					deliveryA2UnchangedState, err := deliveryDao.GetState(deliveryA2)
					assert.Assert(t, deliveryA2UnchangedState == types.DeliveryStateProgress)

					err = deliveryDao.SetState(deliveryA2, types.DeliveryStateFail)
					assert.NilError(t, err)

					deliveryA2UpdatedState, err := deliveryDao.GetState(deliveryA2)
					assert.Assert(t, deliveryA2UpdatedState == types.DeliveryStateFail)
				})
			})
			// #endregion
		},
	)
}
