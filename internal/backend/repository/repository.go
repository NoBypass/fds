package repository

import (
	"github.com/NoBypass/surgo"
	"os"
)

var (
	DB *surgo.DB
)

func init() {
	DB = surgo.MustConnect(
		os.Getenv("db_host"),
		surgo.Password(os.Getenv("db_pwd")),
		surgo.User(os.Getenv("db_user")),
		surgo.Database(os.Getenv("db_name")),
		surgo.Namespace(os.Getenv("db_namespace")),
	)
}
