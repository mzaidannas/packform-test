package repository

import (
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

func CreatePaginator(
	cursor paginator.Cursor,
	order *paginator.Order,
	limit *int,
) *paginator.Paginator {
	p := paginator.New()
	if order != nil {
		p.SetOrder(*order)
	}
	if limit != nil {
		p.SetLimit(*limit)
	}
	if cursor.After != nil {
		p.SetAfterCursor(*cursor.After)
	}
	if cursor.Before != nil {
		p.SetBeforeCursor(*cursor.Before)
	}
	return p
}

func GetDatatable[GeneralModel any](db *gorm.DB, search *string, order *string, limit *int) ([]GeneralModel, paginator.Cursor, error) {

	var objects []GeneralModel

	if *limit == 0 {
		*limit = 10
	}
	if *order == "" {
		*order = "DESC"
	}

	order_p := paginator.Order(*order)

	// create paginator for User model
	p := CreatePaginator(paginator.Cursor{}, &order_p, limit)

	// find objects with pagination
	result, cursor, err := p.Paginate(db, &objects)

	// this is paginator error, e.g., invalid cursor
	if err != nil {
		return nil, paginator.Cursor{}, err
	}

	// this is gorm error
	if result.Error != nil {
		return nil, paginator.Cursor{}, result.Error
	}

	return objects, cursor, nil

}
