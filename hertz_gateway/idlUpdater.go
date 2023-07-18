package main

import (
	"strings"

	"github.com/cloudwego/api_gateway/hertz_gateway/biz/handler"
	registerCenter "github.com/cloudwego/api_gateway/register_center/shared"
	"github.com/cloudwego/thriftgo/parser"
	"github.com/fsnotify/fsnotify"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
)

// to split the path name
func methodSplit(pathName []string) string {
	path := pathName[0]
	parts := strings.Split(path, "/")
	subpath := strings.Join(parts[1:], "/")
	return subpath
}

// create generic client for newly added IDL files
func clientCreation(entryName string, idlPath string) error {
	svcName := strings.ReplaceAll(entryName, ".thrift", "")

	filePath := idlPath + entryName

	fileSyntax, err := parser.ParseFile(filePath, nil, false)
	if err != nil {
		hlog.Errorf("parse file failed: %v", err)
		return err
	}

	for _, svc := range fileSyntax.Services {
		for _, function := range svc.Functions {
			functionName := function.Name
			if handler.PathToMethod[svcName] == nil {
				handler.PathToMethod[svcName] = make(map[string]string)
			}

			var subpath string
			httpMethods := []string{"api.get", "api.post", "api.put", "api.delete", "api.patch", "api.head", "api.options"}

			for _, method := range httpMethods {
				if len(function.Annotations.Get(method)) > 0 {
					subpath = methodSplit(function.Annotations.Get(method))
					handler.PathToMethod[svcName][subpath] = functionName
					break
				}
			}
		}
	}

	provider, err := generic.NewThriftFileProvider(entryName, idlPath)
	if err != nil {
		hlog.Errorf("new thrift file provider failed: %v", err)
		return err
	}

	g, err := generic.JSONThriftGeneric(provider)
	if err != nil {
		hlog.Error(err)
		return err
	}

	loadbalanceropt := client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer())
	// creates new generic client for each IDL
	cli, err := genericclient.NewClient(
		svcName,
		g,
		client.WithResolver(registerCenter.NacosResolver),
		loadbalanceropt,
	)
	if err != nil {
		hlog.Error(err)
		return err
	}

	handler.SvcMap[svcName] = cli
	return nil
}

// remove generic client for deleted IDL files
func clientRemoval(entryName string, idlPath string) {
	svcName := strings.ReplaceAll(entryName, ".thrift", "")
	delete(handler.SvcMap, svcName)
	delete(handler.PathToMethod, svcName)
}

// watchIDLs watches the IDL directory for changes
func watchIDLs(idlPath string) {
	// Create a channel to signal when a change has been detected
	changeChan := make(chan []string, 5)

	// Start a goroutine to watch for changes in the IDL directory
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			hlog.Error(err)
		}
		defer watcher.Close()

		// Watch IDL path for changes
		err = watcher.Add(idlPath)
		if err != nil {
			hlog.Error(err)
		}

		for {
			select {
			case event := <-watcher.Events:
				/*
					if event.Op&fsnotify.Write == fsnotify.Write {}
					if event.Op&fsnotify.Rename == fsnotify.Rename {}
				*/
				if event.Op&fsnotify.Create == fsnotify.Create && event.Op&fsnotify.Rename != fsnotify.Rename {
					// IDL file has been modified, signal change
					name := strings.Split(event.Name, "\\")
					select {
					case changeChan <- []string{name[2], "add"}:
					default:
					}
				}

				if event.Op&fsnotify.Remove == fsnotify.Remove {
					name := strings.Split(event.Name, "\\")
					select {
					case changeChan <- []string{name[2], "remove"}:
					default:
					}
				}
			case err := <-watcher.Errors:
				hlog.Error(err)
			}
		}
	}()

	for entry := range changeChan {
		switch entry[1] {
		case "add":
			clientCreation(entry[0], idlPath)
		case "remove":
			clientRemoval(entry[0], idlPath)
		}
	}
}