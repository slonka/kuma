import{d as g,R as k,r as e,o as s,i as n,w as o,j as r,p as w,n as z,E as V,x as h,k as v}from"./index-adcc6fc8.js";import{_ as C}from"./ResourceCodeBlock.vue_vue_type_style_index_0_lang-1849c988.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-4a09da2d.js";import"./toYaml-4e00099e.js";const K=g({__name:"ConfigView",setup(x){const i=k();return(R,y)=>{const u=e("RouteTitle"),_=e("DataSource"),l=e("KCard"),m=e("AppView"),p=e("RouteView");return s(),n(p,{name:"zone-ingress-config-view",params:{zoneIngress:""}},{default:o(({route:t,t:d})=>[r(m,null,{title:o(()=>[w("h2",null,[r(u,{title:d("zone-ingresses.routes.item.navigation.zone-ingress-config-view"),render:!0},null,8,["title"])])]),default:o(()=>[z(),r(l,{class:"mt-4"},{body:o(()=>[r(_,{src:`/zone-ingresses/${t.params.zoneIngress}`},{default:o(({data:a,error:c})=>[c!==void 0?(s(),n(V,{key:0,error:c},null,8,["error"])):a===void 0?(s(),n(h,{key:1})):(s(),n(C,{key:2,id:"code-block-zone-ingress",resource:a,"resource-fetcher":f=>v(i).getZoneIngress({name:t.params.zoneIngress},f),"is-searchable":""},null,8,["resource","resource-fetcher"]))]),_:2},1032,["src"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{K as default};
