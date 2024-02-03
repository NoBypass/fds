package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/surrealdb/surrealdb.go"
	"strings"
)

type DiscordRepository interface {
	Repository
	Create(member *model.DiscordMember) error
	Get(id string) (*model.DiscordMember, error)
	Update(id string, member *model.DiscordMember) error

	RelatePlayedWith(in *model.MojangProfile, out *model.HypixelPlayer) error
	RelateVerifiedWith(in *model.DiscordMember, out *model.HypixelPlayer) error
}

type discordRepository struct {
	*surrealdb.DB
}

func NewDiscordRepository(db *surrealdb.DB) DiscordRepository {
	return &discordRepository{
		db,
	}
}

func (r *discordRepository) Create(member *model.DiscordMember) error {
	query := strings.Replace(`CREATE discord_member:$discordID CONTENT {
		"discordID": $discordID,
		"name": $name,
		"nick": $nick,
		"xp": $xp,
		"level": $level,
		"streak": $streak,
		"lastDailyAt": $lastDailyAt
	};`, "\n", " ", -1)
	_, err := r.DB.Query(query, map[string]interface{}{
		"discordID":   member.DiscordID,
		"name":        member.Name,
		"nick":        member.Nick,
		"xp":          member.XP,
		"level":       member.Level,
		"streak":      member.Streak,
		"lastDailyAt": member.LastDailyAt,
	})
	return err
}

func (r *discordRepository) Get(id string) (*model.DiscordMember, error) {
	member, err := surrealdb.SmartUnmarshal[model.DiscordMember](r.DB.Select("discord_member:" + id))
	return &member, err
}

func (r *discordRepository) Update(id string, member *model.DiscordMember) error {
	query := strings.Replace(`UPDATE discord_member:$discordID SET {
    	"name": $name,
    	"nick": $nick,
    	"xp": $xp,
    	"level": $level,
    	"streak": $streak,
    	"lastDailyAt": $lastDailyAt
    };`, "\n", " ", -1)
	_, err := r.DB.Query(query, map[string]interface{}{
		"discordID":   id,
		"name":        member.Name,
		"nick":        member.Nick,
		"xp":          member.XP,
		"level":       member.Level,
		"streak":      member.Streak,
		"lastDailyAt": member.LastDailyAt,
	})
	return err
}

func (r *discordRepository) RelatePlayedWith(in *model.MojangProfile, out *model.HypixelPlayer) error {
	_, err := r.DB.Query("RELATE mojang_profile:$profileID->played_with->hypixel_player:$playerID;", map[string]interface{}{
		"profileID": in.UUID,
		"playerID":  out.UUID,
	})
	return err
}

func (r *discordRepository) RelateVerifiedWith(in *model.DiscordMember, out *model.HypixelPlayer) error {
	_, err := r.DB.Query("RELATE discord_member:$memberID->verified_with->hypixel_player:$playerID;", map[string]interface{}{
		"memberID": in.DiscordID,
		"playerID": out.UUID,
	})
	return err
}
