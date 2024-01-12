package model

type Device struct {
	Id          string
	DeviceModel string
	Name        string
	Note        string
	Serial      string
}

type DeviceStoreRequest struct {
	Id          string `validate:"required"`
	DeviceModel string `validate:"required"`
	Name        string `validate:"required"`
	Note        string `validate:"required"`
	Serial      string `validate:"required"`
}
