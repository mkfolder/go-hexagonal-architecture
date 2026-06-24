package driven

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"mkfolder.dev/wire-playground/internal/database"
	"mkfolder.dev/wire-playground/internal/user/core"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"not null"`
	Username  string    `gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u UserModel) toDomain() *core.User {
	return &core.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(pg *database.Postgres) *PostgresRepository {
	return &PostgresRepository{db: pg.GetDB()}
}

func (r *PostgresRepository) FindByID(id uuid.UUID) (*core.User, error) {
	var user UserModel
	err := r.db.First(&user, id).Error
	return user.toDomain(), err
}

func (r *PostgresRepository) FindByUsername(username string) (*core.User, error) {
	var user UserModel
	err := r.db.First(&user, "username = ?", username).Error
	return user.toDomain(), err
}

func (r *PostgresRepository) Create(user *core.User) error {
	model := UserModel{
		Email:    user.Email,
		Username: user.Username,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return err
	}

	*user = *model.toDomain()
	return nil
}
