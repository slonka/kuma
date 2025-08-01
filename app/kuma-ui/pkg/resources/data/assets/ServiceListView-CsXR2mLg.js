import{d as R,r as n,p as d,o as r,w as a,b as t,e as i,l as h,I as B,m as D,af as L,q as N,K as P,t as m,S as T,c as _,N as u,y as q,_ as G}from"./index-Dc7zXr8g.js";import{S as I}from"./SummaryView-B-iORA4W.js";const E=R({__name:"ServiceListView",setup(F){return(K,o)=>{const y=n("RouteTitle"),w=n("XSearch"),v=n("XAction"),f=n("XCopyButton"),C=n("XActionGroup"),k=n("RouterView"),V=n("DataCollection"),z=n("DataLoader"),X=n("XLayout"),S=n("XCard"),x=n("AppView"),b=n("RouteView");return r(),d(b,{name:"service-list-view",params:{page:1,size:Number,mesh:"",service:"",s:""}},{default:a(({route:s,t:p,uri:A,me:l})=>[t(y,{render:!1,title:p("services.routes.items.title")},null,8,["title"]),o[7]||(o[7]=i()),t(x,{docs:p("services.href.docs")},{default:a(()=>[t(S,null,{default:a(()=>[t(X,null,{default:a(()=>[h("search",null,[h("form",{onSubmit:o[0]||(o[0]=B(()=>{},["prevent"]))},[t(w,{class:"search-field",keys:["name"],value:s.params.s,onChange:c=>s.update({page:1,s:c})},null,8,["value","onChange"])],32)]),o[6]||(o[6]=i()),t(z,{src:A(D(L),"/meshes/:mesh/service-insights/of/:serviceType",{mesh:s.params.mesh,serviceType:"internal"},{page:s.params.page,size:s.params.size,search:s.params.s}),variant:"list"},{default:a(({data:c})=>[t(V,{type:"services",items:c.items,page:s.params.page,"page-size":s.params.size,total:c.total,onChange:s.update},{default:a(()=>[t(P,{class:"service-collection","data-testid":"service-collection",headers:[{...l.get("headers.name"),label:"Name",key:"name"},{...l.get("headers.addressPort"),label:"Address",key:"addressPort"},{...l.get("headers.online"),label:"DP proxies (online / total)",key:"online"},{...l.get("headers.status"),label:"Status",key:"status"},{...l.get("headers.actions"),label:"Actions",key:"actions",hideLabel:!0}],items:c.items,"is-selected-row":e=>e.name===s.params.service,onResize:l.set},{name:a(({row:e})=>[t(f,{text:e.name},{default:a(()=>[t(v,{"data-action":"",to:{name:"service-detail-view",params:{mesh:e.mesh,service:e.name},query:{page:s.params.page,size:s.params.size}}},{default:a(()=>[i(m(e.name),1)]),_:2},1032,["to"])]),_:2},1032,["text"])]),addressPort:a(({row:e})=>[e.addressPort?(r(),d(f,{key:0,text:e.addressPort},null,8,["text"])):(r(),_(u,{key:1},[i(m(p("common.collection.none")),1)],64))]),online:a(({row:e})=>[e.dataplanes?(r(),_(u,{key:0},[i(m(e.dataplanes.online||0)+" / "+m(e.dataplanes.total||0),1)],64)):(r(),_(u,{key:1},[i(m(p("common.collection.none")),1)],64))]),status:a(({row:e})=>[t(T,{status:e.status},null,8,["status"])]),actions:a(({row:e})=>[t(C,null,{default:a(()=>[t(v,{to:{name:"service-detail-view",params:{mesh:e.mesh,service:e.name}}},{default:a(()=>[i(m(p("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2,__:[1,2,3,4]},1032,["headers","items","is-selected-row","onResize"]),o[5]||(o[5]=i()),s.params.service?(r(),d(k,{key:0},{default:a(e=>[t(I,{onClose:g=>s.replace({name:"service-list-view",params:{mesh:s.params.mesh},query:{page:s.params.page,size:s.params.size}})},{default:a(()=>[(r(),d(q(e.Component),{name:s.params.service,service:c.items.find(g=>g.name===s.params.service)},null,8,["name","service"]))]),_:2},1032,["onClose"])]),_:2},1024)):N("",!0)]),_:2,__:[5]},1032,["items","page","page-size","total","onChange"])]),_:2},1032,["src"])]),_:2,__:[6]},1024)]),_:2},1024)]),_:2},1032,["docs"])]),_:1,__:[7]})}}}),j=G(E,[["__scopeId","data-v-958f8131"]]);export{j as default};
