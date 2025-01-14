package sdk

type SwaApplicationSettingsApplication struct {
	ButtonField   string `json:"buttonField,omitempty"`
	Checkbox      string `json:"checkbox,omitempty"`
	LoginUrlRegex string `json:"loginUrlRegex,omitempty"`
	PasswordField string `json:"passwordField,omitempty"`
	RedirectUrl   string `json:"redirectUrl,omitempty"`
	Url           string `json:"url,omitempty"`
	UsernameField string `json:"usernameField,omitempty"`
}
