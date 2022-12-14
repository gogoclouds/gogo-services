package r

// StatusCode 响应状态
type StatusCode uint16

const (
	Ok               StatusCode = 0
	LoginErr         StatusCode = 1000
	TokenInvalid     StatusCode = 4010
	TokenMission     StatusCode = 4011
	ParameterIllegal StatusCode = 4013
	Internal         StatusCode = 5000
	// ....
)

// Status {code, msg}
var Status = map[StatusCode]string{
	0:    "操作成功",
	1000: "用户名或密码错误",
	4010: "无效的Token",
	4011: "Token缺失",
	4013: "参数不合法",
	5000: "系统内部错误",
}

const (
	CreateSuccess string = "创建成功"
	CreateFail    string = "创建失败"
	UpdateSuccess string = "更新成功"
	UpdateFail    string = "更新失败"
	DeleteSuccess string = "删除成功"
	DeleteFail    string = "删除失败"
)
