package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/exam-platform/chem-service/data"
	"github.com/exam-platform/chem-service/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var validGrades = map[int]bool{9: true, 10: true, 11: true, 12: true}
var validComplexity = map[string]bool{"easy": true, "medium": true, "hard": true}

// HealthCheck handles GET /health
func HealthCheck(c *gin.Context) {
	// log.Debug("Health check requested")
	c.JSON(http.StatusOK, models.HealthResponse{
		Service: "chem-service",
		Version: "1.0.0",
		Status:  "UP",
	})
}

// GetTopQuestionsByGrade handles GET /api/chem/questions/grade/:grade/top/:n
func GetTopQuestionsByGrade(c *gin.Context) {
	gradeStr := c.Param("grade")
	nStr := c.Param("n")

	grade, err := strconv.Atoi(gradeStr)
	if err != nil || !validGrades[grade] {
		// log.Warnf("Invalid grade requested: %s", gradeStr)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_grade",
			Code:    400,
			Message: "Grade must be 9, 10, 11, or 12",
		})
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 || n > 100 {
		// log.Warnf("Invalid n value requested: %s", nStr)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_count",
			Code:    400,
			Message: "n must be between 1 and 100",
		})
		return
	}

	log.Infof("Fetching top %d chemistry questions for grade %d", n, grade)
	questions := data.GetQuestionsByGrade(grade, n)
	// log.Debugf("Returning %d questions for grade %d", len(questions), grade)

	c.JSON(http.StatusOK, models.QuestionResponse{
		Subject:   "Chemistry",
		Requested: n,
		Returned:  len(questions),
		Questions: questions,
	})
}

// GetQuestionsByTopic handles GET /api/chem/questions/topic/:topic/count/:n
func GetQuestionsByTopic(c *gin.Context) {
	topic := c.Param("topic")
	nStr := c.Param("n")

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 || n > 100 {
		// log.Warnf("Invalid n value requested: %s", nStr)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_count",
			Code:    400,
			Message: "n must be between 1 and 100",
		})
		return
	}

	log.Infof("Fetching %d chemistry questions for topic: %s", n, topic)
	questions := data.GetQuestionsByTopic(topic, n)
	// log.Debugf("Returning %d questions for topic '%s'", len(questions), topic)

	c.JSON(http.StatusOK, models.QuestionResponse{
		Subject:   "Chemistry",
		Requested: n,
		Returned:  len(questions),
		Questions: questions,
	})
}

// GetQuestionsByComplexity handles GET /api/chem/questions/complexity/:complexity/count/:n
func GetQuestionsByComplexity(c *gin.Context) {
	complexity := strings.ToLower(c.Param("complexity"))
	nStr := c.Param("n")

	if !validComplexity[complexity] {
		// log.Warnf("Invalid complexity requested: %s", complexity)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_complexity",
			Code:    400,
			Message: "Complexity must be easy, medium, or hard",
		})
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 || n > 100 {
		// log.Warnf("Invalid n value requested: %s", nStr)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_count",
			Code:    400,
			Message: "n must be between 1 and 100",
		})
		return
	}

	log.Infof("Fetching %d chemistry questions with complexity: %s", n, complexity)
	questions := data.GetQuestionsByComplexity(complexity, n)
	// log.Debugf("Returning %d questions with complexity '%s'", len(questions), complexity)

	c.JSON(http.StatusOK, models.QuestionResponse{
		Subject:   "Chemistry",
		Requested: n,
		Returned:  len(questions),
		Questions: questions,
	})
}

// GetTopics handles GET /api/chem/topics
func GetTopics(c *gin.Context) {
	// log.Info("Fetching all available chemistry topics")
	topics := data.GetAllTopics()
	// log.Debugf("Returning %d topics", len(topics))
	c.JSON(http.StatusOK, gin.H{
		"subject": "Chemistry",
		"topics":  topics,
	})
}
