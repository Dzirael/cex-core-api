package user

import (
	"fmt"

	"cex-core-api/app/internal/models"
	user_v1 "cex-core-api/gen/user/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func modelToUser(data *models.User) (*user_v1.User, error) {
	out := &user_v1.User{
		Id:      data.UserID.String(),
		Email:   data.Email,
		Name:    data.Name,
		Surname: data.Surname,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: data.CreatedAt.UTC().Unix(),
		},
		UpdatedAt: &timestamppb.Timestamp{
			Seconds: data.UpdatedAt.UTC().Unix(),
		},
	}

	if err := out.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	return out, nil
}
