import{d as K,a as p,o as r,b as c,w as t,e as a,a0 as N,f as e,t as o,m as l,c as u,F as g,G as f,W as d,p as k,Y as L,Z as S,K as I,q as m,T as w,_ as $}from"./index-Cd9VyreM.js";import{S as A}from"./StatusBadge-Co3KsckJ.js";import{T}from"./TagList-CHb8KwQz.js";const E={class:"stack"},O={class:"stack-with-borders"},U={class:"status-with-reason"},F={key:0},G={class:"mt-4"},W={class:"stack-with-borders"},Z={class:"mt-4 stack"},q={class:"mt-2 stack-with-borders"},Y=K({__name:"DataPlaneSummaryView",props:{items:{}},setup(x){const b=x;return(j,H)=>{const z=p("RouteTitle"),C=p("RouterLink"),D=p("KTooltip"),y=p("DataCollection"),v=p("KBadge"),P=p("AppView"),V=p("RouteView");return r(),c(V,{name:"data-plane-summary-view",params:{dataPlane:""}},{default:t(({t:s,route:B})=>[a(y,{items:b.items,predicate:h=>h.id===B.params.dataPlane,find:!0},{empty:t(()=>[a(N,null,{title:t(()=>[e(o(s("common.collection.summary.empty_title",{type:"Data Plane Proxy"})),1)]),default:t(()=>[e(),l("p",null,o(s("common.collection.summary.empty_message",{type:"Data Plane Proxy"})),1)]),_:2},1024)]),default:t(({items:h})=>[(r(!0),u(g,null,f([h[0]],n=>(r(),c(P,{key:n.id},{title:t(()=>[l("h2",null,[a(C,{to:{name:"data-plane-detail-view",params:{dataPlane:n.id}}},{default:t(()=>[a(z,{title:s("data-planes.routes.item.title",{name:n.name})},null,8,["title"])]),_:2},1032,["to"])])]),default:t(()=>[e(),l("div",E,[l("div",O,[a(d,{layout:"horizontal"},{title:t(()=>[e(o(s("http.api.property.status")),1)]),body:t(()=>[l("div",U,[a(A,{status:n.status},null,8,["status"]),e(),n.dataplaneType==="standard"?(r(),c(y,{key:0,items:n.dataplane.networking.inbounds,predicate:_=>!_.health.ready,empty:!1},{default:t(({items:_})=>[a(D,{class:"reason-tooltip",placement:"bottomEnd"},{content:t(()=>[l("ul",null,[(r(!0),u(g,null,f(_,i=>(r(),u("li",{key:`${i.service}:${i.port}`},o(s("data-planes.routes.item.unhealthy_inbound",{service:i.service,port:i.port})),1))),128))])]),default:t(()=>[a(k(L),{color:k(S),size:k(I)},null,8,["color","size"]),e()]),_:2},1024)]),_:2},1032,["items","predicate"])):m("",!0)])]),_:2},1024),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(`
                    Type
                  `)]),body:t(()=>[e(o(s(`data-planes.type.${n.dataplaneType}`)),1)]),_:2},1024),e(),n.namespace.length>0?(r(),c(d,{key:0,layout:"horizontal"},{title:t(()=>[e(o(s("data-planes.routes.item.namespace")),1)]),body:t(()=>[e(o(n.namespace),1)]),_:2},1024)):m("",!0),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(o(s("data-planes.routes.item.last_updated")),1)]),body:t(()=>[e(o(s("common.formats.datetime",{value:Date.parse(n.modificationTime)})),1)]),_:2},1024)]),e(),n.dataplane.networking.gateway?(r(),u("div",F,[l("h3",null,o(s("data-planes.routes.item.gateway")),1),e(),l("div",G,[l("div",W,[a(d,{layout:"horizontal"},{title:t(()=>[e(o(s("http.api.property.tags")),1)]),body:t(()=>[a(T,{alignment:"right",tags:n.dataplane.networking.gateway.tags},null,8,["tags"])]),_:2},1024),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(o(s("http.api.property.address")),1)]),body:t(()=>[a(w,{text:`${n.dataplane.networking.address}`},null,8,["text"])]),_:2},1024)])])])):m("",!0),e(),n.dataplaneType==="standard"?(r(),c(y,{key:1,items:n.dataplane.networking.inbounds},{default:t(({items:_})=>[l("div",null,[l("h3",null,o(s("data-planes.routes.item.inbounds")),1),e(),l("div",Z,[(r(!0),u(g,null,f(_,(i,R)=>(r(),u("div",{key:R,class:"inbound"},[l("h4",null,[a(w,{text:i.tags["kuma.io/service"]},{default:t(()=>[e(o(s("data-planes.routes.item.inbound_name",{service:i.tags["kuma.io/service"]})),1)]),_:2},1032,["text"])]),e(),l("div",q,[a(d,{layout:"horizontal"},{title:t(()=>[e(o(s("http.api.property.status")),1)]),body:t(()=>[i.health.ready?(r(),c(v,{key:0,appearance:"success"},{default:t(()=>[e(o(s("data-planes.routes.item.health.ready")),1)]),_:2},1024)):(r(),c(v,{key:1,appearance:"danger"},{default:t(()=>[e(o(s("data-planes.routes.item.health.not_ready")),1)]),_:2},1024))]),_:2},1024),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(o(s("http.api.property.tags")),1)]),body:t(()=>[a(T,{alignment:"right",tags:i.tags},null,8,["tags"])]),_:2},1024),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(o(s("http.api.property.address")),1)]),body:t(()=>[a(w,{text:i.addressPort},null,8,["text"])]),_:2},1024)])]))),128))])])]),_:2},1032,["items"])):m("",!0)])]),_:2},1024))),128))]),_:2},1032,["items","predicate"])]),_:1})}}}),X=$(Y,[["__scopeId","data-v-dc42e0a8"]]);export{X as default};
