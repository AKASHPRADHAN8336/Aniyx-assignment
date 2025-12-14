package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/AKASHPRADHAN8336/aniyxProject/internal/models"
	"github.com/AKASHPRADHAN8336/aniyxProject/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error)
	GetUser(ctx context.Context, id int32) (*models.UserResponse, error)
	ListUsers(ctx context.Context) ([]*models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (*models.UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	years := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		years--
	}
	return years
}

func parseDate(dateStr string) (time.Time, error) {
	// Try parsing as ISO format first
	if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return t, nil
	}

	// Try parsing as date only
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		return t, nil
	}

	// Try parsing with T00:00:00Z suffix
	if t, err := time.Parse("2006-01-02T15:04:05Z", dateStr); err == nil {
		return t, nil
	}

	// Try parsing just the date part if it contains T
	if strings.Contains(dateStr, "T") {
		datePart := strings.Split(dateStr, "T")[0]
		if t, err := time.Parse("2006-01-02", datePart); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	fmt.Println("DEBUG: Creating user in service")

	dobTime, err := parseDate(req.Dob)
	if err != nil {
		return nil, err
	}

	userID, err := s.repo.Create(ctx, req.Name, req.Dob)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Parse the date from repository response
	dobTime, err = parseDate(user.Dob)
	if err != nil {
		return nil, err
	}

	age := calculateAge(dobTime)

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  dobTime.Format("2006-01-02"),
		Age:  age,
	}, nil
}

func (s *userService) GetUser(ctx context.Context, id int32) (*models.UserResponse, error) {
	fmt.Printf("DEBUG: Getting user %d in service\n", id)

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	dobTime, err := parseDate(user.Dob)
	if err != nil {
		fmt.Printf("DEBUG: Date parse error for %s: %v\n", user.Dob, err)
		return nil, err
	}

	age := calculateAge(dobTime)

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  dobTime.Format("2006-01-02"),
		Age:  age,
	}, nil
}

func (s *userService) ListUsers(ctx context.Context) ([]*models.UserResponse, error) {
	fmt.Println("DEBUG: Listing users in service")

	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("DEBUG: Service got %d users from repository\n", len(users))

	var responses []*models.UserResponse
	for _, user := range users {
		dobTime, err := parseDate(user.Dob)
		if err != nil {
			fmt.Printf("DEBUG: Skipping user %d due to date parse error: %v\n", user.ID, err)
			continue
		}

		age := calculateAge(dobTime)

		responses = append(responses, &models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			Dob:  dobTime.Format("2006-01-02"),
			Age:  age,
		})
	}

	fmt.Printf("DEBUG: Service returning %d responses\n", len(responses))
	return responses, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (*models.UserResponse, error) {
	fmt.Printf("DEBUG: Updating user %d\n", id)

	dobTime, err := parseDate(req.Dob)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(ctx, id, req.Name, req.Dob)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Parse the date from repository response
	dobTime, err = parseDate(user.Dob)
	if err != nil {
		return nil, err
	}

	age := calculateAge(dobTime)

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  dobTime.Format("2006-01-02"),
		Age:  age,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	fmt.Printf("DEBUG: Deleting user %d\n", id)
	return s.repo.Delete(ctx, id)
}
