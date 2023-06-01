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
	Bin: "0x60a0604052604051620000129062000051565b604051809103905ff0801580156200002c573d5f803e3d5ffd5b506001600160a01b031660805234801562000045575f80fd5b5060016002556200005f565b6101ff8062003ec883390190565b608051613e496200007f5f395f81816112490152612d2a0152613e495ff3fe60806040526004361061011e575f3560e01c80638f41ec5a1161009d578063bb9fe6bf11610062578063bb9fe6bf14610423578063c23a5cea14610437578063d6383f9414610456578063ee21942314610475578063fc7e286d14610494575f80fd5b80638f41ec5a1461039f578063957122ab146103b35780639b249f69146103d2578063a6193531146103f1578063b760faf914610410575f80fd5b8063205c2878116100e3578063205c2878146101eb57806335567e1a1461020a5780634b1d7cf5146102295780635287ce121461024857806370a0823114610362575f80fd5b80630396cb60146101325780630bd28e3b146101455780631b2e01b8146101645780631d732756146101ad5780631fad948c146101cc575f80fd5b3661012e5761012c33610547565b005b5f80fd5b61012c610140366004612fea565b6105ad565b348015610150575f80fd5b5061012c61015f366004613028565b610838565b34801561016f575f80fd5b5061019a61017e366004613060565b600160209081525f928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b3480156101b8575f80fd5b5061019a6101c736600461324e565b61086e565b3480156101d7575f80fd5b5061012c6101e636600461334b565b6109d5565b3480156101f6575f80fd5b5061012c61020536600461339d565b610b4a565b348015610215575f80fd5b5061019a610224366004613060565b610cc1565b348015610234575f80fd5b5061012c61024336600461334b565b610d06565b348015610253575f80fd5b5061030a6102623660046133c7565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506001600160a01b03165f9081526020818152604091829020825160a08101845281546001600160701b038082168352600160701b820460ff16151594830194909452600160781b90049092169282019290925260019091015463ffffffff81166060830152640100000000900465ffffffffffff16608082015290565b6040805182516001600160701b03908116825260208085015115159083015283830151169181019190915260608083015163ffffffff169082015260809182015165ffffffffffff169181019190915260a0016101a4565b34801561036d575f80fd5b5061019a61037c3660046133c7565b6001600160a01b03165f908152602081905260409020546001600160701b031690565b3480156103aa575f80fd5b5061019a600181565b3480156103be575f80fd5b5061012c6103cd3660046133e2565b611136565b3480156103dd575f80fd5b5061012c6103ec366004613460565b611230565b3480156103fc575f80fd5b5061019a61040b3660046134b5565b6112e7565b61012c61041e3660046133c7565b610547565b34801561042e575f80fd5b5061012c611328565b348015610442575f80fd5b5061012c6104513660046133c7565b61144f565b348015610461575f80fd5b5061012c6104703660046134e6565b611682565b348015610480575f80fd5b5061012c61048f3660046134b5565b611771565b34801561049f575f80fd5b506105016104ae3660046133c7565b5f60208190529081526040902080546001909101546001600160701b0380831692600160701b810460ff1692600160781b9091049091169063ffffffff811690640100000000900465ffffffffffff1685565b604080516001600160701b0396871681529415156020860152929094169183019190915263ffffffff16606082015265ffffffffffff909116608082015260a0016101a4565b610551813461193c565b6001600160a01b0381165f8181526020818152604091829020805492516001600160701b03909316835292917f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c491015b60405180910390a25050565b335f90815260208190526040902063ffffffff82166106135760405162461bcd60e51b815260206004820152601a60248201527f6d757374207370656369667920756e7374616b652064656c617900000000000060448201526064015b60405180910390fd5b600181015463ffffffff90811690831610156106715760405162461bcd60e51b815260206004820152601c60248201527f63616e6e6f7420646563726561736520756e7374616b652074696d6500000000604482015260640161060a565b80545f90610690903490600160781b90046001600160701b0316613556565b90505f81116106d65760405162461bcd60e51b81526020600482015260126024820152711b9bc81cdd185ad9481cdc1958da599a595960721b604482015260640161060a565b6001600160701b0381111561071e5760405162461bcd60e51b815260206004820152600e60248201526d7374616b65206f766572666c6f7760901b604482015260640161060a565b6040805160a08101825283546001600160701b0390811682526001602080840182815286841685870190815263ffffffff808b16606088019081525f608089018181523380835296829052908a902098518954955194518916600160781b02600160781b600160e81b0319951515600160701b026effffffffffffffffffffffffffffff199097169190991617949094179290921695909517865551949092018054925165ffffffffffff166401000000000269ffffffffffffffffffff19909316949093169390931717905590517fa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c019061082b908490879091825263ffffffff16602082015260400190565b60405180910390a2505050565b335f9081526001602090815260408083206001600160c01b0385168452909152812080549161086683613569565b919050555050565b5f805a90503330146108c25760405162461bcd60e51b815260206004820152601760248201527f4141393220696e7465726e616c2063616c6c206f6e6c79000000000000000000604482015260640161060a565b8451604081015160608201518101611388015a10156108ea5763deaddead60e01b5f5260205ffd5b87515f9015610978575f610903845f01515f8c866119d7565b905080610976575f6109166108006119ed565b80519091501561097057845f01516001600160a01b03168a602001517f1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a2018760200151846040516109679291906135ce565b60405180910390a35b60019250505b505b5f88608001515a86030190506109c75f838b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250611a18915050565b9a9950505050505050505050565b6109dd611d02565b815f816001600160401b038111156109f7576109f7613093565b604051908082528060200260200182016040528015610a3057816020015b610a1d612f6c565b815260200190600190039081610a155790505b5090505f5b82811015610aa5575f828281518110610a5057610a506135e6565b602002602001015190505f80610a8a848a8a87818110610a7257610a726135e6565b9050602002810190610a8491906135fa565b85611d59565b91509150610a9a8483835f611f3f565b505050600101610a35565b506040515f907fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972908290a15f5b83811015610b2d57610b2181888884818110610af057610af06135e6565b9050602002810190610b0291906135fa565b858481518110610b1457610b146135e6565b60200260200101516120d9565b90910190600101610ad2565b50610b3884826121f9565b505050610b456001600255565b505050565b335f90815260208190526040902080546001600160701b0316821115610bb25760405162461bcd60e51b815260206004820152601960248201527f576974686472617720616d6f756e7420746f6f206c6172676500000000000000604482015260640161060a565b8054610bc89083906001600160701b0316613619565b81546001600160701b0319166001600160701b0391909116178155604080516001600160a01b03851681526020810184905233917fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb910160405180910390a25f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114610c70576040519150601f19603f3d011682016040523d82523d5f602084013e610c75565b606091505b5050905080610cbb5760405162461bcd60e51b81526020600482015260126024820152716661696c656420746f20776974686472617760701b604482015260640161060a565b50505050565b6001600160a01b0382165f9081526001602090815260408083206001600160c01b038516845290915290819020549082901b67ffffffffffffffff1916175b92915050565b610d0e611d02565b815f805b82811015610e7c5736868683818110610d2d57610d2d6135e6565b9050602002810190610d3f919061362c565b9050365f610d4d8380613640565b90925090505f610d6360408501602086016133c7565b90505f196001600160a01b03821601610dbe5760405162461bcd60e51b815260206004820152601760248201527f4141393620696e76616c69642061676772656761746f72000000000000000000604482015260640161060a565b6001600160a01b03811615610e59576001600160a01b03811663e3563a4f8484610deb6040890189613685565b6040518563ffffffff1660e01b8152600401610e0a9493929190613826565b5f6040518083038186803b158015610e20575f80fd5b505afa925050508015610e31575060015b610e595760405163086a9f7560e41b81526001600160a01b038216600482015260240161060a565b610e638287613556565b9550505050508080610e7490613569565b915050610d12565b505f816001600160401b03811115610e9657610e96613093565b604051908082528060200260200182016040528015610ecf57816020015b610ebc612f6c565b815260200190600190039081610eb45790505b506040519091507fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972905f90a15f805b84811015610fdd5736888883818110610f1957610f196135e6565b9050602002810190610f2b919061362c565b9050365f610f398380613640565b90925090505f610f4f60408501602086016133c7565b9050815f5b81811015610fc4575f898981518110610f6f57610f6f6135e6565b602002602001015190505f80610f918b898987818110610a7257610a726135e6565b91509150610fa184838389611f3f565b8a610fab81613569565b9b50505050508080610fbc90613569565b915050610f54565b5050505050508080610fd590613569565b915050610efe565b505f8091505f5b858110156110f25736898983818110610fff57610fff6135e6565b9050602002810190611011919061362c565b905061102360408201602083016133c7565b6001600160a01b03167f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d60405160405180910390a2365f6110648380613640565b9092509050805f5b818110156110da576110ae8885858481811061108a5761108a6135e6565b905060200281019061109c91906135fa565b8b8b81518110610b1457610b146135e6565b6110b89088613556565b9650876110c481613569565b98505080806110d290613569565b91505061106c565b505050505080806110ea90613569565b915050610fe4565b506040515f907f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d908290a261112786826121f9565b5050505050610b456001600255565b8315801561114c57506001600160a01b0383163b155b156111995760405162461bcd60e51b815260206004820152601960248201527f41413230206163636f756e74206e6f74206465706c6f79656400000000000000604482015260640161060a565b6014811061120f575f6111af60148284866138a2565b6111b8916138c9565b60601c9050803b5f0361120d5760405162461bcd60e51b815260206004820152601b60248201527f41413330207061796d6173746572206e6f74206465706c6f7965640000000000604482015260640161060a565b505b60405162461bcd60e51b8152602060048201525f602482015260440161060a565b604051632b870d1b60e11b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063570e1a369061128090869086906004016138fe565b6020604051808303815f875af115801561129c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112c09190613911565b604051633653dc0360e11b81526001600160a01b038216600482015290915060240161060a565b5f6112f1826122ee565b6040805160208101929092523090820152466060820152608001604051602081830303815290604052805190602001209050919050565b335f9081526020819052604081206001810154909163ffffffff90911690036113805760405162461bcd60e51b815260206004820152600a6024820152691b9bdd081cdd185ad95960b21b604482015260640161060a565b8054600160701b900460ff166113cc5760405162461bcd60e51b8152602060048201526011602482015270616c726561647920756e7374616b696e6760781b604482015260640161060a565b60018101545f906113e39063ffffffff164261392c565b60018301805469ffffffffffff00000000191664010000000065ffffffffffff841690810291909117909155835460ff60701b1916845560405190815290915033907ffa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a906020016105a1565b335f9081526020819052604090208054600160781b90046001600160701b0316806114b35760405162461bcd60e51b81526020600482015260146024820152734e6f207374616b6520746f20776974686472617760601b604482015260640161060a565b6001820154640100000000900465ffffffffffff166115145760405162461bcd60e51b815260206004820152601d60248201527f6d7573742063616c6c20756e6c6f636b5374616b652829206669727374000000604482015260640161060a565b60018201544264010000000090910465ffffffffffff1611156115795760405162461bcd60e51b815260206004820152601b60248201527f5374616b65207769746864726177616c206973206e6f74206475650000000000604482015260640161060a565b60018201805469ffffffffffffffffffff191690558154600160781b600160e81b0319168255604080516001600160a01b03851681526020810183905233917fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3910160405180910390a25f836001600160a01b0316826040515f6040518083038185875af1925050503d805f811461162c576040519150601f19603f3d011682016040523d82523d5f602084013e611631565b606091505b5050905080610cbb5760405162461bcd60e51b815260206004820152601860248201527f6661696c656420746f207769746864726177207374616b650000000000000000604482015260640161060a565b61168a612f6c565b61169385612306565b5f806116a05f8885611d59565b915091505f6116af83836123db565b90506116b9435f52565b5f6116c55f8a876120d9565b90506116cf435f52565b5f60606001600160a01b038a161561174057896001600160a01b031689896040516116fb929190613952565b5f604051808303815f865af19150503d805f8114611734576040519150601f19603f3d011682016040523d82523d5f602084013e611739565b606091505b5090925090505b866080015183856020015186604001518585604051630116f59360e71b815260040161060a96959493929190613961565b611779612f6c565b61178282612306565b5f8061178f5f8585611d59565b915091505f6117a4845f015160a001516124a5565b8451519091505f906117b5906124a5565b90506117d260405180604001604052805f81526020015f81525090565b365f6117e160408a018a613685565b90925090505f60148210156117f6575f611810565b61180360145f84866138a2565b61180c916138c9565b60601c5b905061181b816124a5565b93505050505f61182b86866123db565b90505f815f015190505f60016001600160a01b0316826001600160a01b03161490505f6040518060c001604052808b6080015181526020018b6040015181526020018315158152602001856020015165ffffffffffff168152602001856040015165ffffffffffff1681526020016118a48c6060015190565b905290506001600160a01b038316158015906118ca57506001600160a01b038316600114155b1561191b575f6040518060400160405280856001600160a01b031681526020016118f3866124a5565b81525090508187878a84604051633ebb2d3960e21b815260040161060a959493929190613a01565b8086868960405163e0cff05f60e01b815260040161060a9493929190613a80565b6001600160a01b0382165f908152602081905260408120805490919061196c9084906001600160701b0316613556565b90506001600160701b038111156119b85760405162461bcd60e51b815260206004820152601060248201526f6465706f736974206f766572666c6f7760801b604482015260640161060a565b81546001600160701b0319166001600160701b03919091161790555050565b5f805f845160208601878987f195945050505050565b60603d828111156119fb5750815b604051602082018101604052818152815f602083013e9392505050565b5f805a85519091505f9081611a2c826124f3565b60a08301519091506001600160a01b038116611a4b5782519350611be9565b8093505f88511115611be957868202955060028a6002811115611a7057611a70613ad6565b14611add57606083015160405163a9a2340960e01b81526001600160a01b0383169163a9a2340991611aaa908e908d908c90600401613aea565b5f604051808303815f88803b158015611ac1575f80fd5b5087f1158015611ad3573d5f803e3d5ffd5b5050505050611be9565b606083015160405163a9a2340960e01b81526001600160a01b0383169163a9a2340991611b12908e908d908c90600401613aea565b5f604051808303815f88803b158015611b29575f80fd5b5087f193505050508015611b3b575060015b611be957611b47613b2e565b806308c379a003611ba05750611b5b613b47565b80611b665750611ba2565b8b81604051602001611b789190613bcf565b60408051601f1981840301815290829052631101335b60e11b825261060a92916004016135ce565b505b8a604051631101335b60e11b815260040161060a9181526040602082018190526012908201527110504d4c081c1bdcdd13dc081c995d995c9d60721b606082015260800190565b5a85038701965081870295508589604001511015611c52578a604051631101335b60e11b815260040161060a91815260406020808301829052908201527f414135312070726566756e642062656c6f772061637475616c476173436f7374606082015260800190565b6040890151869003611c64858261193c565b5f808c6002811115611c7857611c78613ad6565b1490508460a001516001600160a01b0316855f01516001600160a01b03168c602001517f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f8860200151858d8f604051611cea949392919093845291151560208401526040830152606082015260800190565b60405180910390a45050505050505095945050505050565b6002805403611d535760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161060a565b60028055565b5f805f5a8451909150611d6c8682612522565b611d75866112e7565b6020860152604081015160608201516080830151171760e087013517610100870135176effffffffffffffffffffffffffffff811115611df75760405162461bcd60e51b815260206004820152601860248201527f41413934206761732076616c756573206f766572666c6f770000000000000000604482015260640161060a565b5f80611e0284612618565b9050611e108a8a8a84612663565b85516020870151919950919350611e279190612893565b611e7d5789604051631101335b60e11b815260040161060a918152604060208201819052601a908201527f4141323520696e76616c6964206163636f756e74206e6f6e6365000000000000606082015260800190565b611e85435f52565b60a08401516060906001600160a01b031615611ead57611ea88b8b8b85876128df565b975090505b5f5a87039050808b60a001351015611f11578b604051631101335b60e11b815260040161060a918152604060208201819052601e908201527f41413430206f76657220766572696669636174696f6e4761734c696d69740000606082015260800190565b60408a018390528160608b015260c08b01355a8803018a608001818152505050505050505050935093915050565b5f80611f4a85612afb565b91509150816001600160a01b0316836001600160a01b031614611fb05785604051631101335b60e11b815260040161060a9181526040602082018190526014908201527320a0991a1039b4b3b730ba3ab9329032b93937b960611b606082015260800190565b80156120085785604051631101335b60e11b815260040161060a9181526040602082018190526017908201527f414132322065787069726564206f72206e6f7420647565000000000000000000606082015260800190565b5f61201285612afb565b925090506001600160a01b0381161561206e5786604051631101335b60e11b815260040161060a9181526040602082018190526014908201527320a0999a1039b4b3b730ba3ab9329032b93937b960611b606082015260800190565b81156120d05786604051631101335b60e11b815260040161060a9181526040602082018190526021908201527f41413332207061796d61737465722065787069726564206f72206e6f742064756060820152606560f81b608082015260a00190565b50505050505050565b5f805a90505f6120ea846060015190565b905030631d7327566120ff6060880188613685565b87856040518563ffffffff1660e01b81526004016121209493929190613c0c565b6020604051808303815f875af192505050801561215a575060408051601f3d908101601f1916820190925261215791810190613cbe565b60015b6121ed575f60205f803e505f51632152215360e01b81016121b95786604051631101335b60e11b815260040161060a918152604060208201819052600f908201526e41413935206f7574206f662067617360881b606082015260800190565b5f85608001515a6121ca9086613619565b6121d49190613556565b90506121e4886002888685611a18565b945050506121f0565b92505b50509392505050565b6001600160a01b03821661224f5760405162461bcd60e51b815260206004820152601860248201527f4141393020696e76616c69642062656e65666963696172790000000000000000604482015260640161060a565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114612298576040519150601f19603f3d011682016040523d82523d5f602084013e61229d565b606091505b5050905080610b455760405162461bcd60e51b815260206004820152601f60248201527f41413931206661696c65642073656e6420746f2062656e656669636961727900604482015260640161060a565b5f6122f882612b4a565b805190602001209050919050565b3063957122ab6123196040840184613685565b61232660208601866133c7565b612334610120870187613685565b6040518663ffffffff1660e01b8152600401612354959493929190613cd5565b5f6040518083038186803b15801561236a575f80fd5b505afa92505050801561237b575060015b6123d857612387613b2e565b806308c379a0036123ce575061239b613b47565b806123a657506123d0565b8051156123ca575f81604051631101335b60e11b815260040161060a9291906135ce565b5050565b505b3d5f803e3d5ffd5b50565b604080516060810182525f808252602082018190529181018290529061240084612c1a565b90505f61240c84612c1a565b82519091506001600160a01b038116612423575080515b602080840151604080860151928501519085015191929165ffffffffffff8083169085161015612451578193505b8065ffffffffffff168365ffffffffffff16111561246d578092505b5050604080516060810182526001600160a01b03909416845265ffffffffffff92831660208501529116908201529250505092915050565b6040805180820182525f80825260208083018281526001600160a01b03959095168252819052919091208054600160781b90046001600160701b031682526001015463ffffffff1690915290565b60c081015160e08201515f919080820361250e575092915050565b61251a82488301612c89565b949350505050565b61252f60208301836133c7565b6001600160a01b0316815260208083013590820152608080830135604083015260a0830135606083015260c0808401359183019190915260e0808401359183019190915261010083013590820152365f61258d610120850185613685565b9092509050801561260c5760148110156125e95760405162461bcd60e51b815260206004820152601d60248201527f4141393320696e76616c6964207061796d6173746572416e6444617461000000604482015260640161060a565b6125f660145f83856138a2565b6125ff916138c9565b60601c60a0840152610cbb565b5f60a084015250505050565b60a08101515f9081906001600160a01b0316612635576001612638565b60035b60ff1690505f8360800151828560600151028560400151010190508360c00151810292505050919050565b5f805f5a8551805191925090612686898861268160408c018c613685565b612ca0565b60a0820151612693435f52565b5f6001600160a01b0382166126d8576001600160a01b0383165f908152602081905260409020546001600160701b03168881116126d2578089036126d4565b5f5b9150505b606084015160208a0151604051633a871cdd60e01b81526001600160a01b03861692633a871cdd929091612712918f918790600401613d0a565b6020604051808303815f8887f19350505050801561274d575060408051601f3d908101601f1916820190925261274a91810190613cbe565b60015b6127d757612759613b2e565b806308c379a00361278a575061276d613b47565b80612778575061278c565b8b81604051602001611b789190613d2e565b505b8a604051631101335b60e11b815260040161060a918152604060208201819052601690820152754141323320726576657274656420286f72204f4f472960501b606082015260800190565b95506001600160a01b038216612880576001600160a01b0383165f90815260208190526040902080546001600160701b0316808a1115612863578c604051631101335b60e11b815260040161060a9181526040602082018190526017908201527f41413231206469646e2774207061792070726566756e64000000000000000000606082015260800190565b81546001600160701b031916908a90036001600160701b03161790555b5a85039650505050505094509492505050565b6001600160a01b0382165f90815260016020908152604080832084821c80855292528220805484916001600160401b0383169190856128d183613569565b909155501495945050505050565b825160608181015190915f9184811161293a5760405162461bcd60e51b815260206004820152601f60248201527f4141343120746f6f206c6974746c6520766572696669636174696f6e47617300604482015260640161060a565b60a08201516001600160a01b0381165f90815260208190526040902080548784039291906001600160701b0316898110156129c1578c604051631101335b60e11b815260040161060a918152604060208201819052601e908201527f41413331207061796d6173746572206465706f73697420746f6f206c6f770000606082015260800190565b898103825f015f6101000a8154816001600160701b0302191690836001600160701b03160217905550826001600160a01b031663f465c77e858e8e602001518e6040518563ffffffff1660e01b8152600401612a1f93929190613d0a565b5f604051808303815f8887f193505050508015612a5d57506040513d5f823e601f3d908101601f19168201604052612a5a9190810190613d64565b60015b612ae757612a69613b2e565b806308c379a003612a9a5750612a7d613b47565b80612a885750612a9c565b8d81604051602001611b789190613dea565b505b8c604051631101335b60e11b815260040161060a918152604060208201819052601690820152754141333320726576657274656420286f72204f4f472960501b606082015260800190565b909e909d509b505050505050505050505050565b5f80825f03612b0e57505f928392509050565b5f612b1884612c1a565b9050806040015165ffffffffffff16421180612b3f5750806020015165ffffffffffff1642105b905194909350915050565b6060813560208301355f612b69612b646040870187613685565b612f5a565b90505f612b7c612b646060880188613685565b9050608086013560a087013560c088013560e08901356101008a01355f612baa612b646101208e018e613685565b604080516001600160a01b039c909c1660208d01528b81019a909a5260608b019890985250608089019590955260a088019390935260c087019190915260e08601526101008501526101208401526101408084019190915281518084039091018152610160909201905292915050565b604080516060810182525f80825260208201819052918101919091528160a081901c65ffffffffffff81165f03612c54575065ffffffffffff5b604080516060810182526001600160a01b03909316835260d09490941c602083015265ffffffffffff16928101929092525090565b5f818310612c975781612c99565b825b9392505050565b8015610cbb578251516001600160a01b0381163b15612d0b5784604051631101335b60e11b815260040161060a918152604060208201819052601f908201527f414131302073656e64657220616c726561647920636f6e737472756374656400606082015260800190565b835160600151604051632b870d1b60e11b81525f916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163570e1a369190612d6290889088906004016138fe565b6020604051808303815f8887f1158015612d7e573d5f803e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190612da39190613911565b90506001600160a01b038116612e055785604051631101335b60e11b815260040161060a918152604060208201819052601b908201527f4141313320696e6974436f6465206661696c6564206f72204f4f470000000000606082015260800190565b816001600160a01b0316816001600160a01b031614612e6f5785604051631101335b60e11b815260040161060a91815260406020808301829052908201527f4141313420696e6974436f6465206d7573742072657475726e2073656e646572606082015260800190565b806001600160a01b03163b5f03612ed15785604051631101335b60e11b815260040161060a91815260406020808301829052908201527f4141313520696e6974436f6465206d757374206372656174652073656e646572606082015260800190565b5f612edf60148286886138a2565b612ee8916138c9565b60601c9050826001600160a01b031686602001517fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d83895f015160a00151604051612f499291906001600160a01b0392831681529116602082015260400190565b60405180910390a350505050505050565b5f604051828085833790209392505050565b6040518060a00160405280612fc96040518061010001604052805f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f6001600160a01b031681526020015f81526020015f81525090565b81526020015f80191681526020015f81526020015f81526020015f81525090565b5f60208284031215612ffa575f80fd5b813563ffffffff81168114612c99575f80fd5b80356001600160c01b0381168114613023575f80fd5b919050565b5f60208284031215613038575f80fd5b612c998261300d565b6001600160a01b03811681146123d8575f80fd5b803561302381613041565b5f8060408385031215613071575f80fd5b823561307c81613041565b915061308a6020840161300d565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b60a081018181106001600160401b03821117156130c6576130c6613093565b60405250565b61010081018181106001600160401b03821117156130c6576130c6613093565b601f8201601f191681016001600160401b038111828210171561311157613111613093565b6040525050565b5f6001600160401b0382111561313057613130613093565b50601f01601f191660200190565b5f818303610180811215613150575f80fd5b60405161315c816130a7565b8092506101008083121561316e575f80fd5b604051925061317c836130cc565b61318585613055565b8352602085013560208401526040850135604084015260608501356060840152608085013560808401526131bb60a08601613055565b60a084015260c085013560c084015260e085013560e084015282825280850135602083015250610120840135604082015261014084013560608201526101608401356080820152505092915050565b5f8083601f84011261321a575f80fd5b5081356001600160401b03811115613230575f80fd5b602083019150836020828501011115613247575f80fd5b9250929050565b5f805f806101c08587031215613262575f80fd5b84356001600160401b0380821115613278575f80fd5b818701915087601f83011261328b575f80fd5b813561329681613118565b6040516132a382826130ec565b8281528a60208487010111156132b7575f80fd5b826020860160208301375f602084830101528098505050506132dc886020890161313e565b94506101a08701359150808211156132f2575f80fd5b506132ff8782880161320a565b95989497509550505050565b5f8083601f84011261331b575f80fd5b5081356001600160401b03811115613331575f80fd5b6020830191508360208260051b8501011115613247575f80fd5b5f805f6040848603121561335d575f80fd5b83356001600160401b03811115613372575f80fd5b61337e8682870161330b565b909450925050602084013561339281613041565b809150509250925092565b5f80604083850312156133ae575f80fd5b82356133b981613041565b946020939093013593505050565b5f602082840312156133d7575f80fd5b8135612c9981613041565b5f805f805f606086880312156133f6575f80fd5b85356001600160401b038082111561340c575f80fd5b61341889838a0161320a565b90975095506020880135915061342d82613041565b90935060408701359080821115613442575f80fd5b5061344f8882890161320a565b969995985093965092949392505050565b5f8060208385031215613471575f80fd5b82356001600160401b03811115613486575f80fd5b6134928582860161320a565b90969095509350505050565b5f61016082840312156134af575f80fd5b50919050565b5f602082840312156134c5575f80fd5b81356001600160401b038111156134da575f80fd5b61251a8482850161349e565b5f805f80606085870312156134f9575f80fd5b84356001600160401b038082111561350f575f80fd5b61351b8883890161349e565b95506020870135915061352d82613041565b909350604086013590808211156132f2575f80fd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d0057610d00613542565b5f6001820161357a5761357a613542565b5060010190565b5f5b8381101561359b578181015183820152602001613583565b50505f910152565b5f81518084526135ba816020860160208601613581565b601f01601f19169290920160200192915050565b828152604060208201525f61251a60408301846135a3565b634e487b7160e01b5f52603260045260245ffd5b5f823561015e1983360301811261360f575f80fd5b9190910192915050565b81810381811115610d0057610d00613542565b5f8235605e1983360301811261360f575f80fd5b5f808335601e19843603018112613655575f80fd5b8301803591506001600160401b0382111561366e575f80fd5b6020019150600581901b3603821315613247575f80fd5b5f808335601e1984360301811261369a575f80fd5b8301803591506001600160401b038211156136b3575f80fd5b602001915036819003821315613247575f80fd5b5f808335601e198436030181126136dc575f80fd5b83016020810192503590506001600160401b038111156136fa575f80fd5b803603821315613247575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f61016061374e8461374185613055565b6001600160a01b03169052565b6020830135602085015261376560408401846136c7565b8260408701526137788387018284613708565b9250505061378960608401846136c7565b858303606087015261379c838284613708565b925050506080830135608085015260a083013560a085015260c083013560c085015260e083013560e08501526101008084013581860152506101206137e3818501856136c7565b868403838801526137f5848284613708565b9350505050610140613809818501856136c7565b8684038388015261381b848284613708565b979650505050505050565b604080825281018490525f6060600586901b830181019083018783805b8981101561388b57868503605f190184528235368c900361015e19018112613869578283fd5b613875868d8301613730565b9550506020938401939290920191600101613843565b50505050828103602084015261381b818587613708565b5f80858511156138b0575f80fd5b838611156138bc575f80fd5b5050820193919092039150565b6bffffffffffffffffffffffff1981358181169160148510156138f65780818660140360031b1b83161692505b505092915050565b602081525f61251a602083018486613708565b5f60208284031215613921575f80fd5b8151612c9981613041565b65ffffffffffff81811683821601908082111561394b5761394b613542565b5092915050565b818382375f9101908152919050565b8681528560208201525f65ffffffffffff8087166040840152808616606084015250831515608083015260c060a083015261399f60c08301846135a3565b98975050505050505050565b80518252602081015160208301526040810151151560408301525f606082015165ffffffffffff8082166060860152806080850151166080860152505060a082015160c060a085015261251a60c08501826135a3565b5f610140808352613a14818401896139ab565b915050613a2e602083018780518252602090810151910152565b845160608301526020948501516080830152835160a08301529284015160c082015281516001600160a01b031660e0820152908301518051610100830152909201516101209092019190915292915050565b60e081525f613a9260e08301876139ab565b9050613aab602083018680518252602090810151910152565b8351606083015260208401516080830152825160a0830152602083015160c083015295945050505050565b634e487b7160e01b5f52602160045260245ffd5b5f60038510613b0757634e487b7160e01b5f52602160045260245ffd5b84825260606020830152613b1e60608301856135a3565b9050826040830152949350505050565b5f60033d1115613b445760045f803e505f5160e01c5b90565b5f60443d1015613b545790565b6040516003193d81016004833e81513d6001600160401b038160248401118184111715613b8357505050505090565b8285019150815181811115613b9b5750505050505090565b843d8701016020828501011115613bb55750505050505090565b613bc4602082860101876130ec565b509095945050505050565b75020a09a98103837b9ba27b8103932bb32b93a32b21d160551b81525f8251613bff816016850160208701613581565b9190910160160192915050565b5f6101c0808352613c208184018789613708565b9050845160018060a01b03808251166020860152602082015160408601526040820151606086015260608201516080860152608082015160a08601528060a08301511660c08601525060c081015160e085015260e08101516101008501525060208501516101208401526040850151610140840152606085015161016084015260808501516101808401528281036101a084015261381b81856135a3565b5f60208284031215613cce575f80fd5b5051919050565b606081525f613ce8606083018789613708565b6001600160a01b0386166020840152828103604084015261399f818587613708565b606081525f613d1c6060830186613730565b60208301949094525060400152919050565b6e020a09919903932bb32b93a32b21d1608d1b81525f8251613d5781600f850160208701613581565b91909101600f0192915050565b5f8060408385031215613d75575f80fd5b82516001600160401b03811115613d8a575f80fd5b8301601f81018513613d9a575f80fd5b8051613da581613118565b604051613db282826130ec565b828152876020848601011115613dc6575f80fd5b613dd7836020830160208701613581565b6020969096015195979596505050505050565b6e020a09999903932bb32b93a32b21d1608d1b81525f8251613d5781600f85016020870161358156fea2646970667358221220bfa2427621ac9a94fe2d5c7d66261ce13472e713a2ae5494743368e24556d97464736f6c63430008140033608060405234801561000f575f80fd5b506101e28061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063570e1a361461002d575b5f80fd5b61004061003b3660046100e4565b61005c565b6040516001600160a01b03909116815260200160405180910390f35b5f8061006b6014828587610150565b61007491610177565b60601c90505f6100878460148188610150565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92018290525084519495509360209350849250905082850182875af190505f519350806100db575f93505b50505092915050565b5f80602083850312156100f5575f80fd5b823567ffffffffffffffff8082111561010c575f80fd5b818501915085601f83011261011f575f80fd5b81358181111561012d575f80fd5b86602082850101111561013e575f80fd5b60209290920196919550909350505050565b5f808585111561015e575f80fd5b8386111561016a575f80fd5b5050820193919092039150565b6bffffffffffffffffffffffff1981358181169160148510156101a45780818660140360031b1b83161692505b50509291505056fea264697066735822122007943da4c2de99a9a15b2799f96edf34820401c76798dbbc41446cf439d58e6964736f6c63430008140033",
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
