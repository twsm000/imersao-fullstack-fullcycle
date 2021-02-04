package grpc

import (
	"context"

	"github.com/twsm000/imersao-fullstack-fullcycle/application/grpc/pb"
	"github.com/twsm000/imersao-fullstack-fullcycle/application/usecase"
)

// NewPixGrpcService ...
func NewPixGrpcService(usecase *usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{UseCase: usecase}
}

// PixGrpcService ...
type PixGrpcService struct {
	UseCase *usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

// RegisterPixKey ...
func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	pix, err := p.UseCase.RegisterKey(in.Key, in.Kind, in.AccountId)
	if err != nil {
		pkc := &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}
		return pkc, err
	}

	pkc := &pb.PixKeyCreatedResult{
		Id:     pix.ID,
		Status: "created",
	}
	return pkc, nil
}

// Find ...
func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pix, err := p.UseCase.FindKeyByKind(in.Key, in.Kind)
	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	pki := &pb.PixKeyInfo{
		Id:   pix.ID,
		Kind: pix.Kind,
		Key:  pix.Key,
		Account: &pb.Account{
			AccountId:     pix.Account.ID,
			AccountNumber: pix.Account.Number,
			BankId:        pix.Account.Bank.ID,
			BankName:      pix.Account.Bank.Name,
			OwnerName:     pix.Account.OwnerName,
			CreatedAt:     pix.Account.CreatedAt.String(),
		},
		CreatedAt: pix.CreatedAt.String(),
	}
	return pki, nil
}
