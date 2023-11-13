import{d as z,a as l,o as t,b as r,w as a,e as i,p as V,f as o,H as x,t as m,a3 as R,c as _,F as d,q as g,$ as B,K as T,J as D,v as f,_ as I}from"./index-dc1529df.js";import{A as L}from"./AppCollection-4e1d9b64.js";import{S as N}from"./StatusBadge-903f8974.js";import{S as A}from"./SummaryView-8ab015d9.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-79530b3b.js";const $=z({__name:"ServiceListView",setup(K){return(q,E)=>{const h=l("RouteTitle"),u=l("RouterLink"),w=l("KCard"),C=l("RouterView"),S=l("AppView"),v=l("DataSource"),b=l("RouteView");return t(),r(v,{src:"/me"},{default:a(({data:y})=>[y?(t(),r(b,{key:0,name:"service-list-view",params:{page:1,size:y.pageSize,mesh:"",service:""}},{default:a(({route:s,t:c})=>[i(v,{src:`/meshes/${s.params.mesh}/service-insights?page=${s.params.page}&size=${s.params.size}`},{default:a(({data:n,error:p})=>[i(S,null,{title:a(()=>[V("h2",null,[i(h,{title:c("services.routes.items.title"),render:!0},null,8,["title"])])]),default:a(()=>[o(),i(w,null,{body:a(()=>[p!==void 0?(t(),r(x,{key:0,error:p},null,8,["error"])):(t(),r(L,{key:1,class:"service-collection","data-testid":"service-collection","empty-state-message":c("common.emptyState.message",{type:"Services"}),headers:[{label:"Name",key:"name"},{label:"Type",key:"serviceType"},{label:"Address",key:"addressPort"},{label:"DP proxies (online / total)",key:"online"},{label:"Status",key:"status"},{label:"Details",key:"details",hideLabel:!0}],"page-number":parseInt(s.params.page),"page-size":parseInt(s.params.size),total:n==null?void 0:n.total,items:n==null?void 0:n.items,error:p,"is-selected-row":e=>e.name===s.params.service,onChange:s.update},{name:a(({row:e})=>[i(u,{to:{name:"service-detail-view",params:{mesh:e.mesh,service:e.name},query:{page:s.params.page,size:s.params.size}}},{default:a(()=>[o(m(e.name),1)]),_:2},1032,["to"])]),serviceType:a(({rowValue:e})=>[o(m(e||"internal"),1)]),addressPort:a(({rowValue:e})=>[e?(t(),r(R,{key:0,text:e},null,8,["text"])):(t(),_(d,{key:1},[o(m(c("common.collection.none")),1)],64))]),online:a(({row:e})=>[e.dataplanes?(t(),_(d,{key:0},[o(m(e.dataplanes.online||0)+" / "+m(e.dataplanes.total||0),1)],64)):(t(),_(d,{key:1},[o(m(c("common.collection.none")),1)],64))]),status:a(({row:e})=>[i(N,{status:e.status||"not_available"},null,8,["status"])]),details:a(({row:e})=>[i(u,{class:"details-link","data-testid":"details-link",to:{name:"service-detail-view",params:{mesh:e.mesh,service:e.name}}},{default:a(()=>[o(m(c("common.collection.details_link"))+" ",1),i(g(B),{display:"inline-block",decorative:"",size:g(T)},null,8,["size"])]),_:2},1032,["to"])]),_:2},1032,["empty-state-message","headers","page-number","page-size","total","items","error","is-selected-row","onChange"]))]),_:2},1024),o(),s.params.service?(t(),r(C,{key:0},{default:a(e=>[i(A,{onClose:k=>s.replace({name:"service-list-view",params:{mesh:s.params.mesh},query:{page:s.params.page,size:s.params.size}})},{default:a(()=>[(t(),r(D(e.Component),{name:s.params.service,service:n==null?void 0:n.items.find(k=>k.name===s.params.service)},null,8,["name","service"]))]),_:2},1032,["onClose"])]),_:2},1024)):f("",!0)]),_:2},1024)]),_:2},1032,["src"])]),_:2},1032,["params"])):f("",!0)]),_:1})}}});const U=I($,[["__scopeId","data-v-78e23fed"]]);export{U as default};
