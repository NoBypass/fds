package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/surgo"
	"golang.org/x/net/context"
	"os"
)

func init() {
	db, err := surgo.New(
		context.TODO(),
		os.Getenv("db_host"),
		surgo.Pass(os.Getenv("db_pwd")),
		surgo.User(os.Getenv("db_user")),
		surgo.Database(os.Getenv("db_name")),
		surgo.Namespace(os.Getenv("db_namespace")),
	)
	if err != nil {
		panic(err)
	}

	Discord = surgo.Model[model.DiscordMember](db)
	Hypixel = surgo.Model[model.HypixelPlayer](db)
	Mojang = surgo.Model[model.MojangProfile](db)
}

var (
	Discord      surgo.DBModel[model.DiscordMember]
	Hypixel      surgo.DBModel[model.HypixelPlayer]
	Mojang       surgo.DBModel[model.MojangProfile]
	PlayedWith   surgo.DBRelation[model.DiscordMember, model.HypixelPlayer, model.PlayedWith]
	VerifiedWith surgo.DBRelation[model.DiscordMember, model.MojangProfile, model.VerifiedWith]
)
