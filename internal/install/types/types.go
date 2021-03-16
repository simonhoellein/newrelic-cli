// Code generated by tutone: DO NOT EDIT
package types

// OpenInstallationOperatingSystem - Operating System of target environment
type OpenInstallationOperatingSystem string

var OpenInstallationOperatingSystemTypes = struct {
	// MacOS operating system
	DARWIN OpenInstallationOperatingSystem
	// Linux-based operating system
	LINUX OpenInstallationOperatingSystem
	// Windows operating system
	WINDOWS OpenInstallationOperatingSystem
}{
	// MacOS operating system
	DARWIN: "DARWIN",
	// Linux-based operating system
	LINUX: "LINUX",
	// Windows operating system
	WINDOWS: "WINDOWS",
}

// OpenInstallationPlatform - Operating System distribution
type OpenInstallationPlatform string

var OpenInstallationPlatformTypes = struct {
	// Amazon Linux operating system
	AMAZON OpenInstallationPlatform
	// CentOS operating system
	CENTOS OpenInstallationPlatform
	// Debian operating system
	DEBIAN OpenInstallationPlatform
	// RedHat Enterprise Linux operating system
	REDHAT OpenInstallationPlatform
	// SUSE operating system
	SUSE OpenInstallationPlatform
	// Ubuntu operating system
	UBUNTU OpenInstallationPlatform
}{
	// Amazon Linux operating system
	AMAZON: "AMAZON",
	// CentOS operating system
	CENTOS: "CENTOS",
	// Debian operating system
	DEBIAN: "DEBIAN",
	// RedHat Enterprise Linux operating system
	REDHAT: "REDHAT",
	// SUSE operating system
	SUSE: "SUSE",
	// Ubuntu operating system
	UBUNTU: "UBUNTU",
}

// OpenInstallationPlatformFamily - Operating System distribution family
type OpenInstallationPlatformFamily string

var OpenInstallationPlatformFamilyTypes = struct {
	// Debian distribution family
	DEBIAN OpenInstallationPlatformFamily
	// RHEL distribution family
	RHEL OpenInstallationPlatformFamily
	// openSUSE distribution family
	SUSE OpenInstallationPlatformFamily
}{
	// Debian distribution family
	DEBIAN: "DEBIAN",
	// RHEL distribution family
	RHEL: "RHEL",
	// openSUSE distribution family
	SUSE: "SUSE",
}

// OpenInstallationTargetType - Installation target type
type OpenInstallationTargetType string

var OpenInstallationTargetTypeTypes = struct {
	// APM agent installation
	APPLICATION OpenInstallationTargetType
	// Cloud provider installation
	CLOUD OpenInstallationTargetType
	// Docker container installation
	DOCKER OpenInstallationTargetType
	// Bare metal, virtual machine, or host-based installation
	HOST OpenInstallationTargetType
	// Kubernetes installation
	KUBERNETES OpenInstallationTargetType
	// Serverless installation
	SERVERLESS OpenInstallationTargetType
}{
	// APM agent installation
	APPLICATION: "APPLICATION",
	// Cloud provider installation
	CLOUD: "CLOUD",
	// Docker container installation
	DOCKER: "DOCKER",
	// Bare metal, virtual machine, or host-based installation
	HOST: "HOST",
	// Kubernetes installation
	KUBERNETES: "KUBERNETES",
	// Serverless installation
	SERVERLESS: "SERVERLESS",
}

// OpenInstallationAttributes - Custom event data attributes
type OpenInstallationAttributes struct {
	// Built-in parsing rulesets
	Logtype string `json:"logtype,omitempty"`
}

// OpenInstallationLogMatch - Matches partial list of the Log forwarding parameters
type OpenInstallationLogMatch struct {
	// List of custom attributes, as key-value pairs, that can be used to send additional data with the logs which you can then query.
	Attributes OpenInstallationAttributes `json:"attributes,omitempty"`
	// Path to the log file or files.
	File string `json:"file,omitempty"`
	// Name of the log or logs.
	Name string `json:"name"`
	// Regular expression for filtering records.
	Pattern string `json:"pattern,omitempty"`
	// Service name (Linux Only).
	Systemd string `json:"systemd,omitempty"`
}

// OpenInstallationPreInstallConfiguration - Optional pre-install configuration items
type OpenInstallationPreInstallConfiguration struct {
	// Message/Docs notice displayed to user prior to running recipe
	Prompt string `json:"prompt,omitempty"`
	// Message/Docs notice displayed to user prior to running recipe
	Info string `json:"info,omitempty"`
}

// OpenInstallationPostInstallConfiguration - Optional post-install configuration items
type OpenInstallationPostInstallConfiguration struct {
	// Message/Docs notice displayed to user after running the recipe
	Info string `json:"info,omitempty"`
}

// OpenInstallationRecipe - Installation instructions and definition of an instrumentation integration
type OpenInstallationRecipe struct {
	// Description of the recipe
	Description string `json:"description"`
	// Friendly name of the integration
	DisplayName string `json:"displayName,omitempty"`
	// The full contents of the recipe file (yaml)
	File string `json:"file"`
	// The ID
	ID string `json:"id,omitempty"`
	// List of variables to prompt for input from the user
	InputVars []OpenInstallationRecipeInputVariable `json:"inputVars"`
	// Go-task's taskfile definition (see https://taskfile.dev/#/usage)
	Install string `json:"install"`
	// Object representing the intended install target
	InstallTargets []OpenInstallationRecipeInstallTarget `json:"installTargets"`
	// Tags
	Keywords []string `json:"keywords"`
	// # Partial list of possible Log forwarding parameters
	LogMatch []OpenInstallationLogMatch `json:"logMatch"`
	// Short unique handle for the name of the integration
	Name string `json:"name,omitempty"`
	// Object representing optional pre-install configuration items
	PreInstall OpenInstallationPreInstallConfiguration `json:"preInstall,omitempty"`
	// Object representing optional post-install configuration items
	PostInstall OpenInstallationPostInstallConfiguration `json:"postInstall,omitempty"`
	// List of process definitions used to match CLI process detection
	ProcessMatch []string `json:"processMatch"`
	// Github repository url
	Repository string `json:"repository"`
	// NRQL the newrelic-cli uses to validate this recipe
	// is successfully sending data to New Relic
	ValidationNRQL NRQL `json:"validationNrql,omitempty"`
}

// OpenInstallationRecipeInputVariable - Recipe input variable prompts displayed to the user prior to execution
type OpenInstallationRecipeInputVariable struct {
	// Default value of variable
	Default string `json:"default,omitempty"`
	// Name of the variable
	Name string `json:"name"`
	// Message to present to the user
	Prompt string `json:"prompt,omitempty"`
	// Indicates a password field
	Secret bool `json:"secret,omitempty"`
}

// OpenInstallationRecipeInstallTarget - Matrix of supported installation criteria for this recipe
type OpenInstallationRecipeInstallTarget struct {
	// OS kernel architecture
	KernelArch string `json:"kernelArch,omitempty"`
	// OS kernel version
	KernelVersion string `json:"kernelVersion,omitempty"`
	// Operating system
	Os OpenInstallationOperatingSystem `json:"os,omitempty"`
	// Operating System distribution
	Platform OpenInstallationPlatform `json:"platform,omitempty"`
	// Operating System distribution family
	PlatformFamily OpenInstallationPlatformFamily `json:"platformFamily,omitempty"`
	// OS distribution version
	PlatformVersion string `json:"platformVersion,omitempty"`
	// Target type
	Type OpenInstallationTargetType `json:"type,omitempty"`
}

// NRQL - This scalar represents a NRQL query string.
//
// See the [NRQL Docs](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/nrql-resources/nrql-syntax-components-functions) for more information about NRQL syntax.
type NRQL string

type SuccessLink struct {
	Type   string `json:"type"`
	Filter string `json:"filter"`
}
