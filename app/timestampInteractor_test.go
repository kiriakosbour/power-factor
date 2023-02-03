package app

import (
	"errors"
	"power-factor/domain"
	"reflect"
	"testing"
)

func Test_TimestampsMatching_1h_OK(t *testing.T) {
	period := "1h"
	timezone := "Europe/Athens"
	t1 := "20210714T204603Z"
	t2 := "20210715T123456Z"
	ts := domain.Timestamp{
		Interval: period,
		Timezone: timezone,
		Start:    t1,
		End:      t2,
	}
	service := NewTimestampInteractor()

	result, err := service.TimestampService(ts)
	if err.Desc != "" {
		t.Error("Unexpected error: ", err)
	}
	if len(result) != len(TestTimestampsmatching1hOkData) {
		t.Errorf("Result items %d not equal to expected items %d", len(result), len(TestTimestampsmatching1hOkData))
	}
	if !reflect.DeepEqual(result, TestTimestampsmatching1hOkData) {
		t.Errorf("Result data not matching expected data")
	}
}

func Test_TimestampsMatching_1mo_OK(t *testing.T) {
	period := "1mo"
	timezone := "Europe/Athens"
	t1 := "20210214T204603Z"
	t2 := "20211115T123456Z"
	ts := domain.Timestamp{
		Interval: period,
		Timezone: timezone,
		Start:    t1,
		End:      t2,
	}
	service := NewTimestampInteractor()

	result, err := service.TimestampService(ts)
	if err.Desc != "" {
		t.Error("Unexpected error: ", err)
	}
	if len(result) != len(TestTimestampsmatching1moOkData) {
		t.Errorf("Result items %d not equal to expected items %d", len(result), len(TestTimestampsmatching1moOkData))
	}
	if !reflect.DeepEqual(result, TestTimestampsmatching1moOkData) {
		t.Errorf("Result data not matching expected data")
	}
}

func Test_TimestampsMatching_UnsupportedPeriod(t *testing.T) {
	period := "1w"
	timezone := "Europe/Athens"
	t1 := "20210714T204603Z"
	t2 := "20210715T123456Z"
	ts := domain.Timestamp{
		Interval: period,
		Timezone: timezone,
		Start:    t1,
		End:      t2,
	}
	service := NewTimestampInteractor()
	expectedError := errors.New("unsupported period")
	_, err := service.TimestampService(ts)
	if err.Desc == "" {
		t.Error("Expected error: ")
	}
	if err.Desc != expectedError.Error() {
		t.Errorf("Error %s not equal to expected error %s", err.Desc, expectedError.Error())
	}
}

func Test_TimestampsMatching_InvalidPeriod(t *testing.T) {
	period := "1w"
	timezone := "Europe/Athens"
	t1 := "20210715T123456Z"
	t2 := "20210714T204603Z"
	ts := domain.Timestamp{
		Interval: period,
		Timezone: timezone,
		Start:    t1,
		End:      t2,
	}
	service := NewTimestampInteractor()
	expectedError := errors.New("start date cannot be before end date")
	_, err := service.TimestampService(ts)
	if err.Desc == "" {
		t.Error("Expected error: ")
	}
	if err.Desc != expectedError.Error() {
		t.Errorf("Error %s not equal to expected error %s", err.Desc, expectedError.Error())
	}
}

func Test_TimestampsMatching_InvalidTimezone(t *testing.T) {
	period := "1d"
	timezone := "Asia/Athens"
	t1 := "20210714T204603Z"
	t2 := "20210715T123456Z"
	ts := domain.Timestamp{
		Interval: period,
		Timezone: timezone,
		Start:    t1,
		End:      t2,
	}
	service := NewTimestampInteractor()
	expectedError := errors.New("illegal location provided")
	_, err := service.TimestampService(ts)
	if err.Desc == "" {
		t.Error("Expected error: ")
	}
	if err.Desc != expectedError.Error() {
		t.Errorf("Error %s not equal to expected error %s", err.Desc, expectedError.Error())
	}
}
