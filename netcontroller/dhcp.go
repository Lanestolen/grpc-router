package netcontroller

type DHCPService struct {
	c *NetController
}

func (d *DHCPService) StartDHCP() error {
	_, err := d.c.exec("systemctl", "start", "isc-dhcp-server")
	return err
}

func (d *DHCPService) StopDHCP() error {
	_, err := d.c.exec("systemctl", "stop", "isc-dhcp-server")
	return err
}

func (d *DHCPService) UpdateConf(filepath string) error {
	_, err := d.c.exec("dhcpd", "-cf", filepath)
	return err
}
