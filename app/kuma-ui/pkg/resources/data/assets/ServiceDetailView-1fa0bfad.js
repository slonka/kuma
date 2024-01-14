import{d as f,k as y,a as d,o as t,c as u,e as r,w as e,b as i,m as k,t as c,l,X as h,f as n,p as w,F as V,R as C}from"./index-0dcf85b4.js";import{_ as I}from"./EmptyBlock.vue_vue_type_script_setup_true_lang-6ff90078.js";import{E as $}from"./ErrorBlock-70a14ad9.js";import{_ as B}from"./LoadingBlock.vue_vue_type_script_setup_true_lang-2e81faa2.js";import{T as D}from"./TagList-5868dd55.js";import{T as x}from"./TextWithCopyButton-da18080b.js";import{S as T}from"./StatusBadge-e7326f60.js";import"./index-fce48c05.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-d54f4898.js";import"./CopyButton-dcdce1f2.js";const b={key:3,class:"columns"},S=f({__name:"ExternalServiceDetails",props:{mesh:{},service:{}},setup(m){const{t:a}=y(),s=m;return(g,v)=>{const p=d("DataSource");return t(),u("div",null,[r(p,{src:`/meshes/${s.mesh}/external-services/for/${s.service}`},{default:e(({data:o,error:_})=>[_?(t(),i($,{key:0,error:_},null,8,["error"])):o===void 0?(t(),i(B,{key:1})):o===null?(t(),i(I,{key:2,"data-testid":"no-matching-external-service"},{title:e(()=>[k("p",null,c(l(a)("services.detail.no_matching_external_service",{name:s.service})),1)]),_:1})):(t(),u("div",b,[r(h,null,{title:e(()=>[n(c(l(a)("http.api.property.address")),1)]),body:e(()=>[r(x,{text:o.networking.address},null,8,["text"])]),_:2},1024),n(),o.tags!==null?(t(),i(h,{key:0},{title:e(()=>[n(c(l(a)("http.api.property.tags")),1)]),body:e(()=>[r(D,{tags:o.tags},null,8,["tags"])]),_:2},1024)):w("",!0)]))]),_:1},8,["src"])])}}}),E={class:"columns"},N=f({__name:"ServiceInsightDetails",props:{serviceInsight:{}},setup(m){const{t:a}=y(),s=m;return(g,v)=>{var p,o;return t(),u("div",E,[r(h,null,{title:e(()=>[n(c(l(a)("http.api.property.status")),1)]),body:e(()=>[r(T,{status:s.serviceInsight.status},null,8,["status"])]),_:1}),n(),r(h,null,{title:e(()=>[n(c(l(a)("http.api.property.address")),1)]),body:e(()=>[s.serviceInsight.addressPort?(t(),i(x,{key:0,text:s.serviceInsight.addressPort},null,8,["text"])):(t(),u(V,{key:1},[n(c(l(a)("common.detail.none")),1)],64))]),_:1}),n(),r(C,{online:((p=s.serviceInsight.dataplanes)==null?void 0:p.online)??0,total:((o=s.serviceInsight.dataplanes)==null?void 0:o.total)??0},{title:e(()=>[n(c(l(a)("http.api.property.dataPlaneProxies")),1)]),_:1},8,["online","total"])])}}}),P={class:"stack"},G=f({__name:"ServiceDetailView",props:{data:{}},setup(m){const a=m;return(s,g)=>{const v=d("KCard"),p=d("AppView"),o=d("RouteView");return t(),i(o,{name:"service-detail-view",params:{mesh:"",service:""}},{default:e(({route:_})=>[r(p,null,{default:e(()=>[k("div",P,[r(v,null,{default:e(()=>[a.data.serviceType==="external"?(t(),i(S,{key:0,mesh:_.params.mesh,service:_.params.service},null,8,["mesh","service"])):(t(),i(N,{key:1,"service-insight":s.data},null,8,["service-insight"]))]),_:2},1024)])]),_:2},1024)]),_:1})}}});export{G as default};
