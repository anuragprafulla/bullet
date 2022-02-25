package users

import (
	"context"
	"log"
	"os"

	"github.com/anuragprafulla/bullet/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pg struct {
	db *gorm.DB
}

func NewPostgresUserStore(conn string) IUserStore {
	db, err := gorm.Open(postgres.Open(conn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Info,
					Colorful: true,
				},
			),
		},
	)

	if err != nil {
		panic("Unable to connect to database with " + conn + err.Error())
	}

	if err := db.AutoMigrate(User{}); err != nil {
		panic("Unable to migrate database: " + err.Error())
	}

	return &pg{db: db}
}

func (p *pg) Get(ctx context.Context, in *GetUserRequest) (*User, error) {
	user := &User{}
	err := p.db.WithContext(ctx).Take(user, "id = ?", in.ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrorBadRequest
	}

	return user, err
}

func (p *pg) List(ctx context.Context, in *ListUserRequest) ([]*User, error) {
	if in.Limit == 0 || in.Limit > MaxListLimit {
		in.Limit = MaxListLimit
	}

	query := p.db.WithContext(ctx).Limit(in.Limit)

	if in.After != "" {
		query = query.Where("id > ?", in.After)
	}

	if in.Name != "" {
		query = query.Where("name ilike ?", "%"+in.Name+"%")
	}

	list := make([]*User, 0, in.Limit)
	err := query.Order("id").Find(&list).Error

	return list, err
}

func (p *pg) Create(ctx context.Context, in *CreateUserRequest) error {
	if in.User == nil {
		return errors.ErrorBadRequest
	}

	in.User.ID = GenerateUniqueID()

	return p.db.WithContext(ctx).
		Create(in.User).Error
}

func (p *pg) Update(ctx context.Context, in *UpdateUserRequest) error {
	user := &User{
		ID:          in.ID,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		EmailId:     in.EmailId,
		PhoneNumber: in.PhoneNumber,
		UpdatedAt:   p.db.NowFunc(),
	}

	return p.db.WithContext(ctx).Model(user).
		Select("first_name", "last_name", "email_id", "phone_number", "updated_at").
		Updates(user).Error
}

func (p *pg) Delete(ctx context.Context, in *DeleteUserRequest) error {
	user := User{ID: in.ID}
	return p.db.WithContext(ctx).Model(user).Delete(user).Error
}
