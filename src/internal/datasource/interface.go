package datasource

import "github.com/google/uuid"

type Repository[T any] interface { // repository for generic data, *domain.GameState in this case
    Save(data T) error
    Load(id uuid.UUID) T
}
