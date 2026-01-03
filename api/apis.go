package api

import (
	"context"
	"fmt"

	"github.com/paveg/goura/oura"
)

// GetPersonalInfo requests GET /v2/usercollection/personal_info
func (client *Client) GetPersonalInfo(ctx context.Context) (*oura.PersonalInfo, error) {
	subURL := fmt.Sprintf("%s/personal_info", V2BasePath)
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PersonalInfo
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetRingConfiguration requests GET /v2/usercollection/ring_configuration
func (client *Client) GetRingConfiguration(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.RingConfiguration], error) {
	subURL := fmt.Sprintf("%s/ring_configuration", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.RingConfiguration]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetDailySleep requests GET /v2/usercollection/daily_sleep
func (client *Client) GetDailySleep(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.DailySleep], error) {
	subURL := fmt.Sprintf("%s/daily_sleep", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.DailySleep]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetSleep requests GET /v2/usercollection/sleep (detailed sleep periods)
func (client *Client) GetSleep(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.Sleep], error) {
	subURL := fmt.Sprintf("%s/sleep", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.Sleep]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetSleepTime requests GET /v2/usercollection/sleep_time
func (client *Client) GetSleepTime(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.SleepTime], error) {
	subURL := fmt.Sprintf("%s/sleep_time", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.SleepTime]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetDailyActivity requests GET /v2/usercollection/daily_activity
func (client *Client) GetDailyActivity(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.DailyActivity], error) {
	subURL := fmt.Sprintf("%s/daily_activity", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.DailyActivity]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetDailyReadiness requests GET /v2/usercollection/daily_readiness
func (client *Client) GetDailyReadiness(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.DailyReadiness], error) {
	subURL := fmt.Sprintf("%s/daily_readiness", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.DailyReadiness]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetHeartRate requests GET /v2/usercollection/heartrate
func (client *Client) GetHeartRate(ctx context.Context, period oura.DateTimePeriod) (*oura.PaginatedResponse[oura.HeartRate], error) {
	subURL := fmt.Sprintf("%s/heartrate", V2BasePath)
	res, err := client.getWithDateTimePeriod(ctx, subURL, period)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.HeartRate]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetDailySpo2 requests GET /v2/usercollection/daily_spo2
func (client *Client) GetDailySpo2(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.DailySpo2], error) {
	subURL := fmt.Sprintf("%s/daily_spo2", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.DailySpo2]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetDailyStress requests GET /v2/usercollection/daily_stress
func (client *Client) GetDailyStress(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.DailyStress], error) {
	subURL := fmt.Sprintf("%s/daily_stress", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.DailyStress]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetDailyResilience requests GET /v2/usercollection/daily_resilience
func (client *Client) GetDailyResilience(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.DailyResilience], error) {
	subURL := fmt.Sprintf("%s/daily_resilience", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.DailyResilience]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetDailyCardiovascularAge requests GET /v2/usercollection/daily_cardiovascular_age
func (client *Client) GetDailyCardiovascularAge(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.DailyCardiovascularAge], error) {
	subURL := fmt.Sprintf("%s/daily_cardiovascular_age", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.DailyCardiovascularAge]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetVo2Max requests GET /v2/usercollection/vO2_max
func (client *Client) GetVo2Max(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.Vo2Max], error) {
	subURL := fmt.Sprintf("%s/vO2_max", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.Vo2Max]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetWorkouts requests GET /v2/usercollection/workout
func (client *Client) GetWorkouts(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.Workout], error) {
	subURL := fmt.Sprintf("%s/workout", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.Workout]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetSessions requests GET /v2/usercollection/session
func (client *Client) GetSessions(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.Session], error) {
	subURL := fmt.Sprintf("%s/session", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.Session]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetTags requests GET /v2/usercollection/tag (deprecated, use GetEnhancedTags)
func (client *Client) GetTags(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.Tag], error) {
	subURL := fmt.Sprintf("%s/tag", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.Tag]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetEnhancedTags requests GET /v2/usercollection/enhanced_tag
func (client *Client) GetEnhancedTags(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.EnhancedTag], error) {
	subURL := fmt.Sprintf("%s/enhanced_tag", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.EnhancedTag]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// GetRestModePeriods requests GET /v2/usercollection/rest_mode_period
func (client *Client) GetRestModePeriods(ctx context.Context, datePeriod oura.DatePeriod) (*oura.PaginatedResponse[oura.RestModePeriod], error) {
	subURL := fmt.Sprintf("%s/rest_mode_period", V2BasePath)
	res, err := client.getWithDatePeriod(ctx, subURL, datePeriod)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	var apiResponse oura.PaginatedResponse[oura.RestModePeriod]
	if err := decodeBody(res, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
