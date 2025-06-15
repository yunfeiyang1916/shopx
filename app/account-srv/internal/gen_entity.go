package internal

// 生成entity(与表结构一一对应)及gorm query
//go:generate gentool -dsn "root:123456@tcp(localhost:3306)/shopx_account" -modelPkgName entity -outPath ./biz/query
