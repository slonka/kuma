import{E as u}from"./EnvoyData-9e2c0d24.js";import{a as i}from"./dataplane-0a086c06.js";import{d as _,a as e,o as h,b as f,w as t,e as s,p as w,f as V,q as y}from"./index-dc1529df.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-415826e6.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-79530b3b.js";const q=_({__name:"DataPlaneStatsView",props:{data:{}},setup(o){const n=o;return(C,g)=>{const r=e("RouteTitle"),p=e("KCard"),l=e("AppView"),c=e("RouteView");return h(),f(c,{name:"data-plane-stats-view",params:{mesh:"",dataPlane:"",codeSearch:""}},{default:t(({route:a,t:d})=>[s(l,null,{title:t(()=>[w("h2",null,[s(r,{title:d("data-planes.routes.item.navigation.data-plane-stats-view"),render:!0},null,8,["title"])])]),default:t(()=>[V(),s(p,null,{body:t(()=>[s(u,{status:y(i)(n.data.dataplane,n.data.dataplaneInsight).status,resource:"Data Plane Proxy",src:`/meshes/${a.params.mesh}/dataplanes/${a.params.dataPlane}/data-path/stats`,query:a.params.codeSearch,onQueryChange:m=>a.update({codeSearch:m})},null,8,["status","src","query","onQueryChange"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{q as default};
