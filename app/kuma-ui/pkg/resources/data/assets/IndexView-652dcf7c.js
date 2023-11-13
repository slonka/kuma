import{d as S,a as m,o,b as a,w as t,e as l,p as R,f as d,H as $,t as p,a3 as C,c as f,F as I,q as b,$ as P,K as B,J as D,v as V,_ as N}from"./index-dc1529df.js";import{A as T}from"./AppCollection-4e1d9b64.js";import{S as K}from"./StatusBadge-903f8974.js";import{S as q}from"./SummaryView-8ab015d9.js";import{g as E}from"./dataplane-0a086c06.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-79530b3b.js";const F=S({__name:"IndexView",setup(L){function h(k){return k.map(_=>{const{name:u}=_,g={name:"zone-ingress-detail-view",params:{zoneIngress:u}},{networking:e}=_.zoneIngress;let z;e!=null&&e.address&&(e!=null&&e.port)&&(z=`${e.address}:${e.port}`);let c;e!=null&&e.advertisedAddress&&(e!=null&&e.advertisedPort)&&(c=`${e.advertisedAddress}:${e.advertisedPort}`);const y=E(_.zoneIngressInsight??{});return{detailViewRoute:g,name:u,addressPort:z,advertisedAddressPort:c,status:y}})}return(k,_)=>{const u=m("RouteTitle"),g=m("RouterLink"),e=m("KCard"),z=m("RouterView"),c=m("DataSource"),y=m("AppView"),w=m("RouteView");return o(),a(c,{src:"/me"},{default:t(({data:A})=>[A?(o(),a(w,{key:0,name:"zone-ingress-list-view",params:{zone:"",zoneIngress:""}},{default:t(({route:n,t:r})=>[l(y,null,{title:t(()=>[R("h2",null,[l(u,{title:r("zone-ingresses.routes.items.title"),render:!0},null,8,["title"])])]),default:t(()=>[d(),l(c,{src:`/zone-cps/${n.params.zone}/ingresses?page=1&size=100`},{default:t(({data:i,error:v})=>[l(e,null,{body:t(()=>[v!==void 0?(o(),a($,{key:0,error:v},null,8,["error"])):(o(),a(T,{key:1,class:"zone-ingress-collection","data-testid":"zone-ingress-collection",headers:[{label:"Name",key:"name"},{label:"Address",key:"addressPort"},{label:"Advertised address",key:"advertisedAddressPort"},{label:"Status",key:"status"},{label:"Details",key:"details",hideLabel:!0}],"page-number":1,"page-size":100,total:i==null?void 0:i.total,items:i?h(i.items):void 0,error:v,"empty-state-message":r("common.emptyState.message",{type:"Zone Ingresses"}),"empty-state-cta-to":r("zone-ingresses.href.docs"),"empty-state-cta-text":r("common.documentation"),"is-selected-row":s=>s.name===n.params.zoneIngress,onChange:n.update},{name:t(({row:s})=>[l(g,{to:{name:"zone-ingress-summary-view",params:{zone:n.params.zone,zoneIngress:s.name},query:{page:1,size:100}}},{default:t(()=>[d(p(s.name),1)]),_:2},1032,["to"])]),addressPort:t(({rowValue:s})=>[s?(o(),a(C,{key:0,text:s},null,8,["text"])):(o(),f(I,{key:1},[d(p(r("common.collection.none")),1)],64))]),advertisedAddressPort:t(({rowValue:s})=>[s?(o(),a(C,{key:0,text:s},null,8,["text"])):(o(),f(I,{key:1},[d(p(r("common.collection.none")),1)],64))]),status:t(({rowValue:s})=>[s?(o(),a(K,{key:0,status:s},null,8,["status"])):(o(),f(I,{key:1},[d(p(r("common.collection.none")),1)],64))]),details:t(({row:s})=>[l(g,{class:"details-link","data-testid":"details-link",to:{name:"zone-ingress-detail-view",params:{zoneIngress:s.name}}},{default:t(()=>[d(p(r("common.collection.details_link"))+" ",1),l(b(P),{display:"inline-block",decorative:"",size:b(B)},null,8,["size"])]),_:2},1032,["to"])]),_:2},1032,["total","items","error","empty-state-message","empty-state-cta-to","empty-state-cta-text","is-selected-row","onChange"]))]),_:2},1024),d(),n.params.zoneIngress?(o(),a(z,{key:0},{default:t(s=>[l(q,{onClose:x=>n.replace({name:"zone-ingress-list-view",params:{zone:n.params.zone},query:{page:1,size:100}})},{default:t(()=>[(o(),a(D(s.Component),{name:n.params.zoneIngress,"zone-ingress-overview":i==null?void 0:i.items.find(x=>x.name===n.params.zoneIngress)},null,8,["name","zone-ingress-overview"]))]),_:2},1032,["onClose"])]),_:2},1024)):V("",!0)]),_:2},1032,["src"])]),_:2},1024)]),_:1},8,["params"])):V("",!0)]),_:1})}}});const j=N(F,[["__scopeId","data-v-208a808b"]]);export{j as default};
