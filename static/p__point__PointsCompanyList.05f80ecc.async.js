(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[829],{87588:function(C,h,n){"use strict";var r=n(22122),i=n(67294),p=n(61144),u=n(65734),R=function(T,d){return i.createElement(u.Z,(0,r.Z)({},T,{ref:d,icon:p.Z}))};h.Z=i.forwardRef(R)},87672:function(C){C.exports={"search-wrapper":"search-wrapper___k6skK",actions:"actions___2Y8k7","table-wrapper":"table-wrapper___1b7q7"}},26555:function(C,h,n){"use strict";n.r(h);var r=n(8963),i=n(38291),p=n(9715),u=n(71481),R=n(47673),g=n(77808),T=n(34792),d=n(48086),P=n(90636),y=n(11849),A=n(3182),o=n(2824),S=n(57663),E=n(71577),W=n(71194),M=n(50146),m=n(67294),v=n(96486),b=n.n(v),Z=n(36773),B=n(87588),O=n(84514),t=n(87672),s=n.n(t),a=n(85893),e=M.Z.confirm,F=function(){var X=[{title:"\u516C\u53F8\u540D\u79F0",dataIndex:"name",key:"name"},{title:"\u91C7\u8D2D\u8BA4\u8BC1\u4EBA\u6570",dataIndex:"totalMember",key:"totalMember"},{title:"\u53D1\u653E\u79EF\u5206\u4EBA\u6570",dataIndex:"pointMember",key:"pointMember"},{title:"\u53EF\u7528\u79EF\u5206\u603B\u989D",dataIndex:"totalPoint",key:"totalPoint"},{title:"\u79EF\u5206\u7533\u8BF7\u8BB0\u5F55",dataIndex:"action",key:"action",render:function(_,I){return(0,a.jsx)(E.Z,{target:"_blank",href:"?id=".concat(I.id,"#/point/company/history"),type:"link",children:"\u67E5\u770B"})}}],G=(0,m.useState)(!1),w=(0,o.Z)(G,2),Y=w[0],L=w[1],J=(0,m.useState)({current:1,pageSize:10,showSizeChanger:!0,showQuickJumper:!0,showTotal:function(_){return"\u603B\u5171 ".concat(_," \u6761")}}),x=(0,o.Z)(J,2),D=x[0],z=x[1],Q=(0,m.useState)([]),N=(0,o.Z)(Q,2),K=N[0],$=N[1],H=function(_){$(_)},q={selectedRowKeys:K,onChange:H},ee=function(_){z({current:_.current||1,pageSize:_.pageSize||10})},ne=(0,m.useState)([]),k=(0,o.Z)(ne,2),te=k[0],ae=k[1],j=(0,v.debounce)((0,A.Z)((0,P.Z)().mark(function f(){var _;return(0,P.Z)().wrap(function(l){for(;;)switch(l.prev=l.next){case 0:return l.prev=0,L(!0),l.next=4,O.jX({name:U,pageNum:D.current,pageSize:D.pageSize});case 4:_=l.sent,_&&(ae(_.list),z((0,y.Z)((0,y.Z)({},D),{},{total:_.totalNum}))),L(!1),l.next=12;break;case 9:l.prev=9,l.t0=l.catch(0),L(!1);case 12:case"end":return l.stop()}},f,null,[[0,9]])}))),re=(0,m.useState)(""),V=(0,o.Z)(re,2),U=V[0],se=V[1],_e=function(_){se(_.companyName),U===_.companyName&&j()};(0,m.useEffect)(function(){j()},[D.current,D.pageSize,U]);var ue=function(){e({title:"\u786E\u5B9A\u6E05\u96F6\u79EF\u5206\u5417\uFF1F",icon:(0,a.jsx)(B.Z,{}),content:"\u6B64\u64CD\u4F5C\u5C06\u6E05\u96F6\u6240\u9009\u516C\u53F8\u540D\u4E0B\u6240\u6709\u79EF\u5206\uFF0C\u8BF7\u8C28\u614E\u64CD\u4F5C",onOk:function(){return(0,A.Z)((0,P.Z)().mark(function I(){return(0,P.Z)().wrap(function(c){for(;;)switch(c.prev=c.next){case 0:return c.prev=0,c.next=3,O.kz({orgID:K[0]});case 3:d.ZP.success("\u64CD\u4F5C\u6210\u529F"),$([]),j(),c.next=10;break;case 8:c.prev=8,c.t0=c.catch(0);case 10:case"end":return c.stop()}},I,null,[[0,8]])}))()}})};return(0,a.jsxs)(Z.ZP,{children:[(0,a.jsx)("div",{className:s()["search-wrapper"],children:(0,a.jsxs)(u.Z,{name:"basic",initialValues:{companyName:U},labelCol:{span:8},layout:"inline",onFinish:_e,wrapperCol:{span:16},autoComplete:"off",children:[(0,a.jsx)(u.Z.Item,{label:"\u516C\u53F8\u540D\u79F0",name:"companyName",rules:[{required:!1}],children:(0,a.jsx)(g.Z,{maxLength:80})}),(0,a.jsx)(u.Z.Item,{children:(0,a.jsx)(E.Z,{htmlType:"submit",type:"primary",children:"\u67E5\u8BE2"})})]})}),(0,a.jsxs)("div",{className:s()["table-wrapper"],children:[(0,a.jsx)("div",{className:s().actions,children:(0,a.jsx)(E.Z,{disabled:K.length!==1,type:"primary",onClick:ue,children:"\u79EF\u5206\u6E05\u96F6"})}),(0,a.jsx)(i.Z,{rowKey:"id",columns:X,dataSource:te,loading:Y,onChange:ee,pagination:D,rowSelection:q,scroll:{x:"max-content"}})]})]})};h.default=F},84514:function(C,h,n){"use strict";n.d(h,{dV:function(){return R},UX:function(){return T},d0:function(){return P},Uf:function(){return A},kz:function(){return S},YL:function(){return W},jX:function(){return m},jh:function(){return b},sI:function(){return B}});var r=n(90636),i=n(3182),p=n(99871),u=n(636);function R(t){return g.apply(this,arguments)}function g(){return g=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/getApplyToVerify"));case 1:case"end":return e.stop()}},t)})),g.apply(this,arguments)}function T(t){return d.apply(this,arguments)}function d(){return d=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/getApplys?".concat((0,p.R)(s))));case 1:case"end":return e.stop()}},t)})),d.apply(this,arguments)}function P(t){return y.apply(this,arguments)}function y(){return y=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/applyPoint",{method:"POST",data:s,headers:{"Content-Type":"multipart/form-data"}}));case 1:case"end":return e.stop()}},t)})),y.apply(this,arguments)}function A(t){return o.apply(this,arguments)}function o(){return o=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/verifyPoint",{method:"POST",data:s}));case 1:case"end":return e.stop()}},t)})),o.apply(this,arguments)}function S(t){return E.apply(this,arguments)}function E(){return E=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/clearPoint",{method:"POST",data:s}));case 1:case"end":return e.stop()}},t)})),E.apply(this,arguments)}function W(t){return M.apply(this,arguments)}function M(){return M=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/getAccountVerifyList?".concat((0,p.R)(s))));case 1:case"end":return e.stop()}},t)})),M.apply(this,arguments)}function m(t){return v.apply(this,arguments)}function v(){return v=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/getOrganizations?".concat((0,p.R)(s))));case 1:case"end":return e.stop()}},t)})),v.apply(this,arguments)}function b(t){return Z.apply(this,arguments)}function Z(){return Z=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/getPointRecordsByUser?".concat((0,p.R)(s))));case 1:case"end":return e.stop()}},t)})),Z.apply(this,arguments)}function B(t){return O.apply(this,arguments)}function O(){return O=(0,i.Z)((0,r.Z)().mark(function t(s){return(0,r.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.Z)("/api/v1/org/getPointRecordsByApply?".concat((0,p.R)(s))));case 1:case"end":return e.stop()}},t)})),O.apply(this,arguments)}}}]);
