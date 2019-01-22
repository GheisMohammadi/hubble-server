package blockchain

import (
	"encoding/hex"
	"fmt"

	github_com_gallactic_gallactic_core_account "github.com/gallactic/gallactic/core/account"
	proto3 "github.com/gallactic/hubble_server/proto3"
	github_com_tendermint_tendermint_types "github.com/tendermint/tendermint/types"
)

func toAccount(acc *github_com_gallactic_gallactic_core_account.Account, retAccount *Account) {
	code := fmt.Sprintf("%s", acc.Code())
	//ID := uint64(i)

	retAccount.Address = acc.Address().String()
	retAccount.Balance = acc.Balance()
	retAccount.Permission = acc.Permissions().String()
	retAccount.Sequence = acc.Sequence()
	retAccount.Code = code
	//retAccount.ID = ID
}

//BlockMetaToBlock convert blockmeta struct to block struct
func BlockMetaToBlock(meta *github_com_tendermint_tendermint_types.BlockMeta, dest *Block) {
	header := meta.Header
	blockID := meta.BlockID

	dest.Hash = hex.EncodeToString(blockID.Hash)
	dest.ChainID = header.ChainID
	dest.Height = header.Height
	dest.Time = header.Time
	dest.TxCounts = header.NumTxs
	dest.LastBlockHash = hex.EncodeToString(header.LastBlockID.Hash)
}

//toBlock convert BlockResponse to Block struct
func toBlock(blockRes *proto3.BlockResponse, b *Block) {
	blockMeta := blockRes.BlockMeta
	blockID := blockMeta.BlockID
	header := blockRes.Block.Header
	//data := blockRes.Block.Data

	b.Hash = hex.EncodeToString(blockID.Hash)
	b.ChainID = header.ChainID
	b.Height = header.Height
	b.Time = header.Time
	b.LastBlockHash = hex.EncodeToString(header.LastBlockID.Hash)
	b.TxCounts = header.NumTxs
}

//toBlockMeta convert BlockMeta from tendermint to BlockMeta struct
func toBlockMeta(meta *github_com_tendermint_tendermint_types.BlockMeta, b *BlockMeta) {
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
	b.BlockHash = hex.EncodeToString(blockID.Hash)
	b.PartsSetTotal = blockID.PartsHeader.Total
	b.PartsSetHash = hex.EncodeToString(blockID.PartsHeader.Hash)
	// basic block info
	b.VersionBlock = header.Version.Block.Uint64()
	b.VersionApp = header.Version.App.Uint64()
	b.ChainID = header.ChainID
	b.Height = header.Height
	b.Time = header.Time
	b.NumTxs = header.NumTxs
	b.TotalTxs = header.TotalTxs
	// prev block info
	b.LastBlockHash = hex.EncodeToString(header.LastBlockID.Hash)
	// hashes of block data
	b.LastCommitHash = hex.EncodeToString(header.LastCommitHash)
	b.DataHash = hex.EncodeToString(header.DataHash)
	// hashes from the app output from the prev block
	b.ValidatorsHash = hex.EncodeToString(header.ValidatorsHash)
	b.NextValidatorsHash = hex.EncodeToString(header.NextValidatorsHash)
	b.ConsensusHash = hex.EncodeToString(header.ConsensusHash)
	b.AppHash = hex.EncodeToString(header.AppHash)
	b.LastResultsHash = hex.EncodeToString(header.LastResultsHash)
	// consensus info
	b.EvidenceHash = hex.EncodeToString(header.EvidenceHash)
	b.ProposerAddress = hex.EncodeToString(header.ProposerAddress)
}
