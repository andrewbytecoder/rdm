package services

import "sync"

type StorageService struct {
}

var s *StorageService
var onceStorage sync.Once

func Storage() *StorageService {
	if s == nil {
		onceStorage.Do(func() {
			s = &StorageService{}
		})
	}
	return s
}
