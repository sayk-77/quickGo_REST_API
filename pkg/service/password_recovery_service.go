package service

import (
	"context"
	"example.com/go/pkg/database"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

type PasswordRecoveryService struct {
	redis *redis.Client
	db    *gorm.DB
}

var ctx = context.Background()

func NewPasswordRecoveryService(redis *redis.Client, db *gorm.DB) *PasswordRecoveryService {
	return &PasswordRecoveryService{
		redis: redis,
		db:    db,
	}
}

func (prs *PasswordRecoveryService) CheckEmail(email string) (int, error) {
	if err := database.NewClientRepository(prs.db).ClientFindByEmail(email); err != nil {
		return 0, err
	}

	code := rand.Intn(900000) + 100000

	if err := prs.redis.Set(ctx, email, code, 0).Err(); err != nil {
		return 0, err
	}

	NewEmailService("smtp.gmail.com", "587", "yawaihv2@gmail.com", "bdkp ntae lrro bswq").SendRecoveryMail(code, email)

	return 0, nil
}

func (prs *PasswordRecoveryService) CheckCode(code int, email string) (string, error) {
	codeRedis, err := prs.redis.Get(ctx, email).Result()
	if err != nil {
		return "", err
	}

	codeRedisInt, err := strconv.Atoi(codeRedis)
	if err != nil {
		return "", err
	}

	if codeRedisInt != code {
		return "Не верный код", nil
	}

	err = prs.redis.Del(ctx, email).Err()
	if err != nil {
		return "", err
	}

	return "Успех", nil
}

func (prs *PasswordRecoveryService) ChangePassword(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Ошибка хеширования пароля")
	}

	if err := database.NewClientRepository(prs.db).ClientPasswordRecovery(email, string(hashedPassword)); err != nil {
		return err
	}

	return nil
}
