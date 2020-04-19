package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jz222/loggy/internal/models"
	"github.com/jz222/loggy/internal/services"
	"github.com/jz222/loggy/internal/store"
	"github.com/jz222/loggy/internal/utils"
)

type LoggingControllers struct {
	LoggingService services.InterfaceLogging
}

func (l *LoggingControllers) RegisterError(c *gin.Context) {
	errorEvent := models.Error{
		Badges:    map[string]string{},
		ClientIP:  c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Count:     1,
		Timestamp: time.Now().Unix(),
	}

	err := json.NewDecoder(c.Request.Body).Decode(&errorEvent)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !errorEvent.IsValid() {
		utils.RespondWithError(c, http.StatusBadRequest, "the provided data is too large")
		return
	}

	go l.LoggingService.SaveError(errorEvent)

	utils.RespondWithSuccess(c)
}

func GetLoggingController(store store.InterfaceStore) LoggingControllers {
	loggingService := services.GetLoggingService(store)

	return LoggingControllers{
		LoggingService: &loggingService,
	}
}