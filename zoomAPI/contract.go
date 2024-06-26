package zoomAPI

import "time"

type ListMeetingsAPIResponse struct {
	PageCount    int       `json:"page_count"`
	PageNumber   int       `json:"page_number"`
	PageSize     int       `json:"page_size"`
	TotalRecords int       `json:"total_records"`
	Meetings     []Meeting `json:"meetings"`
}

type Meeting struct {
	Uuid      string    `json:"uuid"`
	Id        int       `json:"id"`
	HostId    string    `json:"host_id"`
	Topic     string    `json:"topic"`
	Type      int       `json:"type"`
	StartTime time.Time `json:"start_time"`
	Duration  int       `json:"duration"`
	Timezone  string    `json:"timezone"`
	CreatedAt time.Time `json:"created_at"`
	JoinUrl   string    `json:"join_url"`
	Agenda    string    `json:"agenda,omitempty"`
}

type CreateMeetingRequest struct {
	Topic       string     `json:"topic"`
	Type        int        `json:"type"`
	StartTime   string     `json:"start_time"`
	Duration    int        `json:"duration"`
	ScheduleFor string     `json:"schedule_for"`
	Timezone    string     `json:"timezone"`
	Password    string     `json:"password"`
	Agenda      string     `json:"agenda"`
	Recurrence  Recurrence `json:"recurrence"`
	Settings    Settings   `json:"settings"`
}

type CreateWebinarRequest struct {
	Topic      string     `json:"topic"`
	Type       int        `json:"type"`
	StartTime  string     `json:"start_time"`
	Duration   int        `json:"duration"`
	Timezone   string     `json:"timezone"`
	Password   string     `json:"password"`
	Agenda     string     `json:"agenda"`
	Recurrence Recurrence `json:"recurrence"`
	Settings   Settings   `json:"settings"`
}

type Recurrence struct {
	Type           int    `json:"type"`
	RepeatInterval int    `json:"repeat_interval"`
	WeeklyDays     string `json:"weekly_days"`
	MonthlyDay     int    `json:"monthly_day"`
	MonthlyWeek    int    `json:"monthly_week"`
	MonthlyWeekDay int    `json:"monthly_week_day"`
	EndTimes       int    `json:"end_times"`
	EndDateTime    string `json:"end_date_time"`
}

type Settings struct {
	HostVideo                    bool     `json:"host_video"`
	ParticipantVideo             bool     `json:"participant_video"`
	CnMeeting                    bool     `json:"cn_meeting"`
	InMeeting                    bool     `json:"in_meeting"`
	JoinBeforeHost               bool     `json:"join_before_host"`
	MuteUponEntry                bool     `json:"mute_upon_entry"`
	Watermark                    bool     `json:"watermark"`
	UsePmi                       bool     `json:"use_pmi"`
	ApprovalType                 int      `json:"approval_type"`
	RegistrationType             int      `json:"registration_type"`
	Audio                        string   `json:"audio"`
	AutoRecording                string   `json:"auto_recording"`
	EnforceLogin                 bool     `json:"enforce_login"`
	EnforceLoginDomains          string   `json:"enforce_login_domains"`
	AlternativeHosts             string   `json:"alternative_hosts"`
	GlobalDialInCountries        []string `json:"global_dial_in_countries"`
	RegistrantsEmailNotification bool     `json:"registrants_email_notification"`
}

type CreateMeetingResponse struct {
	CreatedAt time.Time `json:"created_at"`
	Duration  int       `json:"duration"`
	HostId    string    `json:"host_id"`
	Id        int       `json:"id"`
	JoinUrl   string    `json:"join_url"`
	Settings  struct {
		AlternativeHosts      string   `json:"alternative_hosts"`
		ApprovalType          int      `json:"approval_type"`
		Audio                 string   `json:"audio"`
		AutoRecording         string   `json:"auto_recording"`
		CloseRegistration     bool     `json:"close_registration"`
		CnMeeting             bool     `json:"cn_meeting"`
		EnforceLogin          bool     `json:"enforce_login"`
		EnforceLoginDomains   string   `json:"enforce_login_domains"`
		GlobalDialInCountries []string `json:"global_dial_in_countries"`
		GlobalDialInNumbers   []struct {
			City        string `json:"city"`
			Country     string `json:"country"`
			CountryName string `json:"country_name"`
			Number      string `json:"number"`
			Type        string `json:"type"`
		} `json:"global_dial_in_numbers"`
		HostVideo                    bool `json:"host_video"`
		InMeeting                    bool `json:"in_meeting"`
		JoinBeforeHost               bool `json:"join_before_host"`
		MuteUponEntry                bool `json:"mute_upon_entry"`
		ParticipantVideo             bool `json:"participant_video"`
		RegistrantsConfirmationEmail bool `json:"registrants_confirmation_email"`
		UsePmi                       bool `json:"use_pmi"`
		WaitingRoom                  bool `json:"waiting_room"`
		Watermark                    bool `json:"watermark"`
		RegistrantsEmailNotification bool `json:"registrants_email_notification"`
	} `json:"settings"`
	StartTime time.Time `json:"start_time"`
	StartUrl  string    `json:"start_url"`
	Status    string    `json:"status"`
	Timezone  string    `json:"timezone"`
	Topic     string    `json:"topic"`
	Type      int       `json:"type"`
	UUID      string    `json:"uuid"`
}

type CreateWebinarResponse struct {
	UUID                         string    `json:"uuid"`
	Id                           int       `json:"id"`
	HostId                       string    `json:"host_id"`
	HostEmail                    string    `json:"host_email"`
	RegistrantsConfirmationEmail bool      `json:"registrants_confirmation_email"`
	TemplateID                   string    `json:"template_id"`
	Topic                        string    `json:"topic"`
	Type                         int       `json:"type"`
	StartTime                    time.Time `json:"start_time"`
	Duration                     int       `json:"duration"`
	Timezone                     string    `json:"timezone"`
	Agenda                       string    `json:"agenda"`
	CreatedAt                    time.Time `json:"created_at"`
	StartUrl                     string    `json:"start_url"`
	JoinUrl                      string    `json:"join_url"`
	TrackingFields               []struct {
		Field string `json:"field"`
		Value string `json:"value"`
	} `json:"tracking_fields"`
	Occurrences []struct {
		OccurrenceID string    `json:"occurrence_id"`
		StartTime    time.Time `json:"start_time"`
		Duration     int       `json:"duration"`
		Status       string    `json:"status"`
	} `json:"occurrences"`
	Settings struct {
		HostVideo                    bool     `json:"host_video"`
		PanelistsVideo               bool     `json:"panelists_video"`
		PracticeSession              bool     `json:"practice_session"`
		HDVideo                      bool     `json:"hd_video"`
		HDVideoForAttendees          bool     `json:"hd_video_for_attendees"`
		Send1080pVideoForAttendees   bool     `json:"send_1080p_video_for_attendees"`
		ApprovalType                 int      `json:"approval_type"`
		RegistrationType             int      `json:"registration_type"`
		Audio                        string   `json:"audio"`
		AutoRecording                string   `json:"auto_recording"`
		EnforceLogin                 bool     `json:"enforce_login"`
		EnforceLoginDomains          string   `json:"enforce_login_domains"`
		AlternativeHosts             string   `json:"alternative_hosts"`
		CloseRegistration            bool     `json:"close_registration"`
		ShowShareButton              bool     `json:"show_share_button"`
		AllowMultipleDevices         bool     `json:"allow_multiple_devices"`
		OnDemand                     bool     `json:"on_demand"`
		GlobalDialInCountries        []string `json:"global_dial_in_countries"`
		ContactName                  string   `json:"contact_name"`
		ContactEmail                 string   `json:"contact_email"`
		RegistrantsConfirmationEmail bool     `json:"registrants_confirmation_email"`
		RegistrantsRestrictNumber    int     `json:"registrants_restrict_number"`
		NotifyRegistrants            bool     `json:"notify_registrants"`
		PostWebinarSurvey            bool     `json:"post_webinar_survey"`
		SurveyURL                    string   `json:"survey_url"`
		RegistrantsEmailNotification bool     `json:"registrants_email_notification"`
		MeetingAuthentication        bool     `json:"meeting_authentication"`
		AuthenticationOption         string   `json:"authentication_option"`
		AuthenticationDomains        string   `json:"authentication_domains"`
		AuthenticationName           string   `json:"authentication_name"`
		// question_and_answer
		EmailLanguage                        string `json:"email_language"`
		PanelistsInvitationEmailNotification bool   `json:"panelists_invitation_email_notification"`
		// attendees_and_panelists_reminder_email_notification
		// follow_up_attendees_email_notification
		// follow_up_absentees_email_notification
	} `json:"settings"`
	// recurrence
	Password string `json:"password"`
}

type GetMeetingResponse struct {
	Agenda    string    `json:"agenda"`
	CreatedAt time.Time `json:"created_at"`
	Duration  int       `json:"duration"`
	HostId    string    `json:"host_id"`
	Id        int       `json:"id"`
	JoinUrl   string    `json:"join_url"`
	Settings  struct {
		AlternativeHosts      string   `json:"alternative_hosts"`
		ApprovalType          int      `json:"approval_type"`
		Audio                 string   `json:"audio"`
		AutoRecording         string   `json:"auto_recording"`
		CloseRegistration     bool     `json:"close_registration"`
		CnMeeting             bool     `json:"cn_meeting"`
		EnforceLogin          bool     `json:"enforce_login"`
		EnforceLoginDomains   string   `json:"enforce_login_domains"`
		GlobalDialInCountries []string `json:"global_dial_in_countries"`
		GlobalDialInNumbers   []struct {
			City        string `json:"city"`
			Country     string `json:"country"`
			CountryName string `json:"country_name"`
			Number      string `json:"number"`
			Type        string `json:"type"`
		} `json:"global_dial_in_numbers"`
		HostVideo                    bool `json:"host_video"`
		InMeeting                    bool `json:"in_meeting"`
		JoinBeforeHost               bool `json:"join_before_host"`
		MuteUponEntry                bool `json:"mute_upon_entry"`
		ParticipantVideo             bool `json:"participant_video"`
		RegistrantsConfirmationEmail bool `json:"registrants_confirmation_email"`
		UsePmi                       bool `json:"use_pmi"`
		WaitingRoom                  bool `json:"waiting_room"`
		Watermark                    bool `json:"watermark"`
		RegistrantsEmailNotification bool `json:"registrants_email_notification"`
	} `json:"settings"`
	StartTime time.Time `json:"start_time"`
	StartUrl  string    `json:"start_url"`
	Status    string    `json:"status"`
	Timezone  string    `json:"timezone"`
	Topic     string    `json:"topic"`
	Type      int       `json:"type"`
	Uuid      string    `json:"uuid"`
}

type GetMeetingInvitationResponse struct {
	Invitation string `json:"invitation"`
}

type CustomQuestion struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type AddMeetingRegistrantRequest struct {
	Email                 string           `json:"email"`
	FirstName             string           `json:"first_name"`
	LastName              string           `json:"last_name"`
	Address               string           `json:"address"`
	City                  string           `json:"city"`
	Country               string           `json:"country"`
	Zip                   string           `json:"zip"`
	State                 string           `json:"state"`
	Phone                 string           `json:"phone"`
	Industry              string           `json:"industry"`
	Org                   string           `json:"org"`
	JobTitle              string           `json:"job_title"`
	PurchasingTimeFrame   string           `json:"purchasing_time_frame"`
	RoleInPurchaseProcess string           `json:"role_in_purchase_process"`
	NoOfEmployees         string           `json:"no_of_employees"`
	Comments              string           `json:"comments"`
	CustomQuestions       []CustomQuestion `json:"custom_questions"`
}

type AddMeetingRegistrantResponse struct {
	Id           int    `json:"id"`
	JoinUrl      string `json:"join_url"`
	RegistrantId string `json:"registrant_id"`
	StartTime    string `json:"start_time"`
	Topic        string `json:"topic"`
}

type MeetingRegistrant struct {
	Email                 string           `json:"email"`
	FirstName             string           `json:"first_name"`
	LastName              string           `json:"last_name"`
	Id                    string           `json:"id"`
	Address               string           `json:"address"`
	City                  string           `json:"city"`
	Country               string           `json:"country"`
	Zip                   string           `json:"zip"`
	State                 string           `json:"state"`
	Phone                 string           `json:"phone"`
	Industry              string           `json:"industry"`
	Org                   string           `json:"org"`
	JobTitle              string           `json:"job_title"`
	PurchasingTimeFrame   string           `json:"purchasing_time_frame"`
	RoleInPurchaseProcess string           `json:"role_in_purchase_process"`
	NoOfEmployees         string           `json:"no_of_employees"`
	Comments              string           `json:"comments"`
	CustomQuestion        []CustomQuestion `json:"custom_questions"`
	Status                string           `json:"status"`
	CreateTime            time.Time        `json:"create_time"`
	JoinUrl               string           `json:"join_url"`
}

type ListMeetingRegistrantsResponse struct {
	PageCount    int                 `json:"page_count"`
	PageNumber   int                 `json:"page_number"`
	PageSize     int                 `json:"page_size"`
	TotalRecords int                 `json:"total_records"`
	Registrants  []MeetingRegistrant `json:"registrants"`
}

type UpdateMeetingStatusRequest struct {
	Action string `json:"action"`
}
