package immortal

// Control interface
type Control interface {
	GetStatus(socket string) (*Status, error)
	SendSignal(socket, signal string) (*SignalResponse, error)
	FindServices(dir string) ([]*ServiceStatus, error)
	PurgeServices(dir string) error
	Run(command string) ([]byte, error)
}

// ServiceStatus struct
type ServiceStatus struct {
	Name           string
	Socket         string
	Status         *Status
	SignalResponse *SignalResponse
}

// Controller implements Control
type Controller struct{}

// GetStatus returns service status in json format
func (c *Controller) GetStatus(socket string) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SendSignal send signal to process
func (c *Controller) SendSignal(socket, signal string) (*SignalResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindServices return [name, socket path] of service
func (c *Controller) FindServices(dir string) ([]*ServiceStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PurgeServices remove unused service directory
func (c *Controller) PurgeServices(dir string) error { _ = "STUB: not implemented"; return nil }

// Run executes a command and print combinedOutput
func (c *Controller) Run(command string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
