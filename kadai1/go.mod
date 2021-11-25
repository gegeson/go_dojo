module main

replace convert => ./cmd

go 1.17

require convert v0.0.0-00010101000000-000000000000

require (
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
)
