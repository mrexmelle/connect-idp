package tenure

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
)

type Service struct {
	Config           *config.Config
	TenureRepository *Repository
}

func NewService(
	cfg *config.Config,
	r *Repository) *Service {
	return &Service{
		Config:           cfg,
		TenureRepository: r,
	}
}

func (s *Service) Create(request Entity) SingleResponseDto {
	result, err := s.TenureRepository.Create(request)
	return SingleResponseDto{
		Tenure: result,
		Status: mapper.ToStatus(err),
	}
}

func (s *Service) RetrieveByEhid(
	ehid string,
) MultipleResponseDto {
	result, err := s.TenureRepository.FindByEhid(ehid)
	return MultipleResponseDto{
		Tenures: result,
		Status:  mapper.ToStatus(err),
	}
}

func (s *Service) UpdateEndDateById(
	req PatchRequestDto,
	id int,
) PatchResponseDto {
	err := s.TenureRepository.UpdateEndDateById(
		req.Value,
		id,
	)
	return PatchResponseDto{
		Status: mapper.ToStatus(err),
	}
}
