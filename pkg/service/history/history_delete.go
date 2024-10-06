package history

import "go-cli-mgt/pkg/store/repository"

func DeleteHistoryById(id uint64) error {
	return repository.GetSingleton().DeleteHistoryById(id)
}
