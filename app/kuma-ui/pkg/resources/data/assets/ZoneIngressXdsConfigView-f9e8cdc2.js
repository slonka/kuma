import{E as m}from"./EnvoyData-e239327f.js";import{d as l,a as e,o as d,b as u,w as o,e as n,p as _,f as g}from"./index-2abc6d0d.js";import"./index-52545d1d.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-f4baddc1.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-3a8027ab.js";import"./ErrorBlock-ec3dc6ed.js";import"./TextWithCopyButton-6f5fcbbb.js";import"./CopyButton-31edefd2.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-1e5409a6.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-851e7c2d.js";const S=l({__name:"ZoneIngressXdsConfigView",setup(f){return(h,w)=>{const s=e("RouteTitle"),r=e("KCard"),a=e("AppView"),i=e("RouteView");return d(),u(i,{name:"zone-ingress-xds-config-view",params:{zoneIngress:"",codeSearch:""}},{default:o(({route:t,t:p})=>[n(a,null,{title:o(()=>[_("h2",null,[n(s,{title:p("zone-ingresses.routes.item.navigation.zone-ingress-xds-config-view")},null,8,["title"])])]),default:o(()=>[g(),n(r,null,{body:o(()=>[n(m,{resource:"Zone",src:`/zone-ingresses/${t.params.zoneIngress}/data-path/xds`,query:t.params.codeSearch,onQueryChange:c=>t.update({codeSearch:c})},null,8,["src","query","onQueryChange"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{S as default};