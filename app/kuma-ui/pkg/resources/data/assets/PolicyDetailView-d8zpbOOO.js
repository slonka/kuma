import{_ as E}from"./ResourceCodeBlock.vue_vue_type_style_index_0_lang-owWKimIX.js";import{d as V,a as i,o as t,b as n,w as o,W as P,e as l,m as _,f as p,E as f,A as g,t as x,a2 as F,F as B,c as d,H as S,p as L}from"./index-1j9z4Egf.js";import"./CodeBlock-ejECcgv-.js";import"./toYaml-sPaYOD3i.js";const I={key:2,class:"stack","data-testid":"detail-view-details"},K={class:"mt-4"},M={key:3,"data-testid":"affected-data-plane-proxies"},j=V({__name:"PolicyDetailView",setup(N){return(T,q)=>{const C=i("RouteTitle"),k=i("KInput"),w=i("RouterLink"),h=i("DataSource"),v=i("KCard"),$=i("AppView"),R=i("RouteView");return t(),n(R,{name:"policy-detail-view",params:{mesh:"",policy:"",policyPath:"",codeSearch:"",codeFilter:!1,codeRegExp:!1,dataplane:""}},{default:o(({route:e,t:r})=>[l($,{breadcrumbs:[{to:{name:"mesh-detail-view",params:{mesh:e.params.mesh}},text:e.params.mesh},{to:{name:"policy-list-view",params:{mesh:e.params.mesh,policyPath:e.params.policyPath}},text:r("policies.routes.item.breadcrumbs")}]},{title:o(()=>[_("h1",null,[l(P,{text:e.params.policy},{default:o(()=>[l(C,{title:r("policies.routes.item.title",{name:e.params.policy})},null,8,["title"])]),_:2},1032,["text"])])]),default:o(()=>[p(),l(h,{src:`/meshes/${e.params.mesh}/policy-path/${e.params.policyPath}/policy/${e.params.policy}`},{default:o(({data:u,error:y})=>[y?(t(),n(f,{key:0,error:y},null,8,["error"])):u===void 0?(t(),n(g,{key:1})):(t(),d("div",I,[l(v,null,{default:o(()=>[_("h2",null,x(r("policies.detail.affected_dpps")),1),p(),_("div",K,[l(k,{"model-value":e.params.dataplane,type:"text",placeholder:r("policies.detail.dataplane_input_placeholder"),required:"","data-testid":"dataplane-search-input",onInput:a=>e.update({dataplane:a})},null,8,["model-value","placeholder","onInput"]),p(),l(h,{src:`/meshes/${e.params.mesh}/policy-path/${e.params.policyPath}/policy/${e.params.policy}/dataplanes`},{default:o(({data:a,error:c})=>[c?(t(),n(f,{key:0,error:c},null,8,["error"])):a===void 0?(t(),n(g,{key:1})):a.items.length===0?(t(),n(F,{key:2})):(t(),d("ul",M,[(t(!0),d(B,null,S(a.items.filter(s=>s.name.toLowerCase().includes(e.params.dataplane.toLowerCase())),(s,m)=>(t(),d("li",{key:m,"data-testid":"dataplane-name"},[l(w,{to:{name:"data-plane-detail-view",params:{mesh:s.mesh,dataPlane:s.name}}},{default:o(()=>[p(x(s.name),1)]),_:2},1032,["to"])]))),128))]))]),_:2},1032,["src"])])]),_:2},1024),p(),l(E,{resource:u.config,"is-searchable":"",query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter,"is-reg-exp-mode":e.params.codeRegExp,onQueryChange:a=>e.update({codeSearch:a}),onFilterModeChange:a=>e.update({codeFilter:a}),onRegExpModeChange:a=>e.update({codeRegExp:a})},{default:o(({copy:a,copying:c})=>[c?(t(),n(h,{key:0,src:`/meshes/${e.params.mesh}/policy-path/${e.params.policyPath}/policy/${e.params.policy}/as/kubernetes?no-store`,onChange:s=>{a(m=>m(s))},onError:s=>{a((m,b)=>b(s))}},null,8,["src","onChange","onError"])):L("",!0)]),_:2},1032,["resource","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])]))]),_:2},1032,["src"])]),_:2},1032,["breadcrumbs"])]),_:1})}}});export{j as default};
