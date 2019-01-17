package blockchain

import (
	"encoding/hex"

	github_com_tendermint_tendermint_types "github.com/tendermint/tendermint/types"
)

func toBlockMeta(meta *github_com_tendermint_tendermint_types.BlockMeta) *BlockMeta {
	/* Tendermint Block Meta Structure
	// basic block info
	Version  version.Consensus `json:"version"`
	ChainID  string            `json:"chain_id"`
	Height   int64             `json:"height"`
	Time     time.Time         `json:"time"`
	NumTxs   int64             `json:"num_txs"`
	TotalTxs int64             `json:"total_txs"`

	// prev block info
	LastBlockID BlockID `json:"last_block_id"`

	// hashes of block data
	LastCommitHash cmn.HexBytes `json:"last_commit_hash"` // commit from validators from the last block
	DataHash       cmn.HexBytes `json:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash     cmn.HexBytes `json:"validators_hash"`      // validators for the current block
	NextValidatorsHash cmn.HexBytes `json:"next_validators_hash"` // validators for the next block
	ConsensusHash      cmn.HexBytes `json:"consensus_hash"`       // consensus params for current block
	AppHash            cmn.HexBytes `json:"app_hash"`             // state after txs from the previous block
	LastResultsHash    cmn.HexBytes `json:"last_results_hash"`    // root hash of all results from the txs from the previous block

	// consensus info
	EvidenceHash    cmn.HexBytes `json:"evidence_hash"`    // evidence included in the block
	ProposerAddress Address      `json:"proposer_address"` // original proposer of the block
	*/

	header := meta.Header
	blockID := meta.BlockID

	// block ID
	BlockHeaderHash := hex.EncodeToString(blockID.Hash)
	PartsSetTotal := blockID.PartsHeader.Total
	PartsSetHash := hex.EncodeToString(blockID.PartsHeader.Hash)
	// basic block info
	VersionBlock := header.Version.Block.Uint64()
	VersionApp := header.Version.App.Uint64()
	ChainID := header.ChainID
	Height := header.Height
	Time := header.Time.String()
	NumTxs := header.NumTxs
	TotalTxs := header.TotalTxs
	// prev block info
	LastBlockID := hex.EncodeToString(header.LastBlockID.Hash)
	// hashes of block data
	LastCommitHash := hex.EncodeToString(header.LastCommitHash)
	DataHash := hex.EncodeToString(header.DataHash)
	// hashes from the app output from the prev block
	ValidatorsHash := hex.EncodeToString(header.ValidatorsHash)
	NextValidatorsHash := hex.EncodeToString(header.NextValidatorsHash)
	ConsensusHash := hex.EncodeToString(header.ConsensusHash)
	AppHash := hex.EncodeToString(header.AppHash)
	LastResultsHash := hex.EncodeToString(header.LastResultsHash)
	// consensus info
	EvidenceHash := hex.EncodeToString(header.EvidenceHash)
	ProposerAddress := hex.EncodeToString(header.ProposerAddress)

	var b *BlockMeta
	b = &BlockMeta{
		BlockHeaderHash:    BlockHeaderHash,
		PartsSetTotal:      PartsSetTotal,
		PartsSetHash:       PartsSetHash,
		VersionBlock:       VersionBlock,
		VersionApp:         VersionApp,
		ChainID:            ChainID,
		Height:             Height,
		Time:               Time,
		NumTxs:             NumTxs,
		TotalTxs:           TotalTxs,
		LastBlockID:        LastBlockID,
		LastCommitHash:     LastCommitHash,
		DataHash:           DataHash,
		ValidatorsHash:     ValidatorsHash,
		NextValidatorsHash: NextValidatorsHash,
		ConsensusHash:      ConsensusHash,
		AppHash:            AppHash,
		LastResultsHash:    LastResultsHash,
		EvidenceHash:       EvidenceHash,
		ProposerAddress:    ProposerAddress}

	return b
}
