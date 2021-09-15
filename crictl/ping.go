package golangcrictl

import (
	"encoding/json"
	"net/http"
	"time"
)

type Ping struct {
	ID                string     `json:"ID"`
	Containers        int        `json:"Containers"`
	ContainersRunning int        `json:"ContainersRunning"`
	ContainersPaused  int        `json:"ContainersPaused"`
	ContainersStopped int        `json:"ContainersStopped"`
	Images            int        `json:"Images"`
	Driver            string     `json:"Driver"`
	DriverStatus      [][]string `json:"DriverStatus"`
	Plugins           struct {
		Volume        []string    `json:"Volume"`
		Network       []string    `json:"Network"`
		Authorization interface{} `json:"Authorization"`
		Log           []string    `json:"Log"`
	} `json:"Plugins"`
	MemoryLimit        bool      `json:"MemoryLimit"`
	SwapLimit          bool      `json:"SwapLimit"`
	KernelMemory       bool      `json:"KernelMemory"`
	KernelMemoryTCP    bool      `json:"KernelMemoryTCP"`
	CPUCfsPeriod       bool      `json:"CpuCfsPeriod"`
	CPUCfsQuota        bool      `json:"CpuCfsQuota"`
	CPUShares          bool      `json:"CPUShares"`
	CPUSet             bool      `json:"CPUSet"`
	PidsLimit          bool      `json:"PidsLimit"`
	IPv4Forwarding     bool      `json:"IPv4Forwarding"`
	BridgeNfIptables   bool      `json:"BridgeNfIptables"`
	BridgeNfIP6Tables  bool      `json:"BridgeNfIp6tables"`
	Debug              bool      `json:"Debug"`
	NFd                int       `json:"NFd"`
	OomKillDisable     bool      `json:"OomKillDisable"`
	NGoroutines        int       `json:"NGoroutines"`
	SystemTime         time.Time `json:"SystemTime"`
	LoggingDriver      string    `json:"LoggingDriver"`
	CgroupDriver       string    `json:"CgroupDriver"`
	CgroupVersion      string    `json:"CgroupVersion"`
	NEventsListener    int       `json:"NEventsListener"`
	KernelVersion      string    `json:"KernelVersion"`
	OperatingSystem    string    `json:"OperatingSystem"`
	OSVersion          string    `json:"OSVersion"`
	OSType             string    `json:"OSType"`
	Architecture       string    `json:"Architecture"`
	IndexServerAddress string    `json:"IndexServerAddress"`
	RegistryConfig     struct {
		AllowNondistributableArtifactsCIDRs     []interface{} `json:"AllowNondistributableArtifactsCIDRs"`
		AllowNondistributableArtifactsHostnames []interface{} `json:"AllowNondistributableArtifactsHostnames"`
		InsecureRegistryCIDRs                   []string      `json:"InsecureRegistryCIDRs"`
		IndexConfigs                            struct {
			DockerIo struct {
				Name     string        `json:"Name"`
				Mirrors  []interface{} `json:"Mirrors"`
				Secure   bool          `json:"Secure"`
				Official bool          `json:"Official"`
			} `json:"docker.io"`
		} `json:"IndexConfigs"`
		Mirrors []interface{} `json:"Mirrors"`
	} `json:"RegistryConfig"`
	Ncpu              int           `json:"NCPU"`
	MemTotal          int           `json:"MemTotal"`
	GenericResources  interface{}   `json:"GenericResources"`
	DockerRootDir     string        `json:"DockerRootDir"`
	HTTPProxy         string        `json:"HttpProxy"`
	HTTPSProxy        string        `json:"HttpsProxy"`
	NoProxy           string        `json:"NoProxy"`
	Name              string        `json:"Name"`
	Labels            []interface{} `json:"Labels"`
	ExperimentalBuild bool          `json:"ExperimentalBuild"`
	ServerVersion     string        `json:"ServerVersion"`
	Runtimes          struct {
		IoContainerdRuncV2 struct {
			Path string `json:"path"`
		} `json:"io.containerd.runc.v2"`
		IoContainerdRuntimeV1Linux struct {
			Path string `json:"path"`
		} `json:"io.containerd.runtime.v1.linux"`
		Runc struct {
			Path string `json:"path"`
		} `json:"runc"`
	} `json:"Runtimes"`
	DefaultRuntime string `json:"DefaultRuntime"`
	Swarm          struct {
		NodeID           string      `json:"NodeID"`
		NodeAddr         string      `json:"NodeAddr"`
		LocalNodeState   string      `json:"LocalNodeState"`
		ControlAvailable bool        `json:"ControlAvailable"`
		Error            string      `json:"Error"`
		RemoteManagers   interface{} `json:"RemoteManagers"`
	} `json:"Swarm"`
	LiveRestoreEnabled bool   `json:"LiveRestoreEnabled"`
	Isolation          string `json:"Isolation"`
	InitBinary         string `json:"InitBinary"`
	ContainerdCommit   struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"ContainerdCommit"`
	RuncCommit struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"RuncCommit"`
	InitCommit struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"InitCommit"`
	SecurityOptions []string    `json:"SecurityOptions"`
	Warnings        interface{} `json:"Warnings"`
}

func (p *Ping) New(client http.Client) error {
	r, err := client.Get("http://cri-o/info")
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&p)
}
