import{d as x,r as a,o,i as p,w as e,j as s,p as z,n as i,E as K,a0 as V,H as c,a5 as B,l as _,F as u,W as D,k as y,a2 as I,K as T,m as L,t as N}from"./index-adcc6fc8.js";const R=x({__name:"ServiceListView",setup(A){return(E,$)=>{const w=a("RouteTitle"),k=a("RouterLink"),g=a("KButton"),b=a("KDropdownItem"),h=a("KDropdownMenu"),f=a("KCard"),C=a("AppView"),d=a("DataSource"),S=a("RouteView");return o(),p(d,{src:"/me"},{default:e(({data:v})=>[v?(o(),p(S,{key:0,name:"service-list-view",params:{page:1,size:v.pageSize,mesh:""}},{default:e(({route:n,t:r})=>[s(d,{src:`/meshes/${n.params.mesh}/service-insights?page=${n.params.page}&size=${n.params.size}`},{default:e(({data:l,error:m})=>[s(C,null,{title:e(()=>[z("h2",null,[s(w,{title:r("services.routes.items.title"),render:!0},null,8,["title"])])]),default:e(()=>[i(),s(f,null,{body:e(()=>[m!==void 0?(o(),p(K,{key:0,error:m},null,8,["error"])):(o(),p(V,{key:1,class:"service-collection","data-testid":"service-collection","empty-state-message":r("common.emptyState.message",{type:"Services"}),headers:[{label:"Name",key:"name"},{label:"Type",key:"serviceType"},{label:"Address",key:"addressPort"},{label:"DP proxies (online / total)",key:"online"},{label:"Status",key:"status"},{label:"Actions",key:"actions",hideLabel:!0}],"page-number":parseInt(n.params.page),"page-size":parseInt(n.params.size),total:l==null?void 0:l.total,items:l==null?void 0:l.items,error:m,onChange:n.update},{name:e(({row:t})=>[s(k,{to:{name:"service-detail-view",params:{service:t.name}}},{default:e(()=>[i(c(t.name),1)]),_:2},1032,["to"])]),serviceType:e(({rowValue:t})=>[i(c(t||"internal"),1)]),addressPort:e(({rowValue:t})=>[t?(o(),p(B,{key:0,text:t},null,8,["text"])):(o(),_(u,{key:1},[i(c(r("common.collection.none")),1)],64))]),online:e(({row:t})=>[t.dataplanes?(o(),_(u,{key:0},[i(c(t.dataplanes.online||0)+" / "+c(t.dataplanes.total||0),1)],64)):(o(),_(u,{key:1},[i(c(r("common.collection.none")),1)],64))]),status:e(({row:t})=>[s(D,{status:t.status||"not_available"},null,8,["status"])]),actions:e(({row:t})=>[s(h,{class:"actions-dropdown","kpop-attributes":{placement:"bottomEnd",popoverClasses:"mt-5 more-actions-popover"},width:"150"},{default:e(()=>[s(g,{class:"non-visual-button",appearance:"secondary",size:"small"},{default:e(()=>[s(y(I),{size:y(T)},null,8,["size"])]),_:1})]),items:e(()=>[s(b,{item:{to:{name:"service-detail-view",params:{service:t.name}},label:r("common.collection.actions.view")}},null,8,["item"])]),_:2},1024)]),_:2},1032,["empty-state-message","headers","page-number","page-size","total","items","error","onChange"]))]),_:2},1024)]),_:2},1024)]),_:2},1032,["src"])]),_:2},1032,["params"])):L("",!0)]),_:1})}}});const F=N(R,[["__scopeId","data-v-72523eb5"]]);export{F as default};
