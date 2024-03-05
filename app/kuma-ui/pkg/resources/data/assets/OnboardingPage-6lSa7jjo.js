import{d as _,S as f,o as s,c as d,m as o,r as i,f as t,q as r,p,_ as u,b,w as c,V as l,e as h,t as v,n as m,x as S,y}from"./index-1j9z4Egf.js";const k={class:"onboarding-heading"},$={class:"onboarding-title","data-testid":"onboarding-header"},x={key:0,class:"onboarding-description"},w=_({__name:"OnboardingHeading",setup(a){const e=f();return(n,g)=>(s(),d("div",k,[o("h1",$,[i(n.$slots,"title",{},void 0,!0)]),t(),r(e).description?(s(),d("div",x,[i(n.$slots,"description",{},void 0,!0)])):p("",!0)]))}}),z=u(w,[["__scopeId","data-v-505a1a6e"]]),B={class:"onboarding-actions"},N={class:"button-list"},q=_({__name:"OnboardingNavigation",props:{shouldAllowNext:{type:Boolean,required:!1,default:!0},showSkip:{type:Boolean,required:!1,default:!0},nextStep:{type:String,required:!0},previousStep:{type:String,required:!1,default:""},nextStepTitle:{type:String,required:!1,default:"Next"},lastStep:{type:Boolean,required:!1,default:!1}},setup(a){const e=a;return(n,g)=>(s(),d("div",B,[e.previousStep?(s(),b(r(l),{key:0,appearance:"secondary",to:{name:e.previousStep},"data-testid":"onboarding-previous-button"},{default:c(()=>[t(`
      Back
    `)]),_:1},8,["to"])):p("",!0),t(),o("div",N,[e.showSkip?(s(),b(r(l),{key:0,appearance:"tertiary","data-testid":"onboarding-skip-button",to:{name:"home"}},{default:c(()=>[t(`
        Skip setup
      `)]),_:1})):p("",!0),t(),h(r(l),{disabled:!e.shouldAllowNext,appearance:"primary",to:{name:e.lastStep?"home":e.nextStep},"data-testid":"onboarding-next-button"},{default:c(()=>[t(v(e.nextStepTitle),1)]),_:1},8,["disabled","to"])])]))}}),D=u(q,[["__scopeId","data-v-4695c7f4"]]),I=a=>(S("data-v-41beef0f"),a=a(),y(),a),O={class:"onboarding-container"},V={class:"onboarding-container__header"},C={class:"onboarding-container__inner-content"},T={class:"mt-4"},A=I(()=>o("div",{class:"background-image"},null,-1)),H=_({__name:"OnboardingPage",props:{withImage:{type:Boolean,required:!1,default:!1}},setup(a){const e=a;return(n,g)=>(s(),d("div",null,[o("div",O,[o("div",V,[i(n.$slots,"header",{},void 0,!0)]),t(),o("div",{class:m(["onboarding-container__content",{"onboarding-container__content--with-image":e.withImage}])},[o("div",C,[i(n.$slots,"content",{},void 0,!0)])],2),t(),o("div",T,[i(n.$slots,"navigation",{},void 0,!0)])]),t(),A]))}}),E=u(H,[["__scopeId","data-v-41beef0f"]]);export{E as O,z as a,D as b};
