package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jkstack/jkframe/utils"
	agent "github.com/jkstack/libagent"
	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var ConfDir string
var Version string

func Install(*cobra.Command, []string) {
	if len(ConfDir) == 0 {
		fmt.Println("missing --conf argument")
		os.Exit(1)
	}

	dir, err := filepath.Abs(ConfDir)
	utils.Assert(err)

	dummy := agent.NewDummyApp(AgentName, dir)

	err = agent.RegisterService(dummy)
	if err != nil {
		fmt.Printf("can not register service: %v\n", err)
		return
	}
	fmt.Println("register service success")
}

func Uninstall(*cobra.Command, []string) {
	err := agent.UnregisterService(agent.NewDummyApp(AgentName, ""))
	if err != nil {
		fmt.Printf("can not unregister service: %v\n", err)
		return
	}
	fmt.Println("unregister service success")
}

func Start(*cobra.Command, []string) {
	err := agent.Start(agent.NewDummyApp(AgentName, ""))
	if err != nil {
		fmt.Printf("can not start service: %v\n", err)
		return
	}
	fmt.Println("start service success")
}

func Stop(*cobra.Command, []string) {
	err := agent.Stop(agent.NewDummyApp(AgentName, ""))
	if err != nil {
		fmt.Printf("can not stop service: %v\n", err)
		return
	}
	fmt.Println("stop service success")
}

func Restart(*cobra.Command, []string) {
	err := agent.Restart(agent.NewDummyApp(AgentName, ""))
	if err != nil {
		fmt.Printf("can not restart service: %v\n", err)
		return
	}
	fmt.Println("restart service success")
}

func Status(*cobra.Command, []string) {
	status, err := agent.Status(agent.NewDummyApp(AgentName, ""))
	if err != nil {
		fmt.Printf("can not get service status: %v\n", err)
		return
	}
	switch status {
	case service.StatusRunning:
		fmt.Println("service is running")
	case service.StatusStopped:
		fmt.Println("service is stopped")
	case service.StatusUnknown:
		fmt.Println("service status unknown")
	}
}

func Run(*cobra.Command, []string) {
	if len(ConfDir) == 0 {
		fmt.Println("missing --conf argument")
		os.Exit(1)
	}

	dir, err := filepath.Abs(ConfDir)
	utils.Assert(err)

	agent.Run(New(dir, Version))
}
