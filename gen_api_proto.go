package shopx

// 生成api的proto文件
//go:generate kratos proto add api/proto/account/v1/front_user2.proto

// 生成 client 源码
//go:generate kratos proto client api/proto/account/v1 --proto_path=./shopx  --proto_path=./third_party

// 生成 server 源码
//go:generate kratos proto server api/account/v1/ -t internal/service
