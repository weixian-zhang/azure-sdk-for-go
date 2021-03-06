package costmanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ConnectorBillingModel enumerates the values for connector billing model.
type ConnectorBillingModel string

const (
	// AutoUpgrade ...
	AutoUpgrade ConnectorBillingModel = "autoUpgrade"
	// Expired ...
	Expired ConnectorBillingModel = "expired"
	// Premium ...
	Premium ConnectorBillingModel = "premium"
	// Trial ...
	Trial ConnectorBillingModel = "trial"
)

// PossibleConnectorBillingModelValues returns an array of possible values for the ConnectorBillingModel const type.
func PossibleConnectorBillingModelValues() []ConnectorBillingModel {
	return []ConnectorBillingModel{AutoUpgrade, Expired, Premium, Trial}
}

// ConnectorStatus enumerates the values for connector status.
type ConnectorStatus string

const (
	// ConnectorStatusActive ...
	ConnectorStatusActive ConnectorStatus = "active"
	// ConnectorStatusError ...
	ConnectorStatusError ConnectorStatus = "error"
	// ConnectorStatusExpired ...
	ConnectorStatusExpired ConnectorStatus = "expired"
	// ConnectorStatusWarning ...
	ConnectorStatusWarning ConnectorStatus = "warning"
)

// PossibleConnectorStatusValues returns an array of possible values for the ConnectorStatus const type.
func PossibleConnectorStatusValues() []ConnectorStatus {
	return []ConnectorStatus{ConnectorStatusActive, ConnectorStatusError, ConnectorStatusExpired, ConnectorStatusWarning}
}

// CostAllocationPolicy enumerates the values for cost allocation policy.
type CostAllocationPolicy string

const (
	// Evenly ...
	Evenly CostAllocationPolicy = "Evenly"
	// Fixed ...
	Fixed CostAllocationPolicy = "Fixed"
	// Proportional ...
	Proportional CostAllocationPolicy = "Proportional"
)

// PossibleCostAllocationPolicyValues returns an array of possible values for the CostAllocationPolicy const type.
func PossibleCostAllocationPolicyValues() []CostAllocationPolicy {
	return []CostAllocationPolicy{Evenly, Fixed, Proportional}
}

// Direction enumerates the values for direction.
type Direction string

const (
	// Ascending ...
	Ascending Direction = "Ascending"
	// Descending ...
	Descending Direction = "Descending"
)

// PossibleDirectionValues returns an array of possible values for the Direction const type.
func PossibleDirectionValues() []Direction {
	return []Direction{Ascending, Descending}
}

// FunctionType enumerates the values for function type.
type FunctionType string

const (
	// AHUB ...
	AHUB FunctionType = "AHUB"
	// All ...
	All FunctionType = "All"
	// None ...
	None FunctionType = "None"
	// Reservations ...
	Reservations FunctionType = "Reservations"
)

// PossibleFunctionTypeValues returns an array of possible values for the FunctionType const type.
func PossibleFunctionTypeValues() []FunctionType {
	return []FunctionType{AHUB, All, None, Reservations}
}

// GranularityType enumerates the values for granularity type.
type GranularityType string

const (
	// Daily ...
	Daily GranularityType = "Daily"
	// Monthly ...
	Monthly GranularityType = "Monthly"
)

// PossibleGranularityTypeValues returns an array of possible values for the GranularityType const type.
func PossibleGranularityTypeValues() []GranularityType {
	return []GranularityType{Daily, Monthly}
}

// OperatorType enumerates the values for operator type.
type OperatorType string

const (
	// Contains ...
	Contains OperatorType = "Contains"
	// In ...
	In OperatorType = "In"
)

// PossibleOperatorTypeValues returns an array of possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{Contains, In}
}

// ReportConfigColumnType enumerates the values for report config column type.
type ReportConfigColumnType string

const (
	// ReportConfigColumnTypeDimension ...
	ReportConfigColumnTypeDimension ReportConfigColumnType = "Dimension"
	// ReportConfigColumnTypeTag ...
	ReportConfigColumnTypeTag ReportConfigColumnType = "Tag"
)

// PossibleReportConfigColumnTypeValues returns an array of possible values for the ReportConfigColumnType const type.
func PossibleReportConfigColumnTypeValues() []ReportConfigColumnType {
	return []ReportConfigColumnType{ReportConfigColumnTypeDimension, ReportConfigColumnTypeTag}
}

// RuleType enumerates the values for rule type.
type RuleType string

const (
	// RuleTypeCostAllocation ...
	RuleTypeCostAllocation RuleType = "CostAllocation"
	// RuleTypeCustomPrice ...
	RuleTypeCustomPrice RuleType = "CustomPrice"
	// RuleTypeShowbackRuleProperties ...
	RuleTypeShowbackRuleProperties RuleType = "ShowbackRuleProperties"
)

// PossibleRuleTypeValues returns an array of possible values for the RuleType const type.
func PossibleRuleTypeValues() []RuleType {
	return []RuleType{RuleTypeCostAllocation, RuleTypeCustomPrice, RuleTypeShowbackRuleProperties}
}

// ShowbackRuleStatus enumerates the values for showback rule status.
type ShowbackRuleStatus string

const (
	// Active ...
	Active ShowbackRuleStatus = "Active"
	// NotActive ...
	NotActive ShowbackRuleStatus = "NotActive"
)

// PossibleShowbackRuleStatusValues returns an array of possible values for the ShowbackRuleStatus const type.
func PossibleShowbackRuleStatusValues() []ShowbackRuleStatus {
	return []ShowbackRuleStatus{Active, NotActive}
}

// TimeframeType enumerates the values for timeframe type.
type TimeframeType string

const (
	// Custom ...
	Custom TimeframeType = "Custom"
	// MonthToDate ...
	MonthToDate TimeframeType = "MonthToDate"
	// WeekToDate ...
	WeekToDate TimeframeType = "WeekToDate"
	// YearToDate ...
	YearToDate TimeframeType = "YearToDate"
)

// PossibleTimeframeTypeValues returns an array of possible values for the TimeframeType const type.
func PossibleTimeframeTypeValues() []TimeframeType {
	return []TimeframeType{Custom, MonthToDate, WeekToDate, YearToDate}
}
