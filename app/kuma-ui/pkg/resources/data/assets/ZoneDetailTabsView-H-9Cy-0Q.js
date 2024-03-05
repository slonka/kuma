import{d as E,l as R,Q as B,aE as I,C as A,a as s,o as c,c as N,e as n,w as e,aF as T,q as a,t as d,f as r,b as m,m as g,p as S,E as Z,A as $,W as L,H as M,X as O,D as K,aB as V}from"./index-1j9z4Egf.js";import{_ as P}from"./DeleteResourceModal.vue_vue_type_script_setup_true_lang-uT_HNzVW.js";import{_ as j}from"./NavTabs.vue_vue_type_script_setup_true_lang-B8bQ_WBr.js";const q=E({__name:"ZoneActionMenu",props:{zoneOverview:{type:Object,required:!0},kpopAttributes:{type:Object,default:()=>({placement:"bottomEnd"})}},setup(C){const{t}=R(),p=B(),f=I(),o=C,i=A(!1);function _(){i.value=!i.value}async function v(){await p.deleteZone({name:o.zoneOverview.name})}function w(){f.push({name:"zone-cp-list-view"})}return(h,x)=>{const b=s("KDropdownItem"),z=s("KDropdown");return c(),N("div",null,[n(z,{"kpop-attributes":o.kpopAttributes,"trigger-text":a(t)("zones.action_menu.toggle_button"),"show-caret":"",width:"280"},{items:e(()=>[n(b,{danger:"","data-testid":"delete-button",onClick:T(_,["prevent"])},{default:e(()=>[r(d(a(t)("zones.action_menu.delete_button")),1)]),_:1})]),_:1},8,["kpop-attributes","trigger-text"]),r(),i.value?(c(),m(P,{key:0,"confirmation-text":o.zoneOverview.name,"delete-function":v,"is-visible":"","action-button-text":a(t)("common.delete_modal.proceed_button"),title:a(t)("common.delete_modal.title",{type:"Zone"}),"data-testid":"delete-zone-modal",onCancel:_,onDelete:w},{default:e(()=>[g("p",null,d(a(t)("common.delete_modal.text1",{type:"Zone",name:o.zoneOverview.name})),1),r(),g("p",null,d(a(t)("common.delete_modal.text2")),1)]),_:1},8,["confirmation-text","action-button-text","title"])):S("",!0)])}}}),F=E({__name:"ZoneDetailTabsView",setup(C){const{t}=R(),p=A([]),f=o=>{const i=[];o.zoneInsight.store==="memory"&&i.push({kind:"ZONE_STORE_TYPE_MEMORY",payload:{}}),V(o.zoneInsight,"version.kumaCp.kumaCpGlobalCompatible","true")||i.push({kind:"INCOMPATIBLE_ZONE_AND_GLOBAL_CPS_VERSIONS",payload:{zoneCpVersion:V(o.zoneInsight,"version.kumaCp.version",t("common.collection.none"))}}),p.value=i};return(o,i)=>{const _=s("RouteTitle"),v=s("RouterLink"),w=s("RouterView"),h=s("AppView"),x=s("DataSource"),b=s("RouteView");return c(),m(b,{name:"zone-cp-detail-tabs-view",params:{zone:""}},{default:e(({can:z,route:u})=>[n(x,{src:`/zone-cps/${u.params.zone}`,onChange:f},{default:e(({data:k,error:y})=>[y!==void 0?(c(),m(Z,{key:0,error:y},null,8,["error"])):k===void 0?(c(),m($,{key:1})):(c(),m(h,{key:2,breadcrumbs:[{to:{name:"zone-cp-list-view"},text:a(t)("zone-cps.routes.item.breadcrumbs")}]},O({title:e(()=>[g("h1",null,[n(L,{text:u.params.zone},{default:e(()=>[n(_,{title:a(t)("zone-cps.routes.item.title",{name:u.params.zone})},null,8,["title"])]),_:2},1032,["text"])])]),default:e(()=>{var D;return[r(),r(),n(j,{"active-route-name":(D=u.active)==null?void 0:D.name},O({_:2},[M(u.children,({name:l})=>({name:`${l}`,fn:e(()=>[n(v,{to:{name:l},"data-testid":`${l}-tab`},{default:e(()=>[r(d(a(t)(`zone-cps.routes.item.navigation.${l}`)),1)]),_:2},1032,["to","data-testid"])])}))]),1032,["active-route-name"]),r(),n(w,null,{default:e(l=>[(c(),m(K(l.Component),{data:k,notifications:p.value},null,8,["data","notifications"]))]),_:2},1024)]}),_:2},[z("create zones")?{name:"actions",fn:e(()=>[n(q,{"zone-overview":k},null,8,["zone-overview"])]),key:"0"}:void 0]),1032,["breadcrumbs"]))]),_:2},1032,["src"])]),_:1})}}});export{F as default};
