package transaction

import "bwastartup/user"

type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"` // id from campaign
	User user.User
}
