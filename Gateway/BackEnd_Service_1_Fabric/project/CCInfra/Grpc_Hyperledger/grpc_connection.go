package Infra_grpc

import (
	"crypto/x509"
	"log"
	"os"
	"path"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/hash"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var ContractPointer *client.Contract

func init() {

	ContractPointer = GetContract()

}

// GetContract retorna o contrato para interagir com o Hyperledger Fabric
func GetContract() *client.Contract {

	// Carrega as configurações
	err := LoadConfig("")
	if err != nil {
		log.Fatalf("erro ao carregar configurações: %w", err)
	}

	clientConnection := newGrpcConnection()
	id := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithHash(hash.SHA256),
		client.WithClientConnection(clientConnection),
	)
	if err != nil {
		panic(err)
	}

	network := gw.GetNetwork(Config.GetChannelName())
	return network.GetContract(Config.GetChaincodeName())

}

func newGrpcConnection() *grpc.ClientConn {

	certificatePEM, err := os.ReadFile(Config.GetTLSCertPath())
	if err != nil {
		log.Fatalf("failed to read TLS certificate file: %w", err)
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		log.Fatalf("failed to parse TLS certificate: %w", err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, Config.GetGatewayPeer())

	connection, err := grpc.Dial(Config.GetPeerEndpoint(), grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		log.Fatalf("failed to create gRPC connection: %w", err)
	}

	return connection

}

func newIdentity() *identity.X509Identity {

	certificatePEM, err := readFirstFile(Config.GetCertPath())
	if err != nil {
		log.Fatalf("failed to read certificate file: %w", err)
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		log.Fatalf("failed to parse certificate: %w", err)
	}

	id, err := identity.NewX509Identity(Config.GetMSPID(), certificate)
	if err != nil {
		log.Fatalf("failed to create identity: %w", err)
	}

	return id

}

func newSign() identity.Sign {

	privateKeyPEM, err := readFirstFile(Config.GetKeyPath())
	if err != nil {
		log.Fatalf("failed to read private key file: %w", err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		log.Fatalf("failed to parse private key: %w", err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		log.Fatalf("failed to create sign: %w", err)
	}

	return sign

}

func readFirstFile(dirPath string) ([]byte, error) {

	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	fileNames, err := dir.Readdirnames(1)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path.Join(dirPath, fileNames[0]))

}
