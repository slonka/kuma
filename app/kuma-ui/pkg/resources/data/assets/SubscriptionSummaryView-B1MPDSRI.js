import{d as S,r as t,o as r,m as l,w as e,b as n,k as v,t as p,e as c,T as y,M as A,E as D,a as I}from"./index-BZS2OlAU.js";const X=S({__name:"SubscriptionSummaryView",props:{data:{}},setup(d){const m=d;return(u,R)=>{const _=t("XAction"),f=t("XTabs"),b=t("RouterView"),w=t("AppView"),V=t("DataCollection"),C=t("RouteView");return r(),l(C,{name:"subscription-summary-view",params:{subscription:""}},{default:e(({route:s,t:h})=>[n(V,{items:m.data,predicate:o=>o.id===s.params.subscription},{item:e(({item:o})=>[n(w,null,{title:e(()=>[v("h2",null,p(o.zoneInstanceId??o.globalInstanceId),1)]),default:e(()=>{var i;return[c(),n(f,{selected:(i=s.child())==null?void 0:i.name},y({_:2},[A(s.children,({name:a})=>({name:`${a}-tab`,fn:e(()=>[n(_,{to:{name:a}},{default:e(()=>[c(p(h(`subscriptions.routes.item.navigation.${a}`)),1)]),_:2},1032,["to"])])}))]),1032,["selected"]),c(),n(b,null,{default:e(({Component:a})=>[(r(),l(D(a),{data:o},{default:e(()=>[I(u.$slots,"default")]),_:2},1032,["data"]))]),_:2},1024)]}),_:2},1024)]),_:2},1032,["items","predicate"])]),_:3})}}});export{X as default};
