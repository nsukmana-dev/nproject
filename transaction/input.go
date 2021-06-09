package transaction

import "nproject/user"

type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
