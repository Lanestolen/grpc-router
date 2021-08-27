package main

import (
	"net"
	"os"
	"strconv"

	"github.com/lanestolen/grpc-router/config"
	"github.com/lanestolen/grpc-router/dhcp"
	"github.com/lanestolen/grpc-router/netcontroller"
	"github.com/lanestolen/grpc-router/wireguard"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	os.Setenv("CONFIG_PATH", "/home/lanestolen/.config/grpc-router/config.yml")
	conf, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("couldn't parse config")
	}

	port := strconv.FormatUint(uint64(conf.ServiceConfig.Domain.Port), 10)

	grpcServer := grpc.NewServer()
	netcontroller.OpenNetController()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal().Err(err).Str("port", port).Msg("failed to listen on port")
	}

	dhcp.NewDHCPServer(grpcServer)
	wireguard.NewWireguardServer(grpcServer, conf)

	grpcServer.Serve(lis)
}
