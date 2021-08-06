package usecase

import (
	"github.com/Ralphbaer/ze-delivery/common"
	"github.com/Ralphbaer/ze-delivery/common/zmock"
	"github.com/Ralphbaer/ze-delivery/partner/gen/mock"
	"github.com/golang/mock/gomock"

	r "github.com/Ralphbaer/ze-delivery/partner/repository"
)

func setupPartnerUseCaseMocks(mockCtrl *gomock.Controller, mockRepo *mock.MockPartnerRepository) PartnerUseCase {
	return PartnerUseCase{
		PartnerRepo:    mockRepo,
	}
}

func setupPartnerUseCase(repo r.PartnerRepository) PartnerUseCase {
	return PartnerUseCase{
		PartnerRepo:    repo,
	}
}

func setupPartnerRepo(mockCtrl *gomock.Controller) *mock.MockPartnerRepository {
	m := mock.NewMockPartnerRepository(mockCtrl)
	m.
		EXPECT().
		Find(gomock.Any(), gomock.Any()).
		Return(defaultPartner, nil).
		AnyTimes()
	m.
		EXPECT().
		Find(gomock.Any(), gomock.Any()).
		Return(nil, common.EntityNotFoundError{}).
		AnyTimes()
	m.
		EXPECT().
		FindNearest(gomock.Any(), float64(-41.59023), float64(-9.75799)).
		Return(nil, ErrNoNearestPartner).
		AnyTimes()
	m.
		EXPECT().
		Save(gomock.Any(), zmock.FieldValueMatcher("Document", "04.433.714/0001-44")).
		Return(nil, ErrPartnerDocumentConflict).
		AnyTimes()
	m.
		EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(&defaultPartner.ID, nil).
		AnyTimes()

	return m
}