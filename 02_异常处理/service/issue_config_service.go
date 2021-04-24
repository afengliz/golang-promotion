package service

import (
	"context"
	"fmt"
	"github.com/liyanfeng123/golang-promotion/02_异常处理/dal"
	"github.com/liyanfeng123/golang-promotion/02_异常处理/errgroup"
	"github.com/liyanfeng123/golang-promotion/02_异常处理/model/dto"
	"time"
)

type IssueConfigService struct {
}

// GetTestData 并发获取基础数据
func (s *IssueConfigService) GetTestData(ctx context.Context) (res dto.TestData, err error) {
	defer func() {
		if err != nil{
			fmt.Printf("堆栈信息:\n%+v\n", err)
		}
	}()
	// 设置3秒超时
	timeCtx,cancel := context.WithTimeout(ctx,time.Second*10)
	defer cancel()
	g:= errgroup.WithContext(timeCtx)
	result := dto.TestData{}
	g.Go(func(ctx context.Context) error {
		configDal := dal.IssueConfigDal{}
		aData, err := configDal.GetA(ctx)
		result.User = aData
		return err
	})
	g.Go(func(ctx context.Context) error {
		configDal := dal.IssueConfigDal{}
		bData, err := configDal.GetB()
		result.B = bData
		return err
	})
	err = g.Wait()
	return result, err
}



