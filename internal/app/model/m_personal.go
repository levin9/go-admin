package model

import (
	"context"

	"github.com/Levin9/go-admin/internal/app/schema"
)

// IPersonal 人员管理存储接口
type IPersonal interface {
	// 查询数据
	Query(ctx context.Context, params schema.PersonalQueryParam, opts ...schema.PersonalQueryOptions) (*schema.PersonalQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.PersonalQueryOptions) (*schema.Personal, error)
	// 创建数据
	Create(ctx context.Context, item schema.Personal) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.Personal) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
