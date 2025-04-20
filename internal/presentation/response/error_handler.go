package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"zoo/internal/application/errs"
)

var errorStatusMap = map[error]int{
	errs.ErrAnimalNotFound:         http.StatusNotFound,
	errs.ErrEnclosureNotFound:      http.StatusNotFound,
	errs.ErrScheduleNotFound:       http.StatusNotFound,
	errs.ErrFeedingAlreadyOccurred: http.StatusConflict,
	errs.ErrEnclosureIsFull:        http.StatusConflict,
	errs.ErrInvalidID:              http.StatusConflict,
	errs.ErrInvalidDate:            http.StatusBadRequest,
	errs.ErrInvalidTime:            http.StatusBadRequest,
	errs.ErrInvalidGender:          http.StatusBadRequest,
	errs.ErrInvalidStatus:          http.StatusBadRequest,
}

func HandleError(c *gin.Context, err error) {
	for e, code := range errorStatusMap {
		if errors.Is(err, e) {
			c.JSON(code, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
