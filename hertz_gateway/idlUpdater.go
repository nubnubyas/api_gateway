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

// initialArraySize is the initial size of the array for storing services
const initialArraySize = 0

// httpMethodsInCaps is a list of HTTP methods in capital letters
// httpMethods is a list of HTTP methods
var (
	httpMethodsInCaps = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}
	httpMethods       = []string{"api.get", "api.post", "api.put", "api.delete", "api.patch", "api.head", "api.options"}
)

// createGenericClient creates a generic client for newly added IDL files
func createGenericClient(entryName string, idlPath string) error {
	thriftName := strings.ReplaceAll(entryName, ".thrift", "")
	filePath := idlPath + entryName
	fileSyntax, err := parser.ParseFile(filePath, nil, false)

	if err != nil {
		hlog.Errorf("parse file failed: %v", err)
		return err
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
		handler.FileToSvc[thriftName] = make([]string, initialArraySize)
	}

	registerServices(fileSyntax, thriftName, cli)

	return nil
}

// registerServices registers services for IDL files
func registerServices(fileSyntax *parser.Thrift, thriftName string, cli genericclient.Client) {
	for _, svc := range fileSyntax.Services {
		handler.FileToSvc[thriftName] = append(handler.FileToSvc[thriftName], svc.Name)
		handler.SvcMap[svc.Name] = cli

		for _, function := range svc.Functions {
			var subpath string
			functionName := function.Name

			if handler.PathToMethod[svc.Name] == nil {
				handler.PathToMethod[svc.Name] = make(map[handler.MethodPath]string)
			}

			for index, method := range httpMethods {
				if len(function.Annotations.Get(method)) > 0 {
					subpath = function.Annotations.Get(method)[0]
					methodParam := handler.MethodPath{
						Path:   subpath,
						Method: httpMethodsInCaps[index],
					}
					handler.PathToMethod[svc.Name][methodParam] = functionName
					break
				}
			}

			if subpath == "" {
				methodParam := handler.MethodPath{
					Path:   "/" + svc.Name + "/" + functionName,
					Method: "POST",
				}
				handler.PathToMethod[svc.Name][methodParam] = functionName
			}
		}
	}
}

// removeGenericClient removes a generic client for deleted IDL files
func removeGenericClient(entryName string, idlPath string) {
	thriftName := strings.ReplaceAll(entryName, ".thrift", "")
	services := handler.FileToSvc[thriftName]

	for _, svc := range services {
		delete(handler.SvcMap, svc)
		delete(handler.PathToMethod, svc)
	}
}

// watchIDLs watches the IDL directory for changes
func watchIDLs(idlPath string, errChan chan<- error) {
	changeChan := make(chan []string, 5)
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		errChan <- err
		return
	}

	defer watcher.Close()

	err = watcher.Add(idlPath)
	if err != nil {
		errChan <- err
		return
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
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

				if event.Op&fsnotify.Rename == fsnotify.Rename || event.Op&fsnotify.Write == fsnotify.Write {
					name := strings.Split(event.Name, "\\")
					select {
					case changeChan <- []string{name[2], "rename"}:
					default:
					}
				}

			case err := <-watcher.Errors:
				errChan <- err
				return
			}
		}
	}()

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
