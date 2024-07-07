package usecase

import "github.com/muizu555/investment/src/repository"

func GetTradeCount(userID string) (int, error) {
	return repository.GetTradeCountByUserID(userID)
}
