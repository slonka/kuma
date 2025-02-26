import{d as b,r as a,o as A,m as V,w as e,b as t,p as X,I as D,s as k,e as l,C as R,t as m}from"./index-DzNnjnmE.js";const L=b({__name:"MeshListView",setup(B){return(T,c)=>{const _=a("RouteTitle"),d=a("XI18n"),p=a("XAction"),h=a("XActionGroup"),u=a("DataCollection"),g=a("DataLoader"),w=a("XCard"),f=a("AppView"),v=a("DataSource"),y=a("RouteView");return A(),V(y,{name:"mesh-list-view",params:{page:1,size:Number,mesh:""}},{default:e(({route:n,t:i,me:r,uri:C})=>[t(v,{src:C(X(D),"/mesh-insights",{},{page:n.params.page,size:n.params.size})},{default:e(({data:s,error:z})=>[t(f,{docs:s!=null&&s.items.length?i("meshes.href.docs"):""},{title:e(()=>[k("h1",null,[t(_,{title:i("meshes.routes.items.title")},null,8,["title"])])]),default:e(()=>[c[3]||(c[3]=l()),t(d,{path:"meshes.routes.items.intro","default-path":"common.i18n.ignore-error"}),c[4]||(c[4]=l()),t(w,null,{default:e(()=>[t(g,{data:[s],errors:[z]},{loadable:e(()=>[t(u,{type:"meshes",items:(s==null?void 0:s.items)??[void 0],page:n.params.page,"page-size":n.params.size,total:s==null?void 0:s.total,onChange:n.update},{default:e(()=>[t(R,{class:"mesh-collection","data-testid":"mesh-collection",headers:[{...r.get("headers.name"),label:i("meshes.common.name"),key:"name"},{...r.get("headers.services"),label:i("meshes.routes.items.collection.services"),key:"services"},{...r.get("headers.dataplanes"),label:i("meshes.routes.items.collection.dataplanes"),key:"dataplanes"},{...r.get("headers.actions"),label:"Actions",key:"actions",hideLabel:!0}],items:s==null?void 0:s.items,"is-selected-row":o=>o.name===n.params.mesh,onResize:r.set},{name:e(({row:o})=>[t(p,{"data-action":"",to:{name:"mesh-detail-view",params:{mesh:o.name},query:{page:n.params.page,size:n.params.size}}},{default:e(()=>[l(m(o.name),1)]),_:2},1032,["to"])]),services:e(({row:o})=>[l(m(o.services.internal),1)]),dataplanes:e(({row:o})=>[l(m(o.dataplanesByType.standard.online)+" / "+m(o.dataplanesByType.standard.total),1)]),actions:e(({row:o})=>[t(h,null,{default:e(()=>[t(p,{to:{name:"mesh-detail-view",params:{mesh:o.name}}},{default:e(()=>[l(m(i("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2},1032,["headers","items","is-selected-row","onResize"])]),_:2},1032,["items","page","page-size","total","onChange"])]),_:2},1032,["data","errors"])]),_:2},1024)]),_:2},1032,["docs"])]),_:2},1032,["src"])]),_:1})}}});export{L as default};
