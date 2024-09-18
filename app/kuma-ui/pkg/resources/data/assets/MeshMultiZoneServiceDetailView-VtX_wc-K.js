import{d as C,r as n,o as r,m,w as a,b as t,k as i,Z as f,e as s,c as w,L as S,M as g,t as o,A as B,H as K,J as M,q as N}from"./index-BZS2OlAU.js";const R=l=>(K("data-v-41029d69"),l=l(),M(),l),D={class:"stack"},I={class:"columns"},L=R(()=>i("h3",null,`
            Mesh Services
          `,-1)),X=C({__name:"MeshMultiZoneServiceDetailView",props:{data:{}},setup(l){const z=l;return(d,Z)=>{const u=n("KBadge"),_=n("KTruncate"),h=n("KCard"),p=n("XAction"),k=n("XActionGroup"),b=n("AppView"),A=n("RouteView");return r(),m(A,{name:"mesh-multi-zone-service-detail-view",params:{mesh:"",service:"",page:1,size:50,s:"",codeSearch:"",codeFilter:!1,codeRegExp:!1}},{default:a(({t:V,me:c,can:y})=>[t(b,null,{default:a(()=>[i("div",D,[t(h,null,{default:a(()=>[i("div",I,[t(f,null,{title:a(()=>[s(`
                Ports
              `)]),body:a(()=>[t(_,null,{default:a(()=>[(r(!0),w(S,null,g(z.data.spec.ports,e=>(r(),m(u,{key:e.port,appearance:"info"},{default:a(()=>[s(o(e.port)+"/"+o(e.appProtocol)+o(e.name&&e.name!==String(e.port)?` (${e.name})`:""),1)]),_:2},1024))),128))]),_:1})]),_:1}),s(),t(f,null,{title:a(()=>[s(`
                Selector
              `)]),body:a(()=>[t(_,null,{default:a(()=>[(r(!0),w(S,null,g(d.data.spec.selector.meshService.matchLabels,(e,v)=>(r(),m(u,{key:`${v}:${e}`,appearance:"info"},{default:a(()=>[s(o(v)+":"+o(e),1)]),_:2},1024))),128))]),_:1})]),_:1})])]),_:1}),s(),i("div",null,[L,s(),t(h,{class:"mt-4"},{default:a(()=>[t(B,{"data-testid":"mesh-service-collection",headers:[{...c.get("headers.name"),label:"Name",key:"name"},...y("use zones")?[{...c.get("headers.zone"),label:"Zone",key:"zone"}]:[],{...c.get("headers.namespace"),label:"Namespace",key:"namespace"},{...c.get("headers.actions"),label:"Actions",key:"actions",hideLabel:!0}],items:d.data.status.meshServices,total:d.data.status.meshServices.length,onResize:c.set},{name:a(({row:e})=>[t(p,{"data-action":"",class:"name-link",to:{name:"mesh-service-detail-view",params:{mesh:e.mesh,service:e.name}}},{default:a(()=>[s(o(e.name),1)]),_:2},1032,["to"])]),zone:a(({row:e})=>[t(p,{to:{name:"zone-cp-detail-view",params:{zone:e.zone}}},{default:a(()=>[s(o(e.zone),1)]),_:2},1032,["to"])]),namespace:a(({row:e})=>[s(o(e.namespace),1)]),actions:a(({row:e})=>[t(k,null,{default:a(()=>[t(p,{to:{name:"mesh-service-detail-view",params:{mesh:e.mesh,service:e.name}}},{default:a(()=>[s(o(V("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2},1032,["headers","items","total","onResize"])]),_:2},1024)])])]),_:2},1024)]),_:1})}}}),T=N(X,[["__scopeId","data-v-41029d69"]]);export{T as default};
