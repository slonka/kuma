import{_ as D}from"./PolicySummary.vue_vue_type_script_setup_true_lang-CZxFLuhS.js";import{d as T,r as t,o as c,m,w as s,b as r,s as i,e as d,t as h,c as B,F as X,v as g,T as q,p as L,a0 as A,q as _,_ as N}from"./index-ChH5weWG.js";import{_ as x}from"./ResourceCodeBlock.vue_vue_type_script_setup_true_lang-rvKSlgEf.js";const Q=T({__name:"DataPlanePolicySummaryView",props:{policyTypes:{}},setup(C){const k=C;return(v,l)=>{const R=t("RouteTitle"),S=t("XAction"),w=t("XSelect"),V=t("XLayout"),y=t("DataLoader"),F=t("AppView"),b=t("DataSource"),E=t("RouteView");return c(),m(E,{name:"data-plane-policy-summary-view",params:{mesh:"",policyPath:"",policy:"",codeSearch:"",codeFilter:!1,codeRegExp:!1,format:String}},{default:s(({route:e,t:n,uri:M})=>[r(b,{src:`/meshes/${e.params.mesh}/policy-path/${e.params.policyPath}/policy/${e.params.policy}`},{default:s(({data:p,error:P})=>[r(F,null,{title:s(()=>[i("h2",null,[r(S,{to:{name:"policy-detail-view",params:{mesh:e.params.mesh,policyPath:e.params.policyPath,policy:e.params.policy}}},{default:s(()=>[r(R,{title:n("policies.routes.item.title",{name:e.params.policy})},null,8,["title"])]),_:2},1032,["to"])])]),default:s(()=>[l[2]||(l[2]=d()),r(y,{data:[p],errors:[P]},{default:s(()=>{var f;return[p?(c(),m(D,{key:0,policy:p,format:e.params.format,legacy:!((f=k.policyTypes.find(({name:o})=>o===(p==null?void 0:p.type)))!=null&&f.policy.isTargetRef)},{header:s(()=>[i("header",null,[r(V,{type:"separated",size:"max"},{default:s(()=>[i("h3",null,h(n("policies.routes.item.config")),1),l[0]||(l[0]=d()),(c(),B(X,null,g([["structured","universal","k8s"]],o=>i("div",{key:typeof o},[r(w,{label:n("policies.routes.item.format"),selected:o.includes(e.params.format)?e.params.format:o[0],onChange:a=>{e.update({format:a})},onVnodeBeforeMount:a=>{var u;return((u=a==null?void 0:a.props)==null?void 0:u.selected)&&o.includes(a.props.selected)&&a.props.selected!==e.params.format&&e.update({format:a.props.selected})}},q({_:2},[g(o,a=>({name:`${a}-option`,fn:s(()=>[d(h(n(`policies.routes.item.formats.${a}`)),1)])}))]),1032,["label","selected","onChange","onVnodeBeforeMount"])])),64))]),_:2},1024)])]),default:s(()=>[l[1]||(l[1]=d()),e.params.format==="universal"?(c(),m(x,{key:0,"data-testid":"codeblock-yaml-universal",language:"yaml",resource:p.config,"is-searchable":"","show-k8s-copy-button":!1,query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter,"is-reg-exp-mode":e.params.codeRegExp,onQueryChange:o=>e.update({codeSearch:o}),onFilterModeChange:o=>e.update({codeFilter:o}),onRegExpModeChange:o=>e.update({codeRegExp:o})},null,8,["resource","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])):e.params.format==="k8s"?(c(),m(y,{key:1,src:M(L(A),"/meshes/:mesh/policy-path/:path/policy/:name/as/kubernetes",{mesh:e.params.mesh,path:e.params.policyPath,name:e.params.policy})},{default:s(({data:o})=>[r(x,{"data-testid":"codeblock-yaml-k8s",language:"yaml",resource:o,"show-k8s-copy-button":!1,"is-searchable":"",query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter,"is-reg-exp-mode":e.params.codeRegExp,onQueryChange:a=>e.update({codeSearch:a}),onFilterModeChange:a=>e.update({codeFilter:a}),onRegExpModeChange:a=>e.update({codeRegExp:a})},null,8,["resource","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])]),_:2},1032,["src"])):_("",!0)]),_:2},1032,["policy","format","legacy"])):_("",!0)]}),_:2},1032,["data","errors"])]),_:2},1024)]),_:2},1032,["src"])]),_:1})}}}),G=N(Q,[["__scopeId","data-v-31f9d775"]]);export{G as default};
