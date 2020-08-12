package repository

type HeartRateRepository struct {
	count uint
}

func NewHeartRateRepository() *HeartRateRepository {
	return &HeartRateRepository{count: 0}
}

func (repo *HeartRateRepository) Get() uint {
	return repo.count
}

func (repo *HeartRateRepository) Set(count uint) {
	repo.count = count
}
