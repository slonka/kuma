import{d as E,r as t,o as l,m as y,w as o,b as i,e as d,k as m,c as _,F as u,s as f,n as F,t as V,E as K,q as T}from"./index-DHg9Fngg.js";const X={class:"policy-list-content"},$={class:"policy-count"},q={class:"policy-list"},z=E({__name:"PolicyTypeListView",setup(j){return(G,H)=>{const C=t("RouteTitle"),D=t("XAction"),R=t("DataCollection"),B=t("DataLoader"),P=t("KCard"),x=t("RouterView"),w=t("DataSource"),L=t("AppView"),A=t("RouteView");return l(),y(A,{name:"policy-list-view",params:{mesh:"",policyPath:"",policy:""}},{default:o(({route:r,t:N})=>[i(C,{render:!1,title:N("policies.routes.types.title")},null,8,["title"]),d(),i(L,null,{default:o(()=>[i(w,{src:`/mesh-insights/${r.params.mesh}`},{default:o(({data:e})=>[i(w,{src:"/policy-types"},{default:o(({data:c,error:S})=>[m("div",X,[i(P,{class:"policy-type-list","data-testid":"policy-type-list"},{default:o(()=>[i(B,{data:[c],errors:[S]},{default:o(()=>[(l(!0),_(u,null,f([typeof(e==null?void 0:e.policies)>"u"?c.policies:c.policies.filter(s=>{var p,a;return!s.isTargetRefBased&&(((a=(p=e.policies)==null?void 0:p[s.name])==null?void 0:a.total)??0)>0})],s=>(l(),y(R,{key:s,predicate:typeof(e==null?void 0:e.policies)>"u"?void 0:p=>s.length>0||p.isTargetRefBased,items:c.policies},{default:o(({items:p})=>[(l(!0),_(u,null,f([p.find(a=>a.path===r.params.policyPath)],a=>(l(),_(u,{key:a},[(l(!0),_(u,null,f(p,(n,b)=>{var v,k;return l(),_("div",{key:n.path,class:F(["policy-type-link-wrapper",{"policy-type-link-wrapper--is-active":a&&a.path===n.path}])},[i(D,{class:"policy-type-link",to:{name:"policy-list-view",params:{mesh:r.params.mesh,policyPath:n.path}},mount:r.params.policyPath.length===0&&b===0?r.replace:void 0,"data-testid":`policy-type-link-${n.name}`},{default:o(()=>[d(V(n.name),1)]),_:2},1032,["to","mount","data-testid"]),d(),m("div",$,V(((k=(v=e==null?void 0:e.policies)==null?void 0:v[n.name])==null?void 0:k.total)??0),1)],2)}),128))],64))),128))]),_:2},1032,["predicate","items"]))),128))]),_:2},1032,["data","errors"])]),_:2},1024),d(),m("div",q,[i(x,null,{default:o(({Component:s})=>[(l(),y(K(s),{"policy-types":c==null?void 0:c.policies},null,8,["policy-types"]))]),_:2},1024)])])]),_:2},1024)]),_:2},1032,["src"])]),_:2},1024)]),_:1})}}}),M=T(z,[["__scopeId","data-v-072b5a15"]]);export{M as default};
