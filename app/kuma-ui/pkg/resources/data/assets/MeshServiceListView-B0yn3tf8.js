import{d as D,r as l,o as t,m as p,w as s,b as n,e as m,l as K,aB as L,A as N,U as S,t as i,c as u,F as _,p as z,s as f,E as X}from"./index-DHg9Fngg.js";import{S as $}from"./SummaryView-DmzBN36G.js";const M=D({__name:"MeshServiceListView",setup(P){return(q,E)=>{const k=l("RouteTitle"),d=l("XAction"),h=l("KTruncate"),w=l("KBadge"),y=l("XActionGroup"),b=l("RouterView"),C=l("DataCollection"),V=l("DataLoader"),A=l("KCard"),R=l("AppView"),T=l("RouteView");return t(),p(T,{name:"mesh-service-list-view",params:{page:1,size:50,mesh:"",service:""}},{default:s(({route:o,t:g,can:x,uri:B,me:c})=>[n(k,{render:!1,title:g("services.routes.mesh-service-list-view.title")},null,8,["title"]),m(),n(R,null,{default:s(()=>[n(A,null,{default:s(()=>[n(V,{src:B(K(L),"/meshes/:mesh/mesh-services",{mesh:o.params.mesh},{page:o.params.page,size:o.params.size})},{loadable:s(({data:a})=>[n(C,{type:"services",items:(a==null?void 0:a.items)??[void 0]},{default:s(()=>[n(N,{"data-testid":"service-collection",headers:[{...c.get("headers.name"),label:"Name",key:"name"},{...c.get("headers.namespace"),label:"Namespace",key:"namespace"},...x("use zones")?[{...c.get("headers.zone"),label:"Zone",key:"zone"}]:[],{...c.get("headers.addresses"),label:"Addresses",key:"addresses"},{...c.get("headers.ports"),label:"Ports",key:"ports"},{...c.get("headers.tags"),label:"Tags",key:"tags"},{...c.get("headers.actions"),label:"Actions",key:"actions",hideLabel:!0}],"page-number":o.params.page,"page-size":o.params.size,total:a==null?void 0:a.total,items:a==null?void 0:a.items,"is-selected-row":e=>e.name===o.params.service,onChange:o.update,onResize:c.set},{name:s(({row:e})=>[n(S,{text:e.name},{default:s(()=>[n(d,{"data-action":"",to:{name:"mesh-service-summary-view",params:{mesh:e.mesh,service:e.id},query:{page:o.params.page,size:o.params.size}}},{default:s(()=>[m(i(e.name),1)]),_:2},1032,["to"])]),_:2},1032,["text"])]),namespace:s(({row:e})=>[m(i(e.namespace),1)]),zone:s(({row:e})=>[e.labels&&e.labels["kuma.io/origin"]==="zone"&&e.labels["kuma.io/zone"]?(t(),u(_,{key:0},[e.labels["kuma.io/zone"]?(t(),p(d,{key:0,to:{name:"zone-cp-detail-view",params:{zone:e.labels["kuma.io/zone"]}}},{default:s(()=>[m(i(e.labels["kuma.io/zone"]),1)]),_:2},1032,["to"])):z("",!0)],64)):(t(),u(_,{key:1},[m(i(g("common.detail.none")),1)],64))]),addresses:s(({row:e})=>[n(h,null,{default:s(()=>[(t(!0),u(_,null,f(e.status.addresses,r=>(t(),u("span",{key:r.hostname},i(r.hostname),1))),128))]),_:2},1024)]),ports:s(({row:e})=>[n(h,null,{default:s(()=>[(t(!0),u(_,null,f(e.spec.ports,r=>(t(),p(w,{key:r.port,appearance:"info"},{default:s(()=>[m(i(r.port)+":"+i(r.targetPort)+"/"+i(r.appProtocol),1)]),_:2},1024))),128))]),_:2},1024)]),tags:s(({row:e})=>[n(h,null,{default:s(()=>[(t(!0),u(_,null,f(e.spec.selector.dataplaneTags,(r,v)=>(t(),p(w,{key:`${v}:${r}`,appearance:"info"},{default:s(()=>[m(i(v)+":"+i(r),1)]),_:2},1024))),128))]),_:2},1024)]),actions:s(({row:e})=>[n(y,null,{default:s(()=>[n(d,{to:{name:"mesh-service-detail-view",params:{mesh:e.mesh,service:e.id}}},{default:s(()=>[m(i(g("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2},1032,["headers","page-number","page-size","total","items","is-selected-row","onChange","onResize"]),m(),a!=null&&a.items&&o.params.service?(t(),p(b,{key:0},{default:s(e=>[n($,{onClose:r=>o.replace({name:"mesh-service-list-view",params:{mesh:o.params.mesh},query:{page:o.params.page,size:o.params.size}})},{default:s(()=>[(t(),p(X(e.Component),{items:a==null?void 0:a.items},null,8,["items"]))]),_:2},1032,["onClose"])]),_:2},1024)):z("",!0)]),_:2},1032,["items"])]),_:2},1032,["src"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{M as default};
