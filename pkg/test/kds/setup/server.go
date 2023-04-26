package setup

import (
	"github.com/kumahq/kuma/pkg/plugins/resources/postgres"
	"time"

	kuma_cp "github.com/kumahq/kuma/pkg/config/app/kuma-cp"
	"github.com/kumahq/kuma/pkg/core"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/store"
	"github.com/kumahq/kuma/pkg/core/runtime"
	"github.com/kumahq/kuma/pkg/core/runtime/component"
	"github.com/kumahq/kuma/pkg/kds/reconcile"
	kds_server "github.com/kumahq/kuma/pkg/kds/server"
	kds_server_v2 "github.com/kumahq/kuma/pkg/kds/v2/server"
	core_metrics "github.com/kumahq/kuma/pkg/metrics"
	"github.com/kumahq/kuma/pkg/multitenant"
)

type testRuntimeContext struct {
	runtime.Runtime
	rom                 manager.ReadOnlyResourceManager
	cfg                 kuma_cp.Config
	components          []component.Component
	metrics             core_metrics.Metrics
	hashingFn             multitenant.Hashing
	configCustomizationFn postgres.PgxConfigCustomizationFn
	tenant                multitenant.TenantFn
}

func (t *testRuntimeContext) Config() kuma_cp.Config {
	return t.cfg
}

func (t *testRuntimeContext) ReadOnlyResourceManager() manager.ReadOnlyResourceManager {
	return t.rom
}

func (t *testRuntimeContext) Metrics() core_metrics.Metrics {
	return t.metrics
}

func (t *testRuntimeContext) HashingFn() multitenant.Hashing {
	return t.hashingFn
}

func (t *testRuntimeContext) ConfigCustomizationFn() postgres.PgxConfigCustomizationFn {
	return t.configCustomizationFn
}

func (t *testRuntimeContext) TenantFn() multitenant.TenantFn {
	return t.tenant
}

func (t *testRuntimeContext) Add(c ...component.Component) error {
	t.components = append(t.components, c...)
	return nil
}

func StartServer(store store.ResourceStore, clusterID string, providedTypes []model.ResourceType, providedFilter reconcile.ResourceFilter, providedMapper reconcile.ResourceMapper) (kds_server.Server, error) {
	metrics, err := core_metrics.NewMetrics("Global")
	if err != nil {
		return nil, err
	}
	rt := &testRuntimeContext{
		rom:                   manager.NewResourceManager(store),
		cfg:                   kuma_cp.Config{},
		metrics:               metrics,
		tenant:                multitenant.SingleTenant,
		hashingFn:             multitenant.DefaultHashingFn,
		configCustomizationFn: postgres.DefaultPgxConfigCustomizationFn,
	}
	return kds_server.New(core.Log.WithName("kds").WithName(clusterID), rt, providedTypes, clusterID, 100*time.Millisecond, providedFilter, providedMapper, false, 1*time.Second)
}

func StartDeltaServer(store store.ResourceStore, clusterID string, providedTypes []model.ResourceType, providedFilter reconcile.ResourceFilter, providedMapper reconcile.ResourceMapper) (kds_server_v2.Server, error) {
	metrics, err := core_metrics.NewMetrics("Global")
	if err != nil {
		return nil, err
	}
	rt := &testRuntimeContext{
		rom:                   manager.NewResourceManager(store),
		cfg:                   kuma_cp.Config{},
		metrics:               metrics,
		tenant:                multitenant.SingleTenant,
		hashingFn:             multitenant.DefaultHashingFn,
		configCustomizationFn: postgres.DefaultPgxConfigCustomizationFn,
	}
	return kds_server_v2.New(core.Log.WithName("kds-delta").WithName(clusterID), rt, providedTypes, clusterID, 100*time.Millisecond, providedFilter, providedMapper, false, 1*time.Second)
}
