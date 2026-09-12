package model

import "time"

type EventHeader struct {
	ScanID   ScanID
	Sequence uint64
}

type ScanEvent interface {
	Header() EventHeader
	ScanEvent()
}

type ScanStarted struct {
	EventHeader
	Root         string
	RootIdentity Identity
	StartedAt    time.Time
	Config       ScanConfig
}

type EntriesUpdated struct {
	EventHeader
	Entries []Entry
}

type IssueRecorded struct {
	EventHeader
	Issues []ScanIssue
}

type ProgressUpdated struct {
	EventHeader
	Progress ScanProgress
}

type ScanFinished struct {
	EventHeader
	Summary ScanSummary
}

func (e ScanStarted) Header() EventHeader {
	return e.EventHeader
}

func (e EntriesUpdated) Header() EventHeader {
	return e.EventHeader
}

func (e IssueRecorded) Header() EventHeader {
	return e.EventHeader
}

func (e ProgressUpdated) Header() EventHeader {
	return e.EventHeader
}

func (e ScanFinished) Header() EventHeader {
	return e.EventHeader
}

func (ScanStarted) ScanEvent()     {}
func (EntriesUpdated) ScanEvent()  {}
func (IssueRecorded) ScanEvent()   {}
func (ProgressUpdated) ScanEvent() {}
func (ScanFinished) ScanEvent()    {}
