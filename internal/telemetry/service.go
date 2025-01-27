package telemetry

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) GetHighestValuePerMetric() ([]Metric, error) {
	return service.repository.getHighestValuePerMetric()
}

func (service *Service) GetAll() ([]Metric, error) {
	return service.repository.getAll()
}

func (service *Service) Create(entry Metric) error {
	return service.repository.create(entry)
}
