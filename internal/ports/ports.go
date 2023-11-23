package ports

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
)

type FranchiseRepository interface {
	Save(ctx context.Context, franchise franchise.Franchise) error
}
