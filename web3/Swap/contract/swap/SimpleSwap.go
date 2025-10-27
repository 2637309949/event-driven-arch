// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// SwapMetaData contains all meta data concerning the Swap contract.
var SwapMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"SwapFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"SwapSucceeded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UNISWAP_ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610f188061001c5f395ff3fe608060405260043610610042575f3560e01c806389a302711461004d5780639d64eb9614610077578063ad5c464814610093578063d8264920146100bd57610049565b3661004957005b5f5ffd5b348015610058575f5ffd5b506100616100e7565b60405161006e919061069e565b60405180910390f35b610091600480360381019061008c91906106fb565b6100ff565b005b34801561009e575f5ffd5b506100a761062f565b6040516100b4919061069e565b60405180910390f35b3480156100c8575f5ffd5b506100d1610647565b6040516100de919061069e565b60405180910390f35b731c7d4b196cb0c7b01d743fbc6116a902379c723881565b5f3411610141576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013890610793565b60405180910390fd5b5f8211610183576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017a90610821565b60405180910390fd5b428110156101c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bd90610889565b60405180910390fd5b5f600267ffffffffffffffff8111156101e2576101e16108a7565b5b6040519080825280602002602001820160405280156102105781602001602082028036833780820191505090505b50905073fff9976782d46cc05630d1f6ebab18b2324d6b14815f8151811061023b5761023a6108d4565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050731c7d4b196cb0c7b01d743fbc6116a902379c72388160018151811061029e5761029d6108d4565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505f73ee567fe1712faf6149d80da1e6934e354124cfe373ffffffffffffffffffffffffffffffffffffffff1663d06ca61f34846040518363ffffffff1660e01b81526004016103289291906109c7565b5f60405180830381865afa158015610342573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061036a9190610b2c565b905083816001835161037c9190610ba0565b8151811061038d5761038c6108d4565b5b602002602001015110156103d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103cd90610c1d565b60405180910390fd5b5f73ee567fe1712faf6149d80da1e6934e354124cfe390508073ffffffffffffffffffffffffffffffffffffffff16637ff36ab534878633896040518663ffffffff1660e01b815260040161042e9493929190610c3b565b5f6040518083038185885af19350505050801561046d57506040513d5f823e3d601f19601f8201168201806040525081019061046a9190610b2c565b60015b6105b057610479610c91565b806308c379a003610525575061048d610cb0565b806104985750610527565b3373ffffffffffffffffffffffffffffffffffffffff167f25e268b6e417ca50f83f2b95db36dfd7e9d897f9b0b714668524e848e863534734836040516104e0929190610d8f565b60405180910390a2806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051c9190610dbd565b60405180910390fd5b505b3373ffffffffffffffffffffffffffffffffffffffff167f25e268b6e417ca50f83f2b95db36dfd7e9d897f9b0b714668524e848e86353473460405161056d9190610e27565b60405180910390a26040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a790610e9d565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f86946606517786ae54b72815847835fee8c1611b7c9856974c531e7e4c5d43c43483600185516105f89190610ba0565b81518110610609576106086108d4565b5b602002602001015160405161061f929190610ebb565b60405180910390a2505050505050565b73fff9976782d46cc05630d1f6ebab18b2324d6b1481565b73ee567fe1712faf6149d80da1e6934e354124cfe381565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106888261065f565b9050919050565b6106988161067e565b82525050565b5f6020820190506106b15f83018461068f565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b6106da816106c8565b81146106e4575f5ffd5b50565b5f813590506106f5816106d1565b92915050565b5f5f60408385031215610711576107106106c0565b5b5f61071e858286016106e7565b925050602061072f858286016106e7565b9150509250929050565b5f82825260208201905092915050565b7f6e6565642045544820746f2073776170000000000000000000000000000000005f82015250565b5f61077d601083610739565b915061078882610749565b602082019050919050565b5f6020820190508181035f8301526107aa81610771565b9050919050565b7f616d6f756e744f75744d696e206d7573742062652067726561746572207468615f8201527f6e20300000000000000000000000000000000000000000000000000000000000602082015250565b5f61080b602383610739565b9150610816826107b1565b604082019050919050565b5f6020820190508181035f830152610838816107ff565b9050919050565b7f696e76616c696420646561646c696e65000000000000000000000000000000005f82015250565b5f610873601083610739565b915061087e8261083f565b602082019050919050565b5f6020820190508181035f8301526108a081610867565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b61090a816106c8565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6109428161067e565b82525050565b5f6109538383610939565b60208301905092915050565b5f602082019050919050565b5f61097582610910565b61097f818561091a565b935061098a8361092a565b805f5b838110156109ba5781516109a18882610948565b97506109ac8361095f565b92505060018101905061098d565b5085935050505092915050565b5f6040820190506109da5f830185610901565b81810360208301526109ec818461096b565b90509392505050565b5f5ffd5b5f601f19601f8301169050919050565b610a12826109f9565b810181811067ffffffffffffffff82111715610a3157610a306108a7565b5b80604052505050565b5f610a436106b7565b9050610a4f8282610a09565b919050565b5f67ffffffffffffffff821115610a6e57610a6d6108a7565b5b602082029050602081019050919050565b5f5ffd5b5f81519050610a91816106d1565b92915050565b5f610aa9610aa484610a54565b610a3a565b90508083825260208201905060208402830185811115610acc57610acb610a7f565b5b835b81811015610af55780610ae18882610a83565b845260208401935050602081019050610ace565b5050509392505050565b5f82601f830112610b1357610b126109f5565b5b8151610b23848260208601610a97565b91505092915050565b5f60208284031215610b4157610b406106c0565b5b5f82015167ffffffffffffffff811115610b5e57610b5d6106c4565b5b610b6a84828501610aff565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610baa826106c8565b9150610bb5836106c8565b9250828203905081811115610bcd57610bcc610b73565b5b92915050565b7f496e73756666696369656e74206f757470757420616d6f756e740000000000005f82015250565b5f610c07601a83610739565b9150610c1282610bd3565b602082019050919050565b5f6020820190508181035f830152610c3481610bfb565b9050919050565b5f608082019050610c4e5f830187610901565b8181036020830152610c60818661096b565b9050610c6f604083018561068f565b610c7c6060830184610901565b95945050505050565b5f8160e01c9050919050565b5f60033d1115610cad5760045f5f3e610caa5f51610c85565b90505b90565b5f60443d10610d3c57610cc16106b7565b60043d036004823e80513d602482011167ffffffffffffffff82111715610ce9575050610d3c565b808201805167ffffffffffffffff811115610d075750505050610d3c565b80602083010160043d038501811115610d24575050505050610d3c565b610d3382602001850186610a09565b82955050505050505b90565b5f81519050919050565b8281835e5f83830152505050565b5f610d6182610d3f565b610d6b8185610739565b9350610d7b818560208601610d49565b610d84816109f9565b840191505092915050565b5f604082019050610da25f830185610901565b8181036020830152610db48184610d57565b90509392505050565b5f6020820190508181035f830152610dd58184610d57565b905092915050565b7f556e6b6e6f776e206572726f72000000000000000000000000000000000000005f82015250565b5f610e11600d83610739565b9150610e1c82610ddd565b602082019050919050565b5f604082019050610e3a5f830184610901565b8181036020830152610e4b81610e05565b905092915050565b7f53776170206661696c6564207769746820756e6b6e6f776e206572726f7200005f82015250565b5f610e87601e83610739565b9150610e9282610e53565b602082019050919050565b5f6020820190508181035f830152610eb481610e7b565b9050919050565b5f604082019050610ece5f830185610901565b610edb6020830184610901565b939250505056fea2646970667358221220afc1075f83a7247e088f16b55165cd61739c464395f231e798cc1bce4945334b64736f6c634300081e0033",
}

// SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMetaData.ABI instead.
var SwapABI = SwapMetaData.ABI

// SwapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapMetaData.Bin instead.
var SwapBin = SwapMetaData.Bin

// DeploySwap deploys a new Ethereum contract, binding an instance of Swap to it.
func DeploySwap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swap, error) {
	parsed, err := SwapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// UNISWAPROUTER is a free data retrieval call binding the contract method 0xd8264920.
//
// Solidity: function UNISWAP_ROUTER() view returns(address)
func (_Swap *SwapCaller) UNISWAPROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "UNISWAP_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPROUTER is a free data retrieval call binding the contract method 0xd8264920.
//
// Solidity: function UNISWAP_ROUTER() view returns(address)
func (_Swap *SwapSession) UNISWAPROUTER() (common.Address, error) {
	return _Swap.Contract.UNISWAPROUTER(&_Swap.CallOpts)
}

// UNISWAPROUTER is a free data retrieval call binding the contract method 0xd8264920.
//
// Solidity: function UNISWAP_ROUTER() view returns(address)
func (_Swap *SwapCallerSession) UNISWAPROUTER() (common.Address, error) {
	return _Swap.Contract.UNISWAPROUTER(&_Swap.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Swap *SwapCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Swap *SwapSession) USDC() (common.Address, error) {
	return _Swap.Contract.USDC(&_Swap.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Swap *SwapCallerSession) USDC() (common.Address, error) {
	return _Swap.Contract.USDC(&_Swap.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swap *SwapCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swap *SwapSession) WETH() (common.Address, error) {
	return _Swap.Contract.WETH(&_Swap.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swap *SwapCallerSession) WETH() (common.Address, error) {
	return _Swap.Contract.WETH(&_Swap.CallOpts)
}

// SwapETHForUSDC is a paid mutator transaction binding the contract method 0x9d64eb96.
//
// Solidity: function swapETHForUSDC(uint256 amountOutMin, uint256 deadline) payable returns()
func (_Swap *SwapTransactor) SwapETHForUSDC(opts *bind.TransactOpts, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swapETHForUSDC", amountOutMin, deadline)
}

// SwapETHForUSDC is a paid mutator transaction binding the contract method 0x9d64eb96.
//
// Solidity: function swapETHForUSDC(uint256 amountOutMin, uint256 deadline) payable returns()
func (_Swap *SwapSession) SwapETHForUSDC(amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SwapETHForUSDC(&_Swap.TransactOpts, amountOutMin, deadline)
}

// SwapETHForUSDC is a paid mutator transaction binding the contract method 0x9d64eb96.
//
// Solidity: function swapETHForUSDC(uint256 amountOutMin, uint256 deadline) payable returns()
func (_Swap *SwapTransactorSession) SwapETHForUSDC(amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SwapETHForUSDC(&_Swap.TransactOpts, amountOutMin, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapSession) Receive() (*types.Transaction, error) {
	return _Swap.Contract.Receive(&_Swap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapTransactorSession) Receive() (*types.Transaction, error) {
	return _Swap.Contract.Receive(&_Swap.TransactOpts)
}

// SwapSwapFailedIterator is returned from FilterSwapFailed and is used to iterate over the raw logs and unpacked data for SwapFailed events raised by the Swap contract.
type SwapSwapFailedIterator struct {
	Event *SwapSwapFailed // Event containing the contract specifics and raw log

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
func (it *SwapSwapFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapSwapFailed)
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
		it.Event = new(SwapSwapFailed)
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
func (it *SwapSwapFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapSwapFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapSwapFailed represents a SwapFailed event raised by the Swap contract.
type SwapSwapFailed struct {
	User      common.Address
	EthAmount *big.Int
	Reason    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapFailed is a free log retrieval operation binding the contract event 0x25e268b6e417ca50f83f2b95db36dfd7e9d897f9b0b714668524e848e8635347.
//
// Solidity: event SwapFailed(address indexed user, uint256 ethAmount, string reason)
func (_Swap *SwapFilterer) FilterSwapFailed(opts *bind.FilterOpts, user []common.Address) (*SwapSwapFailedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "SwapFailed", userRule)
	if err != nil {
		return nil, err
	}
	return &SwapSwapFailedIterator{contract: _Swap.contract, event: "SwapFailed", logs: logs, sub: sub}, nil
}

// WatchSwapFailed is a free log subscription operation binding the contract event 0x25e268b6e417ca50f83f2b95db36dfd7e9d897f9b0b714668524e848e8635347.
//
// Solidity: event SwapFailed(address indexed user, uint256 ethAmount, string reason)
func (_Swap *SwapFilterer) WatchSwapFailed(opts *bind.WatchOpts, sink chan<- *SwapSwapFailed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "SwapFailed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapSwapFailed)
				if err := _Swap.contract.UnpackLog(event, "SwapFailed", log); err != nil {
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

// ParseSwapFailed is a log parse operation binding the contract event 0x25e268b6e417ca50f83f2b95db36dfd7e9d897f9b0b714668524e848e8635347.
//
// Solidity: event SwapFailed(address indexed user, uint256 ethAmount, string reason)
func (_Swap *SwapFilterer) ParseSwapFailed(log types.Log) (*SwapSwapFailed, error) {
	event := new(SwapSwapFailed)
	if err := _Swap.contract.UnpackLog(event, "SwapFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapSwapSucceededIterator is returned from FilterSwapSucceeded and is used to iterate over the raw logs and unpacked data for SwapSucceeded events raised by the Swap contract.
type SwapSwapSucceededIterator struct {
	Event *SwapSwapSucceeded // Event containing the contract specifics and raw log

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
func (it *SwapSwapSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapSwapSucceeded)
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
		it.Event = new(SwapSwapSucceeded)
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
func (it *SwapSwapSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapSwapSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapSwapSucceeded represents a SwapSucceeded event raised by the Swap contract.
type SwapSwapSucceeded struct {
	User        common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapSucceeded is a free log retrieval operation binding the contract event 0x86946606517786ae54b72815847835fee8c1611b7c9856974c531e7e4c5d43c4.
//
// Solidity: event SwapSucceeded(address indexed user, uint256 ethAmount, uint256 tokenAmount)
func (_Swap *SwapFilterer) FilterSwapSucceeded(opts *bind.FilterOpts, user []common.Address) (*SwapSwapSucceededIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "SwapSucceeded", userRule)
	if err != nil {
		return nil, err
	}
	return &SwapSwapSucceededIterator{contract: _Swap.contract, event: "SwapSucceeded", logs: logs, sub: sub}, nil
}

// WatchSwapSucceeded is a free log subscription operation binding the contract event 0x86946606517786ae54b72815847835fee8c1611b7c9856974c531e7e4c5d43c4.
//
// Solidity: event SwapSucceeded(address indexed user, uint256 ethAmount, uint256 tokenAmount)
func (_Swap *SwapFilterer) WatchSwapSucceeded(opts *bind.WatchOpts, sink chan<- *SwapSwapSucceeded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "SwapSucceeded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapSwapSucceeded)
				if err := _Swap.contract.UnpackLog(event, "SwapSucceeded", log); err != nil {
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

// ParseSwapSucceeded is a log parse operation binding the contract event 0x86946606517786ae54b72815847835fee8c1611b7c9856974c531e7e4c5d43c4.
//
// Solidity: event SwapSucceeded(address indexed user, uint256 ethAmount, uint256 tokenAmount)
func (_Swap *SwapFilterer) ParseSwapSucceeded(log types.Log) (*SwapSwapSucceeded, error) {
	event := new(SwapSwapSucceeded)
	if err := _Swap.contract.UnpackLog(event, "SwapSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
