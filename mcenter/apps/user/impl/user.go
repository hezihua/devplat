package impl

import (
	"context"
	"hzh/devcloud/mcenter/apps/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 查询用户列表
func (i *impl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (
	*user.User, error) {

	ins, err := user.New(in)
	if err != nil {
		return nil, err
	}
	_, err = i.col.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 更新用户列表
func (i *impl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (
	*user.User, error) {

	return nil, nil
}

// 删除用户列表
func (i *impl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (
	*user.User, error) {
	return nil, nil
}

// 查询用户列表
func (i *impl) QueryUser(ctx context.Context, in *user.QueryUserRequest) (
	*user.UserSet, error) {
	set := user.NewUserSet()

	// 构造查询条件
	filter := bson.M{}
	if in.Keywords != "" {
		// >   $gt
		filter["username"] = bson.M{"$regex": in.Keywords, "$options": "im"}
		// bson.M{"$regex": r.Path, "$options": "im"}
	}

	opts := &options.FindOptions{}
	opts.SetLimit(int64(in.Page.PageSize))
	opts.SetSkip(in.Page.ComputeOffset())

	cursor, err := i.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	// 执行查询SQL
	for cursor.Next(ctx) {
		ins := user.NewDefaultPageRequest()
		if err := cursor.Decode(ins); err != nil {
			return nil, err
		}
		ins.Desense()
		set.Add(ins)
	}

	// 统计总数
	set.Total, err = i.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return set, nil
}

// 查询用户详情
func (i *impl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (
	*user.User, error) {
	//校验空
	// TODO

	filter := bson.M{}

	switch in.DescribeBy {
	case user.DESCRIBE_BY_USER_ID:
		filter["_id"] = in.DescribeValue
	case user.DESCRIBE_BY_USERNAME:
		filter["username"] = in.DescribeValue
	}
	ins := user.NewDefaultPageRequest()

	err := i.col.FindOne(ctx, filter).Decode(ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
