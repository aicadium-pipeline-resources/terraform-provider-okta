package sdk

type OpenIdConnectApplicationSettings struct {
	App                *ApplicationSettingsApplication         `json:"app,omitempty"`
	ImplicitAssignment *bool                                   `json:"implicitAssignment,omitempty"`
	InlineHookId       string                                  `json:"inlineHookId,omitempty"`
	Notes              *ApplicationSettingsNotes               `json:"notes,omitempty"`
	Notifications      *ApplicationSettingsNotifications       `json:"notifications,omitempty"`
	OauthClient        *OpenIdConnectApplicationSettingsClient `json:"oauthClient,omitempty"`
}