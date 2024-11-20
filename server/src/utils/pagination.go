package utils

import (
	"reflect"
	"server/src/models/dto"
)

func Paginate(data interface{}, page, perPage, total int) dto.PaginatedResponse {

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 1
	}

	offset := (page - 1) * perPage
	pagination := dto.PaginatedResponse{
		Data:    reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(data).Elem()), 0, 0).Interface(),
		Page:    page,
		PerPage: perPage,
		Total:   total,
	}

	dataValue := reflect.ValueOf(data)
	if offset < total {
		end := offset + perPage
		if end > total {
			end = total
		}
		pagination.Data = dataValue.Slice(offset, end).Interface()
	}

	return pagination
}
