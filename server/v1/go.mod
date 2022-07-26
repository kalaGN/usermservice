module server

go 1.18

require google.golang.org/grpc v1.48.0

require (
	github.com/go-ini/ini v1.66.6 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/smartystreets/assertions v1.13.0 // indirect
	github.com/tietang/go-utils v0.1.3 // indirect
	github.com/tietang/props v2.2.0+incompatible // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gorm.io/driver/mysql v1.3.5 // indirect
	gorm.io/gorm v1.23.8 // indirect
)

require login v0.0.0

replace login => ../../pb/users

require pkg/config v0.0.0

replace pkg/config => ../pkg/config

require pkg/database v0.0.0

replace pkg/database => ../pkg/database
