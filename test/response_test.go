package test

import (
	model "go-challenge/models"
	"go-challenge/utility/response"
	"reflect"
	"testing"
)

func TestSuccessRespone(t *testing.T) {

	if response.Success(nil, "").Message != "Success" {
		t.Errorf("Respose should return  Success message")
	}

	if response.Success(nil, "Ok").Message != "Ok" {
		t.Errorf("Respose should return Ok message")
	}

	if response.Success(nil, "").Success != true {
		t.Errorf("Success value should be true")
	}

	if response.Success(nil, "").Data != nil {
		t.Errorf("Data value should be nil")
	}

	device := model.Device{
		Id:          "/devices/id1",
		DeviceModel: "/devicemodels/id1",
		Name:        "Testing a sensor.",
		Note:        "Testing a sensor.",
		Serial:      "A020000102",
	}

	if reflect.TypeOf(response.Success(device, "").Data) != reflect.TypeOf(device) {
		t.Errorf("Data should be a type of model device")
	}
}


func TestFailRespone(t *testing.T) {

	if response.Fail(nil, "").Message != "Fail" {
		t.Errorf("Respose should return  Fail message")
	}

	if response.Fail(nil, "Sorry").Message != "Sorry" {
		t.Errorf("Rspose should return Sorry message")
	}

	if response.Fail(nil, "").Success != false {
		t.Errorf("Success value should be false")
	}

	if response.Fail(nil, "").Data != nil {
		t.Errorf("Data value should be nil")
	}
}

func TestNotFoundRespone(t *testing.T) {

	if response.NotFound(nil, "").Message != "NotFound" {
		t.Errorf("Respose should return  NotFound message")
	}

	if response.NotFound(nil, "not-found").Message != "not-found" {
		t.Errorf("Rspose should return not-found message")
	}

	if response.NotFound(nil, "").Success != false {
		t.Errorf("Success value should be false")
	}

	if response.NotFound(nil, "").Data != nil {
		t.Errorf("Data value should be nil")
	}
}

func TestConflictRespone(t *testing.T) {

	if response.Conflict(nil, "").Message != "Dose exist" {
		t.Errorf("Respose should return  Dose exist message")
	}

	if response.Conflict(nil, "Conflict").Message != "Conflict" {
		t.Errorf("Rspose should return Conflict message")
	}

	if response.Conflict(nil, "").Success != false {
		t.Errorf("Success value should be false")
	}

	if response.Conflict(nil, "").Data != nil {
		t.Errorf("Data value should be nil")
	}
}

