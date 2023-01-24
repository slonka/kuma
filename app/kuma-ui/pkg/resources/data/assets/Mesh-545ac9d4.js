import{m as V}from"./vuex.esm-bundler-df5bd11e.js";import{a as I,P as R,D as x}from"./kongponents.es-3df60cd6.js";import{u as O}from"./index-f6b877ca.js";import{k as P}from"./kumaApi-db784568.js";import{k as U}from"./helpers-32595d9f.js";import{_ as F}from"./CodeBlock.vue_vue_type_style_index_0_lang-98064716.js";import{f as K}from"./formatForCLI-4055776e.js";import{F as z,S as Y,E as j}from"./EntityScanner-9b947b61.js";import{T as W}from"./TabsWidget-aec0fabd.js";import{C as D}from"./ClientStorage-efe299d9.js";import{P as G}from"./constants-31fdaf55.js";import{l as v,h as N,g as t,e as p,w as l,o as u,t as C,f as n,a as h,b as g,S as c,a1 as f,a4 as E,a2 as A,p as Z,m as q}from"./runtime-dom.esm-bundler-91b41870.js";import{_ as H}from"./_plugin-vue_export-helper-c27b6911.js";import"./store-866e3b85.js";import"./dataplane-4aecf58f.js";import"./_commonjsHelpers-87174ba5.js";import"./index-a8834e9c.js";import"./datadogLogEvents-4578cfa7.js";import"./QueryParameter-70743f73.js";import"./ErrorBlock-46fedade.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-d3176fee.js";const Q={mtls:{enabledBackend:null,backends:[]},tracing:{defaultBackend:null,backends:[{name:null,type:null}]},logging:{backends:[{name:null,format:'{ "destination": "%KUMA_DESTINATION_SERVICE%", "destinationAddress": "%UPSTREAM_LOCAL_ADDRESS%", "source": "%KUMA_SOURCE_SERVICE%", "sourceAddress": "%KUMA_SOURCE_ADDRESS%", "bytesReceived": "%BYTES_RECEIVED%", "bytesSent": "%BYTES_SENT%"}',type:null}]},metrics:{enabledBackend:null,backends:[{name:null,type:null}]}};const J=O();function B(){return{meshName:"",meshCAName:"",meshLoggingBackend:"",meshTracingBackend:"",meshMetricsName:"",meshTracingZipkinURL:"",mtlsEnabled:"disabled",meshCA:"builtin",loggingEnabled:"disabled",loggingType:"tcp",meshLoggingPath:"/",meshLoggingAddress:"127.0.0.1:5000",meshLoggingBackendFormat:'{ start_time: "%START_TIME%", source: "%KUMA_SOURCE_SERVICE%", destination: "%KUMA_DESTINATION_SERVICE%", source_address: "%KUMA_SOURCE_ADDRESS_WITHOUT_PORT%", destination_address: "%UPSTREAM_HOST%", duration_millis: "%DURATION%", bytes_received: "%BYTES_RECEIVED%", bytes_sent: "%BYTES_SENT%" }',tracingEnabled:"disabled",meshTracingType:"zipkin",meshTracingSampling:99.9,metricsEnabled:"disabled",meshMetricsType:"prometheus",meshMetricsDataplanePort:5670,meshMetricsDataplanePath:"/metrics"}}function X(i,a){return Object.keys(i).filter(o=>!a.includes(o)).map(o=>Object.assign({},{[o]:i[o]})).reduce((o,T)=>Object.assign(o,T),{})}const $={name:"MeshWizard",components:{CodeBlock:F,FormFragment:z,TabsWidget:W,StepSkeleton:Y,EntityScanner:j,KAlert:I,KButton:R,KCard:x},data(){return{env:J,hasStoredMeshData:!1,productName:G,selectedTab:"",schema:Q,steps:[{label:"General & Security",slug:"general"},{label:"Logging",slug:"logging"},{label:"Tracing",slug:"tracing"},{label:"Metrics",slug:"metrics"},{label:"Install",slug:"complete"}],tabs:[{hash:"#kubernetes",title:"Kubernetes"},{hash:"#universal",title:"Universal"}],sidebarContent:[{name:"mesh"},{name:"did-you-know"}],formConditions:{mtlsEnabled:!1,loggingEnabled:!1,tracingEnabled:!1,metricsEnabled:!1,loggingType:null},startScanner:!1,scanFound:!1,hideScannerSiblings:!1,scanError:!1,isComplete:!1,validate:B(),vmsg:[]}},computed:{...V({title:"config/getTagline",environment:"config/getEnvironment"}),codeOutput(){const i=this.schema,a=Object.assign({},i),o=this.validate;if(!o)return;const T=o.mtlsEnabled==="enabled",e=o.loggingEnabled==="enabled",b=o.tracingEnabled==="enabled",S=o.metricsEnabled==="enabled",M={mtls:T,logging:e,tracing:b,metrics:S},m=[];if(Object.entries(M).forEach(d=>{const _=d[1],s=d[0];_?m.filter(L=>L!==s):m.push(s)}),T){a.mtls.enabled=!0;const d=a.mtls,_=this.validate.meshCA,s=this.validate.meshCAName;d.backends=[],d.enabledBackend=s,_==="provided"?d.backends=[{name:s,type:_,conf:{cert:{secret:""},key:{secret:""}}}]:d.backends=[{name:s,type:_}]}if(e){const d=a.logging.backends[0],_=d.format;d.conf={},d.name=o.meshLoggingBackend,d.type=o.loggingType,d.format=o.meshLoggingBackendFormat||_,o.loggingType==="tcp"?d.conf.address=o.meshLoggingAddress||"127.0.0.1:5000":o.loggingType==="file"&&(d.conf.path=o.meshLoggingPath)}if(b){const d=a.tracing;d.defaultBackend=o.meshTracingBackend,d.backends[0].type=o.meshTracingType||"zipkin",d.backends[0].name=o.meshTracingBackend,d.backends[0].sampling=o.meshTracingSampling||100,d.backends[0].conf={},d.backends[0].conf.url=o.meshTracingZipkinURL}if(S){const d=a.metrics;d.backends[0].conf={},d.enabledBackend=o.meshMetricsName,d.backends[0].type=o.meshMetricsType||"prometheus",d.backends[0].name=o.meshMetricsName,d.backends[0].conf.port=o.meshMetricsDataplanePort||5670,d.backends[0].conf.path=o.meshMetricsDataplanePath||"/metrics"}const k=X(a,m);let y,w;return this.selectedTab==="#kubernetes"?(w="kubectl",y={apiVersion:"kuma.io/v1alpha1",kind:"Mesh",metadata:{name:o.meshName}},Object.keys(k).length>0&&(y.spec=k)):(w="kumactl",y={type:"Mesh",name:o.meshName,...k}),K(y,`" | ${w} apply -f -`)},nextDisabled(){const{meshName:i,meshCAName:a,meshLoggingBackend:o,meshTracingBackend:T,meshTracingZipkinURL:e,meshMetricsName:b,mtlsEnabled:S,loggingEnabled:M,tracingEnabled:m,metricsEnabled:k,meshLoggingPath:y,loggingType:w}=this.validate;return!i.length||S==="enabled"&&!a?!0:this.$route.query.step==="1"?M==="disabled"?!1:o?w==="file"&&!y:!0:this.$route.query.step==="2"?m==="enabled"&&!(T&&e):this.$route.query.step==="3"?k==="enabled"&&!b:!1}},watch:{"validate.meshName"(i){const a=U(i);this.validate.meshName=a,this.validateMeshName(a)},"validate.meshCAName"(i){this.validate.meshCAName=U(i)},"validate.meshLoggingBackend"(i){this.validate.meshLoggingBackend=U(i)},"validate.meshTracingBackend"(i){this.validate.meshTracingBackend=U(i)},"validate.meshMetricsName"(i){this.validate.meshMetricsName=U(i)}},created(){const i=D.get("createMeshData");i!==null&&(this.validate=i,this.hasStoredMeshData=!0)},methods:{updateStoredData(){D.set("createMeshData",this.validate),this.hasStoredMeshData=!0},resetMeshData(){D.remove("createMeshData"),this.hasStoredMeshData=!1,this.validate=B()},onTabChange(i){this.selectedTab=i},hideSiblings(){this.hideScannerSiblings=!0},validateMeshName(i){!i||i===""?this.vmsg.meshName="A Mesh name is required to proceed":this.vmsg.meshName=""},scanForEntity(){const i=this.validate.meshName;this.scanComplete=!1,this.scanError=!1,i&&P.getMesh({name:i}).then(a=>{a&&a.name.length>0?(this.isRunning=!0,this.scanFound=!0):this.scanError=!0}).catch(a=>{this.scanError=!0,console.error(a)}).finally(()=>{this.scanComplete=!0})}}},r=i=>(Z("data-v-cea7385a"),i=i(),q(),i),ee={class:"wizard"},te={class:"wizard__content"},ae=r(()=>t("code",null,"kubectl",-1)),ne=r(()=>t("code",null,"kumactl",-1)),se=r(()=>t("h3",null,`
            To get started, please fill in the following information:
          `,-1)),le={class:"k-input-label mx-2"},ie=r(()=>t("span",null,"Disabled",-1)),oe={class:"k-input-label mx-2"},re=r(()=>t("span",null,"Enabled",-1)),de=r(()=>t("option",{value:"builtin"},`
                    builtin
                  `,-1)),me=r(()=>t("option",{value:"provided"},`
                    provided
                  `,-1)),ce=r(()=>t("p",{class:"help"},`
                  If you've enabled mTLS, you must select a CA.
                `,-1)),ue=r(()=>t("h3",null,`
            Setup Logging
          `,-1)),pe=r(()=>t("p",null,`
            You can setup as many logging backends as you need that you can later use to log traffic via the “TrafficLog” policy. In this wizard, we allow you to configure one backend, but you can add more manually if you wish.
          `,-1)),ge={class:"k-input-label mx-2"},he=r(()=>t("span",null,"Disabled",-1)),be={class:"k-input-label mx-2"},fe=r(()=>t("span",null,"Enabled",-1)),ke={key:1},ye=r(()=>t("option",{value:"tcp"},`
                      TCP
                    `,-1)),_e=r(()=>t("option",{value:"file"},`
                      File
                    `,-1)),ve=r(()=>t("h3",null,`
            Setup Tracing
          `,-1)),Ee=r(()=>t("p",null,`
            You can setup as many tracing backends as you need that you can later use to log traffic via the “TrafficTrace” policy. In this wizard we allow you to configure one backend, but you can add more manually as you wish.
          `,-1)),Te={class:"k-input-label mx-2"},Se=r(()=>t("span",null,"Disabled",-1)),Me={class:"k-input-label mx-2"},we=r(()=>t("span",null,"Enabled",-1)),Ce=r(()=>t("option",{value:"zipkin"},`
                    Zipkin
                  `,-1)),Ue=[Ce],Ne=r(()=>t("h3",null,`
            Setup Metrics
          `,-1)),Ae=r(()=>t("p",null,`
            You can expose metrics from every data-plane on a configurable path
            and port that a metrics service, like Prometheus, can use to fetch them.
          `,-1)),De={class:"k-input-label mx-2"},Be=r(()=>t("span",null,"Disabled",-1)),Le={class:"k-input-label mx-2"},Ve=r(()=>t("span",null,"Enabled",-1)),Ie=r(()=>t("option",{value:"prometheus"},`
                    Prometheus
                  `,-1)),Re=[Ie],xe={key:0},Oe={key:0},Pe=r(()=>t("h3",null,`
                Install a new Mesh
              `,-1)),Fe=r(()=>t("h3",null,"Searching…",-1)),Ke=r(()=>t("p",null,"We are looking for your mesh.",-1)),ze=r(()=>t("h3",null,"Done!",-1)),Ye={key:0},je=r(()=>t("h3",null,"Mesh not found",-1)),We=r(()=>t("p",null,"We were unable to find your mesh.",-1)),Ge=r(()=>t("p",null,`
                You haven't filled any data out yet! Please return to the first
                step and fill out your information.
              `,-1)),Ze=r(()=>t("h3",null,"Mesh",-1)),qe=["href"],He=r(()=>t("h3",null,"Did You Know?",-1)),Qe=r(()=>t("p",null,`
            As you know, the GUI is read-only, but it will be providing instructions
            to create a new Mesh and verify everything worked well.
          `,-1));function Je(i,a,o,T,e,b){const S=v("KButton"),M=v("KAlert"),m=v("FormFragment"),k=v("KCard"),y=v("CodeBlock"),w=v("TabsWidget"),d=v("EntityScanner"),_=v("StepSkeleton");return u(),N("div",ee,[t("div",te,[p(_,{steps:e.steps,"sidebar-content":e.sidebarContent,"footer-enabled":e.hideScannerSiblings===!1,"next-disabled":b.nextDisabled,onGoToStep:b.updateStoredData},{general:l(()=>[t("p",null,`
            Welcome to the wizard for creating a new Mesh resource in `+C(e.productName)+`.
            We will be providing you with a few steps that will get you started.
          `,1),n(),t("p",null,[n(`
            As you know, the `+C(e.productName)+` GUI is read-only, so at the end of this wizard
            we will be generating the configuration that you can apply with either
            `,1),ae,n(` (if you are running in Kubernetes mode) or
            `),ne,n(` / API (if you are running in Universal mode).
          `)]),n(),se,n(),p(k,{class:"my-6",title:"Mesh Information","has-shadow":""},{body:l(()=>[e.hasStoredMeshData?(u(),h(M,{key:0,class:"reset-mesh-data-alert",appearance:"info"},{alertMessage:l(()=>[n(`
                  Want to start with an empty slate?
                `)]),actionButtons:l(()=>[p(S,{apperance:"outline",onClick:b.resetMeshData},{default:l(()=>[n(`
                    Reset to defaults
                  `)]),_:1},8,["onClick"])]),_:1})):g("",!0),n(),p(m,{class:"mt-4",title:"Mesh name","for-attr":"mesh-name"},{default:l(()=>[c(t("input",{id:"mesh-name","onUpdate:modelValue":a[0]||(a[0]=s=>e.validate.meshName=s),type:"text",class:"k-input w-100","data-testid":"mesh-name",placeholder:"your-mesh-name",required:""},null,512),[[f,e.validate.meshName]]),n(),e.vmsg.meshName?(u(),h(M,{key:0,appearance:"danger",size:"small","alert-message":e.vmsg.meshName},null,8,["alert-message"])):g("",!0)]),_:1}),n(),p(m,{class:"mt-4",title:"Mutual TLS"},{default:l(()=>[t("label",le,[c(t("input",{ref:"mtlsDisabled","onUpdate:modelValue":a[1]||(a[1]=s=>e.validate.mtlsEnabled=s),value:"disabled",name:"mtls",type:"radio",class:"k-input mr-2","data-testid":"mesh-mtls-disabled"},null,512),[[E,e.validate.mtlsEnabled]]),n(),ie]),n(),t("label",oe,[c(t("input",{id:"mtls-enabled","onUpdate:modelValue":a[2]||(a[2]=s=>e.validate.mtlsEnabled=s),value:"enabled",name:"mtls",type:"radio",class:"k-input mr-2","data-testid":"mesh-mtls-enabled"},null,512),[[E,e.validate.mtlsEnabled]]),n(),re])]),_:1}),n(),e.validate.mtlsEnabled==="enabled"?(u(),h(m,{key:1,class:"mt-4",title:"Certificate name","for-attr":"certificate-name"},{default:l(()=>[c(t("input",{id:"certificate-name","onUpdate:modelValue":a[3]||(a[3]=s=>e.validate.meshCAName=s),type:"text",class:"k-input w-100",placeholder:"your-certificate-name","data-testid":"mesh-certificate-name"},null,512),[[f,e.validate.meshCAName]])]),_:1})):g("",!0),n(),e.validate.mtlsEnabled==="enabled"?(u(),h(m,{key:2,class:"mt-4",title:"Certificate Authority","for-attr":"certificate-authority"},{default:l(()=>[c(t("select",{id:"certificate-authority","onUpdate:modelValue":a[4]||(a[4]=s=>e.validate.meshCA=s),class:"k-input w-100",name:"certificate-authority"},[de,n(),me],512),[[A,e.validate.meshCA]]),n(),ce]),_:1})):g("",!0)]),_:1})]),logging:l(()=>[ue,n(),pe,n(),p(k,{class:"my-6",title:"Logging Configuration","has-shadow":""},{body:l(()=>[p(m,{title:"Logging"},{default:l(()=>[t("label",ge,[c(t("input",{id:"logging-disabled","onUpdate:modelValue":a[5]||(a[5]=s=>e.validate.loggingEnabled=s),value:"disabled",name:"logging",type:"radio",class:"k-input mr-2","data-testid":"mesh-logging-disabled"},null,512),[[E,e.validate.loggingEnabled]]),n(),he]),n(),t("label",be,[c(t("input",{id:"logging-enabled","onUpdate:modelValue":a[6]||(a[6]=s=>e.validate.loggingEnabled=s),value:"enabled",name:"logging",type:"radio",class:"k-input mr-2","data-testid":"mesh-logging-enabled"},null,512),[[E,e.validate.loggingEnabled]]),n(),fe])]),_:1}),n(),e.validate.loggingEnabled==="enabled"?(u(),h(m,{key:0,class:"mt-4",title:"Backend name","for-attr":"backend-name"},{default:l(()=>[c(t("input",{id:"backend-name","onUpdate:modelValue":a[7]||(a[7]=s=>e.validate.meshLoggingBackend=s),type:"text",class:"k-input w-100",placeholder:"your-backend-name","data-testid":"mesh-logging-backend-name"},null,512),[[f,e.validate.meshLoggingBackend]])]),_:1})):g("",!0),n(),e.validate.loggingEnabled==="enabled"?(u(),N("div",ke,[p(m,{class:"mt-4",title:"Type"},{default:l(()=>[c(t("select",{id:"logging-type",ref:"loggingTypeSelect","onUpdate:modelValue":a[8]||(a[8]=s=>e.validate.loggingType=s),class:"k-input w-100",name:"logging-type"},[ye,n(),_e],512),[[A,e.validate.loggingType]])]),_:1}),n(),e.validate.loggingType==="file"?(u(),h(m,{key:0,class:"mt-4",title:"Path","for-attr":"backend-address"},{default:l(()=>[c(t("input",{id:"backend-address","onUpdate:modelValue":a[9]||(a[9]=s=>e.validate.meshLoggingPath=s),type:"text",class:"k-input w-100"},null,512),[[f,e.validate.meshLoggingPath]])]),_:1})):g("",!0),n(),e.validate.loggingType==="tcp"?(u(),h(m,{key:1,class:"mt-4",title:"Address","for-attr":"backend-address"},{default:l(()=>[c(t("input",{id:"backend-address","onUpdate:modelValue":a[10]||(a[10]=s=>e.validate.meshLoggingAddress=s),type:"text",class:"k-input w-100"},null,512),[[f,e.validate.meshLoggingAddress]])]),_:1})):g("",!0),n(),p(m,{class:"mt-4",title:"Format","for-attr":"backend-format"},{default:l(()=>[c(t("textarea",{id:"backend-format","onUpdate:modelValue":a[11]||(a[11]=s=>e.validate.meshLoggingBackendFormat=s),class:"k-input w-100 code-sample",rows:"12"},null,512),[[f,e.validate.meshLoggingBackendFormat]])]),_:1})])):g("",!0)]),_:1})]),tracing:l(()=>[ve,n(),Ee,n(),p(k,{class:"my-6",title:"Tracing Configuration","has-shadow":""},{body:l(()=>[p(m,{title:"Tracing"},{default:l(()=>[t("label",Te,[c(t("input",{id:"tracing-disabled","onUpdate:modelValue":a[12]||(a[12]=s=>e.validate.tracingEnabled=s),value:"disabled",name:"tracing",type:"radio",class:"k-input mr-2"},null,512),[[E,e.validate.tracingEnabled]]),n(),Se]),n(),t("label",Me,[c(t("input",{id:"tracing-enabled","onUpdate:modelValue":a[13]||(a[13]=s=>e.validate.tracingEnabled=s),value:"enabled",name:"tracing",type:"radio",class:"k-input mr-2","data-testid":"mesh-tracing-enabled"},null,512),[[E,e.validate.tracingEnabled]]),n(),we])]),_:1}),n(),e.validate.tracingEnabled==="enabled"?(u(),h(m,{key:0,class:"mt-4",title:"Backend name","for-attr":"tracing-backend-name"},{default:l(()=>[c(t("input",{id:"tracing-backend-name","onUpdate:modelValue":a[14]||(a[14]=s=>e.validate.meshTracingBackend=s),type:"text",class:"k-input w-100",placeholder:"your-tracing-backend-name","data-testid":"mesh-tracing-backend-name"},null,512),[[f,e.validate.meshTracingBackend]])]),_:1})):g("",!0),n(),e.validate.tracingEnabled==="enabled"?(u(),h(m,{key:1,class:"mt-4",title:"Type","for-attr":"tracing-type"},{default:l(()=>[c(t("select",{id:"tracing-type","onUpdate:modelValue":a[15]||(a[15]=s=>e.validate.meshTracingType=s),class:"k-input w-100",name:"tracing-type"},Ue,512),[[A,e.validate.meshTracingType]])]),_:1})):g("",!0),n(),e.validate.tracingEnabled==="enabled"?(u(),h(m,{key:2,class:"mt-4",title:"Sampling","for-attr":"tracing-sampling"},{default:l(()=>[c(t("input",{id:"tracing-sampling","onUpdate:modelValue":a[16]||(a[16]=s=>e.validate.meshTracingSampling=s),type:"number",class:"k-input w-100",step:"0.1",min:"0",max:"100"},null,512),[[f,e.validate.meshTracingSampling]])]),_:1})):g("",!0),n(),e.validate.tracingEnabled==="enabled"?(u(),h(m,{key:3,class:"mt-4",title:"URL","for-attr":"tracing-zipkin-url"},{default:l(()=>[c(t("input",{id:"tracing-zipkin-url","onUpdate:modelValue":a[17]||(a[17]=s=>e.validate.meshTracingZipkinURL=s),type:"text",class:"k-input w-100",placeholder:"http://zipkin.url:1234","data-testid":"mesh-tracing-url"},null,512),[[f,e.validate.meshTracingZipkinURL]])]),_:1})):g("",!0)]),_:1})]),metrics:l(()=>[Ne,n(),Ae,n(),p(k,{class:"my-6",title:"Metrics Configuration","has-shadow":""},{body:l(()=>[p(m,{title:"Metrics"},{default:l(()=>[t("label",De,[c(t("input",{id:"metrics-disabled","onUpdate:modelValue":a[18]||(a[18]=s=>e.validate.metricsEnabled=s),value:"disabled",name:"metrics",type:"radio",class:"k-input mr-2"},null,512),[[E,e.validate.metricsEnabled]]),n(),Be]),n(),t("label",Le,[c(t("input",{id:"metrics-enabled","onUpdate:modelValue":a[19]||(a[19]=s=>e.validate.metricsEnabled=s),value:"enabled",name:"metrics",type:"radio",class:"k-input mr-2","data-testid":"mesh-metrics-enabled"},null,512),[[E,e.validate.metricsEnabled]]),n(),Ve])]),_:1}),n(),e.validate.metricsEnabled==="enabled"?(u(),h(m,{key:0,class:"mt-4",title:"Backend name","for-attr":"metrics-name"},{default:l(()=>[c(t("input",{id:"metrics-name","onUpdate:modelValue":a[20]||(a[20]=s=>e.validate.meshMetricsName=s),type:"text",class:"k-input w-100",placeholder:"your-metrics-backend-name","data-testid":"mesh-metrics-backend-name"},null,512),[[f,e.validate.meshMetricsName]])]),_:1})):g("",!0),n(),e.validate.metricsEnabled==="enabled"?(u(),h(m,{key:1,class:"mt-4",title:"Type","for-attr":"metrics-type"},{default:l(()=>[c(t("select",{id:"metrics-type","onUpdate:modelValue":a[21]||(a[21]=s=>e.validate.meshMetricsType=s),class:"k-input w-100",name:"metrics-type"},Re,512),[[A,e.validate.meshMetricsType]])]),_:1})):g("",!0),n(),e.validate.metricsEnabled==="enabled"?(u(),h(m,{key:2,class:"mt-4",title:"Dataplane port","for-attr":"metrics-dataplane-port"},{default:l(()=>[c(t("input",{id:"metrics-dataplane-port","onUpdate:modelValue":a[22]||(a[22]=s=>e.validate.meshMetricsDataplanePort=s),type:"number",class:"k-input w-100",step:"1",min:"0",max:"65535",placeholder:"1234"},null,512),[[f,e.validate.meshMetricsDataplanePort]])]),_:1})):g("",!0),n(),e.validate.metricsEnabled==="enabled"?(u(),h(m,{key:3,class:"mt-4",title:"Dataplane path","for-attr":"metrics-dataplane-path"},{default:l(()=>[c(t("input",{id:"metrics-dataplane-path","onUpdate:modelValue":a[23]||(a[23]=s=>e.validate.meshMetricsDataplanePath=s),type:"text",class:"k-input w-100"},null,512),[[f,e.validate.meshMetricsDataplanePath]])]),_:1})):g("",!0)]),_:1})]),complete:l(()=>[b.codeOutput?(u(),N("div",xe,[e.hideScannerSiblings===!1?(u(),N("div",Oe,[Pe,n(),t("p",null,`
                Since the `+C(e.productName)+` GUI is read-only mode to follow Ops best practices,
                please execute the following command in your shell to create the entity.
                `+C(e.productName)+` will automatically detect when the new entity has been created.
              `,1),n(),p(w,{tabs:e.tabs,"initial-tab-override":i.environment,onOnTabChange:b.onTabChange},{kubernetes:l(()=>[p(y,{id:"code-block-kubernetes-command","data-testid":"kubernetes",language:"bash",code:b.codeOutput},null,8,["code"])]),universal:l(()=>[p(y,{id:"code-block-universal-command","data-testid":"universal",language:"bash",code:b.codeOutput},null,8,["code"])]),_:1},8,["tabs","initial-tab-override","onOnTabChange"])])):g("",!0),n(),p(d,{"loader-function":b.scanForEntity,"should-start":!0,"has-error":e.scanError,"can-complete":e.scanFound,onHideSiblings:b.hideSiblings},{"loading-title":l(()=>[Fe]),"loading-content":l(()=>[Ke]),"complete-title":l(()=>[ze]),"complete-content":l(()=>[t("p",null,[n(`
                  Your mesh `),e.validate.meshName?(u(),N("strong",Ye,C(e.validate.meshName),1)):g("",!0),n(` was found!
                `)]),n(),t("p",null,[p(S,{appearance:"primary",to:{name:"mesh-detail-view",params:{mesh:e.validate.meshName}}},{default:l(()=>[n(`
                    Go to mesh `+C(e.validate.meshName),1)]),_:1},8,["to"])])]),"error-title":l(()=>[je]),"error-content":l(()=>[We]),_:1},8,["loader-function","has-error","can-complete","onHideSiblings"])])):(u(),h(M,{key:1,appearance:"danger"},{alertMessage:l(()=>[Ge]),_:1}))]),mesh:l(()=>[Ze,n(),t("p",null,`
            In `+C(i.title)+`, a Mesh resource allows you to define an isolated environment
            for your data-planes and policies. It's isolated because the mTLS CA
            you choose can be different from the one configured for our Meshes.
            Ideally, you will have either a large Mesh with all the workloads, or
            one Mesh per application for better isolation.
          `,1),n(),t("p",null,[t("a",{href:`${e.env("KUMA_DOCS_URL")}/policies/mesh/?${e.env("KUMA_UTM_QUERY_PARAMS")}`,target:"_blank"},`
              Learn More
            `,8,qe)])]),"did-you-know":l(()=>[He,n(),Qe]),_:1},8,["steps","sidebar-content","footer-enabled","next-disabled","onGoToStep"])])])}const yt=H($,[["render",Je],["__scopeId","data-v-cea7385a"]]);export{yt as default};
