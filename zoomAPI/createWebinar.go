package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/webinars/webinarcreate
*/
func (client Client) CreateWebinar(userId string,
	topic string,
	webinarType int,
	startTime string,
	duration int,
	timezone string,
	password string,
	agenda string,
	recurrence *Recurrence,
	settings *Settings) (createWebinarResponse CreateWebinarResponse, err error) {

	if recurrence == nil {
		recurrence = &Recurrence{
			Type:           1,
			RepeatInterval: 0,
			WeeklyDays:     "",
			MonthlyDay:     0,
			MonthlyWeek:    0,
			MonthlyWeekDay: 0,
			EndTimes:       0,
			EndDateTime:    "",
		}
	}

	if settings == nil {
		settings = &Settings{
			HostVideo:                    false,
			ParticipantVideo:             false,
			CnMeeting:                    false,
			InMeeting:                    false,
			JoinBeforeHost:               false,
			MuteUponEntry:                false,
			Watermark:                    false,
			UsePmi:                       false,
			ApprovalType:                 0,
			RegistrationType:             0,
			Audio:                        "",
			AutoRecording:                "",
			EnforceLogin:                 false,
			EnforceLoginDomains:          "",
			AlternativeHosts:             "",
			GlobalDialInCountries:        nil,
			RegistrantsEmailNotification: false,
		}
	}

	createWebinarRequest := CreateWebinarRequest{
		Topic:       topic,
		Type:        webinarType,
		StartTime:   startTime,
		Duration:    duration,
		Timezone:    timezone,
		Password:    password,
		Agenda:      agenda,
		Recurrence:  *recurrence,
		Settings:    *settings,
	}

	var reqBody []byte
	reqBody, err = json.Marshal(createWebinarRequest)
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("/users/%s/webinars", userId)
	httpMethod := http.MethodPost

	var b []byte
	b, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &createWebinarResponse)
	if err != nil {
		return
	}

	return

}
