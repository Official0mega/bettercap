package firewall

import (
	"fmt"
	"strings"

	"github.com/evilsocket/bettercap-ng/core"
	"github.com/evilsocket/bettercap-ng/net"
)

type WindowsFirewall struct {
	iface        *net.Endpoint
	forwarding   bool
	redirections map[string]*Redirection
}

func Make(iface *net.Endpoint) FirewallManager {
	firewall := &WindowsFirewall{
		iface:        iface,
		forwarding:   false,
		redirections: make(map[string]*Redirection, 0),
	}

	firewall.forwarding = firewall.IsForwardingEnabled()

	return firewall
}

func (f WindowsFirewall) IsForwardingEnabled() bool {
	if out, err := core.Exec("netsh", []string{"interface", "ipv4", "dump"}); err != nil {
		fmt.Printf("%s\n", err)
		return false
	} else {
		return strings.Contains(out, "forwarding=enabled")
	}
}

func (f WindowsFirewall) EnableForwarding(enabled bool) error {
	fmt.Printf("iface idx=%d\n", f.iface.Index)
	return fmt.Errorf("Not implemented")
}

func (f *WindowsFirewall) EnableRedirection(r *Redirection, enabled bool) error {
	return fmt.Errorf("Not implemented")
}

func (f WindowsFirewall) Restore() {
	for _, r := range f.redirections {
		if err := f.EnableRedirection(r, false); err != nil {
			fmt.Printf("%s", err)
		}
	}

	if err := f.EnableForwarding(f.forwarding); err != nil {
		fmt.Printf("%s", err)
	}
}