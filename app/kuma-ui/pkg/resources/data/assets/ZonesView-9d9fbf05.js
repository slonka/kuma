import{u as se}from"./vue-router-573afc44.js";import{P as ne,_ as oe,D as q}from"./kongponents.es-3df60cd6.js";import{f as M,a as re,b as ie}from"./helpers-32595d9f.js";import{g as G,I as le}from"./dataplane-4aecf58f.js";import{k as c}from"./kumaApi-db784568.js";import{b as R}from"./constants-31fdaf55.js";import{Q as J}from"./QueryParameter-70743f73.js";import{u as ue}from"./store-866e3b85.js";import{a as ce,A as fe}from"./AccordionItem-5cf952b5.js";import{_ as me}from"./CodeBlock.vue_vue_type_style_index_0_lang-98064716.js";import{D as pe}from"./DataOverview-07eb4ec1.js";import{F as ge}from"./FrameSkeleton-256a9a83.js";import{_ as de}from"./LabelList.vue_vue_type_style_index_0_lang-4346c84c.js";import{_ as ve}from"./MultizoneInfo.vue_vue_type_script_setup_true_lang-f8d86b15.js";import{_ as he,S as ye}from"./SubscriptionHeader.vue_vue_type_script_setup_true_lang-bf23b0a3.js";import{T as _e}from"./TabsWidget-aec0fabd.js";import{_ as be}from"./WarningsWidget.vue_vue_type_script_setup_true_lang-213c05f4.js";import{d as we,r as n,y as ke,_ as Se,h as f,u as g,a as d,w as o,o as i,e as u,f as I,b as C,g as D,t as z,F as Y,v as H}from"./runtime-dom.esm-bundler-91b41870.js";import"./vuex.esm-bundler-df5bd11e.js";import"./ClientStorage-efe299d9.js";import"./_plugin-vue_export-helper-c27b6911.js";import"./_commonjsHelpers-87174ba5.js";import"./datadogLogEvents-4578cfa7.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-4047971f.js";import"./ErrorBlock-46fedade.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-d3176fee.js";import"./StatusBadge-81464ebd.js";import"./TagList-91d1133a.js";import"./index-f6b877ca.js";const Ee={class:"zones"},Ie={class:"entity-heading"},Ce={key:0},ze={key:1},Ae={key:2},na=we({__name:"ZonesView",props:{selectedZoneName:{type:String,required:!1,default:null},offset:{type:Number,required:!1,default:0}},setup(Q){const A=Q,U={title:"No Data",message:"There are no Zones present."},O=[{hash:"#overview",title:"Overview"},{hash:"#insights",title:"Zone Insights"},{hash:"#config",title:"Config"},{hash:"#warnings",title:"Warnings"}],N=se(),Z=ue(),v=n(!0),m=n(!1),_=n(null),b=n(!0),h=n(!1),w=n(!1),k=n(!1),S=n({headers:[{label:"Actions",key:"actions",hideLabel:!0},{label:"Status",key:"status"},{label:"Name",key:"name"},{label:"Zone CP Version",key:"zoneCpVersion"},{label:"Storage type",key:"storeType"},{label:"Ingress",key:"hasIngress"},{label:"Egress",key:"hasEgress"},{label:"Warnings",key:"warnings",hideLabel:!0}],data:[]}),p=n(null),V=n(null),E=n([]),L=n([]),T=n(null),B=n(A.offset),x=n(new Set),P=n(new Set);ke(()=>N.params.mesh,function(){N.name==="zones"&&(v.value=!0,m.value=!1,b.value=!0,h.value=!1,w.value=!1,k.value=!1,_.value=null,W(0))}),Se(function(){W(A.offset)});function W(a){Z.getters["config/getMulticlusterStatus"]&&F(a)}function j(){return E.value.length===0?O.filter(a=>a.hash!=="#warnings"):O}function K(a){let s="-",t="",e=!0;a.zoneInsight.subscriptions&&a.zoneInsight.subscriptions.length>0&&a.zoneInsight.subscriptions.forEach(l=>{if(l.version&&l.version.kumaCp){s=l.version.kumaCp.version;const{kumaCpGlobalCompatible:y=!0}=l.version.kumaCp;e=y,l.config&&(t=JSON.parse(l.config).store.type)}});const r=G(a.zoneInsight);return{...a,status:r,zoneCpVersion:s,storeType:t,hasIngress:x.value.has(a.name)?"Yes":"No",hasEgress:P.value.has(a.name)?"Yes":"No",withWarnings:!e}}function X(a){const s=new Set;a.forEach(({zoneIngress:{zone:t}})=>{s.add(t)}),x.value=s}function ee(a){const s=new Set;a.forEach(({zoneEgress:{zone:t}})=>{s.add(t)}),P.value=s}async function F(a){B.value=a,J.set("offset",a>0?a:null),v.value=!0,m.value=!1;const s=N.query.ns||null,t=R;try{const[{data:e,next:r},{items:l},{items:y}]=await Promise.all([ae(s,t,a),M(c.getAllZoneIngressOverviews.bind(c)),M(c.getAllZoneEgressOverviews.bind(c))]);V.value=r,e.length?(X(l),ee(y),S.value.data=e.map(K),k.value=!1,m.value=!1,await $({name:A.selectedZoneName??e[0].name})):(S.value.data=[],k.value=!0,m.value=!0,h.value=!0)}catch(e){e instanceof Error?_.value=e:console.error(e),m.value=!0}finally{v.value=!1}}async function $({name:a}){var s;w.value=!1,b.value=!0,h.value=!1,E.value=[];try{const t=await c.getZoneOverview({name:a}),e=((s=t.zoneInsight)==null?void 0:s.subscriptions)??[],r=G(t.zoneInsight);if(p.value={...re(t,["type","name"]),status:r,"Authentication Type":ie(t)},J.set("zone",a),L.value=Array.from(e).reverse(),e.length>0){const l=e[e.length-1],y=l.version.kumaCp.version||"-",{kumaCpGlobalCompatible:te=!0}=l.version.kumaCp;te||E.value.push({kind:le,payload:{zoneCpVersion:y,globalCpVersion:Z.getters["config/getVersion"]}}),e[e.length-1].config&&(T.value=JSON.stringify(JSON.parse(e[e.length-1].config),null,2))}}catch(t){console.error(t),p.value=null,w.value=!0,h.value=!0}finally{b.value=!1}}async function ae(a,s,t){if(a)return{data:[await c.getZoneOverview({name:a},{size:s,offset:t})],next:null};{const{items:e,next:r}=await c.getAllZoneOverviews({size:s,offset:t});return{data:e??[],next:r}}}return(a,s)=>(i(),f("div",Ee,[g(Z).getters["config/getMulticlusterStatus"]===!1?(i(),d(ve,{key:0})):(i(),d(ge,{key:1},{default:o(()=>{var t;return[u(pe,{"selected-entity-name":(t=p.value)==null?void 0:t.name,"page-size":g(R),"is-loading":v.value,error:_.value,"empty-state":U,"table-data":S.value,"table-data-is-empty":k.value,"show-warnings":S.value.data.some(e=>e.withWarnings),next:V.value,"page-offset":B.value,onTableAction:$,onLoadData:F},{additionalControls:o(()=>[a.$route.query.ns?(i(),d(g(ne),{key:0,class:"back-button",appearance:"primary",icon:"arrowLeft",to:{name:"zones"}},{default:o(()=>[I(`
            View all
          `)]),_:1})):C("",!0)]),_:1},8,["selected-entity-name","page-size","is-loading","error","table-data","table-data-is-empty","show-warnings","next","page-offset"]),I(),m.value===!1&&p.value!==null?(i(),d(_e,{key:0,"has-error":_.value!==null,"is-loading":v.value,tabs:j()},{tabHeader:o(()=>[D("h1",Ie,`
            Zone: `+z(p.value.name),1)]),overview:o(()=>[u(de,{"has-error":w.value,"is-loading":b.value,"is-empty":h.value},{default:o(()=>[D("div",null,[D("ul",null,[(i(!0),f(Y,null,H(p.value,(e,r)=>(i(),f("li",{key:r},[e?(i(),f("h4",Ce,z(r),1)):C("",!0),I(),r==="status"?(i(),f("p",ze,[u(g(oe),{appearance:e==="Offline"?"danger":"success"},{default:o(()=>[I(z(e),1)]),_:2},1032,["appearance"])])):(i(),f("p",Ae,z(e),1))]))),128))])])]),_:1},8,["has-error","is-loading","is-empty"])]),insights:o(()=>[u(g(q),{"border-variant":"noBorder"},{body:o(()=>[u(ce,{"initially-open":0},{default:o(()=>[(i(!0),f(Y,null,H(L.value,(e,r)=>(i(),d(fe,{key:r},{"accordion-header":o(()=>[u(he,{details:e},null,8,["details"])]),"accordion-content":o(()=>[u(ye,{details:e},null,8,["details"])]),_:2},1024))),128))]),_:1})]),_:1})]),config:o(()=>[T.value?(i(),d(g(q),{key:0,"border-variant":"noBorder"},{body:o(()=>[u(me,{id:"code-block-zone-config",language:"json",code:T.value,"is-searchable":"","query-key":"zone-config"},null,8,["code"])]),_:1})):C("",!0)]),warnings:o(()=>[u(be,{warnings:E.value},null,8,["warnings"])]),_:1},8,["has-error","is-loading","tabs"])):C("",!0)]}),_:1}))]))}});export{na as default};
