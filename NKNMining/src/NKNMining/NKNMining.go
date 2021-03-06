package main

import (
	"github.com/urfave/cli"
	"NKNMining/cli"
	"NKNMining/common"
	"NKNMining/config"
	"os"
	"time"
	"NKNMining/storage"
	"NKNMining/server"
	"NKNMining/container"
	"NKNMining/network/rpcRequest"
	"NKNMining/task"
	"path/filepath"
	"os/exec"
)

func startDaemon() (hasNewProcess bool)  {
	defer func() {
		if !hasNewProcess {
			storage.Temp.SaveParentPid(-1)
		}
	}()

	ppid := os.Getppid()
	needPPidCheck := true

	if nil != storage.Temp.Load() {
		needPPidCheck = false
	}

	if ppid != 1 {
		if needPPidCheck && ppid == storage.Temp.LastPPid {
			return false
		}
		storage.Temp.SaveParentPid(ppid)

		filePath,_ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)

		cmd.Stdin=os.Stdin
		cmd.Stdout=os.Stdout
		cmd.Stderr=os.Stderr
		cmd.Start()

		return true
	}

	return false
}

func getCliApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "NKN Mining"
	app.Version = "0.0.1"
	app.HelpName = "NKNMining"
	app.Usage = "NKN Mining application"
	app.UsageText = "NKNMining [options] [args]"
	app.HideHelp = false
	app.HideVersion = false

	return
}

func main() {
	//get app
	shell := getCliApp()

	//set some flags
	commands.SetFlags(shell)
	commands.SetAction(shell)

	//run
	err := shell.Run(os.Args)
	if nil != err {
		os.Exit(0)
	}

	//init logs
	common.InitLog(config.ShellConfig.LogFile)
	if startDaemon() {
		return
	}

	storage.InitSetupInfo()
	if common.ShellMutexCheck() {
		common.Log.Error("NKNMining is already running！")
		return
	}

	container.InitNodeContainers()

	rpcRequest.Api.Build()

	task.StartAllTask()

	//start server
	server.Start()

	for {
		time.Sleep(time.Second)
	}
}
