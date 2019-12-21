package internal

import (
	"context"

	"github.com/levin9/go-admin/internal/app/errors"
	"github.com/levin9/go-admin/internal/app/model"
	"github.com/levin9/go-admin/internal/app/schema"
	"github.com/levin9/go-admin/pkg/util"
)

// NewPersonal 创建人员管理
func NewPersonal(mPersonal model.IPersonal) *Personal {
	return &Personal{
		PersonalModel: mPersonal,
	}
}

// Personal 人员管理业务逻辑
type Personal struct {
	PersonalModel model.IPersonal
}

// Query 查询数据
func (a *Personal) Query(ctx context.Context, params schema.PersonalQueryParam, opts ...schema.PersonalQueryOptions) (*schema.PersonalQueryResult, error) {
	return a.PersonalModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *Personal) Get(ctx context.Context, recordID string, opts ...schema.PersonalQueryOptions) (*schema.Personal, error) {
	item, err := a.PersonalModel.Get(ctx, recordID, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *Personal) getUpdate(ctx context.Context, recordID string) (*schema.Personal, error) {
	return a.Get(ctx, recordID)
}

// Create 创建数据
func (a *Personal) Create(ctx context.Context, item schema.Personal) (*schema.Personal, error) {
	item.RecordID = util.MustUUID()
	err := a.PersonalModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, item.RecordID)
}

// Update 更新数据
func (a *Personal) Update(ctx context.Context, recordID string, item schema.Personal) (*schema.Personal, error) {
	oldItem, err := a.PersonalModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	}

	err = a.PersonalModel.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, recordID)
}

// Delete 删除数据
func (a *Personal) Delete(ctx context.Context, recordID string) error {
	oldItem, err := a.PersonalModel.Get(ctx, recordID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.PersonalModel.Delete(ctx, recordID)
}
