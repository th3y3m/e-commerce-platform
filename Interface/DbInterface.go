package Interface

import "gorm.io/gorm"

type IDatabase interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Count(count *int64) *gorm.DB
	Model(value interface{}) *gorm.DB
}
