package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/Ralphbaer/ze-delivery/common"
	e "github.com/Ralphbaer/ze-delivery/partner/entity"
	r "github.com/Ralphbaer/ze-delivery/partner/repository"
	"github.com/golang/mock/gomock"
)

func TestPartnerUseCase_GetNearestPartner(t *testing.T) {
	// Check if short flag is present, if so, skip test.
	if !testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	
	partnerRepo := r.NewPartnerMongoRepository(&common.MongoConnection{
		ConnectionString: "mongodb://localhost:27017/partner",
		Verbose: true,
	})

	type fields struct {
		UseCase PartnerUseCase
	}
	type args struct {
		ctx context.Context
		Query r.PartnerQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *e.Partner
		wantErr bool
	}{
		{
			name: "ShouldPassOne",
			fields: fields{
				UseCase: setupPartnerUseCase(partnerRepo),
			},
			args: args{
				Query: r.PartnerQuery{
					Longitude: -43.909345, 
					Latitude:  -19.924216,
				},
				ctx: context.TODO(),
			},
			want: func() *e.Partner {
				p, err := partnerRepo.FindByDocument(context.TODO(), "16.572.053/0001-17")
				if err != nil {
					t.Errorf("PartnerUseCase.GetNearestPartner().FindByDocument error = %v", err)
					return nil
				}
				return p
			}(),
			wantErr: false,
		},
		{
			name: "ShouldPassTwo",
			fields: fields{
				UseCase: setupPartnerUseCase(partnerRepo),
			},
			args: args{
				Query: r.PartnerQuery{
					Longitude: -46.63118,
					Latitude:  -23.64707,
				},
				ctx: context.TODO(),
			},
			want: func() *e.Partner {
				p, err := partnerRepo.FindByDocument(context.TODO(), "07.710.066/0001-14")
				if err != nil {
					t.Errorf("PartnerUseCase.GetNearestPartner().FindByDocument error = %v", err)
					return nil
				}
				return p
			}(),
			wantErr: false,
		},
		{
			name: "ShouldNotFound",
			fields: fields{
				UseCase: setupPartnerUseCase(partnerRepo),
			},
			args: args{
				Query: r.PartnerQuery{
					Longitude: -38.59023,
					Latitude:  -9.75799,
				},
				ctx: context.TODO(),
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "ShouldNotFoundTwo",
			fields: fields{
				UseCase: setupPartnerUseCase(partnerRepo),
			},
			args: args{
				Query: r.PartnerQuery{
					Longitude:  -41.59023,
					Latitude:  -9.75799,
				},
				ctx: context.TODO(),
			},
			want: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := tt.fields.UseCase
			got, err := uc.GetNearestPartner(tt.args.ctx, &tt.args.Query)
			if (err != nil) != tt.wantErr {
				t.Errorf("PartnerUseCase.GetNearestPartner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartnerUseCase.GetNearestPartner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartnerUseCase_CreatePartner(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := setupPartnerRepo(mockCtrl)
	uc := setupPartnerUseCaseMocks(mockCtrl, repo)
	got, err := uc.Create(context.TODO(), createPartnerInput)
	if err != nil {
		t.Error(err)
		return
	}

	m := struct {
		want *e.Partner
		got  *e.Partner
	}{
		&e.Partner{
			ID:           "1",
			TradingName:  createPartnerInput.TradingName,
			OwnerName:    createPartnerInput.OwnerName,
			Document:     createPartnerInput.Document,
			CoverageArea: (*e.CoverageArea)(createPartnerInput.CoverageArea),
			Address:      (*e.Address)(createPartnerInput.Address),
		},
		&e.Partner{
			ID:          got.ID,
			TradingName: got.TradingName,
			OwnerName:   got.OwnerName,
			Document:    got.Document,
			CoverageArea: &e.CoverageArea{
				Type:        got.CoverageArea.Type,
				Coordinates: got.CoverageArea.Coordinates,
			},
			Address: &e.Address{
				Type:        got.Address.Type,
				Coordinates: got.Address.Coordinates,
			},
		},
	}

	if !reflect.DeepEqual(m.want, m.got) {
		t.Error("Got:", m.got, "Want:", m.want)
		return
	}
}


func TestPartnerUseCase_GetNearestPartner_ErrNoNearestPartner(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := setupPartnerRepo(mockCtrl)
	uc := setupPartnerUseCaseMocks(mockCtrl, repo)
	_, got := uc.GetNearestPartner(context.TODO(), &r.PartnerQuery{
		Longitude:  -41.59023,
		Latitude:  -9.75799,
	})
	if got == nil {
		t.Error("Got:", got, ",want: not nil")
		return
	}

	want := ErrNoNearestPartner
	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Error("Got:", got, ",want:", want)
		return
	}
}

func TestPartnerUseCase_CreatePartner_ErrPartnerDocumentConflict(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := setupPartnerRepo(mockCtrl)
	uc := setupPartnerUseCaseMocks(mockCtrl, repo)
	_, got := uc.Create(context.TODO(), createPartnerInputWithConflict)
	if got == nil {
		t.Error("Got:", got, ",want: not nil")
		return
	}

	want := ErrPartnerDocumentConflict
	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Error("Got:", got, ",want:", want)
		return
	}
}

func TestPartnerUseCase_Get(t *testing.T) {
	type fields struct {
		PartnerRepo r.PartnerRepository
	}
	type args struct {
		ctx context.Context
		ID  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *e.Partner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &PartnerUseCase{
				PartnerRepo: tt.fields.PartnerRepo,
			}
			got, err := uc.Get(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PartnerUseCase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartnerUseCase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
