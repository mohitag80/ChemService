package main

import (
	"os"

	"github.com/exam-platform/chem-service/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	log.Info("Starting Chemistry Exam Service v1.0.0")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Health check
	router.GET("/health", handlers.HealthCheck)

	// Chemistry question endpoints
	api := router.Group("/api/chem")
	{
		api.GET("/questions/grade/:grade/top/:n", handlers.GetTopQuestionsByGrade)
		api.GET("/questions/topic/:topic/count/:n", handlers.GetQuestionsByTopic)
		api.GET("/questions/complexity/:complexity/count/:n", handlers.GetQuestionsByComplexity)
		api.GET("/topics", handlers.GetTopics)
	}

	log.Infof("Chemistry Exam Service listening on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
