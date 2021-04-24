package model

type User struct {
	Id    int64  `json:"id"`
	Age   int32  `json:"age"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "user"
}