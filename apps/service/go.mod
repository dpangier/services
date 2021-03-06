module github.com/micro/services/apps/service

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro/v2 v2.5.1-0.20200429144330-1d29f126f944
	github.com/micro/services/payments/provider v0.0.0-20200318105532-9c3078c484d5
)

replace github.com/micro/services/payments/provider => ../../payments/provider
