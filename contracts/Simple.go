// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	core "github.com/core-coin/go-core"
	"github.com/core-coin/go-core/accounts/abi"
	"github.com/core-coin/go-core/accounts/abi/bind"
	"github.com/core-coin/go-core/common"
	"github.com/core-coin/go-core/core/types"
	"github.com/core-coin/go-core/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = core.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SimpleABI is the input ABI used to generate the binding from.
const SimpleABI = "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Simple is an auto generated Go binding around an Core contract.
type Simple struct {
	SimpleCaller     // Read-only binding to the contract
	SimpleTransactor // Write-only binding to the contract
	SimpleFilterer   // Log filterer for contract events
}

// SimpleCaller is an auto generated read-only Go binding around an Core contract.
type SimpleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransactor is an auto generated write-only Go binding around an Core contract.
type SimpleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFilterer is an auto generated log filtering Go binding around an Core contract events.
type SimpleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type SimpleSession struct {
	Contract     *Simple           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type SimpleCallerSession struct {
	Contract *SimpleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SimpleTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type SimpleTransactorSession struct {
	Contract     *SimpleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRaw is an auto generated low-level Go binding around an Core contract.
type SimpleRaw struct {
	Contract *Simple // Generic contract binding to access the raw methods on
}

// SimpleCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type SimpleCallerRaw struct {
	Contract *SimpleCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type SimpleTransactorRaw struct {
	Contract *SimpleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimple creates a new instance of Simple, bound to a specific deployed contract.
func NewSimple(address common.Address, backend bind.ContractBackend) (*Simple, error) {
	contract, err := bindSimple(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// NewSimpleCaller creates a new read-only instance of Simple, bound to a specific deployed contract.
func NewSimpleCaller(address common.Address, caller bind.ContractCaller) (*SimpleCaller, error) {
	contract, err := bindSimple(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCaller{contract: contract}, nil
}

// NewSimpleTransactor creates a new write-only instance of Simple, bound to a specific deployed contract.
func NewSimpleTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTransactor, error) {
	contract, err := bindSimple(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransactor{contract: contract}, nil
}

// NewSimpleFilterer creates a new log filterer instance of Simple, bound to a specific deployed contract.
func NewSimpleFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleFilterer, error) {
	contract, err := bindSimple(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleFilterer{contract: contract}, nil
}

// bindSimple binds a generic wrapper to an already deployed contract.
func bindSimple(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.SimpleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6ebb8d83.
//
// Solidity: function get() constant returns(uint256)
func (_Simple *SimpleCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Simple.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6ebb8d83.
//
// Solidity: function get() constant returns(uint256)
func (_Simple *SimpleSession) Get() (*big.Int, error) {
	return _Simple.Contract.Get(&_Simple.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6ebb8d83.
//
// Solidity: function get() constant returns(uint256)
func (_Simple *SimpleCallerSession) Get() (*big.Int, error) {
	return _Simple.Contract.Get(&_Simple.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x3af6db04.
//
// Solidity: function set(uint256 x) returns()
func (_Simple *SimpleTransactor) Set(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x3af6db04.
//
// Solidity: function set(uint256 x) returns()
func (_Simple *SimpleSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.Set(&_Simple.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x3af6db04.
//
// Solidity: function set(uint256 x) returns()
func (_Simple *SimpleTransactorSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.Set(&_Simple.TransactOpts, x)
}
