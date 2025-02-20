package tester

import (
	"time"
)

// TBState represents the completion state of a `testing.TB`.
type TBState int

const (
	_ = iota
	// TBPassed represents a passed test.
	TBPassed
	// TBFailed represents a failed test.
	TBFailed
	// TBSkipped represents a skipped test.
	TBSkipped
)

func (s TBState) String() string {
	switch s {
	case TBPassed:
		return "passed"
	case TBFailed:
		return "failed"
	case TBSkipped:
		return "skipped"
	}
	return ""
}

// TBCommon is the representation of the common fields of a testing.TB.
type TBCommon struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	State      TBState   `json:"state"`
	Output     []byte    `json:"output"`
}

// Duration returns the run duration the Test.
func (c *TBCommon) Duration() time.Duration {
	return c.FinishedAt.Sub(c.StartedAt)
}

func (c *TBCommon) OutputString() string {
	return string(c.Output)
}

// Test is the representation of a `testing.T`.
type Test struct {
	TBCommon

	SubTests []*Test `json:"sub_tests,omitempty"`
}

// Benchmark is the representation of a `testing.B`.
type Benchmark struct {
	TBCommon

	SubBenchmarks []*Benchmark `json:"sub_benchmarks,omitempty"`
}

// Run is the representation of a pending test or benchmark that has not
// completed.
type Run struct {
	ID         string    `json:"id"`
	Package    Package   `json:"package"`
	Args       []string  `json:"args"`
	EnqueuedAt time.Time `json:"enqueued_at"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	Tests      []*Test   `json:"tests"`
}

// Package represents a go package that can be tested or benchmarked.
type Package struct {
	Name           string              `json:"name"`
	Path           string              `json:"path"`
	DefaultTimeout int                 `json:"default_timeout"`
	Options        map[string][]string `json:"options"`
}
