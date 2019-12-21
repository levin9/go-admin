package entity

import (
	"github.com/levin9/go-admin/internal/app/schema"
	"github.com/levin9/go-admin/pkg/gormplus"
	"context"
)

// GetPersonalDB 人员管理
func GetPersonalDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return getDBWithModel(ctx, defDB, Personal{})
}

// SchemaPersonal 人员管理
type SchemaPersonal schema.Personal

// ToPersonal 转换为人员管理实体
func (a SchemaPersonal) ToPersonal() *Personal {
	item := &Personal{
		PersonId:         &a.PersonId,
		TenantId:         &a.TenantId,
		UserId:           &a.UserId,
		RealName:         &a.RealName,
		Mobile:           &a.Mobile,
		JobNo:            &a.JobNo,
		DeptName:         &a.DeptName,
		JonName:          &a.JonName,
		JobLevel:         &a.JobLevel,
		JoinDate:         &a.JoinDate,
		LastDate:         &a.LastDate,
		RegularDate:      &a.RegularDate,
		DingTalkId:       &a.DingTalkId,
		F_EnabledMark:    &a.F_EnabledMark,
		F_CreateUserId:   &a.F_CreateUserId,
		F_CreateUserName: &a.F_CreateUserName,
		F_ModifyUserId:   &a.F_ModifyUserId,
		F_ModifyUserName: &a.F_ModifyUserName,
		F_ModifyTime:     &a.F_ModifyTime,
		F_CreateTime:     &a.F_CreateTime,
		F_Description:    &a.F_Description,
		F_SortCode:       &a.F_SortCode,
		F_DeleteMark:     &a.F_DeleteMark,
	}
	return item
}

// Personal 人员管理实体
type Personal struct {
	Model
	PersonId         *string   `gorm:"column:person_id;size:50;index;"`    // 人员编号
	TenantId         *string   `gorm:"column:tenant_id;size:50;index;"`    // 租户编号
	UserId           *string   `gorm:"column:user_id;size:50;index;"`      // 用户编号
	RealName         *string   `gorm:"column:real_name;size:50;"`          // 姓名
	Mobile           *string   `gorm:"column:mobile;size:50;"`             // 手机号
	JobNo            *string   `gorm:"column:job_no;size:50;"`             // 工号
	DeptName         *string   `gorm:"column:dept_name;size:50;"`          // 部门
	JonName          *string   `gorm:"column:jon_name;size:50;"`           // 职务
	JobLevel         *string   `gorm:"column:job_level;size:50;"`          // 职级
	JoinDate         *datetime `gorm:"column:join_date;size:10;"`          // 入职时间
	LastDate         *datetime `gorm:"column:last_date;size:10;"`          // 离职时间
	RegularDate      *datetime `gorm:"column:regular_date;size:10;"`       // 转正时间
	DingTalkId       *string   `gorm:"column:ding_talk_id;size:50;"`       // 钉钉ID
	F_EnabledMark    *int      `gorm:"column:f_enabled_mark;"`             // 启用标记
	F_CreateUserId   *string   `gorm:"column:f_create_user_id;size:50;"`   // 创建用户
	F_CreateUserName *string   `gorm:"column:f_create_user_name;size:50;"` // 创建用户
	F_ModifyUserId   *string   `gorm:"column:f_modify_user_id;size:50;"`   // 职级
	F_ModifyUserName *string   `gorm:"column:f_modify_user_name;size:50;"` // 职级
	F_ModifyTime     *string   `gorm:"column:f_modify_time;size:50;"`      // 职级
	F_CreateTime     *datetime `gorm:"column:f_create_time;size:10;"`      // 创建时间
	F_Description    *string   `gorm:"column:f_description;size:200;"`     // 备注
	F_SortCode       *int      `gorm:"column:f_sort_code;size:10;"`        // 序号
	F_DeleteMark     *int      `gorm:"column:f_delete_mark;"`              // 状态(1:启用 2:停用)
}

func (a Personal) String() string {
	return toString(a)
}

// TableName 表名
func (a Personal) TableName() string {
	return a.Model.TableName("personal")
}

// ToSchemaPersonal 转换为人员管理对象
func (a Personal) ToSchemaPersonal() *schema.Personal {
	item := &schema.Personal{
		PersonId:         *a.PersonId,
		TenantId:         *a.TenantId,
		UserId:           *a.UserId,
		RealName:         *a.RealName,
		Mobile:           *a.Mobile,
		JobNo:            *a.JobNo,
		DeptName:         *a.DeptName,
		JonName:          *a.JonName,
		JobLevel:         *a.JobLevel,
		JoinDate:         *a.JoinDate,
		LastDate:         *a.LastDate,
		RegularDate:      *a.RegularDate,
		DingTalkId:       *a.DingTalkId,
		F_EnabledMark:    *a.F_EnabledMark,
		F_CreateUserId:   *a.F_CreateUserId,
		F_CreateUserName: *a.F_CreateUserName,
		F_ModifyUserId:   *a.F_ModifyUserId,
		F_ModifyUserName: *a.F_ModifyUserName,
		F_ModifyTime:     *a.F_ModifyTime,
		F_CreateTime:     *a.F_CreateTime,
		F_Description:    *a.F_Description,
		F_SortCode:       *a.F_SortCode,
		F_DeleteMark:     *a.F_DeleteMark,
	}
	return item
}

// Personals 人员管理列表
type Personals []*Personal

// ToSchemaPersonals 转换为人员管理对象列表
func (a Personals) ToSchemaPersonals() []*schema.Personal {
	list := make([]*schema.Personal, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaPersonal()
	}
	return list
}
