import{d as x,r as t,o as r,p as u,w as e,b as a,m as D,ao as R,c as v,V as T,q as C,e as c,R as k,K as A,t as B,F as S}from"./index-B5ZV0Dbw.js";const X={key:0},q=x({__name:"BuiltinGatewayDetailTabsView",setup(L){return(N,m)=>{const _=t("RouteTitle"),p=t("XAction"),d=t("XTabs"),w=t("RouterView"),b=t("DataLoader"),f=t("AppView"),h=t("DataSource"),y=t("RouteView");return r(),u(y,{name:"builtin-gateway-detail-tabs-view",params:{mesh:"",gateway:""}},{default:e(({route:s,t:i,uri:g})=>[a(h,{src:g(D(R),"/meshes/:mesh/mesh-gateways/:name",{mesh:s.params.mesh,name:s.params.gateway})},{default:e(({data:o,error:V})=>[a(f,{docs:i("builtin-gateways.href.docs"),breadcrumbs:[{to:{name:"mesh-detail-view",params:{mesh:s.params.mesh}},text:s.params.mesh},{to:{name:"builtin-gateway-list-view",params:{mesh:s.params.mesh}},text:i("builtin-gateways.routes.item.breadcrumbs")}]},{title:e(()=>[o?(r(),v("h1",X,[a(T,{text:o.name},{default:e(()=>[a(_,{title:i("builtin-gateways.routes.item.title",{name:o.name})},null,8,["title"])]),_:2},1032,["text"])])):C("",!0)]),default:e(()=>[m[1]||(m[1]=c()),a(b,{data:[o],errors:[V]},{default:e(()=>{var l;return[a(d,{selected:(l=s.child())==null?void 0:l.name},k({_:2},[A(s.children,({name:n})=>({name:`${n}-tab`,fn:e(()=>[a(p,{to:{name:n}},{default:e(()=>[c(B(i(`builtin-gateways.routes.item.navigation.${n}`)),1)]),_:2},1032,["to"])])}))]),1032,["selected"]),m[0]||(m[0]=c()),a(w,null,{default:e(({Component:n})=>[(r(),u(S(n),{gateway:o},null,8,["gateway"]))]),_:2},1024)]}),_:2},1032,["data","errors"])]),_:2},1032,["docs","breadcrumbs"])]),_:2},1032,["src"])]),_:1})}}});export{q as default};