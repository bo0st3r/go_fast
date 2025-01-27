package telemetry

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) GetAll() ([]Entry, error) {
	return service.repository.getAll()
}

func (service *Service) Create(entry Entry) error {
	return service.repository.create(entry)
}
