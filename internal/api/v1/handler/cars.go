package handler

import (
	"effective_mobile_rest/internal/api/v1/handler/response"
	"effective_mobile_rest/types"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
)

// @Summary Create cars
// @Description Create cars based on given registration numbers
// @Tags Cars
// @Accept json
// @Produce json
// @Param regNums body []string true "Array of registration numbers"
// @Success 200 {object} map[string]interface{} "status": "ok"
// @Failure 400 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /cars [post]
func (h *Handler) createCars(c *gin.Context) {
	var req types.CreateCar
	if err := c.BindJSON(&req); err != nil {
		slog.Info("CreateCar error bindJson")
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	regNums := req.RegNum
	cars := make([]types.Cars, 0, len(regNums))
	wg := new(sync.WaitGroup)
	for _, regNum := range regNums {
		wg.Add(1)
		go func(regNum string) {
			defer wg.Done()
			url := fmt.Sprintf("http://localhost:8080/info?regNum=%s", regNum)
			resp, err := http.Get(url)
			if err != nil {
				slog.Info("Bad request: ", err.Error())
				return
			}
			defer resp.Body.Close()
			var carInfo types.Cars
			if err := json.NewDecoder(resp.Body).Decode(&carInfo); err != nil {
				slog.Info("bad answer for %s", regNum)
				return
			}
			cars = append(cars, carInfo)
		}(regNum)
	}
	wg.Wait()
	log.Println(cars)
	if err := h.service.CreateCars(cars); err != nil {
		slog.Info("create cars error: ", err)
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// @Summary Delete car by ID
// @Description Delete car by its ID
// @Tags Cars
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} map[string]interface{} "status": "ok"
// @Failure 400 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /cars/{id} [delete]
func (h *Handler) deleteCarById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Info("delete car by id error:", err.Error())
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.DeleteCar(id); err != nil {
		slog.Info("delete car by id error:", err.Error())
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// @Summary Get all cars
// @Description Get all cars with optional pagination
// @Tags Cars
// @Accept json
// @Produce json
// @Param limit query int false "Limit number of cars to retrieve"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} map[string]interface{} "Data": "array of cars"
// @Failure 400 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /cars [get]
func (h *Handler) getAllCars(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "1000")
	offsetParam := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		slog.Info("get all cars error: invalid limit param", limit)
		response.NewError(c, http.StatusBadRequest, "invalid limit param")
		return
	}
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		slog.Info("get all cars error: invalid offset param", offset)
		response.NewError(c, http.StatusBadRequest, "invalid offset param")
		return
	}

	cars, err := h.service.GetAllCars(limit, offset)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Debug("Cars Data:", cars)
	c.JSON(http.StatusOK, map[string]interface{}{
		"Data": cars,
	})
}

// @Summary Update car by ID
// @Description Update car information by its ID
// @Tags Cars
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Param newCarData body types.UpdateCar true "New car data"
// @Success 200 {object} map[string]interface{} "status": "ok"
// @Failure 400 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /cars/{id} [put]
func (h *Handler) updateCar(c *gin.Context) {
	var newCarData types.UpdateCar
	if err := c.BindJSON(&newCarData); err != nil {
		slog.Info("update car error:", err)
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Info("update car error: invalid id param")
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.UpdateCarById(id, newCarData); err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
