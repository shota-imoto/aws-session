module github.com/shota-imoto/line-app/src/server

replace github.com/shota-imoto/line-app/lib => ../../lib

go 1.17

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/line/line-bot-sdk-go/v7 v7.15.0
	github.com/shota-imoto/line-app/lib v0.0.0-00010101000000-000000000000
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/line/line-bot-sdk-go v7.8.0+incompatible // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
	gorm.io/gorm v1.23.4 // indirect
)
