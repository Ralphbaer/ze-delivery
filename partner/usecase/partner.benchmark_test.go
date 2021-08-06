package usecase

import (
	"context"
	"testing"

	e "github.com/Ralphbaer/ze-delivery/partner/entity"
	"github.com/golang/mock/gomock"
)


var result *e.Partner
func BenchmarkCreatePartner(b *testing.B) {
	mockCtrl := gomock.NewController(b)
	defer mockCtrl.Finish()

	var partner *e.Partner
	repo := setupPartnerRepo(mockCtrl)
	uc := setupPartnerUseCaseMocks(mockCtrl, repo)
	for n := 0; n < b.N; n++ {
		partner, _ = uc.Create(context.TODO(), createPartnerInput)
	}
	result = partner
}