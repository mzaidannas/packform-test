package repository

import (
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

func CreatePaginator(
	key *string,
	cursor paginator.Cursor,
	order *paginator.Order,
	limit *int,
) *paginator.Paginator {
	p := paginator.New()
	if key != nil {
		p.SetKeys(*key)
	}
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

func GetDatatable[GeneralModel any](db *gorm.DB, search *string, orderCol *string, orderDir *string, limit *int) ([]GeneralModel, paginator.Cursor, error) {

	var objects []GeneralModel

	if *limit == 0 {
		*limit = 100
	}
	if *orderCol == "" {
		*orderCol = "ID"
	}
	if *orderDir == "" {
		*orderDir = "DESC"
	}

	order_p := paginator.Order(*orderDir)

	// create paginator for User model
	p := CreatePaginator(orderCol, paginator.Cursor{}, &order_p, limit)

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
