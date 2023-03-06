package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaings, err := s.repository.FindByUserID(userID)
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

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
