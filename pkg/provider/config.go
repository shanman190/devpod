package provider

import (
	"github.com/loft-sh/devpod/pkg/types"
)

const (
	CommandEnv = "COMMAND"
)

type ProviderConfig struct {
	// Name is the name of the provider
	Name string `json:"name,omitempty"`

	// Version is the provider version
	Version string `json:"version,omitempty"`

	// Type defines the type of the provider. Defaults to Server
	Type ProviderType `json:"type,omitempty"`

	// Description is the provider description
	Description string `json:"description,omitempty"`

	// Options are the provider options
	Options map[string]*ProviderOption `json:"options,omitempty"`

	// Agent allows you to override agent configuration
	Agent ProviderAgentConfig `json:"agent,omitempty"`

	// Exec holds the provider commands
	Exec ProviderCommands `json:"exec,omitempty"`

	// Binaries is an optional field to specify a binary to execute the commands
	Binaries []*ProviderBinary `json:"binaries,omitempty"`
}

type ProviderAgentConfig struct {
	// Path is the path inside the server to use for the agent
	Path string `json:"path,omitempty"`

	// DownloadURL is the base url where to download the agent from
	DownloadURL string `json:"downloadURL,omitempty"`

	// Options are extra options that should be available for the agent
	Options map[string]*ProviderOption `json:"options,omitempty"`

	// Inactivity specifies what should happen if the server is inactive for
	// a while.
	Inactivity *ProviderAgentInactivityConfig `json:"inactivity,omitempty"`
}

type ProviderAgentInactivityConfig struct {
	// Timeout is the timeout in minutes to wait until the agent tries
	// to turn of the server. Defaults to 1 hour.
	Timeout string `json:"inactivityTimeout,omitempty"`

	// Shutdown is the remote command to run when the remote machine
	// should shutdown.
	Shutdown types.StrArray `json:"shutdown,omitempty"`
}

type ProviderType string

const (
	ProviderTypeServer    = "Server"
	ProviderTypeWorkspace = "Workspace"
)

type ProviderBinary struct {
	// The current OS
	OS string `json:"os"`

	// The current Arch
	Arch string `json:"arch"`

	// The binary url to download from or relative path to use
	Path string `json:"path"`
}

type ProviderCommands struct {
	// Init is run directly after `devpod use provider`
	Init types.StrArray `json:"init,omitempty"`

	// Command executes a command on the server
	Command types.StrArray `json:"command,omitempty"`

	// Tunnel creates a tunnel to the workspace
	Tunnel types.StrArray `json:"tunnel,omitempty"`

	// Create creates a new server
	Create types.StrArray `json:"create,omitempty"`

	// Delete destroys a server
	Delete types.StrArray `json:"delete,omitempty"`

	// Start starts a stopped server
	Start types.StrArray `json:"start,omitempty"`

	// Stop stops a running server
	Stop types.StrArray `json:"stop,omitempty"`

	// Status retrieves the server status
	Status types.StrArray `json:"status,omitempty"`
}

type ProviderAgentOption struct {
	// Cache is the duration to cache the value before rerunning the command
	Cache string `json:"cache,omitempty"`

	// Command is the command to run to specify
	Command types.StrArray `json:"command,omitempty"`
}

type ProviderOption struct {
	// Default value if the user omits this option from their configuration.
	Default string `json:"default,omitempty"`

	// A description of the option displayed to the user by a supporting tool.
	Description string `json:"description,omitempty"`

	// ValidationPattern is a regex pattern to validate the value
	ValidationPattern string `json:"validationPattern,omitempty"`

	// ValidationMessage is the message that appears if the user enters an invalid option
	ValidationMessage string `json:"validationMessage,omitempty"`

	// Allowed values for this option.
	Enum []string `json:"enum,omitempty"`

	// Hidden specifies if the option should be hidden
	Hidden bool `json:"hidden,omitempty"`

	// Local will never send the option to the server
	Local bool `json:"local,omitempty"`

	// Cache is the duration to cache the value before rerunning the command
	Cache string `json:"cache,omitempty"`

	// Command is the command to run to specify
	Command types.StrArray `json:"command,omitempty"`
}