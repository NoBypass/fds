package operations

import (
	"context"
	"fmt"
	"github.com/NoBypass/fds/internal/adapter"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
)

type DiscordRepository struct {
	trace.Tracable
	adapter.Database
}

func NewDiscordRepository(db adapter.Database) *DiscordRepository {
	return &DiscordRepository{
		Tracable: trace.NewTracable(),
		Database: db,
	}
}

func (r *DiscordRepository) GetMember(ctx context.Context, id string) (*domain.DiscordMember, error) {
	var member domain.DiscordMember
	err := r.DB(ctx).Scan(&member, "SELECT * FROM ONLY $discord_member", map[string]any{
		"discord_member": fmt.Sprintf("discord_member:%s", id),
	})
	return &member, err
}

func (r *DiscordRepository) UpdateMember(ctx context.Context, member *domain.DiscordMember) error {
	_, err := r.DB(ctx).Query("UPSERT ONLY $discord_member CONTENT $content", map[string]any{
		"discord_member": fmt.Sprintf("discord_member:%s", member.DiscordID),
		"content":        member,
	})
	return err
}

func (r *DiscordRepository) CreateMember(ctx context.Context, member *domain.DiscordMember) error {
	_, err := r.DB(ctx).Query("CREATE ONLY $discord_member CONTENT $content", map[string]any{
		"discord_member": fmt.Sprintf("discord_member:%s", member.DiscordID),
		"content":        member,
	})
	return err
}

func (r *DiscordRepository) GetLeaderboard(ctx context.Context, page int) (*domain.Leaderboard, error) {
	var leaderboard domain.Leaderboard
	err := r.DB(ctx).Scan(&leaderboard, "SELECT * FROM discord_member ORDER BY level DESC, xp DESC LIMIT $limit START $start", map[string]any{
		"limit": 10,
		"start": page * 10,
	})
	return &leaderboard, err
}

func (r *DiscordRepository) RemoveMember(ctx context.Context, id string) error {
	_, err := r.DB(ctx).Query("DELETE ONLY $discord_member", map[string]any{
		"discord_member": fmt.Sprintf("discord_member:%s", id),
	})
	return err
}
