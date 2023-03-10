package campaign

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	CreateCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.GoalAmount = input.GoalAmount
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	// Create Slug
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate) // nama campaign 10 => nama-campaign-10

	// Create Campaign
	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil

}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return campaign, err
	}

	// Check owner the campaign
	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not an owner the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}

	return updateCampaign, nil
}

func (s *service) CreateCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	if input.IsPrimary {
		_, err := s.repository.MarkAllImageAsPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}
	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
	}
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.SaveCampaignImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}
