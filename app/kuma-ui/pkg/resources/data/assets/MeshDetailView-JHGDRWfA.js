import{d as $,m as d,o as p,w as e,b as t,e as o,r as i,c as f,v as T,F as _,L as V,t as r,R as w,p as E,s as A,I as M}from"./index-DFpFnkh8.js";import{_ as I}from"./ResourceCodeBlock.vue_vue_type_script_setup_true_lang-Cmq1Ydcr.js";const q=$({__name:"MeshDetailView",props:{mesh:{}},setup(x){const a=x;return(z,s)=>{const L=i("RouteTitle"),C=i("XI18n"),X=i("XNotification"),R=i("XAction"),h=i("XBadge"),D=i("XAboutCard"),y=i("XLayout"),g=i("XCard"),B=i("DataSource"),N=i("AppView"),P=i("RouteView");return p(),d(P,{name:"mesh-detail-view",params:{mesh:""}},{default:e(({route:k,t:l,uri:b})=>[t(L,{title:l("meshes.routes.overview.title"),render:!1},null,8,["title"]),s[9]||(s[9]=o()),t(B,{src:b(A(M),"/mesh-insights/:name",{name:k.params.mesh})},{default:e(({data:n})=>[t(N,{docs:l("meshes.href.docs"),notifications:!0},{default:e(()=>[t(X,{notify:!a.mesh.mtlsBackend,uri:`meshes.notifications.mtls-warning:${a.mesh.id}`},{default:e(()=>[t(C,{path:"meshes.notifications.mtls-warning"})]),_:1},8,["notify","uri"]),s[8]||(s[8]=o()),t(y,{type:"stack"},{default:e(()=>[t(D,{title:l("meshes.routes.item.about.title"),created:a.mesh.creationTime,modified:a.mesh.modificationTime},{default:e(()=>[(p(),f(_,null,T(["MeshTrafficPermission","MeshMetric","MeshAccessLog","MeshTrace"],m=>{var c;return p(),f(_,{key:m},[(p(!0),f(_,null,T([((c=n==null?void 0:n.policies)==null?void 0:c[m])??{total:0}],u=>(p(),f(_,{key:typeof u},[t(X,{notify:m==="MeshTrafficPermission"&&a.mesh.mtlsBackend&&u.total===0,uri:`meshes.notifications.mtp-warning:${a.mesh.id}`},{default:e(()=>[t(C,{path:"meshes.notifications.mtp-warning"})]),_:2},1032,["notify","uri"]),s[1]||(s[1]=o()),t(V,{layout:"horizontal"},{title:e(()=>[t(R,{to:{name:"policy-list-view",params:{mesh:k.params.mesh,policyPath:`${m.toLowerCase()}s`}}},{default:e(()=>[o(r(m),1)]),_:2},1032,["to"])]),body:e(()=>[t(h,{appearance:u.total>0?"success":"neutral"},{default:e(()=>[o(r(u.total>0?l("meshes.detail.enabled"):l("meshes.detail.disabled")),1)]),_:2},1032,["appearance"])]),_:2},1024)],64))),128))],64)}),64)),s[3]||(s[3]=o()),t(V,{layout:"horizontal"},{title:e(()=>[o(r(l("http.api.property.mtls")),1)]),body:e(()=>[a.mesh.mtlsBackend?(p(),d(h,{key:1,appearance:"info"},{default:e(()=>[o(r(a.mesh.mtlsBackend.type)+" / "+r(a.mesh.mtlsBackend.name),1)]),_:1})):(p(),d(h,{key:0,appearance:"neutral"},{default:e(()=>[o(r(l("meshes.detail.disabled")),1)]),_:2},1024))]),_:2},1024)]),_:2},1032,["title","created","modified"]),s[6]||(s[6]=o()),t(g,null,{default:e(()=>[t(y,{type:"stack"},{default:e(()=>[t(y,{type:"columns",class:"columns-with-borders"},{default:e(()=>[t(w,{total:(n==null?void 0:n.services.total)??0,"data-testid":"services-status"},{title:e(()=>[o(r(l("meshes.detail.services")),1)]),_:2},1032,["total"]),s[4]||(s[4]=o()),t(w,{total:(n==null?void 0:n.dataplanesByType.standard.total)??0,online:(n==null?void 0:n.dataplanesByType.standard.online)??0,"data-testid":"data-plane-proxies-status"},{title:e(()=>[o(r(l("meshes.detail.data_plane_proxies")),1)]),_:2},1032,["total","online"]),s[5]||(s[5]=o()),t(w,{total:(n==null?void 0:n.totalPolicyCount)??0,"data-testid":"policies-status"},{title:e(()=>[o(r(l("meshes.detail.policies")),1)]),_:2},1032,["total"])]),_:2},1024)]),_:2},1024)]),_:2},1024),s[7]||(s[7]=o()),t(g,null,{default:e(()=>[t(I,{resource:a.mesh.config},{default:e(({copy:m,copying:c})=>[c?(p(),d(B,{key:0,src:b(A(M),"/meshes/:name/as/kubernetes",{name:k.params.mesh},{cacheControl:"no-store"}),onChange:u=>{m(v=>v(u))},onError:u=>{m((v,S)=>S(u))}},null,8,["src","onChange","onError"])):E("",!0)]),_:2},1032,["resource"])]),_:2},1024)]),_:2},1024)]),_:2},1032,["docs"])]),_:2},1032,["src"])]),_:1})}}});export{q as default};
