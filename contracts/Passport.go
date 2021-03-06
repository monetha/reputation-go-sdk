// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PassportContractABI is the input ABI used to generate the binding from.
const PassportContractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousRegistry\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"PassportLogicRegistryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPassportLogicRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PassportContractBin is the compiled bytecode used for deploying new contracts.
const PassportContractBin = `0x608060405234801561001057600080fd5b5060405160208061092283398101604081815291517f6f72672e6d6f6e657468612e70726f78792e7061757365640000000000000000825291519081900360180190207f9e7945c55c116aa3404b99fe56db7af9613d3b899554a437c2616a4749a94d8a1461007b57fe5b604080517f6f72672e6d6f6e657468612e70726f78792e6f776e6572000000000000000000815290519081900360170190206000805160206108e2833981519152146100c357fe5b6100d5336401000000006101b5810204565b604080517f6f72672e6d6f6e657468612e70726f78792e70656e64696e674f776e657200008152905190819003601e0190207fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb521461012f57fe5b604080517f6f72672e6d6f6e657468612e70617373706f72742e70726f78792e726567697381527f7472790000000000000000000000000000000000000000000000000000000000602082015290519081900360230190206000805160206109028339815191521461019d57fe5b6101af816401000000006101c7810204565b50610279565b6000805160206108e283398151915255565b6000600160a060020a038216151561026657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f43616e6e6f742073657420726567697374727920746f2061207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060008051602061090283398151915255565b61065a806102886000396000f30060806040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634e71e0c8811461009f578063715018a6146100b457806383197ef0146100c957806386d5c5f9146100de5780638da5cb5b1461010f578063e30c397814610124578063f2fde38b14610139578063f5074f411461015a575b61009d61009861017b565b61020c565b005b3480156100ab57600080fd5b5061009d610230565b3480156100c057600080fd5b5061009d610301565b3480156100d557600080fd5b5061009d6103af565b3480156100ea57600080fd5b506100f3610429565b60408051600160a060020a039092168252519081900360200190f35b34801561011b57600080fd5b506100f3610438565b34801561013057600080fd5b506100f3610442565b34801561014557600080fd5b5061009d600160a060020a036004351661044c565b34801561016657600080fd5b5061009d600160a060020a03600435166104bf565b6000610185610532565b600160a060020a031663609725ef6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156101db57600080fd5b505af11580156101ef573d6000803e3d6000fd5b505050506040513d602081101561020557600080fd5b5051905090565b3660008037600080366000845af43d6000803e80801561022b573d6000f35b3d6000fd5b610238610557565b600160a060020a0316331461024c57600080fd5b61025461057c565b15610297576040805160e560020a62461bcd02815260206004820152601d602482015260008051602061060f833981519152604482015290519081900360640190fd5b61029f610557565b600160a060020a03166102b06105a1565b600160a060020a03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36102f56102f0610557565b6105c6565b6102ff60006105ea565b565b6103096105a1565b600160a060020a0316331461031d57600080fd5b61032561057c565b15610368576040805160e560020a62461bcd02815260206004820152601d602482015260008051602061060f833981519152604482015290519081900360640190fd5b6103706105a1565b600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26102ff60006105c6565b6103b76105a1565b600160a060020a031633146103cb57600080fd5b6103d361057c565b15610416576040805160e560020a62461bcd02815260206004820152601d602482015260008051602061060f833981519152604482015290519081900360640190fd5b61041e6105a1565b600160a060020a0316ff5b6000610433610532565b905090565b60006104336105a1565b6000610433610557565b6104546105a1565b600160a060020a0316331461046857600080fd5b61047061057c565b156104b3576040805160e560020a62461bcd02815260206004820152601d602482015260008051602061060f833981519152604482015290519081900360640190fd5b6104bc816105ea565b50565b6104c76105a1565b600160a060020a031633146104db57600080fd5b6104e361057c565b15610526576040805160e560020a62461bcd02815260206004820152601d602482015260008051602061060f833981519152604482015290519081900360640190fd5b80600160a060020a0316ff5b7fa04bab69e45aeb4c94a78ba5bc1be67ef28977c4fdf815a30b829a794eb67a4a5490565b7fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb525490565b7f9e7945c55c116aa3404b99fe56db7af9613d3b899554a437c2616a4749a94d8a5490565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d225490565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d2255565b7fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb52555600636f6e74726163742073686f756c64206e6f7420626520706175736564000000a165627a7a723058204765dfc2d25730d9b971d62801c310ee5636f66cde41b6e586caceb48ad401fc00293ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d22a04bab69e45aeb4c94a78ba5bc1be67ef28977c4fdf815a30b829a794eb67a4a`

// DeployPassportContract deploys a new Ethereum contract, binding an instance of PassportContract to it.
func DeployPassportContract(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *PassportContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PassportContractBin), backend, _registry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PassportContract{PassportContractCaller: PassportContractCaller{contract: contract}, PassportContractTransactor: PassportContractTransactor{contract: contract}, PassportContractFilterer: PassportContractFilterer{contract: contract}}, nil
}

// PassportContract is an auto generated Go binding around an Ethereum contract.
type PassportContract struct {
	PassportContractCaller     // Read-only binding to the contract
	PassportContractTransactor // Write-only binding to the contract
	PassportContractFilterer   // Log filterer for contract events
}

// PassportContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PassportContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PassportContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PassportContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PassportContractSession struct {
	Contract     *PassportContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PassportContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PassportContractCallerSession struct {
	Contract *PassportContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PassportContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PassportContractTransactorSession struct {
	Contract     *PassportContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PassportContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PassportContractRaw struct {
	Contract *PassportContract // Generic contract binding to access the raw methods on
}

// PassportContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PassportContractCallerRaw struct {
	Contract *PassportContractCaller // Generic read-only contract binding to access the raw methods on
}

// PassportContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PassportContractTransactorRaw struct {
	Contract *PassportContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPassportContract creates a new instance of PassportContract, bound to a specific deployed contract.
func NewPassportContract(address common.Address, backend bind.ContractBackend) (*PassportContract, error) {
	contract, err := bindPassportContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PassportContract{PassportContractCaller: PassportContractCaller{contract: contract}, PassportContractTransactor: PassportContractTransactor{contract: contract}, PassportContractFilterer: PassportContractFilterer{contract: contract}}, nil
}

// NewPassportContractCaller creates a new read-only instance of PassportContract, bound to a specific deployed contract.
func NewPassportContractCaller(address common.Address, caller bind.ContractCaller) (*PassportContractCaller, error) {
	contract, err := bindPassportContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PassportContractCaller{contract: contract}, nil
}

// NewPassportContractTransactor creates a new write-only instance of PassportContract, bound to a specific deployed contract.
func NewPassportContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PassportContractTransactor, error) {
	contract, err := bindPassportContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PassportContractTransactor{contract: contract}, nil
}

// NewPassportContractFilterer creates a new log filterer instance of PassportContract, bound to a specific deployed contract.
func NewPassportContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PassportContractFilterer, error) {
	contract, err := bindPassportContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PassportContractFilterer{contract: contract}, nil
}

// bindPassportContract binds a generic wrapper to an already deployed contract.
func bindPassportContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportContract *PassportContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportContract.Contract.PassportContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportContract *PassportContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.Contract.PassportContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportContract *PassportContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportContract.Contract.PassportContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportContract *PassportContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportContract *PassportContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportContract *PassportContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportContract.Contract.contract.Transact(opts, method, params...)
}

// GetPassportLogicRegistry is a free data retrieval call binding the contract method 0x86d5c5f9.
//
// Solidity: function getPassportLogicRegistry() constant returns(address)
func (_PassportContract *PassportContractCaller) GetPassportLogicRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportContract.contract.Call(opts, out, "getPassportLogicRegistry")
	return *ret0, err
}

// GetPassportLogicRegistry is a free data retrieval call binding the contract method 0x86d5c5f9.
//
// Solidity: function getPassportLogicRegistry() constant returns(address)
func (_PassportContract *PassportContractSession) GetPassportLogicRegistry() (common.Address, error) {
	return _PassportContract.Contract.GetPassportLogicRegistry(&_PassportContract.CallOpts)
}

// GetPassportLogicRegistry is a free data retrieval call binding the contract method 0x86d5c5f9.
//
// Solidity: function getPassportLogicRegistry() constant returns(address)
func (_PassportContract *PassportContractCallerSession) GetPassportLogicRegistry() (common.Address, error) {
	return _PassportContract.Contract.GetPassportLogicRegistry(&_PassportContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportContract *PassportContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportContract *PassportContractSession) Owner() (common.Address, error) {
	return _PassportContract.Contract.Owner(&_PassportContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportContract *PassportContractCallerSession) Owner() (common.Address, error) {
	return _PassportContract.Contract.Owner(&_PassportContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportContract *PassportContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportContract.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportContract *PassportContractSession) PendingOwner() (common.Address, error) {
	return _PassportContract.Contract.PendingOwner(&_PassportContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportContract *PassportContractCallerSession) PendingOwner() (common.Address, error) {
	return _PassportContract.Contract.PendingOwner(&_PassportContract.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportContract *PassportContractTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportContract *PassportContractSession) ClaimOwnership() (*types.Transaction, error) {
	return _PassportContract.Contract.ClaimOwnership(&_PassportContract.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportContract *PassportContractTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _PassportContract.Contract.ClaimOwnership(&_PassportContract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PassportContract *PassportContractTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PassportContract *PassportContractSession) Destroy() (*types.Transaction, error) {
	return _PassportContract.Contract.Destroy(&_PassportContract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PassportContract *PassportContractTransactorSession) Destroy() (*types.Transaction, error) {
	return _PassportContract.Contract.Destroy(&_PassportContract.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_PassportContract *PassportContractTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_PassportContract *PassportContractSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.DestroyAndSend(&_PassportContract.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_PassportContract *PassportContractTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.DestroyAndSend(&_PassportContract.TransactOpts, _recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportContract *PassportContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportContract *PassportContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportContract.Contract.RenounceOwnership(&_PassportContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportContract *PassportContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportContract.Contract.RenounceOwnership(&_PassportContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PassportContract *PassportContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PassportContract *PassportContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.TransferOwnership(&_PassportContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PassportContract *PassportContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.TransferOwnership(&_PassportContract.TransactOpts, newOwner)
}

// PassportContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PassportContract contract.
type PassportContractOwnershipRenouncedIterator struct {
	Event *PassportContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PassportContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportContractOwnershipRenounced)
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
		it.Event = new(PassportContractOwnershipRenounced)
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
func (it *PassportContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportContractOwnershipRenounced represents a OwnershipRenounced event raised by the PassportContract contract.
type PassportContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_PassportContract *PassportContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PassportContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportContractOwnershipRenouncedIterator{contract: _PassportContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_PassportContract *PassportContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PassportContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportContractOwnershipRenounced)
				if err := _PassportContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PassportContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PassportContract contract.
type PassportContractOwnershipTransferredIterator struct {
	Event *PassportContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PassportContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportContractOwnershipTransferred)
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
		it.Event = new(PassportContractOwnershipTransferred)
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
func (it *PassportContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportContractOwnershipTransferred represents a OwnershipTransferred event raised by the PassportContract contract.
type PassportContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PassportContract *PassportContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PassportContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportContractOwnershipTransferredIterator{contract: _PassportContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PassportContract *PassportContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PassportContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportContractOwnershipTransferred)
				if err := _PassportContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PassportContractPassportLogicRegistryChangedIterator is returned from FilterPassportLogicRegistryChanged and is used to iterate over the raw logs and unpacked data for PassportLogicRegistryChanged events raised by the PassportContract contract.
type PassportContractPassportLogicRegistryChangedIterator struct {
	Event *PassportContractPassportLogicRegistryChanged // Event containing the contract specifics and raw log

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
func (it *PassportContractPassportLogicRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportContractPassportLogicRegistryChanged)
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
		it.Event = new(PassportContractPassportLogicRegistryChanged)
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
func (it *PassportContractPassportLogicRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportContractPassportLogicRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportContractPassportLogicRegistryChanged represents a PassportLogicRegistryChanged event raised by the PassportContract contract.
type PassportContractPassportLogicRegistryChanged struct {
	PreviousRegistry common.Address
	NewRegistry      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPassportLogicRegistryChanged is a free log retrieval operation binding the contract event 0x5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c2.
//
// Solidity: event PassportLogicRegistryChanged(address indexed previousRegistry, address indexed newRegistry)
func (_PassportContract *PassportContractFilterer) FilterPassportLogicRegistryChanged(opts *bind.FilterOpts, previousRegistry []common.Address, newRegistry []common.Address) (*PassportContractPassportLogicRegistryChangedIterator, error) {

	var previousRegistryRule []interface{}
	for _, previousRegistryItem := range previousRegistry {
		previousRegistryRule = append(previousRegistryRule, previousRegistryItem)
	}
	var newRegistryRule []interface{}
	for _, newRegistryItem := range newRegistry {
		newRegistryRule = append(newRegistryRule, newRegistryItem)
	}

	logs, sub, err := _PassportContract.contract.FilterLogs(opts, "PassportLogicRegistryChanged", previousRegistryRule, newRegistryRule)
	if err != nil {
		return nil, err
	}
	return &PassportContractPassportLogicRegistryChangedIterator{contract: _PassportContract.contract, event: "PassportLogicRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchPassportLogicRegistryChanged is a free log subscription operation binding the contract event 0x5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c2.
//
// Solidity: event PassportLogicRegistryChanged(address indexed previousRegistry, address indexed newRegistry)
func (_PassportContract *PassportContractFilterer) WatchPassportLogicRegistryChanged(opts *bind.WatchOpts, sink chan<- *PassportContractPassportLogicRegistryChanged, previousRegistry []common.Address, newRegistry []common.Address) (event.Subscription, error) {

	var previousRegistryRule []interface{}
	for _, previousRegistryItem := range previousRegistry {
		previousRegistryRule = append(previousRegistryRule, previousRegistryItem)
	}
	var newRegistryRule []interface{}
	for _, newRegistryItem := range newRegistry {
		newRegistryRule = append(newRegistryRule, newRegistryItem)
	}

	logs, sub, err := _PassportContract.contract.WatchLogs(opts, "PassportLogicRegistryChanged", previousRegistryRule, newRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportContractPassportLogicRegistryChanged)
				if err := _PassportContract.contract.UnpackLog(event, "PassportLogicRegistryChanged", log); err != nil {
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
