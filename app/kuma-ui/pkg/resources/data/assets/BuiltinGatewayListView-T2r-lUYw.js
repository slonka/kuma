import{K as w}from"./index-FZCiQto1.js";import{d as b,a as l,o as n,b as m,w as a,t as p,f as c,e as s,F as z,c as h,Q as C,q as y,p as x,_ as V}from"./index-CqXGLTiP.js";import{A as v}from"./AppCollection-kQgAZcfN.js";import{E as B}from"./ErrorBlock-226oEngl.js";import{T as L}from"./TextWithCopyButton-3lrZWnHN.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-2iMUP-XH.js";import"./CopyButton-MfTze-YE.js";const N=b({__name:"BuiltinGatewayListView",setup(S){return(A,D)=>{const r=l("RouterLink"),d=l("KCard"),g=l("AppView"),_=l("DataSource"),f=l("RouteView");return n(),m(_,{src:"/me"},{default:a(({data:k})=>[k?(n(),m(f,{key:0,name:"builtin-gateway-list-view",params:{page:1,size:10,mesh:"",gateway:""}},{default:a(({route:t,t:i})=>[s(_,{src:`/meshes/${t.params.mesh}/mesh-gateways?page=${t.params.page}&size=${t.params.size}`},{default:a(({data:o,error:u})=>[s(g,null,{default:a(()=>[s(d,null,{default:a(()=>[u!==void 0?(n(),m(B,{key:0,error:u},null,8,["error"])):(n(),m(v,{key:1,class:"builtin-gateway-collection","data-testid":"builtin-gateway-collection","empty-state-message":i("common.emptyState.message",{type:"Built-in Gateways"}),"empty-state-cta-to":i("builtin-gateways.href.docs"),"empty-state-cta-text":i("common.documentation"),headers:[{label:"Name",key:"name"},{label:"Zone",key:"zone"},{label:"Details",key:"details",hideLabel:!0}],"page-number":t.params.page,"page-size":t.params.size,total:o==null?void 0:o.total,items:o==null?void 0:o.items,error:u,onChange:t.update},{name:a(({row:e})=>[s(L,{text:e.name},{default:a(()=>[s(r,{to:{name:"builtin-gateway-detail-view",params:{mesh:e.mesh,gateway:e.name},query:{page:t.params.page,size:t.params.size}}},{default:a(()=>[c(p(e.name),1)]),_:2},1032,["to"])]),_:2},1032,["text"])]),zone:a(({row:e})=>[e.labels&&e.labels["kuma.io/origin"]==="zone"&&e.labels["kuma.io/zone"]?(n(),m(r,{key:0,to:{name:"zone-cp-detail-view",params:{zone:e.labels["kuma.io/zone"]}}},{default:a(()=>[c(p(e.labels["kuma.io/zone"]),1)]),_:2},1032,["to"])):(n(),h(z,{key:1},[c(p(i("common.detail.none")),1)],64))]),details:a(({row:e})=>[s(r,{class:"details-link","data-testid":"details-link",to:{name:"builtin-gateway-detail-view",params:{mesh:e.mesh,gateway:e.name}}},{default:a(()=>[c(p(i("common.collection.details_link"))+" ",1),s(y(C),{display:"inline-block",decorative:"",size:y(w)},null,8,["size"])]),_:2},1032,["to"])]),_:2},1032,["empty-state-message","empty-state-cta-to","empty-state-cta-text","page-number","page-size","total","items","error","onChange"]))]),_:2},1024)]),_:2},1024)]),_:2},1032,["src"])]),_:1})):x("",!0)]),_:1})}}}),q=V(N,[["__scopeId","data-v-f83079e6"]]);export{q as default};