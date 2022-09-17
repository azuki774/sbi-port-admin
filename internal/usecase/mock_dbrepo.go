package usecase

type MockDBRepo struct{}

func NewMockDBRepo() *MockDBRepo {
	return &MockDBRepo{}
}

func (m *MockDBRepo) CloseDB() (err error){
	return nil
}
