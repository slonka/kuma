import{d as E,r as p,o as s,m,w as e,b as i,k as g,Y as h,e as n,c as r,F as u,s as k,t as l,p as _,l as I,aC as N,A as K,S as X,E as q,q as F}from"./index-DHg9Fngg.js";import{F as $}from"./FilterBar-9vnNu05r.js";import{S as G}from"./SummaryView-DmzBN36G.js";const M={class:"stack"},O={class:"columns"},j={key:0},J={key:1},W=E({__name:"MeshServiceDetailView",props:{data:{}},setup(x){const v=x;return(f,Y)=>{const y=p("KTruncate"),z=p("KBadge"),w=p("KCard"),C=p("RouterLink"),S=p("XIcon"),P=p("XAction"),A=p("XActionGroup"),D=p("RouterView"),T=p("DataLoader"),L=p("AppView"),R=p("RouteView");return s(),m(R,{name:"mesh-service-detail-view",params:{mesh:"",service:"",page:1,size:50,s:"",dataPlane:"",codeSearch:"",codeFilter:!1,codeRegExp:!1}},{default:e(({can:V,route:o,t:c,uri:B,me:d})=>[i(L,null,{default:e(()=>[g("div",M,[i(w,null,{default:e(()=>[g("div",O,[v.data.status.addresses.length>0?(s(),m(h,{key:0},{title:e(()=>[n(`
                Addresses
              `)]),body:e(()=>[i(y,null,{default:e(()=>[(s(!0),r(u,null,k(v.data.status.addresses,t=>(s(),r("span",{key:t.hostname},l(t.hostname),1))),128))]),_:1})]),_:1})):_("",!0),n(),i(h,null,{title:e(()=>[n(`
                Ports
              `)]),body:e(()=>[i(y,null,{default:e(()=>[(s(!0),r(u,null,k(f.data.spec.ports,t=>(s(),m(z,{key:t.port,appearance:"info"},{default:e(()=>[n(l(t.port)+":"+l(t.targetPort)+"/"+l(t.appProtocol),1)]),_:2},1024))),128))]),_:1})]),_:1}),n(),i(h,null,{title:e(()=>[n(`
                Dataplane Tags
              `)]),body:e(()=>[i(y,null,{default:e(()=>[(s(!0),r(u,null,k(f.data.spec.selector.dataplaneTags,(t,a)=>(s(),m(z,{key:`${a}:${t}`,appearance:"info"},{default:e(()=>[n(l(a)+":"+l(t),1)]),_:2},1024))),128))]),_:1})]),_:1}),n(),f.data.status.vips.length>0?(s(),m(h,{key:1,class:"ip"},{title:e(()=>[n(`
                VIPs
              `)]),body:e(()=>[i(y,null,{default:e(()=>[(s(!0),r(u,null,k(f.data.status.vips,t=>(s(),r("span",{key:t.ip},l(t.ip),1))),128))]),_:1})]),_:1})):_("",!0)])]),_:1}),n(),g("div",null,[g("h3",null,l(c("services.detail.data_plane_proxies")),1),n(),i(w,{class:"mt-4"},{default:e(()=>[i(T,{src:B(I(N),"/meshes/:mesh/dataplanes/for/mesh-service/:tags",{mesh:o.params.mesh,tags:JSON.stringify(v.data.spec.selector.dataplaneTags)},{page:o.params.page,size:o.params.size,search:o.params.s})},{loadable:e(({data:t})=>[i(K,{class:"data-plane-collection","data-testid":"data-plane-collection","page-number":o.params.page,"page-size":o.params.size,headers:[{...d.get("headers.name"),label:"Name",key:"name"},{...d.get("headers.namespace"),label:"Namespace",key:"namespace"},...V("use zones")?[{...d.get("headers.zone"),label:"Zone",key:"zone"}]:[],{...d.get("headers.certificate"),label:"Certificate Info",key:"certificate"},{...d.get("headers.status"),label:"Status",key:"status"},{...d.get("headers.warnings"),label:"Warnings",key:"warnings",hideLabel:!0},{...d.get("headers.actions"),label:"Actions",key:"actions",hideLabel:!0}],items:t==null?void 0:t.items,total:t==null?void 0:t.total,"is-selected-row":a=>a.name===o.params.dataPlane,"summary-route-name":"service-data-plane-summary-view","empty-state-message":c("common.emptyState.message",{type:"Data Plane Proxies"}),"empty-state-cta-to":c("data-planes.href.docs.data_plane_proxy"),"empty-state-cta-text":c("common.documentation"),onChange:o.update,onResize:d.set},{toolbar:e(()=>[i($,{class:"data-plane-proxy-filter",placeholder:"name:dataplane-name",query:o.params.s,fields:{name:{description:"filter by name or parts of a name"},protocol:{description:"filter by “kuma.io/protocol” value"},tag:{description:"filter by tags (e.g. “tag: version:2”)"},...V("use zones")&&{zone:{description:"filter by “kuma.io/zone” value"}}},onChange:a=>o.update({...Object.fromEntries(a.entries())})},null,8,["query","fields","onChange"])]),name:e(({row:a})=>[i(C,{class:"name-link",to:{name:"mesh-service-data-plane-summary-view",params:{mesh:a.mesh,dataPlane:a.id},query:{page:o.params.page,size:o.params.size,s:o.params.s}}},{default:e(()=>[n(l(a.name),1)]),_:2},1032,["to"])]),namespace:e(({row:a})=>[n(l(a.namespace),1)]),zone:e(({row:a})=>[a.zone?(s(),m(C,{key:0,to:{name:"zone-cp-detail-view",params:{zone:a.zone}}},{default:e(()=>[n(l(a.zone),1)]),_:2},1032,["to"])):(s(),r(u,{key:1},[n(l(c("common.collection.none")),1)],64))]),certificate:e(({row:a})=>{var b;return[(b=a.dataplaneInsight.mTLS)!=null&&b.certificateExpirationTime?(s(),r(u,{key:0},[n(l(c("common.formats.datetime",{value:Date.parse(a.dataplaneInsight.mTLS.certificateExpirationTime)})),1)],64)):(s(),r(u,{key:1},[n(l(c("data-planes.components.data-plane-list.certificate.none")),1)],64))]}),status:e(({row:a})=>[i(X,{status:a.status},null,8,["status"])]),warnings:e(({row:a})=>[a.isCertExpired||a.warnings.length>0?(s(),m(S,{key:0,class:"mr-1",name:"warning"},{default:e(()=>[g("ul",null,[a.warnings.length>0?(s(),r("li",j,l(c("data-planes.components.data-plane-list.version_mismatch")),1)):_("",!0),n(),a.isCertExpired?(s(),r("li",J,l(c("data-planes.components.data-plane-list.cert_expired")),1)):_("",!0)])]),_:2},1024)):(s(),r(u,{key:1},[n(l(c("common.collection.none")),1)],64))]),actions:e(({row:a})=>[i(A,null,{default:e(()=>[i(P,{to:{name:"data-plane-detail-view",params:{dataPlane:a.id}}},{default:e(()=>[n(l(c("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2},1032,["page-number","page-size","headers","items","total","is-selected-row","empty-state-message","empty-state-cta-to","empty-state-cta-text","onChange","onResize"]),n(),o.params.dataPlane?(s(),m(D,{key:0},{default:e(a=>[i(G,{onClose:b=>o.replace({name:o.name,params:{mesh:o.params.mesh},query:{page:o.params.page,size:o.params.size,s:o.params.s}})},{default:e(()=>[typeof t<"u"?(s(),m(q(a.Component),{key:0,items:t.items},null,8,["items"])):_("",!0)]),_:2},1032,["onClose"])]),_:2},1024)):_("",!0)]),_:2},1032,["src"])]),_:2},1024)])])]),_:2},1024)]),_:1})}}}),U=F(W,[["__scopeId","data-v-3ba40200"]]);export{U as default};
