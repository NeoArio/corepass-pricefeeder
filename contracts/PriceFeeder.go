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

// PriceFeederABI is the input ABI used to generate the binding from.
const PriceFeederABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"addPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"authorizeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"authorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"deauthorizeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregatedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PriceFeeder is an auto generated Go binding around an Core contract.
type PriceFeeder struct {
	PriceFeederCaller     // Read-only binding to the contract
	PriceFeederTransactor // Write-only binding to the contract
	PriceFeederFilterer   // Log filterer for contract events
}

// PriceFeederCaller is an auto generated read-only Go binding around an Core contract.
type PriceFeederCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederTransactor is an auto generated write-only Go binding around an Core contract.
type PriceFeederTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederFilterer is an auto generated log filtering Go binding around an Core contract events.
type PriceFeederFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type PriceFeederSession struct {
	Contract     *PriceFeeder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeederCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type PriceFeederCallerSession struct {
	Contract *PriceFeederCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceFeederTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type PriceFeederTransactorSession struct {
	Contract     *PriceFeederTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceFeederRaw is an auto generated low-level Go binding around an Core contract.
type PriceFeederRaw struct {
	Contract *PriceFeeder // Generic contract binding to access the raw methods on
}

// PriceFeederCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type PriceFeederCallerRaw struct {
	Contract *PriceFeederCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeederTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type PriceFeederTransactorRaw struct {
	Contract *PriceFeederTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeeder creates a new instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeeder(address common.Address, backend bind.ContractBackend) (*PriceFeeder, error) {
	contract, err := bindPriceFeeder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeeder{PriceFeederCaller: PriceFeederCaller{contract: contract}, PriceFeederTransactor: PriceFeederTransactor{contract: contract}, PriceFeederFilterer: PriceFeederFilterer{contract: contract}}, nil
}

// NewPriceFeederCaller creates a new read-only instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederCaller(address common.Address, caller bind.ContractCaller) (*PriceFeederCaller, error) {
	contract, err := bindPriceFeeder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeederCaller{contract: contract}, nil
}

// NewPriceFeederTransactor creates a new write-only instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeederTransactor, error) {
	contract, err := bindPriceFeeder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeederTransactor{contract: contract}, nil
}

// NewPriceFeederFilterer creates a new log filterer instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeederFilterer, error) {
	contract, err := bindPriceFeeder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeederFilterer{contract: contract}, nil
}

// bindPriceFeeder binds a generic wrapper to an already deployed contract.
func bindPriceFeeder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeederABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeeder *PriceFeederRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeeder.Contract.PriceFeederCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeeder *PriceFeederRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeeder.Contract.PriceFeederTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeeder *PriceFeederRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeeder.Contract.PriceFeederTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeeder *PriceFeederCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeeder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeeder *PriceFeederTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeeder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeeder *PriceFeederTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeeder.Contract.contract.Transact(opts, method, params...)
}

// Authorized is a free data retrieval call binding the contract method 0x9efda55c.
//
// Solidity: function authorized(address _nodeAddress) constant returns(bool)
func (_PriceFeeder *PriceFeederCaller) Authorized(opts *bind.CallOpts, _nodeAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PriceFeeder.contract.Call(opts, out, "authorized", _nodeAddress)
	return *ret0, err
}

// Authorized is a free data retrieval call binding the contract method 0x9efda55c.
//
// Solidity: function authorized(address _nodeAddress) constant returns(bool)
func (_PriceFeeder *PriceFeederSession) Authorized(_nodeAddress common.Address) (bool, error) {
	return _PriceFeeder.Contract.Authorized(&_PriceFeeder.CallOpts, _nodeAddress)
}

// Authorized is a free data retrieval call binding the contract method 0x9efda55c.
//
// Solidity: function authorized(address _nodeAddress) constant returns(bool)
func (_PriceFeeder *PriceFeederCallerSession) Authorized(_nodeAddress common.Address) (bool, error) {
	return _PriceFeeder.Contract.Authorized(&_PriceFeeder.CallOpts, _nodeAddress)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() constant returns(uint256)
func (_PriceFeeder *PriceFeederCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceFeeder.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() constant returns(uint256)
func (_PriceFeeder *PriceFeederSession) Decimals() (*big.Int, error) {
	return _PriceFeeder.Contract.Decimals(&_PriceFeeder.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() constant returns(uint256)
func (_PriceFeeder *PriceFeederCallerSession) Decimals() (*big.Int, error) {
	return _PriceFeeder.Contract.Decimals(&_PriceFeeder.CallOpts)
}

// GetAggregatedPrice is a free data retrieval call binding the contract method 0xd9c1c1c9.
//
// Solidity: function getAggregatedPrice() constant returns(uint256, uint256, uint32)
func (_PriceFeeder *PriceFeederCaller) GetAggregatedPrice(opts *bind.CallOpts) (*big.Int, *big.Int, uint32, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(uint32)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _PriceFeeder.contract.Call(opts, out, "getAggregatedPrice")
	return *ret0, *ret1, *ret2, err
}

// GetAggregatedPrice is a free data retrieval call binding the contract method 0xd9c1c1c9.
//
// Solidity: function getAggregatedPrice() constant returns(uint256, uint256, uint32)
func (_PriceFeeder *PriceFeederSession) GetAggregatedPrice() (*big.Int, *big.Int, uint32, error) {
	return _PriceFeeder.Contract.GetAggregatedPrice(&_PriceFeeder.CallOpts)
}

// GetAggregatedPrice is a free data retrieval call binding the contract method 0xd9c1c1c9.
//
// Solidity: function getAggregatedPrice() constant returns(uint256, uint256, uint32)
func (_PriceFeeder *PriceFeederCallerSession) GetAggregatedPrice() (*big.Int, *big.Int, uint32, error) {
	return _PriceFeeder.Contract.GetAggregatedPrice(&_PriceFeeder.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x677dcf04.
//
// Solidity: function getLatestPrice() constant returns(uint256, uint256)
func (_PriceFeeder *PriceFeederCaller) GetLatestPrice(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _PriceFeeder.contract.Call(opts, out, "getLatestPrice")
	return *ret0, *ret1, err
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x677dcf04.
//
// Solidity: function getLatestPrice() constant returns(uint256, uint256)
func (_PriceFeeder *PriceFeederSession) GetLatestPrice() (*big.Int, *big.Int, error) {
	return _PriceFeeder.Contract.GetLatestPrice(&_PriceFeeder.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x677dcf04.
//
// Solidity: function getLatestPrice() constant returns(uint256, uint256)
func (_PriceFeeder *PriceFeederCallerSession) GetLatestPrice() (*big.Int, *big.Int, error) {
	return _PriceFeeder.Contract.GetLatestPrice(&_PriceFeeder.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() constant returns(string)
func (_PriceFeeder *PriceFeederCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PriceFeeder.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() constant returns(string)
func (_PriceFeeder *PriceFeederSession) Name() (string, error) {
	return _PriceFeeder.Contract.Name(&_PriceFeeder.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() constant returns(string)
func (_PriceFeeder *PriceFeederCallerSession) Name() (string, error) {
	return _PriceFeeder.Contract.Name(&_PriceFeeder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xbe0e67a3.
//
// Solidity: function owner() constant returns(address)
func (_PriceFeeder *PriceFeederCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceFeeder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0xbe0e67a3.
//
// Solidity: function owner() constant returns(address)
func (_PriceFeeder *PriceFeederSession) Owner() (common.Address, error) {
	return _PriceFeeder.Contract.Owner(&_PriceFeeder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xbe0e67a3.
//
// Solidity: function owner() constant returns(address)
func (_PriceFeeder *PriceFeederCallerSession) Owner() (common.Address, error) {
	return _PriceFeeder.Contract.Owner(&_PriceFeeder.CallOpts)
}

// AddPrice is a paid mutator transaction binding the contract method 0xca725b7e.
//
// Solidity: function addPrice(uint256 _price) returns()
func (_PriceFeeder *PriceFeederTransactor) AddPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _PriceFeeder.contract.Transact(opts, "addPrice", _price)
}

// AddPrice is a paid mutator transaction binding the contract method 0xca725b7e.
//
// Solidity: function addPrice(uint256 _price) returns()
func (_PriceFeeder *PriceFeederSession) AddPrice(_price *big.Int) (*types.Transaction, error) {
	return _PriceFeeder.Contract.AddPrice(&_PriceFeeder.TransactOpts, _price)
}

// AddPrice is a paid mutator transaction binding the contract method 0xca725b7e.
//
// Solidity: function addPrice(uint256 _price) returns()
func (_PriceFeeder *PriceFeederTransactorSession) AddPrice(_price *big.Int) (*types.Transaction, error) {
	return _PriceFeeder.Contract.AddPrice(&_PriceFeeder.TransactOpts, _price)
}

// AuthorizeNode is a paid mutator transaction binding the contract method 0x84b54833.
//
// Solidity: function authorizeNode(address _nodeAddress) returns()
func (_PriceFeeder *PriceFeederTransactor) AuthorizeNode(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeder.contract.Transact(opts, "authorizeNode", _nodeAddress)
}

// AuthorizeNode is a paid mutator transaction binding the contract method 0x84b54833.
//
// Solidity: function authorizeNode(address _nodeAddress) returns()
func (_PriceFeeder *PriceFeederSession) AuthorizeNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeder.Contract.AuthorizeNode(&_PriceFeeder.TransactOpts, _nodeAddress)
}

// AuthorizeNode is a paid mutator transaction binding the contract method 0x84b54833.
//
// Solidity: function authorizeNode(address _nodeAddress) returns()
func (_PriceFeeder *PriceFeederTransactorSession) AuthorizeNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeder.Contract.AuthorizeNode(&_PriceFeeder.TransactOpts, _nodeAddress)
}

// DeauthorizeNode is a paid mutator transaction binding the contract method 0xff5cf248.
//
// Solidity: function deauthorizeNode(address _nodeAddress) returns()
func (_PriceFeeder *PriceFeederTransactor) DeauthorizeNode(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeder.contract.Transact(opts, "deauthorizeNode", _nodeAddress)
}

// DeauthorizeNode is a paid mutator transaction binding the contract method 0xff5cf248.
//
// Solidity: function deauthorizeNode(address _nodeAddress) returns()
func (_PriceFeeder *PriceFeederSession) DeauthorizeNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeder.Contract.DeauthorizeNode(&_PriceFeeder.TransactOpts, _nodeAddress)
}

// DeauthorizeNode is a paid mutator transaction binding the contract method 0xff5cf248.
//
// Solidity: function deauthorizeNode(address _nodeAddress) returns()
func (_PriceFeeder *PriceFeederTransactorSession) DeauthorizeNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeder.Contract.DeauthorizeNode(&_PriceFeeder.TransactOpts, _nodeAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x5e0827e9.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeeder *PriceFeederTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeeder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x5e0827e9.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeeder *PriceFeederSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeeder.Contract.RenounceOwnership(&_PriceFeeder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x5e0827e9.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeeder *PriceFeederTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeeder.Contract.RenounceOwnership(&_PriceFeeder.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xaae7857b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeeder *PriceFeederTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeeder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xaae7857b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeeder *PriceFeederSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeeder.Contract.TransferOwnership(&_PriceFeeder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xaae7857b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeeder *PriceFeederTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeeder.Contract.TransferOwnership(&_PriceFeeder.TransactOpts, newOwner)
}

// PriceFeederOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceFeeder contract.
type PriceFeederOwnershipTransferredIterator struct {
	Event *PriceFeederOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeederOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederOwnershipTransferred)
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
		it.Event = new(PriceFeederOwnershipTransferred)
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
func (it *PriceFeederOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederOwnershipTransferred represents a OwnershipTransferred event raised by the PriceFeeder contract.
type PriceFeederOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x6706336cd2535bfac8a54fc906c5d5bf89c8a10509b1188ec4664ee2acadd6b4.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeeder *PriceFeederFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceFeederOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeeder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeederOwnershipTransferredIterator{contract: _PriceFeeder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x6706336cd2535bfac8a54fc906c5d5bf89c8a10509b1188ec4664ee2acadd6b4.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeeder *PriceFeederFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceFeederOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeeder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederOwnershipTransferred)
				if err := _PriceFeeder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x6706336cd2535bfac8a54fc906c5d5bf89c8a10509b1188ec4664ee2acadd6b4.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeeder *PriceFeederFilterer) ParseOwnershipTransferred(log types.Log) (*PriceFeederOwnershipTransferred, error) {
	event := new(PriceFeederOwnershipTransferred)
	if err := _PriceFeeder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
