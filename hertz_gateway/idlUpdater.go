package main

import (
	"fmt"
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

// createGenericClient creates a generic client for newly added IDL files
func createGenericClient(entryName string, idlPath string) error {
	// Remove .thrift extension from entry name
	thriftName := strings.ReplaceAll(entryName, ".thrift", "")

	// Construct file path
	filePath := idlPath + entryName

	// Parse the thrift file
	fileSyntax, err := parser.ParseFile(filePath, nil, false)
	if err != nil {
		hlog.Errorf("parse file failed: %v", err)
		return err
	}

	// Create a new thrift file provider
	provider, err := generic.NewThriftFileProvider(entryName, idlPath)
	if err != nil {
		hlog.Errorf("new thrift file provider failed: %v", err)
		return err
	}

	// Create a JSON thrift generic provider
	g, err := generic.JSONThriftGeneric(provider)
	if err != nil {
		hlog.Error(err)
		return err
	}

	// Create a generic client with a new weighted random balancer
	loadBalancerOpt := client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer())
	cli, err := genericclient.NewClient(
		thriftName,
		g,
		client.WithResolver(registerCenter.NacosResolver),
		loadBalancerOpt,
	)
	if err != nil {
		hlog.Error(err)
		return err
	}

	if handler.FileToSvc[thriftName] == nil {
		handler.FileToSvc[thriftName] = make([]string, 3)
	}

	httpMethodsInCaps := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}
	httpMethods := []string{"api.get", "api.post", "api.put", "api.delete", "api.patch", "api.head", "api.options"}

	fmt.Println(thriftName)
	for _, svc := range fileSyntax.Services {
		handler.FileToSvc[thriftName] = append(handler.FileToSvc[thriftName], svc.Name)
		handler.SvcMap[svc.Name] = cli

		for _, function := range svc.Functions {
			var subpath string
			functionName := function.Name

			// Initialize PathToMethod map if it's nil
			if handler.PathToMethod[svc.Name] == nil {
				handler.PathToMethod[svc.Name] = make(map[handler.MethodPath]string)
			}

			// Check if function has any HTTP method annotation
			for index, method := range httpMethods {
				if len(function.Annotations.Get(method)) > 0 {
					subpath = function.Annotations.Get(method)[0]
					methodParam := handler.MethodPath{
						Path:   subpath,
						Method: httpMethodsInCaps[index],
					}

					handler.PathToMethod[svc.Name][methodParam] = functionName
					//handler.PathToMethod[svcName][subpath] = functionName
					break
				}
			}

			// If function has no HTTP method annotation, default to POST
			if subpath == "" {
				methodParam := handler.MethodPath{
					Path:   "/" + svc.Name + "/" + functionName,
					Method: "POST",
				}
				handler.PathToMethod[svc.Name][methodParam] = functionName
			}
		}
	}
	fmt.Println(handler.PathToMethod)
	return nil
}

// removeGenericClient removes a generic client for deleted IDL files
func removeGenericClient(entryName string, idlPath string) {
	// Remove .thrift extension from entry name
	thriftName := strings.ReplaceAll(entryName, ".thrift", "")

	// Get services associated with the thrift file
	services := handler.FileToSvc[thriftName]

	// Remove services from SvcMap and PathToMethod
	for _, svc := range services {
		delete(handler.SvcMap, svc)
		delete(handler.PathToMethod, svc)
	}
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
				// Handle different types of file changes
				if event.Op&fsnotify.Create == fsnotify.Create {
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

				// Versions to be used must exist in the idl folder
				if event.Op&fsnotify.Rename == fsnotify.Rename || event.Op&fsnotify.Write == fsnotify.Write {
					name := strings.Split(event.Name, "\\")
					select {
					case changeChan <- []string{name[2], "rename"}:
					default:
					}
				}
			case err := <-watcher.Errors:
				hlog.Error(err)
			}
		}
	}()

	// Handle file changes
	for entry := range changeChan {
		switch entry[1] {
		case "add":
			createGenericClient(entry[0], idlPath)
		case "remove":
			removeGenericClient(entry[0], idlPath)
		case "rename":
			removeGenericClient(entry[0], idlPath)
		}
	}
}
