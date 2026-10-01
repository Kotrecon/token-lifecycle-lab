package main

type AppMode int

const (
	LifecycleLabMode AppMode = iota + 1
	ScenarioSimulatorMode
	ExitMode
)
