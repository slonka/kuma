import{d as _,u as d,r as i,v as u,o,j as c,b as l,g as k}from"./index-0b474b9f.js";import{_ as w}from"./ZoneDetails.vue_vue_type_script_setup_true_lang-3a595c00.js";import{_ as z}from"./EmptyBlock.vue_vue_type_script_setup_true_lang-0fb89700.js";import{E as h}from"./ErrorBlock-b6b5f4b2.js";import{_ as y}from"./LoadingBlock.vue_vue_type_script_setup_true_lang-4f0d596b.js";import{u as g}from"./store-2a6e4467.js";import{u as B}from"./index-a5f95863.js";import"./kongponents.es-dca28cba.js";import"./AccordionList-05424288.js";import"./_plugin-vue_export-helper-c27b6911.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-f1268e20.js";import"./DefinitionListItem-27d09fac.js";import"./SubscriptionHeader.vue_vue_type_script_setup_true_lang-58a8d14c.js";import"./TabsWidget-da6e4bda.js";import"./QueryParameter-70743f73.js";import"./TextWithCopyButton-5142404b.js";import"./WarningsWidget.vue_vue_type_script_setup_true_lang-a80d5a0a.js";const E={class:"zone-details"},$={key:3,class:"kcard-border"},G=_({__name:"ZoneDetailView",setup(b){const p=B(),e=d(),f=g(),a=i(null),n=i(!0),r=i(null);u(()=>e.params.mesh,function(){e.name==="zone-detail-view"&&s()}),u(()=>e.params.name,function(){e.name==="zone-detail-view"&&s()}),v();function v(){f.dispatch("updatePageTitle",e.params.zone),s()}async function s(){n.value=!0,r.value=null;const m=e.params.zone;try{a.value=await p.getZoneOverview({name:m})}catch(t){a.value=null,t instanceof Error?r.value=t:console.error(t)}finally{n.value=!1}}return(m,t)=>(o(),c("div",E,[n.value?(o(),l(y,{key:0})):r.value!==null?(o(),l(h,{key:1,error:r.value},null,8,["error"])):a.value===null?(o(),l(z,{key:2})):(o(),c("div",$,[k(w,{"zone-overview":a.value},null,8,["zone-overview"])]))]))}});export{G as default};
