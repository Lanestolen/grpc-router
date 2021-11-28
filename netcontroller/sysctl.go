package netcontroller

import "fmt"

// IPService wraps the "IP" commands
type SysService struct {
	c *NetController
}

func (ipc *SysService) Enable(iface string) error {
	enableString := fmt.Sprintf("wg-quick@%s", iface)
	cmds := []string{"enable", enableString}
	_, err := ipc.exec(cmds...)
	return err
}

// exec executes an ExecFunc using 'ip'.
func (ipc *SysService) exec(args ...string) ([]byte, error) {
	return ipc.c.exec("systemctl", args...)
}
