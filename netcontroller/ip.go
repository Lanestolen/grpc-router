package netcontroller

// IPService wraps the "IP" commands
type IPService struct {
	c *NetController
}

// AddTunTap function adds a tap interface with the given mode
// i.e:  ip tuntap add tap0 mode tap
func (ipc *IPService) AddTunTap(tap, mode string) error {
	cmds := []string{"tuntap", "add", tap, "mode", mode}
	//_, err := ipc.exec(fmt.Sprintf("tuntap add %s mode %s", tap, mode))
	_, err := ipc.exec(cmds...)
	return err
}

// DelTuntap deletes tap with given name and mode
func (ipc *IPService) DelTuntap(tap, mode string) error {
	cmds := []string{"tuntap", "del", tap, "mode", mode}
	_, err := ipc.exec(cmds...)

	return err
}

// exec executes an ExecFunc using 'ip'.
func (ipc *IPService) exec(args ...string) ([]byte, error) {
	return ipc.c.exec("ip", args...)
}
