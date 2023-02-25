package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
)

func AccountToAccountResponse(account *entity.Account) *response.Account {
	r := &response.Account{
		Id:        account.Id,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}

	return r
}

func AccountsToAccountResponses(accounts *[]entity.Account) *[]response.Account {
	rs := make([]response.Account, 0)

	for _, account := range *accounts {
		rs = append(rs, *AccountToAccountResponse(&account))
	}

	return &rs
}
