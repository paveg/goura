package api

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
	// Hr5min            []int   `json:"hr_5min"`
	// Rmssd5min         []int   `json:"rmssd_5min"`
}

// DatePeriod struct for required date fields with api request.
type DatePeriod struct {
	StartDate string
	EndDate   string
}

// Activities represents user activities
type Activities struct {
	Activity []Activity `json:"activity"`
}

// Activity represents user activity
type Activity struct {
	SummaryDate            string  `json:"summary_date"`
	DayStart               string  `json:"day_start"`
	DayEnd                 string  `json:"day_end"`
	TimeZone               int     `json:"timezone"`
	Score                  int     `json:"score"`
	ScoreStayActive        int     `json:"score_stay_active"`
	ScoreMoveEveryHour     int     `json:"score_move_every_hour"`
	ScoreMeetDailyTargets  int     `json:"score_meet_daily_targets"`
	ScoreTrainingFrequency int     `json:"score_training_frequency"`
	ScoreTrainingVolume    int     `json:"score_training_volume"`
	ScoreRecoveryTime      int     `json:"score_recovery_time"`
	DailyMovement          int     `json:"daily_movement"`
	NonWear                int     `json:"non_wear"`
	Rest                   int     `json:"rest"`
	Inactive               int     `json:"inactive"`
	InactivityAlerts       int     `json:"inactivity_alerts"`
	Low                    int     `json:"low"`
	Medium                 int     `json:"medium"`
	High                   int     `json:"high"`
	Steps                  int     `json:"steps"`
	CalTotal               int     `json:"cal_total"`
	CalActive              int     `json:"cal_active"`
	MetMinInactive         int     `json:"met_min_inactive"`
	MetMinLow              int     `json:"met_min_low"`
	MetMinMediumPlus       int     `json:"met_min_medium_plus"`
	MetMinMedium           int     `json:"met_min_medium"`
	MetMinHigh             int     `json:"met_min_high"`
	AverageMet             float64 `json:"average_met"`
	Class5min              string  `json:"class_5min"`
	// Met1min                []float64 `json:"met_1min"`
}
