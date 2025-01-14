package sdk

import "encoding/json"

type LifecycleExpirationPolicyRuleCondition struct {
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	Number          int64  `json:"-"`
	NumberPtr       *int64 `json:"number,omitempty"`
	Unit            string `json:"unit,omitempty"`
}

func NewLifecycleExpirationPolicyRuleCondition() *LifecycleExpirationPolicyRuleCondition {
	return &LifecycleExpirationPolicyRuleCondition{}
}

func (a *LifecycleExpirationPolicyRuleCondition) IsPolicyInstance() bool {
	return true
}

func (a *LifecycleExpirationPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	type Alias LifecycleExpirationPolicyRuleCondition
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.Number != 0 {
		result.NumberPtr = Int64Ptr(a.Number)
	}
	return json.Marshal(&result)
}

func (a *LifecycleExpirationPolicyRuleCondition) UnmarshalJSON(data []byte) error {
	type Alias LifecycleExpirationPolicyRuleCondition

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.NumberPtr != nil {
		a.Number = *result.NumberPtr
		a.NumberPtr = result.NumberPtr
	}
	return nil
}
