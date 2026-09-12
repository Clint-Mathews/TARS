package model

import "time"

type RunState uint8

const (
	RunUnknown RunState = iota
	RunScanning
	RunComplete
	RunPartial
	RunCancelled
	RunFailed
)

type ScanConfig struct {
	Root                string
	Workers             int
	WorkQueue           int
	ReadBatch           int
	EventQueue          int
	EventBatch          int
	MaxRetries          int
	MaxPathBytes        int64
	MaxDepth            int
	MaxIssueSamples     int
	MaxDiagnosticsBytes int
	LargeFilesLimit     int
	ProgressInterval    time.Duration
}

func DefaultScanConfig(root string) ScanConfig {
	return ScanConfig{
		Root:                root,
		Workers:             8,
		WorkQueue:           68,
		ReadBatch:           128,
		EventQueue:          32,
		EventBatch:          128,
		MaxRetries:          200_000,
		MaxPathBytes:        64 << 20,
		MaxDepth:            256,
		MaxIssueSamples:     1_000,
		MaxDiagnosticsBytes: 512,
		LargeFilesLimit:     50,
		ProgressInterval:    100 * time.Millisecond,
	}
}

type ScanIssue struct {
	ScanID    ScanID
	EntryID   EntryID
	Path      string
	Operation string
	Kind      IssueKind
	Message   string
}

type ScanProgress struct {
	ObservedEntries int64
	ObserevedDirs   int64
	ObserverdPaths  int64
	UniqueFiles     int64
	CurrentPath     string
	RootTotal       Measurements
	Coverage        Coverage
	LargeFileIDs    []EntryID
}

type ScanSummary struct {
	ScanID          ScanID
	Root            string
	RootIdentity    Identity
	StartedAt       time.Duration
	FinishedAt      time.Duration
	State           RunState
	RootTotal       Measurements
	Coverage        Coverage
	ObservedEntries int64
	ObservedDirs    int64
	RegularPaths    int64
	UniqueFiles     int64
	IssueCounts     map[IssueKind]int64
	LimitsReached   []string
}

type VolumeInfo struct {
	MountPOint     string
	FilesystemType string
	MeasuredAt     time.Time
	Capacity       ByteValue
	Free           ByteValue
	Available      ByteValue
	Used           ByteValue
	Err            string
}
