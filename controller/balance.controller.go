package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mr687/privy-be-test-go/dto"
	"gitlab.com/mr687/privy-be-test-go/entity"
	"gitlab.com/mr687/privy-be-test-go/repository"
	"gitlab.com/mr687/privy-be-test-go/response"
	"gitlab.com/mr687/privy-be-test-go/service"
)

func (s *Server) GetBalance(c *gin.Context) {
	repo := repository.NewBalanceRepository(s.DB)

	userId, _ := service.GetTokenId(c)
	if userId == 0 {
		response.Unauthorized(c)
		return
	}

	balance, err := repo.GetCurrentBalance(userId)
	var amount int64
	if err != nil {
		amount = 0
	}

	if balance == nil {
		balance = entity.EmptyUserBalance(userId)
	}

	amount = balance.BalanceAchieve
	res := gin.H{"balance": amount}
	response.Ok(c, "Get current balance successfully", res)
}

func (s *Server) AddBalance(c *gin.Context) {
	var request dto.AddBalanceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ValidationError(c, err)
		return
	}

	repo := repository.NewBalanceRepository(s.DB)
	userId, _ := service.GetTokenId(c)
	if userId == 0 {
		response.Unauthorized(c)
		return
	}

	var newBalance *entity.UserBalance
	currentBalance, err := repo.GetCurrentBalance(userId)
	if err != nil {
		newBalance = entity.EmptyUserBalance(userId)
	} else {
		newBalance = currentBalance
	}

	newBalance.Balance = request.Amount
	newBalance.BalanceAchieve += request.Amount

	err = repo.SaveBalance(c, newBalance)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Ok(c, "Add balance successfully", nil)
}
