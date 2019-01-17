package blockchain

//Account struct
type Account struct {
	Address    string
	PublicKey  string
	Balance    uint64
	Permission string
	Sequence   uint64
	Code       string
	ID         uint64
}

//BlockMeta struct in blocks
type BlockMeta struct {
	//block ID
	BlockHeaderHash string
	PartsSetTotal   int
	PartsSetHash    string
	// basic block info
	VersionBlock uint64
	VersionApp   uint64
	ChainID      string
	Height       int64
	Time         string
	NumTxs       int64
	TotalTxs     int64
	// prev block info
	LastBlockID string
	// hashes of block data
	LastCommitHash string
	DataHash       string
	// hashes from the app output from the prev block
	ValidatorsHash     string
	NextValidatorsHash string
	ConsensusHash      string
	AppHash            string
	LastResultsHash    string
	// consensus info
	EvidenceHash    string
	ProposerAddress string
}

//Block struct
type Block struct {
	ID            uint64
	Hash          string
	ChainID       string
	Height        int64
	Time          string
	LastBlockHash string
	TxCounts      int64
	DataHash      string
}

//Adapter for data base
type Adapter interface {
	CreateGRPCClient() error

	Update() error

	GetBlocksLastHeight() uint64
	GetBlocksMeta(id int) (*BlockMeta, error)
	GetBlock(height uint64) (*Block, error)
	GetAccountsCount() int
	GetAccount(id int) (*Account, error)
}
