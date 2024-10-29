import{d as C,e as t,o as w,m as x,w as s,a as o,b as c,l as R,a3 as $}from"./index-z3MQYrrd.js";const V=C({__name:"ConnectionInboundSummaryStatsView",props:{data:{},dataplaneOverview:{}},setup(r){const e=r;return(y,k)=>{const p=t("RouteTitle"),i=t("KButton"),l=t("XCodeBlock"),m=t("DataCollection"),_=t("DataLoader"),u=t("AppView"),h=t("RouteView");return w(),x(h,{params:{codeSearch:"",codeFilter:!1,codeRegExp:!1,mesh:"",dataPlane:"",connection:""},name:"connection-inbound-summary-stats-view"},{default:s(({route:a})=>[o(p,{render:!1,title:"Stats"}),c(),o(u,null,{default:s(()=>[o(_,{src:`/meshes/${a.params.mesh}/dataplanes/${a.params.dataPlane}/stats/${e.dataplaneOverview.dataplane.networking.inboundAddress}`},{default:s(({data:g,refresh:f})=>[o(m,{items:g.raw.split(`
`),predicate:d=>[`listener.${e.data.listenerAddress.length>0?e.data.listenerAddress:a.params.connection}`,`cluster.${e.data.name}.`,`http.${e.data.name}.`,`tcp.${e.data.name}.`].some(n=>d.startsWith(n))&&(!d.includes(".rds.")||d.includes(`_${e.data.port}`))},{default:s(({items:d})=>[o(l,{language:"json",code:d.map(n=>n.replace(`${e.data.listenerAddress.length>0?e.data.listenerAddress:a.params.connection}.`,"").replace(`${e.data.name}.`,"")).join(`
`),"is-searchable":"",query:a.params.codeSearch,"is-filter-mode":a.params.codeFilter,"is-reg-exp-mode":a.params.codeRegExp,onQueryChange:n=>a.update({codeSearch:n}),onFilterModeChange:n=>a.update({codeFilter:n}),onRegExpModeChange:n=>a.update({codeRegExp:n})},{"primary-actions":s(()=>[o(i,{appearance:"primary",onClick:f},{default:s(()=>[o(R($)),c(`

                Refresh
              `)]),_:2},1032,["onClick"])]),_:2},1032,["code","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])]),_:2},1032,["items","predicate"])]),_:2},1032,["src"])]),_:2},1024)]),_:1})}}});export{V as default};