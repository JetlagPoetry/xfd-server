(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[417],{12310:function(ce){ce.exports={"check-point-page":"check-point-page___3LLI8","actions-wrapper":"actions-wrapper___oqRsY","module-content":"module-content___1ojif",content:"content___lfVFi"}},13717:function(ce,re,a){"use strict";a.r(re),a.d(re,{default:function(){return I}});var F=a(57663),U=a(71577),r=a(88983),Q=a(47933),K=a(11849),b=a(49111),de=a(19650),H=a(12968),Y=a(26141),Ne=a(34792),ie=a(48086),p=a(90636),w=a(3182),Se=a(9715),g=a(71481),X=a(2824),fe=a(47673),me=a(60345),z=a(67294),se=a(95916),q=a(636),ve=null,oe=null;function he(u){return _.apply(this,arguments)}function _(){return _=(0,w.Z)((0,p.Z)().mark(function u(f){return(0,p.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,q.Z)("/api/v1/org/getAccountVerify?id=".concat(f.id)));case 1:case"end":return s.stop()}},u)})),_.apply(this,arguments)}function le(u){return ee.apply(this,arguments)}function ee(){return ee=(0,w.Z)((0,p.Z)().mark(function u(f){return(0,p.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,q.Z)("/api/v1/org/getAccountToVerify"));case 1:case"end":return s.stop()}},u)})),ee.apply(this,arguments)}function xe(u){return J.apply(this,arguments)}function J(){return J=_asyncToGenerator(_regeneratorRuntime().mark(function u(f){var C,s;return _regeneratorRuntime().wrap(function(h){for(;;)switch(h.prev=h.next){case 0:return C=f.current,s=_objectWithoutProperties(f,ve),h.abrupt("return",request("/api/v1/org/getAccountVerifyList?".concat(objectToUrlParams(_objectSpread(_objectSpread({},s),{},{pageNum:C})))));case 2:case"end":return h.stop()}},u)})),J.apply(this,arguments)}function ge(u){return te.apply(this,arguments)}function te(){return te=(0,w.Z)((0,p.Z)().mark(function u(f){return(0,p.Z)().wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.abrupt("return",(0,q.Z)("/api/v1/org/verifyAccount",{method:"POST",data:f}));case 1:case"end":return s.stop()}},u)})),te.apply(this,arguments)}function e(u){return i.apply(this,arguments)}function i(){return i=_asyncToGenerator(_regeneratorRuntime().mark(function u(f){var C,s;return _regeneratorRuntime().wrap(function(h){for(;;)switch(h.prev=h.next){case 0:return C=f.current,s=_objectWithoutProperties(f,oe),h.abrupt("return",request("/api/v1/org/getOrganizations?".concat(objectToUrlParams(_objectSpread(_objectSpread({},s),{},{pageNum:C})))));case 2:case"end":return h.stop()}},u)})),i.apply(this,arguments)}var n=a(31134),l=a(12310),d=a.n(l),m=a(81910),y=a(975),A=a(99871),t=a(85893),x={wrapperCol:{span:14}},R=me.Z.TextArea,j=function(){var f=(0,A.D)("id"),C=g.Z.useForm(),s=(0,X.Z)(C,1),N=s[0],h=g.Z.useWatch("verifyStatus",N),T=(0,z.useState)({hasNext:!1,id:"",identityUrlA:"",identityUrlB:"",organization:"",organizationCode:"",organizationUrl:"",phone:"",position:"",realName:"",role:y.i4.Buyers}),S=(0,X.Z)(T,2),v=S[0],P=S[1],D=(0,z.useState)(!1),$=(0,X.Z)(D,2),k=$[0],M=$[1],L=function(){var W=(0,w.Z)((0,p.Z)().mark(function G(){var c;return(0,p.Z)().wrap(function(o){for(;;)switch(o.prev=o.next){case 0:if(o.prev=0,!f){o.next=9;break}return o.next=4,he({id:f});case 4:c=o.sent,N.setFieldsValue({verifyComment:c.comment,verifyStatus:c.status}),c.status===1?M(!1):M(!0),o.next=12;break;case 9:return o.next=11,le();case 11:c=o.sent;case 12:c&&P({hasNext:c.hasNext,id:c.id,identityUrlA:c.identityUrlA,identityUrlB:c.identityUrlB,organization:c.organization,organizationCode:c.organizationCode,organizationUrl:c.organizationUrl,phone:c.phone,position:c.position,realName:c.realName,role:c.role}),o.next=17;break;case 15:o.prev=15,o.t0=o.catch(0);case 17:case"end":return o.stop()}},G,null,[[0,15]])}));return function(){return W.apply(this,arguments)}}();(0,z.useEffect)(function(){L()},[]);var B=function(){m.m8.goBack()},V=(0,z.useState)(!1),E=(0,X.Z)(V,2),ne=E[0],O=E[1],ae=function(){var W=(0,w.Z)((0,p.Z)().mark(function G(){var c;return(0,p.Z)().wrap(function(o){for(;;)switch(o.prev=o.next){case 0:return o.next=2,N.validateFields();case 2:return c=o.sent,o.prev=3,O(!0),o.next=7,ge({id:v.id,comment:c.verifyComment,status:c.verifyStatus});case 7:ie.ZP.success("\u64CD\u4F5C\u6210\u529F"),O(!1),N.resetFields(),v.hasNext?L():B(),o.next=15;break;case 13:o.prev=13,o.t0=o.catch(3);case 15:case"end":return o.stop()}},G,null,[[3,13]])}));return function(){return W.apply(this,arguments)}}();return(0,t.jsx)(se.ZP,{children:(0,t.jsxs)("div",{className:d()["check-point-page"],children:[(0,t.jsx)("h2",{children:"\u8BA4\u8BC1\u4FE1\u606F"}),(0,t.jsxs)(g.Z,(0,K.Z)((0,K.Z)({},x),{},{children:[(0,t.jsx)(g.Z.Item,{label:"\u8BA4\u8BC1\u89D2\u8272",children:(0,t.jsx)("p",{children:n.Zv.get(v.role)})}),(0,t.jsx)(g.Z.Item,{label:"\u673A\u6784\u540D\u79F0",children:(0,t.jsx)("p",{children:v.organization})}),(0,t.jsx)(g.Z.Item,{label:"\u8425\u4E1A\u6267\u7167",children:(0,t.jsx)(Y.Z,{width:100,height:100,src:v.organizationUrl})}),(0,t.jsx)(g.Z.Item,{label:"\u793E\u4F1A\u4FE1\u7528\u4EE3\u7801",children:(0,t.jsx)("p",{children:v.organizationCode})}),(0,t.jsx)(g.Z.Item,{label:"\u672C\u4EBA\u59D3\u540D",children:(0,t.jsx)("p",{children:v.realName})}),(0,t.jsx)(g.Z.Item,{label:"\u672C\u4EBA\u8EAB\u4EFD\u8BC1\u6B63\u53CD\u9762",children:(0,t.jsxs)(de.Z,{children:[(0,t.jsx)(Y.Z,{width:100,height:100,src:v.identityUrlA}),(0,t.jsx)(Y.Z,{width:100,height:100,src:v.identityUrlB})]})}),(0,t.jsx)(g.Z.Item,{label:"\u62C5\u4EFB\u804C\u4F4D",children:(0,t.jsx)("p",{children:v.position})}),(0,t.jsx)(g.Z.Item,{label:"\u8054\u7CFB\u7535\u8BDD",children:(0,t.jsx)("p",{children:v.phone})})]})),(0,t.jsx)("h2",{children:"\u5BA1\u6838"}),(0,t.jsxs)(g.Z,(0,K.Z)((0,K.Z)({form:N},x),{},{children:[(0,t.jsx)(g.Z.Item,{name:"verifyStatus",label:"\u5BA1\u6838\u7ED3\u8BBA",rules:[{required:!0,message:"\u8BF7\u9009\u62E9"}],children:(0,t.jsxs)(Q.ZP.Group,{disabled:k,children:[(0,t.jsx)(Q.ZP,{value:3,children:"\u5BA1\u6838\u901A\u8FC7"}),(0,t.jsx)(Q.ZP,{value:2,children:"\u5BA1\u6838\u62D2\u7EDD"})]})}),h===2&&(0,t.jsx)(g.Z.Item,{name:"verifyComment",label:"\u539F\u56E0",children:(0,t.jsx)(R,{showCount:!0,disabled:k,rows:4,maxLength:50})}),(0,t.jsx)("div",{className:d()["actions-wrapper"],children:k?(0,t.jsx)(U.Z,{block:!0,disabled:!0,children:"\u5DF2\u5BA1\u6838"}):(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(U.Z,{onClick:B,children:"\u9000\u51FA"}),(0,t.jsx)(U.Z,{htmlType:"submit",loading:ne,onClick:ae,type:"primary",children:v.hasNext?"\u63D0\u4EA4\u5E76\u5BA1\u6838\u4E0B\u4E00\u4E2A":"\u63D0\u4EA4"})]})})]}))]})})},I=j},57315:function(ce,re,a){"use strict";a.d(re,{Z:function(){return te}});var F=a(22122),U=a(28481),r=a(67294),Q=a.t(r,2),K=a(38475),b=a(28991),de=a(94184),H=a.n(de),Y=a(15105);function Ne(){var e=(0,b.Z)({},Q);return e.useId}var ie=0;function p(){}var w=Ne(),Se=w?function(i){var n=w();return i||n}:function(i){var n=r.useState("ssr-id"),l=(0,U.Z)(n,2),d=l[0],m=l[1];return r.useEffect(function(){var y=ie;ie+=1,m("rc_unique_".concat(y))},[]),i||d},g=a(94999),X=a(64217),fe=a(5461);function me(e){var i=e.prefixCls,n=e.style,l=e.visible,d=e.maskProps,m=e.motionName;return r.createElement(fe.ZP,{key:"mask",visible:l,motionName:m,leavedClassName:"".concat(i,"-mask-hidden")},function(y,A){var t=y.className,x=y.style;return r.createElement("div",(0,F.Z)({ref:A,style:(0,b.Z)((0,b.Z)({},x),n),className:H()("".concat(i,"-mask"),t)},d))})}function z(e,i,n){var l=i;return!l&&n&&(l="".concat(e,"-").concat(n)),l}function se(e,i){var n=e["page".concat(i?"Y":"X","Offset")],l="scroll".concat(i?"Top":"Left");if(typeof n!="number"){var d=e.document;n=d.documentElement[l],typeof n!="number"&&(n=d.body[l])}return n}function q(e){var i=e.getBoundingClientRect(),n={left:i.left,top:i.top},l=e.ownerDocument,d=l.defaultView||l.parentWindow;return n.left+=se(d),n.top+=se(d,!0),n}var ve=r.memo(function(e){var i=e.children;return i},function(e,i){var n=i.shouldUpdate;return!n}),oe={width:0,height:0,overflow:"hidden",outline:"none"},he=r.forwardRef(function(e,i){var n=e.prefixCls,l=e.className,d=e.style,m=e.title,y=e.ariaId,A=e.footer,t=e.closable,x=e.closeIcon,R=e.onClose,j=e.children,I=e.bodyStyle,u=e.bodyProps,f=e.modalRender,C=e.onMouseDown,s=e.onMouseUp,N=e.holderRef,h=e.visible,T=e.forceRender,S=e.width,v=e.height,P=(0,r.useRef)(),D=(0,r.useRef)();r.useImperativeHandle(i,function(){return{focus:function(){var E;(E=P.current)===null||E===void 0||E.focus()},changeActive:function(E){var ne=document,O=ne.activeElement;E&&O===D.current?P.current.focus():!E&&O===P.current&&D.current.focus()}}});var $={};S!==void 0&&($.width=S),v!==void 0&&($.height=v);var k;A&&(k=r.createElement("div",{className:"".concat(n,"-footer")},A));var M;m&&(M=r.createElement("div",{className:"".concat(n,"-header")},r.createElement("div",{className:"".concat(n,"-title"),id:y},m)));var L;t&&(L=r.createElement("button",{type:"button",onClick:R,"aria-label":"Close",className:"".concat(n,"-close")},x||r.createElement("span",{className:"".concat(n,"-close-x")})));var B=r.createElement("div",{className:"".concat(n,"-content")},L,M,r.createElement("div",(0,F.Z)({className:"".concat(n,"-body"),style:I},u),j),k);return r.createElement("div",{key:"dialog-element",role:"dialog","aria-labelledby":m?y:null,"aria-modal":"true",ref:N,style:(0,b.Z)((0,b.Z)({},d),$),className:H()(n,l),onMouseDown:C,onMouseUp:s},r.createElement("div",{tabIndex:0,ref:P,style:oe,"aria-hidden":"true"}),r.createElement(ve,{shouldUpdate:h||T},f?f(B):B),r.createElement("div",{tabIndex:0,ref:D,style:oe,"aria-hidden":"true"}))}),_=he,le=r.forwardRef(function(e,i){var n=e.prefixCls,l=e.title,d=e.style,m=e.className,y=e.visible,A=e.forceRender,t=e.destroyOnClose,x=e.motionName,R=e.ariaId,j=e.onVisibleChanged,I=e.mousePosition,u=(0,r.useRef)(),f=r.useState(),C=(0,U.Z)(f,2),s=C[0],N=C[1],h={};s&&(h.transformOrigin=s);function T(){var S=q(u.current);N(I?"".concat(I.x-S.left,"px ").concat(I.y-S.top,"px"):"")}return r.createElement(fe.ZP,{visible:y,onVisibleChanged:j,onAppearPrepare:T,onEnterPrepare:T,forceRender:A,motionName:x,removeOnLeave:t,ref:u},function(S,v){var P=S.className,D=S.style;return r.createElement(_,(0,F.Z)({},e,{ref:i,title:l,ariaId:R,prefixCls:n,holderRef:v,style:(0,b.Z)((0,b.Z)((0,b.Z)({},D),d),h),className:H()(m,P)}))})});le.displayName="Content";var ee=le;function xe(e){var i=e.prefixCls,n=i===void 0?"rc-dialog":i,l=e.zIndex,d=e.visible,m=d===void 0?!1:d,y=e.keyboard,A=y===void 0?!0:y,t=e.focusTriggerAfterClose,x=t===void 0?!0:t,R=e.wrapStyle,j=e.wrapClassName,I=e.wrapProps,u=e.onClose,f=e.afterClose,C=e.transitionName,s=e.animation,N=e.closable,h=N===void 0?!0:N,T=e.mask,S=T===void 0?!0:T,v=e.maskTransitionName,P=e.maskAnimation,D=e.maskClosable,$=D===void 0?!0:D,k=e.maskStyle,M=e.maskProps,L=e.rootClassName,B=(0,r.useRef)(),V=(0,r.useRef)(),E=(0,r.useRef)(),ne=r.useState(m),O=(0,U.Z)(ne,2),ae=O[0],W=O[1],G=Se();function c(){(0,g.Z)(V.current,document.activeElement)||(B.current=document.activeElement)}function ye(){if(!(0,g.Z)(V.current,document.activeElement)){var Z;(Z=E.current)===null||Z===void 0||Z.focus()}}function o(Z){if(Z)ye();else{if(W(!1),S&&B.current&&x){try{B.current.focus({preventScroll:!0})}catch(Ae){}B.current=null}ae&&(f==null||f())}}function Ce(Z){u==null||u(Z)}var ue=(0,r.useRef)(!1),Ze=(0,r.useRef)(),be=function(){clearTimeout(Ze.current),ue.current=!0},pe=function(){Ze.current=setTimeout(function(){ue.current=!1})},Ee=null;$&&(Ee=function(Ae){ue.current?ue.current=!1:V.current===Ae.target&&Ce(Ae)});function Re(Z){if(A&&Z.keyCode===Y.Z.ESC){Z.stopPropagation(),Ce(Z);return}m&&Z.keyCode===Y.Z.TAB&&E.current.changeActive(!Z.shiftKey)}return(0,r.useEffect)(function(){m&&(W(!0),c())},[m]),(0,r.useEffect)(function(){return function(){clearTimeout(Ze.current)}},[]),r.createElement("div",(0,F.Z)({className:H()("".concat(n,"-root"),L)},(0,X.Z)(e,{data:!0})),r.createElement(me,{prefixCls:n,visible:S&&m,motionName:z(n,v,P),style:(0,b.Z)({zIndex:l},k),maskProps:M}),r.createElement("div",(0,F.Z)({tabIndex:-1,onKeyDown:Re,className:H()("".concat(n,"-wrap"),j),ref:V,onClick:Ee,style:(0,b.Z)((0,b.Z)({zIndex:l},R),{},{display:ae?null:"none"})},I),r.createElement(ee,(0,F.Z)({},e,{onMouseDown:be,onMouseUp:pe,ref:E,closable:h,ariaId:G,prefixCls:n,visible:m&&ae,onClose:Ce,onVisibleChanged:o,motionName:z(n,C,s)}))))}var J=function(i){var n=i.visible,l=i.getContainer,d=i.forceRender,m=i.destroyOnClose,y=m===void 0?!1:m,A=i.afterClose,t=r.useState(n),x=(0,U.Z)(t,2),R=x[0],j=x[1];return r.useEffect(function(){n&&j(!0)},[n]),!d&&y&&!R?null:r.createElement(K.Z,{open:n||d||R,autoDestroy:!1,getContainer:l,autoLock:n||R},r.createElement(xe,(0,F.Z)({},i,{destroyOnClose:y,afterClose:function(){A==null||A(),j(!1)}})))};J.displayName="Dialog";var ge=J,te=ge}}]);