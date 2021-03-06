package drivers

import (
	"errors"
	"path/filepath"
)

const (
	DefaultSSHUser          = "root"
	DefaultSSHPort          = 22
	DefaultEngineInstallURL = "https://get.docker.com"
	DefaultMaxAttempt       = 60
)

// BaseDriver - Embed this struct into drivers to provide the common set
// of fields and functions.
type BaseDriver struct {
	IPAddress      string
	MachineName    string
	SSHUser        string
	SSHPort        int
	SSHKeyPath     string
	StorePath      string
	SwarmMaster    bool
	SwarmHost      string
	SwarmDiscovery string
	MaxAttempt     int
}

// DriverName returns the name of the driver
func (d *BaseDriver) DriverName() string {
	return "unknown"
}

// GetMachineName returns the machine name
func (d *BaseDriver) GetMachineName() string {
	return d.MachineName
}

// GetMaxAttempt returns the maximum attempt
func (d *BaseDriver) GetMaxAttempt() int {
	if d.MaxAttempt == 0 {
		d.MaxAttempt = DefaultMaxAttempt
	}
	return d.MaxAttempt
}

// GetIP returns the ip
func (d *BaseDriver) GetIP() (string, error) {
	if d.IPAddress == "" {
		return "", errors.New("IP address is not set")
	}
	return d.IPAddress, nil
}

// GetSSHKeyPath returns the ssh key path
func (d *BaseDriver) GetSSHKeyPath() string {
	if d.SSHKeyPath == "" {
		d.SSHKeyPath = d.ResolveStorePath("id_rsa")
	}
	return d.SSHKeyPath
}

// GetSSHPort returns the ssh port, 22 if not specified
func (d *BaseDriver) GetSSHPort() (int, error) {
	if d.SSHPort == 0 {
		d.SSHPort = DefaultSSHPort
	}

	return d.SSHPort, nil
}

// GetSSHUsername returns the ssh user name, root if not specified
func (d *BaseDriver) GetSSHUsername() string {
	if d.SSHUser == "" {
		d.SSHUser = DefaultSSHUser
	}
	return d.SSHUser
}

// PreCreateCheck is called to enforce pre-creation steps
func (d *BaseDriver) PreCreateCheck() error {
	return nil
}

// ResolveStorePath returns the store path where the machine is
func (d *BaseDriver) ResolveStorePath(file string) string {
	return filepath.Join(d.StorePath, "machines", d.MachineName, file)
}

// SetExtraConfigFromFlags configures the driver for extra params as swarm or max attempts
func (d *BaseDriver) SetExtraConfigFromFlags(flags DriverOptions) {
	d.SwarmMaster = flags.Bool("swarm-master")
	d.SwarmHost = flags.String("swarm-host")
	d.SwarmDiscovery = flags.String("swarm-discovery")
	d.MaxAttempt = flags.Int("max-attempts")
}

func EngineInstallURLFlagSet(flags DriverOptions) bool {
	engineInstallURLFlag := flags.String("engine-install-url")
	return engineInstallURLFlag != DefaultEngineInstallURL && engineInstallURLFlag != ""
}
