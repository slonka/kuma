import{m as y,$ as b,_ as f}from"./kongponents.es-3df60cd6.js";import{A as d}from"./kumaApi-db784568.js";import{d as v,c as h,h as t,e as i,a0 as E,u as r,w as n,f as o,a as w,b as u,o as s,g as a,t as c,F as B,v as x,p as S,m as C}from"./runtime-dom.esm-bundler-91b41870.js";import{_ as I}from"./_plugin-vue_export-helper-c27b6911.js";const k=e=>(S("data-v-9de3b600"),e=e(),C(),e),N={class:"error-block"},V=k(()=>a("p",null,"An error has occurred while trying to load this data.",-1)),A={class:"error-block-details"},D=k(()=>a("summary",null,"Details",-1)),F={key:0},$={key:0,class:"badge-list"},j=v({__name:"ErrorBlock",props:{error:{type:[Error,d],required:!1,default:null}},setup(e){const l=e,_=h(()=>l.error instanceof Error),m=h(()=>l.error instanceof d?l.error.causes:[]);return(q,z)=>(s(),t("div",N,[i(r(b),{"cta-is-hidden":""},E({title:n(()=>[i(r(y),{class:"mb-3",icon:"warning",color:"var(--black-75)","secondary-color":"var(--yellow-300)",size:"42"}),o(),V]),_:2},[r(_)||r(m).length>0?{name:"message",fn:n(()=>[a("details",A,[D,o(),r(_)?(s(),t("p",F,c(e.error.message),1)):u("",!0),o(),a("ul",null,[(s(!0),t(B,null,x(r(m),(p,g)=>(s(),t("li",{key:g},[a("b",null,[a("code",null,c(p.field),1)]),o(": "+c(p.message),1)]))),128))])])]),key:"0"}:void 0]),1024),o(),e.error instanceof r(d)?(s(),t("div",$,[e.error.code?(s(),w(r(f),{key:0,appearance:"warning"},{default:n(()=>[o(c(e.error.code),1)]),_:1})):u("",!0),o(),i(r(f),{appearance:"warning"},{default:n(()=>[o(c(e.error.statusCode),1)]),_:1})])):u("",!0)]))}});const H=I(j,[["__scopeId","data-v-9de3b600"]]);export{H as E};
