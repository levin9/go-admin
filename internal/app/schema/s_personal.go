package schema

// Personal 人员管理
type Personal struct {
	PersonId         string   `json:"person_id" binding:"required" swaggo:"true,人员编号"`
	TenantId         string   `json:"tenant_id" binding:"required" swaggo:"true,租户编号"`
	UserId           string   `json:"user_id" binding:"required" swaggo:"true,用户编号"`
	RealName         string   `json:"real_name" swaggo:"false,姓名"`
	Mobile           string   `json:"mobile" swaggo:"false,手机号"`
	JobNo            string   `json:"job_no" swaggo:"false,工号"`
	DeptName         string   `json:"dept_name" swaggo:"false,部门"`
	JonName          string   `json:"jon_name" swaggo:"false,职务"`
	JobLevel         string   `json:"job_level" swaggo:"false,职级"`
	JoinDate         datetime `json:"join_date" swaggo:"false,入职时间"`
	LastDate         datetime `json:"last_date" swaggo:"false,离职时间"`
	RegularDate      datetime `json:"regular_date" swaggo:"false,转正时间"`
	DingTalkId       string   `json:"ding_talk_id" swaggo:"false,钉钉ID"`
	F_EnabledMark    int      `json:"f_enabled_mark" binding:"required" swaggo:"true,启用标记"`
	F_CreateUserId   string   `json:"f_create_user_id" swaggo:"false,创建用户"`
	F_CreateUserName string   `json:"f_create_user_name" swaggo:"false,创建用户"`
	F_ModifyUserId   string   `json:"f_modify_user_id" swaggo:"false,职级"`
	F_ModifyUserName string   `json:"f_modify_user_name" swaggo:"false,职级"`
	F_ModifyTime     string   `json:"f_modify_time" swaggo:"false,职级"`
	F_CreateTime     datetime `json:"f_create_time" swaggo:"false,创建时间"`
	F_Description    string   `json:"f_description" swaggo:"false,备注"`
	F_SortCode       int      `json:"f_sort_code" swaggo:"false,序号"`
	F_DeleteMark     int      `json:"f_delete_mark" binding:"required" swaggo:"true,状态(1:启用 2:停用)"`
}

// PersonalQueryParam 查询条件
type PersonalQueryParam struct {
}

// PersonalQueryOptions 查询可选参数项
type PersonalQueryOptions struct {
	PageParam *PaginationParam // 分页参数
}

// PersonalQueryResult 查询结果
type PersonalQueryResult struct {
	Data       Personals
	PageResult *PaginationResult
}

// Personals 人员管理列表
type Personals []*Personal
