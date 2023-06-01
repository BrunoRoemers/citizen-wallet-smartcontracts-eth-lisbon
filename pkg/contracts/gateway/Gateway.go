// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// GatewayMemoryUserOp is an auto generated low-level Go binding around an user-defined struct.
type GatewayMemoryUserOp struct {
	Sender               common.Address
	Nonce                *big.Int
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	Paymaster            common.Address
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}

// GatewayUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type GatewayUserOpInfo struct {
	MUserOp       GatewayMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

// IEntryPointUserOpsPerAggregator is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointUserOpsPerAggregator struct {
	UserOps    []UserOperation
	Aggregator common.Address
	Signature  []byte
}

// IStakeManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerDepositInfo struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"targetSuccess\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"targetResult\",\"type\":\"bytes\"}],\"name\":\"ExecutionResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureValidationFailed\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sigFailed\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"paymasterContext\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.ReturnInfo\",\"name\":\"returnInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"senderInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"factoryInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"paymasterInfo\",\"type\":\"tuple\"}],\"name\":\"ValidationResult\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sigFailed\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"paymasterContext\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.ReturnInfo\",\"name\":\"returnInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"senderInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"factoryInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"paymasterInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"stakeInfo\",\"type\":\"tuple\"}],\"internalType\":\"structIEntryPoint.AggregatorStakeInfo\",\"name\":\"aggregatorInfo\",\"type\":\"tuple\"}],\"name\":\"ValidationResultWithAggregation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureAggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SIG_VALIDATION_FAILED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"}],\"name\":\"_validateSenderAndPaymaster\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"deposit\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"deposit\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"getSenderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAggregator\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleAggregatedOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"}],\"internalType\":\"structGateway.MemoryUserOp\",\"name\":\"mUserOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contextOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"}],\"internalType\":\"structGateway.UserOpInfo\",\"name\":\"opInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"innerHandleOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"simulateHandleOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"simulateValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052604051620000129062000055565b604051809103906000f0801580156200002f573d6000803e3d6000fd5b506001600160a01b03166080523480156200004957600080fd5b50600160025562000063565b610213806200402183390190565b608051613f9b620000866000396000818161129c0152612df20152613f9b6000f3fe6080604052600436106101235760003560e01c80638f41ec5a116100a0578063bb9fe6bf11610064578063bb9fe6bf1461043a578063c23a5cea1461044f578063d6383f941461046f578063ee2194231461048f578063fc7e286d146104af57600080fd5b80638f41ec5a146103b2578063957122ab146103c75780639b249f69146103e7578063a619353114610407578063b760faf91461042757600080fd5b8063205c2878116100e7578063205c2878146101f657806335567e1a146102165780634b1d7cf5146102365780635287ce121461025657806370a082311461037357600080fd5b80630396cb60146101385780630bd28e3b1461014b5780631b2e01b81461016b5780631d732756146101b65780631fad948c146101d657600080fd5b366101335761013133610564565b005b600080fd5b6101316101463660046130c5565b6105cb565b34801561015757600080fd5b50610131610166366004613107565b61085a565b34801561017757600080fd5b506101a3610186366004613142565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b3480156101c257600080fd5b506101a36101d136600461333c565b610891565b3480156101e257600080fd5b506101316101f1366004613445565b610a04565b34801561020257600080fd5b5061013161021136600461349b565b610b80565b34801561022257600080fd5b506101a3610231366004613142565b610cfc565b34801561024257600080fd5b50610131610251366004613445565b610d42565b34801561026257600080fd5b5061031b6102713660046134c7565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506001600160a01b031660009081526020818152604091829020825160a08101845281546001600160701b038082168352600160701b820460ff16151594830194909452600160781b90049092169282019290925260019091015463ffffffff81166060830152640100000000900465ffffffffffff16608082015290565b6040805182516001600160701b03908116825260208085015115159083015283830151169181019190915260608083015163ffffffff169082015260809182015165ffffffffffff169181019190915260a0016101ad565b34801561037f57600080fd5b506101a361038e3660046134c7565b6001600160a01b03166000908152602081905260409020546001600160701b031690565b3480156103be57600080fd5b506101a3600181565b3480156103d357600080fd5b506101316103e23660046134e4565b611185565b3480156103f357600080fd5b50610131610402366004613568565b611282565b34801561041357600080fd5b506101a36104223660046135c2565b61133d565b6101316104353660046134c7565b610564565b34801561044657600080fd5b5061013161137f565b34801561045b57600080fd5b5061013161046a3660046134c7565b6114a8565b34801561047b57600080fd5b5061013161048a3660046135f6565b6116e0565b34801561049b57600080fd5b506101316104aa3660046135c2565b6117db565b3480156104bb57600080fd5b5061051e6104ca3660046134c7565b600060208190529081526040902080546001909101546001600160701b0380831692600160701b810460ff1692600160781b9091049091169063ffffffff811690640100000000900465ffffffffffff1685565b604080516001600160701b0396871681529415156020860152929094169183019190915263ffffffff16606082015265ffffffffffff909116608082015260a0016101ad565b61056e81346119b7565b6001600160a01b03811660008181526020818152604091829020805492516001600160701b03909316835292917f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c491015b60405180910390a25050565b33600090815260208190526040902063ffffffff82166106325760405162461bcd60e51b815260206004820152601a60248201527f6d757374207370656369667920756e7374616b652064656c617900000000000060448201526064015b60405180910390fd5b600181015463ffffffff90811690831610156106905760405162461bcd60e51b815260206004820152601c60248201527f63616e6e6f7420646563726561736520756e7374616b652074696d65000000006044820152606401610629565b80546000906106b0903490600160781b90046001600160701b031661366d565b9050600081116106f75760405162461bcd60e51b81526020600482015260126024820152711b9bc81cdd185ad9481cdc1958da599a595960721b6044820152606401610629565b6001600160701b0381111561073f5760405162461bcd60e51b815260206004820152600e60248201526d7374616b65206f766572666c6f7760901b6044820152606401610629565b6040805160a08101825283546001600160701b0390811682526001602080840182815286841685870190815263ffffffff808b16606088019081526000608089018181523380835296829052908a902098518954955194518916600160781b02600160781b600160e81b0319951515600160701b026effffffffffffffffffffffffffffff199097169190991617949094179290921695909517865551949092018054925165ffffffffffff166401000000000269ffffffffffffffffffff19909316949093169390931717905590517fa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c019061084d908490879091825263ffffffff16602082015260400190565b60405180910390a2505050565b3360009081526001602090815260408083206001600160c01b0385168452909152812080549161088983613680565b919050555050565b6000805a90503330146108e65760405162461bcd60e51b815260206004820152601760248201527f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000006044820152606401610629565b8451604081015160608201518101611388015a10156109105763deaddead60e01b60005260206000fd5b8751600090156109a457600061092d846000015160008c86611a53565b9050806109a2576000610941610800611a6b565b80519091501561099c5784600001516001600160a01b03168a602001517f1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a2018760200151846040516109939291906136e9565b60405180910390a35b60019250505b505b600088608001515a86030190506109f66000838b8b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250889250611a97915050565b9a9950505050505050505050565b610a0c611d8e565b816000816001600160401b03811115610a2757610a27613177565b604051908082528060200260200182016040528015610a6057816020015b610a4d61303b565b815260200190600190039081610a455790505b50905060005b82811015610ad9576000828281518110610a8257610a82613702565b60200260200101519050600080610abd848a8a87818110610aa557610aa5613702565b9050602002810190610ab79190613718565b85611de5565b91509150610ace8483836000611fd0565b505050600101610a66565b506040516000907fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972908290a160005b83811015610b6357610b5781888884818110610b2657610b26613702565b9050602002810190610b389190613718565b858481518110610b4a57610b4a613702565b602002602001015161216c565b90910190600101610b08565b50610b6e8482612293565b505050610b7b6001600255565b505050565b33600090815260208190526040902080546001600160701b0316821115610be95760405162461bcd60e51b815260206004820152601960248201527f576974686472617720616d6f756e7420746f6f206c61726765000000000000006044820152606401610629565b8054610bff9083906001600160701b0316613739565b81546001600160701b0319166001600160701b0391909116178155604080516001600160a01b03851681526020810184905233917fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb910160405180910390a26000836001600160a01b03168360405160006040518083038185875af1925050503d8060008114610cab576040519150601f19603f3d011682016040523d82523d6000602084013e610cb0565b606091505b5050905080610cf65760405162461bcd60e51b81526020600482015260126024820152716661696c656420746f20776974686472617760701b6044820152606401610629565b50505050565b6001600160a01b03821660009081526001602090815260408083206001600160c01b038516845290915290819020549082901b67ffffffffffffffff1916175b92915050565b610d4a611d8e565b816000805b82811015610ebe5736868683818110610d6a57610d6a613702565b9050602002810190610d7c919061374c565b9050366000610d8b8380613762565b90925090506000610da260408501602086016134c7565b90506000196001600160a01b03821601610dfe5760405162461bcd60e51b815260206004820152601760248201527f4141393620696e76616c69642061676772656761746f720000000000000000006044820152606401610629565b6001600160a01b03811615610e9b576001600160a01b03811663e3563a4f8484610e2b60408901896137ab565b6040518563ffffffff1660e01b8152600401610e4a9493929190613956565b60006040518083038186803b158015610e6257600080fd5b505afa925050508015610e73575060015b610e9b5760405163086a9f7560e41b81526001600160a01b0382166004820152602401610629565b610ea5828761366d565b9550505050508080610eb690613680565b915050610d4f565b506000816001600160401b03811115610ed957610ed9613177565b604051908082528060200260200182016040528015610f1257816020015b610eff61303b565b815260200190600190039081610ef75790505b506040519091507fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97290600090a16000805b848110156110275736888883818110610f5e57610f5e613702565b9050602002810190610f70919061374c565b9050366000610f7f8380613762565b90925090506000610f9660408501602086016134c7565b90508160005b8181101561100e576000898981518110610fb857610fb8613702565b60200260200101519050600080610fdb8b898987818110610aa557610aa5613702565b91509150610feb84838389611fd0565b8a610ff581613680565b9b5050505050808061100690613680565b915050610f9c565b505050505050808061101f90613680565b915050610f43565b50600080915060005b85811015611140573689898381811061104b5761104b613702565b905060200281019061105d919061374c565b905061106f60408201602083016134c7565b6001600160a01b03167f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d60405160405180910390a23660006110b18380613762565b90925090508060005b81811015611128576110fc888585848181106110d8576110d8613702565b90506020028101906110ea9190613718565b8b8b81518110610b4a57610b4a613702565b611106908861366d565b96508761111281613680565b985050808061112090613680565b9150506110ba565b5050505050808061113890613680565b915050611030565b506040516000907f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d908290a26111768682612293565b5050505050610b7b6001600255565b8315801561119b57506001600160a01b0383163b155b156111e85760405162461bcd60e51b815260206004820152601960248201527f41413230206163636f756e74206e6f74206465706c6f796564000000000000006044820152606401610629565b601481106112605760006111ff60148284866139d3565b611208916139fd565b60601c9050803b60000361125e5760405162461bcd60e51b815260206004820152601b60248201527f41413330207061796d6173746572206e6f74206465706c6f79656400000000006044820152606401610629565b505b60405162461bcd60e51b81526020600482015260006024820152604401610629565b604051632b870d1b60e11b81526000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063570e1a36906112d39086908690600401613a32565b6020604051808303816000875af11580156112f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113169190613a46565b604051633653dc0360e11b81526001600160a01b0382166004820152909150602401610629565b60006113488261238c565b6040805160208101929092523090820152466060820152608001604051602081830303815290604052805190602001209050919050565b3360009081526020819052604081206001810154909163ffffffff90911690036113d85760405162461bcd60e51b815260206004820152600a6024820152691b9bdd081cdd185ad95960b21b6044820152606401610629565b8054600160701b900460ff166114245760405162461bcd60e51b8152602060048201526011602482015270616c726561647920756e7374616b696e6760781b6044820152606401610629565b600181015460009061143c9063ffffffff1642613a63565b60018301805469ffffffffffff00000000191664010000000065ffffffffffff841690810291909117909155835460ff60701b1916845560405190815290915033907ffa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a906020016105bf565b3360009081526020819052604090208054600160781b90046001600160701b03168061150d5760405162461bcd60e51b81526020600482015260146024820152734e6f207374616b6520746f20776974686472617760601b6044820152606401610629565b6001820154640100000000900465ffffffffffff1661156e5760405162461bcd60e51b815260206004820152601d60248201527f6d7573742063616c6c20756e6c6f636b5374616b6528292066697273740000006044820152606401610629565b60018201544264010000000090910465ffffffffffff1611156115d35760405162461bcd60e51b815260206004820152601b60248201527f5374616b65207769746864726177616c206973206e6f742064756500000000006044820152606401610629565b60018201805469ffffffffffffffffffff191690558154600160781b600160e81b0319168255604080516001600160a01b03851681526020810183905233917fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3910160405180910390a26000836001600160a01b03168260405160006040518083038185875af1925050503d806000811461168a576040519150601f19603f3d011682016040523d82523d6000602084013e61168f565b606091505b5050905080610cf65760405162461bcd60e51b815260206004820152601860248201527f6661696c656420746f207769746864726177207374616b6500000000000000006044820152606401610629565b6116e861303b565b6116f1856123a5565b60008061170060008885611de5565b915091506000611710838361247f565b905061171b43600052565b600061172960008a8761216c565b905061173443600052565b600060606001600160a01b038a16156117aa57896001600160a01b03168989604051611761929190613a89565b6000604051808303816000865af19150503d806000811461179e576040519150601f19603f3d011682016040523d82523d6000602084013e6117a3565b606091505b5090925090505b866080015183856020015186604001518585604051630116f59360e71b815260040161062996959493929190613a99565b6117e361303b565b6117ec826123a5565b6000806117fb60008585611de5565b915091506000611812846000015160a0015161254b565b8451519091506000906118249061254b565b9050611843604051806040016040528060008152602001600081525090565b36600061185360408a018a6137ab565b90925090506000601482101561186a576000611885565b6118786014600084866139d3565b611881916139fd565b60601c5b90506118908161254b565b935050505060006118a1868661247f565b9050600081600001519050600060016001600160a01b0316826001600160a01b031614905060006040518060c001604052808b6080015181526020018b6040015181526020018315158152602001856020015165ffffffffffff168152602001856040015165ffffffffffff16815260200161191e8c6060015190565b905290506001600160a01b0383161580159061194457506001600160a01b038316600114155b156119965760006040518060400160405280856001600160a01b0316815260200161196e8661254b565b81525090508187878a84604051633ebb2d3960e21b8152600401610629959493929190613b3b565b8086868960405163e0cff05f60e01b81526004016106299493929190613bbb565b6001600160a01b038216600090815260208190526040812080549091906119e89084906001600160701b031661366d565b90506001600160701b03811115611a345760405162461bcd60e51b815260206004820152601060248201526f6465706f736974206f766572666c6f7760801b6044820152606401610629565b81546001600160701b0319166001600160701b03919091161790555050565b6000806000845160208601878987f195945050505050565b60603d82811115611a795750815b604051602082018101604052818152816000602083013e9392505050565b6000805a855190915060009081611aad8261259a565b60a08301519091506001600160a01b038116611acc5782519350611c73565b809350600088511115611c7357868202955060028a6002811115611af257611af2613c12565b14611b6457606083015160405163a9a2340960e01b81526001600160a01b0383169163a9a2340991611b2c908e908d908c90600401613c28565b600060405180830381600088803b158015611b4657600080fd5b5087f1158015611b5a573d6000803e3d6000fd5b5050505050611c73565b606083015160405163a9a2340960e01b81526001600160a01b0383169163a9a2340991611b99908e908d908c90600401613c28565b600060405180830381600088803b158015611bb357600080fd5b5087f193505050508015611bc5575060015b611c7357611bd1613c6f565b806308c379a003611c2a5750611be5613c8b565b80611bf05750611c2c565b8b81604051602001611c029190613d14565b60408051601f1981840301815290829052631101335b60e11b825261062992916004016136e9565b505b8a604051631101335b60e11b81526004016106299181526040602082018190526012908201527110504d4c081c1bdcdd13dc081c995d995c9d60721b606082015260800190565b5a85038701965081870295508589604001511015611cdc578a604051631101335b60e11b815260040161062991815260406020808301829052908201527f414135312070726566756e642062656c6f772061637475616c476173436f7374606082015260800190565b6040890151869003611cee85826119b7565b6000808c6002811115611d0357611d03613c12565b1490508460a001516001600160a01b031685600001516001600160a01b03168c602001517f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f8860200151858d8f604051611d76949392919093845291151560208401526040830152606082015260800190565b60405180910390a45050505050505095945050505050565b6002805403611ddf5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610629565b60028055565b60008060005a8451909150611dfa86826125ca565b611e038661133d565b6020860152604081015160608201516080830151171760e087013517610100870135176effffffffffffffffffffffffffffff811115611e855760405162461bcd60e51b815260206004820152601860248201527f41413934206761732076616c756573206f766572666c6f7700000000000000006044820152606401610629565b600080611e91846126c3565b9050611e9f8a8a8a84612710565b85516020870151919950919350611eb69190612948565b611f0c5789604051631101335b60e11b8152600401610629918152604060208201819052601a908201527f4141323520696e76616c6964206163636f756e74206e6f6e6365000000000000606082015260800190565b611f1543600052565b60a08401516060906001600160a01b031615611f3d57611f388b8b8b8587612995565b975090505b60005a87039050808b60a001351015611fa2578b604051631101335b60e11b8152600401610629918152604060208201819052601e908201527f41413430206f76657220766572696669636174696f6e4761734c696d69740000606082015260800190565b60408a018390528160608b015260c08b01355a8803018a608001818152505050505050505050935093915050565b600080611fdc85612bb8565b91509150816001600160a01b0316836001600160a01b0316146120425785604051631101335b60e11b81526004016106299181526040602082018190526014908201527320a0991a1039b4b3b730ba3ab9329032b93937b960611b606082015260800190565b801561209a5785604051631101335b60e11b81526004016106299181526040602082018190526017908201527f414132322065787069726564206f72206e6f7420647565000000000000000000606082015260800190565b60006120a585612bb8565b925090506001600160a01b038116156121015786604051631101335b60e11b81526004016106299181526040602082018190526014908201527320a0999a1039b4b3b730ba3ab9329032b93937b960611b606082015260800190565b81156121635786604051631101335b60e11b81526004016106299181526040602082018190526021908201527f41413332207061796d61737465722065787069726564206f72206e6f742064756060820152606560f81b608082015260a00190565b50505050505050565b6000805a9050600061217f846060015190565b905030631d73275661219460608801886137ab565b87856040518563ffffffff1660e01b81526004016121b59493929190613d52565b6020604051808303816000875af19250505080156121f0575060408051601f3d908101601f191682019092526121ed91810190613e05565b60015b61228757600060206000803e50600051632152215360e01b81016122525786604051631101335b60e11b8152600401610629918152604060208201819052600f908201526e41413935206f7574206f662067617360881b606082015260800190565b600085608001515a6122649086613739565b61226e919061366d565b905061227e886002888685611a97565b9450505061228a565b92505b50509392505050565b6001600160a01b0382166122e95760405162461bcd60e51b815260206004820152601860248201527f4141393020696e76616c69642062656e656669636961727900000000000000006044820152606401610629565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114612336576040519150601f19603f3d011682016040523d82523d6000602084013e61233b565b606091505b5050905080610b7b5760405162461bcd60e51b815260206004820152601f60248201527f41413931206661696c65642073656e6420746f2062656e6566696369617279006044820152606401610629565b600061239782612c0b565b805190602001209050919050565b3063957122ab6123b860408401846137ab565b6123c560208601866134c7565b6123d36101208701876137ab565b6040518663ffffffff1660e01b81526004016123f3959493929190613e1e565b60006040518083038186803b15801561240b57600080fd5b505afa92505050801561241c575060015b61247c57612428613c6f565b806308c379a003612470575061243c613c8b565b806124475750612472565b80511561246c57600081604051631101335b60e11b81526004016106299291906136e9565b5050565b505b3d6000803e3d6000fd5b50565b60408051606081018252600080825260208201819052918101829052906124a584612cde565b905060006124b284612cde565b82519091506001600160a01b0381166124c9575080515b602080840151604080860151928501519085015191929165ffffffffffff80831690851610156124f7578193505b8065ffffffffffff168365ffffffffffff161115612513578092505b5050604080516060810182526001600160a01b03909416845265ffffffffffff92831660208501529116908201529250505092915050565b604080518082018252600080825260208083018281526001600160a01b03959095168252819052919091208054600160781b90046001600160701b031682526001015463ffffffff1690915290565b60c081015160e0820151600091908082036125b6575092915050565b6125c282488301612d4f565b949350505050565b6125d760208301836134c7565b6001600160a01b0316815260208083013590820152608080830135604083015260a0830135606083015260c0808401359183019190915260e08084013591830191909152610100830135908201523660006126366101208501856137ab565b909250905080156126b65760148110156126925760405162461bcd60e51b815260206004820152601d60248201527f4141393320696e76616c6964207061796d6173746572416e64446174610000006044820152606401610629565b6126a06014600083856139d3565b6126a9916139fd565b60601c60a0840152610cf6565b600060a084015250505050565b60a081015160009081906001600160a01b03166126e15760016126e4565b60035b60ff16905060008360800151828560600151028560400151010190508360c00151810292505050919050565b60008060005a8551805191925090612735898861273060408c018c6137ab565b612d67565b60a082015161274343600052565b60006001600160a01b03821661278b576001600160a01b0383166000908152602081905260409020546001600160701b031688811161278457808903612787565b60005b9150505b606084015160208a0151604051633a871cdd60e01b81526001600160a01b03861692633a871cdd9290916127c5918f918790600401613e54565b60206040518083038160008887f193505050508015612801575060408051601f3d908101601f191682019092526127fe91810190613e05565b60015b61288b5761280d613c6f565b806308c379a00361283e5750612821613c8b565b8061282c5750612840565b8b81604051602001611c029190613e79565b505b8a604051631101335b60e11b8152600401610629918152604060208201819052601690820152754141323320726576657274656420286f72204f4f472960501b606082015260800190565b95506001600160a01b038216612935576001600160a01b038316600090815260208190526040902080546001600160701b0316808a1115612918578c604051631101335b60e11b81526004016106299181526040602082018190526017908201527f41413231206469646e2774207061792070726566756e64000000000000000000606082015260800190565b81546001600160701b031916908a90036001600160701b03161790555b5a85039650505050505094509492505050565b6001600160a01b038216600090815260016020908152604080832084821c80855292528220805484916001600160401b03831691908561298783613680565b909155501495945050505050565b825160608181015190916000918481116129f15760405162461bcd60e51b815260206004820152601f60248201527f4141343120746f6f206c6974746c6520766572696669636174696f6e476173006044820152606401610629565b60a08201516001600160a01b038116600090815260208190526040902080548784039291906001600160701b031689811015612a79578c604051631101335b60e11b8152600401610629918152604060208201819052601e908201527f41413331207061796d6173746572206465706f73697420746f6f206c6f770000606082015260800190565b8981038260000160006101000a8154816001600160701b0302191690836001600160701b03160217905550826001600160a01b031663f465c77e858e8e602001518e6040518563ffffffff1660e01b8152600401612ad993929190613e54565b60006040518083038160008887f193505050508015612b1a57506040513d6000823e601f3d908101601f19168201604052612b179190810190613eb0565b60015b612ba457612b26613c6f565b806308c379a003612b575750612b3a613c8b565b80612b455750612b59565b8d81604051602001611c029190613f3b565b505b8c604051631101335b60e11b8152600401610629918152604060208201819052601690820152754141333320726576657274656420286f72204f4f472960501b606082015260800190565b909e909d509b505050505050505050505050565b60008082600003612bce57506000928392509050565b6000612bd984612cde565b9050806040015165ffffffffffff16421180612c005750806020015165ffffffffffff1642105b905194909350915050565b6060813560208301356000612c2b612c2660408701876137ab565b613028565b90506000612c3f612c2660608801886137ab565b9050608086013560a087013560c088013560e08901356101008a01356000612c6e612c266101208e018e6137ab565b604080516001600160a01b039c909c1660208d01528b81019a909a5260608b019890985250608089019590955260a088019390935260c087019190915260e08601526101008501526101208401526101408084019190915281518084039091018152610160909201905292915050565b60408051606081018252600080825260208201819052918101919091528160a081901c65ffffffffffff8116600003612d1a575065ffffffffffff5b604080516060810182526001600160a01b03909316835260d09490941c602083015265ffffffffffff16928101929092525090565b6000818310612d5e5781612d60565b825b9392505050565b8015610cf6578251516001600160a01b0381163b15612dd25784604051631101335b60e11b8152600401610629918152604060208201819052601f908201527f414131302073656e64657220616c726561647920636f6e737472756374656400606082015260800190565b835160600151604051632b870d1b60e11b81526000916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163570e1a369190612e2a9088908890600401613a32565b60206040518083038160008887f1158015612e49573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190612e6e9190613a46565b90506001600160a01b038116612ed05785604051631101335b60e11b8152600401610629918152604060208201819052601b908201527f4141313320696e6974436f6465206661696c6564206f72204f4f470000000000606082015260800190565b816001600160a01b0316816001600160a01b031614612f3a5785604051631101335b60e11b815260040161062991815260406020808301829052908201527f4141313420696e6974436f6465206d7573742072657475726e2073656e646572606082015260800190565b806001600160a01b03163b600003612f9d5785604051631101335b60e11b815260040161062991815260406020808301829052908201527f4141313520696e6974436f6465206d757374206372656174652073656e646572606082015260800190565b6000612fac60148286886139d3565b612fb5916139fd565b60601c9050826001600160a01b031686602001517fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d83896000015160a001516040516130179291906001600160a01b0392831681529116602082015260400190565b60405180910390a350505050505050565b6000604051828085833790209392505050565b6040518060a001604052806130a060405180610100016040528060006001600160a01b031681526020016000815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160008152602001600081525090565b8152602001600080191681526020016000815260200160008152602001600081525090565b6000602082840312156130d757600080fd5b813563ffffffff81168114612d6057600080fd5b80356001600160c01b038116811461310257600080fd5b919050565b60006020828403121561311957600080fd5b612d60826130eb565b6001600160a01b038116811461247c57600080fd5b803561310281613122565b6000806040838503121561315557600080fd5b823561316081613122565b915061316e602084016130eb565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b60a081018181106001600160401b03821117156131ac576131ac613177565b60405250565b61010081018181106001600160401b03821117156131ac576131ac613177565b601f8201601f191681016001600160401b03811182821017156131f7576131f7613177565b6040525050565b60006001600160401b0382111561321757613217613177565b50601f01601f191660200190565b600081830361018081121561323957600080fd5b6040516132458161318d565b8092506101008083121561325857600080fd5b6040519250613266836131b2565b61326f85613137565b8352602085013560208401526040850135604084015260608501356060840152608085013560808401526132a560a08601613137565b60a084015260c085013560c084015260e085013560e084015282825280850135602083015250610120840135604082015261014084013560608201526101608401356080820152505092915050565b60008083601f84011261330657600080fd5b5081356001600160401b0381111561331d57600080fd5b60208301915083602082850101111561333557600080fd5b9250929050565b6000806000806101c0858703121561335357600080fd5b84356001600160401b038082111561336a57600080fd5b818701915087601f83011261337e57600080fd5b8135613389816131fe565b60405161339682826131d2565b8281528a60208487010111156133ab57600080fd5b826020860160208301376000602084830101528098505050506133d18860208901613225565b94506101a08701359150808211156133e857600080fd5b506133f5878288016132f4565b95989497509550505050565b60008083601f84011261341357600080fd5b5081356001600160401b0381111561342a57600080fd5b6020830191508360208260051b850101111561333557600080fd5b60008060006040848603121561345a57600080fd5b83356001600160401b0381111561347057600080fd5b61347c86828701613401565b909450925050602084013561349081613122565b809150509250925092565b600080604083850312156134ae57600080fd5b82356134b981613122565b946020939093013593505050565b6000602082840312156134d957600080fd5b8135612d6081613122565b6000806000806000606086880312156134fc57600080fd5b85356001600160401b038082111561351357600080fd5b61351f89838a016132f4565b90975095506020880135915061353482613122565b9093506040870135908082111561354a57600080fd5b50613557888289016132f4565b969995985093965092949392505050565b6000806020838503121561357b57600080fd5b82356001600160401b0381111561359157600080fd5b61359d858286016132f4565b90969095509350505050565b600061016082840312156135bc57600080fd5b50919050565b6000602082840312156135d457600080fd5b81356001600160401b038111156135ea57600080fd5b6125c2848285016135a9565b6000806000806060858703121561360c57600080fd5b84356001600160401b038082111561362357600080fd5b61362f888389016135a9565b95506020870135915061364182613122565b909350604086013590808211156133e857600080fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610d3c57610d3c613657565b60006001820161369257613692613657565b5060010190565b60005b838110156136b457818101518382015260200161369c565b50506000910152565b600081518084526136d5816020860160208601613699565b601f01601f19169290920160200192915050565b8281526040602082015260006125c260408301846136bd565b634e487b7160e01b600052603260045260246000fd5b6000823561015e1983360301811261372f57600080fd5b9190910192915050565b81810381811115610d3c57610d3c613657565b60008235605e1983360301811261372f57600080fd5b6000808335601e1984360301811261377957600080fd5b8301803591506001600160401b0382111561379357600080fd5b6020019150600581901b360382131561333557600080fd5b6000808335601e198436030181126137c257600080fd5b8301803591506001600160401b038211156137dc57600080fd5b60200191503681900382131561333557600080fd5b6000808335601e1984360301811261380857600080fd5b83016020810192503590506001600160401b0381111561382757600080fd5b80360382131561333557600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600061016061387e8461387185613137565b6001600160a01b03169052565b6020830135602085015261389560408401846137f1565b8260408701526138a88387018284613836565b925050506138b960608401846137f1565b85830360608701526138cc838284613836565b925050506080830135608085015260a083013560a085015260c083013560c085015260e083013560e0850152610100808401358186015250610120613913818501856137f1565b86840383880152613925848284613836565b9350505050610140613939818501856137f1565b8684038388015261394b848284613836565b979650505050505050565b6040808252810184905260006060600586901b830181019083018783805b898110156139bc57868503605f190184528235368c900361015e1901811261399a578283fd5b6139a6868d830161385f565b9550506020938401939290920191600101613974565b50505050828103602084015261394b818587613836565b600080858511156139e357600080fd5b838611156139f057600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015613a2a5780818660140360031b1b83161692505b505092915050565b6020815260006125c2602083018486613836565b600060208284031215613a5857600080fd5b8151612d6081613122565b65ffffffffffff818116838216019080821115613a8257613a82613657565b5092915050565b8183823760009101908152919050565b868152856020820152600065ffffffffffff8087166040840152808616606084015250831515608083015260c060a0830152613ad860c08301846136bd565b98975050505050505050565b80518252602081015160208301526040810151151560408301526000606082015165ffffffffffff8082166060860152806080850151166080860152505060a082015160c060a08501526125c260c08501826136bd565b6000610140808352613b4f81840189613ae4565b915050613b69602083018780518252602090810151910152565b845160608301526020948501516080830152835160a08301529284015160c082015281516001600160a01b031660e0820152908301518051610100830152909201516101209092019190915292915050565b60e081526000613bce60e0830187613ae4565b9050613be7602083018680518252602090810151910152565b8351606083015260208401516080830152825160a0830152602083015160c083015295945050505050565b634e487b7160e01b600052602160045260246000fd5b600060038510613c4857634e487b7160e01b600052602160045260246000fd5b84825260606020830152613c5f60608301856136bd565b9050826040830152949350505050565b600060033d1115613c885760046000803e5060005160e01c5b90565b600060443d1015613c995790565b6040516003193d81016004833e81513d6001600160401b038160248401118184111715613cc857505050505090565b8285019150815181811115613ce05750505050505090565b843d8701016020828501011115613cfa5750505050505090565b613d09602082860101876131d2565b509095945050505050565b75020a09a98103837b9ba27b8103932bb32b93a32b21d160551b815260008251613d45816016850160208701613699565b9190910160160192915050565b60006101c0808352613d678184018789613836565b9050845160018060a01b03808251166020860152602082015160408601526040820151606086015260608201516080860152608082015160a08601528060a08301511660c08601525060c081015160e085015260e08101516101008501525060208501516101208401526040850151610140840152606085015161016084015260808501516101808401528281036101a084015261394b81856136bd565b600060208284031215613e1757600080fd5b5051919050565b606081526000613e32606083018789613836565b6001600160a01b03861660208401528281036040840152613ad8818587613836565b606081526000613e67606083018661385f565b60208301949094525060400152919050565b6e020a09919903932bb32b93a32b21d1608d1b815260008251613ea381600f850160208701613699565b91909101600f0192915050565b60008060408385031215613ec357600080fd5b82516001600160401b03811115613ed957600080fd5b8301601f81018513613eea57600080fd5b8051613ef5816131fe565b604051613f0282826131d2565b828152876020848601011115613f1757600080fd5b613f28836020830160208701613699565b6020969096015195979596505050505050565b6e020a09999903932bb32b93a32b21d1608d1b815260008251613ea381600f85016020870161369956fea264697066735822122093137699f4d80f59493e9f82b932a19d63400f29cdae716d1b4ec342978e71d164736f6c63430008140033608060405234801561001057600080fd5b506101f3806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063570e1a3614610030575b600080fd5b61004361003e3660046100ec565b61005f565b6040516001600160a01b03909116815260200160405180910390f35b60008061006f601482858761015e565b61007891610188565b60601c9050600061008c846014818861015e565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525084519495509360209350849250905082850182875af190506000519350806100e357600093505b50505092915050565b600080602083850312156100ff57600080fd5b823567ffffffffffffffff8082111561011757600080fd5b818501915085601f83011261012b57600080fd5b81358181111561013a57600080fd5b86602082850101111561014c57600080fd5b60209290920196919550909350505050565b6000808585111561016e57600080fd5b8386111561017b57600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff1981358181169160148510156101b55780818660140360031b1b83161692505b50509291505056fea2646970667358221220c2a91e75d8a003b4feb49f0ddaeb7b7807f4af0cb22723acc8fff844e85c9c5e64736f6c63430008140033",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayMetaData.Bin instead.
var GatewayBin = GatewayMetaData.Bin

// DeployGateway deploys a new Ethereum contract, binding an instance of Gateway to it.
func DeployGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gateway, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Gateway *GatewayCaller) SIGVALIDATIONFAILED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "SIG_VALIDATION_FAILED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Gateway *GatewaySession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _Gateway.Contract.SIGVALIDATIONFAILED(&_Gateway.CallOpts)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Gateway *GatewayCallerSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _Gateway.Contract.SIGVALIDATIONFAILED(&_Gateway.CallOpts)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Gateway *GatewayCaller) ValidateSenderAndPaymaster(opts *bind.CallOpts, initCode []byte, sender common.Address, paymasterAndData []byte) error {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "_validateSenderAndPaymaster", initCode, sender, paymasterAndData)

	if err != nil {
		return err
	}

	return err

}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Gateway *GatewaySession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _Gateway.Contract.ValidateSenderAndPaymaster(&_Gateway.CallOpts, initCode, sender, paymasterAndData)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Gateway *GatewayCallerSession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _Gateway.Contract.ValidateSenderAndPaymaster(&_Gateway.CallOpts, initCode, sender, paymasterAndData)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gateway *GatewayCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gateway *GatewaySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gateway.Contract.BalanceOf(&_Gateway.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gateway *GatewayCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gateway.Contract.BalanceOf(&_Gateway.CallOpts, account)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Gateway *GatewayCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Deposit         *big.Int
		Staked          bool
		Stake           *big.Int
		UnstakeDelaySec uint32
		WithdrawTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Staked = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeDelaySec = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.WithdrawTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Gateway *GatewaySession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _Gateway.Contract.Deposits(&_Gateway.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Gateway *GatewayCallerSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _Gateway.Contract.Deposits(&_Gateway.CallOpts, arg0)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Gateway *GatewayCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IStakeManagerDepositInfo, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Gateway *GatewaySession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _Gateway.Contract.GetDepositInfo(&_Gateway.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Gateway *GatewayCallerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _Gateway.Contract.GetDepositInfo(&_Gateway.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Gateway *GatewayCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Gateway *GatewaySession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _Gateway.Contract.GetNonce(&_Gateway.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Gateway *GatewayCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _Gateway.Contract.GetNonce(&_Gateway.CallOpts, sender, key)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Gateway *GatewayCaller) GetUserOpHash(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Gateway *GatewaySession) GetUserOpHash(userOp UserOperation) ([32]byte, error) {
	return _Gateway.Contract.GetUserOpHash(&_Gateway.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Gateway *GatewayCallerSession) GetUserOpHash(userOp UserOperation) ([32]byte, error) {
	return _Gateway.Contract.GetUserOpHash(&_Gateway.CallOpts, userOp)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Gateway *GatewayCaller) NonceSequenceNumber(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "nonceSequenceNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Gateway *GatewaySession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Gateway.Contract.NonceSequenceNumber(&_Gateway.CallOpts, arg0, arg1)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Gateway *GatewayCallerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Gateway.Contract.NonceSequenceNumber(&_Gateway.CallOpts, arg0, arg1)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Gateway *GatewayTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Gateway *GatewaySession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Gateway.Contract.AddStake(&_Gateway.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Gateway *GatewayTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Gateway.Contract.AddStake(&_Gateway.TransactOpts, unstakeDelaySec)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Gateway *GatewayTransactor) DepositTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "depositTo", account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Gateway *GatewaySession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.DepositTo(&_Gateway.TransactOpts, account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Gateway *GatewayTransactorSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.DepositTo(&_Gateway.TransactOpts, account)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Gateway *GatewayTransactor) GetSenderAddress(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "getSenderAddress", initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Gateway *GatewaySession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _Gateway.Contract.GetSenderAddress(&_Gateway.TransactOpts, initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Gateway *GatewayTransactorSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _Gateway.Contract.GetSenderAddress(&_Gateway.TransactOpts, initCode)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Gateway *GatewayTransactor) HandleAggregatedOps(opts *bind.TransactOpts, opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "handleAggregatedOps", opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Gateway *GatewaySession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.HandleAggregatedOps(&_Gateway.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Gateway *GatewayTransactorSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.HandleAggregatedOps(&_Gateway.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Gateway *GatewayTransactor) HandleOps(opts *bind.TransactOpts, ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Gateway *GatewaySession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.HandleOps(&_Gateway.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Gateway *GatewayTransactorSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.HandleOps(&_Gateway.TransactOpts, ops, beneficiary)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Gateway *GatewayTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Gateway *GatewaySession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.IncrementNonce(&_Gateway.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Gateway *GatewayTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.IncrementNonce(&_Gateway.TransactOpts, key)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Gateway *GatewayTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo GatewayUserOpInfo, context []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Gateway *GatewaySession) InnerHandleOp(callData []byte, opInfo GatewayUserOpInfo, context []byte) (*types.Transaction, error) {
	return _Gateway.Contract.InnerHandleOp(&_Gateway.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Gateway *GatewayTransactorSession) InnerHandleOp(callData []byte, opInfo GatewayUserOpInfo, context []byte) (*types.Transaction, error) {
	return _Gateway.Contract.InnerHandleOp(&_Gateway.TransactOpts, callData, opInfo, context)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Gateway *GatewayTransactor) SimulateHandleOp(opts *bind.TransactOpts, op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "simulateHandleOp", op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Gateway *GatewaySession) SimulateHandleOp(op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Gateway.Contract.SimulateHandleOp(&_Gateway.TransactOpts, op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Gateway *GatewayTransactorSession) SimulateHandleOp(op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Gateway.Contract.SimulateHandleOp(&_Gateway.TransactOpts, op, target, targetCallData)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_Gateway *GatewayTransactor) SimulateValidation(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "simulateValidation", userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_Gateway *GatewaySession) SimulateValidation(userOp UserOperation) (*types.Transaction, error) {
	return _Gateway.Contract.SimulateValidation(&_Gateway.TransactOpts, userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_Gateway *GatewayTransactorSession) SimulateValidation(userOp UserOperation) (*types.Transaction, error) {
	return _Gateway.Contract.SimulateValidation(&_Gateway.TransactOpts, userOp)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Gateway *GatewayTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Gateway *GatewaySession) UnlockStake() (*types.Transaction, error) {
	return _Gateway.Contract.UnlockStake(&_Gateway.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Gateway *GatewayTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _Gateway.Contract.UnlockStake(&_Gateway.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Gateway *GatewayTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Gateway *GatewaySession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawStake(&_Gateway.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Gateway *GatewayTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawStake(&_Gateway.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Gateway *GatewayTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Gateway *GatewaySession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawTo(&_Gateway.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Gateway *GatewayTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawTo(&_Gateway.TransactOpts, withdrawAddress, withdrawAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewaySession) Receive() (*types.Transaction, error) {
	return _Gateway.Contract.Receive(&_Gateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _Gateway.Contract.Receive(&_Gateway.TransactOpts)
}

// GatewayAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the Gateway contract.
type GatewayAccountDeployedIterator struct {
	Event *GatewayAccountDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayAccountDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayAccountDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayAccountDeployed represents a AccountDeployed event raised by the Gateway contract.
type GatewayAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_Gateway *GatewayFilterer) FilterAccountDeployed(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*GatewayAccountDeployedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayAccountDeployedIterator{contract: _Gateway.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_Gateway *GatewayFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *GatewayAccountDeployed, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayAccountDeployed)
				if err := _Gateway.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountDeployed is a log parse operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_Gateway *GatewayFilterer) ParseAccountDeployed(log types.Log) (*GatewayAccountDeployed, error) {
	event := new(GatewayAccountDeployed)
	if err := _Gateway.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayBeforeExecutionIterator is returned from FilterBeforeExecution and is used to iterate over the raw logs and unpacked data for BeforeExecution events raised by the Gateway contract.
type GatewayBeforeExecutionIterator struct {
	Event *GatewayBeforeExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayBeforeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayBeforeExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayBeforeExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayBeforeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayBeforeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayBeforeExecution represents a BeforeExecution event raised by the Gateway contract.
type GatewayBeforeExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeExecution is a free log retrieval operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_Gateway *GatewayFilterer) FilterBeforeExecution(opts *bind.FilterOpts) (*GatewayBeforeExecutionIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return &GatewayBeforeExecutionIterator{contract: _Gateway.contract, event: "BeforeExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeExecution is a free log subscription operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_Gateway *GatewayFilterer) WatchBeforeExecution(opts *bind.WatchOpts, sink chan<- *GatewayBeforeExecution) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayBeforeExecution)
				if err := _Gateway.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeforeExecution is a log parse operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_Gateway *GatewayFilterer) ParseBeforeExecution(log types.Log) (*GatewayBeforeExecution, error) {
	event := new(GatewayBeforeExecution)
	if err := _Gateway.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Gateway contract.
type GatewayDepositedIterator struct {
	Event *GatewayDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDeposited represents a Deposited event raised by the Gateway contract.
type GatewayDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Gateway *GatewayFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*GatewayDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayDepositedIterator{contract: _Gateway.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Gateway *GatewayFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDeposited)
				if err := _Gateway.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Gateway *GatewayFilterer) ParseDeposited(log types.Log) (*GatewayDeposited, error) {
	event := new(GatewayDeposited)
	if err := _Gateway.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewaySignatureAggregatorChangedIterator is returned from FilterSignatureAggregatorChanged and is used to iterate over the raw logs and unpacked data for SignatureAggregatorChanged events raised by the Gateway contract.
type GatewaySignatureAggregatorChangedIterator struct {
	Event *GatewaySignatureAggregatorChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewaySignatureAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewaySignatureAggregatorChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewaySignatureAggregatorChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewaySignatureAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewaySignatureAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewaySignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the Gateway contract.
type GatewaySignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignatureAggregatorChanged is a free log retrieval operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Gateway *GatewayFilterer) FilterSignatureAggregatorChanged(opts *bind.FilterOpts, aggregator []common.Address) (*GatewaySignatureAggregatorChangedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &GatewaySignatureAggregatorChangedIterator{contract: _Gateway.contract, event: "SignatureAggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchSignatureAggregatorChanged is a free log subscription operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Gateway *GatewayFilterer) WatchSignatureAggregatorChanged(opts *bind.WatchOpts, sink chan<- *GatewaySignatureAggregatorChanged, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewaySignatureAggregatorChanged)
				if err := _Gateway.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignatureAggregatorChanged is a log parse operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Gateway *GatewayFilterer) ParseSignatureAggregatorChanged(log types.Log) (*GatewaySignatureAggregatorChanged, error) {
	event := new(GatewaySignatureAggregatorChanged)
	if err := _Gateway.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the Gateway contract.
type GatewayStakeLockedIterator struct {
	Event *GatewayStakeLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStakeLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayStakeLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStakeLocked represents a StakeLocked event raised by the Gateway contract.
type GatewayStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Gateway *GatewayFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*GatewayStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayStakeLockedIterator{contract: _Gateway.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Gateway *GatewayFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *GatewayStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStakeLocked)
				if err := _Gateway.contract.UnpackLog(event, "StakeLocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Gateway *GatewayFilterer) ParseStakeLocked(log types.Log) (*GatewayStakeLocked, error) {
	event := new(GatewayStakeLocked)
	if err := _Gateway.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the Gateway contract.
type GatewayStakeUnlockedIterator struct {
	Event *GatewayStakeUnlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStakeUnlocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayStakeUnlocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStakeUnlocked represents a StakeUnlocked event raised by the Gateway contract.
type GatewayStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Gateway *GatewayFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*GatewayStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayStakeUnlockedIterator{contract: _Gateway.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Gateway *GatewayFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *GatewayStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStakeUnlocked)
				if err := _Gateway.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeUnlocked is a log parse operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Gateway *GatewayFilterer) ParseStakeUnlocked(log types.Log) (*GatewayStakeUnlocked, error) {
	event := new(GatewayStakeUnlocked)
	if err := _Gateway.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Gateway contract.
type GatewayStakeWithdrawnIterator struct {
	Event *GatewayStakeWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStakeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayStakeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStakeWithdrawn represents a StakeWithdrawn event raised by the Gateway contract.
type GatewayStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*GatewayStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayStakeWithdrawnIterator{contract: _Gateway.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStakeWithdrawn)
				if err := _Gateway.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) ParseStakeWithdrawn(log types.Log) (*GatewayStakeWithdrawn, error) {
	event := new(GatewayStakeWithdrawn)
	if err := _Gateway.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayUserOperationEventIterator is returned from FilterUserOperationEvent and is used to iterate over the raw logs and unpacked data for UserOperationEvent events raised by the Gateway contract.
type GatewayUserOperationEventIterator struct {
	Event *GatewayUserOperationEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayUserOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayUserOperationEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayUserOperationEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayUserOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayUserOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayUserOperationEvent represents a UserOperationEvent event raised by the Gateway contract.
type GatewayUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Gateway *GatewayFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (*GatewayUserOperationEventIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return &GatewayUserOperationEventIterator{contract: _Gateway.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Gateway *GatewayFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *GatewayUserOperationEvent, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayUserOperationEvent)
				if err := _Gateway.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserOperationEvent is a log parse operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Gateway *GatewayFilterer) ParseUserOperationEvent(log types.Log) (*GatewayUserOperationEvent, error) {
	event := new(GatewayUserOperationEvent)
	if err := _Gateway.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayUserOperationRevertReasonIterator is returned from FilterUserOperationRevertReason and is used to iterate over the raw logs and unpacked data for UserOperationRevertReason events raised by the Gateway contract.
type GatewayUserOperationRevertReasonIterator struct {
	Event *GatewayUserOperationRevertReason // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayUserOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayUserOperationRevertReason)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayUserOperationRevertReason)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayUserOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayUserOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayUserOperationRevertReason represents a UserOperationRevertReason event raised by the Gateway contract.
type GatewayUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Gateway *GatewayFilterer) FilterUserOperationRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*GatewayUserOperationRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayUserOperationRevertReasonIterator{contract: _Gateway.contract, event: "UserOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Gateway *GatewayFilterer) WatchUserOperationRevertReason(opts *bind.WatchOpts, sink chan<- *GatewayUserOperationRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayUserOperationRevertReason)
				if err := _Gateway.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserOperationRevertReason is a log parse operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Gateway *GatewayFilterer) ParseUserOperationRevertReason(log types.Log) (*GatewayUserOperationRevertReason, error) {
	event := new(GatewayUserOperationRevertReason)
	if err := _Gateway.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Gateway contract.
type GatewayWithdrawnIterator struct {
	Event *GatewayWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWithdrawn represents a Withdrawn event raised by the Gateway contract.
type GatewayWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*GatewayWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWithdrawnIterator{contract: _Gateway.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWithdrawn)
				if err := _Gateway.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Gateway *GatewayFilterer) ParseWithdrawn(log types.Log) (*GatewayWithdrawn, error) {
	event := new(GatewayWithdrawn)
	if err := _Gateway.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
