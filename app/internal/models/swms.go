package models

import (
	"encoding/json"
	"time"
)

type SWMSTableData struct {
	Name   string      `json:"name"`
	ID     int         `json:"id"`
	Values []SWMSValue `json:"values"`
}

type SWMSValue struct {
	SubId            int              `json:"subId"`
	ParentId         int              `json:"parentId"`
	Task             []string         `json:"task"`
	PotentialHazards []string         `json:"potentialHazards"`
	RiskBefore       string           `json:"riskBefore"`
	ControlMeasures  []ControlMeasure `json:"controlMeasures"`
	RiskAfter        string           `json:"riskAfter"`
}

type ControlMeasure struct {
	Name   string   `json:"name"`
	Values []string `json:"values,omitempty"`
}

type Swms struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	SwmsType        string          `json:"swms_type"`
	FileName        *string         `json:"file_name,omitempty"`
	FilePath        *string         `json:"file_path,omitempty"`
	UserId          int             `json:"user_id"`
	GeneratorStatus string          `json:"generator_status"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	SwmsData        json.RawMessage `json:"swms_data"` // JSONB type is best represented by json.RawMessage
}

type SwmsSchema struct {
	ID               int              `json:"id"`
	SubId            int              `json:"sub_id"`
	Name             string           `json:"name"`
	Task             []string         `json:"task"`
	PotentialHazards []string         `json:"potential_hazards"`
	RiskBefore       string           `json:"risk_before"`
	ControlMeasures  []ControlMeasure `json:"control_measures"`
	RiskAfter        string           `json:"risk_after"`
}
