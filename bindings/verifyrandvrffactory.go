// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// VerifyRandVRFFactoryMetaData contains all meta data concerning the VerifyRandVRFFactory contract.
var VerifyRandVRFFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createProxy\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifyRandAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ProxyCreated\",\"inputs\":[{\"name\":\"mintProxyAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x6080806040523461001657610198908161001b8239f35b5f80fdfe60806040908082526004361015610014575f80fd5b5f91823560e01c6325b5672714610029575f80fd5b3461015e578060031936011261015e576004356001600160a01b03808216820361015a57602435918183168093036101565780763d602d80600a3d3981f3363d3d373d3d3d363d7300000062ffffff6e5af43d82803e903d91602b57fd5bf39360881c1617875260781b176020526037600986f0169283156101475750823b156101435781519063485cc95560e01b82523360048301526024820152838160448183875af180156101395761010a575b602083837efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349838251848152a151908152f35b67ffffffffffffffff8111610125576020935081525f6100d9565b634e487b7160e01b84526041600452602484fd5b82513d86823e3d90fd5b8380fd5b63b06ebf3d60e01b8152600490fd5b8580fd5b8480fd5b8280fdfea264697066735822122024b6d8f6b176082bf5d95a5dd7021f57ae2ef7b7acb6c015613d72bc8d7a692f64736f6c63430008140033",
}

// VerifyRandVRFFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyRandVRFFactoryMetaData.ABI instead.
var VerifyRandVRFFactoryABI = VerifyRandVRFFactoryMetaData.ABI

// VerifyRandVRFFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifyRandVRFFactoryMetaData.Bin instead.
var VerifyRandVRFFactoryBin = VerifyRandVRFFactoryMetaData.Bin

// DeployVerifyRandVRFFactory deploys a new Ethereum contract, binding an instance of VerifyRandVRFFactory to it.
func DeployVerifyRandVRFFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerifyRandVRFFactory, error) {
	parsed, err := VerifyRandVRFFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifyRandVRFFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifyRandVRFFactory{VerifyRandVRFFactoryCaller: VerifyRandVRFFactoryCaller{contract: contract}, VerifyRandVRFFactoryTransactor: VerifyRandVRFFactoryTransactor{contract: contract}, VerifyRandVRFFactoryFilterer: VerifyRandVRFFactoryFilterer{contract: contract}}, nil
}

// VerifyRandVRFFactory is an auto generated Go binding around an Ethereum contract.
type VerifyRandVRFFactory struct {
	VerifyRandVRFFactoryCaller     // Read-only binding to the contract
	VerifyRandVRFFactoryTransactor // Write-only binding to the contract
	VerifyRandVRFFactoryFilterer   // Log filterer for contract events
}

// VerifyRandVRFFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyRandVRFFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyRandVRFFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyRandVRFFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyRandVRFFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyRandVRFFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyRandVRFFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifyRandVRFFactorySession struct {
	Contract     *VerifyRandVRFFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VerifyRandVRFFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyRandVRFFactoryCallerSession struct {
	Contract *VerifyRandVRFFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// VerifyRandVRFFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyRandVRFFactoryTransactorSession struct {
	Contract     *VerifyRandVRFFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// VerifyRandVRFFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyRandVRFFactoryRaw struct {
	Contract *VerifyRandVRFFactory // Generic contract binding to access the raw methods on
}

// VerifyRandVRFFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyRandVRFFactoryCallerRaw struct {
	Contract *VerifyRandVRFFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyRandVRFFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyRandVRFFactoryTransactorRaw struct {
	Contract *VerifyRandVRFFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifyRandVRFFactory creates a new instance of VerifyRandVRFFactory, bound to a specific deployed contract.
func NewVerifyRandVRFFactory(address common.Address, backend bind.ContractBackend) (*VerifyRandVRFFactory, error) {
	contract, err := bindVerifyRandVRFFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFFactory{VerifyRandVRFFactoryCaller: VerifyRandVRFFactoryCaller{contract: contract}, VerifyRandVRFFactoryTransactor: VerifyRandVRFFactoryTransactor{contract: contract}, VerifyRandVRFFactoryFilterer: VerifyRandVRFFactoryFilterer{contract: contract}}, nil
}

// NewVerifyRandVRFFactoryCaller creates a new read-only instance of VerifyRandVRFFactory, bound to a specific deployed contract.
func NewVerifyRandVRFFactoryCaller(address common.Address, caller bind.ContractCaller) (*VerifyRandVRFFactoryCaller, error) {
	contract, err := bindVerifyRandVRFFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFFactoryCaller{contract: contract}, nil
}

// NewVerifyRandVRFFactoryTransactor creates a new write-only instance of VerifyRandVRFFactory, bound to a specific deployed contract.
func NewVerifyRandVRFFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyRandVRFFactoryTransactor, error) {
	contract, err := bindVerifyRandVRFFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFFactoryTransactor{contract: contract}, nil
}

// NewVerifyRandVRFFactoryFilterer creates a new log filterer instance of VerifyRandVRFFactory, bound to a specific deployed contract.
func NewVerifyRandVRFFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyRandVRFFactoryFilterer, error) {
	contract, err := bindVerifyRandVRFFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFFactoryFilterer{contract: contract}, nil
}

// bindVerifyRandVRFFactory binds a generic wrapper to an already deployed contract.
func bindVerifyRandVRFFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyRandVRFFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyRandVRFFactory.Contract.VerifyRandVRFFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyRandVRFFactory.Contract.VerifyRandVRFFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyRandVRFFactory.Contract.VerifyRandVRFFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyRandVRFFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyRandVRFFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyRandVRFFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(address implementation, address verifyRandAddress) returns(address)
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryTransactor) CreateProxy(opts *bind.TransactOpts, implementation common.Address, verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRFFactory.contract.Transact(opts, "createProxy", implementation, verifyRandAddress)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(address implementation, address verifyRandAddress) returns(address)
func (_VerifyRandVRFFactory *VerifyRandVRFFactorySession) CreateProxy(implementation common.Address, verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRFFactory.Contract.CreateProxy(&_VerifyRandVRFFactory.TransactOpts, implementation, verifyRandAddress)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(address implementation, address verifyRandAddress) returns(address)
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryTransactorSession) CreateProxy(implementation common.Address, verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRFFactory.Contract.CreateProxy(&_VerifyRandVRFFactory.TransactOpts, implementation, verifyRandAddress)
}

// VerifyRandVRFFactoryProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the VerifyRandVRFFactory contract.
type VerifyRandVRFFactoryProxyCreatedIterator struct {
	Event *VerifyRandVRFFactoryProxyCreated // Event containing the contract specifics and raw log

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
func (it *VerifyRandVRFFactoryProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyRandVRFFactoryProxyCreated)
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
		it.Event = new(VerifyRandVRFFactoryProxyCreated)
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
func (it *VerifyRandVRFFactoryProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyRandVRFFactoryProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyRandVRFFactoryProxyCreated represents a ProxyCreated event raised by the VerifyRandVRFFactory contract.
type VerifyRandVRFFactoryProxyCreated struct {
	MintProxyAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address mintProxyAddress)
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryFilterer) FilterProxyCreated(opts *bind.FilterOpts) (*VerifyRandVRFFactoryProxyCreatedIterator, error) {

	logs, sub, err := _VerifyRandVRFFactory.contract.FilterLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFFactoryProxyCreatedIterator{contract: _VerifyRandVRFFactory.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address mintProxyAddress)
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *VerifyRandVRFFactoryProxyCreated) (event.Subscription, error) {

	logs, sub, err := _VerifyRandVRFFactory.contract.WatchLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyRandVRFFactoryProxyCreated)
				if err := _VerifyRandVRFFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address mintProxyAddress)
func (_VerifyRandVRFFactory *VerifyRandVRFFactoryFilterer) ParseProxyCreated(log types.Log) (*VerifyRandVRFFactoryProxyCreated, error) {
	event := new(VerifyRandVRFFactoryProxyCreated)
	if err := _VerifyRandVRFFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
