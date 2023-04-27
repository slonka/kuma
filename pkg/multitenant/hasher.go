package multitenant

import (
    context "context"
    envoy_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    "github.com/kumahq/kuma/pkg/core"
    v3 "github.com/kumahq/kuma/pkg/util/xds/v3"
)

var log = core.Log.WithName("custom-hasher")

type CtxIDHash struct {
    nodeHash v3.NodeHash
    ctx      context.Context
    resourceHashKey func(ctx context.Context) string
}

func (c CtxIDHash) ID(node *envoy_v3.Node) string {
    log.Info("hasher", "tenantId", TenantFromCtx(c.ctx))
    return c.nodeHash.ID(node) + ":" + c.resourceHashKey(c.ctx)
}

var _ v3.NodeHash = CtxIDHash{}

func NewCtxIDHash(ctx context.Context, resourceHashKey func(ctx context.Context) string, nodeHash v3.NodeHash) v3.NodeHash {
    return CtxIDHash{
        ctx:             ctx,
        resourceHashKey: resourceHashKey,
        nodeHash:        nodeHash,
    }
}
