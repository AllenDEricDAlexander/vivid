package persistence

import (
	"context"
	"fmt"
	"time"
	"user/model"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDAOStrategy interface {
	Create(user model.User) error
	Init() error
}

type MySQLDAO struct {
	DB *gorm.DB
}

func (p *MySQLDAO) Create(user model.User) error {
	p.DB.Create(user)
	return nil
}

func (p *MySQLDAO) Init() error {
	dsn := "root:MQa12345@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	p.DB = db
	return nil
}

type RedisDAO struct {
	DB *redis.Client
}

func (r *RedisDAO) Create(user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	err := r.DB.Do(ctx, "set", user.ID, user.Nickname, "EX", 3600).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisDAO) Init() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	r.DB = rdb
	return nil
}

type PostgreSQLDAO struct {
	DB *gorm.DB
}

func (w *PostgreSQLDAO) Create(user model.User) error {
	return nil
}

func (w *PostgreSQLDAO) Init() error {
	dsn := "root:MQa12345@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	w.DB = db
	return nil
}

type UserDAOContext struct {
	strategy UserDAOStrategy
}

func (pc *UserDAOContext) SetDB(strategy UserDAOStrategy) {
	pc.strategy = strategy
	pc.strategy.Init()
}

func (pc *UserDAOContext) DoCreate(user model.User) error {
	err := pc.strategy.Create(user)
	if err != nil {
		fmt.Println("Error creating user persistence context")
		return err
	}
	return nil
}
