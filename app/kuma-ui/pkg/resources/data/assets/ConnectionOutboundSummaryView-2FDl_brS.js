import{_ as C}from"./NavTabs.vue_vue_type_script_setup_true_lang-B8bQ_WBr.js";import{d as R,a,o as r,b as m,w as e,t as p,m as D,f as c,e as n,H as h,X as k,D as x}from"./index-1j9z4Egf.js";const N=R({__name:"ConnectionOutboundSummaryView",props:{data:{}},setup(l){const u=l;return(y,b)=>{const _=a("RouterLink"),d=a("DataCollection"),f=a("RouterView"),v=a("AppView"),w=a("RouteView");return r(),m(w,{name:"connection-outbound-summary-view",params:{service:""}},{default:e(({route:o,t:V})=>[n(v,null,{title:e(()=>[D("h2",null,p(o.params.service),1)]),default:e(()=>{var i;return[c(),n(C,{"active-route-name":(i=o.active)==null?void 0:i.name},k({_:2},[h(o.children,t=>({name:`${t.name}`,fn:e(()=>[n(_,{to:{name:t.name}},{default:e(()=>[c(p(V(`connections.routes.item.navigation.${t.name.split("-")[3]}`)),1)]),_:2},1032,["to"])])}))]),1032,["active-route-name"]),c(),n(f,null,{default:e(({Component:t})=>[n(d,{items:u.data,predicate:s=>s.name===o.params.service,find:!0},{default:e(({items:s})=>[(r(),m(x(t),{data:s[0]},null,8,["data"]))]),_:2},1032,["items","predicate"])]),_:2},1024)]}),_:2},1024)]),_:1})}}});export{N as default};
