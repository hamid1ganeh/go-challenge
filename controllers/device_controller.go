package controller

import (
	model "go-challenge/models"
	service "go-challenge/services"
	"go-challenge/utility/response"
	"net/http"
	"github.com/labstack/echo/v4"
)

func Store(c echo.Context) error {

	deviceStoreRequest := new(model.DeviceStoreRequest)

	if err := c.Bind(deviceStoreRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(nil, "Bad request"))
	}

	if err := c.Validate(deviceStoreRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Fail(nil, err.Error()))
	}

	device := model.Device{
		Id:          deviceStoreRequest.Id,
		DeviceModel: deviceStoreRequest.DeviceModel,
		Name:        deviceStoreRequest.Name,
		Note:        deviceStoreRequest.Note,
		Serial:      deviceStoreRequest.Serial,
	}

	show, _ := service.GetById(deviceStoreRequest.Id)

	if show.Id == "" {
		_,err := service.Create(device)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.Fail(nil,err.Error()))
	
		}else
		{
			return c.String(http.StatusCreated, "The new device is created successfully")
		}
	}else{

		return c.JSON(http.StatusConflict, response.Conflict(nil, ""))
	}

}

func Show(c echo.Context) error {
	deviceId := "/devices/"+c.Param("id")
	show, err := service.GetById(deviceId)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(nil,err.Error()))

	}else if show.Id == "" {

		return c.JSON(http.StatusNotFound, response.NotFound(nil,""))

	}else{

		return c.JSON(http.StatusOK, response.Success(show, ""))

	}
}
