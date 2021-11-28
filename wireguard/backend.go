package wireguard

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/lanestolen/grpc-router/config"
	"github.com/lanestolen/grpc-router/netcontroller"
	"github.com/lanestolen/grpc-router/wireguard/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type wireguard struct {
	config config.Config
	proto.UnimplementedWireguardServer
}

func NewWireguardServer(s *grpc.Server, conf config.Config) {
	proto.RegisterWireguardServer(s, &wireguard{config: conf})
}

// InitializeI creates interface configuration and make it UP.
func (w *wireguard) InitializeI(ctx context.Context, r *proto.IReq) (*proto.IResp, error) {

	log.Info().Msgf("Initializing interface for %s ", r.IName)
	privKey, err := generatePrivateKey(w.config.WireGuardDir + r.IName + "_priv")
	if err != nil {
		log.Error().Err(err).Str("keyFile", privKey).Msg("Problem generating the privatekey")
		return &proto.IResp{}, err
	}
	log.Info().Msgf("Private key is generated %s with name %s", w.config.WireGuardDir, r.IName)
	if err := generatePublicKey(ctx, w.config.WireGuardDir+r.IName+"_priv", w.config.WireGuardDir+r.IName+"_pub"); err != nil {
		return &proto.IResp{}, err
	}

	wgI := Interface{
		address:    r.Address,
		listenPort: r.ListenPort,
		privateKey: privKey,
		eth:        r.Eth,
		saveConfig: r.SaveConfig,
		iName:      r.IName,
	}

	out, err := w.genInterfaceConf(wgI, w.config.WireGuardDir)
	if err != nil {
		log.Error().Err(err).Str("interface", wgI.iName).Msg("Problem in configuration of the interface")

		return &proto.IResp{Message: out}, fmt.Errorf("Problem in configuration of the interface -- %v", err)
	}

	out, err = upDown(r.IName, "up")

	if err != nil {

		log.Error().Err(err).Str("interface", out).Msg("Problem in making the interface UP")
		return &proto.IResp{Message: out}, fmt.Errorf("PROBLEM IN THE FUNCTION upDown -- %v", err)
	}

	netcontroller.Controller.Sysctl.Enable(r.IName)

	log.Debug().Str("Address: ", r.Address).
		Uint32("ListenPort: ", r.ListenPort).
		Str("Ethernet I: ", r.Eth).
		Str("PrivateKey: ", r.PrivateKey).
		Bool("SaveConfig", r.SaveConfig).Msgf("Interface %s created and it is up", r.IName)

	return &proto.IResp{Message: out}, nil
}

// AddPeer adds peer to given wireguard interface
func (w *wireguard) AddPeer(ctx context.Context, r *proto.AddPReq) (*proto.AddPResp, error) {

	out, err := addPeer(r.Nic, r.PublicKey, r.AllowedIPs)
	out = strings.Replace(out, "/n", "", -1)
	if err != nil {
		return &proto.AddPResp{Message: out}, err
	}
	log.Info().Msgf("Peer with public key: { %s } is added to interface: { %s } from allowed-ips: { %s }", r.PublicKey, r.Nic, r.AllowedIPs)
	if err := saveConfig(r.Nic); err != nil {
		return &proto.AddPResp{Message: "error"}, fmt.Errorf("problem saving config")
	}
	return &proto.AddPResp{Message: out}, nil
}

// DelPeer deletes peer from given wireguard interface
func (w *wireguard) DelPeer(ctx context.Context, r *proto.DelPReq) (*proto.DelPResp, error) {
	out, err := removePeer(r.PeerPublicKey, r.IpAddress)
	if err != nil {
		return &proto.DelPResp{Message: out}, err
	}
	log.Info().Msgf("Peer with public key: { %s } is deleted from ip-address: { %s }", r.PeerPublicKey, r.IpAddress)
	return &proto.DelPResp{Message: out}, nil
}

// GetNICInfo returns general information about given wireguard interface
func (w *wireguard) GetNICInfo(ctx context.Context, r *proto.NICInfoReq) (*proto.NICInfoResp, error) {
	out, err := nicInfo(r.Interface)
	if err != nil {
		return &proto.NICInfoResp{Message: string(out)}, err
	}
	log.Debug().Msgf("NIC Information for { %s } is printed ", r.Interface)
	return &proto.NICInfoResp{Message: string(out)}, nil
}

func (w *wireguard) GetPeerStatus(ctx context.Context, in *proto.PeerStatusReq) (*proto.PeerStatusResp, error) {
	publicKey := in.PublicKey
	nicName := in.NicName
	isConnected, err := checkStatus(nicName, publicKey)
	if err != nil {
		return &proto.PeerStatusResp{Status: false}, err
	}
	return &proto.PeerStatusResp{Status: isConnected}, nil
}

// ManageNIC is managing (up & down) given wireguard interface
func (w *wireguard) ManageNIC(ctx context.Context, r *proto.ManageNICReq) (*proto.ManageNICResp, error) {
	out, err := upDown(r.Nic, r.Cmd)
	if err != nil {
		return &proto.ManageNICResp{Message: string(out)}, err
	}
	log.Info().Msgf("ManageNIC: interface %s is called to be %s", r.Nic, r.Cmd)
	return &proto.ManageNICResp{Message: out}, nil
}

// wg show <interface-name>
// if interface-name is not provided by user list for all.
func (w *wireguard) ListPeers(ctx context.Context, r *proto.ListPeersReq) (*proto.ListPeersResp, error) {
	out, err := listPeers(r.Nicname)
	if err != nil {
		log.Printf("Error in listing peers in gRPC %v", err)
		return &proto.ListPeersResp{}, err
	}
	log.Info().Msgf("ListPeers: listing peers for %s interface", r.Nicname)
	return &proto.ListPeersResp{Response: out}, nil
}

// GenPrivateKey generates PrivateKey for wireguard interface
func (w *wireguard) GenPrivateKey(ctx context.Context, r *proto.PrivKeyReq) (*proto.PrivKeyResp, error) {

	_, err := generatePrivateKey(w.config.WireGuardDir + r.PrivateKeyName + "_priv")
	if err != nil {
		return &proto.PrivKeyResp{}, err
	}
	log.Info().Msgf("GenPrivateKey is called to generate new private key with filename %s", r.PrivateKeyName)
	return &proto.PrivKeyResp{Message: "Private Key is created with name " + w.config.WireGuardDir + r.PrivateKeyName}, nil
}

// GenPublicKey generates PublicKey for wireguard interface
func (w *wireguard) GenPublicKey(ctx context.Context, r *proto.PubKeyReq) (*proto.PubKeyResp, error) {
	// check whether private key exists or not, if not generate one
	if _, err := os.Stat(w.config.WireGuardDir + r.PrivKeyName + "_pub"); os.IsNotExist(err) {
		fmt.Printf("PrivateKeyFile is not exists, creating one ... %s\n", r.PrivKeyName)
		_, err := generatePrivateKey(w.config.WireGuardDir + r.PrivKeyName + "_priv")
		if err != nil {
			return &proto.PubKeyResp{Message: "Error"}, fmt.Errorf("error in generation of private key %v", err)
		}
	}

	if err := generatePublicKey(ctx, w.config.WireGuardDir+r.PrivKeyName+"_priv", w.config.WireGuardDir+r.PubKeyName+"_pub"); err != nil {
		return &proto.PubKeyResp{Message: "Error"}, fmt.Errorf("error in generation of public key %v", err)
	}
	return &proto.PubKeyResp{Message: "Public key is generated with " + w.config.WireGuardDir + r.PubKeyName + " name"}, nil
}

// GetPublicKey returns content of given PublicKey
func (w *wireguard) GetPublicKey(ctx context.Context, req *proto.PubKeyReq) (*proto.PubKeyResp, error) {

	out, err := w.getContent(req.PubKeyName + "_pub")
	if err != nil {
		return &proto.PubKeyResp{}, err
	}

	out = strings.Replace(out, "\n", "", 1)
	return &proto.PubKeyResp{Message: out}, nil
}

// GetPrivateKey returns content of given PrivateKey
func (w *wireguard) GetPrivateKey(ctx context.Context, req *proto.PrivKeyReq) (*proto.PrivKeyResp, error) {

	log.Info().Msgf("Getting private key information for team %s", req.PrivateKeyName)
	out, err := w.getContent(req.PrivateKeyName + "_priv")
	if err != nil {
		return &proto.PrivKeyResp{}, err
	}
	return &proto.PrivKeyResp{Message: out}, nil
}
