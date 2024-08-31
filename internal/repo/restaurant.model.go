package repo

import "github.com/tuanchill/lofola-api/internal/models"

type IRestaurantRepo interface {
	CreateRestaurant(restaurant *models.Restaurant) error
	GetRestaurantByID(id uint) (*models.Restaurant, error)
	GetRestaurants() ([]*models.Restaurant, error)
	UpdateRestaurant(restaurant *models.Restaurant) error
	DeleteRestaurant(id uint) error
}

type restaurantRepo struct {
}

func NewRestaurantRepo() IRestaurantRepo {
	return &restaurantRepo{}
}

// CreateRestaurant implements IRestaurantRepo.
func (r *restaurantRepo) CreateRestaurant(restaurant *models.Restaurant) error {
	panic("unimplemented")
}

// DeleteRestaurant implements IRestaurantRepo.
func (r *restaurantRepo) DeleteRestaurant(id uint) error {
	panic("unimplemented")
}

// GetRestaurantByID implements IRestaurantRepo.
func (r *restaurantRepo) GetRestaurantByID(id uint) (*models.Restaurant, error) {
	panic("unimplemented")
}

// GetRestaurants implements IRestaurantRepo.
func (r *restaurantRepo) GetRestaurants() ([]*models.Restaurant, error) {
	panic("unimplemented")
}

// UpdateRestaurant implements IRestaurantRepo.
func (r *restaurantRepo) UpdateRestaurant(restaurant *models.Restaurant) error {
	panic("unimplemented")
}
