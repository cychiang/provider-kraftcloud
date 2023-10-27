package compute

import ujconfig "github.com/upbound/upjet/pkg/config"

// Configure configures the compute group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("kraftcloud_instance", func(r *ujconfig.Resource) {
		r.Kind = "Instance"
		r.ShortGroup = "compute"
	})
}
