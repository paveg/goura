package api

import "time"

// Age represents user age
type Age int

// Weight represents user weight
type Weight float32

// Height represents user height
type Height float32

// Gender represents user gender
type Gender string

// Email represents user email address
type Email string

// UserID represents user identifier
type UserID string

// UserInfo represents user information
type UserInfo struct {
	Age    Age    `json:"age"`
	Weight Weight `json:"weight"`
	Height Height `json:"height"`
	Gender Gender `json:"gender"`
	Email  Email  `json:"email"`
	UserID UserID `json:"user_id"`
}

// SleepPeriods represents user sleep period information
type SleepPeriods struct {
	SleepPeriods []Sleep `json:"sleep"`
}

// Sleep represents user sleep information
// TODO: implements under construction
// @see: https://cloud.ouraring.com/docs/sleep
//  "hr_5min": [0, 53, 51, 0, 50, 50, 49, 49, 50, 50, 51, 52, 52, 51, 53, 58, 60, 60, 59, 58, 58, 58, 58, 55, 55, 55, 55, 56, 56, 55, 53, 53, 53, 53, 53, 53, 57, 58, 60, 60, 59, 57, 59, 58, 56, 56, 56, 56, 55, 55, 56, 56, 57, 58, 55, 56, 57, 60, 58, 58, 59, 57, 54, 54, 53, 52, 52, 55, 53, 54, 56, 0]
//  "rmssd_5min": [0, 0, 62, 0, 75, 52, 56, 56, 64, 57, 55, 78, 77, 83, 70, 35, 21, 25, 49, 44, 48, 48, 62, 69, 66, 64, 79, 59, 67, 66, 70, 63, 53, 57, 53, 57, 38, 26, 18, 24, 30, 35, 36, 46, 53, 59, 50, 50, 53, 53, 57, 52, 41, 37, 49, 47, 48, 35, 32, 34, 52, 57, 62, 57, 70, 81, 81, 65, 69, 72, 64, 0]
type Sleep struct {
	SummaryDate       string  `json:"summary_date"`
	PeriodID          int     `json:"period_id"`
	IsLongest         int     `json:"is_longest"`
	TimeZone          int     `json:"time_zone"`
	BedTimeStart      string  `json:"bedtime_start"`
	BedTimeEnd        string  `json:"bedtime_end"`
	Score             int     `json:"score"`
	ScoreTotal        int     `json:"score_total"`
	ScoreDisturbances int     `json:"score_disturbances"`
	ScoreEfficiency   int     `json:"score_efficiency"`
	ScoreLatency      int     `json:"score_latency"`
	ScoreRem          int     `json:"score_rem"`
	ScoreDeep         int     `json:"score_deep"`
	ScoreAlignment    int     `json:"score_alignment"`
	Total             int     `json:"total"`
	Duration          int     `json:"duration"`
	Awake             int     `json:"awake"`
	Light             int     `json:"light"`
	Rem               int     `json:"rem"`
	Deep              int     `json:"deep"`
	OnsetLatency      int     `json:"onset_latency"`
	Restless          int     `json:"restless"`
	Efficiency        int     `json:"efficiency"`
	MidpointTime      int     `json:"midpoint_time"`
	HrLowest          int     `json:"hr_lowest"`
	HrAverage         float64 `json:"hr_average"`
	Rmssd             int     `json:"rmssd"`
	BreathAverage     float64 `json:"breath_average"`
	TemperatureDelta  float64 `json:"temperature_delta"`
	Hypnogram5min     string  `json:"hypnogram_5min"`
}

// DatePeriod struct for required date fields with api request.
type DatePeriod struct {
	StartDate time.Time
	EndDate   time.Time
}
