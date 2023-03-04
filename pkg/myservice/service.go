package myservice

// RepositoryInternal provides access to MySQL repository
type RepositoryMySQL interface {
	// CreateAccount(context.Context, adding.Order) (int, error)
}

// RepositoryMemory provides access to in-memory repository
type RepositoryRedis interface {
}

// Service provides equity listing operations
type Service interface {
}

type service struct {
	rmy RepositoryMySQL
	rre RepositoryRedis
}

// NewService creates an listing service with the necessary dependencies
func NewService(rmy RepositoryMySQL, rre RepositoryRedis) Service {
	return &service{rmy, rre}
}
