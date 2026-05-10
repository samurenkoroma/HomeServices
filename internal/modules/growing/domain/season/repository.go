package season

import (
	"context"
)

type Repository interface {
	Save(context.Context, *Season) error
	FindByID(ctx context.Context, id SeasonID) (*Season, error)
	//FindAll(context.Context, Filter) ([]*Season, error)

	//Delete(ctx context.Context, id SeasonID) error
	//
	//// Поиск
	//FindByName(ctx context.Context, name string) (*Season, error)
	//FindByStatus(ctx context.Context, status SeasonStatus) ([]*Season, error)
	//FindActive(ctx context.Context) (*Season, error)                              // текущий активный сезон
	//FindOverlapping(ctx context.Context, start, end time.Time) ([]*Season, error) // пересекающиеся
	//
	//// Проверки
	//ExistsByName(ctx context.Context, name string) (bool, error)
	//ExistsOverlapping(ctx context.Context, start, end time.Time) (bool, error)
}
