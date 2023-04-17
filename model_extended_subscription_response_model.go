/*
 * ElevenLabs API Documentation
 *
 * This is the documentation for the ElevenLabs API. You can use this API to use our service programmatically, this is done by using your xi-api-key. <br/> You can view your xi-api-key using the 'Profile' tab on https://beta.elevenlabs.io. Our API is experimental so all endpoints are subject to change.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ExtendedSubscriptionResponseModel struct {
	Tier string `json:"tier"`
	CharacterCount int32 `json:"character_count"`
	CharacterLimit int32 `json:"character_limit"`
	CanExtendCharacterLimit bool `json:"can_extend_character_limit"`
	AllowedToExtendCharacterLimit bool `json:"allowed_to_extend_character_limit"`
	NextCharacterCountResetUnix int32 `json:"next_character_count_reset_unix"`
	VoiceLimit int32 `json:"voice_limit"`
	ProfessionalVoiceLimit int32 `json:"professional_voice_limit"`
	CanExtendVoiceLimit bool `json:"can_extend_voice_limit"`
	CanUseInstantVoiceCloning bool `json:"can_use_instant_voice_cloning"`
	CanUseProfessionalVoiceCloning bool `json:"can_use_professional_voice_cloning"`
	AvailableModels []TtsModelResponseModel `json:"available_models"`
	CanUseDelayedPaymentMethods bool `json:"can_use_delayed_payment_methods"`
	Currency string `json:"currency"`
	Status string `json:"status"`
	NextInvoice *InvoiceResponseModel `json:"next_invoice"`
}