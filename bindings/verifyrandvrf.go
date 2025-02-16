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

// VerifyRandVRFMetaData contains all meta data concerning the VerifyRandVRF contract.
var VerifyRandVRFMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequestStatus\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_verifyRandAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVerifyRand\",\"inputs\":[{\"name\":\"_verifyRandAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyRandAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestSent\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608080604052346100b8577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100a957506001600160401b036002600160401b031982821601610064575b604051610a2890816100bd8239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f8080610055565b63f92ee8a960e01b8152600490fd5b5f80fdfe60406080815260049081361015610014575f80fd5b5f91823560e01c9081631b739ef11461065c57816327f68ce11461061757816338ba461414610410578163485cc955146102ae578163715018a61461024557816382e215ab146102185781638796ba8c146101dc5781638da5cb5b146101a7578163c45a61831461017e578163d8a4676f146100f557508063f2fde38b146100c55763fc2a88c3146100a4575f80fd5b346100c157816003193601126100c1576020906001549051908152f35b5080fd5b82346100f25760203660031901126100f2576100ef6100e2610800565b6100ea6108e8565b610920565b80f35b80fd5b8284346100f257602092836003193601126100c157919091358252600383528082209060ff82541692600180930190825193848784549182815201938352878320925b8882821061016b57886101678989610152828b038361084a565b808051958695151586528501528301906108b5565b0390f35b8454865290940193928201928201610138565b5050346100c157816003193601126100c15760025490516001600160a01b039091168152602090f35b5050346100c157816003193601126100c1575f805160206109d38339815191525490516001600160a01b039091168152602090f35b90503461021457602036600319011261021457359180548310156100f2575061020660209261086c565b91905490519160031b1c8152f35b8280fd5b905034610214576020366003190112610214578160209360ff923581526003855220541690519015158152f35b83346100f257806003193601126100f25761025e6108e8565b5f805160206109d383398151915280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b9050346102145781600319360112610214576102c8610800565b906024356001600160a01b0381169081900361040c577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0092835460ff81871c16159367ffffffffffffffff821680159081610404575b60011490816103fa575b1590816103f1575b506103e3575067ffffffffffffffff19811660011785556103639190846103c4575b5061035b610991565b6100ea610991565b6bffffffffffffffffffffffff60a01b6002541617600255610383578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f808280f35b68ffffffffffffffffff1916680100000000000000011785555f610352565b865163f92ee8a960e01b8152fd5b9050155f610330565b303b159150610328565b86915061031e565b8480fd5b839150346100c157826003193601126100c157803560249384359467ffffffffffffffff8087116106135736602388011215610613578685013594818611610601578560051b978451986020976104698983018c61084a565b8a5284888b0191830101913683116105fd5785899101915b8383106105ed5750506002546001600160a01b03163303915061059c9050578351906104ac8261081a565b60019384835284888401938b8552888b5260038a52878b209051151560ff8019835416911617815501925191825194851161058b57600160401b851161058b57505086908254848455808510610561575b5001908752858720875b83811061055057887ff3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a898961054a8e8b80805195869586528501528301906108b5565b0390a180f35b825182820155918701918401610507565b838a528585848c2092830192015b82811061057d5750506104fd565b8b81558a945087910161056f565b634e487b7160e01b8a526041905288fd5b835162461bcd60e51b81529081018690526026818401527f4f6e6c792076657269667952616e642063616e2063616c6c20746869732066756044820152653731ba34b7b760d11b6064820152608490fd5b8235815291810191899101610481565b8980fd5b634e487b7160e01b8752604190528186fd5b8580fd5b83346100f25760203660031901126100f257610631610800565b6106396108e8565b60018060a01b03166bffffffffffffffffffffffff60a01b600254161760025580f35b8383346100c157806003193601126100c1578235906106796108e8565b805160209067ffffffffffffffff90828101828111828210176107ed5784528581528351916106a78361081a565b868352838301918252858752600384528487209251151560ff80198554169116178355600180930191519182519182116107da57600160401b928383116107c7578590825484845580851061079d575b500190885284882084895b84811061078b5750505050508554908110156107785791848261074e85606097957fe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f0169997018a5561086c565b81549060031b9084821b915f19901b1916179055558151928352602435908301523090820152a180f35b634e487b7160e01b865260418752602486fd5b87845194019381840155018590610702565b838b528685848d2092830192015b8281106107b95750506106f7565b8c81558994508891016107ab565b634e487b7160e01b895260418a52602489fd5b634e487b7160e01b885260418952602488fd5b634e487b7160e01b875260418852602487fd5b600435906001600160a01b038216820361081657565b5f80fd5b6040810190811067ffffffffffffffff82111761083657604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761083657604052565b905f9182548110156108a1578280527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019190565b634e487b7160e01b83526032600452602483fd5b9081518082526020808093019301915f5b8281106108d4575050505090565b8351855293810193928101926001016108c6565b5f805160206109d3833981519152546001600160a01b0316330361090857565b60405163118cdaa760e01b8152336004820152602490fd5b6001600160a01b03908116908115610979575f805160206109d383398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b604051631e4fbdf760e01b81525f6004820152602490fd5b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c16156109c057565b604051631afcd79f60e31b8152600490fdfe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300a26469706673582212205acb79525ce43330e6115ce166cbd856bed1d0587f678bed3d351a34f2f38e4064736f6c63430008140033",
}

// VerifyRandVRFABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyRandVRFMetaData.ABI instead.
var VerifyRandVRFABI = VerifyRandVRFMetaData.ABI

// VerifyRandVRFBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifyRandVRFMetaData.Bin instead.
var VerifyRandVRFBin = VerifyRandVRFMetaData.Bin

// DeployVerifyRandVRF deploys a new Ethereum contract, binding an instance of VerifyRandVRF to it.
func DeployVerifyRandVRF(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerifyRandVRF, error) {
	parsed, err := VerifyRandVRFMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifyRandVRFBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifyRandVRF{VerifyRandVRFCaller: VerifyRandVRFCaller{contract: contract}, VerifyRandVRFTransactor: VerifyRandVRFTransactor{contract: contract}, VerifyRandVRFFilterer: VerifyRandVRFFilterer{contract: contract}}, nil
}

// VerifyRandVRF is an auto generated Go binding around an Ethereum contract.
type VerifyRandVRF struct {
	VerifyRandVRFCaller     // Read-only binding to the contract
	VerifyRandVRFTransactor // Write-only binding to the contract
	VerifyRandVRFFilterer   // Log filterer for contract events
}

// VerifyRandVRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyRandVRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyRandVRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyRandVRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyRandVRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyRandVRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyRandVRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifyRandVRFSession struct {
	Contract     *VerifyRandVRF    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyRandVRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyRandVRFCallerSession struct {
	Contract *VerifyRandVRFCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VerifyRandVRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyRandVRFTransactorSession struct {
	Contract     *VerifyRandVRFTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VerifyRandVRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyRandVRFRaw struct {
	Contract *VerifyRandVRF // Generic contract binding to access the raw methods on
}

// VerifyRandVRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyRandVRFCallerRaw struct {
	Contract *VerifyRandVRFCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyRandVRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyRandVRFTransactorRaw struct {
	Contract *VerifyRandVRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifyRandVRF creates a new instance of VerifyRandVRF, bound to a specific deployed contract.
func NewVerifyRandVRF(address common.Address, backend bind.ContractBackend) (*VerifyRandVRF, error) {
	contract, err := bindVerifyRandVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRF{VerifyRandVRFCaller: VerifyRandVRFCaller{contract: contract}, VerifyRandVRFTransactor: VerifyRandVRFTransactor{contract: contract}, VerifyRandVRFFilterer: VerifyRandVRFFilterer{contract: contract}}, nil
}

// NewVerifyRandVRFCaller creates a new read-only instance of VerifyRandVRF, bound to a specific deployed contract.
func NewVerifyRandVRFCaller(address common.Address, caller bind.ContractCaller) (*VerifyRandVRFCaller, error) {
	contract, err := bindVerifyRandVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFCaller{contract: contract}, nil
}

// NewVerifyRandVRFTransactor creates a new write-only instance of VerifyRandVRF, bound to a specific deployed contract.
func NewVerifyRandVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyRandVRFTransactor, error) {
	contract, err := bindVerifyRandVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFTransactor{contract: contract}, nil
}

// NewVerifyRandVRFFilterer creates a new log filterer instance of VerifyRandVRF, bound to a specific deployed contract.
func NewVerifyRandVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyRandVRFFilterer, error) {
	contract, err := bindVerifyRandVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFFilterer{contract: contract}, nil
}

// bindVerifyRandVRF binds a generic wrapper to an already deployed contract.
func bindVerifyRandVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyRandVRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyRandVRF *VerifyRandVRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyRandVRF.Contract.VerifyRandVRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyRandVRF *VerifyRandVRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.VerifyRandVRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyRandVRF *VerifyRandVRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.VerifyRandVRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyRandVRF *VerifyRandVRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyRandVRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyRandVRF *VerifyRandVRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyRandVRF *VerifyRandVRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.contract.Transact(opts, method, params...)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_VerifyRandVRF *VerifyRandVRFCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	var out []interface{}
	err := _VerifyRandVRF.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(struct {
		Fulfilled   bool
		RandomWords []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_VerifyRandVRF *VerifyRandVRFSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _VerifyRandVRF.Contract.GetRequestStatus(&_VerifyRandVRF.CallOpts, _requestId)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_VerifyRandVRF *VerifyRandVRFCallerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _VerifyRandVRF.Contract.GetRequestStatus(&_VerifyRandVRF.CallOpts, _requestId)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VerifyRandVRF *VerifyRandVRFCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerifyRandVRF.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VerifyRandVRF *VerifyRandVRFSession) LastRequestId() (*big.Int, error) {
	return _VerifyRandVRF.Contract.LastRequestId(&_VerifyRandVRF.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VerifyRandVRF *VerifyRandVRFCallerSession) LastRequestId() (*big.Int, error) {
	return _VerifyRandVRF.Contract.LastRequestId(&_VerifyRandVRF.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyRandVRF *VerifyRandVRFCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyRandVRF.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyRandVRF *VerifyRandVRFSession) Owner() (common.Address, error) {
	return _VerifyRandVRF.Contract.Owner(&_VerifyRandVRF.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyRandVRF *VerifyRandVRFCallerSession) Owner() (common.Address, error) {
	return _VerifyRandVRF.Contract.Owner(&_VerifyRandVRF.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VerifyRandVRF *VerifyRandVRFCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VerifyRandVRF.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VerifyRandVRF *VerifyRandVRFSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _VerifyRandVRF.Contract.RequestIds(&_VerifyRandVRF.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VerifyRandVRF *VerifyRandVRFCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _VerifyRandVRF.Contract.RequestIds(&_VerifyRandVRF.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_VerifyRandVRF *VerifyRandVRFCaller) RequestMapping(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VerifyRandVRF.contract.Call(opts, &out, "requestMapping", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_VerifyRandVRF *VerifyRandVRFSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _VerifyRandVRF.Contract.RequestMapping(&_VerifyRandVRF.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_VerifyRandVRF *VerifyRandVRFCallerSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _VerifyRandVRF.Contract.RequestMapping(&_VerifyRandVRF.CallOpts, arg0)
}

// VerifyRandAddress is a free data retrieval call binding the contract method 0xc45a6183.
//
// Solidity: function verifyRandAddress() view returns(address)
func (_VerifyRandVRF *VerifyRandVRFCaller) VerifyRandAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyRandVRF.contract.Call(opts, &out, "verifyRandAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyRandAddress is a free data retrieval call binding the contract method 0xc45a6183.
//
// Solidity: function verifyRandAddress() view returns(address)
func (_VerifyRandVRF *VerifyRandVRFSession) VerifyRandAddress() (common.Address, error) {
	return _VerifyRandVRF.Contract.VerifyRandAddress(&_VerifyRandVRF.CallOpts)
}

// VerifyRandAddress is a free data retrieval call binding the contract method 0xc45a6183.
//
// Solidity: function verifyRandAddress() view returns(address)
func (_VerifyRandVRF *VerifyRandVRFCallerSession) VerifyRandAddress() (common.Address, error) {
	return _VerifyRandVRF.Contract.VerifyRandAddress(&_VerifyRandVRF.CallOpts)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactor) FulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VerifyRandVRF.contract.Transact(opts, "fulfillRandomWords", _requestId, _randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_VerifyRandVRF *VerifyRandVRFSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.FulfillRandomWords(&_VerifyRandVRF.TransactOpts, _requestId, _randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactorSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.FulfillRandomWords(&_VerifyRandVRF.TransactOpts, _requestId, _randomWords)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _verifyRandAddress) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.contract.Transact(opts, "initialize", initialOwner, _verifyRandAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _verifyRandAddress) returns()
func (_VerifyRandVRF *VerifyRandVRFSession) Initialize(initialOwner common.Address, _verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.Initialize(&_VerifyRandVRF.TransactOpts, initialOwner, _verifyRandAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _verifyRandAddress) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactorSession) Initialize(initialOwner common.Address, _verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.Initialize(&_VerifyRandVRF.TransactOpts, initialOwner, _verifyRandAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyRandVRF *VerifyRandVRFTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyRandVRF.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyRandVRF *VerifyRandVRFSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.RenounceOwnership(&_VerifyRandVRF.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyRandVRF *VerifyRandVRFTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.RenounceOwnership(&_VerifyRandVRF.TransactOpts)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactor) RequestRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _VerifyRandVRF.contract.Transact(opts, "requestRandomWords", _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_VerifyRandVRF *VerifyRandVRFSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.RequestRandomWords(&_VerifyRandVRF.TransactOpts, _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactorSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.RequestRandomWords(&_VerifyRandVRF.TransactOpts, _requestId, _numWords)
}

// SetVerifyRand is a paid mutator transaction binding the contract method 0x27f68ce1.
//
// Solidity: function setVerifyRand(address _verifyRandAddress) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactor) SetVerifyRand(opts *bind.TransactOpts, _verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.contract.Transact(opts, "setVerifyRand", _verifyRandAddress)
}

// SetVerifyRand is a paid mutator transaction binding the contract method 0x27f68ce1.
//
// Solidity: function setVerifyRand(address _verifyRandAddress) returns()
func (_VerifyRandVRF *VerifyRandVRFSession) SetVerifyRand(_verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.SetVerifyRand(&_VerifyRandVRF.TransactOpts, _verifyRandAddress)
}

// SetVerifyRand is a paid mutator transaction binding the contract method 0x27f68ce1.
//
// Solidity: function setVerifyRand(address _verifyRandAddress) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactorSession) SetVerifyRand(_verifyRandAddress common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.SetVerifyRand(&_VerifyRandVRF.TransactOpts, _verifyRandAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyRandVRF *VerifyRandVRFSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.TransferOwnership(&_VerifyRandVRF.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyRandVRF *VerifyRandVRFTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyRandVRF.Contract.TransferOwnership(&_VerifyRandVRF.TransactOpts, newOwner)
}

// VerifyRandVRFFillRandomWordsIterator is returned from FilterFillRandomWords and is used to iterate over the raw logs and unpacked data for FillRandomWords events raised by the VerifyRandVRF contract.
type VerifyRandVRFFillRandomWordsIterator struct {
	Event *VerifyRandVRFFillRandomWords // Event containing the contract specifics and raw log

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
func (it *VerifyRandVRFFillRandomWordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyRandVRFFillRandomWords)
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
		it.Event = new(VerifyRandVRFFillRandomWords)
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
func (it *VerifyRandVRFFillRandomWordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyRandVRFFillRandomWordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyRandVRFFillRandomWords represents a FillRandomWords event raised by the VerifyRandVRF contract.
type VerifyRandVRFFillRandomWords struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFillRandomWords is a free log retrieval operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VerifyRandVRF *VerifyRandVRFFilterer) FilterFillRandomWords(opts *bind.FilterOpts) (*VerifyRandVRFFillRandomWordsIterator, error) {

	logs, sub, err := _VerifyRandVRF.contract.FilterLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFFillRandomWordsIterator{contract: _VerifyRandVRF.contract, event: "FillRandomWords", logs: logs, sub: sub}, nil
}

// WatchFillRandomWords is a free log subscription operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VerifyRandVRF *VerifyRandVRFFilterer) WatchFillRandomWords(opts *bind.WatchOpts, sink chan<- *VerifyRandVRFFillRandomWords) (event.Subscription, error) {

	logs, sub, err := _VerifyRandVRF.contract.WatchLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyRandVRFFillRandomWords)
				if err := _VerifyRandVRF.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
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

// ParseFillRandomWords is a log parse operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VerifyRandVRF *VerifyRandVRFFilterer) ParseFillRandomWords(log types.Log) (*VerifyRandVRFFillRandomWords, error) {
	event := new(VerifyRandVRFFillRandomWords)
	if err := _VerifyRandVRF.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyRandVRFInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VerifyRandVRF contract.
type VerifyRandVRFInitializedIterator struct {
	Event *VerifyRandVRFInitialized // Event containing the contract specifics and raw log

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
func (it *VerifyRandVRFInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyRandVRFInitialized)
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
		it.Event = new(VerifyRandVRFInitialized)
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
func (it *VerifyRandVRFInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyRandVRFInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyRandVRFInitialized represents a Initialized event raised by the VerifyRandVRF contract.
type VerifyRandVRFInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VerifyRandVRF *VerifyRandVRFFilterer) FilterInitialized(opts *bind.FilterOpts) (*VerifyRandVRFInitializedIterator, error) {

	logs, sub, err := _VerifyRandVRF.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFInitializedIterator{contract: _VerifyRandVRF.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VerifyRandVRF *VerifyRandVRFFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VerifyRandVRFInitialized) (event.Subscription, error) {

	logs, sub, err := _VerifyRandVRF.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyRandVRFInitialized)
				if err := _VerifyRandVRF.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VerifyRandVRF *VerifyRandVRFFilterer) ParseInitialized(log types.Log) (*VerifyRandVRFInitialized, error) {
	event := new(VerifyRandVRFInitialized)
	if err := _VerifyRandVRF.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyRandVRFOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifyRandVRF contract.
type VerifyRandVRFOwnershipTransferredIterator struct {
	Event *VerifyRandVRFOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifyRandVRFOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyRandVRFOwnershipTransferred)
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
		it.Event = new(VerifyRandVRFOwnershipTransferred)
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
func (it *VerifyRandVRFOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyRandVRFOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyRandVRFOwnershipTransferred represents a OwnershipTransferred event raised by the VerifyRandVRF contract.
type VerifyRandVRFOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyRandVRF *VerifyRandVRFFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifyRandVRFOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyRandVRF.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFOwnershipTransferredIterator{contract: _VerifyRandVRF.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyRandVRF *VerifyRandVRFFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifyRandVRFOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyRandVRF.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyRandVRFOwnershipTransferred)
				if err := _VerifyRandVRF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifyRandVRF *VerifyRandVRFFilterer) ParseOwnershipTransferred(log types.Log) (*VerifyRandVRFOwnershipTransferred, error) {
	event := new(VerifyRandVRFOwnershipTransferred)
	if err := _VerifyRandVRF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyRandVRFRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the VerifyRandVRF contract.
type VerifyRandVRFRequestSentIterator struct {
	Event *VerifyRandVRFRequestSent // Event containing the contract specifics and raw log

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
func (it *VerifyRandVRFRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyRandVRFRequestSent)
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
		it.Event = new(VerifyRandVRFRequestSent)
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
func (it *VerifyRandVRFRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyRandVRFRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyRandVRFRequestSent represents a RequestSent event raised by the VerifyRandVRF contract.
type VerifyRandVRFRequestSent struct {
	RequestId *big.Int
	NumWords  *big.Int
	Current   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_VerifyRandVRF *VerifyRandVRFFilterer) FilterRequestSent(opts *bind.FilterOpts) (*VerifyRandVRFRequestSentIterator, error) {

	logs, sub, err := _VerifyRandVRF.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &VerifyRandVRFRequestSentIterator{contract: _VerifyRandVRF.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_VerifyRandVRF *VerifyRandVRFFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *VerifyRandVRFRequestSent) (event.Subscription, error) {

	logs, sub, err := _VerifyRandVRF.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyRandVRFRequestSent)
				if err := _VerifyRandVRF.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_VerifyRandVRF *VerifyRandVRFFilterer) ParseRequestSent(log types.Log) (*VerifyRandVRFRequestSent, error) {
	event := new(VerifyRandVRFRequestSent)
	if err := _VerifyRandVRF.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
