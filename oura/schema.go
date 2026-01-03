package oura

// DatePeriod struct for required date fields with api request.
type DatePeriod struct {
	StartDate string
	EndDate   string
}

// DateTimePeriod struct for datetime range queries (used by heartrate endpoint).
type DateTimePeriod struct {
	StartDateTime string
	EndDateTime   string
}

// PaginatedResponse is a generic wrapper for paginated API responses.
type PaginatedResponse[T any] struct {
	Data      []T     `json:"data"`
	NextToken *string `json:"next_token,omitempty"`
}

// PersonalInfo represents user personal information from v2 API.
type PersonalInfo struct {
	ID            string  `json:"id"`
	Age           int     `json:"age,omitempty"`
	Weight        float64 `json:"weight,omitempty"`
	Height        float64 `json:"height,omitempty"`
	BiologicalSex string  `json:"biological_sex,omitempty"`
	Email         string  `json:"email,omitempty"`
}

// RingConfiguration represents ring hardware and firmware details.
type RingConfiguration struct {
	ID              string `json:"id"`
	Color           string `json:"color,omitempty"`
	Design          string `json:"design,omitempty"`
	FirmwareVersion string `json:"firmware_version,omitempty"`
	HardwareType    string `json:"hardware_type,omitempty"`
	SetUpAt         string `json:"set_up_at,omitempty"`
	Size            int    `json:"size,omitempty"`
}

// DailySleepContributors represents the contributor scores for daily sleep.
type DailySleepContributors struct {
	DeepSleep   *int `json:"deep_sleep,omitempty"`
	Efficiency  *int `json:"efficiency,omitempty"`
	Latency     *int `json:"latency,omitempty"`
	RemSleep    *int `json:"rem_sleep,omitempty"`
	Restfulness *int `json:"restfulness,omitempty"`
	Timing      *int `json:"timing,omitempty"`
	TotalSleep  *int `json:"total_sleep,omitempty"`
}

// DailySleep represents daily sleep summary from v2 API.
type DailySleep struct {
	ID           string                  `json:"id"`
	Day          string                  `json:"day"`
	Score        *int                    `json:"score,omitempty"`
	Timestamp    string                  `json:"timestamp,omitempty"`
	Contributors *DailySleepContributors `json:"contributors,omitempty"`
}

// SampleData represents time-series sample data.
type SampleData struct {
	Interval  float64   `json:"interval"`
	Items     []float64 `json:"items"`
	Timestamp string    `json:"timestamp"`
}

// Sleep represents detailed sleep period data from v2 API.
type Sleep struct {
	ID                    string      `json:"id"`
	AverageBreath         *float64    `json:"average_breath,omitempty"`
	AverageHeartRate      *float64    `json:"average_heart_rate,omitempty"`
	AverageHRV            *int        `json:"average_hrv,omitempty"`
	AwakeTime             *int        `json:"awake_time,omitempty"`
	BedtimeEnd            string      `json:"bedtime_end,omitempty"`
	BedtimeStart          string      `json:"bedtime_start,omitempty"`
	Day                   string      `json:"day"`
	DeepSleepDuration     *int        `json:"deep_sleep_duration,omitempty"`
	Efficiency            *int        `json:"efficiency,omitempty"`
	HeartRate             *SampleData `json:"heart_rate,omitempty"`
	HRV                   *SampleData `json:"hrv,omitempty"`
	Latency               *int        `json:"latency,omitempty"`
	LightSleepDuration    *int        `json:"light_sleep_duration,omitempty"`
	LowBatteryAlert       bool        `json:"low_battery_alert,omitempty"`
	LowestHeartRate       *int        `json:"lowest_heart_rate,omitempty"`
	Movement30Sec         *string     `json:"movement_30_sec,omitempty"`
	Period                *int        `json:"period,omitempty"`
	ReadinessScoreDelta   *int        `json:"readiness_score_delta,omitempty"`
	REMSleepDuration      *int        `json:"rem_sleep_duration,omitempty"`
	RestlessPeriods       *int        `json:"restless_periods,omitempty"`
	SleepPhase5Min        *string     `json:"sleep_phase_5_min,omitempty"`
	SleepScoreDelta       *int        `json:"sleep_score_delta,omitempty"`
	TimeInBed             *int        `json:"time_in_bed,omitempty"`
	TotalSleepDuration    *int        `json:"total_sleep_duration,omitempty"`
	Type                  string      `json:"type,omitempty"`
}

// SleepTime represents sleep time recommendations.
type SleepTime struct {
	ID                string  `json:"id"`
	Day               string  `json:"day"`
	OptimalBedtime    *Window `json:"optimal_bedtime,omitempty"`
	Recommendation    string  `json:"recommendation,omitempty"`
	Status            string  `json:"status,omitempty"`
}

// Window represents a time window.
type Window struct {
	DayTz    int    `json:"day_tz,omitempty"`
	EndOffset int   `json:"end_offset,omitempty"`
	StartOffset int `json:"start_offset,omitempty"`
}

// DailyActivityContributors represents the contributor scores for daily activity.
type DailyActivityContributors struct {
	MeetDailyTargets  *int `json:"meet_daily_targets,omitempty"`
	MoveEveryHour     *int `json:"move_every_hour,omitempty"`
	RecoveryTime      *int `json:"recovery_time,omitempty"`
	StayActive        *int `json:"stay_active,omitempty"`
	TrainingFrequency *int `json:"training_frequency,omitempty"`
	TrainingVolume    *int `json:"training_volume,omitempty"`
}

// DailyActivity represents daily activity data from v2 API.
type DailyActivity struct {
	ID                      string                     `json:"id"`
	Day                     string                     `json:"day"`
	Score                   *int                       `json:"score,omitempty"`
	ActiveCalories          int                        `json:"active_calories,omitempty"`
	AverageMETMinutes       float64                    `json:"average_met_minutes,omitempty"`
	Contributors            *DailyActivityContributors `json:"contributors,omitempty"`
	EquivalentWalkingDistance int                      `json:"equivalent_walking_distance,omitempty"`
	HighActivityMETMinutes  int                        `json:"high_activity_met_minutes,omitempty"`
	HighActivityTime        int                        `json:"high_activity_time,omitempty"`
	InactivityAlerts        int                        `json:"inactivity_alerts,omitempty"`
	LowActivityMETMinutes   int                        `json:"low_activity_met_minutes,omitempty"`
	LowActivityTime         int                        `json:"low_activity_time,omitempty"`
	MediumActivityMETMinutes int                       `json:"medium_activity_met_minutes,omitempty"`
	MediumActivityTime      int                        `json:"medium_activity_time,omitempty"`
	Met                     *SampleData                `json:"met,omitempty"`
	MetersToTarget          int                        `json:"meters_to_target,omitempty"`
	NonWearTime             int                        `json:"non_wear_time,omitempty"`
	RestingTime             int                        `json:"resting_time,omitempty"`
	SedentaryMETMinutes     int                        `json:"sedentary_met_minutes,omitempty"`
	SedentaryTime           int                        `json:"sedentary_time,omitempty"`
	Steps                   int                        `json:"steps,omitempty"`
	TargetCalories          int                        `json:"target_calories,omitempty"`
	TargetMeters            int                        `json:"target_meters,omitempty"`
	Timestamp               string                     `json:"timestamp,omitempty"`
	TotalCalories           int                        `json:"total_calories,omitempty"`
}

// DailyReadinessContributors represents the contributor scores for daily readiness.
type DailyReadinessContributors struct {
	ActivityBalance     *int `json:"activity_balance,omitempty"`
	BodyTemperature     *int `json:"body_temperature,omitempty"`
	HRVBalance          *int `json:"hrv_balance,omitempty"`
	PreviousDayActivity *int `json:"previous_day_activity,omitempty"`
	PreviousNight       *int `json:"previous_night,omitempty"`
	RecoveryIndex       *int `json:"recovery_index,omitempty"`
	RestingHeartRate    *int `json:"resting_heart_rate,omitempty"`
	SleepBalance        *int `json:"sleep_balance,omitempty"`
}

// DailyReadiness represents daily readiness data from v2 API.
type DailyReadiness struct {
	ID                        string                      `json:"id"`
	Day                       string                      `json:"day"`
	Score                     *int                        `json:"score,omitempty"`
	TemperatureDeviation      *float64                    `json:"temperature_deviation,omitempty"`
	TemperatureTrendDeviation *float64                    `json:"temperature_trend_deviation,omitempty"`
	Timestamp                 string                      `json:"timestamp,omitempty"`
	Contributors              *DailyReadinessContributors `json:"contributors,omitempty"`
}

// HeartRate represents heart rate data point from v2 API.
type HeartRate struct {
	BPM       int    `json:"bpm"`
	Source    string `json:"source,omitempty"`
	Timestamp string `json:"timestamp"`
}

// DailySpo2 represents daily SpO2 (blood oxygen) data from v2 API.
type DailySpo2 struct {
	ID                string           `json:"id"`
	Day               string           `json:"day"`
	SpO2Percentage    *Spo2Percentage  `json:"spo2_percentage,omitempty"`
	BreathingDisturbanceIndex *float64 `json:"breathing_disturbance_index,omitempty"`
}

// Spo2Percentage represents SpO2 percentage range.
type Spo2Percentage struct {
	Average float64 `json:"average,omitempty"`
}

// DailyStress represents daily stress data from v2 API.
type DailyStress struct {
	ID               string `json:"id"`
	Day              string `json:"day"`
	StressHigh       *int   `json:"stress_high,omitempty"`
	RecoveryHigh     *int   `json:"recovery_high,omitempty"`
	DaySummary       string `json:"day_summary,omitempty"`
}

// DailyResilience represents daily resilience data from v2 API.
type DailyResilience struct {
	ID           string                       `json:"id"`
	Day          string                       `json:"day"`
	Level        string                       `json:"level,omitempty"`
	Contributors *DailyResilienceContributors `json:"contributors,omitempty"`
}

// DailyResilienceContributors represents resilience contributors.
type DailyResilienceContributors struct {
	SleepRecovery    float64 `json:"sleep_recovery,omitempty"`
	DaytimeRecovery  float64 `json:"daytime_recovery,omitempty"`
	Stress           float64 `json:"stress,omitempty"`
}

// DailyCardiovascularAge represents daily cardiovascular age data.
type DailyCardiovascularAge struct {
	ID     string `json:"id"`
	Day    string `json:"day"`
	VascularAge *int `json:"vascular_age,omitempty"`
}

// Vo2Max represents VO2 max data from v2 API.
type Vo2Max struct {
	ID        string   `json:"id"`
	Day       string   `json:"day"`
	Vo2Max    *float64 `json:"vo2_max,omitempty"`
	Timestamp string   `json:"timestamp,omitempty"`
}

// Workout represents workout data from v2 API.
type Workout struct {
	ID            string   `json:"id"`
	Activity      string   `json:"activity,omitempty"`
	Calories      *float64 `json:"calories,omitempty"`
	Day           string   `json:"day"`
	Distance      *float64 `json:"distance,omitempty"`
	EndDatetime   string   `json:"end_datetime,omitempty"`
	Intensity     string   `json:"intensity,omitempty"`
	Label         *string  `json:"label,omitempty"`
	Source        string   `json:"source,omitempty"`
	StartDatetime string   `json:"start_datetime,omitempty"`
}

// Session represents guided/unguided session data from v2 API.
type Session struct {
	ID                 string      `json:"id"`
	Day                string      `json:"day"`
	StartDatetime      string      `json:"start_datetime,omitempty"`
	EndDatetime        string      `json:"end_datetime,omitempty"`
	Type               string      `json:"type,omitempty"`
	HeartRate          *SampleData `json:"heart_rate,omitempty"`
	HRV                *SampleData `json:"hrv,omitempty"`
	Mood               *string     `json:"mood,omitempty"`
	MotionCount        *SampleData `json:"motion_count,omitempty"`
}

// Tag represents user-entered tag data (deprecated, use EnhancedTag).
type Tag struct {
	ID        string   `json:"id"`
	Day       string   `json:"day"`
	Text      *string  `json:"text,omitempty"`
	Timestamp string   `json:"timestamp,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// EnhancedTag represents enhanced tag data from v2 API.
type EnhancedTag struct {
	ID             string  `json:"id"`
	Day            string  `json:"day"`
	TagTypeCode    *string `json:"tag_type_code,omitempty"`
	StartTime      *string `json:"start_time,omitempty"`
	EndTime        *string `json:"end_time,omitempty"`
	StartDay       *string `json:"start_day,omitempty"`
	EndDay         *string `json:"end_day,omitempty"`
	Comment        *string `json:"comment,omitempty"`
}

// RestModePeriod represents rest mode period data from v2 API.
type RestModePeriod struct {
	ID            string    `json:"id"`
	EndDay        *string   `json:"end_day,omitempty"`
	EndTime       *string   `json:"end_time,omitempty"`
	Episodes      []Episode `json:"episodes,omitempty"`
	StartDay      string    `json:"start_day,omitempty"`
	StartTime     *string   `json:"start_time,omitempty"`
}

// Episode represents a rest mode episode.
type Episode struct {
	Tags      []string `json:"tags,omitempty"`
	Timestamp string   `json:"timestamp,omitempty"`
}

// Webhook types for subscription management.

// WebhookSubscription represents a webhook subscription.
type WebhookSubscription struct {
	ID            string `json:"id"`
	CallbackURL   string `json:"callback_url"`
	EventType     string `json:"event_type"`
	DataType      string `json:"data_type"`
	ExpirationTime string `json:"expiration_time,omitempty"`
}

// WebhookSubscriptionRequest represents a request to create/update a webhook subscription.
type WebhookSubscriptionRequest struct {
	CallbackURL     string `json:"callback_url"`
	VerificationToken string `json:"verification_token"`
	EventType       string `json:"event_type"`
	DataType        string `json:"data_type"`
}
