package app

import (
	"context"
	"github.com/NoBypass/fds/internal/common/errs"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
)

type DiscordRepository interface {
	GetMember(ctx context.Context, id string) (*domain.DiscordMember, error)
	UpdateMember(ctx context.Context, member *domain.DiscordMember) error
	CreateMember(ctx context.Context, member *domain.DiscordMember) error
	GetLeaderboard(ctx context.Context, page int) (*domain.Leaderboard, error)
}

type DiscordUseCase struct {
	trace.Tracable

	pwd    string
	signer *domain.JWTSigner
	repo   DiscordRepository
}

func NewDiscordUseCase(repo DiscordRepository, jwtSecret, pwd string) *DiscordUseCase {
	return &DiscordUseCase{
		Tracable: trace.NewTracable(),
		signer:   domain.NewJWTSigner(jwtSecret),
		repo:     repo,
		pwd:      pwd,
	}
}

func (uc *DiscordUseCase) Auth(ctx context.Context, pwd string) (string, bool) {
	sp, ctx := uc.StartSpan(ctx, uc.Auth)
	defer sp.Finish()

	if uc.pwd != pwd {
		return "", false
	}

	return uc.signer.NewJWT("discord_bot", domain.RoleBot), true
}

func (uc *DiscordUseCase) Daily(ctx context.Context, id string) (*domain.DiscordMember, error) {
	sp, ctx := uc.StartSpan(ctx, uc.Daily)
	defer sp.Finish()

	member, err := uc.repo.GetMember(ctx, id)
	if err != nil {
		return nil, errs.NotFound
	}

	if !member.CanClaimDaily() {
		return nil, domain.ErrAlreadyClaimed
	}

	member.ClaimDaily()
	return member, nil
}

func (uc *DiscordUseCase) GetLeaderboard(ctx context.Context, page int) (*domain.Leaderboard, error) {
	sp, ctx := uc.StartSpan(ctx, uc.GetLeaderboard)
	defer sp.Finish()

	return uc.repo.GetLeaderboard(ctx, page)
}

func (uc *DiscordUseCase) Verify(ctx context.Context, discordID, ign string) (bool, error) {
	sp, ctx := uc.StartSpan(ctx, uc.Verify)
	defer sp.Finish()

	// TODO
	return false, nil
}
