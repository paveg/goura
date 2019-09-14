package api

import (
	"context"
	"fmt"
	"log"

	"github.com/paveg/goura/oura"
)

// GetUserInfo requests GET /v1/userinfo
func (client *Client) GetUserInfo(ctx context.Context) (*oura.UserInfo, error) {
	subURL := fmt.Sprint("/v1/userinfo")
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	res, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse oura.UserInfo
	log.Printf("HTTP Request: %v", res.Status)
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetSleep requests GET /v1/sleep
func (client *Client) GetSleep(ctx context.Context, datePeriod oura.DatePeriod) (*oura.SleepPeriods, error) {
	subURL := fmt.Sprintf("/v1/sleep")
	var apiResponse oura.SleepPeriods
	res, err := client.getRequestWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}
	log.Printf("HTTP Request: %v", res.Status)
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetActivity requests GET /v1/activity
func (client *Client) GetActivity(ctx context.Context, datePeriod oura.DatePeriod) (*oura.Activities, error) {
	subURL := fmt.Sprintf("/v1/activity")
	var apiResponse oura.Activities
	res, err := client.getRequestWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}
	log.Printf("HTTP Request: %v", res.Status)
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetReadiness requests GET /v1/readiness
func (client *Client) GetReadiness(ctx context.Context, datePeriod oura.DatePeriod) (*oura.Readinesses, error) {
	subURL := fmt.Sprintf("/v1/readiness")
	var apiResponse oura.Readinesses
	res, err := client.getRequestWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}
	log.Printf("HTTP Request: %v", res.Status)
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
