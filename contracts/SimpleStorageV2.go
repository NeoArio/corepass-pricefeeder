// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// SimpleeventMetaData contains all meta data concerning the Simpleevent contract.
var SimpleeventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SimpleeventABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleeventMetaData.ABI instead.
var SimpleeventABI = SimpleeventMetaData.ABI

// Simpleevent is an auto generated Go binding around an Ethereum contract.
type Simpleevent struct {
	SimpleeventCaller     // Read-only binding to the contract
	SimpleeventTransactor // Write-only binding to the contract
	SimpleeventFilterer   // Log filterer for contract events
}

// SimpleeventCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleeventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleeventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleeventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleeventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleeventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleeventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleeventSession struct {
	Contract     *Simpleevent      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleeventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleeventCallerSession struct {
	Contract *SimpleeventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SimpleeventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleeventTransactorSession struct {
	Contract     *SimpleeventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SimpleeventRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleeventRaw struct {
	Contract *Simpleevent // Generic contract binding to access the raw methods on
}

// SimpleeventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleeventCallerRaw struct {
	Contract *SimpleeventCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleeventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleeventTransactorRaw struct {
	Contract *SimpleeventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleevent creates a new instance of Simpleevent, bound to a specific deployed contract.
func NewSimpleevent(address common.Address, backend bind.ContractBackend) (*Simpleevent, error) {
	contract, err := bindSimpleevent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simpleevent{SimpleeventCaller: SimpleeventCaller{contract: contract}, SimpleeventTransactor: SimpleeventTransactor{contract: contract}, SimpleeventFilterer: SimpleeventFilterer{contract: contract}}, nil
}

// NewSimpleeventCaller creates a new read-only instance of Simpleevent, bound to a specific deployed contract.
func NewSimpleeventCaller(address common.Address, caller bind.ContractCaller) (*SimpleeventCaller, error) {
	contract, err := bindSimpleevent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleeventCaller{contract: contract}, nil
}

// NewSimpleeventTransactor creates a new write-only instance of Simpleevent, bound to a specific deployed contract.
func NewSimpleeventTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleeventTransactor, error) {
	contract, err := bindSimpleevent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleeventTransactor{contract: contract}, nil
}

// NewSimpleeventFilterer creates a new log filterer instance of Simpleevent, bound to a specific deployed contract.
func NewSimpleeventFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleeventFilterer, error) {
	contract, err := bindSimpleevent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleeventFilterer{contract: contract}, nil
}

// bindSimpleevent binds a generic wrapper to an already deployed contract.
func bindSimpleevent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleeventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simpleevent *SimpleeventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simpleevent.Contract.SimpleeventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simpleevent *SimpleeventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpleevent.Contract.SimpleeventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simpleevent *SimpleeventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simpleevent.Contract.SimpleeventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simpleevent *SimpleeventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simpleevent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simpleevent *SimpleeventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpleevent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simpleevent *SimpleeventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simpleevent.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Simpleevent *SimpleeventCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simpleevent.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Simpleevent *SimpleeventSession) Get() (*big.Int, error) {
	return _Simpleevent.Contract.Get(&_Simpleevent.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Simpleevent *SimpleeventCallerSession) Get() (*big.Int, error) {
	return _Simpleevent.Contract.Get(&_Simpleevent.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Simpleevent *SimpleeventTransactor) Set(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _Simpleevent.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Simpleevent *SimpleeventSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Simpleevent.Contract.Set(&_Simpleevent.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Simpleevent *SimpleeventTransactorSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Simpleevent.Contract.Set(&_Simpleevent.TransactOpts, x)
}

// SimpleeventItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Simpleevent contract.
type SimpleeventItemSetIterator struct {
	Event *SimpleeventItemSet // Event containing the contract specifics and raw log

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
func (it *SimpleeventItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleeventItemSet)
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
		it.Event = new(SimpleeventItemSet)
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
func (it *SimpleeventItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleeventItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleeventItemSet represents a ItemSet event raised by the Simpleevent contract.
type SimpleeventItemSet struct {
	Key   string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Simpleevent *SimpleeventFilterer) FilterItemSet(opts *bind.FilterOpts) (*SimpleeventItemSetIterator, error) {

	logs, sub, err := _Simpleevent.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &SimpleeventItemSetIterator{contract: _Simpleevent.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Simpleevent *SimpleeventFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *SimpleeventItemSet) (event.Subscription, error) {

	logs, sub, err := _Simpleevent.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleeventItemSet)
				if err := _Simpleevent.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Simpleevent *SimpleeventFilterer) ParseItemSet(log types.Log) (*SimpleeventItemSet, error) {
	event := new(SimpleeventItemSet)
	if err := _Simpleevent.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
