import{d as C,e as n,o as w,m as x,w as o,a as t,b as s}from"./index-CgC5RQPZ.js";const V=C({__name:"ConnectionOutboundSummaryStatsView",props:{dataplaneOverview:{}},setup(r){const i=r;return(R,y)=>{const p=n("RouteTitle"),d=n("XAction"),l=n("XCodeBlock"),m=n("DataCollection"),_=n("DataLoader"),u=n("AppView"),h=n("RouteView");return w(),x(h,{params:{codeSearch:"",codeFilter:!1,codeRegExp:!1,mesh:"",dataPlane:"",connection:""},name:"connection-outbound-summary-stats-view"},{default:o(({route:e})=>[t(p,{render:!1,title:"Stats"}),s(),t(u,null,{default:o(()=>[t(_,{src:`/meshes/${e.params.mesh}/dataplanes/${e.params.dataPlane}/stats/${i.dataplaneOverview.dataplane.networking.inboundAddress}`},{default:o(({data:f,refresh:g})=>[t(m,{items:f.raw.split(`
`),predicate:c=>c.includes(`.${e.params.connection}.`)},{default:o(({items:c})=>[t(l,{language:"json",code:c.map(a=>a.replace(`${e.params.connection}.`,"")).join(`
`),"is-searchable":"",query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter,"is-reg-exp-mode":e.params.codeRegExp,onQueryChange:a=>e.update({codeSearch:a}),onFilterModeChange:a=>e.update({codeFilter:a}),onRegExpModeChange:a=>e.update({codeRegExp:a})},{"primary-actions":o(()=>[t(d,{action:"refresh",appearance:"primary",onClick:g},{default:o(()=>[s(`
                Refresh
              `)]),_:2},1032,["onClick"])]),_:2},1032,["code","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])]),_:2},1032,["items","predicate"])]),_:2},1032,["src"])]),_:2},1024)]),_:1})}}});export{V as default};