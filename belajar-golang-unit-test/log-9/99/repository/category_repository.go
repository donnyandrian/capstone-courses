package repository

import "belajar-golang-unit-test/log-9/99/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
