package main

type ScenarioCategory int

const (
	LifecycleSessionCategory ScenarioCategory = iota + 1
	APIValidationCategory
	LeakageStorageCategory
)

func (category ScenarioCategory) Label() string {
	switch category {
	case LifecycleSessionCategory:
		return "Lifecycle and session scenarios"
	case APIValidationCategory:
		return "API validation scenarios"
	case LeakageStorageCategory:
		return "Leakage and storage scenarios"
	default:
		return "Unknown scenario category"
	}
}
