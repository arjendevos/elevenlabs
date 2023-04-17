/*
 * ElevenLabs API Documentation
 *
 * This is the documentation for the ElevenLabs API. You can use this API to use our service programmatically, this is done by using your xi-api-key. <br/> You can view your xi-api-key using the 'Profile' tab on https://beta.elevenlabs.io. Our API is experimental so all endpoints are subject to change.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type HistoryItemResponseModel struct {
	HistoryItemId string `json:"history_item_id"`
	RequestId string `json:"request_id"`
	VoiceId string `json:"voice_id"`
	VoiceName string `json:"voice_name"`
	Text string `json:"text"`
	DateUnix int32 `json:"date_unix"`
	CharacterCountChangeFrom int32 `json:"character_count_change_from"`
	CharacterCountChangeTo int32 `json:"character_count_change_to"`
	ContentType string `json:"content_type"`
	State string `json:"state"`
	Settings *interface{} `json:"settings"`
	Feedback *FeedbackResponseModel `json:"feedback"`
}