package store_test

import (
	"context"
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	"github.com/kumahq/kuma/pkg/core"
	"github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/apis/system"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/store"
	"github.com/kumahq/kuma/pkg/kds/util"
	client_v2 "github.com/kumahq/kuma/pkg/kds/v2/client"
	sync_store "github.com/kumahq/kuma/pkg/kds/v2/store"
	core_metrics "github.com/kumahq/kuma/pkg/metrics"
	"github.com/kumahq/kuma/pkg/plugins/resources/memory"
	. "github.com/kumahq/kuma/pkg/test/matchers"
	model2 "github.com/kumahq/kuma/pkg/test/resources/model"
	test_store "github.com/kumahq/kuma/pkg/test/store"
)

var meshBuilder = func(idx int) *mesh.MeshResource {
	ca := fmt.Sprintf("ca-%d", idx)
	meshName := fmt.Sprintf("mesh-%d", idx)
	return &mesh.MeshResource{
		Meta: &model2.ResourceMeta{
			Name: meshName,
		},
		Spec: &mesh_proto.Mesh{
			Mtls: &mesh_proto.Mesh_Mtls{
				EnabledBackend: ca,
				Backends: []*mesh_proto.CertificateAuthorityBackend{
					{
						Name: ca,
						Type: "builtin",
					},
				},
			},
		},
	}
}

var _ = Describe("SyncResourceStoreDelta", func() {
	var syncer sync_store.ResourceSyncer
	var resourceStore store.ResourceStore

	BeforeEach(func() {
		resourceStore = memory.NewStore()
		metrics, err := core_metrics.NewMetrics("")
		Expect(err).ToNot(HaveOccurred())
		syncer, err = sync_store.NewResourceSyncer(core.Log, resourceStore, store.NoTransactions{}, metrics, context.Background())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should create new resources in empty store", func() {
		upstreamResponse := client_v2.UpstreamResponse{}
		upstream := &mesh.MeshResourceList{}
		idxs := []int{1, 2, 3, 4}
		for _, i := range idxs {
			m := meshBuilder(i)
			err := upstream.AddItem(m)
			Expect(err).ToNot(HaveOccurred())
		}
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream

		err, nackError := syncer.Sync(context.Background(), upstreamResponse)
		Expect(err).ToNot(HaveOccurred())
		Expect(nackError).ToNot(HaveOccurred())

		actual := &mesh.MeshResourceList{}
		err = resourceStore.List(context.Background(), actual)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual.Items).To(Equal(upstream.Items))
	})

	It("should delete all resources", func() {
		upstreamResponse := client_v2.UpstreamResponse{}
		removedResources := []model.ResourceKey{}
		for i := 0; i < 10; i++ {
			m := meshBuilder(i)
			removedResources = append(removedResources, model.WithoutMesh(fmt.Sprintf("mesh-%d", i)))
			err := resourceStore.Create(context.Background(), m, store.CreateBy(model.MetaToResourceKey(m.GetMeta())))
			Expect(err).ToNot(HaveOccurred())
		}
		upstream := &mesh.MeshResourceList{}
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream
		upstreamResponse.RemovedResourcesKey = removedResources

		err, nackError := syncer.Sync(context.Background(), upstreamResponse)
		Expect(err).ToNot(HaveOccurred())
		Expect(nackError).ToNot(HaveOccurred())

		actual := &mesh.MeshResourceList{}
		err = resourceStore.List(context.Background(), actual)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual.Items).To(BeEmpty())
	})

	It("should delete resources which are not represented in upstream and create new", func() {
		for i := 0; i < 10; i++ {
			m := meshBuilder(i)
			err := resourceStore.Create(context.Background(), m, store.CreateBy(model.MetaToResourceKey(m.GetMeta())))
			Expect(err).ToNot(HaveOccurred())
		}

		upstream := &mesh.MeshResourceList{}
		idxs := []int{1, 2, 7, 12}
		for _, i := range idxs {
			m := meshBuilder(i)
			err := upstream.AddItem(m)
			Expect(err).ToNot(HaveOccurred())
		}
		upstreamResponse := client_v2.UpstreamResponse{}
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream
		upstreamResponse.RemovedResourcesKey = []model.ResourceKey{
			model.WithoutMesh("mesh-0"),
			model.WithoutMesh("mesh-3"),
			model.WithoutMesh("mesh-4"),
			model.WithoutMesh("mesh-5"),
			model.WithoutMesh("mesh-6"),
			model.WithoutMesh("mesh-8"),
			model.WithoutMesh("mesh-9"),
			model.WithoutMesh("mesh-10"),
		}

		err, nackError := syncer.Sync(context.Background(), upstreamResponse)
		Expect(err).ToNot(HaveOccurred())
		Expect(nackError).ToNot(HaveOccurred())

		actual := &mesh.MeshResourceList{}
		err = resourceStore.List(context.Background(), actual)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual.Items).To(HaveLen(len(upstream.Items)))
		for i, item := range actual.Items {
			Expect(item.Spec).To(MatchProto(upstream.Items[i].Spec))
		}
	})

	It("should delete resources which are not represented in upstream and create new when is an initial request", func() {
		for i := 0; i < 10; i++ {
			m := meshBuilder(i)
			err := resourceStore.Create(context.Background(), m, store.CreateBy(model.MetaToResourceKey(m.GetMeta())))
			Expect(err).ToNot(HaveOccurred())
		}

		upstream := &mesh.MeshResourceList{}
		idxs := []int{1, 2, 7, 12}
		for _, i := range idxs {
			m := meshBuilder(i)
			err := upstream.AddItem(m)
			Expect(err).ToNot(HaveOccurred())
		}
		upstreamResponse := client_v2.UpstreamResponse{}
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream
		upstreamResponse.IsInitialRequest = true

		err, nackError := syncer.Sync(context.Background(), upstreamResponse)
		Expect(err).ToNot(HaveOccurred())
		Expect(nackError).ToNot(HaveOccurred())

		actual := &mesh.MeshResourceList{}
		err = resourceStore.List(context.Background(), actual)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual.Items).To(HaveLen(len(upstream.Items)))
		for i, item := range actual.Items {
			Expect(item.Spec).To(MatchProto(upstream.Items[i].Spec))
		}
	})

	It("should ignore resources from upstream that it does not support", func() {
		// given
		upstream := &mesh.MeshResourceList{}
		Expect(upstream.AddItem(meshBuilder(1))).To(Succeed())
		upstreamResponse := client_v2.UpstreamResponse{}
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream

		// when
		err, nackError := syncer.Sync(context.Background(), upstreamResponse, sync_store.PrefilterBy(func(r model.Resource) bool {
			return r.GetMeta().GetName() != "mesh-1"
		}))

		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(nackError).ToNot(HaveOccurred())
		actual := &mesh.MeshResourceList{}
		Expect(resourceStore.List(context.Background(), actual)).To(Succeed())
		Expect(actual.GetItems()).To(BeEmpty())
	})

	It("should add all resources and skip the conflict one", func() {
		mesh1 := meshBuilder(1)
		mesh2 := meshBuilder(2)
		mesh3 := meshBuilder(3)
		Expect(resourceStore.Create(
			context.Background(),
			mesh2,
			store.CreateBy(model.MetaToResourceKey(mesh2.GetMeta())),
			store.CreateWithLabels(map[string]string{mesh_proto.ResourceOriginLabel: "zone"}),
		)).ToNot(HaveOccurred())

		// given
		upstream := &mesh.MeshResourceList{}

		// try to add resource without the label
		mesh2 = meshBuilder(2)
		mesh2.Spec.MeshServices = &mesh_proto.Mesh_MeshServices{
			Mode: mesh_proto.Mesh_MeshServices_Exclusive,
		}
		Expect(upstream.AddItem(mesh1)).To(Succeed())
		Expect(upstream.AddItem(mesh2)).To(Succeed())
		Expect(upstream.AddItem(mesh3)).To(Succeed())
		upstreamResponse := client_v2.UpstreamResponse{}
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream

		// when
		err, nackError := syncer.Sync(context.Background(), upstreamResponse, sync_store.PrefilterBy(func(r model.Resource) bool {
			return r.GetMeta().GetLabels()[mesh_proto.ResourceOriginLabel] != "zone"
		}), sync_store.SkipConflictResource())

		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(nackError).To(HaveOccurred())

		actual := &mesh.MeshResourceList{}
		Expect(resourceStore.List(context.Background(), actual)).To(Succeed())
		Expect(actual.GetItems()).To(HaveLen(3))
		Expect(actual.GetItems()[0].GetSpec()).To(MatchProto(meshBuilder(2).GetSpec()))
		Expect(actual.GetItems()[1].GetSpec()).To(MatchProto(mesh1.GetSpec()))
		// should not update resource since mesh-2 already exists
		Expect(actual.GetItems()[2].GetSpec()).To(MatchProto(mesh3.GetSpec()))
	})

	It("should ignore invalid resource from upstream and add only valid", func() {
		// given
		upstream := &mesh.MeshResourceList{}
		mesh1 := meshBuilder(1)
		mesh2 := meshBuilder(2)
		mesh3 := meshBuilder(3)
		Expect(upstream.AddItem(mesh1)).To(Succeed())
		Expect(upstream.AddItem(mesh2)).To(Succeed())
		Expect(upstream.AddItem(mesh3)).To(Succeed())
		upstreamResponse := client_v2.UpstreamResponse{}
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream
		upstreamResponse.InvalidResourcesKey = []model.ResourceKey{model.MetaToResourceKey(mesh2.GetMeta())}

		// when
		err, nackError := syncer.Sync(context.Background(), upstreamResponse)

		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(nackError).ToNot(HaveOccurred())
		actual := &mesh.MeshResourceList{}
		Expect(resourceStore.List(context.Background(), actual)).To(Succeed())
		Expect(actual.GetItems()).To(HaveLen(2))
		Expect(actual.GetItems()[0].GetSpec()).To(MatchProto(mesh1.GetSpec()))
		Expect(actual.GetItems()[1].GetSpec()).To(MatchProto(mesh3.GetSpec()))
	})

	It("should not update resource with the equal spec", func() {
		// given resource in the store
		res := meshBuilder(1)
		key := model.MetaToResourceKey(res.GetMeta())
		Expect(resourceStore.Create(context.Background(), res, store.CreateBy(key))).To(Succeed())
		existing := mesh.NewMeshResource()
		Expect(resourceStore.Get(context.Background(), existing, store.GetBy(key))).To(Succeed())

		// when sync the resource with equal 'spec'
		upstream := &mesh.MeshResourceList{}
		Expect(upstream.AddItem(meshBuilder(1))).To(Succeed())

		upstreamResponse := client_v2.UpstreamResponse{}
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream

		Expect(syncer.Sync(context.Background(), upstreamResponse)).To(Succeed())

		// then resource's version is the same
		actual := mesh.NewMeshResource()
		Expect(resourceStore.Get(context.Background(), actual, store.GetBy(key))).To(Succeed())
		Expect(actual.GetMeta().GetVersion()).To(Equal(existing.GetMeta().GetVersion()))
	})
})

var _ = Describe("SyncResourceStoreDelta errors", func() {
	var syncer sync_store.ResourceSyncer
	var resourceStore store.ResourceStore

	BeforeEach(func() {
		resourceStore = &test_store.FailingStore{CreateErr: errors.Join(store.ErrorResourceAlreadyExists(system.GlobalSecretType, "zone-token-signing-public-key-1", ""))}
		metrics, err := core_metrics.NewMetrics("")
		Expect(err).ToNot(HaveOccurred())
		syncer, err = sync_store.NewResourceSyncer(core.Log, resourceStore, store.NoTransactions{}, metrics, context.Background())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should correctly recognize user errors", func() {
		upstreamResponse := client_v2.UpstreamResponse{}
		upstream := &mesh.MeshResourceList{}
		m := meshBuilder(1)
		err := upstream.AddItem(m)
		Expect(err).ToNot(HaveOccurred())
		upstreamResponse.Type = upstream.GetItemType()
		upstreamResponse.AddedResources = upstream

		err, nackError := syncer.Sync(context.Background(), upstreamResponse, sync_store.SkipConflictResource())
		Expect(err).ToNot(HaveOccurred())
		Expect(util.IsUserError(nackError)).To(BeTrue())
		Expect(util.IsUserErrorMessage(nackError.Error())).To(BeTrue())
		Expect(nackError).To(MatchError(`user error
resource already exists: type="GlobalSecret" name="zone-token-signing-public-key-1" mesh=""`))
	})
})
