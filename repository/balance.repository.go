package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/mr687/privy-be-test-go/entity"
	"gitlab.com/mr687/privy-be-test-go/service"
	"gorm.io/gorm"
)

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) *balanceRepository {
	return &balanceRepository{db}
}

func (repo *balanceRepository) GetCurrentBalance(userId uint64) (*entity.UserBalance, error) {
	var balance *entity.UserBalance
	err := repo.db.Table("user_balances").Where("user_id = ?", userId).Last(&balance).Error
	if err != nil {
		return entity.EmptyUserBalance(userId), err
	}

	return balance, nil
}

func (repo *balanceRepository) SaveBalance(c *gin.Context, newBalance *entity.UserBalance) error {
	tx := repo.db.Begin()

	if err := tx.Save(&newBalance).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := repo.SaveBalanceHistory(tx, c, newBalance); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (repo *balanceRepository) SaveBalanceHistory(tx *gorm.DB, c *gin.Context, newBalance *entity.UserBalance) error {
	geoInfo := c.MustGet("geoip").(*service.GeoIp)

	balanceHistory := entity.UserBalanceHistory{
		UserBalanceId: newBalance.Id,
		BalanceBefore: newBalance.BalanceAchieve - newBalance.Balance,
		BalanceAfter:  newBalance.BalanceAchieve,
		Ip:            geoInfo.Ip,
		Location:      fmt.Sprintf("%s, %s", geoInfo.City, geoInfo.CountryName),
		UserAgent:     c.GetHeader("User-Agent"),
		Type:          "debit",
		Author:        "System",
		Activtiy:      "Add balance",
	}
	return tx.Save(&balanceHistory).Error
}
