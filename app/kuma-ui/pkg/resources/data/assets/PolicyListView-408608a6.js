import{K as x}from"./index-fce48c05.js";import{d as V,k as L,y as I,H as N,a as g,o,c as b,e as m,w as a,F as w,C as E,n as K,l,f as s,t as p,m as u,b as n,p as v,T as A,_ as F,O as M,B as O}from"./index-671fe4fd.js";import{D as q,A as D}from"./AppCollection-a9aee10e.js";import{E as $}from"./ErrorBlock-12fc0e3a.js";import{P as G}from"./PolicyTypeTag-6da32bfa.js";import{_ as Z}from"./LoadingBlock.vue_vue_type_script_setup_true_lang-e439d559.js";import{S as H}from"./SummaryView-61bb43be.js";import"./TextWithCopyButton-fcf481cf.js";import"./CopyButton-2a28adb7.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-1e209f96.js";const U={class:"policy-list-content"},j={class:"policy-count"},J={class:"policy-list"},Q={class:"stack"},W={class:"description"},X={class:"description-content"},Y={class:"description-actions"},ee={class:"visually-hidden"},te={key:0},ae=V({__name:"PolicyList",props:{pageNumber:{},pageSize:{},policyTypes:{},currentPolicyType:{},policyCollection:{},policyError:{},meshInsight:{},isSelectedRow:{type:[Function,null],default:null}},emits:["change"],setup(B,{emit:S}){const{t:c}=L(),T=I(),e=B,P=S,C=N(()=>{const z=e.policyTypes.filter(y=>!y.isTargetRefBased&&y.name!=="MeshGateway").some(y=>{var t,h,r;return(((r=(h=(t=e.meshInsight)==null?void 0:t.policies)==null?void 0:h[y.name])==null?void 0:r.total)??0)>0});return e.policyTypes.filter(y=>y.isTargetRefBased||y.name==="MeshGateway"?!0:z)});return(z,y)=>{const t=g("RouterLink"),h=g("KCard"),r=g("KBadge");return o(),b("div",U,[m(h,{class:"policy-type-list","data-testid":"policy-type-list"},{default:a(()=>[(o(!0),b(w,null,E(C.value,(d,_)=>{var i,k,f;return o(),b("div",{key:_,class:K(["policy-type-link-wrapper",{"policy-type-link-wrapper--is-active":d.path===e.currentPolicyType.path}])},[m(t,{class:"policy-type-link",to:{name:"policy-list-view",params:{mesh:l(T).params.mesh,policyPath:d.path}},"data-testid":`policy-type-link-${d.name}`},{default:a(()=>[s(p(d.name),1)]),_:2},1032,["to","data-testid"]),s(),u("div",j,p(((f=(k=(i=e.meshInsight)==null?void 0:i.policies)==null?void 0:k[d.name])==null?void 0:f.total)??0),1)],2)}),128))]),_:1}),s(),u("div",J,[u("div",Q,[m(h,null,{default:a(()=>[u("div",W,[u("div",X,[u("h3",null,[m(G,{"policy-type":e.currentPolicyType.name},{default:a(()=>[s(p(l(c)("policies.collection.title",{name:e.currentPolicyType.name})),1)]),_:1},8,["policy-type"])]),s(),u("p",null,p(l(c)(`policies.type.${e.currentPolicyType.name}.description`,void 0,{defaultMessage:l(c)("policies.collection.description")})),1)]),s(),u("div",Y,[e.currentPolicyType.isExperimental?(o(),n(r,{key:0,appearance:"warning"},{default:a(()=>[s(p(l(c)("policies.collection.beta")),1)]),_:1})):v("",!0),s(),e.currentPolicyType.isInbound?(o(),n(r,{key:1,appearance:"neutral"},{default:a(()=>[s(p(l(c)("policies.collection.inbound")),1)]),_:1})):v("",!0),s(),e.currentPolicyType.isOutbound?(o(),n(r,{key:2,appearance:"neutral"},{default:a(()=>[s(p(l(c)("policies.collection.outbound")),1)]),_:1})):v("",!0),s(),m(q,{href:l(c)("policies.href.docs",{name:e.currentPolicyType.name}),"data-testid":"policy-documentation-link"},{default:a(()=>[u("span",ee,p(l(c)("common.documentation")),1)]),_:1},8,["href"])])])]),_:1}),s(),m(h,null,{default:a(()=>{var d,_;return[e.policyError!==void 0?(o(),n($,{key:0,error:e.policyError},null,8,["error"])):(o(),n(D,{key:1,class:"policy-collection","data-testid":"policy-collection","empty-state-message":l(c)("common.emptyState.message",{type:`${e.currentPolicyType.name} policies`}),"empty-state-cta-to":l(c)("policies.href.docs",{name:e.currentPolicyType.name}),"empty-state-cta-text":l(c)("common.documentation"),headers:[{label:"Name",key:"name"},...e.currentPolicyType.isTargetRefBased?[{label:"Zone",key:"zone"}]:[],...e.currentPolicyType.isTargetRefBased?[{label:"Target ref",key:"targetRef"}]:[],{label:"Details",key:"details",hideLabel:!0}],"page-number":e.pageNumber,"page-size":e.pageSize,total:(d=e.policyCollection)==null?void 0:d.total,items:(_=e.policyCollection)==null?void 0:_.items,error:e.policyError,"is-selected-row":e.isSelectedRow,onChange:y[0]||(y[0]=i=>P("change",i))},{name:a(({row:i})=>[m(t,{to:{name:"policy-summary-view",params:{mesh:i.mesh,policyPath:e.currentPolicyType.path,policy:i.name},query:{page:e.pageNumber,size:e.pageSize}}},{default:a(()=>[s(p(i.name),1)]),_:2},1032,["to"])]),targetRef:a(({row:i})=>{var k;return[e.currentPolicyType.isTargetRefBased&&typeof((k=i.spec)==null?void 0:k.targetRef)<"u"?(o(),n(r,{key:0,appearance:"neutral"},{default:a(()=>[s(p(i.spec.targetRef.kind),1),i.spec.targetRef.name?(o(),b("span",te,[s(":"),u("b",null,p(i.spec.targetRef.name),1)])):v("",!0)]),_:2},1024)):(o(),b(w,{key:1},[s(p(l(c)("common.detail.none")),1)],64))]}),zone:a(({row:i})=>[i.labels&&i.labels["kuma.io/origin"]==="zone"&&i.labels["kuma.io/zone"]?(o(),n(t,{key:0,to:{name:"zone-cp-detail-view",params:{zone:i.labels["kuma.io/zone"]}}},{default:a(()=>[s(p(i.labels["kuma.io/zone"]),1)]),_:2},1032,["to"])):(o(),b(w,{key:1},[s(p(l(c)("common.detail.none")),1)],64))]),details:a(({row:i})=>[m(t,{class:"details-link","data-testid":"details-link",to:{name:"policy-detail-view",params:{mesh:i.mesh,policyPath:e.currentPolicyType.path,policy:i.name}}},{default:a(()=>[s(p(l(c)("common.collection.details_link"))+" ",1),m(l(A),{display:"inline-block",decorative:"",size:l(x)},null,8,["size"])]),_:2},1032,["to"])]),_:1},8,["empty-state-message","empty-state-cta-to","empty-state-cta-text","headers","page-number","page-size","total","items","error","is-selected-row"]))]}),_:1})])])])}}});const ie=F(ae,[["__scopeId","data-v-efdb2ab3"]]),ue=V({__name:"PolicyListView",setup(B){return(S,c)=>{const T=g("RouteTitle"),e=g("RouterView"),P=g("DataSource"),C=g("AppView"),z=g("RouteView");return o(),n(P,{src:"/me"},{default:a(({data:y})=>[y?(o(),n(z,{key:0,name:"policy-list-view",params:{page:1,size:y.pageSize,mesh:"",policyPath:"",policy:""}},{default:a(({route:t,t:h})=>[m(C,null,{title:a(()=>[u("h2",null,[m(T,{title:h("policies.routes.items.title")},null,8,["title"])])]),default:a(()=>[s(),m(P,{src:"/*/policy-types"},{default:a(({data:r,error:d})=>[d?(o(),n($,{key:0,error:d},null,8,["error"])):r===void 0?(o(),n(Z,{key:1})):r.policies.length===0?(o(),n(M,{key:2})):(o(),n(P,{key:3,src:`/meshes/${t.params.mesh}/policy-path/${t.params.policyPath}?page=${t.params.page}&size=${t.params.size}`},{default:a(({data:_,error:i})=>[m(P,{src:`/mesh-insights/${t.params.mesh}`},{default:a(({data:k})=>[(o(),n(ie,{key:t.params.policyPath,"page-number":t.params.page,"page-size":t.params.size,"current-policy-type":r.policies.find(f=>f.path===t.params.policyPath)??r.policies[0],"policy-types":r.policies,"mesh-insight":k,"policy-collection":_,"policy-error":i,"is-selected-row":f=>f.name===t.params.policy,onChange:t.update},null,8,["page-number","page-size","current-policy-type","policy-types","mesh-insight","policy-collection","policy-error","is-selected-row","onChange"])),s(),t.params.policy?(o(),n(e,{key:0},{default:a(f=>[m(H,{onClose:R=>t.replace({name:"policy-list-view",params:{mesh:t.params.mesh,policyPath:t.params.policyPath},query:{page:t.params.page,size:t.params.size}})},{default:a(()=>[(o(),n(O(f.Component),{policy:_==null?void 0:_.items.find(R=>R.name===t.params.policy),"policy-type":r.policies.find(R=>R.path===t.params.policyPath)},null,8,["policy","policy-type"]))]),_:2},1032,["onClose"])]),_:2},1024)):v("",!0)]),_:2},1032,["src"])]),_:2},1032,["src"]))]),_:2},1024)]),_:2},1024)]),_:2},1032,["params"])):v("",!0)]),_:1})}}});export{ue as default};
