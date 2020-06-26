package utils

type ZoomMeeting struct {
	ClientEmail  string `json:"clientEmail,omitempty"`
	UserEmail    string `json:"userEmail,omitempty"`
	SelectedTime string `json:"selectedTime,omitempty"`
}