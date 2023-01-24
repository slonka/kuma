import{s as w}from"./helpers-32595d9f.js";import{S as B}from"./StatusBadge-81464ebd.js";import{T as C}from"./TagList-91d1133a.js";import{_ as D}from"./YamlView.vue_vue_type_script_setup_true_lang-d444e9d2.js";import{d as I,c as i,l as N,h as l,g as s,f as t,e as x,w as O,u as n,a as S,b as u,t as d,F as g,p as V,m as j,o as a}from"./runtime-dom.esm-bundler-91b41870.js";import{_ as P}from"./_plugin-vue_export-helper-c27b6911.js";const _=o=>(V("data-v-aeb75a49"),o=o(),j(),o),$={class:"entity-summary entity-section-list"},q={class:"block-list"},E={class:"entity-title"},F={class:"definition"},L=_(()=>s("span",null,"Mesh:",-1)),A={class:"definition"},M=_(()=>s("span",null,"Address:",-1)),R={key:0,class:"definition"},z=_(()=>s("span",null,"TLS:",-1)),G={key:1,class:"definition"},H=_(()=>s("span",null,"Data plane proxies:",-1)),J={key:0},K=_(()=>s("h2",null,"Tags",-1)),Q={class:"config-section"},U=I({__name:"ServiceSummary",props:{service:{type:Object,required:!0},externalService:{type:Object,required:!1,default:null}},setup(o){const e=o,k=i(()=>({name:"service-detail-view",params:{service:e.service.name,mesh:e.service.mesh}})),v=i(()=>e.service.serviceType==="external"&&e.externalService!==null?e.externalService.networking.address:e.service.addressPort??null),m=i(()=>{var r;return e.service.serviceType==="external"&&e.externalService!==null?(r=e.externalService.networking.tls)!=null&&r.enabled?"Enabled":"Disabled":null}),h=i(()=>{var r,c;if(e.service.serviceType==="external")return null;{const p=((r=e.service.dataplanes)==null?void 0:r.online)??0,T=((c=e.service.dataplanes)==null?void 0:c.total)??0;return`${p} online / ${T} total`}}),f=i(()=>e.service.serviceType==="external"?null:e.service.status??null),y=i(()=>e.service.serviceType==="external"&&e.externalService!==null?Object.entries(e.externalService.tags).map(([r,c])=>({label:r,value:c})):[]),b=i(()=>w(e.externalService??e.service));return(r,c)=>{const p=N("router-link");return a(),l("div",$,[s("section",null,[s("div",q,[s("div",null,[s("h1",E,[s("span",null,[t(`
              Service:

              `),x(p,{to:n(k)},{default:O(()=>[t(d(e.service.name),1)]),_:1},8,["to"])]),t(),n(f)?(a(),S(B,{key:0,status:n(f)},null,8,["status"])):u("",!0)]),t(),s("div",F,[L,t(),s("span",null,d(e.service.mesh),1)]),t(),s("div",A,[M,t(),s("span",null,[n(v)!==null?(a(),l(g,{key:0},[t(d(n(v)),1)],64)):(a(),l(g,{key:1},[t("—")],64))])]),t(),n(m)!==null?(a(),l("div",R,[z,t(),s("span",null,d(n(m)),1)])):u("",!0),t(),n(h)!==null?(a(),l("div",G,[H,t(),s("span",null,d(n(h)),1)])):u("",!0)]),t(),n(y).length>0?(a(),l("div",J,[K,t(),x(C,{tags:n(y)},null,8,["tags"])])):u("",!0)])]),t(),s("section",Q,[e.service.serviceType==="external"?(a(),S(D,{key:0,id:"code-block-service",content:n(b),"is-searchable":"","code-max-height":"250px"},null,8,["content"])):u("",!0)])])}}});const te=P(U,[["__scopeId","data-v-aeb75a49"]]);export{te as S};
