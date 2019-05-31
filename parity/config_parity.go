package parity

import (
	xchain ".."
	"github.com/ethereum/go-ethereum/common"
)

// Config is the data structure for Parity-Ethereum's chain configuration.
type Config struct {
	Name      string        `json:"name"`
	DataDir   string        `json:"dataDir"`
	EngineOpt ConfigEngines `json:"engine"`
	Params    *ConfigParams `json:"params"`
}

type ConfigEngines struct {
	ParityConfigEngineEthash         *ConfigEngineEthash         `json:"Ethash,omitempty"`
	ParityConfigEngineInstantSeal    *ConfigEngineInstantSeal    `json:"instantSeal,omitempty"`
	ParityConfigEngineClique         *ConfigEngineClique         `json:"Clique,omitempty"`
	ParityConfigEngineAuthorityRound *ConfigEngineAuthorityRound `json:"authorityRound,omitempty"`
}

// ParityConfigEngine is the data structure for a consensus engine.
type ConfigEngineEthash struct {
	Params ConfigEngineEthashParams `json:"params"`
}

// ParityConfigEngineParamsEthash is the data structure for the Ethash consensus engine parameters.
type ConfigEngineEthashParams struct {
	MinimumDifficulty                    xchain.Uint64 `json:"minimumDifficulty,omitempty"`
	DifficultyBoundDivisor               xchain.Uint64 `json:"difficultyBoundDivisor,omitempty"`
	DifficultyIncrementDivisor           xchain.Uint64 `json:"difficultyIncrementDivisor,omitempty"`
	MetropolisDifficultyIncrementDivisor xchain.Uint64 `json:"metropolisDifficultyIncrementDivisor,omitempty"`
	DurationLimit                        xchain.Uint64 `json:"durationLimit,omitempty"`

	HomesteadTransition           xchain.Uint64       `json:"homesteadTransition,omitempty"`
	BlockReward                   *xchain.BlockReward `json:"blockReward,omitempty"`
	BlockRewardContractTransition xchain.Uint64       `json:"blockRewardContractTransition,omitempty"`
	BlockRewardContractAddress    *common.Address     `json:"blockRewardContractAddress,omitempty"`
	BlockRewardContractCode       []byte              `json:"blockRewardContractCode,omitempty"`

	DaoHardforkTransition  xchain.Uint64    `json:"daoHardforkTransition,omitempty"`
	DaoHardforkBeneficiary *common.Address  `json:"daoHardforkBeneficiary,omitempty"`
	DaoHardforkAccounts    []common.Address `json:"daoHardforkAccounts,omitempty"`

	DifficultyHardforkTransition   xchain.Uint64 `json:"difficultyHardforkTransition,omitempty"`
	DifficultyHardforkBoundDivisor xchain.Uint64 `json:"difficultyHardforkBoundDivisor,omitempty"`
	BombDefuseTransition           xchain.Uint64 `json:"bombDefuseTransition,omitempty"`

	EIP100BTransition xchain.Uint64 `json:"eip100bTransition,omitempty"`

	Ecip1010PauseTransition    xchain.Uint64 `json:"ecip1010PauseTransition,omitempty"`
	Ecip1010ContinueTransition xchain.Uint64 `json:"ecip1010ContinueTransition,omitempty"`

	Ecip1017EraRounds xchain.Uint64 `json:"ecip1017EraRounds,omitempty"`

	DifficultyBombDelays *xchain.BTreeMap `json:"difficultyBombDelays,omitempty"`

	EXPIP2Transition    xchain.Uint64 `json:"expip2Transition,omitempty"`
	EXPIP2DurationLimit xchain.Uint64 `json:"expip2DurationLimit,omitempty"`
	ProgPowTransition   xchain.Uint64 `json:"progPowTransition,omitempty"`
}

type ConfigEngineInstantSeal struct {
	Params ConfigEngineInstantSealParams `json:"params"`
}

type ConfigEngineInstantSealParams struct {
	MillisecondTimestamp bool `json:"millisecondTimestamp,omitempty"`
}

type ConfigEngineClique struct {
	Params ConfigEngineCliqueParams `json:"params"`
}

type ConfigEngineCliqueParams struct {
	Period xchain.Uint64 `json:"period,omitempty"`
	Epoch  xchain.Uint64 `json:"epoch,omitempty"`
}

type ConfigEngineAuthorityRound struct {
	Params ConfigEngineAuthorityRoundParams `json:"params"`
}

type ConfigEngineAuthorityRoundParams struct {
}

type ConfigParams struct {
	AccountStartNonce    xchain.Uint64 `json:"accountStartNonce,omitempty"`
	MaximumExtraDataSize xchain.Uint64 `json:"maximumExtraDataSize,omitempty"`
	MinGasLimit          xchain.Uint64 `json:"minGasLimit,omitempty"`

	NetworkID xchain.Uint64 `json:"networkID,omitempty"`
	ChainID   xchain.Uint64 `json:"chainID,omitempty"`

	SubProtocolName string `json:"subprotocolName,omitempty"`

	ForkBlock xchain.Uint64 `json:"forkBlock,omitempty"`
	ForkHash  xchain.Uint64 `json:"forkHash,omitempty"`

	EIP150Transition    xchain.Uint64 `json:"eip150Transition,omitempty"`
	EIP160Transition    xchain.Uint64 `json:"eip160Transition,omitempty"`
	EIP161abcTransition xchain.Uint64 `json:"eip161abcTransition,omitempty"`
	EIP161dTransition   xchain.Uint64 `json:"eip161dTransition,omitempty"`

	EIP98Transition xchain.Uint64 `json:"eip98Transition,omitempty"`

	EIP155Transition                xchain.Uint64 `json:"eip155Transition,omitempty"`
	ValidateChainIDTransition       xchain.Uint64 `json:"validateChainIdTransition,omitempty"`
	ValidateChainReceiptsTransition xchain.Uint64 `json:"validateChainReceiptsTransition,omitempty"`

	EIP140Transition         xchain.Uint64   `json:"eip140Transition,omitempty"`
	EIP145Transition         xchain.Uint64   `json:"eip145Transition,omitempty"`
	EIP210Transition         xchain.Uint64   `json:"eip210Transition,omitempty"`
	EIP210ContractAddress    *common.Address `json:"eip210contractAddress,omitempty"`
	EIP210ContractCode       xchain.Uint64   `json:"eip210contractCode,omitempty"`
	EIP211Transition         xchain.Uint64   `json:"eip211Transition,omitempty"`
	EIP214Transition         xchain.Uint64   `json:"eip214Transition,omitempty"`
	EIP658Transition         xchain.Uint64   `json:"eip658Transition,omitempty"`
	EIP1014Transition        xchain.Uint64   `json:"eip1014Transition,omitempty"`
	EIP1052Transition        xchain.Uint64   `json:"eip1052Transition,omitempty"`
	EIP1283Transition        xchain.Uint64   `json:"eip1283Transition,omitempty"`
	EIP1283DisableTransition xchain.Uint64   `json:"eip1283DisableTransition,omitempty"`

	DustProtectionTransition xchain.Uint64 `json:"dustProtectionTransition,omitempty"`
	NonceCapIncrement        xchain.Uint64 `json:"nonceCapIncrement,omitempty"`
	RemoveDustContracts      bool          `json:"removeDustContracts,omitempty"`

	GasLimitBoundDivisor xchain.Uint64 `json:"gasLimitBoundDivisor,omitempty"`

	Registrar *common.Address `json:"registrar,omitempty"`

	ApplyReward bool `json:"applyReward,omitempty"`

	MaxCodeSize                             xchain.Uint64   `json:"maxCodeSize,omitempty"`
	MaxTransactionSize                      xchain.Uint64   `json:"maxTransactionSize,omitempty"`
	MaxCodeSizeTransition                   xchain.Uint64   `json:"maxCodeSizeTransition,omitempty"`
	TransactionPermissionContract           *common.Address `json:"transactionPermissionContract,omitempty"`
	TransactionPermissionContractTransition xchain.Uint64   `json:"transactionPermissionContractTransition,omitempty"`
	WASMActivationTransition                xchain.Uint64   `json:"wasmActivationTransition,omitempty"`
	KIP4Transition                          xchain.Uint64   `json:"kip4Transition,omitempty"`
	KIP6Transition                          xchain.Uint64   `json:"kip6Transition,omitempty"`
}
