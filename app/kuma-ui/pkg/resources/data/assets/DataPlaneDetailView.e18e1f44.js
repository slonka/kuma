import{d as G,o as e,c as b,w as l,k as a,n as D,cm as j,I as Z,l as t,F as _,t as o,a as m,b as P,y as C,u as v,K as le,M as oe,T as J,cl as W,j as X,C as x,A as Q,B as ee,r as L,g as K,i as E,q as z,f as T,cA as ie,cz as H,cC as re,x as ce,cK as ue,z as de,O as pe,cL as me,ck as he,cH as _e,cI as ye,cM as ve,cJ as fe,cv as ge,p as Pe}from"./index.ea0d4a24.js";import{_ as ae}from"./EmptyBlock.vue_vue_type_script_setup_true_lang.b7d1d2f8.js";import{E as te}from"./ErrorBlock.6da0472d.js";import{_ as se}from"./LoadingBlock.vue_vue_type_script_setup_true_lang.c9f72cfb.js";import{_ as we}from"./EntityURLControl.vue_vue_type_script_setup_true_lang.3f293302.js";import{S as ke,E as V}from"./EnvoyData.af20fc05.js";import{_ as Y}from"./LabelList.vue_vue_type_style_index_0_lang.cd975a49.js";import{a as De,S as be}from"./SubscriptionHeader.f45b5cb6.js";import{T as $e}from"./TabsWidget.2b9f88d3.js";import{W as Le}from"./WarningsWidget.5aa49a70.js";import{Y as Ie}from"./YamlView.a58d7ecd.js";import"./CodeBlock.vue_vue_type_style_index_0_lang.f474aa6d.js";import"./_commonjsHelpers.f037b798.js";import"./index.58caa11d.js";const Se={class:"dataplane-policy-header"},Te={class:"dataplane-policy-subtitle"},Ce={key:0,class:"badge-list"},Me={class:"policy-wrapper"},Oe={class:"policy-type"},Ne={key:0},Ge=G({__name:"SidecarDataplanePolicyList",props:{sidecarDataplanePolicies:{type:Array,required:!0}},setup(c){const i=c,I={inbound:"Policies applied on incoming connection on address",outbound:"Policies applied on outgoing connection to the address",service:"Policies applied on outgoing connections to service",dataplane:"Policies applied on all incoming and outgoing connections to the selected data plane proxy"};return(w,f)=>{const h=X("router-link");return e(),b(W,{"initially-open":[],"multiple-open":""},{default:l(()=>[(e(!0),a(_,null,D(i.sidecarDataplanePolicies,(s,p)=>(e(),b(j,{key:p},Z({"accordion-header":l(()=>[t("div",Se,[t("div",null,[t("p",null,[s.type==="dataplane"?(e(),a(_,{key:0},[P(" Dataplane ")],64)):(e(),a(_,{key:1},[P(o(s.service),1)],64))]),t("p",Te,[s.type==="inbound"||s.type==="outbound"?(e(),a(_,{key:0},[P(o(s.type)+" "+o(s.name),1)],64)):s.type==="service"||s.type==="dataplane"?(e(),a(_,{key:1},[P(o(s.type),1)],64)):C("",!0),m(v(oe),{width:"300",placement:"right",trigger:"hover"},{content:l(()=>[P(o(I[s.type]),1)]),default:l(()=>[m(v(le),{icon:"help",size:"16",class:"ml-1"})]),_:2},1024)])]),s.matchedPolicies.length>0?(e(),a("div",Ce,[(e(!0),a(_,null,D(s.matchedPolicies,(g,y)=>(e(),b(v(J),{key:`${p}-${y}`},{default:l(()=>[P(o(g.name),1)]),_:2},1024))),128))])):C("",!0)])]),_:2},[s.matchedPolicies.length>0?{name:"accordion-content",fn:l(()=>[t("div",Me,[(e(!0),a(_,null,D(s.matchedPolicies,(g,y)=>(e(),a("div",{key:`${p}-${y}`},[t("p",Oe,o(g.pluralName),1),g.policies.length>0?(e(),a("ul",Ne,[(e(!0),a(_,null,D(g.policies,(k,O)=>(e(),a("li",{key:`${p}-${y}-${O}`,"data-testid":"policy-name"},[m(h,{to:k.route},{default:l(()=>[P(o(k.name),1)]),_:2},1032,["to"])]))),128))])):C("",!0)]))),128))])]),key:"0"}:void 0]),1024))),128))]),_:1})}}});const Ae=x(Ge,[["__scopeId","data-v-933a6e18"]]),N=c=>(Q("data-v-55dd6da4"),c=c(),ee(),c),Ee={class:"mesh-gateway-policy-list"},xe=N(()=>t("h3",null,"Gateway policies",-1)),Be={key:0,class:"policy-list"},qe=N(()=>t("h3",{class:"mt-6"}," Listeners ",-1)),Re=N(()=>t("b",null,"Host",-1)),Ve=N(()=>t("h4",{class:"mt-2"}," Routes ",-1)),Ke={class:"dataplane-policy-header"},je=N(()=>t("b",null,"Route",-1)),We=N(()=>t("b",null,"Service",-1)),ze={key:0,class:"badge-list"},Fe={class:"policy-list mt-1"},Ue=G({__name:"MeshGatewayDataplanePolicyList",props:{meshGatewayDataplane:{type:Object,required:!0},meshGatewayListenerEntries:{type:Array,required:!0},meshGatewayRoutePolicies:{type:Array,required:!0}},setup(c){const i=c;return(I,w)=>{const f=X("router-link");return e(),a("div",Ee,[xe,c.meshGatewayRoutePolicies.length>0?(e(),a("ul",Be,[(e(!0),a(_,null,D(c.meshGatewayRoutePolicies,(h,s)=>(e(),a("li",{key:s},[t("span",null,o(h.type),1),P(": "),m(f,{to:h.route},{default:l(()=>[P(o(h.name),1)]),_:2},1032,["to"])]))),128))])):C("",!0),qe,t("div",null,[(e(!0),a(_,null,D(i.meshGatewayListenerEntries,(h,s)=>(e(),a("div",{key:s},[t("div",null,[t("div",null,[Re,P(": "+o(h.hostName)+":"+o(h.port)+" ("+o(h.protocol)+") ",1)]),h.routeEntries.length>0?(e(),a(_,{key:0},[Ve,m(W,{"initially-open":[],"multiple-open":""},{default:l(()=>[(e(!0),a(_,null,D(h.routeEntries,(p,g)=>(e(),b(j,{key:g},Z({"accordion-header":l(()=>[t("div",Ke,[t("div",null,[t("div",null,[je,P(": "),m(f,{to:p.route},{default:l(()=>[P(o(p.routeName),1)]),_:2},1032,["to"])]),t("div",null,[We,P(": "+o(p.service),1)])]),p.policies.length>0?(e(),a("div",ze,[(e(!0),a(_,null,D(p.policies,(y,k)=>(e(),b(v(J),{key:`${s}-${k}`},{default:l(()=>[P(o(y.type),1)]),_:2},1024))),128))])):C("",!0)])]),_:2},[p.policies.length>0?{name:"accordion-content",fn:l(()=>[t("ul",Fe,[(e(!0),a(_,null,D(p.policies,(y,k)=>(e(),a("li",{key:`${s}-${k}`},[P(o(y.type)+": ",1),m(f,{to:y.route},{default:l(()=>[P(o(y.name),1)]),_:2},1032,["to"])]))),128))])]),key:"0"}:void 0]),1024))),128))]),_:2},1024)],64)):C("",!0)])]))),128))])])}}});const He=x(Ue,[["__scopeId","data-v-55dd6da4"]]),Ye={key:2,class:"policies-list"},Ze={key:3,class:"policies-list"},Je=G({__name:"DataplanePolicies",props:{dataPlane:{type:Object,required:!0}},setup(c){const i=c,I=z(),w=L(null),f=L([]),h=L([]),s=L([]),p=L(!0),g=L(null);K(()=>i.dataPlane.name,function(){y()}),y();async function y(){var $,u;g.value=null,p.value=!0,f.value=[],h.value=[],s.value=[];try{if(((u=($=i.dataPlane.networking.gateway)==null?void 0:$.type)==null?void 0:u.toUpperCase())==="BUILTIN")w.value=await E.getMeshGatewayDataplane({mesh:i.dataPlane.mesh,name:i.dataPlane.name}),h.value=k(w.value),s.value=O(w.value.policies);else{const{items:n}=await E.getSidecarDataplanePolicies({mesh:i.dataPlane.mesh,name:i.dataPlane.name});f.value=B(n)}}catch(r){r instanceof Error?g.value=r:console.error(r)}finally{p.value=!1}}function k($){const u=[];for(const r of $.listeners)for(const n of r.hosts)for(const d of n.routes){const M=[];for(const S of d.destinations){const A=O(S.policies),q={routeName:d.route,route:{name:"meshgatewayroutes",params:{mesh:$.gateway.mesh},query:{ns:d.route}},service:S.tags["kuma.io/service"],policies:A};M.push(q)}u.push({protocol:r.protocol,port:r.port,hostName:n.hostName,routeEntries:M})}return u}function O($){if($===void 0)return[];const u=[];for(const r of Object.values($)){const n=I.state.policiesByType[r.type];u.push({type:r.type,name:r.name,route:{name:n.path,params:{mesh:r.mesh},query:{ns:r.name}}})}return u}function B($){const u=[];for(const r of $){const n=[];for(const[A,q]of Object.entries(r.matchedPolicies)){const F=I.state.policiesByType[A],U=[];for(const R of q)U.push({name:R.name,route:{name:F.path,query:{ns:R.name},params:{mesh:R.mesh}}});n.push({name:A,pluralName:F.pluralDisplayName,policies:U})}const{name:d,type:M,service:S}=r;u.push({name:d,type:M,service:S,matchedPolicies:n})}return u}return($,u)=>p.value?(e(),b(se,{key:0})):g.value!==null?(e(),b(te,{key:1,error:g.value},null,8,["error"])):f.value.length>0?(e(),a("div",Ye,[m(Ae,{"sidecar-dataplane-policies":f.value},null,8,["sidecar-dataplane-policies"])])):h.value.length>0&&w.value!==null?(e(),a("div",Ze,[m(He,{"mesh-gateway-dataplane":w.value,"mesh-gateway-listener-entries":h.value,"mesh-gateway-route-policies":s.value},null,8,["mesh-gateway-dataplane","mesh-gateway-listener-entries","mesh-gateway-route-policies"])])):(e(),b(ae,{key:4}))}});const Xe=x(Je,[["__scopeId","data-v-041b4164"]]),ne=c=>(Q("data-v-c16c9c3e"),c=c(),ee(),c),Qe={key:0},ea={key:1},aa=ne(()=>t("h4",null,"Tags",-1)),ta=ne(()=>t("h4",null,"Versions",-1)),sa={class:"config-wrapper"},na={key:0},la=["href"],oa=G({__name:"DataPlaneDetails",props:{dataPlane:{type:Object,required:!0},dataPlaneOverview:{type:Object,required:!0}},setup(c){const i=c,I=z(),w=[{hash:"#overview",title:"Overview"},{hash:"#insights",title:"DPP Insights"},{hash:"#dpp-policies",title:"Policies"},{hash:"#xds-configuration",title:"XDS Configuration"},{hash:"#envoy-stats",title:"Stats"},{hash:"#envoy-clusters",title:"Clusters"},{hash:"#mtls",title:"Certificate Insights"},{hash:"#warnings",title:"Warnings"}],f=L([]),h=T(()=>{const{type:u,name:r,mesh:n}=i.dataPlane,d=ie(i.dataPlane,i.dataPlaneOverview.dataplaneInsight);return{type:u,name:r,mesh:n,status:d}}),s=T(()=>H(i.dataPlane)),p=T(()=>re(i.dataPlaneOverview.dataplaneInsight)),g=T(()=>ce(i.dataPlane)),y=T(()=>ue(i.dataPlaneOverview)),k=T(()=>{const u=Array.from(i.dataPlaneOverview.dataplaneInsight.subscriptions);return u.reverse(),u}),O=T(()=>{const u=I.getters.getKumaDocsVersion;return u!==null?u:"latest"}),B=T(()=>f.value.length===0?w.filter(u=>u.hash!=="#warnings"):w);function $(){const u=i.dataPlaneOverview.dataplaneInsight.subscriptions;if(u.length===0||!("version"in u[0]))return;const r=u[0].version;if(r&&r.kumaDp&&r.envoy){const d=_e(r);d.kind!==ye&&d.kind!==ve&&f.value.push(d)}I.getters["config/getMulticlusterStatus"]&&r&&H(i.dataPlane).find(S=>S.label===fe)&&typeof r.kumaDp.kumaCpCompatible=="boolean"&&!r.kumaDp.kumaCpCompatible&&f.value.push({kind:ge,payload:{kumaDp:r.kumaDp.version}})}return $(),(u,r)=>(e(),b($e,{tabs:v(B),"initial-tab-override":"overview"},{tabHeader:l(()=>[t("div",null,[t("h3",null," DPP: "+o(c.dataPlane.name),1)]),t("div",null,[m(we,{name:c.dataPlane.name,mesh:c.dataPlane.mesh},null,8,["name","mesh"])])]),overview:l(()=>[m(Y,null,{default:l(()=>[t("div",null,[t("ul",null,[(e(!0),a(_,null,D(v(h),(n,d)=>(e(),a("li",{key:d},[t("h4",null,o(d),1),d==="status"&&typeof n!="string"?(e(),a("div",Qe,[t("div",{class:de(["entity-status",{"is-offline":n.status.toLowerCase()==="offline","is-online":n.status.toLowerCase()==="online","is-degraded":n.status.toLowerCase()==="partially degraded","is-not-available":n.status.toLowerCase()==="not available"}])},[t("span",null,o(n.status),1)],2),(e(!0),a(_,null,D(n.reason,(M,S)=>(e(),a("div",{key:S,class:"reason"},o(M),1))),128))])):(e(),a("div",ea,o(n),1))]))),128))])]),t("div",null,[v(s).length>0?(e(),a(_,{key:0},[aa,t("p",null,[(e(!0),a(_,null,D(v(s),(n,d)=>(e(),a("span",{key:d,class:"tag-cols"},[t("span",null,o(n.label)+": ",1),t("span",null,o(n.value),1)]))),128))])],64)):C("",!0),v(p)?(e(),a(_,{key:1},[ta,t("p",null,[(e(!0),a(_,null,D(v(p),(n,d)=>(e(),a("span",{key:d,class:"tag-cols"},[t("span",null,o(d)+": ",1),t("span",null,o(n),1)]))),128))])],64)):C("",!0)])]),_:1}),t("div",sa,[m(Ie,{id:"code-block-data-plane",content:v(g),"is-searchable":""},null,8,["content"])])]),insights:l(()=>[m(ke,{"is-empty":v(k).length===0},{default:l(()=>[m(v(pe),{"border-variant":"noBorder"},{body:l(()=>[m(W,{"initially-open":0},{default:l(()=>[(e(!0),a(_,null,D(v(k),(n,d)=>(e(),b(j,{key:d},{"accordion-header":l(()=>[m(De,{details:n},null,8,["details"])]),"accordion-content":l(()=>[m(be,{details:n,"is-discovery-subscription":""},null,8,["details"])]),_:2},1024))),128))]),_:1})]),_:1})]),_:1},8,["is-empty"])]),"dpp-policies":l(()=>[m(Xe,{"data-plane":c.dataPlane},null,8,["data-plane"])]),"xds-configuration":l(()=>[m(V,{"data-path":"xds",mesh:c.dataPlane.mesh,"dpp-name":c.dataPlane.name,"query-key":"envoy-data-data-plane"},null,8,["mesh","dpp-name"])]),"envoy-stats":l(()=>[m(V,{"data-path":"stats",mesh:c.dataPlane.mesh,"dpp-name":c.dataPlane.name,"query-key":"envoy-data-data-plane"},null,8,["mesh","dpp-name"])]),"envoy-clusters":l(()=>[m(V,{"data-path":"clusters",mesh:c.dataPlane.mesh,"dpp-name":c.dataPlane.name,"query-key":"envoy-data-data-plane"},null,8,["mesh","dpp-name"])]),mtls:l(()=>[m(Y,null,{default:l(()=>[v(y)!==null?(e(),a("ul",na,[(e(!0),a(_,null,D(v(y),(n,d)=>(e(),a("li",{key:d},[t("h4",null,o(n.label),1),t("p",null,o(n.value),1)]))),128))])):(e(),b(v(me),{key:1,appearance:"danger"},{alertMessage:l(()=>[P(" This data plane proxy does not yet have mTLS configured \u2014 "),t("a",{href:`https://kuma.io/docs/${v(O)}/documentation/security/#certificates`,class:"external-link",target:"_blank"}," Learn About Certificates in "+o(v(he)),9,la)]),_:1}))]),_:1})]),warnings:l(()=>[m(Le,{warnings:f.value},null,8,["warnings"])]),_:1},8,["tabs"]))}});const ia=x(oa,[["__scopeId","data-v-c16c9c3e"]]),ra={class:"component-frame"},Da=G({__name:"DataPlaneDetailView",setup(c){const i=Pe(),I=z(),w=L(null),f=L(null),h=L(!0),s=L(null);async function p(){s.value=null,h.value=!0;const g=i.params.mesh,y=i.params.dataPlane;try{w.value=await E.getDataplaneFromMesh({mesh:g,name:y}),f.value=await E.getDataplaneOverviewFromMesh({mesh:g,name:y})}catch(k){w.value=null,k instanceof Error?s.value=k:console.error(k)}finally{h.value=!1}}return K(()=>i.params.mesh,function(){i.name==="data-plane-detail-view"&&p()}),K(()=>i.params.dataPlane,function(){i.name==="data-plane-detail-view"&&p()}),p(),I.dispatch("updatePageTitle",i.params.dataPlane),(g,y)=>(e(),a("div",ra,[h.value?(e(),b(se,{key:0})):s.value!==null?(e(),b(te,{key:1,error:s.value},null,8,["error"])):w.value===null||f.value===null?(e(),b(ae,{key:2})):(e(),b(ia,{key:3,"data-plane":w.value,"data-plane-overview":f.value},null,8,["data-plane","data-plane-overview"]))]))}});export{Da as default};
