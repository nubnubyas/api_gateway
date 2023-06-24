package registerCenter

import (
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

var (
	NacosRegistry, ErrRegistry = registry.NewDefaultNacosRegistry()
	NacosResolver, ErrResolver = resolver.NewDefaultNacosResolver()
)
