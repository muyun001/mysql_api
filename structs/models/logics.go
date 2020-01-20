package models

const (
	CodeRequestFormError = 400 // 访问此接口时参数格式不正确

	CodeInsertDataSucceed  = 200 // 获取结果成功
	CodeInsertDataRepeated = 201 // 重复插入数据
	CodeInsertDataFailed   = 501 // 数据插入失败

	CodeQueryDataSucceed = 200 // 查询数据成功
	CodeQueryDataFailed  = 500 // 查询数据失败

	CodeUpdateDataSucceed = 200 // 数据更新成功
	CodeUpdateDataFailed  = 500 // 数据更新失败
)

const (
	StatusUnSearched    = 0  //未查询
	StatusSearching     = 1  // 正在查询
	StatusSearchSucceed = 2  // 查询成功
	StatusSearchFailed  = -1 // 查询失败
	StatusNoNeedSearch  = -2 // 不需要再查询
)

const EachQueryProductsNum = 100 // 每次获取products的数量
