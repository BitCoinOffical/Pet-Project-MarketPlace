package sqlbuilder

import (
	"fmt"
	"strings"

	"main.go/interfaces/http/dto"
)

func BuildGetAllQuery(filter *dto.ProductFilterDTO) (string, string, []interface{}, error) {
	query := []string{}
	args := []interface{}{}
	num := 1
	page := filter.Page
	if filter.Page <= 0 {
		page = 1
	}

	if filter.Category != nil {
		query = append(query, fmt.Sprintf("category = $%d", num))
		args = append(args, filter.Category)
		num++
	}

	if filter.Search != nil {
		query = append(query, fmt.Sprintf("name ILIKE $%d", num))
		args = append(args, "%"+*filter.Search+"%")
		num++
	}

	if filter.MinPrice != nil {
		query = append(query, fmt.Sprintf("price >= $%d", num))
		args = append(args, filter.MinPrice)
		num++
	}

	if filter.MaxPrice != nil {
		query = append(query, fmt.Sprintf("price <= $%d", num))
		args = append(args, filter.MaxPrice)
		num++
	}

	where := ""
	if len(query) > 0 {
		where = "WHERE " + strings.Join(query, " AND ")
	}

	offset := (page - 1) * filter.Limit
	limitOffset := fmt.Sprintf("LIMIT $%d OFFSET $%d", num, num+1)
	args = append(args, filter.Limit, offset)
	return where, limitOffset, args, nil
}
