// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simple_ocr

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

var SimpleOCRMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000816100665760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610096576100968161009d565b5050610146565b336001600160a01b038216036100f55760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005d565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6116db806101556000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b14610161578063afcb95d714610189578063e3d0e712146101a9578063f2fde38b146101bc57600080fd5b8063181f5a771461008d57806379ba5097146100d557806381411834146100df57806381ff7048146100f4575b600080fd5b604080518082018252601081527f53696d706c65204f435220312e302e3000000000000000000000000000000000602082015290516100cc9190611147565b60405180910390f35b6100dd6101cf565b005b6100e76102d1565b6040516100cc91906111b3565b61013e60015460025463ffffffff74010000000000000000000000000000000000000000830481169378010000000000000000000000000000000000000000000000009093041691565b6040805163ffffffff9485168152939092166020840152908201526060016100cc565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100cc565b6040805160018152600060208201819052918101919091526060016100cc565b6100dd6101b73660046113ab565b610340565b6100dd6101ca366004611478565b610eac565b60015473ffffffffffffffffffffffffffffffffffffffff163314610255576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600680548060200260200160405190810160405280929190818152602001828054801561033657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161030b575b5050505050905090565b855185518560ff16601f8311156103b3576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161024c565b8060000361041d576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161024c565b8183146104ab576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161024c565b6104b68160036114c2565b831161051e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161024c565b610526610ec0565b60006040518060c001604052808b81526020018a81526020018960ff1681526020018881526020018767ffffffffffffffff1681526020018681525090505b6005541561071a5760055460009061057f906001906114df565b9050600060058281548110610596576105966114f2565b60009182526020822001546006805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106105d0576105d06114f2565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff85811684526004909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009081169091559290911680845292208054909116905560058054919250908061065057610650611521565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905560068054806106b9576106b9611521565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550610565915050565b60005b815151811015610cc95781518051600091908390811061073f5761073f6114f2565b602002602001015173ffffffffffffffffffffffffffffffffffffffff16036107c4576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7369676e6572206d757374206e6f7420626520656d7074790000000000000000604482015260640161024c565b600073ffffffffffffffffffffffffffffffffffffffff16826020015182815181106107f2576107f26114f2565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603610877576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f7472616e736d6974746572206d757374206e6f7420626520656d707479000000604482015260640161024c565b60006004600084600001518481518110610893576108936114f2565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156108dd576108dd611550565b14610944576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161024c565b6040805180820190915260ff82168152600160208201528251805160049160009185908110610975576109756114f2565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610a1657610a16611550565b021790555060009150610a269050565b6004600084602001518481518110610a4057610a406114f2565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115610a8a57610a8a611550565b14610af1576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161024c565b6040805180820190915260ff821681526020810160028152506004600084602001518481518110610b2457610b246114f2565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610bc557610bc5611550565b021790555050825180516005925083908110610be357610be36114f2565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915582015180516006919083908110610c5f57610c5f6114f2565b60209081029190910181015182546001808201855560009485529290932090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091550161071d565b506040810151600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600180547fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff8116780100000000000000000000000000000000000000000000000063ffffffff4381168202929092178085559204811692918291601491610d819184917401000000000000000000000000000000000000000090041661157f565b92506101000a81548163ffffffff021916908363ffffffff160217905550610de04630600160149054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151610f43565b600281905582518051600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010060ff9093169290920291909117905560015460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598610e97988b9891977401000000000000000000000000000000000000000090920463ffffffff169690959194919391926115a3565b60405180910390a15050505050505050505050565b610eb4610ec0565b610ebd81610fee565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610f41576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161024c565b565b6000808a8a8a8a8a8a8a8a8a604051602001610f6799989796959493929190611639565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff82160361106d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161024c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000815180845260005b81811015611109576020818501810151868301820152016110ed565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061115a60208301846110e3565b9392505050565b60008151808452602080850194506020840160005b838110156111a857815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611176565b509495945050505050565b60208152600061115a6020830184611161565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561123c5761123c6111c6565b604052919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461126857600080fd5b919050565b600082601f83011261127e57600080fd5b8135602067ffffffffffffffff82111561129a5761129a6111c6565b8160051b6112a98282016111f5565b92835284810182019282810190878511156112c357600080fd5b83870192505b848310156112e9576112da83611244565b825291830191908301906112c9565b979650505050505050565b803560ff8116811461126857600080fd5b600082601f83011261131657600080fd5b813567ffffffffffffffff811115611330576113306111c6565b61136160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016111f5565b81815284602083860101111561137657600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff8116811461126857600080fd5b60008060008060008060c087890312156113c457600080fd5b863567ffffffffffffffff808211156113dc57600080fd5b6113e88a838b0161126d565b975060208901359150808211156113fe57600080fd5b61140a8a838b0161126d565b965061141860408a016112f4565b9550606089013591508082111561142e57600080fd5b61143a8a838b01611305565b945061144860808a01611393565b935060a089013591508082111561145e57600080fd5b5061146b89828a01611305565b9150509295509295509295565b60006020828403121561148a57600080fd5b61115a82611244565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820281158282048414176114d9576114d9611493565b92915050565b818103818111156114d9576114d9611493565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b63ffffffff81811683821601908082111561159c5761159c611493565b5092915050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526115d38184018a611161565b905082810360808401526115e78189611161565b905060ff871660a084015282810360c084015261160481876110e3565b905067ffffffffffffffff851660e084015282810361010084015261162981856110e3565b9c9b505050505050505050505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526116808285018b611161565b91508382036080850152611694828a611161565b915060ff881660a085015283820360c08501526116b182886110e3565b90861660e0850152838103610100850152905061162981856110e356fea164736f6c6343000818000a",
}

var SimpleOCRABI = SimpleOCRMetaData.ABI

var SimpleOCRBin = SimpleOCRMetaData.Bin

func DeploySimpleOCR(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleOCR, error) {
	parsed, err := SimpleOCRMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleOCRBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleOCR{address: address, abi: *parsed, SimpleOCRCaller: SimpleOCRCaller{contract: contract}, SimpleOCRTransactor: SimpleOCRTransactor{contract: contract}, SimpleOCRFilterer: SimpleOCRFilterer{contract: contract}}, nil
}

type SimpleOCR struct {
	address common.Address
	abi     abi.ABI
	SimpleOCRCaller
	SimpleOCRTransactor
	SimpleOCRFilterer
}

type SimpleOCRCaller struct {
	contract *bind.BoundContract
}

type SimpleOCRTransactor struct {
	contract *bind.BoundContract
}

type SimpleOCRFilterer struct {
	contract *bind.BoundContract
}

type SimpleOCRSession struct {
	Contract     *SimpleOCR
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SimpleOCRCallerSession struct {
	Contract *SimpleOCRCaller
	CallOpts bind.CallOpts
}

type SimpleOCRTransactorSession struct {
	Contract     *SimpleOCRTransactor
	TransactOpts bind.TransactOpts
}

type SimpleOCRRaw struct {
	Contract *SimpleOCR
}

type SimpleOCRCallerRaw struct {
	Contract *SimpleOCRCaller
}

type SimpleOCRTransactorRaw struct {
	Contract *SimpleOCRTransactor
}

func NewSimpleOCR(address common.Address, backend bind.ContractBackend) (*SimpleOCR, error) {
	abi, err := abi.JSON(strings.NewReader(SimpleOCRABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSimpleOCR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleOCR{address: address, abi: abi, SimpleOCRCaller: SimpleOCRCaller{contract: contract}, SimpleOCRTransactor: SimpleOCRTransactor{contract: contract}, SimpleOCRFilterer: SimpleOCRFilterer{contract: contract}}, nil
}

func NewSimpleOCRCaller(address common.Address, caller bind.ContractCaller) (*SimpleOCRCaller, error) {
	contract, err := bindSimpleOCR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleOCRCaller{contract: contract}, nil
}

func NewSimpleOCRTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleOCRTransactor, error) {
	contract, err := bindSimpleOCR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleOCRTransactor{contract: contract}, nil
}

func NewSimpleOCRFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleOCRFilterer, error) {
	contract, err := bindSimpleOCR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleOCRFilterer{contract: contract}, nil
}

func bindSimpleOCR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleOCRMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SimpleOCR *SimpleOCRRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleOCR.Contract.SimpleOCRCaller.contract.Call(opts, result, method, params...)
}

func (_SimpleOCR *SimpleOCRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleOCR.Contract.SimpleOCRTransactor.contract.Transfer(opts)
}

func (_SimpleOCR *SimpleOCRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleOCR.Contract.SimpleOCRTransactor.contract.Transact(opts, method, params...)
}

func (_SimpleOCR *SimpleOCRCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleOCR.Contract.contract.Call(opts, result, method, params...)
}

func (_SimpleOCR *SimpleOCRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleOCR.Contract.contract.Transfer(opts)
}

func (_SimpleOCR *SimpleOCRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleOCR.Contract.contract.Transact(opts, method, params...)
}

func (_SimpleOCR *SimpleOCRCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _SimpleOCR.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_SimpleOCR *SimpleOCRSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _SimpleOCR.Contract.LatestConfigDetails(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _SimpleOCR.Contract.LatestConfigDetails(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _SimpleOCR.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_SimpleOCR *SimpleOCRSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _SimpleOCR.Contract.LatestConfigDigestAndEpoch(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _SimpleOCR.Contract.LatestConfigDigestAndEpoch(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleOCR.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SimpleOCR *SimpleOCRSession) Owner() (common.Address, error) {
	return _SimpleOCR.Contract.Owner(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCallerSession) Owner() (common.Address, error) {
	return _SimpleOCR.Contract.Owner(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SimpleOCR.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SimpleOCR *SimpleOCRSession) Transmitters() ([]common.Address, error) {
	return _SimpleOCR.Contract.Transmitters(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCallerSession) Transmitters() ([]common.Address, error) {
	return _SimpleOCR.Contract.Transmitters(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleOCR.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SimpleOCR *SimpleOCRSession) TypeAndVersion() (string, error) {
	return _SimpleOCR.Contract.TypeAndVersion(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRCallerSession) TypeAndVersion() (string, error) {
	return _SimpleOCR.Contract.TypeAndVersion(&_SimpleOCR.CallOpts)
}

func (_SimpleOCR *SimpleOCRTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleOCR.contract.Transact(opts, "acceptOwnership")
}

func (_SimpleOCR *SimpleOCRSession) AcceptOwnership() (*types.Transaction, error) {
	return _SimpleOCR.Contract.AcceptOwnership(&_SimpleOCR.TransactOpts)
}

func (_SimpleOCR *SimpleOCRTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SimpleOCR.Contract.AcceptOwnership(&_SimpleOCR.TransactOpts)
}

func (_SimpleOCR *SimpleOCRTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _SimpleOCR.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_SimpleOCR *SimpleOCRSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _SimpleOCR.Contract.SetConfig(&_SimpleOCR.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_SimpleOCR *SimpleOCRTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _SimpleOCR.Contract.SetConfig(&_SimpleOCR.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_SimpleOCR *SimpleOCRTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SimpleOCR.contract.Transact(opts, "transferOwnership", to)
}

func (_SimpleOCR *SimpleOCRSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SimpleOCR.Contract.TransferOwnership(&_SimpleOCR.TransactOpts, to)
}

func (_SimpleOCR *SimpleOCRTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SimpleOCR.Contract.TransferOwnership(&_SimpleOCR.TransactOpts, to)
}

type SimpleOCRConfigSetIterator struct {
	Event *SimpleOCRConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SimpleOCRConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleOCRConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(SimpleOCRConfigSet)
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

func (it *SimpleOCRConfigSetIterator) Error() error {
	return it.fail
}

func (it *SimpleOCRConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SimpleOCRConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_SimpleOCR *SimpleOCRFilterer) FilterConfigSet(opts *bind.FilterOpts) (*SimpleOCRConfigSetIterator, error) {

	logs, sub, err := _SimpleOCR.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &SimpleOCRConfigSetIterator{contract: _SimpleOCR.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_SimpleOCR *SimpleOCRFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *SimpleOCRConfigSet) (event.Subscription, error) {

	logs, sub, err := _SimpleOCR.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SimpleOCRConfigSet)
				if err := _SimpleOCR.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_SimpleOCR *SimpleOCRFilterer) ParseConfigSet(log types.Log) (*SimpleOCRConfigSet, error) {
	event := new(SimpleOCRConfigSet)
	if err := _SimpleOCR.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SimpleOCROwnershipTransferRequestedIterator struct {
	Event *SimpleOCROwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SimpleOCROwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleOCROwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(SimpleOCROwnershipTransferRequested)
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

func (it *SimpleOCROwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SimpleOCROwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SimpleOCROwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SimpleOCR *SimpleOCRFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleOCROwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleOCR.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleOCROwnershipTransferRequestedIterator{contract: _SimpleOCR.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SimpleOCR *SimpleOCRFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SimpleOCROwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleOCR.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SimpleOCROwnershipTransferRequested)
				if err := _SimpleOCR.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SimpleOCR *SimpleOCRFilterer) ParseOwnershipTransferRequested(log types.Log) (*SimpleOCROwnershipTransferRequested, error) {
	event := new(SimpleOCROwnershipTransferRequested)
	if err := _SimpleOCR.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SimpleOCROwnershipTransferredIterator struct {
	Event *SimpleOCROwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SimpleOCROwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleOCROwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(SimpleOCROwnershipTransferred)
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

func (it *SimpleOCROwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SimpleOCROwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SimpleOCROwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SimpleOCR *SimpleOCRFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleOCROwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleOCR.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleOCROwnershipTransferredIterator{contract: _SimpleOCR.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SimpleOCR *SimpleOCRFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SimpleOCROwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleOCR.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SimpleOCROwnershipTransferred)
				if err := _SimpleOCR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SimpleOCR *SimpleOCRFilterer) ParseOwnershipTransferred(log types.Log) (*SimpleOCROwnershipTransferred, error) {
	event := new(SimpleOCROwnershipTransferred)
	if err := _SimpleOCR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_SimpleOCR *SimpleOCR) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _SimpleOCR.abi.Events["ConfigSet"].ID:
		return _SimpleOCR.ParseConfigSet(log)
	case _SimpleOCR.abi.Events["OwnershipTransferRequested"].ID:
		return _SimpleOCR.ParseOwnershipTransferRequested(log)
	case _SimpleOCR.abi.Events["OwnershipTransferred"].ID:
		return _SimpleOCR.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (SimpleOCRConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (SimpleOCROwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SimpleOCROwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_SimpleOCR *SimpleOCR) Address() common.Address {
	return _SimpleOCR.address
}

type SimpleOCRInterface interface {
	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*SimpleOCRConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *SimpleOCRConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*SimpleOCRConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleOCROwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SimpleOCROwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SimpleOCROwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleOCROwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SimpleOCROwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SimpleOCROwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
