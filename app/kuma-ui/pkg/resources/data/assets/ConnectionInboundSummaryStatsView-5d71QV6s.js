import{d as R,a as o,o as i,b as l,w as s,e as r,m as _,f,E as S,A as k,a4 as E,q as h,K as V}from"./index-1j9z4Egf.js";import{C as v}from"./CodeBlock-ejECcgv-.js";const F=R({__name:"ConnectionInboundSummaryStatsView",props:{inbound:{},gateway:{}},setup(g){const C=g,y=(p,d)=>{if(C.gateway){const a=d.split(":");return p.raw.split(`
`).filter(e=>e.includes(`.${a[0]}.`)).filter(e=>!e.startsWith("listener.")||e.includes(`_${a[1]}.`)).filter(e=>!e.includes(".rds.")||e.includes(`_${a[1]}_*.`)).map(e=>e.replace(`${a[0]}.`,"")).map(e=>e.replace(`_${a[1]}.`,".")).map(e=>{if(e.includes(".rds.")){const c=e.split(".");return c.splice(2,1),c.join(".")}else return e}).join(`
`)}else{const a=`localhost_${d.split(":")[1]}`;return p.raw.split(`
`).filter(e=>e.includes(`.${a}.`)).map(e=>e.replace(`${a}.`,"")).join(`
`)}};return(p,d)=>{const a=o("RouteTitle"),e=o("KButton"),c=o("DataSource"),w=o("AppView"),$=o("RouteView");return i(),l($,{params:{codeSearch:"",codeFilter:!1,codeRegExp:!1,mesh:"",dataPlane:"",service:""},name:"connection-inbound-summary-stats-view"},{default:s(({route:n})=>[r(w,null,{title:s(()=>[_("h3",null,[r(a,{title:"Stats"})])]),default:s(()=>[f(),_("div",null,[r(c,{src:`/meshes/${n.params.mesh}/dataplanes/${n.params.dataPlane}/stats/${n.params.service}`},{default:s(({data:m,error:u,refresh:x})=>[u?(i(),l(S,{key:0,error:u},null,8,["error"])):m===void 0?(i(),l(k,{key:1})):(i(),l(v,{key:2,language:"json",code:y(m,n.params.service),"is-searchable":"",query:n.params.codeSearch,"is-filter-mode":n.params.codeFilter,"is-reg-exp-mode":n.params.codeRegExp,onQueryChange:t=>n.update({codeSearch:t}),onFilterModeChange:t=>n.update({codeFilter:t}),onRegExpModeChange:t=>n.update({codeRegExp:t})},{"primary-actions":s(()=>[r(e,{appearance:"primary",onClick:x},{default:s(()=>[r(h(E),{size:h(V)},null,8,["size"]),f(`
                Refresh
              `)]),_:2},1032,["onClick"])]),_:2},1032,["code","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"]))]),_:2},1032,["src"])])]),_:2},1024)]),_:1})}}});export{F as default};
