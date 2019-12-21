package model

import (
	"context"

	"github.com/levin9/go-admin/internal/app/errors"
	"github.com/levin9/go-admin/internal/app/model/impl/gorm/internal/entity"
	"github.com/levin9/go-admin/internal/app/schema"
	"github.com/levin9/go-admin/pkg/gormplus"
)

// NewPersonal 创建人员管理存储实例
func NewPersonal(db *gormplus.DB) *Personal {
	return &Personal{db}
}

// Personal 人员管理存储
type Personal struct {
	db *gormplus.DB
}

func (a *Personal) getQueryOption(opts ...schema.PersonalQueryOptions) schema.PersonalQueryOptions {
	var opt schema.PersonalQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Personal) Query(ctx context.Context, params schema.PersonalQueryParam, opts ...schema.PersonalQueryOptions) (*schema.PersonalQueryResult, error) {
	db := entity.GetPersonalDB(ctx, a.db).DB

	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.Personals
	pr, err := WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.PersonalQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaPersonals(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Personal) Get(ctx context.Context, recordID string, opts ...schema.PersonalQueryOptions) (*schema.Personal, error) {
	db := entity.GetPersonalDB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.Personal
	ok, err := a.db.FindOne(db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaPersonal(), nil
}

// Create 创建数据
func (a *Personal) Create(ctx context.Context, item schema.Personal) error {
	Personal := entity.SchemaPersonal(item).ToPersonal()
	result := entity.GetPersonalDB(ctx, a.db).Create(Personal)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *Personal) Update(ctx context.Context, recordID string, item schema.Personal) error {
	Personal := entity.SchemaPersonal(item).ToPersonal()
	result := entity.GetPersonalDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates(Personal)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *Personal) Delete(ctx context.Context, recordID string) error {
	result := entity.GetPersonalDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.Personal{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
