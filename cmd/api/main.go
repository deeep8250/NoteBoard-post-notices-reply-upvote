package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	authhandler "github.com/threadpulse/internal/auth/handlers"
	authrepo "github.com/threadpulse/internal/auth/repository"
	authservice "github.com/threadpulse/internal/auth/services"
	"github.com/threadpulse/internal/db/redis"
	"github.com/threadpulse/internal/middleware"
	"github.com/threadpulse/internal/routes"

	threadhandler "github.com/threadpulse/internal/threads/handlers"
	threadrepo "github.com/threadpulse/internal/threads/repository"
	threadservice "github.com/threadpulse/internal/threads/services"

	replieshandler "github.com/threadpulse/internal/replies/handlers"
	repliesrepo "github.com/threadpulse/internal/replies/repository"
	repliesservice "github.com/threadpulse/internal/replies/services"

	upvotehandler "github.com/threadpulse/internal/upvotes/handlers"
	upvoterepo "github.com/threadpulse/internal/upvotes/repositories"
	upvoteservice "github.com/threadpulse/internal/upvotes/services"

	_ "github.com/threadpulse/internal/routes"
)

func main() {

	dsn := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	var db *sqlx.DB
	var err error

	for range 5 {
		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("error while connecting to db : ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("database connection failed")

	}

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal("Migration files not found ", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migtaion failed ", err.Error())
	}

	log.Println("Migrations ran successfully")

	fmt.Println("database connect successfully")

	// pinging redis
	redisClient := redis.NewRedisClient()
	ctx := context.Background()
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("redis connection failed: ", err.Error())
	}

	log.Println("Redis connected successfully")

	//

	// dependency injections

	AuthRepo := authrepo.NewAuthRepo(db)
	AuthService := authservice.NewAuthService(AuthRepo)
	AuthHandler := authhandler.NewAuthHandler(AuthService)

	//threads
	ThreadRepo := threadrepo.NewThreadRepo(db)

	ThreadService := threadservice.NewThreadsService(ThreadRepo, redisClient)
	ThreadHandler := threadhandler.NewThreadHandler(ThreadService)

	// replies
	RepliesRepo := repliesrepo.NewRepliesRepo(db)
	repliesService := repliesservice.NewRepliesService(RepliesRepo)
	repliesHandler := replieshandler.NewRepliesHandler(repliesService)

	// upvotes
	UpvoteRepo := upvoterepo.NewUpvotesRepository(db)
	UpvoteWorker := upvoterepo.NewUpvoteWorker(UpvoteRepo)
	UpvoteWorker.Start()
	UpvoteService := upvoteservice.NewUpvoteService(UpvoteRepo, UpvoteWorker)
	UpvoteHandler := upvotehandler.NewUpvoteHandler(UpvoteService)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	routes.Routes(r, AuthHandler, ThreadHandler, repliesHandler, UpvoteHandler)

	r.Run(":8080")

}
