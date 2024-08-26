package app

import (
	"context"
	"github.com/NoBypass/fds/internal/common/errs"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
	"time"
)

type DiscordRepository interface {
	GetMember(ctx context.Context, id string) (*domain.DiscordMember, error)
	UpdateMember(ctx context.Context, member *domain.DiscordMember) error
	CreateMember(ctx context.Context, member *domain.DiscordMember) error
	GetLeaderboard(ctx context.Context, page int) (*domain.Leaderboard, error)
	RemoveMember(ctx context.Context, id string) error
}
type HypixelService interface {
	PlayerByName(ctx context.Context, name string) (*domain.HypixelPlayerResp, error)
}

type DiscordUseCase struct {
	trace.Tracable

	pwd     string
	signer  *domain.JWTSigner
	repo    DiscordRepository
	hypixel HypixelService
}

func NewDiscordUseCase(repo DiscordRepository, svc HypixelService, jwtSecret, pwd string) *DiscordUseCase {
	return &DiscordUseCase{
		Tracable: trace.NewTracable(),
		signer:   domain.NewJWTSigner(jwtSecret),
		hypixel:  svc,
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
		return nil, errs.AlreadyClaimed
	}

	member.ClaimDaily()
	return member, nil
}

func (uc *DiscordUseCase) GetLeaderboard(ctx context.Context, page int) (*domain.Leaderboard, error) {
	sp, ctx := uc.StartSpan(ctx, uc.GetLeaderboard)
	defer sp.Finish()

	return uc.repo.GetLeaderboard(ctx, page)
}

func (uc *DiscordUseCase) Verify(ctx context.Context, discordID, discordName, ign string) (string, error) {
	sp, ctx := uc.StartSpan(ctx, uc.Verify)
	defer sp.Finish()

	playerResp, err := uc.hypixel.PlayerByName(ctx, ign)
	if err != nil {
		return "", err
	}

	if playerResp.Player.SocialMedia.Links["DISCORD"] != discordName {
		return "", nil
	}

	err = uc.repo.CreateMember(ctx, &domain.DiscordMember{
		DiscordID:   discordID,
		LastDailyAt: time.Now().Add(time.Hour * -24),
	})
	if err != nil {
		return "", err
	}

	return playerResp.Player.DisplayName, nil
}

func (uc *DiscordUseCase) Revoke(ctx context.Context, discordID string) error {
	sp, ctx := uc.StartSpan(ctx, uc.Revoke)
	defer sp.Finish()

	return uc.repo.RemoveMember(ctx, discordID)
}

func (uc *DiscordUseCase) GiveXP(ctx context.Context, discordID string, amount int) error {
	sp, ctx := uc.StartSpan(ctx, uc.GiveXP)
	defer sp.Finish()

	member, err := uc.repo.GetMember(ctx, discordID)
	if err != nil {
		return err
	}

	member.AddXP(amount)
	return uc.repo.UpdateMember(ctx, member)
}
