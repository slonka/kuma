import{E as d}from"./EnvoyData-d1af290d.js";import{g as l}from"./dataplane-dcd0858b.js";import{d as g,a as e,o as _,b as f,w as o,e as t,p as h,f as w,q as C}from"./index-784d2bbf.js";import"./index-9dd3e7d3.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-226d1ddf.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-f6a2a033.js";import"./ErrorBlock-d38c2168.js";import"./TextWithCopyButton-7ef74197.js";import"./CopyButton-9c00109a.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-9960c4c9.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-8f5d9bcc.js";const b=g({__name:"ZoneEgressXdsConfigView",props:{data:{}},setup(a){const n=a;return(V,x)=>{const r=e("RouteTitle"),i=e("KCard"),p=e("AppView"),c=e("RouteView");return _(),f(c,{name:"zone-egress-xds-config-view",params:{zoneEgress:"",codeSearch:""}},{default:o(({route:s,t:m})=>[t(p,null,{title:o(()=>[h("h2",null,[t(r,{title:m("zone-egresses.routes.item.navigation.zone-egress-xds-config-view")},null,8,["title"])])]),default:o(()=>[w(),t(i,null,{body:o(()=>[t(d,{status:C(l)(n.data.zoneEgressInsight),resource:"Zone",src:`/zone-egresses/${s.params.zoneEgress}/data-path/xds`,query:s.params.codeSearch,onQueryChange:u=>s.update({codeSearch:u})},null,8,["status","src","query","onQueryChange"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{b as default};
