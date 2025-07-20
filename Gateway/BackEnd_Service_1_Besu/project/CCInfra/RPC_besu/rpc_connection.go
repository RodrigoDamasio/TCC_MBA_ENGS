package Infra_rpc

import (
	"context"
	"math/big"

	contract "project/CCEntitys/Contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Blockchain struct {
	Client           *ethclient.Client
	Auth             *bind.TransactOpts
	ContractInstance *contract.Contract
}

// Cria CallOpts pronto para consulta, podendo customizar address e contexto
func (b *Blockchain) NewCallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{
		From:    b.Auth.From, // padrão: mesmo endereço do Auth
		Context: ctx,
	}
}

func NewBlockchain() (*Blockchain, error) {

	// Carrega as configurações do arquivo YAML

	config, err := LoadRPCConfig("CCInfra/RPC_besu/rpc_config.yaml")
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(config.GetRPCURL())
	if err != nil {
		return nil, err
	}

	privateKeyHex := config.GetPrivateKey()
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.GetChainID()))
	if err != nil {
		return nil, err
	}

	contractInstance, err := contract.NewContract(common.HexToAddress(config.GetContractAddress()), client)
	if err != nil {
		return nil, err
	}

	return &Blockchain{
		Client:           client,
		Auth:             auth,
		ContractInstance: contractInstance,
	}, nil
}

// Cria uma nova instância de Blockchain

func NewBlockchainWithConfig(configPath string) (*Blockchain, error) {

	// Carrega as configurações do arquivo YAML

	config, err := LoadRPCConfig(configPath)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(config.GetRPCURL())
	if err != nil {
		return nil, err
	}

	privateKeyHex := config.GetPrivateKey()
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.GetChainID()))
	if err != nil {
		return nil, err
	}

	contractInstance, err := contract.NewContract(common.HexToAddress(config.GetContractAddress()), client)
	if err != nil {
		return nil, err
	}

	return &Blockchain{
		Client:           client,
		Auth:             auth,
		ContractInstance: contractInstance,
	}, nil

}
