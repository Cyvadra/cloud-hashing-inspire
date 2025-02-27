// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"
	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/user-special-reduction" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserSpecialReduction(ctx context.Context, in *npool.CreateUserSpecialReductionRequest) (*npool.CreateUserSpecialReductionResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create coupon pool error: %w", err)
		return &npool.CreateUserSpecialReductionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateUserSpecialReduction(ctx context.Context, in *npool.UpdateUserSpecialReductionRequest) (*npool.UpdateUserSpecialReductionResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update coupon pool error: %w", err)
		return &npool.UpdateUserSpecialReductionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserSpecialReduction(ctx context.Context, in *npool.GetUserSpecialReductionRequest) (*npool.GetUserSpecialReductionResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon pool error: %w", err)
		return &npool.GetUserSpecialReductionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserSpecialReductionsByApp(ctx context.Context, in *npool.GetUserSpecialReductionsByAppRequest) (*npool.GetUserSpecialReductionsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon pool by app error: %w", err)
		return &npool.GetUserSpecialReductionsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserSpecialReductionsByAppReleaser(ctx context.Context, in *npool.GetUserSpecialReductionsByAppReleaserRequest) (*npool.GetUserSpecialReductionsByAppReleaserResponse, error) {
	resp, err := crud.GetByAppReleaser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon pool by app releaser error: %w", err)
		return &npool.GetUserSpecialReductionsByAppReleaserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserSpecialReductionsByAppUser(ctx context.Context, in *npool.GetUserSpecialReductionsByAppUserRequest) (*npool.GetUserSpecialReductionsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon pool by app user error: %w", err)
		return &npool.GetUserSpecialReductionsByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
