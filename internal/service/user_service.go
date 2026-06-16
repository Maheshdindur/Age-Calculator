package service

import (
	"context"
	"backend-project/db/sqlc"
	"backend-project/internal/models"
	"backend-project/internal/repository"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.UserRequest) (models.UserResponse, error)
	GetUserByID(ctx context.Context, id int32) (models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req models.UserRequest) (models.UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context) ([]models.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req models.UserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  pgtype.Date{Time: dob, Valid: true},
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
	}, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob.Time),
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  pgtype.Date{Time: dob, Valid: true},
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var res []models.UserResponse
	for _, u := range users {
		res = append(res, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob.Time.Format("2006-01-02"),
			Age:  CalculateAge(u.Dob.Time),
		})
	}
	return res, nil
}

func CalculateAge(dob time.Time) string {
	now := time.Now()

	years := now.Year() - dob.Year()
	months := int(now.Month()) - int(dob.Month())
	days := now.Day() - dob.Day()

	if days < 0 {
		months--
		// Get days in previous month
		daysInPrevMonth := time.Date(now.Year(), now.Month(), 0, 0, 0, 0, 0, time.UTC).Day()
		days += daysInPrevMonth
	}

	if months < 0 {
		years--
		months += 12
	}

	return fmt.Sprintf("%d years %d months %d days", years, months, days)
}
