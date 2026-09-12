package model

import "time"

type ScanID uint64
type EntryID uint64

type Identity struct {
	Device uint64
	Inode  uint64
	Known  bool
}

type ByteValue struct {
	Bytes int64
	Known bool
}

type Measurements struct {
	Logical   ByteValue
	Allocated ByteValue
}

type EntryKind uint8

const (
	EntryUnknown EntryKind = iota
	EntryRegular
	EnrtyDirectory
	EntrySymLink
	EntryOther
)

type CoverageState uint8

const (
	CoverageUnknown CoverageState = iota
	CoveragePending
	CoverageComplete
	CoveragePartial
	CoverageExcluded
)

type IssueKind uint8

const (
	IssueUnknown IssueKind = iota
	IssuePermission
	IssueDisappeared
	IssueIO
	IssueBoundary
	IssueRepeatedDirectory
	IssueChangedIdentity
	IssueChangedMetaData
	IssueCloudUnavailable
	IssueResourceLimit
	IssueCancelled
	IssueUnsupportedMetadata
	IssueOverflow
)

type CloudState uint8

const (
	CloudUnknown CloudState = iota
	CloudLocal
	CloudPlaceholder
)

type Coverage struct {
	State         CoverageState
	PendingDirs   int64
	FailedItems   int64
	ExcludedItems int64
	Reasons       []IssueKind
}
type Entry struct {
	ID               EntryID
	ParentID         EntryID
	Name             string
	Path             string
	Kind             EntryKind
	Identity         Identity
	ModTime          time.Time
	ModTimeKnown     bool
	Own              Measurements
	Total            Measurements
	RegularFileCount int64
	Coverage         Coverage
	Cloud            CloudState
	LinkText         string
	SharedWith       EntryID
}

func (e Entry) CloneEntry() Entry {
	ret := e // shallow copy struct value
	if e.Coverage.Reasons != nil {
		ret.Coverage.Reasons = make([]IssueKind, len(e.Coverage.Reasons))
		copy(ret.Coverage.Reasons, e.Coverage.Reasons) // Deep copy slice contents
	}
	return ret
}
