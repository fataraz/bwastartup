package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaings, err := s.repository.FindByID(userID)
		if err != nil {
			return campaings, err
		}
		return campaings, nil
	}

	campaings, err := s.repository.FindAll()
	if err != nil {
		return campaings, err
	}

	return campaings, nil

}
