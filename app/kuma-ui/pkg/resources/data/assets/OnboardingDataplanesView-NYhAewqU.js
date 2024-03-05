import{L as x}from"./LoadingBox-AQ9YRhzj.js";import{O as V,a as D,b as B}from"./OnboardingPage-6lSa7jjo.js";import{d as O,C as m,M as S,a as l,o as s,b as r,w as t,e as n,f as d,E as A,A as C,F as N,t as c,m as u,H as T,c as p,_ as R}from"./index-1j9z4Egf.js";import{S as $}from"./StatusBadge-j_uZdXEO.js";const E={key:0,class:"status-loading-box mb-4"},F={key:1},K={class:"mb-4"},L=O({__name:"OnboardingDataplanesView",setup(P){const e=m(),_=m(0),v=S(()=>Array.isArray(e.value)?e.value.some(o=>o.status==="offline"):!0);function f(o){o&&(e.value=o.items,_.value++)}return(o,H)=>{const b=l("RouteTitle"),h=l("KTable"),w=l("DataSource"),k=l("AppView"),y=l("RouteView");return s(),r(y,{name:"onboarding-dataplanes-view"},{default:t(({t:i})=>[n(b,{title:i("onboarding.routes.dataplanes-overview.title"),render:!1},null,8,["title"]),d(),n(k,null,{default:t(()=>[n(w,{src:v.value?"/dataplanes/poll?page=1&size=10":"",onChange:f},{default:t(({error:g})=>[g!==void 0?(s(),r(A,{key:0,error:g},null,8,["error"])):e.value===void 0?(s(),r(C,{key:1})):(s(),r(V,{key:2},{header:t(()=>[(s(!0),p(N,null,T([v.value?"waiting":"success"],a=>(s(),r(D,{key:a,"data-testid":`state-${a}`},{title:t(()=>[d(c(i(`onboarding.routes.dataplanes-overview.header.${a}.title`)),1)]),description:t(()=>[u("p",null,c(i(`onboarding.routes.dataplanes-overview.header.${a}.description`)),1)]),_:2},1032,["data-testid"]))),128))]),content:t(()=>[e.value.length===0?(s(),p("div",E,[n(x)])):(s(),p("div",F,[u("p",K,[u("b",null,"Found "+c(e.value.length)+" DPPs:",1)]),d(),n(h,{class:"mb-4","data-testid":"dataplanes-table","fetcher-cache-key":String(_.value),fetcher:()=>{var a;return{data:e.value,total:(a=e.value)==null?void 0:a.length}},headers:[{label:"Mesh",key:"mesh"},{label:"Name",key:"name"},{label:"Status",key:"status"}],"disable-pagination":""},{status:t(({row:a})=>[n($,{status:a.status},null,8,["status"])]),_:1},8,["fetcher-cache-key","fetcher"])]))]),navigation:t(()=>[n(B,{"next-step":"onboarding-completed-view","previous-step":"onboarding-add-new-services-code-view","should-allow-next":e.value.length>0},null,8,["should-allow-next"])]),_:2},1024))]),_:2},1032,["src"])]),_:2},1024)]),_:1})}}}),q=R(L,[["__scopeId","data-v-f140547b"]]);export{q as default};
