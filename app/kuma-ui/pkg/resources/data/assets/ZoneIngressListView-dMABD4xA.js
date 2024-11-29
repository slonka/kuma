import{d as I,e as a,o as d,p as m,w as t,a as s,b as r,l as b,m as x,au as R,A as D,t as p,a0 as _,c as z,J as k,S as L,F as B,q as w}from"./index-DGJfQfev.js";import{S as N}from"./SummaryView-Bw7MaIhE.js";const T=["innerHTML"],H=I({__name:"ZoneIngressListView",props:{data:{}},setup(X){return(q,l)=>{const f=a("RouteTitle"),u=a("XAction"),A=a("XActionGroup"),g=a("RouterView"),y=a("DataCollection"),v=a("DataLoader"),C=a("KCard"),V=a("AppView"),h=a("RouteView");return d(),m(h,{name:"zone-ingress-list-view",params:{zone:"",zoneIngress:""}},{default:t(({route:n,t:i,me:c,uri:S})=>[s(f,{render:!1,title:i("zone-ingresses.routes.items.title")},null,8,["title"]),l[6]||(l[6]=r()),s(V,{docs:i("zone-ingresses.href.docs")},{default:t(()=>[b("div",{innerHTML:i("zone-ingresses.routes.items.intro",{},{defaultMessage:""})},null,8,T),l[5]||(l[5]=r()),s(C,null,{default:t(()=>[s(v,{src:S(x(R),"/zone-cps/:name/ingresses",{name:n.params.zone},{page:1,size:100})},{loadable:t(({data:o})=>[s(y,{type:"zone-ingresses",items:(o==null?void 0:o.items)??[void 0],total:o==null?void 0:o.total,onChange:n.update},{default:t(()=>[s(D,{class:"zone-ingress-collection","data-testid":"zone-ingress-collection",headers:[{...c.get("headers.name"),label:"Name",key:"name"},{...c.get("headers.socketAddress"),label:"Address",key:"socketAddress"},{...c.get("headers.advertisedSocketAddress"),label:"Advertised address",key:"advertisedSocketAddress"},{...c.get("headers.status"),label:"Status",key:"status"},{...c.get("headers.actions"),label:"Actions",key:"actions",hideLabel:!0}],items:o==null?void 0:o.items,"is-selected-row":e=>e.name===n.params.zoneIngress,onResize:c.set},{name:t(({row:e})=>[s(u,{"data-action":"",to:{name:"zone-ingress-summary-view",params:{zone:n.params.zone,zoneIngress:e.id},query:{page:1,size:100}}},{default:t(()=>[r(p(e.name),1)]),_:2},1032,["to"])]),socketAddress:t(({row:e})=>[e.zoneIngress.socketAddress.length>0?(d(),m(_,{key:0,text:e.zoneIngress.socketAddress},null,8,["text"])):(d(),z(k,{key:1},[r(p(i("common.collection.none")),1)],64))]),advertisedSocketAddress:t(({row:e})=>[e.zoneIngress.advertisedSocketAddress.length>0?(d(),m(_,{key:0,text:e.zoneIngress.advertisedSocketAddress},null,8,["text"])):(d(),z(k,{key:1},[r(p(i("common.collection.none")),1)],64))]),status:t(({row:e})=>[s(L,{status:e.state},null,8,["status"])]),actions:t(({row:e})=>[s(A,null,{default:t(()=>[s(u,{to:{name:"zone-ingress-detail-view",params:{zoneIngress:e.id}}},{default:t(()=>[r(p(i("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2},1032,["headers","items","is-selected-row","onResize"]),l[4]||(l[4]=r()),n.child()?(d(),m(g,{key:0},{default:t(({Component:e})=>[s(N,{onClose:M=>n.replace({name:"zone-ingress-list-view",params:{zone:n.params.zone},query:{page:1,size:100}})},{default:t(()=>[typeof o<"u"?(d(),m(B(e),{key:0,items:o.items},null,8,["items"])):w("",!0)]),_:2},1032,["onClose"])]),_:2},1024)):w("",!0)]),_:2},1032,["items","total","onChange"])]),_:2},1032,["src"])]),_:2},1024)]),_:2},1032,["docs"])]),_:1})}}});export{H as default};
