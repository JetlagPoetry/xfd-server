(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[240,626],{3122:function(g){g.exports={"check-point-page":"check-point-page___2IgQP","module-title":"module-title___1pW4G","actions-wrapper":"actions-wrapper___13-_v","points-process":"points-process___1exmi",form:"form___k5C6u",flex:"flex___JLoLX",key:"key___1pqq0","ant-radio-wrapper":"ant-radio-wrapper___-ePag"}},10873:function(g,d,r){"use strict";r.r(d);var s=r(11849),c=r(57663),m=r(71577),i=r(88983),n=r(47933),h=r(98858),o=r(4914),v=r(34792),y=r(48086),f=r(90636),R=r(3182),T=r(9715),P=r(71481),E=r(2824),L=r(47673),D=r(77808),Z=r(67294),O=r(27484),I=r.n(O),M=r(36773),j=r(84514),B=r(3122),a=r.n(B),_=r(81910),t=r(85893),e={wrapperCol:{span:14}},x=D.Z.TextArea,S=function(){var $=P.Z.useForm(),k=(0,E.Z)($,1),C=k[0],N=P.Z.useWatch("verifyStatus",C),z=(0,Z.useState)({applyURL:"",comment:"",hasNext:!1,id:"",organizationName:"",totalPoint:0,submitTime:"",username:"",userPhone:"",userPosition:""}),W=(0,E.Z)(z,2),p=W[0],V=W[1],K=function(){var U=(0,R.Z)((0,f.Z)().mark(function A(){var l;return(0,f.Z)().wrap(function(u){for(;;)switch(u.prev=u.next){case 0:return u.prev=0,u.next=3,j.dV();case 3:l=u.sent,l&&V({applyURL:l.applyURL,comment:l.comment,hasNext:l.hasNext,id:l.id,organizationName:l.organizationName,totalPoint:l.totalPoint,submitTime:I()(l.submitTime).format("YYYY-MM-DD HH:mm"),username:l.username,userPhone:l.userPhone,userPosition:l.userPosition}),u.next=9;break;case 7:u.prev=7,u.t0=u.catch(0);case 9:case"end":return u.stop()}},A,null,[[0,7]])}));return function(){return U.apply(this,arguments)}}();(0,Z.useEffect)(function(){K()},[]);var b=function(){_.m8.goBack()},Y=(0,Z.useState)(!1),w=(0,E.Z)(Y,2),G=w[0],F=w[1],X=function(){var U=(0,R.Z)((0,f.Z)().mark(function A(){var l;return(0,f.Z)().wrap(function(u){for(;;)switch(u.prev=u.next){case 0:return u.next=2,C.validateFields();case 2:return l=u.sent,u.prev=3,F(!0),u.next=7,j.Uf({id:p.id,verifyComment:l.verifyComment,verifyStatus:l.verifyStatus});case 7:y.ZP.success("\u64CD\u4F5C\u6210\u529F"),F(!1),C.resetFields(),p.hasNext?K():b(),u.next=15;break;case 13:u.prev=13,u.t0=u.catch(3);case 15:case"end":return u.stop()}},A,null,[[3,13]])}));return function(){return U.apply(this,arguments)}}();return(0,t.jsx)(M.ZP,{children:(0,t.jsxs)("div",{className:a()["check-point-page"],children:[(0,t.jsxs)(o.Z,{title:"\u79EF\u5206\u7533\u8BF7",children:[(0,t.jsx)(o.Z.Item,{label:"\u7533\u8BF7\u5185\u5BB9",children:(0,t.jsx)("a",{target:"_blank",href:p.applyURL,type:"link",children:"\u67E5\u770B"})}),(0,t.jsx)(o.Z.Item,{label:"\u7533\u8BF7\u4E3B\u4F53",children:p.organizationName}),(0,t.jsx)(o.Z.Item,{label:"\u7533\u8BF7\u8BF4\u660E",children:p.comment}),(0,t.jsx)(o.Z.Item,{label:"\u7533\u8BF7\u65B0\u589E\u79EF\u5206",children:p.totalPoint}),(0,t.jsx)(o.Z.Item,{label:"\u7533\u8BF7\u65F6\u95F4",children:p.submitTime}),(0,t.jsx)(o.Z.Item,{label:"\u7533\u8BF7\u4EBA",children:p.username}),(0,t.jsx)(o.Z.Item,{label:"\u62C5\u4EFB\u804C\u4F4D",children:p.userPosition}),(0,t.jsx)(o.Z.Item,{label:"\u8054\u7CFB\u7535\u8BDD",children:p.userPhone})]}),(0,t.jsx)("p",{className:a()["module-title"],children:"\u5BA1\u6838"}),(0,t.jsxs)(P.Z,(0,s.Z)((0,s.Z)({form:C},e),{},{children:[(0,t.jsx)(P.Z.Item,{name:"verifyStatus",label:"\u5BA1\u6838\u7ED3\u8BBA",rules:[{required:!0,message:"\u8BF7\u9009\u62E9"}],children:(0,t.jsxs)(n.ZP.Group,{children:[(0,t.jsx)(n.ZP,{value:1,children:"\u5BA1\u6838\u901A\u8FC7"}),(0,t.jsx)(n.ZP,{value:2,children:"\u5BA1\u6838\u62D2\u7EDD"})]})}),N===2&&(0,t.jsx)(P.Z.Item,{name:"verifyComment",label:"\u539F\u56E0",children:(0,t.jsx)(x,{rows:4,maxLength:6})}),(0,t.jsxs)("div",{className:a()["actions-wrapper"],children:[(0,t.jsx)(m.Z,{onClick:b,children:"\u9000\u51FA"}),(0,t.jsx)(m.Z,{htmlType:"submit",loading:G,onClick:X,type:"primary",children:p.hasNext?"\u63D0\u4EA4\u5E76\u5BA1\u6838\u4E0B\u4E00\u4E2A":"\u63D0\u4EA4"})]})]}))]})})};d.default=S},84514:function(g,d,r){"use strict";r.d(d,{dV:function(){return n},UX:function(){return o},d0:function(){return y},Uf:function(){return R},kz:function(){return P},YL:function(){return L},jX:function(){return Z},jh:function(){return I},sI:function(){return j}});var s=r(90636),c=r(3182),m=r(99871),i=r(636);function n(a){return h.apply(this,arguments)}function h(){return h=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/getApplyToVerify"));case 1:case"end":return e.stop()}},a)})),h.apply(this,arguments)}function o(a){return v.apply(this,arguments)}function v(){return v=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/getApplys?".concat((0,m.R)(_))));case 1:case"end":return e.stop()}},a)})),v.apply(this,arguments)}function y(a){return f.apply(this,arguments)}function f(){return f=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/applyPoint",{method:"POST",data:_,headers:{"Content-Type":"multipart/form-data"}}));case 1:case"end":return e.stop()}},a)})),f.apply(this,arguments)}function R(a){return T.apply(this,arguments)}function T(){return T=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/verifyPoint",{method:"POST",data:_}));case 1:case"end":return e.stop()}},a)})),T.apply(this,arguments)}function P(a){return E.apply(this,arguments)}function E(){return E=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/clearPoint",{method:"POST",data:_}));case 1:case"end":return e.stop()}},a)})),E.apply(this,arguments)}function L(a){return D.apply(this,arguments)}function D(){return D=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/getAccountVerifyList?".concat((0,m.R)(_))));case 1:case"end":return e.stop()}},a)})),D.apply(this,arguments)}function Z(a){return O.apply(this,arguments)}function O(){return O=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/getOrganizations?".concat((0,m.R)(_))));case 1:case"end":return e.stop()}},a)})),O.apply(this,arguments)}function I(a){return M.apply(this,arguments)}function M(){return M=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/getPointRecordsByUser?".concat((0,m.R)(_))));case 1:case"end":return e.stop()}},a)})),M.apply(this,arguments)}function j(a){return B.apply(this,arguments)}function B(){return B=(0,c.Z)((0,s.Z)().mark(function a(_){return(0,s.Z)().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,i.Z)("/api/v1/org/getPointRecordsByApply?".concat((0,m.R)(_))));case 1:case"end":return e.stop()}},a)})),B.apply(this,arguments)}},99871:function(g,d,r){"use strict";r.d(d,{R:function(){return s},D:function(){return c}});function s(m){var i=Object.keys(m).map(function(n){return"".concat(n,"=").concat(m[n])});return i.join("&")}function c(m){var i=new RegExp("(^|&)"+m+"=([^&]*)(&|$)"),n=window.location.search.substr(1).match(i);return n!=null?decodeURIComponent(n[2]):null}},636:function(g,d,r){"use strict";var s=r(34792),c=r(48086),m=r(12666),i=m.Z.create({baseURL:"/",timeout:3e4,withCredentials:!1});i.interceptors.request.use(function(n){n&&n.headers&&(n.headers["Content-Type"]||(n.headers["Content-Type"]="application/json"));var h=localStorage.getItem("token");return(n==null?void 0:n.url)!=="/api/v1/user/login"&&(n.headers.Authorization="".concat(h)),n},function(n){return Promise.reject(n)}),i.interceptors.response.use(function(n){var h=n.data,o=n.data,v=n.status,y=n.statusText;return v!==200?(c.ZP.error(y),null):o.status===10010?o:o.status!==200?(c.ZP.error(o.msg),null):o.data},function(n){return console.log("err"+n),Promise.reject(n)}),d.Z=i}}]);
