package orchestrator

import (
	"time"
)

type InstanceKey struct {
	Hostname string
	Port     int
}

type BinlogType int

const (
	BinaryLog BinlogType = iota
	RelayLog
)

// BinlogCoordinates described binary log coordinates in the form of log file & log position.
type BinlogCoordinates struct {
	LogFile string
	LogPos  int64
	Type    BinlogType
}

type NullInt64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}

type CandidatePromotionRule string

const (
	MustPromoteRule      CandidatePromotionRule = "must"
	PreferPromoteRule                           = "prefer"
	NeutralPromoteRule                          = "neutral"
	PreferNotPromoteRule                        = "prefer_not"
	MustNotPromoteRule                          = "must_not"
)

type Instance struct {
	Key                    InstanceKey
	InstanceAlias          string
	Uptime                 uint
	ServerID               uint
	ServerUUID             string
	Version                string
	VersionComment         string
	FlavorName             string
	ReadOnly               bool
	Binlog_format          string
	BinlogRowImage         string
	LogBinEnabled          bool
	LogSlaveUpdatesEnabled bool
	SelfBinlogCoordinates  BinlogCoordinates
	MasterKey              InstanceKey
	IsDetachedMaster       bool
	Slave_SQL_Running      bool
	Slave_IO_Running       bool
	HasReplicationFilters  bool
	GTIDMode               string
	SupportsOracleGTID     bool
	UsingOracleGTID        bool
	UsingMariaDBGTID       bool
	UsingPseudoGTID        bool
	ReadBinlogCoordinates  BinlogCoordinates
	ExecBinlogCoordinates  BinlogCoordinates
	IsDetached             bool
	RelaylogCoordinates    BinlogCoordinates
	LastSQLError           string
	LastIOError            string
	SecondsBehindMaster    NullInt64
	SQLDelay               uint
	ExecutedGtidSet        string
	GtidPurged             string

	SlaveLagSeconds                 NullInt64
	SlaveHosts                      InstanceKeyMap
	ClusterName                     string
	SuggestedClusterAlias           string
	DataCenter                      string
	PhysicalEnvironment             string
	ReplicationDepth                uint
	IsCoMaster                      bool
	HasReplicationCredentials       bool
	ReplicationCredentialsAvailable bool
	SemiSyncEnforced                bool
	SemiSyncMasterEnabled           bool
	SemiSyncReplicaEnabled          bool

	LastSeenTimestamp    string
	IsLastCheckValid     bool
	IsUpToDate           bool
	IsRecentlyChecked    bool
	SecondsSinceLastSeen NullInt64
	CountMySQLSnapshots  int

	// Careful. IsCandidate and PromotionRule are used together
	// and probably need to be merged. IsCandidate's value may
	// be picked up from daabase_candidate_instance's value when
	// reading an instance from the db.
	IsCandidate          bool
	PromotionRule        CandidatePromotionRule
	IsDowntimed          bool
	DowntimeReason       string
	DowntimeOwner        string
	DowntimeEndTimestamp string
	ElapsedDowntime      time.Duration
	UnresolvedHostname   string
	AllowTLS             bool

	LastDiscoveryLatency time.Duration
}