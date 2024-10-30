import{d as f,e as s,o as r,m as _,w as a,a as t,l as k,an as C,A,$ as v,b as l,t as c,c as V,H as x}from"./index-CgC5RQPZ.js";const X=f({__name:"BuiltinGatewayListView",setup(D){return(B,L)=>{const m=s("XAction"),u=s("XActionGroup"),g=s("DataCollection"),d=s("DataLoader"),w=s("KCard"),h=s("AppView"),y=s("RouteView");return r(),_(y,{name:"builtin-gateway-list-view",params:{page:1,size:50,mesh:"",gateway:""}},{default:a(({route:n,t:p,can:z,me:i,uri:b})=>[t(h,{docs:p("builtin-gateways.href.docs")},{default:a(()=>[t(w,null,{default:a(()=>[t(d,{src:b(k(C),"/meshes/:mesh/mesh-gateways",{mesh:n.params.mesh},{page:n.params.page,size:n.params.size})},{loadable:a(({data:o})=>[t(g,{type:"gateways",items:(o==null?void 0:o.items)??[void 0],page:n.params.page,"page-size":n.params.size,total:o==null?void 0:o.total,onChange:n.update},{default:a(()=>[t(A,{class:"builtin-gateway-collection","data-testid":"builtin-gateway-collection",headers:[{...i.get("headers.name"),label:"Name",key:"name"},...z("use zones")?[{...i.get("headers.zone"),label:"Zone",key:"zone"}]:[],{...i.get("headers.actions"),label:"Actions",key:"actions",hideLabel:!0}],items:o==null?void 0:o.items,onResize:i.set},{name:a(({row:e})=>[t(v,{text:e.name},{default:a(()=>[t(m,{"data-action":"",to:{name:"builtin-gateway-detail-view",params:{mesh:e.mesh,gateway:e.id}}},{default:a(()=>[l(c(e.name),1)]),_:2},1032,["to"])]),_:2},1032,["text"])]),zone:a(({row:e})=>[e.labels&&e.labels["kuma.io/origin"]==="zone"&&e.labels["kuma.io/zone"]?(r(),_(m,{key:0,to:{name:"zone-cp-detail-view",params:{zone:e.labels["kuma.io/zone"]}}},{default:a(()=>[l(c(e.labels["kuma.io/zone"]),1)]),_:2},1032,["to"])):(r(),V(x,{key:1},[l(c(p("common.detail.none")),1)],64))]),actions:a(({row:e})=>[t(u,null,{default:a(()=>[t(m,{to:{name:"builtin-gateway-detail-view",params:{mesh:e.mesh,gateway:e.name}}},{default:a(()=>[l(c(p("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2},1032,["headers","items","onResize"])]),_:2},1032,["items","page","page-size","total","onChange"])]),_:2},1032,["src"])]),_:2},1024)]),_:2},1032,["docs"])]),_:1})}}});export{X as default};