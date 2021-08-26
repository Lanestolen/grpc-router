package netcontroller

// IfConfigService wraps the "ifconfig" command
type IfConfigService struct {
	c *NetController
}

// Up sets the given interface up
func (ipc *IfConfigService) Up(iface string) error {
	cmds := []string{iface, "up"}
	_, err := ipc.exec(cmds...)
	return err
}

// Down takes the given interface down
func (ipc *IfConfigService) Down(iface string) error {
	cmds := []string{iface, "down"}
	_, err := ipc.exec(cmds...)
	return err
}

// SetIP sets the ipaddress of the given interface.
// The IP should be given in the CIDR format
func (ipc *IfConfigService) SetIP(iface, ip string) error {
	cmds := []string{iface, ip}
	_, err := ipc.exec(cmds...)
	return err
}

// exec executes an ExecFunc using 'ifconfig'.
func (ipc *IfConfigService) exec(args ...string) ([]byte, error) {

	return ipc.c.exec("ifconfig", args...)
}
