import{d as C,k as x,V as D,a as c,o as f,b as w,w as o,e as r,m as p,f as d,t as R,l as S,C as N,x as g,y as I,a2 as k,_ as B}from"./index-4198b723.js";import{N as O}from"./NavTabs-f0ec179f.js";const P=t=>(g("data-v-03904fc2"),t=t(),I(),t),A={class:"summary-title-wrapper"},T=P(()=>p("img",{"aria-hidden":"true",src:k},null,-1)),$={class:"summary-title"},j=C({__name:"DataPlaneOutboundSummaryView",props:{data:{}},setup(t){var l;const{t:v}=x(),y=D(),V=t,h=(((l=y.getRoutes().find(e=>e.name==="data-plane-outbound-summary-view"))==null?void 0:l.children)??[]).map(e=>{var n,a;const i=typeof e.name>"u"?(n=e.children)==null?void 0:n[0]:e,s=i.name,u=((a=i.meta)==null?void 0:a.module)??"";return{title:v(`data-planes.routes.item.navigation.${s}`),routeName:s,module:u}});return(e,i)=>{const s=c("DataCollection"),u=c("RouterView"),_=c("AppView"),n=c("RouteView");return f(),w(n,{name:"data-plane-outbound-summary-view",params:{service:""}},{default:o(({route:a})=>[r(_,null,{title:o(()=>[p("div",A,[T,d(),p("h2",$,R(a.params.service),1)])]),default:o(()=>[d(),r(O,{tabs:S(h)},null,8,["tabs"]),d(),r(u,null,{default:o(b=>[r(s,{items:V.data,predicate:m=>m.name===a.params.service,find:!0},{default:o(({items:m})=>[(f(),w(N(b.Component),{data:m[0]},null,8,["data"]))]),_:2},1032,["items","predicate"])]),_:2},1024)]),_:2},1024)]),_:1})}}});const F=B(j,[["__scopeId","data-v-03904fc2"]]);export{F as default};
