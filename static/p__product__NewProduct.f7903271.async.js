(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[202],{51650:function(re){re.exports={productDetailWrapper:"productDetailWrapper___3KAzI",detailWrap:"detailWrap___3vr7o",detailTitle:"detailTitle___A1sIC"}},39562:function(re,X,i){"use strict";i.r(X),i.d(X,{default:function(){return yt}});var C=i(88983),w=i(47933),K=i(63185),D=i(9676),ve=i(77576),G=i(12028),ot=i(57663),A=i(71577),ge=i(48736),R=i(27049),fe=i(58024),b=i(91894),me=i(36877),O=i(18480),xe=i(9715),m=i(71481),Ce=i(71194),J=i(48889),ye=i(43185),E=i(48215),v=i(22385),x=i(45777),d=i(86582),s=i(90636),_a=i(34792),L=i(48086),H=i(3182),g=i(2824),en=i(47673),S=i(60345),h=i(67294),dt=i(95916),ct=i(51042),Y=i(93695),Se=i(70213),pt=i(37266),ht=i(71879),Ze=i(77673),Fe=i(57865),vt=i(34778),gt=i(7145),je=i.n(gt),ft=i(51650),q=i.n(ft),we=function(ee){var te=["image/jpeg","image/jpg","image/png","image/bmp","image/webp"];if(!te.includes(ee.type))return L.ZP.error("\u56FE\u7247\u683C\u5F0F\u4EC5\u652F\u6301 jpeg\u3001jpg\u3001png\u3001bmp\u3001webp"),E.Z.LIST_IGNORE;var Z=ee.size/1024/1024<3;return Z?!0:(L.ZP.error("Image must smaller than 2MB!"),E.Z.LIST_IGNORE)},De=["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"],Q=i(69083),mt=i(81910),a=i(85893),xt=S.Z.TextArea,ue=10,_=20,ke="DragableUploadList",Ct=function(){var ee=(0,h.useState)([]),te=(0,g.Z)(ee,2),Z=te[0],ie=te[1],Zt=(0,h.useState)([]),Ie=(0,g.Z)(Zt,2),Ee=Ie[0],Le=Ie[1],Ft=(0,h.useState)(),Te=(0,g.Z)(Ft,2),k=Te[0],jt=Te[1],wt=(0,h.useState)(),Be=(0,g.Z)(wt,2),se=Be[0],Dt=Be[1],kt=(0,h.useState)(""),Ae=(0,g.Z)(kt,2),oe=Ae[0],It=Ae[1],Et=(0,h.useState)(""),Re=(0,g.Z)(Et,2),be=Re[0],Lt=Re[1],Tt=(0,h.useState)([]),Pe=(0,g.Z)(Tt,2),P=Pe[0],de=Pe[1],Bt=(0,h.useState)([{text:"",children:[{text:""}]}]),We=(0,g.Z)(Bt,2),f=We[0],U=We[1],At=(0,h.useState)([]),Ne=(0,g.Z)(At,2),I=Ne[0],V=Ne[1],Rt=(0,h.useState)(!1),Ge=(0,g.Z)(Rt,2),bt=Ge[0],Pt=Ge[1],Wt=(0,h.useState)(""),Nt=(0,g.Z)(Wt,1),Gt=Nt[0],Ot=(0,h.useState)(""),Ut=(0,g.Z)(Ot,1),Mt=Ut[0],$t=(0,h.useState)([{text:"",children:[{text:""}]}]),Oe=(0,g.Z)($t,2),o=Oe[0],M=Oe[1],zt=(0,h.useState)([]),Ue=(0,g.Z)(zt,2),F=Ue[0],ae=Ue[1],Kt=(0,h.useState)([{title:o[0].text||"\u54C1\u7C7B",key:"retailsaleType1"},{title:"\u5E93\u5B58",key:"stock"},{title:"\u96F6\u552E\u4EF7",key:"price"},{title:(0,a.jsxs)("div",{children:["\u72B6\u6001",(0,a.jsx)("div",{style:{fontSize:"12px",color:"#000"},children:"(\u542F\u7528\u540E\u672C\u884C\u9700\u5FC5\u586B)"})]}),key:"status"}]),Me=(0,g.Z)(Kt,2),T=Me[0],ne=Me[1],Jt=(0,h.useState)([]),$e=(0,g.Z)(Jt,2),ze=$e[0],Ke=$e[1],Ht=(0,h.useState)([{title:f[0].text||"\u54C1\u7C7B",dataIndex:"wholesaleType1",key:"wholesaleType1"},{title:"\u8BA1\u91CF\u5355\u4F4D",dataIndex:"unit",key:"unit"},{title:"\u8D77\u6279\u91CF",dataIndex:"package",key:"package"},{title:"\u4EF7\u683C",dataIndex:"price",key:"price"},{title:(0,a.jsxs)("div",{children:["\u72B6\u6001",(0,a.jsx)("div",{style:{fontSize:"12px",color:"#000"},children:"(\u542F\u7528\u540E\u672C\u884C\u9700\u5FC5\u586B)"})]}),dataIndex:"status",key:"status"}]),Je=(0,g.Z)(Ht,2),B=Je[0],le=Je[1],Qt=(0,h.useState)(["1","2"]),He=(0,g.Z)(Qt,2),ce=He[0],Vt=He[1],Xt=(0,h.useState)([]),Qe=(0,g.Z)(Xt,2),j=Qe[0],Yt=Qe[1],qt=(0,h.useState)([]),Ve=(0,g.Z)(qt,2),pe=Ve[0],_t=Ve[1],ea=(0,h.useState)("1"),Xe=(0,g.Z)(ea,2),Ye=Xe[0],ta=Xe[1],aa=(0,h.useState)(""),qe=(0,g.Z)(aa,2),he=qe[0],na=qe[1],la=function(){var l=(0,H.Z)((0,s.Z)().mark(function e(){var t;return(0,s.Z)().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return t={goodsDetail:{categoryName:se==null?void 0:se.join("/"),categoryAID:k==null?void 0:k[0],categoryBID:k==null?void 0:k[1],categoryCID:k==null?void 0:k[2],name:oe,description:be,images:Z.map(function(u){return u.response.data.link}),descImages:P.map(function(u){return u.response.data.link}),wholesaleShipping:pe==null?void 0:pe.join("/"),wholesaleAreaCodeA:j==null?void 0:j[0],wholesaleAreaCodeB:j==null?void 0:j[1],wholesaleAreaCodeC:j==null?void 0:j[2],retailShippingFee:o[0].text?he:void 0,retailShippingTime:o[0].text?Number(Ye):void 0,wholesaleLogistics:ce.map(Number)},wholesaleProducts:I.filter(function(u){return u.status}).map(function(u){return{price:u.price,minOrderQuantity:Number(u.package),status:u.status?1:0,unit:u.unit,productAttr:f.length===1?[{key:f[0].text,value:u.wholesaleType1}]:[{key:f[0].text,value:u.wholesaleType1},{key:f[1].text,value:u.wholesaleType2}]}}),retailProducts:o[0].text?F.filter(function(u){return u.status}).map(function(u){return{price:u.price,status:u.status?1:0,stock:Number(u.stock),productAttr:o.length===1?[{key:o[0].text,value:u.retailsaleType1}]:[{key:o[0].text,value:u.retailsaleType1},{key:o[1].text,value:u.retailsaleType2}]}}):void 0},r.next=3,(0,Q.JJ)(t);case 3:L.ZP.success("\u6DFB\u52A0\u6210\u529F"),mt.m8.push("/product/list");case 5:case"end":return r.stop()}},e)}));return function(){return l.apply(this,arguments)}}(),ra=function(){if(Z.length<=0||oe==="")return!0;var e=!1;if(f.forEach(function(r){r.text===""&&(e=!0),r.children.forEach(function(u){u.text===""&&(e=!0)})}),I.forEach(function(r){r.status&&(r.unit===""||r.price===""||r.package==="")&&(e=!0)}),o[0].text&&(o[0].children.forEach(function(r){r.text===""&&(e=!0)}),F[0].status===!0)){var t,n;(((t=F[0].children)===null||t===void 0?void 0:t.stock)===""||((n=F[0].children)===null||n===void 0?void 0:n.price)==="")&&(e=!0)}return!!(e||j.length===0||o[0].text&&he==="")},_e=function(){var e=!0;return o[0].text&&(e=!1,o[0].children.forEach(function(t){t.text===""&&(e=!0)}),F.find(function(t){return t.status})||(e=!0),F.forEach(function(t){t.status&&((t==null?void 0:t.stock)===""||(t==null?void 0:t.price)==="")&&(e=!0)})),e},et=function(e){return[{title:e[0].text||"\u54C1\u7C7B",dataIndex:"wholesaleType1",key:"wholesaleType1",width:150},{title:"\u8BA1\u91CF\u5355\u4F4D",dataIndex:"unit",key:"unit",width:150},{title:"\u8D77\u6279\u91CF",dataIndex:"package",key:"package",width:150},{title:"\u4EF7\u683C",dataIndex:"price",key:"price",width:150},{title:(0,a.jsxs)("div",{children:["\u72B6\u6001",(0,a.jsx)("div",{style:{fontSize:"12px",color:"#000"},children:"(\u542F\u7528\u540E\u672C\u884C\u9700\u5FC5\u586B)"})]}),dataIndex:"status",key:"status"}]},tt=function(e){var t;return[{title:e[0].text||"\u54C1\u7C7B",dataIndex:"wholesaleType1",key:"wholesaleType1",width:150},{title:(e==null||(t=e[1])===null||t===void 0?void 0:t.text)||"\u54C1\u7C7B",dataIndex:"wholesaleType2",key:"wholesaleType2",width:150},{title:"\u8BA1\u91CF\u5355\u4F4D",dataIndex:"unit",key:"unit",width:150},{title:"\u8D77\u6279\u91CF",dataIndex:"package",key:"package",width:150},{title:"\u4EF7\u683C",dataIndex:"price",key:"price",width:150},{title:(0,a.jsxs)("div",{children:["\u72B6\u6001",(0,a.jsx)("div",{style:{fontSize:"12px",color:"#000"},children:"(\u542F\u7528\u540E\u672C\u884C\u9700\u5FC5\u586B)"})]}),dataIndex:"status",key:"status"}]},at=function(e){return[{title:e[0].text||"\u54C1\u7C7B",key:"retailsaleType1"},{title:"\u5E93\u5B58",key:"stock"},{title:"\u96F6\u552E\u4EF7",key:"price"},{title:(0,a.jsxs)("div",{children:["\u72B6\u6001",(0,a.jsx)("div",{style:{fontSize:"12px",color:"#000"},children:"(\u542F\u7528\u540E\u672C\u884C\u9700\u5FC5\u586B)"})]}),key:"status"}]},nt=function(e){var t;return[{title:e[0].text||"\u54C1\u7C7B",key:"retailsaleType1"},{title:(e==null||(t=e[1])===null||t===void 0?void 0:t.text)||"\u54C1\u7C7B",key:"retailsaleType2"},{title:"\u5E93\u5B58",key:"stock"},{title:"\u96F6\u552E\u4EF7",key:"price"},{title:(0,a.jsxs)("div",{children:["\u72B6\u6001",(0,a.jsx)("div",{style:{fontSize:"12px",color:"#000"},children:"(\u542F\u7528\u540E\u672C\u884C\u9700\u5FC5\u586B)"})]}),key:"status"}]},W=function(e){var t=f;e&&(t=e);var n=[];if(t.length===1)t[0].children.forEach(function(y,$){n.push({key:$,wholesaleType1:y.text||"\u793A\u4F8B\u54C1\u7C7B",wholesaleType2:y.text||"\u793A\u4F8B\u54C1\u7C7B",unit:y.unit||"",package:y.package||"",price:y.price||"",status:!0})});else for(var r=t[0].children,u=t[1].children,c=0;c<r.length;c++)for(var p=0;p<u.length;p++)n.push({key:c+"_"+p,wholesaleType1:r[c].text||"\u793A\u4F8B\u54C1\u7C7B",wholesaleType2:u[p].text||"\u793A\u4F8B\u54C1\u7C7B",unit:r[c].unit||"",package:r[p].package||"",price:r[p].price||"",status:!0});V(n)},N=function(e){var t=o;e&&(t=e);var n=[];if(t.length===1)t[0].children.forEach(function(y,$){n.push({key:$,retailsaleType1:y.text||"\u793A\u4F8B\u54C1\u7C7B",retailsaleType2:y.text||"\u793A\u4F8B\u54C1\u7C7B",unit:y.unit||"",package:y.package||"",price:y.price||"",status:!1})});else for(var r=t[0].children,u=t[1].children,c=0;c<r.length;c++)for(var p=0;p<u.length;p++)n.push({key:c+"_"+p,retailsaleType1:r[c].text||"\u793A\u4F8B\u54C1\u7C7B",retailsaleType2:u[p].text||"\u793A\u4F8B\u54C1\u7C7B",unit:r[c].unit||"",package:r[p].package||"",price:r[p].price||"",status:!1});ae(n)},ua=function(){var l=(0,H.Z)((0,s.Z)().mark(function e(t){var n,r;return(0,s.Z)().wrap(function(c){for(;;)switch(c.prev=c.next){case 0:return n=t[t.length-1],n.loading=!0,c.next=4,(0,Q.mZ)({code:n.value});case 4:r=c.sent,n.loading=!1,n.children=r.map(function(p){return{label:p.name,value:p.code,isLeaf:p.level===3}}),Ke((0,d.Z)(ze));case 8:case"end":return c.stop()}},e)}));return function(t){return l.apply(this,arguments)}}(),ia=function(){var l=(0,H.Z)((0,s.Z)().mark(function e(t){var n,r;return(0,s.Z)().wrap(function(c){for(;;)switch(c.prev=c.next){case 0:return n=t[t.length-1],n.loading=!0,c.next=4,(0,Q.BH)({level:n.level+1,parentId:n.id});case 4:r=c.sent,n.loading=!1,n.children=r.map(function(p){return{label:p.name,value:p.id,id:p.id,level:p.level,isLeaf:p.level===3}}),Le((0,d.Z)(Ee));case 8:case"end":return c.stop()}},e)}));return function(t){return l.apply(this,arguments)}}(),sa=function(){var l=(0,H.Z)((0,s.Z)().mark(function e(){var t;return(0,s.Z)().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return r.next=2,(0,Q.mZ)({});case 2:t=r.sent,Ke(t.map(function(u){return{label:u.name,value:u.code,isLeaf:u.level===3}}));case 4:case"end":return r.stop()}},e)}));return function(){return l.apply(this,arguments)}}(),oa=function(){var l=(0,H.Z)((0,s.Z)().mark(function e(){var t;return(0,s.Z)().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return r.next=2,(0,Q.BH)({level:1});case 2:t=r.sent,Le(t.map(function(u){return{label:u.name,value:u.id,id:u.id,level:u.level,isLeaf:!1}}));case 4:case"end":return r.stop()}},e)}));return function(){return l.apply(this,arguments)}}();(0,h.useEffect)(function(){W(),N(),sa(),oa()},[]);var lt=(0,a.jsxs)("div",{children:[(0,a.jsx)(ct.Z,{}),(0,a.jsx)("div",{style:{marginTop:8},children:"Upload"})]}),da=function(e){ie(e.fileList)},ca=function(e){ie(Z.filter(function(t){return t!==e}))},pa=function(e){de(e.fileList)},ha=function(e){de(Z.filter(function(t){return t!==e}))},va=function(){return Pt(!1)},ga=(0,h.useCallback)(function(l,e){var t=Z[l];ie(je()(Z,{$splice:[[l,1],[e,0,t]]}))},[Z]),fa=(0,h.useCallback)(function(l,e){var t=P[l];de(je()(P,{$splice:[[l,1],[e,0,t]]}))},[P]),rt=function(e){var t=e.originNode,n=e.moveRow,r=e.file,u=e.fileList,c=(0,h.useRef)(null),p=u.indexOf(r),y=(0,pt.L)({accept:ke,collect:function(z){var qa=z.getItem()||{},st=qa.index;return st===p?{}:{isOver:z.isOver(),dropClassName:st<p?" drop-over-downward":" drop-over-upward"}},drop:function(z){n(z.index,p)}}),$=(0,g.Z)(y,2),ut=$[0],Ka=ut.isOver,Ja=ut.dropClassName,Ha=$[1],Qa=(0,ht.c)({type:ke,item:{index:p},collect:function(z){return{isDragging:z.isDragging()}}}),Va=(0,g.Z)(Qa,2),Xa=Va[1];Ha(Xa(c));var Ya=(0,a.jsx)(x.Z,{title:"Upload Error",children:t.props.children});return(0,a.jsx)("div",{ref:c,className:"ant-upload-draggable-list-item ".concat(Ka?Ja:""),style:{cursor:"move"},children:r.status==="error"?Ya:t})},ma=function(){var e=[].concat((0,d.Z)(f),[{text:"",children:[{text:""}]}]);U(e),le(tt(e)),W(e)},xa=function(e){var t=(0,d.Z)(f);t.splice(e,1),U(t),le(et(t)),W(t)},Ca=function(e,t){var n=(0,d.Z)(f);n[e].children=n[e].children.filter(function(r,u){return u!==t}),U(n),W()},ya=function(e){var t=(0,d.Z)(f);t[e].children=[].concat((0,d.Z)(t[e].children),[{lable:"",children:[{text:""}]}]),U(t),W()},Sa=function(e,t){var n=(0,d.Z)(f);n[t].text=e.target.value,U(n),f.length===1?le(et(n)):le(tt(n)),W()},Za=function(e,t,n){var r=(0,d.Z)(f);r[t].children[n].text=e.target.value,U(r),W()},Fa=function(e,t){var n=(0,d.Z)(I);n[t].unit=e.target.value,V(n)},ja=function(e,t){if(e.target.value&&!/^(\d){1,7}(\.)?(\d{1,2})?$/.test(e.target.value)||parseFloat(e.target.value)>999999999e-2){L.ZP.warning("\u8BF7\u8F93\u51650-9999999.99\u4E4B\u95F4\u7684\u6570\u5B57");return}var n=(0,d.Z)(I);n[t].package=e.target.value,V(n)},wa=function(e,t){if(e.target.value&&!/^(\d){1,7}(\.)?(\d{1,2})?$/.test(e.target.value)||parseFloat(e.target.value)>999999999e-2){L.ZP.warning("\u8BF7\u8F93\u51650-9999999.99\u4E4B\u95F4\u7684\u6570\u5B57");return}var n=(0,d.Z)(I);n[t].price=e.target.value,V(n)},Da=function(e,t){var n=(0,d.Z)(I);n[t].status=e,e||(n[t].unit="",n[t].package="",n[t].price=""),V(n)},ka=function(){var e=[].concat((0,d.Z)(o),[{text:"",children:[{text:""}]}]);M(e),ne(nt(e)),N(e)},Ia=function(e){var t=(0,d.Z)(o);t.splice(e,1),M(t),ne(at(t)),N(t)},Ea=function(e,t){var n=(0,d.Z)(o);n[e].children=n[e].children.filter(function(r,u){return u!==t}),M(n),N()},La=function(e){var t=(0,d.Z)(o);t[e].children=[].concat((0,d.Z)(t[e].children),[{lable:"",children:[{text:""}]}]),M(t),N()},Ta=function(e,t){var n=(0,d.Z)(o);n[t].text=e.target.value,M(n),o.length===1?ne(at(n)):ne(nt(n)),N()},Ba=function(e,t,n){var r=(0,d.Z)(o);r[t].children[n].text=e.target.value,M(r),N()},Aa=function(e,t){if(e.target.value&&!/^(\d){1,7}(\.)?(\d{1,2})?$/.test(e.target.value)||parseFloat(e.target.value)>999999999e-2){L.ZP.warning("\u8BF7\u8F93\u51650-9999999.99\u4E4B\u95F4\u7684\u6570\u5B57");return}var n=(0,d.Z)(F);n[t].stock=e.target.value,ae(n)},Ra=function(e,t){if(e.target.value&&!/^(\d){1,7}(\.)?(\d{1,2})?$/.test(e.target.value)||parseFloat(e.target.value)>999999999e-2){L.ZP.warning("\u8BF7\u8F93\u51650-9999999.99\u4E4B\u95F4\u7684\u6570\u5B57");return}var n=(0,d.Z)(F);n[t].price=e.target.value,ae(n)},ba=function(e,t){var n=(0,d.Z)(F);n[t].status=e,e||(n[t].stock="",n[t].price=""),ae(n)},Pa=[{label:"\u6574\u8F66",value:"1"},{label:"\u7269\u6D41/\u4E13\u7EBF",value:"2"},{label:"\u5FEB\u9012",value:"3"},{label:"\u7A7A\u8FD0",value:"4"},{label:"\u94C1\u8DEF",value:"5"},{label:"\u5176\u4ED6\u8FD0\u8F93",value:"6"}],Wa=[{label:"24\u5C0F\u65F6\u5185",value:"1"},{label:"48\u5C0F\u65F6\u5185",value:"2"},{label:"7\u5929\u5185",value:"3"}],Na=function(e){It(e.target.value)},Ga=function(e){Vt(e)},Oa=function(e,t){Yt(e),_t(t==null?void 0:t.map(function(n){return n.label}))},Ua=function(e,t){jt(e),Dt(t==null?void 0:t.map(function(n){return n.label}))},Ma=function(e){ta(e.target.value)},$a=function(e){Lt(e.target.value)},za=function(e){if(e.target.value&&!/^(\d){1,7}(\.)?(\d{1,2})?$/.test(e.target.value)||parseFloat(e.target.value)>999999999e-2){L.ZP.warning("\u8BF7\u8F93\u51650-9999999.99\u4E4B\u95F4\u7684\u6570\u5B57");return}na(e.target.value)};return(0,a.jsx)(dt.ZP,{children:(0,a.jsxs)(m.Z,{labelCol:{span:3},colon:!1,layout:"horizontal",children:[(0,a.jsxs)(b.Z,{title:(0,a.jsx)("div",{style:{fontSize:"20px"},children:"\u5546\u54C1\u4FE1\u606F"}),style:{marginBottom:"16px"},children:[(0,a.jsxs)(m.Z.Item,{label:"\u5546\u54C1\u8F6E\u64AD\u56FE",wrapperCol:{span:24},required:!0,children:[(0,a.jsxs)("div",{style:{margin:"5px 0"},children:["\u56FE\u7247\u8981\u6C42\uFF1A\u5BBD\u9AD8\u6BD4\u4F8B\u4E3A1:1\u3002\u5DF2\u4E0A\u4F20",Z.length,"/",ue,"\u5F20\uFF0C\u62D6\u62FD\u53EF\u8C03\u6574\u987A\u5E8F"]}),(0,a.jsx)(Ze.W,{backend:Fe.PD,children:(0,a.jsx)(vt.Z,{rotationSlider:!0,children:(0,a.jsx)(E.Z,{action:"/api/v1/common/uploadFile",listType:"picture-card",maxCount:ue,headers:{Authorization:localStorage.getItem("token")||""},showUploadList:{showPreviewIcon:!1},onChange:da,beforeUpload:we,onRemove:function(e){return ca(e)},fileList:Z,itemRender:function(e,t,n){return(0,a.jsx)(rt,{originNode:e,file:t,fileList:n,moveRow:ga})},children:Z.length>=ue?null:lt})})}),(0,a.jsx)(J.Z,{open:bt,title:Mt,footer:null,onCancel:va,children:(0,a.jsx)("img",{alt:"example",style:{width:"100%"},src:Gt})})]}),(0,a.jsx)(m.Z.Item,{required:!0,label:"\u5546\u54C1\u5206\u7C7B",wrapperCol:{span:12},rules:[{required:!0,message:"\u8BF7\u9009\u62E9\u5546\u54C1\u5206\u7C7B"}],children:(0,a.jsx)(O.Z,{allowClear:!1,options:Ee,placeholder:"\u8BF7\u9009\u62E9\u5546\u54C1\u5206\u7C7B",loadData:ia,onChange:Ua})}),(0,a.jsx)(m.Z.Item,{label:"\u5546\u54C1\u6807\u9898",wrapperCol:{span:12},required:!0,children:(0,a.jsx)(S.Z,{value:oe,onChange:function(e){Na(e)},placeholder:"\u6700\u591A\u8F93\u5165100\u4E2A\u5B57\u7B26",maxLength:100})}),(0,a.jsx)(m.Z.Item,{label:"\u5546\u54C1\u8BE6\u60C5",wrapperCol:{span:24},children:(0,a.jsxs)("div",{className:q().productDetailWrapper,children:[(0,a.jsx)("div",{children:"\u8BE6\u60C5\u4ECB\u7ECD\u5546\u54C1\u4EE5\u63D0\u5347\u8F6C\u5316\u3002\u82E5\u672A\u7F16\u8F91\uFF0C\u5546\u54C1\u53D1\u5E03\u540E\u8F6E\u64AD\u56FE\u5C06\u81EA\u52A8\u586B\u5145\u81F3\u56FE\u6587\u8BE6\u60C5"}),(0,a.jsxs)("div",{className:q().detailWrap,children:[(0,a.jsx)("div",{className:q().detailTitle,children:"\u6587\u5B57\u63CF\u8FF0"}),(0,a.jsx)(xt,{rows:4,placeholder:"\u6700\u591A\u8F93\u5165200\u4E2A\u5B57\u7B26",maxLength:200,value:be,onChange:$a}),(0,a.jsx)("div",{className:q().detailTitle,children:"\u8BE6\u60C5\u56FE\u7247"}),(0,a.jsxs)("div",{style:{marginBottom:"4px"},children:["\u4EC5\u652F\u6301\u4E0A\u4F20\u56FE\u7247\uFF0C\u6700\u591A\u53EF\u4E0A\u4F20",_,"\u5F20\uFF0C\u5DF2\u4E0A\u4F20",P.length,"/",_,"\u5F20\uFF0C\u62D6\u62FD\u53EF\u8C03\u6574\u987A\u5E8F"]}),(0,a.jsx)(Ze.W,{backend:Fe.PD,children:(0,a.jsx)(E.Z,{action:"/api/v1/common/uploadFile",listType:"picture-card",multiple:!0,maxCount:_,headers:{Authorization:localStorage.getItem("token")||""},showUploadList:{showPreviewIcon:!1},onChange:pa,beforeUpload:we,onRemove:function(e){return ha(e)},fileList:P,itemRender:function(e,t,n){return(0,a.jsx)(rt,{originNode:e,file:t,fileList:n,moveRow:fa})},children:P.length>=_?null:lt})})]})]})})]}),(0,a.jsxs)(b.Z,{title:(0,a.jsx)("div",{style:{fontSize:"20px"},children:"\u4EF7\u683C\u4FE1\u606F"}),style:{marginBottom:"16px"},children:[(0,a.jsx)(m.Z.Item,{required:!0,label:"\u6279\u53D1\u89C4\u683C",children:(0,a.jsxs)("div",{style:{backgroundColor:"#f0f0f0",padding:"24px"},children:[(0,a.jsx)("div",{style:{marginBottom:"8px"},children:"\u6700\u591A\u6DFB\u52A02\u4E2A\u5546\u54C1\u89C4\u683C\u7C7B\u578B"}),(0,a.jsx)("div",{children:f==null?void 0:f.map(function(l,e){var t;return(0,a.jsxs)("div",{style:{padding:"8px 16px",marginBottom:"8px",backgroundColor:"#fff"},children:[(0,a.jsx)(S.Z,{value:l.text,onChange:function(r){return Sa(r,e)},style:{width:"120px",marginRight:"16px"},maxLength:2,showCount:!0,placeholder:"\u54C1\u7C7B"}),f.length>1&&(0,a.jsx)(Y.Z,{style:{fontSize:16},onClick:function(){return xa(e)}}),(0,a.jsx)(R.Z,{style:{color:"#f0f0f0",margin:"6px"}}),(0,a.jsxs)("div",{style:{display:"flex",alignItems:"center"},children:[(t=l.children)===null||t===void 0?void 0:t.map(function(n,r){return(0,a.jsx)(S.Z,{value:n.text,onChange:function(c){return Za(c,e,r)},style:{width:"156px",marginRight:"8px"},maxLength:20,placeholder:"\u793A\u4F8B\u54C1\u7C7B".concat(De[r]),addonAfter:l.children.length>1?(0,a.jsx)(Y.Z,{onClick:function(){return Ca(e,r)}}):null},r)}),(0,a.jsx)(Se.Z,{style:{fontSize:16},onClick:function(){return ya(e)}})]})]},e)})}),(0,a.jsxs)(A.Z,{type:"primary",style:{marginTop:"12px"},onClick:ma,disabled:f.length===2,children:["\u6DFB\u52A0\u5546\u54C1\u89C4\u683C\uFF08",f.length,"/2\uFF09"]})]})}),(0,a.jsx)(m.Z.Item,{required:!0,label:"\u6279\u53D1\u4EF7\u683C",children:(0,a.jsxs)("div",{style:{backgroundColor:"#f0f0f0",padding:"24px",marginBottom:"36px"},children:[(0,a.jsx)("div",{style:{backgroundColor:"#fff",display:"grid",gridTemplateColumns:"repeat(".concat(B.length,", 1fr)")},children:B==null?void 0:B.map(function(l,e){return(0,a.jsx)("div",{style:{display:"flex",alignItems:"center",padding:"0 16px",height:"50px",border:"1px solid #f0f0f0"},children:l.title},e)})}),I==null?void 0:I.map(function(l,e){return(0,a.jsx)("div",{style:{backgroundColor:"#fff",display:"grid",gridTemplateColumns:"repeat(".concat(B.length,", 1fr)")},children:B==null?void 0:B.map(function(t,n){return(0,a.jsxs)("div",{style:{display:"flex",alignItems:"center",padding:"0 16px",height:"40px",border:"1px solid #f0f0f0"},children:[t.key==="wholesaleType1"&&(0,a.jsx)("span",{children:l.wholesaleType1}),t.key==="wholesaleType2"&&(0,a.jsx)("span",{children:l.wholesaleType2}),t.key==="unit"&&(0,a.jsx)(S.Z,{value:l.unit,maxLength:2,showCount:!0,onChange:function(u){return Fa(u,e)},placeholder:"\u8BF7\u8F93\u5165"}),t.key==="package"&&(0,a.jsxs)("div",{style:{display:"flex",alignItems:"center"},children:[(0,a.jsx)(S.Z,{style:{flex:1,marginRight:"4px"},maxLength:10,value:l.package,onChange:function(u){return ja(u,e)},placeholder:"\u8BF7\u8F93\u5165"}),(0,a.jsx)("span",{children:l.unit||"\u65A4"})]}),t.key==="price"&&(0,a.jsxs)("div",{style:{display:"flex",alignItems:"center"},children:[(0,a.jsx)(S.Z,{style:{flex:1,marginRight:"4px"},maxLength:10,value:l.price,onChange:function(u){return wa(u,e)},placeholder:"\u8BF7\u8F93\u5165"}),(0,a.jsxs)("span",{children:["\u5143/",l.unit||"\u65A4"]})]}),t.key==="status"&&(0,a.jsx)(G.Z,{checkedChildren:"\u5DF2\u542F\u7528",unCheckedChildren:"\u672A\u542F\u7528",checked:l.status,onChange:function(u){return Da(u,e)}})]},n)})},e)})]})}),(0,a.jsx)(m.Z.Item,{label:"\u96F6\u552E\u89C4\u683C",required:o[0].text,children:(0,a.jsxs)("div",{style:{backgroundColor:"#f0f0f0",padding:"24px"},children:[(0,a.jsx)("div",{style:{marginBottom:"8px"},children:"\u6700\u591A\u6DFB\u52A02\u4E2A\u5546\u54C1\u89C4\u683C\u7C7B\u578B"}),(0,a.jsx)("div",{children:o==null?void 0:o.map(function(l,e){var t;return(0,a.jsxs)("div",{style:{padding:"8px 16px",marginBottom:"8px",backgroundColor:"#fff"},children:[(0,a.jsx)(S.Z,{value:l.text,onChange:function(r){return Ta(r,e)},style:{width:"120px",marginRight:"16px"},maxLength:2,showCount:!0,placeholder:"\u54C1\u7C7B"}),o.length>1&&(0,a.jsx)(Y.Z,{style:{fontSize:16},onClick:function(){return Ia(e)}}),(0,a.jsx)(R.Z,{style:{color:"#f0f0f0",margin:"6px"}}),(0,a.jsxs)("div",{style:{display:"flex",alignItems:"center"},children:[(t=l.children)===null||t===void 0?void 0:t.map(function(n,r){return(0,a.jsx)(S.Z,{value:n.text,onChange:function(c){return Ba(c,e,r)},style:{width:"156px",marginRight:"8px"},maxLength:20,placeholder:"\u793A\u4F8B\u54C1\u7C7B".concat(De[r]),addonAfter:l.children.length>1?(0,a.jsx)(Y.Z,{onClick:function(){return Ea(e,r)}}):null},r)}),(0,a.jsx)(Se.Z,{style:{fontSize:16},onClick:function(){return La(e)}})]})]},e)})}),(0,a.jsxs)(A.Z,{type:"primary",style:{marginTop:"12px"},onClick:ka,disabled:o.length===2,children:["\u6DFB\u52A0\u5546\u54C1\u89C4\u683C\uFF08",o.length,"/2\uFF09"]})]})}),(0,a.jsx)(m.Z.Item,{label:"\u96F6\u552E\u4EF7\u683C",required:o[0].text,children:(0,a.jsxs)("div",{style:{backgroundColor:"#f0f0f0",padding:"24px",marginBottom:"36px"},children:[(0,a.jsx)("div",{style:{backgroundColor:"#fff",display:"grid",gridTemplateColumns:"repeat(".concat(T.length,", 1fr)")},children:T==null?void 0:T.map(function(l,e){return(0,a.jsx)("div",{style:{display:"flex",alignItems:"center",padding:"0 16px",height:"50px",border:"1px solid #f0f0f0"},children:l.title},e)})}),F==null?void 0:F.map(function(l,e){return(0,a.jsx)("div",{style:{backgroundColor:"#fff",display:"grid",gridTemplateColumns:"repeat(".concat(T.length,", 1fr)")},children:T==null?void 0:T.map(function(t,n){return(0,a.jsxs)("div",{style:{display:"flex",alignItems:"center",padding:"0 16px",height:"40px",border:"1px solid #f0f0f0"},children:[t.key==="retailsaleType1"&&(0,a.jsx)("span",{children:l.retailsaleType1}),t.key==="retailsaleType2"&&(0,a.jsx)("span",{children:l.retailsaleType2}),t.key==="stock"&&(0,a.jsx)(S.Z,{value:l.stock,maxLength:10,onChange:function(u){return Aa(u,e)},disabled:l.status===!1,placeholder:"\u8BF7\u8F93\u5165"}),t.key==="price"&&(0,a.jsxs)("div",{style:{display:"flex",alignItems:"center"},children:[(0,a.jsx)(S.Z,{style:{flex:1,marginRight:"4px"},maxLength:10,value:l.price,onChange:function(u){return Ra(u,e)},disabled:l.status===!1,placeholder:"\u8BF7\u8F93\u5165"}),(0,a.jsx)("span",{children:"\u5143"})]}),t.key==="status"&&(0,a.jsx)(G.Z,{checkedChildren:"\u5DF2\u542F\u7528",unCheckedChildren:"\u672A\u542F\u7528",checked:l.status,onChange:function(u){return ba(u,e)}})]},n)})},e)})]})})]}),(0,a.jsxs)(b.Z,{title:(0,a.jsx)("div",{style:{fontSize:"20px"},children:"\u670D\u52A1\u548C\u627F\u8BFA"}),style:{marginBottom:"16px"},children:[(0,a.jsx)(m.Z.Item,{required:!0,label:"\u6279\u53D1\u7269\u6D41",children:(0,a.jsx)(D.Z.Group,{options:Pa,defaultValue:ce,value:ce,onChange:Ga})}),(0,a.jsx)(m.Z.Item,{wrapperCol:{span:12},required:!0,label:"\u6279\u53D1\u53D1\u8D27\u5730",children:(0,a.jsx)(O.Z,{allowClear:!1,options:ze,placeholder:"\u8BF7\u9009\u62E9\u6279\u53D1\u53D1\u8D27\u5730",loadData:ua,onChange:Oa})}),(0,a.jsx)(R.Z,{}),(0,a.jsx)(m.Z.Item,{required:o[0].text,label:"\u96F6\u552E\u53D1\u8D27\u65F6\u95F4",rules:[{required:o[0].text}],children:(0,a.jsx)(w.ZP.Group,{disabled:_e(),options:Wa,value:Ye,onChange:Ma})}),(0,a.jsxs)(m.Z.Item,{required:o[0].text,label:"\u96F6\u552E\u7269\u6D41",children:["\u5317\u4EAC\u5305\u90AE\uFF0C\u5176\u4ED6\u5730\u533A\u7EDF\u4E00\u90AE\u8D39",(0,a.jsx)(S.Z,{disabled:_e(),style:{width:"60px",display:"inline",margin:"0 8px"},value:he,onChange:za}),"\u5143"]})]}),(0,a.jsx)("div",{style:{display:"flex",justifyContent:"center",marginTop:"20px"},children:(0,a.jsx)(A.Z,{size:"large",type:"primary",disabled:ra(),onClick:la,children:"\u63D0\u4EA4\u4E0A\u67B6"})})]})})},yt=Ct},69083:function(re,X,i){"use strict";i.d(X,{g2:function(){return ve},k1:function(){return ge},mZ:function(){return fe},BH:function(){return me},JJ:function(){return xe},pm:function(){return Ce},ys:function(){return ye}});var C=i(90636),w=i(3182),K=i(99871),D=i(636);function ve(v){return G.apply(this,arguments)}function G(){return G=(0,w.Z)((0,C.Z)().mark(function v(x){return(0,C.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,D.Z)("/api/v1/goods/getMyGoodsDetail?".concat((0,K.R)(x))));case 1:case"end":return s.stop()}},v)})),G.apply(this,arguments)}function ot(v){return A.apply(this,arguments)}function A(){return A=_asyncToGenerator(_regeneratorRuntime().mark(function v(x){return _regeneratorRuntime().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",request("/api/v1/goods/getGoodsList?".concat(objectToUrlParams(x))));case 1:case"end":return s.stop()}},v)})),A.apply(this,arguments)}function ge(v){return R.apply(this,arguments)}function R(){return R=(0,w.Z)((0,C.Z)().mark(function v(x){return(0,C.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,D.Z)("/api/v1/goods/getMyGoodsList?".concat((0,K.R)(x))));case 1:case"end":return s.stop()}},v)})),R.apply(this,arguments)}function fe(v){return b.apply(this,arguments)}function b(){return b=(0,w.Z)((0,C.Z)().mark(function v(x){return(0,C.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,D.Z)("/api/v1/common/area?".concat((0,K.R)(x))));case 1:case"end":return s.stop()}},v)})),b.apply(this,arguments)}function me(v){return O.apply(this,arguments)}function O(){return O=(0,w.Z)((0,C.Z)().mark(function v(x){return(0,C.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,D.Z)("/api/v1/mall/categories?".concat((0,K.R)(x))));case 1:case"end":return s.stop()}},v)})),O.apply(this,arguments)}function xe(v){return m.apply(this,arguments)}function m(){return m=(0,w.Z)((0,C.Z)().mark(function v(x){return(0,C.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,D.Z)("/api/v1/goods/addGoods",{method:"POST",data:x}));case 1:case"end":return s.stop()}},v)})),m.apply(this,arguments)}function Ce(v){return J.apply(this,arguments)}function J(){return J=(0,w.Z)((0,C.Z)().mark(function v(x){return(0,C.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,D.Z)("/api/v1/goods/modifyMyGoodsStatus",{method:"POST",data:x}));case 1:case"end":return s.stop()}},v)})),J.apply(this,arguments)}function ye(v){return E.apply(this,arguments)}function E(){return E=(0,w.Z)((0,C.Z)().mark(function v(x){return(0,C.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,D.Z)("/api/v1/goods/deleteMyGoods",{method:"DELETE",data:x}));case 1:case"end":return s.stop()}},v)})),E.apply(this,arguments)}}}]);