// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// AaveMetaData contains all meta data concerning the Aave contract.
var AaveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddressesProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositedToAave\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"WithdrawnFromAave\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToAave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provider\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawFromAave\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161125e38038061125e8339818101604052810190610031919061013e565b60015f819055505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161009d906101c3565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506101e1565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61010d826100e4565b9050919050565b61011d81610103565b8114610127575f5ffd5b50565b5f8151905061013881610114565b92915050565b5f60208284031215610153576101526100e0565b5b5f6101608482850161012a565b91505092915050565b5f82825260208201905092915050565b7f70726f7669646572207a65726f000000000000000000000000000000000000005f82015250565b5f6101ad600d83610169565b91506101b882610179565b602082019050919050565b5f6020820190508181035f8301526101da816101a1565b9050919050565b60805161105f6101ff5f395f818160d601526107be015261105f5ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063085d48831461004e57806377662ffc1461006c578063ad9b3af714610088578063eeb149e7146100b8575b5f5ffd5b6100566100d4565b604051610063919061093a565b60405180910390f35b610086600480360381019061008191906109c5565b6100f8565b005b6100a2600480360381019061009d91906109c5565b6101f7565b6040516100af9190610a24565b60405180910390f35b6100d260048036038101906100cd9190610a3d565b610389565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b610100610764565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361016e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016590610ad5565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb82846040518363ffffffff1660e01b81526004016101a9929190610b02565b6020604051808303815f875af11580156101c5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e99190610b5e565b506101f26107b1565b505050565b5f610200610764565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361026e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026590610ad5565b60405180910390fd5b5f6102776107ba565b90505f8173ffffffffffffffffffffffffffffffffffffffff166369328dec8787876040518463ffffffff1660e01b81526004016102b793929190610b89565b6020604051808303815f875af11580156102d3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102f79190610bd2565b90508373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ff4233c2492c8fe129470230aba79a18b71f74c3b98de673b67a264211d833c698460405161036d9190610a24565b60405180910390a480925050506103826107b1565b9392505050565b610391610764565b5f81116103d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ca90610c47565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610441576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043890610caf565b60405180910390fd5b5f8290505f8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b815260040161048393929190610ccd565b6020604051808303815f875af115801561049f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104c39190610b5e565b905080610505576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104fc90610d4c565b60405180910390fd5b5f61050e6107ba565b90508273ffffffffffffffffffffffffffffffffffffffff1663095ea7b3825f6040518363ffffffff1660e01b815260040161054b929190610da3565b6020604051808303815f875af1158015610567573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058b9190610b5e565b6105ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c190610e14565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b382866040518363ffffffff1660e01b8152600401610605929190610b02565b6020604051808303815f875af1158015610621573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106459190610b5e565b610684576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067b90610e7c565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663617ba0378686335f6040518563ffffffff1660e01b81526004016106c39493929190610ed7565b5f604051808303815f87803b1580156106da575f5ffd5b505af11580156106ec573d5f5f3e3d5ffd5b505050508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fc6fd639034f160abf697c51f38caecb0a761839f3d541152d245177e0611bf3e8660405161074d9190610a24565b60405180910390a35050506107606107b1565b5050565b60025f54036107a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079f90610f64565b60405180910390fd5b60025f81905550565b60015f81905550565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663026b1d5f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610825573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108499190610f96565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b09061100b565b60405180910390fd5b8091505090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f6109026108fd6108f8846108c0565b6108df565b6108c0565b9050919050565b5f610913826108e8565b9050919050565b5f61092482610909565b9050919050565b6109348161091a565b82525050565b5f60208201905061094d5f83018461092b565b92915050565b5f5ffd5b5f610961826108c0565b9050919050565b61097181610957565b811461097b575f5ffd5b50565b5f8135905061098c81610968565b92915050565b5f819050919050565b6109a481610992565b81146109ae575f5ffd5b50565b5f813590506109bf8161099b565b92915050565b5f5f5f606084860312156109dc576109db610953565b5b5f6109e98682870161097e565b93505060206109fa868287016109b1565b9250506040610a0b8682870161097e565b9150509250925092565b610a1e81610992565b82525050565b5f602082019050610a375f830184610a15565b92915050565b5f5f60408385031215610a5357610a52610953565b5b5f610a608582860161097e565b9250506020610a71858286016109b1565b9150509250929050565b5f82825260208201905092915050565b7f7a65726f20746f000000000000000000000000000000000000000000000000005f82015250565b5f610abf600783610a7b565b9150610aca82610a8b565b602082019050919050565b5f6020820190508181035f830152610aec81610ab3565b9050919050565b610afc81610957565b82525050565b5f604082019050610b155f830185610af3565b610b226020830184610a15565b9392505050565b5f8115159050919050565b610b3d81610b29565b8114610b47575f5ffd5b50565b5f81519050610b5881610b34565b92915050565b5f60208284031215610b7357610b72610953565b5b5f610b8084828501610b4a565b91505092915050565b5f606082019050610b9c5f830186610af3565b610ba96020830185610a15565b610bb66040830184610af3565b949350505050565b5f81519050610bcc8161099b565b92915050565b5f60208284031215610be757610be6610953565b5b5f610bf484828501610bbe565b91505092915050565b7f7a65726f20616d6f756e740000000000000000000000000000000000000000005f82015250565b5f610c31600b83610a7b565b9150610c3c82610bfd565b602082019050919050565b5f6020820190508181035f830152610c5e81610c25565b9050919050565b7f7a65726f206173736574000000000000000000000000000000000000000000005f82015250565b5f610c99600a83610a7b565b9150610ca482610c65565b602082019050919050565b5f6020820190508181035f830152610cc681610c8d565b9050919050565b5f606082019050610ce05f830186610af3565b610ced6020830185610af3565b610cfa6040830184610a15565b949350505050565b7f7472616e7366657246726f6d206661696c6564000000000000000000000000005f82015250565b5f610d36601383610a7b565b9150610d4182610d02565b602082019050919050565b5f6020820190508181035f830152610d6381610d2a565b9050919050565b5f819050919050565b5f610d8d610d88610d8384610d6a565b6108df565b610992565b9050919050565b610d9d81610d73565b82525050565b5f604082019050610db65f830185610af3565b610dc36020830184610d94565b9392505050565b7f617070726f7665207265736574206661696c65640000000000000000000000005f82015250565b5f610dfe601483610a7b565b9150610e0982610dca565b602082019050919050565b5f6020820190508181035f830152610e2b81610df2565b9050919050565b7f617070726f7665206661696c65640000000000000000000000000000000000005f82015250565b5f610e66600e83610a7b565b9150610e7182610e32565b602082019050919050565b5f6020820190508181035f830152610e9381610e5a565b9050919050565b5f61ffff82169050919050565b5f610ec1610ebc610eb784610d6a565b6108df565b610e9a565b9050919050565b610ed181610ea7565b82525050565b5f608082019050610eea5f830187610af3565b610ef76020830186610a15565b610f046040830185610af3565b610f116060830184610ec8565b95945050505050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c005f82015250565b5f610f4e601f83610a7b565b9150610f5982610f1a565b602082019050919050565b5f6020820190508181035f830152610f7b81610f42565b9050919050565b5f81519050610f9081610968565b92915050565b5f60208284031215610fab57610faa610953565b5b5f610fb884828501610f82565b91505092915050565b7f706f6f6c206e6f742073657400000000000000000000000000000000000000005f82015250565b5f610ff5600c83610a7b565b915061100082610fc1565b602082019050919050565b5f6020820190508181035f83015261102281610fe9565b905091905056fea2646970667358221220b41ba76cb5fb3e4b364e653c8e9fc3bc3ba1f08034cb67756e900ed329bd6cdb64736f6c634300081e0033",
}

// AaveABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveMetaData.ABI instead.
var AaveABI = AaveMetaData.ABI

// AaveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AaveMetaData.Bin instead.
var AaveBin = AaveMetaData.Bin

// DeployAave deploys a new Ethereum contract, binding an instance of Aave to it.
func DeployAave(auth *bind.TransactOpts, backend bind.ContractBackend, _poolAddressesProvider common.Address) (common.Address, *types.Transaction, *Aave, error) {
	parsed, err := AaveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AaveBin), backend, _poolAddressesProvider)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aave{AaveCaller: AaveCaller{contract: contract}, AaveTransactor: AaveTransactor{contract: contract}, AaveFilterer: AaveFilterer{contract: contract}}, nil
}

// Aave is an auto generated Go binding around an Ethereum contract.
type Aave struct {
	AaveCaller     // Read-only binding to the contract
	AaveTransactor // Write-only binding to the contract
	AaveFilterer   // Log filterer for contract events
}

// AaveCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveSession struct {
	Contract     *Aave             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveCallerSession struct {
	Contract *AaveCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AaveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveTransactorSession struct {
	Contract     *AaveTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveRaw struct {
	Contract *Aave // Generic contract binding to access the raw methods on
}

// AaveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveCallerRaw struct {
	Contract *AaveCaller // Generic read-only contract binding to access the raw methods on
}

// AaveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveTransactorRaw struct {
	Contract *AaveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAave creates a new instance of Aave, bound to a specific deployed contract.
func NewAave(address common.Address, backend bind.ContractBackend) (*Aave, error) {
	contract, err := bindAave(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aave{AaveCaller: AaveCaller{contract: contract}, AaveTransactor: AaveTransactor{contract: contract}, AaveFilterer: AaveFilterer{contract: contract}}, nil
}

// NewAaveCaller creates a new read-only instance of Aave, bound to a specific deployed contract.
func NewAaveCaller(address common.Address, caller bind.ContractCaller) (*AaveCaller, error) {
	contract, err := bindAave(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveCaller{contract: contract}, nil
}

// NewAaveTransactor creates a new write-only instance of Aave, bound to a specific deployed contract.
func NewAaveTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveTransactor, error) {
	contract, err := bindAave(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveTransactor{contract: contract}, nil
}

// NewAaveFilterer creates a new log filterer instance of Aave, bound to a specific deployed contract.
func NewAaveFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveFilterer, error) {
	contract, err := bindAave(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveFilterer{contract: contract}, nil
}

// bindAave binds a generic wrapper to an already deployed contract.
func bindAave(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.AaveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transact(opts, method, params...)
}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() view returns(address)
func (_Aave *AaveCaller) Provider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() view returns(address)
func (_Aave *AaveSession) Provider() (common.Address, error) {
	return _Aave.Contract.Provider(&_Aave.CallOpts)
}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() view returns(address)
func (_Aave *AaveCallerSession) Provider() (common.Address, error) {
	return _Aave.Contract.Provider(&_Aave.CallOpts)
}

// DepositToAave is a paid mutator transaction binding the contract method 0xeeb149e7.
//
// Solidity: function depositToAave(address asset, uint256 amount) returns()
func (_Aave *AaveTransactor) DepositToAave(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "depositToAave", asset, amount)
}

// DepositToAave is a paid mutator transaction binding the contract method 0xeeb149e7.
//
// Solidity: function depositToAave(address asset, uint256 amount) returns()
func (_Aave *AaveSession) DepositToAave(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.DepositToAave(&_Aave.TransactOpts, asset, amount)
}

// DepositToAave is a paid mutator transaction binding the contract method 0xeeb149e7.
//
// Solidity: function depositToAave(address asset, uint256 amount) returns()
func (_Aave *AaveTransactorSession) DepositToAave(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.DepositToAave(&_Aave.TransactOpts, asset, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0x77662ffc.
//
// Solidity: function rescueERC20(address tokenAddress, uint256 amount, address to) returns()
func (_Aave *AaveTransactor) RescueERC20(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "rescueERC20", tokenAddress, amount, to)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0x77662ffc.
//
// Solidity: function rescueERC20(address tokenAddress, uint256 amount, address to) returns()
func (_Aave *AaveSession) RescueERC20(tokenAddress common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Aave.Contract.RescueERC20(&_Aave.TransactOpts, tokenAddress, amount, to)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0x77662ffc.
//
// Solidity: function rescueERC20(address tokenAddress, uint256 amount, address to) returns()
func (_Aave *AaveTransactorSession) RescueERC20(tokenAddress common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Aave.Contract.RescueERC20(&_Aave.TransactOpts, tokenAddress, amount, to)
}

// WithdrawFromAave is a paid mutator transaction binding the contract method 0xad9b3af7.
//
// Solidity: function withdrawFromAave(address asset, uint256 amount, address to) returns(uint256)
func (_Aave *AaveTransactor) WithdrawFromAave(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "withdrawFromAave", asset, amount, to)
}

// WithdrawFromAave is a paid mutator transaction binding the contract method 0xad9b3af7.
//
// Solidity: function withdrawFromAave(address asset, uint256 amount, address to) returns(uint256)
func (_Aave *AaveSession) WithdrawFromAave(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Aave.Contract.WithdrawFromAave(&_Aave.TransactOpts, asset, amount, to)
}

// WithdrawFromAave is a paid mutator transaction binding the contract method 0xad9b3af7.
//
// Solidity: function withdrawFromAave(address asset, uint256 amount, address to) returns(uint256)
func (_Aave *AaveTransactorSession) WithdrawFromAave(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Aave.Contract.WithdrawFromAave(&_Aave.TransactOpts, asset, amount, to)
}

// AaveDepositedToAaveIterator is returned from FilterDepositedToAave and is used to iterate over the raw logs and unpacked data for DepositedToAave events raised by the Aave contract.
type AaveDepositedToAaveIterator struct {
	Event *AaveDepositedToAave // Event containing the contract specifics and raw log

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
func (it *AaveDepositedToAaveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveDepositedToAave)
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
		it.Event = new(AaveDepositedToAave)
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
func (it *AaveDepositedToAaveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveDepositedToAaveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveDepositedToAave represents a DepositedToAave event raised by the Aave contract.
type AaveDepositedToAave struct {
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositedToAave is a free log retrieval operation binding the contract event 0xc6fd639034f160abf697c51f38caecb0a761839f3d541152d245177e0611bf3e.
//
// Solidity: event DepositedToAave(address indexed user, address indexed asset, uint256 amount)
func (_Aave *AaveFilterer) FilterDepositedToAave(opts *bind.FilterOpts, user []common.Address, asset []common.Address) (*AaveDepositedToAaveIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "DepositedToAave", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveDepositedToAaveIterator{contract: _Aave.contract, event: "DepositedToAave", logs: logs, sub: sub}, nil
}

// WatchDepositedToAave is a free log subscription operation binding the contract event 0xc6fd639034f160abf697c51f38caecb0a761839f3d541152d245177e0611bf3e.
//
// Solidity: event DepositedToAave(address indexed user, address indexed asset, uint256 amount)
func (_Aave *AaveFilterer) WatchDepositedToAave(opts *bind.WatchOpts, sink chan<- *AaveDepositedToAave, user []common.Address, asset []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "DepositedToAave", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveDepositedToAave)
				if err := _Aave.contract.UnpackLog(event, "DepositedToAave", log); err != nil {
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

// ParseDepositedToAave is a log parse operation binding the contract event 0xc6fd639034f160abf697c51f38caecb0a761839f3d541152d245177e0611bf3e.
//
// Solidity: event DepositedToAave(address indexed user, address indexed asset, uint256 amount)
func (_Aave *AaveFilterer) ParseDepositedToAave(log types.Log) (*AaveDepositedToAave, error) {
	event := new(AaveDepositedToAave)
	if err := _Aave.contract.UnpackLog(event, "DepositedToAave", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveWithdrawnFromAaveIterator is returned from FilterWithdrawnFromAave and is used to iterate over the raw logs and unpacked data for WithdrawnFromAave events raised by the Aave contract.
type AaveWithdrawnFromAaveIterator struct {
	Event *AaveWithdrawnFromAave // Event containing the contract specifics and raw log

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
func (it *AaveWithdrawnFromAaveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveWithdrawnFromAave)
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
		it.Event = new(AaveWithdrawnFromAave)
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
func (it *AaveWithdrawnFromAaveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveWithdrawnFromAaveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveWithdrawnFromAave represents a WithdrawnFromAave event raised by the Aave contract.
type AaveWithdrawnFromAave struct {
	Caller common.Address
	Asset  common.Address
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFromAave is a free log retrieval operation binding the contract event 0xf4233c2492c8fe129470230aba79a18b71f74c3b98de673b67a264211d833c69.
//
// Solidity: event WithdrawnFromAave(address indexed caller, address indexed asset, uint256 amount, address indexed to)
func (_Aave *AaveFilterer) FilterWithdrawnFromAave(opts *bind.FilterOpts, caller []common.Address, asset []common.Address, to []common.Address) (*AaveWithdrawnFromAaveIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "WithdrawnFromAave", callerRule, assetRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveWithdrawnFromAaveIterator{contract: _Aave.contract, event: "WithdrawnFromAave", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFromAave is a free log subscription operation binding the contract event 0xf4233c2492c8fe129470230aba79a18b71f74c3b98de673b67a264211d833c69.
//
// Solidity: event WithdrawnFromAave(address indexed caller, address indexed asset, uint256 amount, address indexed to)
func (_Aave *AaveFilterer) WatchWithdrawnFromAave(opts *bind.WatchOpts, sink chan<- *AaveWithdrawnFromAave, caller []common.Address, asset []common.Address, to []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "WithdrawnFromAave", callerRule, assetRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveWithdrawnFromAave)
				if err := _Aave.contract.UnpackLog(event, "WithdrawnFromAave", log); err != nil {
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

// ParseWithdrawnFromAave is a log parse operation binding the contract event 0xf4233c2492c8fe129470230aba79a18b71f74c3b98de673b67a264211d833c69.
//
// Solidity: event WithdrawnFromAave(address indexed caller, address indexed asset, uint256 amount, address indexed to)
func (_Aave *AaveFilterer) ParseWithdrawnFromAave(log types.Log) (*AaveWithdrawnFromAave, error) {
	event := new(AaveWithdrawnFromAave)
	if err := _Aave.contract.UnpackLog(event, "WithdrawnFromAave", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
