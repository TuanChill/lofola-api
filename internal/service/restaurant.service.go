package service

import "github.com/tuanchill/lofola-api/internal/models"

type IRestaurantService interface {
	CreateRestaurant(restaurant *models.Restaurant) error
	GetRestaurantByID(id uint) (*models.Restaurant, error)
	GetRestaurants() ([]*models.Restaurant, error)
	UpdateRestaurant(restaurant *models.Restaurant) error
	DeleteRestaurant(id uint) error
}

type restaurantService struct {
}

func NewRestaurantService() IRestaurantService {
	return &restaurantService{}
}

// CreateRestaurant implements IRestaurantService.
func (r *restaurantService) CreateRestaurant(restaurant *models.Restaurant) error {
	panic("unimplemented")
}

// DeleteRestaurant implements IRestaurantService.
func (r *restaurantService) DeleteRestaurant(id uint) error {
	panic("unimplemented")
}

// GetRestaurantByID implements IRestaurantService.
func (r *restaurantService) GetRestaurantByID(id uint) (*models.Restaurant, error) {
	panic("unimplemented")
}

// GetRestaurants implements IRestaurantService.
func (r *restaurantService) GetRestaurants() ([]*models.Restaurant, error) {
	panic("unimplemented")
}

// UpdateRestaurant implements IRestaurantService.
func (r *restaurantService) UpdateRestaurant(restaurant *models.Restaurant) error {
	panic("unimplemented")
}
