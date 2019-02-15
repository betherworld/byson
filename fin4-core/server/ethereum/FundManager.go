// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// FundmanagerABI is the input ABI used to generate the binding from.
const FundmanagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setHuntContrAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDevFund\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"purpose\",\"type\":\"string\"}],\"name\":\"payCommProj\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"switchManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConFund\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addEmployee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setProjManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getEmployee\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumEmp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"perc\",\"type\":\"uint8\"}],\"name\":\"donate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"delEmployee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"category\",\"type\":\"string\"},{\"name\":\"purpose\",\"type\":\"string\"}],\"name\":\"payOpCosts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"perc\",\"type\":\"uint8\"}],\"name\":\"donationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amdin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"category\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"purpose\",\"type\":\"string\"}],\"name\":\"costsPayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"purpose\",\"type\":\"string\"}],\"name\":\"bonusDisbursed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"newlyRecruited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"honorablyDischarged\",\"type\":\"event\"}]"

// FundmanagerBin is the compiled bytecode used for deploying new contracts.
const FundmanagerBin = `6080604052600080556000600155600060055534801561001e57600080fd5b50611c238061002e6000396000f3fe6080604052600436106100df576000357c010000000000000000000000000000000000000000000000000000000090048063a502e8081161009c578063d22a2bfc11610076578063d22a2bfc1461055d578063e5b0b36b14610588578063e73dff3d146105b9578063fdcbf57c146105f4576100df565b8063a502e80814610329578063a7952f8914610425578063bca9a5c514610476576100df565b80632046d5f0146100e457806334220ceb146101355780633fc818031461019057806352a3e876146101bb5780635df11291146102ad5780638246b900146102fe575b600080fd5b3480156100f057600080fd5b506101336004803603602081101561010757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061077d565b005b34801561014157600080fd5b5061018e6004803603604081101561015857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610887565b005b34801561019c57600080fd5b506101a5610a21565b6040518082815260200191505060405180910390f35b3480156101c757600080fd5b506102ab600480360360608110156101de57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561022557600080fd5b82018360208201111561023757600080fd5b8035906020019184600183028401116401000000008311171561025957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610a2a565b005b3480156102b957600080fd5b506102fc600480360360208110156102d057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c98565b005b34801561030a57600080fd5b50610313610d38565b6040518082815260200191505060405180910390f35b34801561033557600080fd5b5061040f6004803603604081101561034c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561038957600080fd5b82018360208201111561039b57600080fd5b803590602001918460018302840111640100000000831117156103bd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610d42565b6040518082815260200191505060405180910390f35b34801561043157600080fd5b506104746004803603602081101561044857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fa6565b005b34801561048257600080fd5b506104af6004803603602081101561049957600080fd5b81019080803590602001909291905050506110b0565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610521578082015181840152602081019050610506565b50505050905090810190601f16801561054e5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561056957600080fd5b506105726111ba565b6040518082815260200191505060405180910390f35b6105b76004803603602081101561059e57600080fd5b81019080803560ff1690602001909291905050506111c4565b005b3480156105c557600080fd5b506105f2600480360360208110156105dc57600080fd5b8101908080359060200190929190505050611377565b005b34801561060057600080fd5b5061077b6004803603608081101561061757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561065e57600080fd5b82018360208201111561067057600080fd5b8035906020019184600183028401116401000000008311171561069257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156106f557600080fd5b82018360208201111561070757600080fd5b8035906020019184600183028401116401000000008311171561072957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611780565b005b600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610843576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f43616e206f6e6c7920626520736574206f6e636521000000000000000000000081525060200191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561094c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f416363657373207265737472696374656421000000000000000000000000000081525060200191505060405180910390fd5b60015481111515156109c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f4f7574206f662066756e6473210000000000000000000000000000000000000081525060200191505060405180910390fd5b806001600082825403925050819055508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610a1c573d6000803e3d6000fd5b505050565b60008054905090565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610aeb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f416363657373207265737472696374656421000000000000000000000000000081525060200191505060405180910390fd5b6001548211151515610b65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f4f7574206f662066756e6473210000000000000000000000000000000000000081525060200191505060405180910390fd5b7ff4e62d273949cbb5d95f6d383e9ca88c6db0a4b446bf7d1ddf50a4245050aaa5338383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c01578082015181840152602081019050610be6565b50505050905090810190601f168015610c2e5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a18160008082825403925050819055508273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610c92573d6000803e3d6000fd5b50505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610cf457600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600154905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610da057600080fd5b7f94b9abd59d46991632b031ccc683365d31f06b06c1aaba6ebc37d38fecef4f748383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610e35578082015181840152602081019050610e1a565b50505050905090810190601f168015610e625780820380516001836020036101000a031916815260200191505b50935050505060405180910390a16001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600460408051908101604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610f85929190611a5d565b50505090506005600081548092919060010191905055508091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561106c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f43616e206f6e6c7920626520736574206f6e636521000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060606004838154811015156110c357fe5b906000526020600020906002020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660048481548110151561110457fe5b9060005260206000209060020201600101808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111aa5780601f1061117f576101008083540402835291602001916111aa565b820191906000526020600020905b81548152906001019060200180831161118d57829003601f168201915b5050505050905091509150915091565b6000600554905090565b60006064348115156111d257fe5b0614151561122b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611bd26026913960400191505060405180910390fd5b60008160ff1610158015611243575060648160ff1611155b15156112b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f496e76616c69642070657263656e74616765210000000000000000000000000081525060200191505060405180910390fd5b60648160ff1634028115156112c857fe5b04600080828254019250508190555060648160640360ff1634028115156112eb57fe5b046001600082825401925050819055507f89fdcdfe5b4e613a0bfd27e274ae5a26670ad7fa455498e9a5056765ad7ca5f9333483604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018260ff1660ff168152602001935050505060405180910390a150565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156113d357600080fd5b600081101580156113e5575060055481105b1515611459576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f496e76616c69642061646d696e2069642100000000000000000000000000000081525060200191505060405180910390fd5b7f3f80228e8dadcc90607f1d05c451d6b8d094b8f401267433aa4ff1e6f8045a7c60048281548110151561148957fe5b906000526020600020906002020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166004838154811015156114ca57fe5b9060005260206000209060020201600101604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156115965780601f1061156b57610100808354040283529160200191611596565b820191906000526020600020905b81548152906001019060200180831161157957829003601f168201915b5050935050505060405180910390a16000600660006004848154811015156115ba57fe5b906000526020600020906002020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600460016005540381548110151561165157fe5b906000526020600020906002020160048281548110151561166e57fe5b90600052602060002090600202016000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001820181600101908054600181600116156101000203166002900461170a929190611add565b50905050600460016005540381548110151561172257fe5b9060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006117689190611b64565b50506005600081548092919060019003919050555050565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611841576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f416363657373207265737472696374656421000000000000000000000000000081525060200191505060405180910390fd5b60015483111515156118bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f4f7574206f662066756e6473210000000000000000000000000000000000000081525060200191505060405180910390fd5b826001600082825403925050819055507f870e929233079f78d790a9edba5412dd48d44dfca832307c489510d1fbfff50f33848484604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561196c578082015181840152602081019050611951565b50505050905090810190601f1680156119995780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156119d25780820151818401526020810190506119b7565b50505050905090810190601f1680156119ff5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a18373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050158015611a56573d6000803e3d6000fd5b5050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a9e57805160ff1916838001178555611acc565b82800160010185558215611acc579182015b82811115611acb578251825591602001919060010190611ab0565b5b509050611ad99190611bac565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611b165780548555611b53565b82800160010185558215611b5357600052602060002091601f016020900482015b82811115611b52578254825591600101919060010190611b37565b5b509050611b609190611bac565b5090565b50805460018160011615610100020316600290046000825580601f10611b8a5750611ba9565b601f016020900490600052602060002090810190611ba89190611bac565b5b50565b611bce91905b80821115611bca576000816000905550600101611bb2565b5090565b9056fe446f6e6174696f6e73206e6565647320746f206265206d756c7469706c65206f662031303021a165627a7a7230582041930b7dc7ccb3e8b1aacde5306dd7b0c48976220f40dd2ba456b36d5814f1ca0029`

// DeployFundmanager deploys a new Ethereum contract, binding an instance of Fundmanager to it.
func DeployFundmanager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Fundmanager, error) {
	parsed, err := abi.JSON(strings.NewReader(FundmanagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FundmanagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fundmanager{FundmanagerCaller: FundmanagerCaller{contract: contract}, FundmanagerTransactor: FundmanagerTransactor{contract: contract}, FundmanagerFilterer: FundmanagerFilterer{contract: contract}}, nil
}

// Fundmanager is an auto generated Go binding around an Ethereum contract.
type Fundmanager struct {
	FundmanagerCaller     // Read-only binding to the contract
	FundmanagerTransactor // Write-only binding to the contract
	FundmanagerFilterer   // Log filterer for contract events
}

// FundmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FundmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FundmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FundmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FundmanagerSession struct {
	Contract     *Fundmanager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FundmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FundmanagerCallerSession struct {
	Contract *FundmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FundmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FundmanagerTransactorSession struct {
	Contract     *FundmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FundmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FundmanagerRaw struct {
	Contract *Fundmanager // Generic contract binding to access the raw methods on
}

// FundmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FundmanagerCallerRaw struct {
	Contract *FundmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// FundmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FundmanagerTransactorRaw struct {
	Contract *FundmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFundmanager creates a new instance of Fundmanager, bound to a specific deployed contract.
func NewFundmanager(address common.Address, backend bind.ContractBackend) (*Fundmanager, error) {
	contract, err := bindFundmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fundmanager{FundmanagerCaller: FundmanagerCaller{contract: contract}, FundmanagerTransactor: FundmanagerTransactor{contract: contract}, FundmanagerFilterer: FundmanagerFilterer{contract: contract}}, nil
}

// NewFundmanagerCaller creates a new read-only instance of Fundmanager, bound to a specific deployed contract.
func NewFundmanagerCaller(address common.Address, caller bind.ContractCaller) (*FundmanagerCaller, error) {
	contract, err := bindFundmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FundmanagerCaller{contract: contract}, nil
}

// NewFundmanagerTransactor creates a new write-only instance of Fundmanager, bound to a specific deployed contract.
func NewFundmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FundmanagerTransactor, error) {
	contract, err := bindFundmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FundmanagerTransactor{contract: contract}, nil
}

// NewFundmanagerFilterer creates a new log filterer instance of Fundmanager, bound to a specific deployed contract.
func NewFundmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FundmanagerFilterer, error) {
	contract, err := bindFundmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FundmanagerFilterer{contract: contract}, nil
}

// bindFundmanager binds a generic wrapper to an already deployed contract.
func bindFundmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FundmanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fundmanager *FundmanagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fundmanager.Contract.FundmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fundmanager *FundmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fundmanager.Contract.FundmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fundmanager *FundmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fundmanager.Contract.FundmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fundmanager *FundmanagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fundmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fundmanager *FundmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fundmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fundmanager *FundmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fundmanager.Contract.contract.Transact(opts, method, params...)
}

// GetConFund is a free data retrieval call binding the contract method 0x8246b900.
//
// Solidity: function getConFund() constant returns(uint256)
func (_Fundmanager *FundmanagerCaller) GetConFund(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fundmanager.contract.Call(opts, out, "getConFund")
	return *ret0, err
}

// GetConFund is a free data retrieval call binding the contract method 0x8246b900.
//
// Solidity: function getConFund() constant returns(uint256)
func (_Fundmanager *FundmanagerSession) GetConFund() (*big.Int, error) {
	return _Fundmanager.Contract.GetConFund(&_Fundmanager.CallOpts)
}

// GetConFund is a free data retrieval call binding the contract method 0x8246b900.
//
// Solidity: function getConFund() constant returns(uint256)
func (_Fundmanager *FundmanagerCallerSession) GetConFund() (*big.Int, error) {
	return _Fundmanager.Contract.GetConFund(&_Fundmanager.CallOpts)
}

// GetDevFund is a free data retrieval call binding the contract method 0x3fc81803.
//
// Solidity: function getDevFund() constant returns(uint256)
func (_Fundmanager *FundmanagerCaller) GetDevFund(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fundmanager.contract.Call(opts, out, "getDevFund")
	return *ret0, err
}

// GetDevFund is a free data retrieval call binding the contract method 0x3fc81803.
//
// Solidity: function getDevFund() constant returns(uint256)
func (_Fundmanager *FundmanagerSession) GetDevFund() (*big.Int, error) {
	return _Fundmanager.Contract.GetDevFund(&_Fundmanager.CallOpts)
}

// GetDevFund is a free data retrieval call binding the contract method 0x3fc81803.
//
// Solidity: function getDevFund() constant returns(uint256)
func (_Fundmanager *FundmanagerCallerSession) GetDevFund() (*big.Int, error) {
	return _Fundmanager.Contract.GetDevFund(&_Fundmanager.CallOpts)
}

// GetEmployee is a free data retrieval call binding the contract method 0xbca9a5c5.
//
// Solidity: function getEmployee(uint256 id) constant returns(address, string)
func (_Fundmanager *FundmanagerCaller) GetEmployee(opts *bind.CallOpts, id *big.Int) (common.Address, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Fundmanager.contract.Call(opts, out, "getEmployee", id)
	return *ret0, *ret1, err
}

// GetEmployee is a free data retrieval call binding the contract method 0xbca9a5c5.
//
// Solidity: function getEmployee(uint256 id) constant returns(address, string)
func (_Fundmanager *FundmanagerSession) GetEmployee(id *big.Int) (common.Address, string, error) {
	return _Fundmanager.Contract.GetEmployee(&_Fundmanager.CallOpts, id)
}

// GetEmployee is a free data retrieval call binding the contract method 0xbca9a5c5.
//
// Solidity: function getEmployee(uint256 id) constant returns(address, string)
func (_Fundmanager *FundmanagerCallerSession) GetEmployee(id *big.Int) (common.Address, string, error) {
	return _Fundmanager.Contract.GetEmployee(&_Fundmanager.CallOpts, id)
}

// GetNumEmp is a free data retrieval call binding the contract method 0xd22a2bfc.
//
// Solidity: function getNumEmp() constant returns(uint256)
func (_Fundmanager *FundmanagerCaller) GetNumEmp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fundmanager.contract.Call(opts, out, "getNumEmp")
	return *ret0, err
}

// GetNumEmp is a free data retrieval call binding the contract method 0xd22a2bfc.
//
// Solidity: function getNumEmp() constant returns(uint256)
func (_Fundmanager *FundmanagerSession) GetNumEmp() (*big.Int, error) {
	return _Fundmanager.Contract.GetNumEmp(&_Fundmanager.CallOpts)
}

// GetNumEmp is a free data retrieval call binding the contract method 0xd22a2bfc.
//
// Solidity: function getNumEmp() constant returns(uint256)
func (_Fundmanager *FundmanagerCallerSession) GetNumEmp() (*big.Int, error) {
	return _Fundmanager.Contract.GetNumEmp(&_Fundmanager.CallOpts)
}

// AddEmployee is a paid mutator transaction binding the contract method 0xa502e808.
//
// Solidity: function addEmployee(address account, string name) returns(uint256)
func (_Fundmanager *FundmanagerTransactor) AddEmployee(opts *bind.TransactOpts, account common.Address, name string) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "addEmployee", account, name)
}

// AddEmployee is a paid mutator transaction binding the contract method 0xa502e808.
//
// Solidity: function addEmployee(address account, string name) returns(uint256)
func (_Fundmanager *FundmanagerSession) AddEmployee(account common.Address, name string) (*types.Transaction, error) {
	return _Fundmanager.Contract.AddEmployee(&_Fundmanager.TransactOpts, account, name)
}

// AddEmployee is a paid mutator transaction binding the contract method 0xa502e808.
//
// Solidity: function addEmployee(address account, string name) returns(uint256)
func (_Fundmanager *FundmanagerTransactorSession) AddEmployee(account common.Address, name string) (*types.Transaction, error) {
	return _Fundmanager.Contract.AddEmployee(&_Fundmanager.TransactOpts, account, name)
}

// DelEmployee is a paid mutator transaction binding the contract method 0xe73dff3d.
//
// Solidity: function delEmployee(uint256 id) returns()
func (_Fundmanager *FundmanagerTransactor) DelEmployee(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "delEmployee", id)
}

// DelEmployee is a paid mutator transaction binding the contract method 0xe73dff3d.
//
// Solidity: function delEmployee(uint256 id) returns()
func (_Fundmanager *FundmanagerSession) DelEmployee(id *big.Int) (*types.Transaction, error) {
	return _Fundmanager.Contract.DelEmployee(&_Fundmanager.TransactOpts, id)
}

// DelEmployee is a paid mutator transaction binding the contract method 0xe73dff3d.
//
// Solidity: function delEmployee(uint256 id) returns()
func (_Fundmanager *FundmanagerTransactorSession) DelEmployee(id *big.Int) (*types.Transaction, error) {
	return _Fundmanager.Contract.DelEmployee(&_Fundmanager.TransactOpts, id)
}

// Donate is a paid mutator transaction binding the contract method 0xe5b0b36b.
//
// Solidity: function donate(uint8 perc) returns()
func (_Fundmanager *FundmanagerTransactor) Donate(opts *bind.TransactOpts, perc uint8) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "donate", perc)
}

// Donate is a paid mutator transaction binding the contract method 0xe5b0b36b.
//
// Solidity: function donate(uint8 perc) returns()
func (_Fundmanager *FundmanagerSession) Donate(perc uint8) (*types.Transaction, error) {
	return _Fundmanager.Contract.Donate(&_Fundmanager.TransactOpts, perc)
}

// Donate is a paid mutator transaction binding the contract method 0xe5b0b36b.
//
// Solidity: function donate(uint8 perc) returns()
func (_Fundmanager *FundmanagerTransactorSession) Donate(perc uint8) (*types.Transaction, error) {
	return _Fundmanager.Contract.Donate(&_Fundmanager.TransactOpts, perc)
}

// PayCommProj is a paid mutator transaction binding the contract method 0x52a3e876.
//
// Solidity: function payCommProj(address recipient, uint256 amount, string purpose) returns()
func (_Fundmanager *FundmanagerTransactor) PayCommProj(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "payCommProj", recipient, amount, purpose)
}

// PayCommProj is a paid mutator transaction binding the contract method 0x52a3e876.
//
// Solidity: function payCommProj(address recipient, uint256 amount, string purpose) returns()
func (_Fundmanager *FundmanagerSession) PayCommProj(recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _Fundmanager.Contract.PayCommProj(&_Fundmanager.TransactOpts, recipient, amount, purpose)
}

// PayCommProj is a paid mutator transaction binding the contract method 0x52a3e876.
//
// Solidity: function payCommProj(address recipient, uint256 amount, string purpose) returns()
func (_Fundmanager *FundmanagerTransactorSession) PayCommProj(recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _Fundmanager.Contract.PayCommProj(&_Fundmanager.TransactOpts, recipient, amount, purpose)
}

// PayOpCosts is a paid mutator transaction binding the contract method 0xfdcbf57c.
//
// Solidity: function payOpCosts(address recipient, uint256 amount, string category, string purpose) returns()
func (_Fundmanager *FundmanagerTransactor) PayOpCosts(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, category string, purpose string) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "payOpCosts", recipient, amount, category, purpose)
}

// PayOpCosts is a paid mutator transaction binding the contract method 0xfdcbf57c.
//
// Solidity: function payOpCosts(address recipient, uint256 amount, string category, string purpose) returns()
func (_Fundmanager *FundmanagerSession) PayOpCosts(recipient common.Address, amount *big.Int, category string, purpose string) (*types.Transaction, error) {
	return _Fundmanager.Contract.PayOpCosts(&_Fundmanager.TransactOpts, recipient, amount, category, purpose)
}

// PayOpCosts is a paid mutator transaction binding the contract method 0xfdcbf57c.
//
// Solidity: function payOpCosts(address recipient, uint256 amount, string category, string purpose) returns()
func (_Fundmanager *FundmanagerTransactorSession) PayOpCosts(recipient common.Address, amount *big.Int, category string, purpose string) (*types.Transaction, error) {
	return _Fundmanager.Contract.PayOpCosts(&_Fundmanager.TransactOpts, recipient, amount, category, purpose)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x34220ceb.
//
// Solidity: function redeemEther(address recipient, uint256 amount) returns()
func (_Fundmanager *FundmanagerTransactor) RedeemEther(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "redeemEther", recipient, amount)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x34220ceb.
//
// Solidity: function redeemEther(address recipient, uint256 amount) returns()
func (_Fundmanager *FundmanagerSession) RedeemEther(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fundmanager.Contract.RedeemEther(&_Fundmanager.TransactOpts, recipient, amount)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x34220ceb.
//
// Solidity: function redeemEther(address recipient, uint256 amount) returns()
func (_Fundmanager *FundmanagerTransactorSession) RedeemEther(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fundmanager.Contract.RedeemEther(&_Fundmanager.TransactOpts, recipient, amount)
}

// SetHuntContrAddr is a paid mutator transaction binding the contract method 0x2046d5f0.
//
// Solidity: function setHuntContrAddr(address account) returns()
func (_Fundmanager *FundmanagerTransactor) SetHuntContrAddr(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "setHuntContrAddr", account)
}

// SetHuntContrAddr is a paid mutator transaction binding the contract method 0x2046d5f0.
//
// Solidity: function setHuntContrAddr(address account) returns()
func (_Fundmanager *FundmanagerSession) SetHuntContrAddr(account common.Address) (*types.Transaction, error) {
	return _Fundmanager.Contract.SetHuntContrAddr(&_Fundmanager.TransactOpts, account)
}

// SetHuntContrAddr is a paid mutator transaction binding the contract method 0x2046d5f0.
//
// Solidity: function setHuntContrAddr(address account) returns()
func (_Fundmanager *FundmanagerTransactorSession) SetHuntContrAddr(account common.Address) (*types.Transaction, error) {
	return _Fundmanager.Contract.SetHuntContrAddr(&_Fundmanager.TransactOpts, account)
}

// SetProjManager is a paid mutator transaction binding the contract method 0xa7952f89.
//
// Solidity: function setProjManager(address account) returns()
func (_Fundmanager *FundmanagerTransactor) SetProjManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "setProjManager", account)
}

// SetProjManager is a paid mutator transaction binding the contract method 0xa7952f89.
//
// Solidity: function setProjManager(address account) returns()
func (_Fundmanager *FundmanagerSession) SetProjManager(account common.Address) (*types.Transaction, error) {
	return _Fundmanager.Contract.SetProjManager(&_Fundmanager.TransactOpts, account)
}

// SetProjManager is a paid mutator transaction binding the contract method 0xa7952f89.
//
// Solidity: function setProjManager(address account) returns()
func (_Fundmanager *FundmanagerTransactorSession) SetProjManager(account common.Address) (*types.Transaction, error) {
	return _Fundmanager.Contract.SetProjManager(&_Fundmanager.TransactOpts, account)
}

// SwitchManager is a paid mutator transaction binding the contract method 0x5df11291.
//
// Solidity: function switchManager(address account) returns()
func (_Fundmanager *FundmanagerTransactor) SwitchManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Fundmanager.contract.Transact(opts, "switchManager", account)
}

// SwitchManager is a paid mutator transaction binding the contract method 0x5df11291.
//
// Solidity: function switchManager(address account) returns()
func (_Fundmanager *FundmanagerSession) SwitchManager(account common.Address) (*types.Transaction, error) {
	return _Fundmanager.Contract.SwitchManager(&_Fundmanager.TransactOpts, account)
}

// SwitchManager is a paid mutator transaction binding the contract method 0x5df11291.
//
// Solidity: function switchManager(address account) returns()
func (_Fundmanager *FundmanagerTransactorSession) SwitchManager(account common.Address) (*types.Transaction, error) {
	return _Fundmanager.Contract.SwitchManager(&_Fundmanager.TransactOpts, account)
}

// FundmanagerBonusDisbursedIterator is returned from FilterBonusDisbursed and is used to iterate over the raw logs and unpacked data for BonusDisbursed events raised by the Fundmanager contract.
type FundmanagerBonusDisbursedIterator struct {
	Event *FundmanagerBonusDisbursed // Event containing the contract specifics and raw log

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
func (it *FundmanagerBonusDisbursedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundmanagerBonusDisbursed)
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
		it.Event = new(FundmanagerBonusDisbursed)
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
func (it *FundmanagerBonusDisbursedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundmanagerBonusDisbursedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundmanagerBonusDisbursed represents a BonusDisbursed event raised by the Fundmanager contract.
type FundmanagerBonusDisbursed struct {
	Admin   common.Address
	Amount  *big.Int
	Purpose string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBonusDisbursed is a free log retrieval operation binding the contract event 0xf4e62d273949cbb5d95f6d383e9ca88c6db0a4b446bf7d1ddf50a4245050aaa5.
//
// Solidity: event bonusDisbursed(address admin, uint256 amount, string purpose)
func (_Fundmanager *FundmanagerFilterer) FilterBonusDisbursed(opts *bind.FilterOpts) (*FundmanagerBonusDisbursedIterator, error) {

	logs, sub, err := _Fundmanager.contract.FilterLogs(opts, "bonusDisbursed")
	if err != nil {
		return nil, err
	}
	return &FundmanagerBonusDisbursedIterator{contract: _Fundmanager.contract, event: "bonusDisbursed", logs: logs, sub: sub}, nil
}

// WatchBonusDisbursed is a free log subscription operation binding the contract event 0xf4e62d273949cbb5d95f6d383e9ca88c6db0a4b446bf7d1ddf50a4245050aaa5.
//
// Solidity: event bonusDisbursed(address admin, uint256 amount, string purpose)
func (_Fundmanager *FundmanagerFilterer) WatchBonusDisbursed(opts *bind.WatchOpts, sink chan<- *FundmanagerBonusDisbursed) (event.Subscription, error) {

	logs, sub, err := _Fundmanager.contract.WatchLogs(opts, "bonusDisbursed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundmanagerBonusDisbursed)
				if err := _Fundmanager.contract.UnpackLog(event, "bonusDisbursed", log); err != nil {
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

// FundmanagerCostsPayedIterator is returned from FilterCostsPayed and is used to iterate over the raw logs and unpacked data for CostsPayed events raised by the Fundmanager contract.
type FundmanagerCostsPayedIterator struct {
	Event *FundmanagerCostsPayed // Event containing the contract specifics and raw log

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
func (it *FundmanagerCostsPayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundmanagerCostsPayed)
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
		it.Event = new(FundmanagerCostsPayed)
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
func (it *FundmanagerCostsPayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundmanagerCostsPayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundmanagerCostsPayed represents a CostsPayed event raised by the Fundmanager contract.
type FundmanagerCostsPayed struct {
	Amdin    common.Address
	Amount   *big.Int
	Category string
	Purpose  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCostsPayed is a free log retrieval operation binding the contract event 0x870e929233079f78d790a9edba5412dd48d44dfca832307c489510d1fbfff50f.
//
// Solidity: event costsPayed(address amdin, uint256 amount, string category, string purpose)
func (_Fundmanager *FundmanagerFilterer) FilterCostsPayed(opts *bind.FilterOpts) (*FundmanagerCostsPayedIterator, error) {

	logs, sub, err := _Fundmanager.contract.FilterLogs(opts, "costsPayed")
	if err != nil {
		return nil, err
	}
	return &FundmanagerCostsPayedIterator{contract: _Fundmanager.contract, event: "costsPayed", logs: logs, sub: sub}, nil
}

// WatchCostsPayed is a free log subscription operation binding the contract event 0x870e929233079f78d790a9edba5412dd48d44dfca832307c489510d1fbfff50f.
//
// Solidity: event costsPayed(address amdin, uint256 amount, string category, string purpose)
func (_Fundmanager *FundmanagerFilterer) WatchCostsPayed(opts *bind.WatchOpts, sink chan<- *FundmanagerCostsPayed) (event.Subscription, error) {

	logs, sub, err := _Fundmanager.contract.WatchLogs(opts, "costsPayed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundmanagerCostsPayed)
				if err := _Fundmanager.contract.UnpackLog(event, "costsPayed", log); err != nil {
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

// FundmanagerDonationReceivedIterator is returned from FilterDonationReceived and is used to iterate over the raw logs and unpacked data for DonationReceived events raised by the Fundmanager contract.
type FundmanagerDonationReceivedIterator struct {
	Event *FundmanagerDonationReceived // Event containing the contract specifics and raw log

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
func (it *FundmanagerDonationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundmanagerDonationReceived)
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
		it.Event = new(FundmanagerDonationReceived)
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
func (it *FundmanagerDonationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundmanagerDonationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundmanagerDonationReceived represents a DonationReceived event raised by the Fundmanager contract.
type FundmanagerDonationReceived struct {
	Donor  common.Address
	Amount *big.Int
	Perc   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonationReceived is a free log retrieval operation binding the contract event 0x89fdcdfe5b4e613a0bfd27e274ae5a26670ad7fa455498e9a5056765ad7ca5f9.
//
// Solidity: event donationReceived(address donor, uint256 amount, uint8 perc)
func (_Fundmanager *FundmanagerFilterer) FilterDonationReceived(opts *bind.FilterOpts) (*FundmanagerDonationReceivedIterator, error) {

	logs, sub, err := _Fundmanager.contract.FilterLogs(opts, "donationReceived")
	if err != nil {
		return nil, err
	}
	return &FundmanagerDonationReceivedIterator{contract: _Fundmanager.contract, event: "donationReceived", logs: logs, sub: sub}, nil
}

// WatchDonationReceived is a free log subscription operation binding the contract event 0x89fdcdfe5b4e613a0bfd27e274ae5a26670ad7fa455498e9a5056765ad7ca5f9.
//
// Solidity: event donationReceived(address donor, uint256 amount, uint8 perc)
func (_Fundmanager *FundmanagerFilterer) WatchDonationReceived(opts *bind.WatchOpts, sink chan<- *FundmanagerDonationReceived) (event.Subscription, error) {

	logs, sub, err := _Fundmanager.contract.WatchLogs(opts, "donationReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundmanagerDonationReceived)
				if err := _Fundmanager.contract.UnpackLog(event, "donationReceived", log); err != nil {
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

// FundmanagerHonorablyDischargedIterator is returned from FilterHonorablyDischarged and is used to iterate over the raw logs and unpacked data for HonorablyDischarged events raised by the Fundmanager contract.
type FundmanagerHonorablyDischargedIterator struct {
	Event *FundmanagerHonorablyDischarged // Event containing the contract specifics and raw log

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
func (it *FundmanagerHonorablyDischargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundmanagerHonorablyDischarged)
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
		it.Event = new(FundmanagerHonorablyDischarged)
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
func (it *FundmanagerHonorablyDischargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundmanagerHonorablyDischargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundmanagerHonorablyDischarged represents a HonorablyDischarged event raised by the Fundmanager contract.
type FundmanagerHonorablyDischarged struct {
	Admin common.Address
	Name  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHonorablyDischarged is a free log retrieval operation binding the contract event 0x3f80228e8dadcc90607f1d05c451d6b8d094b8f401267433aa4ff1e6f8045a7c.
//
// Solidity: event honorablyDischarged(address admin, string name)
func (_Fundmanager *FundmanagerFilterer) FilterHonorablyDischarged(opts *bind.FilterOpts) (*FundmanagerHonorablyDischargedIterator, error) {

	logs, sub, err := _Fundmanager.contract.FilterLogs(opts, "honorablyDischarged")
	if err != nil {
		return nil, err
	}
	return &FundmanagerHonorablyDischargedIterator{contract: _Fundmanager.contract, event: "honorablyDischarged", logs: logs, sub: sub}, nil
}

// WatchHonorablyDischarged is a free log subscription operation binding the contract event 0x3f80228e8dadcc90607f1d05c451d6b8d094b8f401267433aa4ff1e6f8045a7c.
//
// Solidity: event honorablyDischarged(address admin, string name)
func (_Fundmanager *FundmanagerFilterer) WatchHonorablyDischarged(opts *bind.WatchOpts, sink chan<- *FundmanagerHonorablyDischarged) (event.Subscription, error) {

	logs, sub, err := _Fundmanager.contract.WatchLogs(opts, "honorablyDischarged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundmanagerHonorablyDischarged)
				if err := _Fundmanager.contract.UnpackLog(event, "honorablyDischarged", log); err != nil {
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

// FundmanagerNewlyRecruitedIterator is returned from FilterNewlyRecruited and is used to iterate over the raw logs and unpacked data for NewlyRecruited events raised by the Fundmanager contract.
type FundmanagerNewlyRecruitedIterator struct {
	Event *FundmanagerNewlyRecruited // Event containing the contract specifics and raw log

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
func (it *FundmanagerNewlyRecruitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundmanagerNewlyRecruited)
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
		it.Event = new(FundmanagerNewlyRecruited)
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
func (it *FundmanagerNewlyRecruitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundmanagerNewlyRecruitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundmanagerNewlyRecruited represents a NewlyRecruited event raised by the Fundmanager contract.
type FundmanagerNewlyRecruited struct {
	Admin common.Address
	Name  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewlyRecruited is a free log retrieval operation binding the contract event 0x94b9abd59d46991632b031ccc683365d31f06b06c1aaba6ebc37d38fecef4f74.
//
// Solidity: event newlyRecruited(address admin, string name)
func (_Fundmanager *FundmanagerFilterer) FilterNewlyRecruited(opts *bind.FilterOpts) (*FundmanagerNewlyRecruitedIterator, error) {

	logs, sub, err := _Fundmanager.contract.FilterLogs(opts, "newlyRecruited")
	if err != nil {
		return nil, err
	}
	return &FundmanagerNewlyRecruitedIterator{contract: _Fundmanager.contract, event: "newlyRecruited", logs: logs, sub: sub}, nil
}

// WatchNewlyRecruited is a free log subscription operation binding the contract event 0x94b9abd59d46991632b031ccc683365d31f06b06c1aaba6ebc37d38fecef4f74.
//
// Solidity: event newlyRecruited(address admin, string name)
func (_Fundmanager *FundmanagerFilterer) WatchNewlyRecruited(opts *bind.WatchOpts, sink chan<- *FundmanagerNewlyRecruited) (event.Subscription, error) {

	logs, sub, err := _Fundmanager.contract.WatchLogs(opts, "newlyRecruited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundmanagerNewlyRecruited)
				if err := _Fundmanager.contract.UnpackLog(event, "newlyRecruited", log); err != nil {
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
