package proxmoxve

import (
	"github.com/rancher/machine/libmachine/drivers"
	"github.com/rancher/machine/libmachine/mcnflag"
	"github.com/rancher/machine/libmachine/state"
)

// Driver for Proxmox VE
type Driver struct {
	*drivers.BaseDriver
	driver *ProxmoxVE

	// Top-level strategy for proisioning a new node
	ProvisionStrategy string

	// Basic Authentication for Proxmox VE
	Host     string // Host to connect to
	Node     string // optional, node to create VM on, host used if omitted but must match internal node name
	User     string // username
	Password string // password
	Realm    string // realm, e.g. pam, pve, etc.

	// File to load as boot image RancherOS/Boot2Docker
	ImageFile string // in the format <storagename>:iso/<filename>.iso

	Pool            string // pool to add the VM to (necessary for users with only pool permission)
	Storage         string // internal PVE storage name
	StorageType     string // Type of the storage (currently QCOW2 and RAW)
	DiskSize        string // disk size in GB
	Memory          int    // memory in GB
	StorageFilename string
	Onboot          string // Specifies whether a VM will be started during system bootup.
	Protection      string // Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.
	Citype          string // Specifies the cloud-init configuration format.
	NUMA            string // Enable/disable NUMA

	CiEnabled string

	NetModel    string // Net Interface Model, [e1000, virtio, realtek, etc...]
	NetFirewall string // Enable/disable firewall
	NetMtu      string // set nic MTU
	NetBridge   string // bridge applied to network interface
	NetVlanTag  int    // vlan tag

	ScsiController string
	ScsiAttributes string

	VMID          string // VM ID only filled by create()
	VMIDRange     string // acceptable range of VMIDs
	VMUUID        string // UUID to confirm
	CloneVMID     string // VM ID to clone
	CloneFull     int    // Make a full (detached) clone from parent (defaults to true if VMID is not a template, otherwise false)
	GuestUsername string // user to log into the guest OS to copy the public key
	GuestPassword string // password to log into the guest OS to copy the public key
	GuestSSHPort  int    // ssh port to log into the guest OS to copy the public key
	CPU           string // Emulated CPU type.
	CPUSockets    string // The number of cpu sockets.
	CPUCores      string // The number of cores per socket.
	driverDebug   bool   // driver debugging
	restyDebug    bool   // enable resty debugging
}

// Create implements drivers.Driver
func (d *Driver) Create() error {
	panic("unimplemented")
}

// GetCreateFlags implements drivers.Driver
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	panic("unimplemented")
}

// GetSSHHostname implements drivers.Driver
func (d *Driver) GetSSHHostname() (string, error) {
	panic("unimplemented")
}

// GetState implements drivers.Driver
func (d *Driver) GetState() (state.State, error) {
	panic("unimplemented")
}

// GetURL implements drivers.Driver
func (*Driver) GetURL() (string, error) {
	panic("unimplemented")
}

// Kill implements drivers.Driver
func (d *Driver) Kill() error {
	panic("unimplemented")
}

// Remove implements drivers.Driver
func (*Driver) Remove() error {
	panic("unimplemented")
}

// Restart implements drivers.Driver
func (d *Driver) Restart() error {
	panic("unimplemented")
}

// SetConfigFromFlags implements drivers.Driver
func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	panic("unimplemented")
}

// Start implements drivers.Driver
func (*Driver) Start() error {
	panic("unimplemented")
}

// Stop implements drivers.Driver
func (d *Driver) Stop() error {
	panic("unimplemented")
}

// Upgrade implements drivers.Driver
func (d *Driver) Upgrade() error {
	panic("unimplemented")
}

// Validate implements drivers.Driver
func (d *Driver) Validate() error {
	panic("unimplemented")
}

// GetIP implements drivers.Driver
func (d *Driver) GetIP() (string, error) {
	panic("unimplemented")
}

// GetSSHKeyPath implements drivers.Driver
func (d *Driver) GetSSHKeyPath() string {
	panic("unimplemented")
}

// GetSSHUsername implements drivers.Driver
func (d *Driver) GetSSHUsername() string {
	panic("unimplemented")
}

// DriverName implements drivers.Driver
func (d *Driver) DriverName() string {
	panic("unimplemented")
}

// PreCreateCheck implements drivers.Driver
func (d *Driver) PreCreateCheck() error {
	panic("unimplemented")
}

// GetMachineName implements drivers.Driver
func (d *Driver) GetMachineName() string {
	panic("unimplemented")
}

// GetSSHCommand implements drivers.Driver
func (d *Driver) GetSSHCommand(args ...string) (string, error) {
	panic("unimplemented")
}
