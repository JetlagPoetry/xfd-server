(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[955],{79027:function(L){L.exports={"order-detail-container":"order-detail-container___LGnty","operate-btn":"operate-btn___uyekI"}},31134:function(L,C,r){"use strict";r.d(C,{Zv:function(){return l},_y:function(){return j}});var t=r(32059),m=r(975),I,d,S,o,f,Z={WAIT_DELIVER:3,WAIT_RECEIVE:4,RECEIVED:5,AFTER_SALE:6},q=(I={},(0,t.Z)(I,Z.WAIT_DELIVER,"\u4EE3\u53D1\u8D27"),(0,t.Z)(I,Z.WAIT_RECEIVE,"\u5F85\u6536\u8D27/\u5DF2\u53D1\u8D27"),(0,t.Z)(I,Z.RECEIVED," \u5DF2\u7B7E\u6536"),I),D={UNKNOWN:0,APPROVED:1,REJECTED:2,CLEAR:3},K=(d={},(0,t.Z)(d,D.UNKNOWN,"\u672A\u77E5"),(0,t.Z)(d,D.APPROVED,"\u5BA1\u6838\u901A\u8FC7"),(0,t.Z)(d,D.REJECTED,"\u5BA1\u6838\u5931\u8D25"),(0,t.Z)(d,D.CLEAR,"\u6E05\u96F6"),d),c={UNKNOWN:0,SUBMITTED:1,REJECTED:2,SUCCESS:3},$=(S={},(0,t.Z)(S,c.UNKNOWN,"\u672A\u77E5"),(0,t.Z)(S,c.SUBMITTED,"\u5DF2\u63D0\u4EA4"),(0,t.Z)(S,c.REJECTED,"\u88AB\u62D2\u7EDD"),(0,t.Z)(S,c.SUCCESS,"\u6210\u529F"),S),a={UNKNOWN:0,CUSTOMER:1,SUPPLIER:2,BUYER:3,ADMIN:4,ROOT:5},A=(o={},(0,t.Z)(o,a.UNKNOWN,"\u672A\u77E5"),(0,t.Z)(o,a.CUSTOMER,"\u5BA2\u6237"),(0,t.Z)(o,a.SUPPLIER,"\u4F9B\u5E94\u5546"),(0,t.Z)(o,a.BUYER,"\u4E70\u5BB6"),(0,t.Z)(o,a.ADMIN,"\u7BA1\u7406\u5458"),(0,t.Z)(o,a.ROOT,"\u8D85\u7EA7\u7BA1\u7406\u5458"),o),p={APPLICATION:0,SPEND:1,EXPIRED:2,QUIT:3,CANCEL:4},s=(f={},(0,t.Z)(f,p.APPLICATION,"\u65B0\u589E"),(0,t.Z)(f,p.SPEND,"\u6D88\u8D39"),(0,t.Z)(f,p.EXPIRED,"\u8FC7\u671F"),(0,t.Z)(f,p.QUIT,"\u79BB\u804C"),(0,t.Z)(f,p.CANCEL,"\u6E05\u96F6"),f),l=new Map([[m.i4.Buyers,"\u91C7\u8D2D\u5546"],[m.i4.Supplier,"\u4F9B\u8D27\u5546"],[m.i4.Management,"\u7BA1\u7406\u5458"],[m.i4.Super_Management,"\u8D85\u7EA7\u7BA1\u7406\u5458"]]),j=new Map([[3,"\u5F85\u53D1\u8D27"],[4,"\u5F85\u6536\u8D27"],[5,"\u5DF2\u7B7E\u6536"],[6,"\u552E\u540E/\u7ED3\u675F"]])},66754:function(L,C,r){"use strict";r.r(C),r.d(C,{default:function(){return ye}});var t=r(57663),m=r(71577),I=r(87593),d=r(37636),S=r(13062),o=r(71230),f=r(89032),Z=r(15746),q=r(12968),D=r(26141),K=r(48736),c=r(27049),$=r(98858),a=r(4914),A=r(90636),p=r(3182),s=r(2824),l=r(67294),j=r(27484),n=r.n(j),le=r(95916),ie=r(99871),N=r(1987),xe=r(71194),de=r(48889),ee=r(11849),Ne=r(88983),X=r(47933),Be=r(34792),oe=r(48086),Ue=r(9715),B=r(71481),Me=r(43358),re=r(34041),be=r(47673),ce=r(60345),F=r(975),e=r(85893),me={labelCol:{span:6},wrapperCol:{span:16}},ve=ce.Z.TextArea,H=re.Z.Option,pe=new Map([[3,"\u7ED3\u675F\u8BA2\u5355"],[4,"\u552E\u540E\u8BB0\u5F55"],[5,"\u552E\u540E\u8BB0\u5F55"]]),Ee=new Map([[3,"\u6C9F\u901A\u5907\u6CE8"],[4,"\u552E\u540E\u8BB0\u5F55"],[5,"\u552E\u540E\u8BB0\u5F55"]]),E;(function(g){g[g.Apply_Exchange=1]="Apply_Exchange",g[g.Apply_Refund=2]="Apply_Refund",g[g.Other=3]="Other"})(E||(E={}));var he=function(h){var J=B.Z.useForm(),Y=(0,s.Z)(J,1),P=Y[0],z=function(){var M=(0,p.Z)((0,A.Z)().mark(function v(){var R,O;return(0,A.Z)().wrap(function(u){for(;;)switch(u.prev=u.next){case 0:return u.prev=0,u.next=3,P.validateFields();case 3:R=u.sent,O={queryOrderID:Number(h.id),reason:R.reason},u.t0=h.orderStatus,u.next=u.t0===F.iF.To_Be_Delivered?8:u.t0===F.iF.Awaiting_Receipt||u.t0===F.iF.Have_Been_Received?12:27;break;case 8:return Object.assign(O,{returnPointType:R.returnPointType}),u.next=11,N.qJ(O);case 11:return u.abrupt("break",28);case 12:if(R.afterSalesType!==E.Apply_Refund){u.next=18;break}return Object.assign(O,{returnPointType:R.returnPointType,afterSaleType:E.Apply_Refund}),u.next=16,N.ui(O);case 16:u.next=26;break;case 18:if(R.afterSalesType!==E.Other){u.next=24;break}return Object.assign(O,{returnPointType:R.returnPointType,afterSaleType:E.Other}),u.next=22,N.ui(O);case 22:u.next=26;break;case 24:return u.next=26,N.NN(O);case 26:return u.abrupt("break",28);case 27:return u.abrupt("break",28);case 28:oe.ZP.success("\u64CD\u4F5C\u6210\u529F"),h.onOk(),u.next=34;break;case 32:u.prev=32,u.t1=u.catch(0);case 34:case"end":return u.stop()}},v,null,[[0,32]])}));return function(){return M.apply(this,arguments)}}(),Q=(0,l.useState)(!1),U=(0,s.Z)(Q,2),x=U[0],W=U[1],w=function(v){v===E.Apply_Refund||v===E.Other?W(!0):W(!1)};return(0,e.jsx)(de.Z,{centered:!0,onOk:z,onCancel:h.onCancel,title:pe.get(h.orderStatus),width:600,open:h.open,bodyStyle:{padding:"24px"},children:(0,e.jsxs)(B.Z,(0,ee.Z)((0,ee.Z)({form:P,initialValues:{afterSalesType:E.Apply_Exchange,returnPointType:2}},me),{},{children:[[F.iF.Awaiting_Receipt,F.iF.Have_Been_Received].includes(h.orderStatus)&&(0,e.jsx)(e.Fragment,{children:(0,e.jsx)(B.Z.Item,{label:"\u552E\u540E\u7C7B\u578B",name:"afterSalesType",children:(0,e.jsxs)(re.Z,{onChange:w,children:[(0,e.jsx)(H,{value:E.Apply_Exchange,children:"\u6362\u8D27"}),(0,e.jsx)(H,{value:E.Apply_Refund,children:"\u9000\u8D27\u9000\u6B3E"}),(0,e.jsx)(H,{value:E.Other,children:"\u5176\u4ED6"})]})})}),(h.orderStatus===F.iF.To_Be_Delivered||x)&&(0,e.jsx)(B.Z.Item,{name:"returnPointType",label:"\u5904\u7406\u65B9\u5F0F",children:(0,e.jsxs)(X.ZP.Group,{children:[(0,e.jsx)(X.ZP,{value:2,children:"\u8FD4\u8FD8\u79EF\u5206"}),(0,e.jsx)(X.ZP,{value:1,children:"\u4E0D\u8FD4\u8FD8\u79EF\u5206"})]})}),(0,e.jsx)(B.Z.Item,{name:"reason",label:Ee.get(h.orderStatus),rules:[{required:!0,message:"\u8BF7\u8F93\u5165"}],children:(0,e.jsx)(ve,{maxLength:200})})]}))})},Ze=he,Te=r(79027),G=r.n(Te),fe=r(31134),De=new Map([[3,"\u7ED3\u675F\u8BA2\u5355"],[4,"\u552E\u540E\u8BB0\u5F55"],[5,"\u552E\u540E\u8BB0\u5F55"]]),Oe=function(){var h=(0,ie.D)("id")||"",J=(0,l.useState)(3),Y=(0,s.Z)(J,2),P=Y[0],z=Y[1],Q=(0,l.useState)({}),U=(0,s.Z)(Q,2),x=U[0],W=U[1],w=(0,l.useState)({}),M=(0,s.Z)(w,2),v=M[0],R=M[1],O=(0,l.useState)({}),k=(0,s.Z)(O,2),u=k[0],Se=k[1],je=(0,l.useState)({}),ae=(0,s.Z)(je,2),V=ae[0],Re=ae[1],Ie=(0,l.useState)([]),te=(0,s.Z)(Ie,2),ge=te[0],Pe=te[1],ne=function(){var T=(0,p.Z)((0,A.Z)().mark(function b(){var i;return(0,A.Z)().wrap(function(y){for(;;)switch(y.prev=y.next){case 0:return y.prev=0,y.next=3,N.dz(h);case 3:i=y.sent,z(i==null?void 0:i.order.status),W(i==null?void 0:i.order),R(i==null?void 0:i.goods),Se(i==null?void 0:i.buyer),Re(i==null?void 0:i.seller),Pe(i.orderRecord),console.log(i.orderRecord),y.next=15;break;case 13:y.prev=13,y.t0=y.catch(0);case 15:case"end":return y.stop()}},b,null,[[0,13]])}));return function(){return T.apply(this,arguments)}}();(0,l.useEffect)(function(){ne()},[]);var Ae=(0,l.useState)(!1),ue=(0,s.Z)(Ae,2),se=ue[0],_=ue[1],Ce=function(){_(!0)},Fe=function(){_(!1),ne()};return(0,e.jsxs)(le.ZP,{children:[(0,e.jsxs)("div",{className:G()["order-detail-container"],children:[(0,e.jsx)("h1",{children:fe._y.get(P)}),(0,e.jsxs)(a.Z,{column:2,title:"\u8BA2\u5355\u4FE1\u606F",children:[(0,e.jsx)(a.Z.Item,{label:"\u8BA2\u5355\u7F16\u53F7",children:x.orderSn}),(0,e.jsx)(a.Z.Item,{label:"\u4E0B\u5355\u65F6\u95F4",children:n()(x.createdAt).format("YYYY-MM-DD HH:mm")}),(0,e.jsxs)(a.Z.Item,{label:"\u8BA2\u5355\u91D1\u989D",children:["\xA5 ",x.totalPrice]}),(0,e.jsx)(a.Z.Item,{label:"\u4ED8\u6B3E\u65F6\u95F4",children:n()(x.payedAt).format("YYYY-MM-DD HH:mm")})]}),(0,e.jsx)(c.Z,{}),(0,e.jsxs)(a.Z,{column:2,title:"\u5546\u54C1\u4FE1\u606F",children:[(0,e.jsx)(a.Z.Item,{label:"\u5546\u54C1\u540D\u79F0",children:v.name}),(0,e.jsx)(a.Z.Item,{label:"\u5546\u54C1\u7F16\u53F7",children:v.goodsID}),(0,e.jsx)(a.Z.Item,{label:"\u8D2D\u4E70\u89C4\u683C",children:v.productAttr}),(0,e.jsx)(a.Z.Item,{label:"\u8D2D\u4E70\u6570\u91CF",children:v.quantity}),(0,e.jsxs)(a.Z.Item,{label:"\u5546\u54C1\u5355\u4EF7",children:["\xA5 ",v.unitPrice]}),(0,e.jsx)(a.Z.Item,{label:"\u5546\u54C1\u8FD0\u8D39",children:v.postPrice}),(0,e.jsx)(a.Z.Item,{label:"\u5546\u54C1\u4E3B\u56FE",children:(0,e.jsx)(D.Z,{width:100,src:v.image})})]}),(0,e.jsx)(c.Z,{}),(0,e.jsxs)(a.Z,{column:2,title:"\u4E70\u5BB6\u4FE1\u606F",children:[(0,e.jsx)(a.Z.Item,{label:"\u7528\u6237\u8D26\u53F7",children:u.userName}),(0,e.jsx)(a.Z.Item,{label:"\u7528\u6237\u7535\u8BDD",children:u.userPhone}),(0,e.jsx)(a.Z.Item,{label:"\u6240\u5C5E\u516C\u53F8",children:u.userOrganizationName}),(0,e.jsx)(a.Z.Item,{label:"\u6536\u4EF6\u59D3\u540D",children:u.singerName}),(0,e.jsx)(a.Z.Item,{label:"\u6536\u4EF6\u7535\u8BDD",children:u.singerPhone}),(0,e.jsx)(a.Z.Item,{label:"\u6536\u4EF6\u5730\u5740",children:u.singerAddr})]}),(0,e.jsx)(c.Z,{}),(0,e.jsxs)(a.Z,{column:2,title:"\u5356\u5BB6\u4FE1\u606F",children:[(0,e.jsx)(a.Z.Item,{label:"\u7528\u6237\u8D26\u53F7",children:V.name}),(0,e.jsx)(a.Z.Item,{label:"\u5356\u5BB6\u7535\u8BDD",children:V.phone}),(0,e.jsx)(a.Z.Item,{label:"\u8BA4\u8BC1\u4F01\u4E1A",children:V.organizationName}),(0,e.jsx)(a.Z.Item,{label:"\u8BA4\u8BC1\u804C\u4F4D",children:V.position})]}),(0,e.jsx)(c.Z,{}),(0,e.jsx)(a.Z,{column:2,title:"\u8BA2\u5355\u8BB0\u5F55"}),(0,e.jsx)(d.Z,{children:ge.map(function(T){return(0,e.jsx)(d.Z.Item,{children:(0,e.jsxs)(o.Z,{children:[(0,e.jsx)(Z.Z,{span:3,children:T.recordName}),(0,e.jsx)(Z.Z,{span:4,children:n()(T.recordTime).format("YYYY-MM-DD HH:mm")}),(0,e.jsx)(Z.Z,{span:8,children:T.records&&T.records.map(function(b){return(0,e.jsx)("p",{children:b},b)})})]})},T.recordTime)})}),(0,e.jsx)(c.Z,{}),[3,4,5].includes(P)&&(0,e.jsx)("div",{className:G()["operate-btn"],children:(0,e.jsx)(m.Z,{block:!0,type:"primary",onClick:Ce,children:De.get(P)})}),[6].includes(P)&&(0,e.jsx)("div",{className:G()["operate-btn"],children:(0,e.jsx)(m.Z,{block:!0,disabled:!0,type:"primary",children:"\u8BA2\u5355\u5DF2\u7ED3\u675F"})})]}),se&&(0,e.jsx)(Ze,{id:h,open:se,onCancel:function(){return _(!1)},orderStatus:P,onOk:function(){return Fe()}})]})},ye=Oe},1987:function(L,C,r){"use strict";r.d(C,{dz:function(){return S},Fw:function(){return f},qJ:function(){return K},NN:function(){return $},ui:function(){return A}});var t=r(90636),m=r(3182),I=r(99871),d=r(636);function S(s){return o.apply(this,arguments)}function o(){return o=(0,m.Z)((0,t.Z)().mark(function s(l){return(0,t.Z)().wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return n.abrupt("return",(0,d.Z)("/api/v1/order/getOrderDetail?queryOrderID=".concat(l)));case 1:case"end":return n.stop()}},s)})),o.apply(this,arguments)}function f(s){return Z.apply(this,arguments)}function Z(){return Z=(0,m.Z)((0,t.Z)().mark(function s(l){return(0,t.Z)().wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return n.abrupt("return",(0,d.Z)("/api/v1/order/getOrderList?".concat((0,I.R)(l))));case 1:case"end":return n.stop()}},s)})),Z.apply(this,arguments)}function q(s){return D.apply(this,arguments)}function D(){return D=_asyncToGenerator(_regeneratorRuntime().mark(function s(l){return _regeneratorRuntime().wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return n.abrupt("return",request("/api/v1/order/exportOrder?".concat(objectToUrlParams(l)),{responseType:"arraybuffer"}));case 1:case"end":return n.stop()}},s)})),D.apply(this,arguments)}function K(s){return c.apply(this,arguments)}function c(){return c=(0,m.Z)((0,t.Z)().mark(function s(l){return(0,t.Z)().wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return n.abrupt("return",(0,d.Z)("api/v1/order/closeOrder",{method:"POST",data:l}));case 1:case"end":return n.stop()}},s)})),c.apply(this,arguments)}function $(s){return a.apply(this,arguments)}function a(){return a=(0,m.Z)((0,t.Z)().mark(function s(l){return(0,t.Z)().wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return n.abrupt("return",(0,d.Z)("api/v1/order/applyExchange",{method:"POST",data:l}));case 1:case"end":return n.stop()}},s)})),a.apply(this,arguments)}function A(s){return p.apply(this,arguments)}function p(){return p=(0,m.Z)((0,t.Z)().mark(function s(l){return(0,t.Z)().wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return n.abrupt("return",(0,d.Z)("api/v1/order/applyRefund",{method:"POST",data:l}));case 1:case"end":return n.stop()}},s)})),p.apply(this,arguments)}}}]);
