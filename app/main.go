package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"todo/lib/database_transaction"
	"todo/lib/redis"
	"todo/utils/validators"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/subosito/gotenv"
	"go.uber.org/fx"

	"todo/app/middleware"
	"todo/app/server"

	activityHTTP "todo/service/activity/delivery/http"
	activityModule "todo/service/activity/module"
)

type libs struct {
	fx.Out

	Redis              redis.Client
	TransactionManager database_transaction.Client
}

type handlers struct {
	fx.In

	// OhlcHandler *todoHTTP.Handler
	ActivityHandler *activityHTTP.Handler
}

func main() {
	log.Println("server is starting")

	loadEnv()

	app := fx.New(
		fx.Provide(
			setupDatabase,
			initLibs,
		),
		activityModule.Module,
		fx.Invoke(
			validators.NewValidator,
			startServer,
		),
	)

	app.Run()
}

func startServer(lc fx.Lifecycle, db *gorm.DB, handlers handlers) {
	m := middleware.New(middleware.Config{
		Db: db,
	})

	h := server.BuildHandler(m,
		handlers.ActivityHandler,
	)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      h,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func(s *http.Server) {
				log.Printf("api is available at %s\n", s.Addr)
				if err := s.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatal(err)
				}
			}(s)
			return nil
		},
		OnStop: func(c context.Context) error {
			_ = s.Shutdown(c)
			log.Println("api gracefully stopped")
			return nil
		},
	})
}

func loadEnv() {
	err := gotenv.Load()

	if err != nil {
		log.Println("failed to load from .env")
	}
}

func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Println(err)
		panic(err)
	}

	sqlDB, _ := db.DB()

	// Set the maximum number of concurrently idle connections. Setting this
	// to less than or equal to 0 will mean that no idle connections are retained.
	sqlDB.SetMaxIdleConns(50)

	// Set the number of open connections (in-use + idle).
	sqlDB.SetMaxOpenConns(50)

	// Set the maximum lifetime of a connection to 1 hour. Setting it to 0
	// means that there is no maximum lifetime and the connection is reused
	// forever (which is the default behavior).
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func initLibs(lc fx.Lifecycle, db *gorm.DB) libs {
	l := libs{
		Redis: redis.NewClient(redis.Credentials{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}, os.Getenv("APP_ENV")),
		TransactionManager: database_transaction.New(db),
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			_ = l.Redis.Close()

			return nil
		},
	})

	return l
}
