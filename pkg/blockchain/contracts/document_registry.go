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
	_ = abi.ConvertType
)

// BlockchainMetaData contains all meta data concerning the Blockchain contract.
var BlockchainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsCID\",\"type\":\"string\"}],\"name\":\"DocumentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"DocumentVerified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"}],\"name\":\"getDocument\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsCID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"recordVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsCID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registerDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"}],\"name\":\"verifyOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506111d98061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80636a243d5b1461004e5780639409d3101461006a578063b10d6b411461009b578063c4fd1289146100ce575b5f80fd5b610068600480360381019061006391906109f4565b6100ea565b005b610084600480360381019061007f9190610a7c565b6102b2565b604051610092929190610b00565b60405180910390f35b6100b560048036038101906100b09190610a7c565b6104a2565b6040516100c59493929190610bf1565b60405180910390f35b6100e860048036038101906100e39190610c6c565b6106e3565b005b5f73ffffffffffffffffffffffffffffffffffffffff165f808581526020019081526020015f206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461018a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018190610cf4565b60405180910390fd5b6040518060a001604052808481526020013373ffffffffffffffffffffffffffffffffffffffff168152602001428152602001838152602001828152505f808581526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030190816102449190610f0c565b50608082015181600401908161025a9190611033565b509050503373ffffffffffffffffffffffffffffffffffffffff16837fdba629cc54aad70846b805286cd90bf0920d3b29656d9b042814d2a2dedf4acb846040516102a59190611102565b60405180910390a3505050565b5f805f805f8581526020019081526020015f206040518060a00160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815260200160038201805461034890610d3f565b80601f016020809104026020016040519081016040528092919081815260200182805461037490610d3f565b80156103bf5780601f10610396576101008083540402835291602001916103bf565b820191905f5260205f20905b8154815290600101906020018083116103a257829003601f168201915b505050505081526020016004820180546103d890610d3f565b80601f016020809104026020016040519081016040528092919081815260200182805461040490610d3f565b801561044f5780601f106104265761010080835404028352916020019161044f565b820191905f5260205f20905b81548152906001019060200180831161043257829003601f168201915b50505050508152505090505f8073ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff1614159050808260200151935093505050915091565b5f806060805f805f8781526020019081526020015f206040518060a00160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815260200160038201805461053b90610d3f565b80601f016020809104026020016040519081016040528092919081815260200182805461056790610d3f565b80156105b25780601f10610589576101008083540402835291602001916105b2565b820191905f5260205f20905b81548152906001019060200180831161059557829003601f168201915b505050505081526020016004820180546105cb90610d3f565b80601f01602080910402602001604051908101604052809291908181526020018280546105f790610d3f565b80156106425780601f1061061957610100808354040283529160200191610642565b820191905f5260205f20905b81548152906001019060200180831161062557829003601f168201915b50505050508152505090505f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16036106bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b69061116c565b60405180910390fd5b80602001518160400151826060015183608001519450945094509450509193509193565b5f73ffffffffffffffffffffffffffffffffffffffff165f808481526020019081526020015f206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610783576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077a9061116c565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16827f6894ebbaf406311560473c5438f7c75d70dc611e024968e8f8596dbd034c5084836040516107ca919061118a565b60405180910390a35050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6107f9816107e7565b8114610803575f80fd5b50565b5f81359050610814816107f0565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61086882610822565b810181811067ffffffffffffffff8211171561088757610886610832565b5b80604052505050565b5f6108996107d6565b90506108a5828261085f565b919050565b5f67ffffffffffffffff8211156108c4576108c3610832565b5b6108cd82610822565b9050602081019050919050565b828183375f83830152505050565b5f6108fa6108f5846108aa565b610890565b9050828152602081018484840111156109165761091561081e565b5b6109218482856108da565b509392505050565b5f82601f83011261093d5761093c61081a565b5b813561094d8482602086016108e8565b91505092915050565b5f67ffffffffffffffff8211156109705761096f610832565b5b61097982610822565b9050602081019050919050565b5f61099861099384610956565b610890565b9050828152602081018484840111156109b4576109b361081e565b5b6109bf8482856108da565b509392505050565b5f82601f8301126109db576109da61081a565b5b81356109eb848260208601610986565b91505092915050565b5f805f60608486031215610a0b57610a0a6107df565b5b5f610a1886828701610806565b935050602084013567ffffffffffffffff811115610a3957610a386107e3565b5b610a4586828701610929565b925050604084013567ffffffffffffffff811115610a6657610a656107e3565b5b610a72868287016109c7565b9150509250925092565b5f60208284031215610a9157610a906107df565b5b5f610a9e84828501610806565b91505092915050565b5f8115159050919050565b610abb81610aa7565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610aea82610ac1565b9050919050565b610afa81610ae0565b82525050565b5f604082019050610b135f830185610ab2565b610b206020830184610af1565b9392505050565b5f819050919050565b610b3981610b27565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610b7182610b3f565b610b7b8185610b49565b9350610b8b818560208601610b59565b610b9481610822565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f610bc382610b9f565b610bcd8185610ba9565b9350610bdd818560208601610b59565b610be681610822565b840191505092915050565b5f608082019050610c045f830187610af1565b610c116020830186610b30565b8181036040830152610c238185610b67565b90508181036060830152610c378184610bb9565b905095945050505050565b610c4b81610aa7565b8114610c55575f80fd5b50565b5f81359050610c6681610c42565b92915050565b5f8060408385031215610c8257610c816107df565b5b5f610c8f85828601610806565b9250506020610ca085828601610c58565b9150509250929050565b7f446f63756d656e7420616c7265616479207265676973746572656400000000005f82015250565b5f610cde601b83610b49565b9150610ce982610caa565b602082019050919050565b5f6020820190508181035f830152610d0b81610cd2565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610d5657607f821691505b602082108103610d6957610d68610d12565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610dcb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d90565b610dd58683610d90565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610e10610e0b610e0684610b27565b610ded565b610b27565b9050919050565b5f819050919050565b610e2983610df6565b610e3d610e3582610e17565b848454610d9c565b825550505050565b5f90565b610e51610e45565b610e5c818484610e20565b505050565b5b81811015610e7f57610e745f82610e49565b600181019050610e62565b5050565b601f821115610ec457610e9581610d6f565b610e9e84610d81565b81016020851015610ead578190505b610ec1610eb985610d81565b830182610e61565b50505b505050565b5f82821c905092915050565b5f610ee45f1984600802610ec9565b1980831691505092915050565b5f610efc8383610ed5565b9150826002028217905092915050565b610f1582610b3f565b67ffffffffffffffff811115610f2e57610f2d610832565b5b610f388254610d3f565b610f43828285610e83565b5f60209050601f831160018114610f74575f8415610f62578287015190505b610f6c8582610ef1565b865550610fd3565b601f198416610f8286610d6f565b5f5b82811015610fa957848901518255600182019150602085019450602081019050610f84565b86831015610fc65784890151610fc2601f891682610ed5565b8355505b6001600288020188555050505b505050505050565b5f819050815f5260205f209050919050565b601f82111561102e57610fff81610fdb565b61100884610d81565b81016020851015611017578190505b61102b61102385610d81565b830182610e61565b50505b505050565b61103c82610b9f565b67ffffffffffffffff81111561105557611054610832565b5b61105f8254610d3f565b61106a828285610fed565b5f60209050601f83116001811461109b575f8415611089578287015190505b6110938582610ef1565b8655506110fa565b601f1984166110a986610fdb565b5f5b828110156110d0578489015182556001820191506020850194506020810190506110ab565b868310156110ed57848901516110e9601f891682610ed5565b8355505b6001600288020188555050505b505050505050565b5f6020820190508181035f83015261111a8184610b67565b905092915050565b7f446f63756d656e74206e6f7420726567697374657265640000000000000000005f82015250565b5f611156601783610b49565b915061116182611122565b602082019050919050565b5f6020820190508181035f8301526111838161114a565b9050919050565b5f60208201905061119d5f830184610ab2565b9291505056fea26469706673582212204bcb6398486058a26e44e4f70130956a9e43dfb352b50c656fb48fed927a4b0c64736f6c634300081a0033",
}

// BlockchainABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainMetaData.ABI instead.
var BlockchainABI = BlockchainMetaData.ABI

// BlockchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockchainMetaData.Bin instead.
var BlockchainBin = BlockchainMetaData.Bin

// DeployBlockchain deploys a new Ethereum contract, binding an instance of Blockchain to it.
func DeployBlockchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Blockchain, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// Blockchain is an auto generated Go binding around an Ethereum contract.
type Blockchain struct {
	BlockchainCaller     // Read-only binding to the contract
	BlockchainTransactor // Write-only binding to the contract
	BlockchainFilterer   // Log filterer for contract events
}

// BlockchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainSession struct {
	Contract     *Blockchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainCallerSession struct {
	Contract *BlockchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlockchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainTransactorSession struct {
	Contract     *BlockchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainRaw struct {
	Contract *Blockchain // Generic contract binding to access the raw methods on
}

// BlockchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainCallerRaw struct {
	Contract *BlockchainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainTransactorRaw struct {
	Contract *BlockchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchain creates a new instance of Blockchain, bound to a specific deployed contract.
func NewBlockchain(address common.Address, backend bind.ContractBackend) (*Blockchain, error) {
	contract, err := bindBlockchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// NewBlockchainCaller creates a new read-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainCaller(address common.Address, caller bind.ContractCaller) (*BlockchainCaller, error) {
	contract, err := bindBlockchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainCaller{contract: contract}, nil
}

// NewBlockchainTransactor creates a new write-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainTransactor, error) {
	contract, err := bindBlockchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransactor{contract: contract}, nil
}

// NewBlockchainFilterer creates a new log filterer instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFilterer, error) {
	contract, err := bindBlockchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFilterer{contract: contract}, nil
}

// bindBlockchain binds a generic wrapper to an already deployed contract.
func bindBlockchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.BlockchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transact(opts, method, params...)
}

// GetDocument is a free data retrieval call binding the contract method 0xb10d6b41.
//
// Solidity: function getDocument(bytes32 documentHash) view returns(address owner, uint256 timestamp, string ipfsCID, bytes publicKey)
func (_Blockchain *BlockchainCaller) GetDocument(opts *bind.CallOpts, documentHash [32]byte) (struct {
	Owner     common.Address
	Timestamp *big.Int
	IpfsCID   string
	PublicKey []byte
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getDocument", documentHash)

	outstruct := new(struct {
		Owner     common.Address
		Timestamp *big.Int
		IpfsCID   string
		PublicKey []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IpfsCID = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PublicKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetDocument is a free data retrieval call binding the contract method 0xb10d6b41.
//
// Solidity: function getDocument(bytes32 documentHash) view returns(address owner, uint256 timestamp, string ipfsCID, bytes publicKey)
func (_Blockchain *BlockchainSession) GetDocument(documentHash [32]byte) (struct {
	Owner     common.Address
	Timestamp *big.Int
	IpfsCID   string
	PublicKey []byte
}, error) {
	return _Blockchain.Contract.GetDocument(&_Blockchain.CallOpts, documentHash)
}

// GetDocument is a free data retrieval call binding the contract method 0xb10d6b41.
//
// Solidity: function getDocument(bytes32 documentHash) view returns(address owner, uint256 timestamp, string ipfsCID, bytes publicKey)
func (_Blockchain *BlockchainCallerSession) GetDocument(documentHash [32]byte) (struct {
	Owner     common.Address
	Timestamp *big.Int
	IpfsCID   string
	PublicKey []byte
}, error) {
	return _Blockchain.Contract.GetDocument(&_Blockchain.CallOpts, documentHash)
}

// VerifyOwnership is a free data retrieval call binding the contract method 0x9409d310.
//
// Solidity: function verifyOwnership(bytes32 documentHash) view returns(bool isRegistered, address owner)
func (_Blockchain *BlockchainCaller) VerifyOwnership(opts *bind.CallOpts, documentHash [32]byte) (struct {
	IsRegistered bool
	Owner        common.Address
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "verifyOwnership", documentHash)

	outstruct := new(struct {
		IsRegistered bool
		Owner        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRegistered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// VerifyOwnership is a free data retrieval call binding the contract method 0x9409d310.
//
// Solidity: function verifyOwnership(bytes32 documentHash) view returns(bool isRegistered, address owner)
func (_Blockchain *BlockchainSession) VerifyOwnership(documentHash [32]byte) (struct {
	IsRegistered bool
	Owner        common.Address
}, error) {
	return _Blockchain.Contract.VerifyOwnership(&_Blockchain.CallOpts, documentHash)
}

// VerifyOwnership is a free data retrieval call binding the contract method 0x9409d310.
//
// Solidity: function verifyOwnership(bytes32 documentHash) view returns(bool isRegistered, address owner)
func (_Blockchain *BlockchainCallerSession) VerifyOwnership(documentHash [32]byte) (struct {
	IsRegistered bool
	Owner        common.Address
}, error) {
	return _Blockchain.Contract.VerifyOwnership(&_Blockchain.CallOpts, documentHash)
}

// RecordVerification is a paid mutator transaction binding the contract method 0xc4fd1289.
//
// Solidity: function recordVerification(bytes32 documentHash, bool verified) returns()
func (_Blockchain *BlockchainTransactor) RecordVerification(opts *bind.TransactOpts, documentHash [32]byte, verified bool) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "recordVerification", documentHash, verified)
}

// RecordVerification is a paid mutator transaction binding the contract method 0xc4fd1289.
//
// Solidity: function recordVerification(bytes32 documentHash, bool verified) returns()
func (_Blockchain *BlockchainSession) RecordVerification(documentHash [32]byte, verified bool) (*types.Transaction, error) {
	return _Blockchain.Contract.RecordVerification(&_Blockchain.TransactOpts, documentHash, verified)
}

// RecordVerification is a paid mutator transaction binding the contract method 0xc4fd1289.
//
// Solidity: function recordVerification(bytes32 documentHash, bool verified) returns()
func (_Blockchain *BlockchainTransactorSession) RecordVerification(documentHash [32]byte, verified bool) (*types.Transaction, error) {
	return _Blockchain.Contract.RecordVerification(&_Blockchain.TransactOpts, documentHash, verified)
}

// RegisterDocument is a paid mutator transaction binding the contract method 0x6a243d5b.
//
// Solidity: function registerDocument(bytes32 documentHash, string ipfsCID, bytes publicKey) returns()
func (_Blockchain *BlockchainTransactor) RegisterDocument(opts *bind.TransactOpts, documentHash [32]byte, ipfsCID string, publicKey []byte) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "registerDocument", documentHash, ipfsCID, publicKey)
}

// RegisterDocument is a paid mutator transaction binding the contract method 0x6a243d5b.
//
// Solidity: function registerDocument(bytes32 documentHash, string ipfsCID, bytes publicKey) returns()
func (_Blockchain *BlockchainSession) RegisterDocument(documentHash [32]byte, ipfsCID string, publicKey []byte) (*types.Transaction, error) {
	return _Blockchain.Contract.RegisterDocument(&_Blockchain.TransactOpts, documentHash, ipfsCID, publicKey)
}

// RegisterDocument is a paid mutator transaction binding the contract method 0x6a243d5b.
//
// Solidity: function registerDocument(bytes32 documentHash, string ipfsCID, bytes publicKey) returns()
func (_Blockchain *BlockchainTransactorSession) RegisterDocument(documentHash [32]byte, ipfsCID string, publicKey []byte) (*types.Transaction, error) {
	return _Blockchain.Contract.RegisterDocument(&_Blockchain.TransactOpts, documentHash, ipfsCID, publicKey)
}

// BlockchainDocumentRegisteredIterator is returned from FilterDocumentRegistered and is used to iterate over the raw logs and unpacked data for DocumentRegistered events raised by the Blockchain contract.
type BlockchainDocumentRegisteredIterator struct {
	Event *BlockchainDocumentRegistered // Event containing the contract specifics and raw log

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
func (it *BlockchainDocumentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainDocumentRegistered)
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
		it.Event = new(BlockchainDocumentRegistered)
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
func (it *BlockchainDocumentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainDocumentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainDocumentRegistered represents a DocumentRegistered event raised by the Blockchain contract.
type BlockchainDocumentRegistered struct {
	DocumentHash [32]byte
	Owner        common.Address
	IpfsCID      string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDocumentRegistered is a free log retrieval operation binding the contract event 0xdba629cc54aad70846b805286cd90bf0920d3b29656d9b042814d2a2dedf4acb.
//
// Solidity: event DocumentRegistered(bytes32 indexed documentHash, address indexed owner, string ipfsCID)
func (_Blockchain *BlockchainFilterer) FilterDocumentRegistered(opts *bind.FilterOpts, documentHash [][32]byte, owner []common.Address) (*BlockchainDocumentRegisteredIterator, error) {

	var documentHashRule []interface{}
	for _, documentHashItem := range documentHash {
		documentHashRule = append(documentHashRule, documentHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "DocumentRegistered", documentHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainDocumentRegisteredIterator{contract: _Blockchain.contract, event: "DocumentRegistered", logs: logs, sub: sub}, nil
}

// WatchDocumentRegistered is a free log subscription operation binding the contract event 0xdba629cc54aad70846b805286cd90bf0920d3b29656d9b042814d2a2dedf4acb.
//
// Solidity: event DocumentRegistered(bytes32 indexed documentHash, address indexed owner, string ipfsCID)
func (_Blockchain *BlockchainFilterer) WatchDocumentRegistered(opts *bind.WatchOpts, sink chan<- *BlockchainDocumentRegistered, documentHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var documentHashRule []interface{}
	for _, documentHashItem := range documentHash {
		documentHashRule = append(documentHashRule, documentHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "DocumentRegistered", documentHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainDocumentRegistered)
				if err := _Blockchain.contract.UnpackLog(event, "DocumentRegistered", log); err != nil {
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

// ParseDocumentRegistered is a log parse operation binding the contract event 0xdba629cc54aad70846b805286cd90bf0920d3b29656d9b042814d2a2dedf4acb.
//
// Solidity: event DocumentRegistered(bytes32 indexed documentHash, address indexed owner, string ipfsCID)
func (_Blockchain *BlockchainFilterer) ParseDocumentRegistered(log types.Log) (*BlockchainDocumentRegistered, error) {
	event := new(BlockchainDocumentRegistered)
	if err := _Blockchain.contract.UnpackLog(event, "DocumentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainDocumentVerifiedIterator is returned from FilterDocumentVerified and is used to iterate over the raw logs and unpacked data for DocumentVerified events raised by the Blockchain contract.
type BlockchainDocumentVerifiedIterator struct {
	Event *BlockchainDocumentVerified // Event containing the contract specifics and raw log

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
func (it *BlockchainDocumentVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainDocumentVerified)
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
		it.Event = new(BlockchainDocumentVerified)
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
func (it *BlockchainDocumentVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainDocumentVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainDocumentVerified represents a DocumentVerified event raised by the Blockchain contract.
type BlockchainDocumentVerified struct {
	DocumentHash [32]byte
	Verifier     common.Address
	Verified     bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDocumentVerified is a free log retrieval operation binding the contract event 0x6894ebbaf406311560473c5438f7c75d70dc611e024968e8f8596dbd034c5084.
//
// Solidity: event DocumentVerified(bytes32 indexed documentHash, address indexed verifier, bool verified)
func (_Blockchain *BlockchainFilterer) FilterDocumentVerified(opts *bind.FilterOpts, documentHash [][32]byte, verifier []common.Address) (*BlockchainDocumentVerifiedIterator, error) {

	var documentHashRule []interface{}
	for _, documentHashItem := range documentHash {
		documentHashRule = append(documentHashRule, documentHashItem)
	}
	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "DocumentVerified", documentHashRule, verifierRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainDocumentVerifiedIterator{contract: _Blockchain.contract, event: "DocumentVerified", logs: logs, sub: sub}, nil
}

// WatchDocumentVerified is a free log subscription operation binding the contract event 0x6894ebbaf406311560473c5438f7c75d70dc611e024968e8f8596dbd034c5084.
//
// Solidity: event DocumentVerified(bytes32 indexed documentHash, address indexed verifier, bool verified)
func (_Blockchain *BlockchainFilterer) WatchDocumentVerified(opts *bind.WatchOpts, sink chan<- *BlockchainDocumentVerified, documentHash [][32]byte, verifier []common.Address) (event.Subscription, error) {

	var documentHashRule []interface{}
	for _, documentHashItem := range documentHash {
		documentHashRule = append(documentHashRule, documentHashItem)
	}
	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "DocumentVerified", documentHashRule, verifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainDocumentVerified)
				if err := _Blockchain.contract.UnpackLog(event, "DocumentVerified", log); err != nil {
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

// ParseDocumentVerified is a log parse operation binding the contract event 0x6894ebbaf406311560473c5438f7c75d70dc611e024968e8f8596dbd034c5084.
//
// Solidity: event DocumentVerified(bytes32 indexed documentHash, address indexed verifier, bool verified)
func (_Blockchain *BlockchainFilterer) ParseDocumentVerified(log types.Log) (*BlockchainDocumentVerified, error) {
	event := new(BlockchainDocumentVerified)
	if err := _Blockchain.contract.UnpackLog(event, "DocumentVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
