import{_ as f}from"./CodeBlock.vue_vue_type_style_index_0_lang-226d1ddf.js";import{E as g}from"./ErrorBlock-d38c2168.js";import{_ as b}from"./LoadingBlock.vue_vue_type_script_setup_true_lang-8f5d9bcc.js";import{d as h,a as e,o as t,b as a,w as o,e as s,p as y,f as k}from"./index-784d2bbf.js";import"./index-9dd3e7d3.js";import"./TextWithCopyButton-7ef74197.js";import"./CopyButton-9c00109a.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-9960c4c9.js";const v=h({__name:"DiagnosticsDetailView",setup(V){return(w,C)=>{const d=e("RouteTitle"),m=e("KCard"),l=e("AppView"),u=e("DataSource"),p=e("RouteView");return t(),a(p,{name:"diagnostics",params:{codeSearch:""}},{default:o(({route:c,t:n})=>[s(u,{src:"/config"},{default:o(({data:r,error:i})=>[s(l,{breadcrumbs:[{to:{name:"diagnostics"},text:n("diagnostics.routes.item.breadcrumbs")}]},{title:o(()=>[y("h1",null,[s(d,{title:n("diagnostics.routes.item.title")},null,8,["title"])])]),default:o(()=>[k(),s(m,null,{body:o(()=>[i?(t(),a(g,{key:0,error:i},null,8,["error"])):r===void 0?(t(),a(b,{key:1})):(t(),a(f,{key:2,id:"code-block-diagnostics","data-testid":"code-block-diagnostics",language:"json",code:JSON.stringify(r,null,2),"is-searchable":"",query:c.params.codeSearch,onQueryChange:_=>c.update({codeSearch:_})},null,8,["code","query","onQueryChange"]))]),_:2},1024)]),_:2},1032,["breadcrumbs"])]),_:2},1024)]),_:1})}}});export{v as default};
