import{m as x,a as w}from"./kongponents.es-3df60cd6.js";import{h as f,c as I,r as g}from"./helpers-32595d9f.js";import{d as v,c as C,o as s,h as o,f as n,g as t,t as a,b as l,u as i,F as y,v as b,a as B,w as k,e as V,p as $,m as L}from"./runtime-dom.esm-bundler-91b41870.js";import{_ as N}from"./_plugin-vue_export-helper-c27b6911.js";const d=e=>($("data-v-0e767065"),e=e(),L(),e),P={key:0},j=d(()=>t("h5",{class:"overview-tertiary-title"},`
        General Information:
      `,-1)),q={key:0},E=d(()=>t("strong",null,"Global Instance ID:",-1)),F={class:"mono"},G={key:1},O=d(()=>t("strong",null,"Control Plane Instance ID:",-1)),R={class:"mono"},H={key:2},M=d(()=>t("strong",null,"Last Connected:",-1)),U={key:3},W=d(()=>t("strong",null,"Last Disconnected:",-1)),z={key:1},A={class:"overview-stat-grid"},J={class:"overview-tertiary-title"},K={class:"mono"},Q=v({__name:"SubscriptionDetails",props:{details:{type:Object,required:!0},isDiscoverySubscription:{type:Boolean,default:!1}},setup(e){const r=e,m=C(()=>{var c;if(r.isDiscoverySubscription){const{lastUpdateTime:S,total:h,...u}=r.details.status;return u}return(c=r.details.status)==null?void 0:c.stat});function _(c){return c?parseInt(c,10).toLocaleString("en").toString():"0"}function D(c){return c==="--"?"error calculating":c}return(c,S)=>(s(),o("div",null,[e.details.globalInstanceId||e.details.connectTime||e.details.disconnectTime?(s(),o("div",P,[j,n(),t("ul",null,[e.details.globalInstanceId?(s(),o("li",q,[E,n(` 
          `),t("span",F,a(e.details.globalInstanceId),1)])):l("",!0),n(),e.details.controlPlaneInstanceId?(s(),o("li",G,[O,n(` 
          `),t("span",R,a(e.details.controlPlaneInstanceId),1)])):l("",!0),n(),e.details.connectTime?(s(),o("li",H,[M,n(` 
          `+a(i(f)(e.details.connectTime)),1)])):l("",!0),n(),e.details.disconnectTime?(s(),o("li",U,[W,n(` 
          `+a(i(f)(e.details.disconnectTime)),1)])):l("",!0)])])):l("",!0),n(),i(m)?(s(),o("div",z,[t("ul",A,[(s(!0),o(y,null,b(i(m),(h,u)=>(s(),o("li",{key:u},[t("h6",J,a(i(I)(u))+`:
          `,1),n(),t("ul",null,[(s(!0),o(y,null,b(h,(T,p)=>(s(),o("li",{key:p},[t("strong",null,a(i(I)(p))+":",1),n(` 
              `),t("span",K,a(D(_(T))),1)]))),128))])]))),128))])])):(s(),B(i(w),{key:2,appearance:"info",class:"mt-4"},{alertIcon:k(()=>[V(i(x),{icon:"portal"})]),alertMessage:k(()=>[n(`
        There are no subscription statistics for `),t("strong",null,a(e.details.id),1)]),_:1}))]))}});const oe=N(Q,[["__scopeId","data-v-0e767065"]]),X={class:"text-lg font-medium"},Y={class:"color-green-500"},Z={key:0,class:"ml-4 color-red-600"},ae=v({__name:"SubscriptionHeader",props:{details:{type:Object,required:!0}},setup(e){const r=e;return(m,_)=>(s(),o("h4",X,[t("span",Y,`
      Connect time: `+a(i(g)(r.details.connectTime)),1),n(),r.details.disconnectTime?(s(),o("span",Z,`
      Disconnect time: `+a(i(g)(r.details.disconnectTime)),1)):l("",!0)]))}});export{oe as S,ae as _};
