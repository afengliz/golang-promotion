package dal

import (
	"context"
	"github.com/liyanfeng123/golang-promotion/02_异常处理/db"
	"github.com/liyanfeng123/golang-promotion/02_异常处理/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IssueConfigDal struct {
}

var NotFoundError  = errors.New("未找到记录")
func (p *IssueConfigDal) GetA(ctx context.Context) (model.User, error) {
	var result model.User
	if err :=db.GetDb().WithContext(ctx).Model(&model.User{}).
		Where("name = ?", "afengna").First(&result).Error;err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result,errors.WithStack(NotFoundError)
		}
		return result,errors.WithStack(err)
	}
	return result, nil
}
func (p *IssueConfigDal) GetB() (string, error) {
	return "B",nil
	//return "I m B", errors.WithStack(xerrors.New("逗你玩"))
}
