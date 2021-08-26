package dhcp

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"text/template"

	"github.com/lanestolen/grpc-router/dhcp/proto"
	"github.com/lanestolen/grpc-router/netcontroller"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	//go:embed dhcpd.conf.tmpl
	dhcpTemplate embed.FS
)

type dhcpServer struct {
	proto.UnimplementedDHCPServer
}

func NewDHCPServer(s *grpc.Server) {
	proto.RegisterDHCPServer(s, &dhcpServer{})
}

func (d *dhcpServer) StartDHCP(ctx context.Context, r *proto.StartReq) (*proto.Resp, error) {
	log.Info().Msg("starting dhcp server")

	if err := netcontroller.Controller.DHCP.StopDHCP(); err != nil {
		return &proto.Resp{Success: false, Message: fmt.Sprintf("error while stopping dhcp: %v", err)}, err
	}

	if err := createDHCPFile(r); err != nil {
		return &proto.Resp{Success: false, Message: fmt.Sprintf("error while creating dhcp config: %v", err)}, err
	}

	if err := setupIfaces(r.Networks); err != nil {
		return &proto.Resp{Success: false, Message: fmt.Sprintf("error defining dhcp interfaces: %v", err)}, err
	}

	if err := netcontroller.Controller.DHCP.StartDHCP(); err != nil {
		return &proto.Resp{Success: false, Message: fmt.Sprintf("error while starting dhcp: %v", err)}, err
	}

	return &proto.Resp{Success: true, Message: "started dhcp server"}, nil
}

func (d *dhcpServer) StopDHCP(ctx context.Context, r *proto.StopReq) (*proto.Resp, error) {
	log.Info().Msg("stopping dhcp server")

	if err := netcontroller.Controller.DHCP.StopDHCP(); err != nil {
		return &proto.Resp{Success: false, Message: fmt.Sprintf("error while stopping dhcp: %v", err)}, err
	}

	return &proto.Resp{Success: true, Message: "started dhcp server"}, nil
}

func createDHCPConfig(nets *proto.StartReq) ([]byte, error) {
	var tpl bytes.Buffer
	tmpl, err := template.ParseFS(dhcpTemplate, "dhcpd.conf.tmpl")
	if err != nil {
		return []byte{}, err
	}
	tmpl.Execute(&tpl, nets)
	return tpl.Bytes(), nil
}

func createDHCPFile(nets *proto.StartReq) error {
	content, err := createDHCPConfig(nets)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("/etc/dhcp/dhcpd.conf", content, 0644); err != nil {
		return err
	}

	return nil
}

func createIfaceConf(ifaces []string) error {
	setupstring := "INTERFACES=\"" + strings.Join(ifaces, " ") + "\""
	err := ioutil.WriteFile("/etc/default/isc-dhcp-server", []byte(setupstring), 0644)
	if err != nil {
		return err
	}
	return nil
}

func setupIfaces(o []*proto.Network) error {
	var ifaces []string
	var usedifaces []string
	infs, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, f := range infs {
		if f.Name == "lo" || f.Name == "enp0s3" || f.Name == "docker0" {
			continue
		}
		ifaces = append(ifaces, f.Name)
	}

	for i, nets := range o {
		log.Debug().Str("iface", ifaces[i]).Str("ip", fmt.Sprintf("%s/24", nets.Router)).Msg("setting up interface")
		if err := netcontroller.Controller.IfConfig.SetIP(ifaces[i], fmt.Sprintf("%s/24", nets.Router)); err != nil {
			return err
		}
		usedifaces = append(usedifaces, ifaces[i])
	}
	if err := createIfaceConf(usedifaces); err != nil {
		return err
	}
	return nil
}
