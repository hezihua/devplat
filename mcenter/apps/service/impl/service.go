package impl

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/service"
	"go.mongodb.org/mongo-driver/bson"
)

// 查询用户列表
func (i *impl) CreateService(ctx context.Context, in *service.CreateServiceRequest) (
	*service.Service, error) {

	ins := service.New(in)
	_, err := i.col.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 查询用户详情
func (i *impl) DescribeService(ctx context.Context, in *service.DescribeServiceRequest) (
	*service.Service, error) {
	//校验空
	// TODO

	filter := bson.M{}

	switch in.DescribeBy {
	case service.DESCRIBE_BY_SERVICE_ID:
		filter["_id"] = in.DescribeValue
	case service.DESCRIBE_BY_SERVICE_CREDENTIAL_ID:
		filter["credential.client_id"] = in.DescribeValue
	}

	ins := service.NewDefaultService()

	err := i.col.FindOne(ctx, filter).Decode(ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
