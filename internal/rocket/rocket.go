//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket  github.com/sdivyansh59/grpc-microservice-go/internal/rocket Store

package rocket

import "context"

// Rocket - should contains the defination of our
// rocket
type Rocket struct {
	ID string
	Name string
	Type string
	Flights int
}

// Store - defines the interface we expect
// our database implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

//Service - our rocket service responsible for 
// updating the rocket inventory
type Service struct {
	Store Store
}

// New - return a new instance of our rocket service 
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID - retrieve a rocket based on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{},err
	}

	return rkt, nil
}

// InsertRocket - insert a new  a rocket into store
func (s Service) InsertRocket(ctx context.Context,rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{},err
	}

	return rkt, nil
}

// DeleteRocket - delete a rocket from inventory
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}

	return nil
}