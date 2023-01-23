package repository

import (
	"fmt"
	"packform-test/src/models"

	"gorm.io/gorm"
)

func GetDatatable(relation *gorm.DB, search string, order string) *gorm.DB {
	if search != "" {
		relation = relation.Where("(order_name || ' ' || customer_name || ' ' || customer_company) ILIKE ?", "%"+search+"%")
	}
	if order == "" {
		order = "desc"
	}
	return relation.Debug().Model(&models.Report{}).Order(fmt.Sprintf("order_date %s", order))
}
