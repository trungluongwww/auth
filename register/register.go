package register

import (
	"github.com/trungluongwww/auth/config"
	"github.com/trungluongwww/auth/internal/signer"
	"github.com/trungluongwww/auth/pkg/repository"
	"github.com/trungluongwww/auth/third_party/social"
	"gorm.io/gorm"
)

type Register struct {
	DB         *gorm.DB
	Repository repository.Repository
	Signer     signer.Signer
	Social     social.Social
}

func NewRegister(db *gorm.DB, cfg config.Env) *Register {
	return &Register{
		DB:         db,
		Repository: repository.NewRepository(db),
		Signer:     signer.NewSigner(cfg),
		Social:     social.NewSocial(),
	}
}
