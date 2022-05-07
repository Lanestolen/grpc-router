package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/lanestolen/grpc-router/dhcp/proto"
	wg "github.com/lanestolen/grpc-router/wireguard/proto"
	"google.golang.org/grpc"
)

type Creds struct {
	Token    string
	Insecure bool
}

func (c Creds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"token": string(c.Token),
	}, nil
}

func (c Creds) RequireTransportSecurity() bool {
	return !c.Insecure
}

func main() {
	// change the endpoint address with your instance ip
	endpointAddress := "localhost"
	var conn *grpc.ClientConn
	// wg is AUTH_KEY from vpn/auth.go
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"wg": "deneme",
	})

	tokenString, err := token.SignedString([]byte("test"))
	if err != nil {
		fmt.Println("Error creating the token")
	}

	authCreds := Creds{Token: tokenString}
	dialOpts := []grpc.DialOption{}
	authCreds.Insecure = true
	dialOpts = append(dialOpts,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(authCreds))

	conn, err = grpc.Dial(fmt.Sprintf("%s:5353", endpointAddress), dialOpts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	if err == nil {
		fmt.Printf("Client is connected successfully !")
	}
	defer conn.Close()

	client := proto.NewDHCPClient(conn)
	ctx := context.TODO()

	var Networks []*proto.Network
	net := proto.Network{
		Router:  "10.10.11.1",
		Min:     "10.10.11.6",
		Max:     "10.10.11.255",
		Network: "10.10.11.0",
		DnsServer: "10.10.11.3",

	}
	Networks = append(Networks, &net)

	client.StopDHCP(ctx, &proto.StopReq{})
	resp, err := client.StartDHCP(ctx, &proto.StartReq{Networks: Networks})
	fmt.Println(resp, err)

	wgclient := wg.NewWireguardClient(conn)
	_, err = wgclient.InitializeI(ctx, &wg.IReq{
		Address:            "10.2.11.1/24", // this should be randomized and should not collide with lab subnet like 124.5.6.0/24
		ListenPort:         51820,          // this should be randomized and should not collide with any used ports by host
		SaveConfig:         true,
		Eth:                "eth0",
		IName:              "wg",
		DownInterfacesFile: "/etc/network/downinterfaces",
	})
	if err != nil {
		fmt.Println("Error while initializing interface")
		panic(err)
	}

	fmt.Println("Getting server public key...")
	serverPubKey, err := wgclient.GetPublicKey(ctx, &wg.PubKeyReq{PubKeyName: "wg", PrivKeyName: "wg"})
	if err != nil {
		panic(err)
	}

	wgresp, err := wgclient.GenPublicKey(ctx, &wg.PubKeyReq{
		PubKeyName:  "client1",
		PrivKeyName: "client1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf(wgresp.Message)

	publicKey, err := wgclient.GetPublicKey(ctx, &wg.PubKeyReq{PubKeyName: "client1"})
	if err != nil {
		panic(err)
	}
	clientPrivKey, err := wgclient.GetPrivateKey(ctx, &wg.PrivKeyReq{PrivateKeyName: "client1"})
	if err != nil {
		panic(err)
	}

	peerIP := "10.2.11.3/32"

	_, err = wgclient.AddPeer(ctx, &wg.AddPReq{
		Nic:        "wg",
		AllowedIPs: peerIP,
		PublicKey:  publicKey.Message,
	})
	if err != nil {
		panic(err)
	}

	allowedIps := "10.142.233.1/24"
	clientConfig := fmt.Sprintf(
		`[Interface]
Address = %s
PrivateKey = %s
DNS = 1.1.1.1
MTU = 1500
[Peer]
PublicKey = %s
AllowedIps = %s
Endpoint =  %s
PersistentKeepalive = 25`, peerIP, clientPrivKey.Message, serverPubKey.Message, allowedIps, fmt.Sprintf("%s:51820", endpointAddress))

	fmt.Printf("Client Config \n %s ", clientConfig)
}
