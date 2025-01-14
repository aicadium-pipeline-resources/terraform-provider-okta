package sdk

import "encoding/json"

type PasswordPolicyPasswordSettingsComplexity struct {
	Dictionary        *PasswordDictionary `json:"dictionary,omitempty"`
	ExcludeAttributes []string            `json:"excludeAttributes,omitempty"`
	ExcludeUsername   *bool               `json:"excludeUsername,omitempty"`
	MinLength         int64               `json:"-"`
	MinLengthPtr      *int64              `json:"minLength,omitempty"`
	MinLowerCase      int64               `json:"-"`
	MinLowerCasePtr   *int64              `json:"minLowerCase"`
	MinNumber         int64               `json:"-"`
	MinNumberPtr      *int64              `json:"minNumber"`
	MinSymbol         int64               `json:"-"`
	MinSymbolPtr      *int64              `json:"minSymbol"`
	MinUpperCase      int64               `json:"-"`
	MinUpperCasePtr   *int64              `json:"minUpperCase"`
}

func NewPasswordPolicyPasswordSettingsComplexity() *PasswordPolicyPasswordSettingsComplexity {
	return &PasswordPolicyPasswordSettingsComplexity{
		ExcludeAttributes: []string{},
		ExcludeUsername:   boolPtr(true),
		MinLength:         8,
		MinLengthPtr:      Int64Ptr(8),
		MinLowerCase:      1,
		MinLowerCasePtr:   Int64Ptr(1),
		MinNumber:         1,
		MinNumberPtr:      Int64Ptr(1),
		MinSymbol:         1,
		MinSymbolPtr:      Int64Ptr(1),
		MinUpperCase:      1,
		MinUpperCasePtr:   Int64Ptr(1),
	}
}

func (a *PasswordPolicyPasswordSettingsComplexity) IsPolicyInstance() bool {
	return true
}

func (a *PasswordPolicyPasswordSettingsComplexity) MarshalJSON() ([]byte, error) {
	type Alias PasswordPolicyPasswordSettingsComplexity
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.MinLength != 0 {
		result.MinLengthPtr = Int64Ptr(a.MinLength)
	}
	if a.MinLowerCase != 0 {
		result.MinLowerCasePtr = Int64Ptr(a.MinLowerCase)
	}
	if a.MinNumber != 0 {
		result.MinNumberPtr = Int64Ptr(a.MinNumber)
	}
	if a.MinSymbol != 0 {
		result.MinSymbolPtr = Int64Ptr(a.MinSymbol)
	}
	if a.MinUpperCase != 0 {
		result.MinUpperCasePtr = Int64Ptr(a.MinUpperCase)
	}
	return json.Marshal(&result)
}

func (a *PasswordPolicyPasswordSettingsComplexity) UnmarshalJSON(data []byte) error {
	type Alias PasswordPolicyPasswordSettingsComplexity

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.MinLengthPtr != nil {
		a.MinLength = *result.MinLengthPtr
		a.MinLengthPtr = result.MinLengthPtr
	}
	if result.MinLowerCasePtr != nil {
		a.MinLowerCase = *result.MinLowerCasePtr
		a.MinLowerCasePtr = result.MinLowerCasePtr
	}
	if result.MinNumberPtr != nil {
		a.MinNumber = *result.MinNumberPtr
		a.MinNumberPtr = result.MinNumberPtr
	}
	if result.MinSymbolPtr != nil {
		a.MinSymbol = *result.MinSymbolPtr
		a.MinSymbolPtr = result.MinSymbolPtr
	}
	if result.MinUpperCasePtr != nil {
		a.MinUpperCase = *result.MinUpperCasePtr
		a.MinUpperCasePtr = result.MinUpperCasePtr
	}
	return nil
}
