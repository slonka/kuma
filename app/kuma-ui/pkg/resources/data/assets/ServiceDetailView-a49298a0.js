import{d as g,a2 as S,o as t,j as y,h as i,w as e,m as x,a8 as _,l as a,D as l,i as o,g as r,k as D,Y as w,U as I,a4 as $,F as b,ad as B,r as h,E as C,t as T,ae as E}from"./index-70c86013.js";import{T as N}from"./TagList-59f59704.js";const P={class:"stack"},R={class:"columns",style:{"--columns":"2"}},A=g({__name:"ExternalServiceDetails",props:{externalService:{}},setup(c){const s=c,{t:n}=S();return(k,u)=>(t(),y("div",P,[i(o(w),null,{body:e(()=>[x("div",R,[i(_,null,{title:e(()=>[a(l(o(n)("http.api.property.address")),1)]),body:e(()=>[a(l(s.externalService.networking.address),1)]),_:1}),a(),s.externalService.tags!==null?(t(),r(_,{key:0},{title:e(()=>[a(l(o(n)("http.api.property.tags")),1)]),body:e(()=>[i(N,{tags:s.externalService.tags},null,8,["tags"])]),_:1})):D("",!0)])]),_:1})]))}}),F={class:"stack"},j={class:"columns",style:{"--columns":"3"}},L=g({__name:"ServiceInsightDetails",props:{serviceInsight:{}},setup(c){const s=c,{t:n}=S();return(k,u)=>(t(),y("div",F,[i(o(w),null,{body:e(()=>{var p,d;return[x("div",j,[i(_,null,{title:e(()=>[a(l(o(n)("http.api.property.status")),1)]),body:e(()=>[i(I,{status:s.serviceInsight.status??"not_available"},null,8,["status"])]),_:1}),a(),i(_,null,{title:e(()=>[a(l(o(n)("http.api.property.address")),1)]),body:e(()=>[s.serviceInsight.addressPort?(t(),r($,{key:0,text:s.serviceInsight.addressPort},null,8,["text"])):(t(),y(b,{key:1},[a(l(o(n)("common.detail.none")),1)],64))]),_:1}),a(),i(B,{online:((p=s.serviceInsight.dataplanes)==null?void 0:p.online)??0,total:((d=s.serviceInsight.dataplanes)==null?void 0:d.total)??0},{title:e(()=>[a(l(o(n)("http.api.property.dataPlaneProxies")),1)]),_:1},8,["online","total"])])]}),_:1})]))}}),Y=g({__name:"ServiceDetailView",props:{data:{}},setup(c){const s=c;return(n,k)=>{const u=h("DataSource"),p=h("AppView"),d=h("RouteView");return t(),r(d,{name:"service-detail-view",params:{mesh:"",service:""}},{default:e(({route:m,t:V})=>[i(p,null,{default:e(()=>[s.data.serviceType==="external"?(t(),r(u,{key:0,src:`/meshes/${m.params.mesh}/external-services/for/${m.params.service}`},{default:e(({data:v,error:f})=>[f?(t(),r(C,{key:0,error:f},null,8,["error"])):v===void 0?(t(),r(T,{key:1})):v===null?(t(),r(E,{key:2,"data-testid":"no-matching-external-service"},{title:e(()=>[x("p",null,l(V("services.detail.no_matching_external_service",{name:m.params.service})),1)]),_:2},1024)):(t(),r(A,{key:3,"external-service":v},null,8,["external-service"]))]),_:2},1032,["src"])):(t(),r(L,{key:1,"service-insight":n.data},null,8,["service-insight"]))]),_:2},1024)]),_:1})}}});export{Y as default};