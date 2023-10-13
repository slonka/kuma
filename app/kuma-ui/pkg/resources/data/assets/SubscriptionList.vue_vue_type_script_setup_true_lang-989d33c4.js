import{A as j,a as B}from"./AccordionList-1a487a67.js";import{d as S,g as x,h as b,o as c,l as u,i as g,w as v,j as f,k as a,an as L,n,H as s,ac as P,p as t,F as y,I as A,t as T,m as $,D as R,G as H}from"./index-adcc6fc8.js";const N={class:"stack"},U={key:1},V={class:"row"},C={class:"header"},O={class:"header"},q=["data-testid"],F={class:"type"},E=S({__name:"SubscriptionDetails",props:{subscription:{type:Object,required:!0},isDiscoverySubscription:{type:Boolean,default:!1}},setup(r){const e=r,{t:o}=x(),p=b(()=>{var d;let l;if("controlPlaneInstanceId"in e.subscription){const{lastUpdateTime:i,total:_,...m}=e.subscription.status;l=m}else l=((d=e.subscription.status)==null?void 0:d.stat)??{};return l?Object.entries(l).map(([i,_])=>{const{responsesSent:m="0",responsesAcknowledged:h="0",responsesRejected:I="0"}=_;return{type:i,responsesSent:m,responsesAcknowledged:h,responsesRejected:I}}):[]});return(l,d)=>(c(),u("div",N,[p.value.length===0?(c(),g(a(P),{key:0,appearance:"info"},{alertIcon:v(()=>[f(a(L))]),alertMessage:v(()=>[n(s(a(o)("common.detail.subscriptions.no_stats",{id:e.subscription.id})),1)]),_:1})):(c(),u("div",U,[t("div",V,[t("div",C,s(a(o)("common.detail.subscriptions.type")),1),n(),t("div",O,s(a(o)("common.detail.subscriptions.responses_sent_acknowledged")),1)]),n(),(c(!0),u(y,null,A(p.value,(i,_)=>(c(),u("div",{key:_,class:"row","data-testid":`subscription-status-${i.type}`},[t("div",F,s(a(o)(`http.api.property.${i.type}`)),1),n(),t("div",null,s(i.responsesSent)+"/"+s(i.responsesAcknowledged),1)],8,q))),128))]))]))}});const G=T(E,[["__scopeId","data-v-ae0695b5"]]),M=""+new URL("icon-deployed-code-e3c999ba.svg",import.meta.url).href,z=""+new URL("icon-connected-037e001a.svg",import.meta.url).href,J=""+new URL("icon-disconnected-ba3c2624.svg",import.meta.url).href,k=r=>(R("data-v-0cbdfb03"),r=r(),H(),r),K={class:"subscription-header"},Q={class:"instance-id"},W=k(()=>t("img",{src:M},null,-1)),X=k(()=>t("img",{src:z},null,-1)),Y={key:0},Z=k(()=>t("img",{src:J},null,-1)),ss={class:"responses-sent-acknowledged"},es=S({__name:"SubscriptionHeader",props:{subscription:{type:Object,required:!0}},setup(r){const e=r,{t:o,formatIsoDate:p}=x(),l=b(()=>"globalInstanceId"in e.subscription?e.subscription.globalInstanceId:null),d=b(()=>"controlPlaneInstanceId"in e.subscription?e.subscription.controlPlaneInstanceId:null),i=b(()=>e.subscription.connectTime?p(e.subscription.connectTime):null),_=b(()=>e.subscription.disconnectTime?p(e.subscription.disconnectTime):null),m=b(()=>{var w;const{responsesSent:h=0,responsesAcknowledged:I=0,responsesRejected:D=0}=((w=e.subscription.status)==null?void 0:w.total)??{};return{responsesSent:h,responsesAcknowledged:I,responsesRejected:D}});return(h,I)=>(c(),u("header",K,[t("span",Q,[W,n(),l.value?(c(),u(y,{key:0},[t("b",null,s(a(o)("http.api.property.globalInstanceId")),1),n(": "+s(l.value),1)],64)):d.value?(c(),u(y,{key:1},[t("b",null,s(a(o)("http.api.property.controlPlaneInstanceId")),1),n(": "+s(d.value),1)],64)):$("",!0)]),n(),t("span",null,[X,n(),t("b",null,s(a(o)("common.detail.subscriptions.connect_time")),1),n(": "+s(i.value),1)]),n(),_.value?(c(),u("span",Y,[Z,n(),t("b",null,s(a(o)("common.detail.subscriptions.disconnect_time")),1),n(": "+s(_.value),1)])):$("",!0),n(),t("span",ss,s(a(o)("common.detail.subscriptions.responses_sent_acknowledged"))+`:

      `+s(m.value.responsesSent)+"/"+s(m.value.responsesAcknowledged),1)]))}});const ts=T(es,[["__scopeId","data-v-0cbdfb03"]]),cs=S({__name:"SubscriptionList",props:{subscriptions:{}},setup(r){const e=r,o=b(()=>{const p=Array.from(e.subscriptions);return p.reverse(),p});return(p,l)=>(c(),g(B,null,{default:v(()=>[(c(!0),u(y,null,A(o.value,(d,i)=>(c(),g(j,{key:i},{"accordion-header":v(()=>[f(ts,{subscription:d},null,8,["subscription"])]),"accordion-content":v(()=>[f(G,{subscription:d},null,8,["subscription"])]),_:2},1024))),128))]),_:1}))}});export{cs as _};
