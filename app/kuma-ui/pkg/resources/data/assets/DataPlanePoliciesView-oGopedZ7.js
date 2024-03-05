import{A as V,a as A}from"./AccordionList-dEVBZa9v.js";import{d as F,a as k,o as e,c as s,m as l,f as a,F as u,H as B,t as o,e as d,w as t,p as b,b as i,X as D,x as j,y as G,_ as M,M as Y,q as w,l as z,n as J,a2 as Q}from"./index-1j9z4Egf.js";import{C as H}from"./CodeBlock-ejECcgv-.js";import{P as q}from"./PolicyTypeTag-gz0KMh93.js";import{T as x}from"./TagList-X1LjfAG5.js";import{t as X}from"./toYaml-sPaYOD3i.js";const C=y=>(j("data-v-f3e7afbb"),y=y(),G(),y),U={class:"policies-list"},W={class:"mesh-gateway-policy-list"},Z=C(()=>l("h3",{class:"mb-2"},`
        Gateway policies
      `,-1)),E={key:0},ee=C(()=>l("h3",{class:"mt-6 mb-2"},`
        Listeners
      `,-1)),te=C(()=>l("b",null,"Host",-1)),ae=C(()=>l("h4",{class:"mt-2 mb-2"},`
                Routes
              `,-1)),se={class:"dataplane-policy-header"},ne=C(()=>l("b",null,"Route",-1)),le=C(()=>l("b",null,"Service",-1)),oe={key:0,class:"badge-list"},ie={class:"mt-1"},ce=F({__name:"BuiltinGatewayPolicies",props:{gatewayDataplane:{},policyTypesByName:{}},setup(y){const m=y;return(c,N)=>{const f=k("RouterLink"),g=k("KBadge");return e(),s("div",U,[l("div",W,[Z,a(),c.gatewayDataplane.routePolicies.length>0?(e(),s("ul",E,[(e(!0),s(u,null,B(c.gatewayDataplane.routePolicies,(r,h)=>(e(),s("li",{key:h},[l("span",null,o(r.type),1),a(`:

          `),d(f,{to:{name:"policy-detail-view",params:{mesh:r.mesh,policyPath:m.policyTypesByName[r.type].path,policy:r.name}}},{default:t(()=>[a(o(r.name),1)]),_:2},1032,["to"])]))),128))])):b("",!0),a(),ee,a(),l("div",null,[(e(!0),s(u,null,B(c.gatewayDataplane.listenerEntries,(r,h)=>(e(),s("div",{key:h},[l("div",null,[l("div",null,[te,a(": "+o(r.hostName)+":"+o(r.port)+" ("+o(r.protocol)+`)
            `,1)]),a(),r.routeEntries.length>0?(e(),s(u,{key:0},[ae,a(),d(A,{"initially-open":[],"multiple-open":""},{default:t(()=>[(e(!0),s(u,null,B(r.routeEntries,(p,T)=>(e(),i(V,{key:T},D({"accordion-header":t(()=>[l("div",se,[l("div",null,[l("div",null,[ne,a(": "),d(f,{to:{name:"policy-detail-view",params:{mesh:p.route.mesh,policyPath:m.policyTypesByName[p.route.type].path,policy:p.route.name}}},{default:t(()=>[a(o(p.route.name),1)]),_:2},1032,["to"])]),a(),l("div",null,[le,a(": "+o(p.service),1)])]),a(),p.origins.length>0?(e(),s("div",oe,[(e(!0),s(u,null,B(p.origins,(n,_)=>(e(),i(g,{key:`${h}-${_}`},{default:t(()=>[a(o(n.type),1)]),_:2},1024))),128))])):b("",!0)])]),_:2},[p.origins.length>0?{name:"accordion-content",fn:t(()=>[l("ul",ie,[(e(!0),s(u,null,B(p.origins,(n,_)=>(e(),s("li",{key:`${h}-${_}`},[a(o(n.type)+`:

                        `,1),d(f,{to:{name:"policy-detail-view",params:{mesh:n.mesh,policyPath:m.policyTypesByName[n.type].path,policy:n.name}}},{default:t(()=>[a(o(n.name),1)]),_:2},1032,["to"])]))),128))])]),key:"0"}:void 0]),1024))),128))]),_:2},1024)],64)):b("",!0)])]))),128))])])])}}}),re=M(ce,[["__scopeId","data-v-f3e7afbb"]]),pe={class:"policy-type-heading"},ue={class:"policy-list"},de={key:0},me=F({__name:"PolicyTypeEntryList",props:{policyTypeEntries:{},policyTypesByName:{}},setup(y){const m=y,c=Y(()=>m.policyTypeEntries.filter(f=>{var g;return((g=m.policyTypesByName[f.type])==null?void 0:g.isTargetRefBased)===!1}));function N({headerKey:f}){return{class:`cell-${f}`}}return(f,g)=>{const r=k("RouterLink"),h=k("KTable");return e(),i(A,{"initially-open":[],"multiple-open":""},{default:t(()=>[(e(!0),s(u,null,B(c.value,(p,T)=>(e(),i(V,{key:T},{"accordion-header":t(()=>[l("h3",pe,[d(q,{"policy-type":p.type},{default:t(()=>[a(o(p.type)+" ("+o(p.connections.length)+`)
          `,1)]),_:2},1032,["policy-type"])])]),"accordion-content":t(()=>[l("div",ue,[d(h,{class:"policy-type-table",fetcher:()=>({data:p.connections,total:p.connections.length}),headers:[{label:"From",key:"sourceTags"},{label:"To",key:"destinationTags"},{label:"On",key:"name"},{label:"Conf",key:"config"},{label:"Origin policies",key:"origins"}],"cell-attrs":N,"disable-pagination":"","is-clickable":""},{sourceTags:t(({row:n})=>[n.sourceTags.length>0?(e(),i(x,{key:0,class:"tag-list","should-truncate":"",tags:n.sourceTags},null,8,["tags"])):(e(),s(u,{key:1},[a(`
                —
              `)],64))]),destinationTags:t(({row:n})=>[n.destinationTags.length>0?(e(),i(x,{key:0,class:"tag-list","should-truncate":"",tags:n.destinationTags},null,8,["tags"])):(e(),s(u,{key:1},[a(`
                —
              `)],64))]),name:t(({row:n})=>[n.name!==null?(e(),s(u,{key:0},[a(o(n.name),1)],64)):(e(),s(u,{key:1},[a(`
                —
              `)],64))]),origins:t(({row:n})=>[n.origins.length>0?(e(),s("ul",de,[(e(!0),s(u,null,B(n.origins,(_,$)=>(e(),s("li",{key:`${T}-${$}`},[d(r,{to:{name:"policy-detail-view",params:{mesh:_.mesh,policyPath:m.policyTypesByName[_.type].path,policy:_.name}}},{default:t(()=>[a(o(_.name),1)]),_:2},1032,["to"])]))),128))])):(e(),s(u,{key:1},[a(`
                —
              `)],64))]),config:t(({row:n})=>[n.config?(e(),i(H,{key:0,code:w(X)(n.config),language:"yaml","show-copy-button":!1},null,8,["code"])):(e(),s(u,{key:1},[a(`
                —
              `)],64))]),_:2},1032,["fetcher"])])]),_:2},1024))),128))]),_:1})}}}),ye=M(me,[["__scopeId","data-v-18983fd0"]]),_e=y=>(j("data-v-c78ed2e9"),y=y(),G(),y),he={class:"policy-type-heading"},fe={class:"policy-list"},ge={key:0,class:"matcher"},ke={key:0,class:"matcher__and"},be=_e(()=>l("br",null,null,-1)),$e={key:1,class:"matcher__not"},ve={class:"matcher__term"},Re={key:1},Be={key:0},Te=F({__name:"RuleEntryList",props:{ruleEntries:{},policyTypesByName:{},showMatchers:{type:Boolean,default:!0}},setup(y){const{t:m}=z(),c=y;function N({headerKey:f}){return{class:`cell-${f}`}}return(f,g)=>{const r=k("RouterLink"),h=k("KTable");return e(),i(A,{"initially-open":[],"multiple-open":""},{default:t(()=>[(e(!0),s(u,null,B(c.ruleEntries,(p,T)=>(e(),i(V,{key:T},{"accordion-header":t(()=>[l("h3",he,[d(q,{"policy-type":p.type},{default:t(()=>[a(o(p.type),1)]),_:2},1032,["policy-type"])])]),"accordion-content":t(()=>[l("div",fe,[d(h,{class:J(["policy-type-table",{"has-matchers":c.showMatchers}]),fetcher:()=>({data:p.rules,total:p.rules.length}),headers:[...c.showMatchers?[{label:"Matchers",key:"matchers"}]:[],{label:"Origin policies",key:"origins"},{label:"Conf",key:"config"}],"cell-attrs":N,"disable-pagination":""},D({origins:t(({row:n})=>[n.origins.length>0?(e(),s("ul",Be,[(e(!0),s(u,null,B(n.origins,(_,$)=>(e(),s("li",{key:`${T}-${$}`},[d(r,{to:{name:"policy-detail-view",params:{mesh:_.mesh,policyPath:c.policyTypesByName[_.type].path,policy:_.name}}},{default:t(()=>[a(o(_.name),1)]),_:2},1032,["to"])]))),128))])):(e(),s(u,{key:1},[a(o(w(m)("common.collection.none")),1)],64))]),config:t(({row:n})=>[n.config?(e(),i(H,{key:0,code:w(X)(n.config),language:"yaml","show-copy-button":!1},null,8,["code"])):(e(),s(u,{key:1},[a(o(w(m)("common.collection.none")),1)],64))]),_:2},[c.showMatchers?{name:"matchers",fn:t(({row:n})=>[n.matchers&&n.matchers.length>0?(e(),s("span",ge,[(e(!0),s(u,null,B(n.matchers,({key:_,value:$,not:I},v)=>(e(),s(u,{key:v},[v>0?(e(),s("span",ke,[a(" and"),be])):b("",!0),I?(e(),s("span",$e,"!")):b("",!0),l("span",ve,o(`${_}:${$}`),1)],64))),128))])):(e(),s("i",Re,o(w(m)("data-planes.routes.item.matches_everything")),1))]),key:"0"}:void 0]),1032,["class","fetcher","headers"])])]),_:2},1024))),128))]),_:1})}}}),S=M(Te,[["__scopeId","data-v-c78ed2e9"]]),we={class:"stack"},Pe={class:"mb-2"},Ne=F({__name:"StandardDataplanePolicies",props:{inspectRulesForDataplane:{},policyTypesByName:{}},setup(y){const{t:m}=z(),c=y;return(N,f)=>{const g=k("KCard");return e(),s("div",we,[c.inspectRulesForDataplane.proxyRules.length>0?(e(),i(g,{key:0},{default:t(()=>[l("h3",null,o(w(m)("data-planes.routes.item.proxy_rule")),1),a(),d(S,{class:"mt-2","rule-entries":c.inspectRulesForDataplane.proxyRules,"policy-types-by-name":c.policyTypesByName,"show-matchers":!1,"data-testid":"proxy-rule-list"},null,8,["rule-entries","policy-types-by-name"])]),_:1})):b("",!0),a(),c.inspectRulesForDataplane.toRules.length>0?(e(),i(g,{key:1},{default:t(()=>[l("h3",null,o(w(m)("data-planes.routes.item.to_rules")),1),a(),d(S,{class:"mt-2","rule-entries":c.inspectRulesForDataplane.toRules,"policy-types-by-name":c.policyTypesByName,"data-testid":"to-rule-list"},null,8,["rule-entries","policy-types-by-name"])]),_:1})):b("",!0),a(),c.inspectRulesForDataplane.fromRuleInbounds.length>0?(e(),i(g,{key:2},{default:t(()=>[l("h3",Pe,o(w(m)("data-planes.routes.item.from_rules")),1),a(),(e(!0),s(u,null,B(c.inspectRulesForDataplane.fromRuleInbounds,(r,h)=>(e(),s("div",{key:h},[l("h4",null,o(w(m)("data-planes.routes.item.port",{port:r.port})),1),a(),d(S,{class:"mt-2","rule-entries":r.ruleEntries,"policy-types-by-name":c.policyTypesByName,"data-testid":`from-rule-list-${h}`},null,8,["rule-entries","policy-types-by-name","data-testid"])]))),128))]),_:1})):b("",!0)])}}}),Ce={class:"stack"},Le={key:0},Me=F({__name:"DataPlanePoliciesView",props:{data:{}},setup(y){const m=y;return(c,N)=>{const f=k("RouteTitle"),g=k("DataCollection"),r=k("DataLoader"),h=k("KCard"),p=k("DataSource"),T=k("AppView"),n=k("RouteView");return e(),i(n,{name:"data-plane-policies-view",params:{mesh:"",dataPlane:""}},{default:t(({can:_,route:$,t:I})=>[d(T,null,{title:t(()=>[l("h2",null,[d(f,{title:I("data-planes.routes.item.navigation.data-plane-policies-view")},null,8,["title"])])]),default:t(()=>[a(),l("div",Ce,[d(p,{src:"/*/policy-types"},{default:t(({data:v,error:K})=>[d(r,{src:`/meshes/${$.params.mesh}/dataplanes/${$.params.dataPlane}/rules`,data:[v],errors:[K]},{default:t(({data:R})=>[R&&v?(e(),i(g,{key:0,items:R.rules},{default:t(()=>[d(Ne,{"policy-types-by-name":v.policies.reduce((L,P)=>Object.assign(L,{[P.name]:P}),{}),"inspect-rules-for-dataplane":R,"data-testid":"rules-based-policies"},null,8,["policy-types-by-name","inspect-rules-for-dataplane"])]),_:2},1032,["items"])):b("",!0)]),_:2},1032,["src","data","errors"]),a(),_("use zones")?b("",!0):(e(),s("div",Le,[l("h3",null,o(I("data-planes.routes.item.legacy_policies")),1),a(),m.data.dataplaneType==="builtin"?(e(),i(r,{key:0,src:`/meshes/${$.params.mesh}/dataplanes/${$.params.dataPlane}/gateway-dataplane-policies`,data:[v],errors:[K]},{default:t(({data:R})=>[R?(e(),i(g,{key:0,items:R.routePolicies},{default:t(()=>[R.listenerEntries.length===0?(e(),i(Q,{key:0})):(e(),i(h,{key:1,class:"mt-4"},{default:t(()=>[v?(e(),i(re,{key:0,"policy-types-by-name":v.policies.reduce((L,P)=>Object.assign(L,{[P.name]:P}),{}),"gateway-dataplane":R,"data-testid":"builtin-gateway-dataplane-policies"},null,8,["policy-types-by-name","gateway-dataplane"])):b("",!0)]),_:2},1024))]),_:2},1032,["items"])):b("",!0)]),_:2},1032,["src","data","errors"])):(e(),i(r,{key:1,src:`/meshes/${$.params.mesh}/dataplanes/${$.params.dataPlane}/sidecar-dataplane-policies`,data:[v],errors:[K]},{default:t(({data:R})=>[R?(e(),i(g,{key:0,items:R.policyTypeEntries},{default:t(({items:L})=>[d(h,{class:"mt-4"},{default:t(()=>[v?(e(),i(ye,{key:0,"policy-type-entries":L,"policy-types-by-name":v.policies.reduce((P,O)=>Object.assign(P,{[O.name]:O}),{}),"data-testid":"sidecar-dataplane-policies"},null,8,["policy-type-entries","policy-types-by-name"])):b("",!0)]),_:2},1024)]),_:2},1032,["items"])):b("",!0)]),_:2},1032,["src","data","errors"]))]))]),_:2},1024)])]),_:2},1024)]),_:1})}}});export{Me as default};
