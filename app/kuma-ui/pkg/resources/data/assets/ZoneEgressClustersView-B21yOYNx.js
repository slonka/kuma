import{_ as p}from"./EnvoyData.vue_vue_type_script_setup_true_lang-__EFnqg3.js";import{d,i as s,o as m,a as u,w as a,j as t,g,k as _}from"./index-Bfp4yKlY.js";import"./kong-icons.es350-CHTYrByy.js";import"./CodeBlock-CCCo0pfs.js";const V=d({__name:"ZoneEgressClustersView",setup(f){return(h,C)=>{const n=s("RouteTitle"),r=s("KCard"),c=s("AppView"),i=s("RouteView");return m(),u(i,{name:"zone-egress-clusters-view",params:{zoneEgress:"",codeSearch:"",codeFilter:!1,codeRegExp:!1}},{default:a(({route:e,t:l})=>[t(c,null,{title:a(()=>[g("h2",null,[t(n,{title:l("zone-egresses.routes.item.navigation.zone-egress-clusters-view")},null,8,["title"])])]),default:a(()=>[_(),t(r,null,{default:a(()=>[t(p,{resource:"Zone",src:`/zone-egresses/${e.params.zoneEgress}/data-path/clusters`,query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter,"is-reg-exp-mode":e.params.codeRegExp,onQueryChange:o=>e.update({codeSearch:o}),onFilterModeChange:o=>e.update({codeFilter:o}),onRegExpModeChange:o=>e.update({codeRegExp:o})},null,8,["src","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{V as default};