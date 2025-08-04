package server

import (
	leaderboard "contest/supervisor"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContestServer struct {
	leaderboard *leaderboard.Leaderboard
	router      *gin.Engine
}

func Init(db *sql.DB) (*ContestServer, error) {

	l, err := leaderboard.Init(db)
	if err != nil {
		return nil, fmt.Errorf("error initiating leaderboard: %w", err)
	}

	server := ContestServer{
		leaderboard: l,
		router:      router(l),
	}

	return &server, nil
}

func (cs *ContestServer) Run() {
	cs.router.Run(":8080")
}

func router(l *leaderboard.Leaderboard) *gin.Engine {
	router := gin.Default()

	// Default
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, "Hit me !!!")
	})

	router.GET("/users", func(ctx *gin.Context) {
		users, err := l.RetrieveUsers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	router.GET("/problems", func(ctx *gin.Context) {
		problems, err := l.RetrieveProblems()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, problems)
	})

	return router
}
