package device

import "github.com/jackc/pgx/v5/pgxpool"

type Controller struct{
	Service *Service
}

func NewController(service *Service) *Controller{
	return &Controller{
		Service: service,
	}
}

type Module struct{
	Controller *Controller
	Service *Service
	Repository *Repository
}

func NewModule(db *pgxpool.Pool) *Module{
	Repo := NewRepository(db)
	Service := NewService(Repo)
	Controller := NewController(Service)
	return &Module{
		Controller: Controller,
		Service: Service,
		Repository: Repo,
	}
}