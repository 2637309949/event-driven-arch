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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"SwapFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"SwapSucceeded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UNISWAP_ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"swapETHForUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610b418061001c5f395ff3fe608060405260043610610042575f3560e01c8063451ae68d1461004d57806389a3027114610069578063ad5c464814610093578063d8264920146100bd57610049565b3661004957005b5f5ffd5b610067600480360381019061006291906104c0565b6100e7565b005b348015610074575f5ffd5b5061007d610434565b60405161008a919061052a565b60405180910390f35b34801561009e575f5ffd5b506100a761044c565b6040516100b4919061052a565b60405180910390f35b3480156100c8575f5ffd5b506100d1610464565b6040516100de919061052a565b60405180910390f35b5f3411610129576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101209061059d565b60405180910390fd5b5f600267ffffffffffffffff811115610145576101446105bb565b5b6040519080825280602002602001820160405280156101735781602001602082028036833780820191505090505b50905073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2815f8151811061019e5761019d6105e8565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507307865c6e87b9f70255377e024ace6630c1eaa37f81600181518110610201576102006105e8565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505f73e592427a0aece92de3edee1f18e0157c0586156490508073ffffffffffffffffffffffffffffffffffffffff16637ff36ab53485853361012c426102819190610642565b6040518663ffffffff1660e01b81526004016102a0949392919061073b565b5f6040518083038185885af1935050505080156102df57506040513d5f823e3d601f19601f820116820180604052508101906102dc91906108bc565b60015b6103b6576102eb61090f565b806308c379a00361036057506102ff61092e565b8061030a5750610362565b3373ffffffffffffffffffffffffffffffffffffffff167f25e268b6e417ca50f83f2b95db36dfd7e9d897f9b0b714668524e848e86353473483604051610352929190610a0d565b60405180910390a2506103b1565b505b3373ffffffffffffffffffffffffffffffffffffffff167f25e268b6e417ca50f83f2b95db36dfd7e9d897f9b0b714668524e848e8635347346040516103a89190610a85565b60405180910390a25b61042f565b3373ffffffffffffffffffffffffffffffffffffffff167f86946606517786ae54b72815847835fee8c1611b7c9856974c531e7e4c5d43c43483600185516103fe9190610ab1565b8151811061040f5761040e6105e8565b5b6020026020010151604051610425929190610ae4565b60405180910390a2505b505050565b7307865c6e87b9f70255377e024ace6630c1eaa37f81565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281565b73e592427a0aece92de3edee1f18e0157c0586156481565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61049f8161048d565b81146104a9575f5ffd5b50565b5f813590506104ba81610496565b92915050565b5f602082840312156104d5576104d4610485565b5b5f6104e2848285016104ac565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610514826104eb565b9050919050565b6105248161050a565b82525050565b5f60208201905061053d5f83018461051b565b92915050565b5f82825260208201905092915050565b7f4e6565642045544820746f2073776170000000000000000000000000000000005f82015250565b5f610587601083610543565b915061059282610553565b602082019050919050565b5f6020820190508181035f8301526105b48161057b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61064c8261048d565b91506106578361048d565b925082820190508082111561066f5761066e610615565b5b92915050565b61067e8161048d565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6106b68161050a565b82525050565b5f6106c783836106ad565b60208301905092915050565b5f602082019050919050565b5f6106e982610684565b6106f3818561068e565b93506106fe8361069e565b805f5b8381101561072e57815161071588826106bc565b9750610720836106d3565b925050600181019050610701565b5085935050505092915050565b5f60808201905061074e5f830187610675565b818103602083015261076081866106df565b905061076f604083018561051b565b61077c6060830184610675565b95945050505050565b5f5ffd5b5f601f19601f8301169050919050565b6107a282610789565b810181811067ffffffffffffffff821117156107c1576107c06105bb565b5b80604052505050565b5f6107d361047c565b90506107df8282610799565b919050565b5f67ffffffffffffffff8211156107fe576107fd6105bb565b5b602082029050602081019050919050565b5f5ffd5b5f8151905061082181610496565b92915050565b5f610839610834846107e4565b6107ca565b9050808382526020820190506020840283018581111561085c5761085b61080f565b5b835b8181101561088557806108718882610813565b84526020840193505060208101905061085e565b5050509392505050565b5f82601f8301126108a3576108a2610785565b5b81516108b3848260208601610827565b91505092915050565b5f602082840312156108d1576108d0610485565b5b5f82015167ffffffffffffffff8111156108ee576108ed610489565b5b6108fa8482850161088f565b91505092915050565b5f8160e01c9050919050565b5f60033d111561092b5760045f5f3e6109285f51610903565b90505b90565b5f60443d106109ba5761093f61047c565b60043d036004823e80513d602482011167ffffffffffffffff821117156109675750506109ba565b808201805167ffffffffffffffff81111561098557505050506109ba565b80602083010160043d0385018111156109a25750505050506109ba565b6109b182602001850186610799565b82955050505050505b90565b5f81519050919050565b8281835e5f83830152505050565b5f6109df826109bd565b6109e98185610543565b93506109f98185602086016109c7565b610a0281610789565b840191505092915050565b5f604082019050610a205f830185610675565b8181036020830152610a3281846109d5565b90509392505050565b7f556e6b6e6f776e206572726f72000000000000000000000000000000000000005f82015250565b5f610a6f600d83610543565b9150610a7a82610a3b565b602082019050919050565b5f604082019050610a985f830184610675565b8181036020830152610aa981610a63565b905092915050565b5f610abb8261048d565b9150610ac68361048d565b9250828203905081811115610ade57610add610615565b5b92915050565b5f604082019050610af75f830185610675565b610b046020830184610675565b939250505056fea2646970667358221220c5e3d46be7c22d53b2789ad26c3226e9091099b0ee331160b0f3ddb0b30893fe64736f6c634300081e0033",
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

// SwapETHForUSDC is a paid mutator transaction binding the contract method 0x451ae68d.
//
// Solidity: function swapETHForUSDC(uint256 amountOutMin) payable returns()
func (_Swap *SwapTransactor) SwapETHForUSDC(opts *bind.TransactOpts, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swapETHForUSDC", amountOutMin)
}

// SwapETHForUSDC is a paid mutator transaction binding the contract method 0x451ae68d.
//
// Solidity: function swapETHForUSDC(uint256 amountOutMin) payable returns()
func (_Swap *SwapSession) SwapETHForUSDC(amountOutMin *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SwapETHForUSDC(&_Swap.TransactOpts, amountOutMin)
}

// SwapETHForUSDC is a paid mutator transaction binding the contract method 0x451ae68d.
//
// Solidity: function swapETHForUSDC(uint256 amountOutMin) payable returns()
func (_Swap *SwapTransactorSession) SwapETHForUSDC(amountOutMin *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SwapETHForUSDC(&_Swap.TransactOpts, amountOutMin)
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
