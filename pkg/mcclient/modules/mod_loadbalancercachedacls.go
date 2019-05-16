package modules

type LoadbalancerCachedAclManager struct {
	ResourceManager
}

var (
	LoadbalancerCachedAcls LoadbalancerCachedAclManager
)

func init() {
	LoadbalancerCachedAcls = LoadbalancerCachedAclManager{
		NewComputeManager(
			"cached_loadbalancer_acl",
			"cached_loadbalancer_acls",
			[]string{
				"id",
				"acl_id",
				"name",
				"acl_entries",
			},
			[]string{"tenant"},
		),
	}
	registerCompute(&LoadbalancerCachedAcls)
}
