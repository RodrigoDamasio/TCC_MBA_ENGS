// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// QualityTestContractSampleResult is an auto generated low-level Go binding around an user-defined struct.
type QualityTestContractSampleResult struct {
	Value  *big.Int
	Result string
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"finalResult\",\"type\":\"string\"}],\"name\":\"QualityTestBatchCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"minLimit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"maxLimit\",\"type\":\"int256\"},{\"internalType\":\"int256[]\",\"name\":\"values\",\"type\":\"int256[]\"},{\"internalType\":\"string\",\"name\":\"productionOrder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"testDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"testDescription\",\"type\":\"string\"}],\"name\":\"performBatchQualityTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"queryQualityTestBatch\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"internalType\":\"structQualityTestContract.SampleResult[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// QueryQualityTestBatch is a free data retrieval call binding the contract method 0xc11c4e23.
//
// Solidity: function queryQualityTestBatch(string id) view returns(string, int256, int256, int256[], string, string, string, (int256,string)[], string)
func (_Contract *ContractCaller) QueryQualityTestBatch(opts *bind.CallOpts, id string) (string, *big.Int, *big.Int, []*big.Int, string, string, string, []QualityTestContractSampleResult, string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "queryQualityTestBatch", id)

	if err != nil {
		return *new(string), *new(*big.Int), *new(*big.Int), *new([]*big.Int), *new(string), *new(string), *new(string), *new([]QualityTestContractSampleResult), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(string)).(*string)
	out6 := *abi.ConvertType(out[6], new(string)).(*string)
	out7 := *abi.ConvertType(out[7], new([]QualityTestContractSampleResult)).(*[]QualityTestContractSampleResult)
	out8 := *abi.ConvertType(out[8], new(string)).(*string)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// QueryQualityTestBatch is a free data retrieval call binding the contract method 0xc11c4e23.
//
// Solidity: function queryQualityTestBatch(string id) view returns(string, int256, int256, int256[], string, string, string, (int256,string)[], string)
func (_Contract *ContractSession) QueryQualityTestBatch(id string) (string, *big.Int, *big.Int, []*big.Int, string, string, string, []QualityTestContractSampleResult, string, error) {
	return _Contract.Contract.QueryQualityTestBatch(&_Contract.CallOpts, id)
}

// QueryQualityTestBatch is a free data retrieval call binding the contract method 0xc11c4e23.
//
// Solidity: function queryQualityTestBatch(string id) view returns(string, int256, int256, int256[], string, string, string, (int256,string)[], string)
func (_Contract *ContractCallerSession) QueryQualityTestBatch(id string) (string, *big.Int, *big.Int, []*big.Int, string, string, string, []QualityTestContractSampleResult, string, error) {
	return _Contract.Contract.QueryQualityTestBatch(&_Contract.CallOpts, id)
}

// PerformBatchQualityTest is a paid mutator transaction binding the contract method 0xc6b20dd8.
//
// Solidity: function performBatchQualityTest(string id, int256 minLimit, int256 maxLimit, int256[] values, string productionOrder, string testDate, string testDescription) returns()
func (_Contract *ContractTransactor) PerformBatchQualityTest(opts *bind.TransactOpts, id string, minLimit *big.Int, maxLimit *big.Int, values []*big.Int, productionOrder string, testDate string, testDescription string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "performBatchQualityTest", id, minLimit, maxLimit, values, productionOrder, testDate, testDescription)
}

// PerformBatchQualityTest is a paid mutator transaction binding the contract method 0xc6b20dd8.
//
// Solidity: function performBatchQualityTest(string id, int256 minLimit, int256 maxLimit, int256[] values, string productionOrder, string testDate, string testDescription) returns()
func (_Contract *ContractSession) PerformBatchQualityTest(id string, minLimit *big.Int, maxLimit *big.Int, values []*big.Int, productionOrder string, testDate string, testDescription string) (*types.Transaction, error) {
	return _Contract.Contract.PerformBatchQualityTest(&_Contract.TransactOpts, id, minLimit, maxLimit, values, productionOrder, testDate, testDescription)
}

// PerformBatchQualityTest is a paid mutator transaction binding the contract method 0xc6b20dd8.
//
// Solidity: function performBatchQualityTest(string id, int256 minLimit, int256 maxLimit, int256[] values, string productionOrder, string testDate, string testDescription) returns()
func (_Contract *ContractTransactorSession) PerformBatchQualityTest(id string, minLimit *big.Int, maxLimit *big.Int, values []*big.Int, productionOrder string, testDate string, testDescription string) (*types.Transaction, error) {
	return _Contract.Contract.PerformBatchQualityTest(&_Contract.TransactOpts, id, minLimit, maxLimit, values, productionOrder, testDate, testDescription)
}

// ContractQualityTestBatchCreatedIterator is returned from FilterQualityTestBatchCreated and is used to iterate over the raw logs and unpacked data for QualityTestBatchCreated events raised by the Contract contract.
type ContractQualityTestBatchCreatedIterator struct {
	Event *ContractQualityTestBatchCreated // Event containing the contract specifics and raw log

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
func (it *ContractQualityTestBatchCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractQualityTestBatchCreated)
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
		it.Event = new(ContractQualityTestBatchCreated)
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
func (it *ContractQualityTestBatchCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractQualityTestBatchCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractQualityTestBatchCreated represents a QualityTestBatchCreated event raised by the Contract contract.
type ContractQualityTestBatchCreated struct {
	Id          string
	FinalResult string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterQualityTestBatchCreated is a free log retrieval operation binding the contract event 0x73f13c3c8e49d923017340abe0fbaeb0e7dca4e9c565ae6417b193cb02bb11f9.
//
// Solidity: event QualityTestBatchCreated(string id, string finalResult)
func (_Contract *ContractFilterer) FilterQualityTestBatchCreated(opts *bind.FilterOpts) (*ContractQualityTestBatchCreatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "QualityTestBatchCreated")
	if err != nil {
		return nil, err
	}
	return &ContractQualityTestBatchCreatedIterator{contract: _Contract.contract, event: "QualityTestBatchCreated", logs: logs, sub: sub}, nil
}

// WatchQualityTestBatchCreated is a free log subscription operation binding the contract event 0x73f13c3c8e49d923017340abe0fbaeb0e7dca4e9c565ae6417b193cb02bb11f9.
//
// Solidity: event QualityTestBatchCreated(string id, string finalResult)
func (_Contract *ContractFilterer) WatchQualityTestBatchCreated(opts *bind.WatchOpts, sink chan<- *ContractQualityTestBatchCreated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "QualityTestBatchCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractQualityTestBatchCreated)
				if err := _Contract.contract.UnpackLog(event, "QualityTestBatchCreated", log); err != nil {
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

// ParseQualityTestBatchCreated is a log parse operation binding the contract event 0x73f13c3c8e49d923017340abe0fbaeb0e7dca4e9c565ae6417b193cb02bb11f9.
//
// Solidity: event QualityTestBatchCreated(string id, string finalResult)
func (_Contract *ContractFilterer) ParseQualityTestBatchCreated(log types.Log) (*ContractQualityTestBatchCreated, error) {
	event := new(ContractQualityTestBatchCreated)
	if err := _Contract.contract.UnpackLog(event, "QualityTestBatchCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
