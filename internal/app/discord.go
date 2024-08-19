package app

import (
	"context"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
)

type DiscordRepository interface {
}

type DiscordService interface {
}

type DiscordUseCase struct {
	trace.Tracable

	pwd    string
	signer *domain.JWTSigner
	//repo   DiscordRepository
	//svc    DiscordService
}

func NewDiscordUseCase( /*repo DiscordRepository, svc DiscordService,*/ jwtSecret, pwd string) *DiscordUseCase {
	return &DiscordUseCase{
		Tracable: trace.NewTracable(),
		signer:   domain.NewJWTSigner(jwtSecret),
		//repo:     repo,
		pwd: pwd,
		//svc:      svc,
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
