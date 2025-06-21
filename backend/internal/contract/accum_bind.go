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

// AccumMetaData contains all meta data concerning the Accum contract.
var AccumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Adding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"Exceed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"checkBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendEthers\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610903380380610903833981810160405281019061003291906101e2565b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a55760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161009c919061021e565b60405180910390fd5b6100b4816100bb60201b60201c565b5050610239565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101af82610184565b9050919050565b6101bf816101a4565b81146101ca57600080fd5b50565b6000815190506101dc816101b6565b92915050565b6000602082840312156101f8576101f761017f565b5b6000610206848285016101cd565b91505092915050565b610218816101a4565b82525050565b6000602082019050610233600083018461020f565b92915050565b6106bb806102486000396000f3fe6080604052600436106100595760003560e01c80630783cda71461006f578063715018a6146100795780637362377b146100905780638da5cb5b146100a7578063c71daccb146100d2578063f2fde38b146100fd5761006a565b3661006a576100683334610126565b005b600080fd5b6100776101aa565b005b34801561008557600080fd5b5061008e6101b6565b005b34801561009c57600080fd5b506100a56101ca565b005b3480156100b357600080fd5b506100bc61026d565b6040516100c991906104b8565b60405180910390f35b3480156100de57600080fd5b506100e7610296565b6040516100f491906104ec565b60405180910390f35b34801561010957600080fd5b50610124600480360381019061011f9190610538565b61029e565b005b7ff5319c94a6e3d822f1909045eea12c668e7851fa33a70d99000329ceabd51c508282604051610157929190610565565b60405180910390a167b469471f8014000047106101a6577fb0428161e04cd6a37a41d41efbaba21a3404b34b15594f16e373792db46de6ed4760405161019d91906105eb565b60405180910390a15b5050565b6101b43334610126565b565b6101be610324565b6101c860006103ab565b565b6101d2610324565b67b469471f8014000047101561021d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021490610665565b60405180910390fd5b61022561026d565b73ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561026a573d6000803e3d6000fd5b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600047905090565b6102a6610324565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036103185760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161030f91906104b8565b60405180910390fd5b610321816103ab565b50565b61032c61046f565b73ffffffffffffffffffffffffffffffffffffffff1661034a61026d565b73ffffffffffffffffffffffffffffffffffffffff16146103a95761036d61046f565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016103a091906104b8565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104a282610477565b9050919050565b6104b281610497565b82525050565b60006020820190506104cd60008301846104a9565b92915050565b6000819050919050565b6104e6816104d3565b82525050565b600060208201905061050160008301846104dd565b92915050565b600080fd5b61051581610497565b811461052057600080fd5b50565b6000813590506105328161050c565b92915050565b60006020828403121561054e5761054d610507565b5b600061055c84828501610523565b91505092915050565b600060408201905061057a60008301856104a9565b61058760208301846104dd565b9392505050565b600082825260208201905092915050565b7f506c616e6b20646f6e6521000000000000000000000000000000000000000000600082015250565b60006105d5600b8361058e565b91506105e08261059f565b602082019050919050565b600060408201905061060060008301846104dd565b8181036020830152610611816105c8565b905092915050565b7f42616c616e6365206c6f776572207468616e20706c616e6b2100000000000000600082015250565b600061064f60198361058e565b915061065a82610619565b602082019050919050565b6000602082019050818103600083015261067e81610642565b905091905056fea264697066735822122059495910b70d402dd0e0353386ff33ba0272abb2ef07023933fecd945e59374464736f6c634300081c0033",
}

// AccumABI is the input ABI used to generate the binding from.
// Deprecated: Use AccumMetaData.ABI instead.
var AccumABI = AccumMetaData.ABI

// AccumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccumMetaData.Bin instead.
var AccumBin = AccumMetaData.Bin

// DeployAccum deploys a new Ethereum contract, binding an instance of Accum to it.
func DeployAccum(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address) (common.Address, *types.Transaction, *Accum, error) {
	parsed, err := AccumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccumBin), backend, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accum{AccumCaller: AccumCaller{contract: contract}, AccumTransactor: AccumTransactor{contract: contract}, AccumFilterer: AccumFilterer{contract: contract}}, nil
}

// Accum is an auto generated Go binding around an Ethereum contract.
type Accum struct {
	AccumCaller     // Read-only binding to the contract
	AccumTransactor // Write-only binding to the contract
	AccumFilterer   // Log filterer for contract events
}

// AccumCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccumSession struct {
	Contract     *Accum            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccumCallerSession struct {
	Contract *AccumCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AccumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccumTransactorSession struct {
	Contract     *AccumTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccumRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccumRaw struct {
	Contract *Accum // Generic contract binding to access the raw methods on
}

// AccumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccumCallerRaw struct {
	Contract *AccumCaller // Generic read-only contract binding to access the raw methods on
}

// AccumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccumTransactorRaw struct {
	Contract *AccumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccum creates a new instance of Accum, bound to a specific deployed contract.
func NewAccum(address common.Address, backend bind.ContractBackend) (*Accum, error) {
	contract, err := bindAccum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accum{AccumCaller: AccumCaller{contract: contract}, AccumTransactor: AccumTransactor{contract: contract}, AccumFilterer: AccumFilterer{contract: contract}}, nil
}

// NewAccumCaller creates a new read-only instance of Accum, bound to a specific deployed contract.
func NewAccumCaller(address common.Address, caller bind.ContractCaller) (*AccumCaller, error) {
	contract, err := bindAccum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccumCaller{contract: contract}, nil
}

// NewAccumTransactor creates a new write-only instance of Accum, bound to a specific deployed contract.
func NewAccumTransactor(address common.Address, transactor bind.ContractTransactor) (*AccumTransactor, error) {
	contract, err := bindAccum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccumTransactor{contract: contract}, nil
}

// NewAccumFilterer creates a new log filterer instance of Accum, bound to a specific deployed contract.
func NewAccumFilterer(address common.Address, filterer bind.ContractFilterer) (*AccumFilterer, error) {
	contract, err := bindAccum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccumFilterer{contract: contract}, nil
}

// bindAccum binds a generic wrapper to an already deployed contract.
func bindAccum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accum *AccumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accum.Contract.AccumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accum *AccumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accum.Contract.AccumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accum *AccumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accum.Contract.AccumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accum *AccumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accum *AccumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accum *AccumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accum.Contract.contract.Transact(opts, method, params...)
}

// CheckBalance is a free data retrieval call binding the contract method 0xc71daccb.
//
// Solidity: function checkBalance() view returns(uint256)
func (_Accum *AccumCaller) CheckBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Accum.contract.Call(opts, &out, "checkBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckBalance is a free data retrieval call binding the contract method 0xc71daccb.
//
// Solidity: function checkBalance() view returns(uint256)
func (_Accum *AccumSession) CheckBalance() (*big.Int, error) {
	return _Accum.Contract.CheckBalance(&_Accum.CallOpts)
}

// CheckBalance is a free data retrieval call binding the contract method 0xc71daccb.
//
// Solidity: function checkBalance() view returns(uint256)
func (_Accum *AccumCallerSession) CheckBalance() (*big.Int, error) {
	return _Accum.Contract.CheckBalance(&_Accum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accum *AccumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accum.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accum *AccumSession) Owner() (common.Address, error) {
	return _Accum.Contract.Owner(&_Accum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accum *AccumCallerSession) Owner() (common.Address, error) {
	return _Accum.Contract.Owner(&_Accum.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accum *AccumTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accum.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accum *AccumSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accum.Contract.RenounceOwnership(&_Accum.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accum *AccumTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accum.Contract.RenounceOwnership(&_Accum.TransactOpts)
}

// SendEthers is a paid mutator transaction binding the contract method 0x0783cda7.
//
// Solidity: function sendEthers() payable returns()
func (_Accum *AccumTransactor) SendEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accum.contract.Transact(opts, "sendEthers")
}

// SendEthers is a paid mutator transaction binding the contract method 0x0783cda7.
//
// Solidity: function sendEthers() payable returns()
func (_Accum *AccumSession) SendEthers() (*types.Transaction, error) {
	return _Accum.Contract.SendEthers(&_Accum.TransactOpts)
}

// SendEthers is a paid mutator transaction binding the contract method 0x0783cda7.
//
// Solidity: function sendEthers() payable returns()
func (_Accum *AccumTransactorSession) SendEthers() (*types.Transaction, error) {
	return _Accum.Contract.SendEthers(&_Accum.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accum *AccumTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Accum.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accum *AccumSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accum.Contract.TransferOwnership(&_Accum.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accum *AccumTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accum.Contract.TransferOwnership(&_Accum.TransactOpts, newOwner)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Accum *AccumTransactor) WithdrawEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accum.contract.Transact(opts, "withdrawEther")
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Accum *AccumSession) WithdrawEther() (*types.Transaction, error) {
	return _Accum.Contract.WithdrawEther(&_Accum.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Accum *AccumTransactorSession) WithdrawEther() (*types.Transaction, error) {
	return _Accum.Contract.WithdrawEther(&_Accum.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Accum *AccumTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accum.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Accum *AccumSession) Receive() (*types.Transaction, error) {
	return _Accum.Contract.Receive(&_Accum.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Accum *AccumTransactorSession) Receive() (*types.Transaction, error) {
	return _Accum.Contract.Receive(&_Accum.TransactOpts)
}

// AccumAddingIterator is returned from FilterAdding and is used to iterate over the raw logs and unpacked data for Adding events raised by the Accum contract.
type AccumAddingIterator struct {
	Event *AccumAdding // Event containing the contract specifics and raw log

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
func (it *AccumAddingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccumAdding)
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
		it.Event = new(AccumAdding)
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
func (it *AccumAddingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccumAddingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccumAdding represents a Adding event raised by the Accum contract.
type AccumAdding struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAdding is a free log retrieval operation binding the contract event 0xf5319c94a6e3d822f1909045eea12c668e7851fa33a70d99000329ceabd51c50.
//
// Solidity: event Adding(address _from, uint256 _amount)
func (_Accum *AccumFilterer) FilterAdding(opts *bind.FilterOpts) (*AccumAddingIterator, error) {

	logs, sub, err := _Accum.contract.FilterLogs(opts, "Adding")
	if err != nil {
		return nil, err
	}
	return &AccumAddingIterator{contract: _Accum.contract, event: "Adding", logs: logs, sub: sub}, nil
}

// WatchAdding is a free log subscription operation binding the contract event 0xf5319c94a6e3d822f1909045eea12c668e7851fa33a70d99000329ceabd51c50.
//
// Solidity: event Adding(address _from, uint256 _amount)
func (_Accum *AccumFilterer) WatchAdding(opts *bind.WatchOpts, sink chan<- *AccumAdding) (event.Subscription, error) {

	logs, sub, err := _Accum.contract.WatchLogs(opts, "Adding")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccumAdding)
				if err := _Accum.contract.UnpackLog(event, "Adding", log); err != nil {
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

// ParseAdding is a log parse operation binding the contract event 0xf5319c94a6e3d822f1909045eea12c668e7851fa33a70d99000329ceabd51c50.
//
// Solidity: event Adding(address _from, uint256 _amount)
func (_Accum *AccumFilterer) ParseAdding(log types.Log) (*AccumAdding, error) {
	event := new(AccumAdding)
	if err := _Accum.contract.UnpackLog(event, "Adding", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccumExceedIterator is returned from FilterExceed and is used to iterate over the raw logs and unpacked data for Exceed events raised by the Accum contract.
type AccumExceedIterator struct {
	Event *AccumExceed // Event containing the contract specifics and raw log

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
func (it *AccumExceedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccumExceed)
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
		it.Event = new(AccumExceed)
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
func (it *AccumExceedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccumExceedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccumExceed represents a Exceed event raised by the Accum contract.
type AccumExceed struct {
	Amount  *big.Int
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExceed is a free log retrieval operation binding the contract event 0xb0428161e04cd6a37a41d41efbaba21a3404b34b15594f16e373792db46de6ed.
//
// Solidity: event Exceed(uint256 _amount, string _message)
func (_Accum *AccumFilterer) FilterExceed(opts *bind.FilterOpts) (*AccumExceedIterator, error) {

	logs, sub, err := _Accum.contract.FilterLogs(opts, "Exceed")
	if err != nil {
		return nil, err
	}
	return &AccumExceedIterator{contract: _Accum.contract, event: "Exceed", logs: logs, sub: sub}, nil
}

// WatchExceed is a free log subscription operation binding the contract event 0xb0428161e04cd6a37a41d41efbaba21a3404b34b15594f16e373792db46de6ed.
//
// Solidity: event Exceed(uint256 _amount, string _message)
func (_Accum *AccumFilterer) WatchExceed(opts *bind.WatchOpts, sink chan<- *AccumExceed) (event.Subscription, error) {

	logs, sub, err := _Accum.contract.WatchLogs(opts, "Exceed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccumExceed)
				if err := _Accum.contract.UnpackLog(event, "Exceed", log); err != nil {
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

// ParseExceed is a log parse operation binding the contract event 0xb0428161e04cd6a37a41d41efbaba21a3404b34b15594f16e373792db46de6ed.
//
// Solidity: event Exceed(uint256 _amount, string _message)
func (_Accum *AccumFilterer) ParseExceed(log types.Log) (*AccumExceed, error) {
	event := new(AccumExceed)
	if err := _Accum.contract.UnpackLog(event, "Exceed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccumOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Accum contract.
type AccumOwnershipTransferredIterator struct {
	Event *AccumOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccumOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccumOwnershipTransferred)
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
		it.Event = new(AccumOwnershipTransferred)
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
func (it *AccumOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccumOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccumOwnershipTransferred represents a OwnershipTransferred event raised by the Accum contract.
type AccumOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accum *AccumFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccumOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accum.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccumOwnershipTransferredIterator{contract: _Accum.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accum *AccumFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccumOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accum.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccumOwnershipTransferred)
				if err := _Accum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accum *AccumFilterer) ParseOwnershipTransferred(log types.Log) (*AccumOwnershipTransferred, error) {
	event := new(AccumOwnershipTransferred)
	if err := _Accum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
