(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[818],{76504:function(W){W.exports={actions:"actions___VoCkO","search-wrapper":"search-wrapper___2CWWw","table-wrapper":"table-wrapper___I8xyZ"}},45842:function(W,Z,a){"use strict";a.r(Z);var m=a(8963),I=a(35008),j=a(88983),C=a(47933),B=a(13062),T=a(71230),Y=a(49111),y=a(19650),ee=a(89032),E=a(15746),K=a(47673),w=a(60345),te=a(34792),H=a(48086),U=a(90636),A=a(3182),O=a(9715),f=a(71481),v=a(11849),o=a(57663),z=a(71577),R=a(2824),d=a(67294),_=a(12666),X=a(27484),re=a.n(X),u=a(95916),n=a(99871),e=a(1987),s=a(76504),r=a.n(s),t=a(85893),i=function(){var h=[{label:"\u5168\u90E8",value:""},{label:"\u5F85\u53D1\u8D27",value:"3"},{label:"\u5F85\u6536\u8D27",value:"4"},{label:"\u5DF2\u7B7E\u6536",value:"5"},{label:"\u552E\u540E/\u7ED3\u675F",value:"6"}],D=(0,d.useState)(""),g=(0,R.Z)(D,2),$=g[0],F=g[1],N=function(l){var p=l.target.value;F(p)},J=[{title:"\u64CD\u4F5C",dataIndex:"queryOrderID",key:"queryOrderID",render:function(l){return(0,t.jsx)(z.Z,{type:"link",href:"?id=".concat(l,"#/order/list/detail"),children:"\u67E5\u770B"})}},{title:"\u8BA2\u5355\u72B6\u6001",dataIndex:"status",key:"status",render:function(l){var p=h.find(function(L){return L.value===l.toString()});return p?p.label:""}},{title:"\u8BA2\u5355\u7F16\u53F7",dataIndex:"orderSn",key:"orderSn"},{title:"\u5546\u54C1\u540D\u79F0",dataIndex:"name",key:"name"},{title:"\u5546\u6237\u540D\u79F0",dataIndex:"supplierOrganizationName",key:"supplierOrganizationName"},{title:"\u652F\u4ED8\u91D1\u989D",dataIndex:"verifyTime",key:"verifyTime",render:function(l,p){var L=parseFloat(p.totalPrice)+parseFloat(p.postPrice);return"\xA5".concat(L.toFixed(2))}},{title:"\u5458\u5DE5\u7535\u8BDD",dataIndex:"consumerUserPhone",key:"consumerUserPhone"},{title:"\u5458\u5DE5\u7EC4\u7EC7",dataIndex:"consumerOrganizationName",key:"consumerOrganizationName"},{title:"\u652F\u4ED8\u65F6\u95F4",dataIndex:"payedAt",key:"payedAt",render:function(l){return re()(l).format("YYYY-MM-DD HH:mm")}}],k=(0,d.useState)([]),P=(0,R.Z)(k,2),x=P[0],V=P[1],q=(0,d.useState)(!1),ue=(0,R.Z)(q,2),pe=ue[0],ne=ue[1],me=(0,d.useState)({current:1,pageSize:10,showSizeChanger:!0,showQuickJumper:!0,showTotal:function(l){return"\u603B\u5171 ".concat(l," \u6761")}}),ie=(0,R.Z)(me,2),oe=ie[0],_e=ie[1],ve=function(l){_e({current:l.current||1,pageSize:l.pageSize||10}),de((0,v.Z)((0,v.Z)({},G),{},{pageNum:l.current||1,pageSize:l.pageSize||10}))},Oe=(0,d.useState)([]),le=(0,R.Z)(Oe,2),ae=le[0],De=le[1],Ee=function(l){De(l)},ge={selectedRowKeys:ae,onChange:Ee},Me=f.Z.useForm(),ye=(0,R.Z)(Me,1),se=ye[0],Pe=(0,d.useState)({orderSn:"",goodName:"",pageNum:1,pageSize:10,supplierOrganizationName:"",userPhone:"",userOrganizationName:""}),ce=(0,R.Z)(Pe,2),G=ce[0],de=ce[1],fe=function(){var S=(0,A.Z)((0,U.Z)().mark(function l(){var p;return(0,U.Z)().wrap(function(M){for(;;)switch(M.prev=M.next){case 0:return M.next=2,se.getFieldsValue();case 2:p=M.sent,de((0,v.Z)((0,v.Z)({},G),{},{pageNum:1,orderSn:p.orderSn||"",goodName:p.goodName||"",supplierOrganizationName:p.supplierOrganizationName||"",userPhone:p.userPhone||"",userOrganizationName:p.userOrganizationName||""}));case 4:case"end":return M.stop()}},l)}));return function(){return S.apply(this,arguments)}}(),$e=function(){var S=(0,A.Z)((0,U.Z)().mark(function l(){return(0,U.Z)().wrap(function(L){for(;;)switch(L.prev=L.next){case 0:return L.next=2,se.resetFields();case 2:fe();case 3:case"end":return L.stop()}},l)}));return function(){return S.apply(this,arguments)}}(),Se=function(){var S=(0,A.Z)((0,U.Z)().mark(function l(){var p;return(0,U.Z)().wrap(function(M){for(;;)switch(M.prev=M.next){case 0:return M.prev=0,ne(!0),M.next=4,e.Fw((0,v.Z)((0,v.Z)({},G),{},{status:$,pageNum:G.pageNum,pageSize:G.pageSize}));case 4:p=M.sent,p&&(V(p.list),_e((0,v.Z)((0,v.Z)({},oe),{},{current:G.pageNum,total:p.totalNum}))),ne(!1),M.next=12;break;case 9:M.prev=9,M.t0=M.catch(0),ne(!1);case 12:case"end":return M.stop()}},l,null,[[0,9]])}));return function(){return S.apply(this,arguments)}}();(0,d.useEffect)(function(){Se()},[$,JSON.stringify(G)]);var Ze=function(){var S=(0,A.Z)((0,U.Z)().mark(function l(){var p,L,M,he,Q;return(0,U.Z)().wrap(function(b){for(;;)switch(b.prev=b.next){case 0:if(!(ae.length<1)){b.next=3;break}return H.ZP.warn("\u8BF7\u52FE\u9009\u8981\u5BFC\u51FA\u7684\u8BA2\u5355"),b.abrupt("return");case 3:return b.prev=3,p=_.Z.create({baseURL:"/",timeout:3e4,withCredentials:!1}),L=localStorage.getItem("token"),b.next=8,p("/api/v1/order/exportOrder?".concat((0,n.R)({ids:ae})),{responseType:"arraybuffer",headers:{Authorization:"".concat(L)}});case 8:M=b.sent,he=window.URL.createObjectURL(new Blob([M.data],{type:"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"})),Q=document.createElement("a"),Q.style.display="none",Q.href=he,Q.setAttribute("download","excel.xlsx"),document.body.appendChild(Q),Q.click(),document.body.removeChild(Q),b.next=21;break;case 19:b.prev=19,b.t0=b.catch(3);case 21:case"end":return b.stop()}},l,null,[[3,19]])}));return function(){return S.apply(this,arguments)}}();return(0,t.jsxs)(u.ZP,{children:[(0,t.jsx)("div",{className:r()["search-wrapper"],children:(0,t.jsx)(f.Z,{form:se,labelCol:{span:8},wrapperCol:{span:16},children:(0,t.jsxs)(T.Z,{children:[(0,t.jsx)(E.Z,{span:6,children:(0,t.jsx)(f.Z.Item,{label:"\u8BA2\u5355\u7F16\u53F7",name:"orderSn",rules:[{required:!1}],children:(0,t.jsx)(w.Z,{})})}),(0,t.jsx)(E.Z,{span:6,children:(0,t.jsx)(f.Z.Item,{label:"\u5546\u54C1\u540D\u79F0",name:"goodName",rules:[{required:!1}],children:(0,t.jsx)(w.Z,{})})}),(0,t.jsx)(E.Z,{span:6,children:(0,t.jsx)(f.Z.Item,{label:"\u5546\u6237\u540D\u79F0",name:"supplierOrganizationName",rules:[{required:!1}],children:(0,t.jsx)(w.Z,{})})}),(0,t.jsx)(E.Z,{span:6,children:(0,t.jsx)(f.Z.Item,{label:"\u5458\u5DE5\u7535\u8BDD",name:"userPhone",rules:[{required:!1}],children:(0,t.jsx)(w.Z,{})})}),(0,t.jsx)(E.Z,{span:6,children:(0,t.jsx)(f.Z.Item,{label:"\u5458\u5DE5\u7EC4\u7EC7",name:"userOrganizationName",rules:[{required:!1}],children:(0,t.jsx)(w.Z,{})})}),(0,t.jsx)(E.Z,{offset:12,span:6,style:{textAlign:"right"},children:(0,t.jsxs)(y.Z,{children:[(0,t.jsx)(z.Z,{type:"primary",onClick:fe,children:"\u67E5\u8BE2"}),(0,t.jsx)(z.Z,{onClick:$e,children:"\u91CD\u7F6E"})]})})]})})}),(0,t.jsxs)("div",{className:r()["table-wrapper"],children:[(0,t.jsxs)("div",{className:r().actions,children:[(0,t.jsx)(C.ZP.Group,{options:h,optionType:"button",buttonStyle:"solid",onChange:N,value:$}),(0,t.jsx)(z.Z,{type:"primary",onClick:Ze,children:"\u5BFC\u51FA"})]}),(0,t.jsx)(I.Z,{columns:J,dataSource:x,loading:pe,onChange:ve,pagination:oe,rowKey:"queryOrderID",rowSelection:ge,scroll:{x:"max-content"}})]})]})};Z.default=i},1987:function(W,Z,a){"use strict";a.d(Z,{dz:function(){return B},Fw:function(){return Y},qJ:function(){return K},NN:function(){return te},ui:function(){return U}});var m=a(90636),I=a(3182),j=a(99871),C=a(636);function B(O){return T.apply(this,arguments)}function T(){return T=(0,I.Z)((0,m.Z)().mark(function O(f){return(0,m.Z)().wrap(function(o){for(;;)switch(o.prev=o.next){case 0:return o.abrupt("return",(0,C.Z)("/api/v1/order/getOrderDetail?queryOrderID=".concat(f)));case 1:case"end":return o.stop()}},O)})),T.apply(this,arguments)}function Y(O){return y.apply(this,arguments)}function y(){return y=(0,I.Z)((0,m.Z)().mark(function O(f){return(0,m.Z)().wrap(function(o){for(;;)switch(o.prev=o.next){case 0:return o.abrupt("return",(0,C.Z)("/api/v1/order/getOrderList?".concat((0,j.R)(f))));case 1:case"end":return o.stop()}},O)})),y.apply(this,arguments)}function ee(O){return E.apply(this,arguments)}function E(){return E=_asyncToGenerator(_regeneratorRuntime().mark(function O(f){return _regeneratorRuntime().wrap(function(o){for(;;)switch(o.prev=o.next){case 0:return o.abrupt("return",request("/api/v1/order/exportOrder?".concat(objectToUrlParams(f)),{responseType:"arraybuffer"}));case 1:case"end":return o.stop()}},O)})),E.apply(this,arguments)}function K(O){return w.apply(this,arguments)}function w(){return w=(0,I.Z)((0,m.Z)().mark(function O(f){return(0,m.Z)().wrap(function(o){for(;;)switch(o.prev=o.next){case 0:return o.abrupt("return",(0,C.Z)("api/v1/order/closeOrder",{method:"POST",data:f}));case 1:case"end":return o.stop()}},O)})),w.apply(this,arguments)}function te(O){return H.apply(this,arguments)}function H(){return H=(0,I.Z)((0,m.Z)().mark(function O(f){return(0,m.Z)().wrap(function(o){for(;;)switch(o.prev=o.next){case 0:return o.abrupt("return",(0,C.Z)("api/v1/order/applyExchange",{method:"POST",data:f}));case 1:case"end":return o.stop()}},O)})),H.apply(this,arguments)}function U(O){return A.apply(this,arguments)}function A(){return A=(0,I.Z)((0,m.Z)().mark(function O(f){return(0,m.Z)().wrap(function(o){for(;;)switch(o.prev=o.next){case 0:return o.abrupt("return",(0,C.Z)("api/v1/order/applyRefund",{method:"POST",data:f}));case 1:case"end":return o.stop()}},O)})),A.apply(this,arguments)}},15746:function(W,Z,a){"use strict";var m=a(21584);Z.Z=m.Z},89032:function(W,Z,a){"use strict";var m=a(38663),I=a.n(m),j=a(6999)},71230:function(W,Z,a){"use strict";var m=a(92820);Z.Z=m.Z},13062:function(W,Z,a){"use strict";var m=a(38663),I=a.n(m),j=a(6999)},27484:function(W){(function(Z,a){W.exports=a()})(this,function(){"use strict";var Z=1e3,a=6e4,m=36e5,I="millisecond",j="second",C="minute",B="hour",T="day",Y="week",y="month",ee="quarter",E="year",K="date",w="Invalid Date",te=/^(\d{4})[-/]?(\d{1,2})?[-/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$/,H=/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g,U={name:"en",weekdays:"Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),months:"January_February_March_April_May_June_July_August_September_October_November_December".split("_"),ordinal:function(u){var n=["th","st","nd","rd"],e=u%100;return"["+u+(n[(e-20)%10]||n[e]||n[0])+"]"}},A=function(u,n,e){var s=String(u);return!s||s.length>=n?u:""+Array(n+1-s.length).join(e)+u},O={s:A,z:function(u){var n=-u.utcOffset(),e=Math.abs(n),s=Math.floor(e/60),r=e%60;return(n<=0?"+":"-")+A(s,2,"0")+":"+A(r,2,"0")},m:function u(n,e){if(n.date()<e.date())return-u(e,n);var s=12*(e.year()-n.year())+(e.month()-n.month()),r=n.clone().add(s,y),t=e-r<0,i=n.clone().add(s+(t?-1:1),y);return+(-(s+(e-r)/(t?r-i:i-r))||0)},a:function(u){return u<0?Math.ceil(u)||0:Math.floor(u)},p:function(u){return{M:y,y:E,w:Y,d:T,D:K,h:B,m:C,s:j,ms:I,Q:ee}[u]||String(u||"").toLowerCase().replace(/s$/,"")},u:function(u){return u===void 0}},f="en",v={};v[f]=U;var o="$isDayjsObject",z=function(u){return u instanceof X||!(!u||!u[o])},R=function u(n,e,s){var r;if(!n)return f;if(typeof n=="string"){var t=n.toLowerCase();v[t]&&(r=t),e&&(v[t]=e,r=t);var i=n.split("-");if(!r&&i.length>1)return u(i[0])}else{var c=n.name;v[c]=n,r=c}return!s&&r&&(f=r),r||!s&&f},d=function(u,n){if(z(u))return u.clone();var e=typeof n=="object"?n:{};return e.date=u,e.args=arguments,new X(e)},_=O;_.l=R,_.i=z,_.w=function(u,n){return d(u,{locale:n.$L,utc:n.$u,x:n.$x,$offset:n.$offset})};var X=function(){function u(e){this.$L=R(e.locale,null,!0),this.parse(e),this.$x=this.$x||e.x||{},this[o]=!0}var n=u.prototype;return n.parse=function(e){this.$d=function(s){var r=s.date,t=s.utc;if(r===null)return new Date(NaN);if(_.u(r))return new Date;if(r instanceof Date)return new Date(r);if(typeof r=="string"&&!/Z$/i.test(r)){var i=r.match(te);if(i){var c=i[2]-1||0,h=(i[7]||"0").substring(0,3);return t?new Date(Date.UTC(i[1],c,i[3]||1,i[4]||0,i[5]||0,i[6]||0,h)):new Date(i[1],c,i[3]||1,i[4]||0,i[5]||0,i[6]||0,h)}}return new Date(r)}(e),this.init()},n.init=function(){var e=this.$d;this.$y=e.getFullYear(),this.$M=e.getMonth(),this.$D=e.getDate(),this.$W=e.getDay(),this.$H=e.getHours(),this.$m=e.getMinutes(),this.$s=e.getSeconds(),this.$ms=e.getMilliseconds()},n.$utils=function(){return _},n.isValid=function(){return this.$d.toString()!==w},n.isSame=function(e,s){var r=d(e);return this.startOf(s)<=r&&r<=this.endOf(s)},n.isAfter=function(e,s){return d(e)<this.startOf(s)},n.isBefore=function(e,s){return this.endOf(s)<d(e)},n.$g=function(e,s,r){return _.u(e)?this[s]:this.set(r,e)},n.unix=function(){return Math.floor(this.valueOf()/1e3)},n.valueOf=function(){return this.$d.getTime()},n.startOf=function(e,s){var r=this,t=!!_.u(s)||s,i=_.p(e),c=function(k,P){var x=_.w(r.$u?Date.UTC(r.$y,P,k):new Date(r.$y,P,k),r);return t?x:x.endOf(T)},h=function(k,P){return _.w(r.toDate()[k].apply(r.toDate("s"),(t?[0,0,0,0]:[23,59,59,999]).slice(P)),r)},D=this.$W,g=this.$M,$=this.$D,F="set"+(this.$u?"UTC":"");switch(i){case E:return t?c(1,0):c(31,11);case y:return t?c(1,g):c(0,g+1);case Y:var N=this.$locale().weekStart||0,J=(D<N?D+7:D)-N;return c(t?$-J:$+(6-J),g);case T:case K:return h(F+"Hours",0);case B:return h(F+"Minutes",1);case C:return h(F+"Seconds",2);case j:return h(F+"Milliseconds",3);default:return this.clone()}},n.endOf=function(e){return this.startOf(e,!1)},n.$set=function(e,s){var r,t=_.p(e),i="set"+(this.$u?"UTC":""),c=(r={},r[T]=i+"Date",r[K]=i+"Date",r[y]=i+"Month",r[E]=i+"FullYear",r[B]=i+"Hours",r[C]=i+"Minutes",r[j]=i+"Seconds",r[I]=i+"Milliseconds",r)[t],h=t===T?this.$D+(s-this.$W):s;if(t===y||t===E){var D=this.clone().set(K,1);D.$d[c](h),D.init(),this.$d=D.set(K,Math.min(this.$D,D.daysInMonth())).$d}else c&&this.$d[c](h);return this.init(),this},n.set=function(e,s){return this.clone().$set(e,s)},n.get=function(e){return this[_.p(e)]()},n.add=function(e,s){var r,t=this;e=Number(e);var i=_.p(s),c=function(g){var $=d(t);return _.w($.date($.date()+Math.round(g*e)),t)};if(i===y)return this.set(y,this.$M+e);if(i===E)return this.set(E,this.$y+e);if(i===T)return c(1);if(i===Y)return c(7);var h=(r={},r[C]=a,r[B]=m,r[j]=Z,r)[i]||1,D=this.$d.getTime()+e*h;return _.w(D,this)},n.subtract=function(e,s){return this.add(-1*e,s)},n.format=function(e){var s=this,r=this.$locale();if(!this.isValid())return r.invalidDate||w;var t=e||"YYYY-MM-DDTHH:mm:ssZ",i=_.z(this),c=this.$H,h=this.$m,D=this.$M,g=r.weekdays,$=r.months,F=r.meridiem,N=function(P,x,V,q){return P&&(P[x]||P(s,t))||V[x].slice(0,q)},J=function(P){return _.s(c%12||12,P,"0")},k=F||function(P,x,V){var q=P<12?"AM":"PM";return V?q.toLowerCase():q};return t.replace(H,function(P,x){return x||function(V){switch(V){case"YY":return String(s.$y).slice(-2);case"YYYY":return _.s(s.$y,4,"0");case"M":return D+1;case"MM":return _.s(D+1,2,"0");case"MMM":return N(r.monthsShort,D,$,3);case"MMMM":return N($,D);case"D":return s.$D;case"DD":return _.s(s.$D,2,"0");case"d":return String(s.$W);case"dd":return N(r.weekdaysMin,s.$W,g,2);case"ddd":return N(r.weekdaysShort,s.$W,g,3);case"dddd":return g[s.$W];case"H":return String(c);case"HH":return _.s(c,2,"0");case"h":return J(1);case"hh":return J(2);case"a":return k(c,h,!0);case"A":return k(c,h,!1);case"m":return String(h);case"mm":return _.s(h,2,"0");case"s":return String(s.$s);case"ss":return _.s(s.$s,2,"0");case"SSS":return _.s(s.$ms,3,"0");case"Z":return i}return null}(P)||i.replace(":","")})},n.utcOffset=function(){return 15*-Math.round(this.$d.getTimezoneOffset()/15)},n.diff=function(e,s,r){var t,i=this,c=_.p(s),h=d(e),D=(h.utcOffset()-this.utcOffset())*a,g=this-h,$=function(){return _.m(i,h)};switch(c){case E:t=$()/12;break;case y:t=$();break;case ee:t=$()/3;break;case Y:t=(g-D)/6048e5;break;case T:t=(g-D)/864e5;break;case B:t=g/m;break;case C:t=g/a;break;case j:t=g/Z;break;default:t=g}return r?t:_.a(t)},n.daysInMonth=function(){return this.endOf(y).$D},n.$locale=function(){return v[this.$L]},n.locale=function(e,s){if(!e)return this.$L;var r=this.clone(),t=R(e,s,!0);return t&&(r.$L=t),r},n.clone=function(){return _.w(this.$d,this)},n.toDate=function(){return new Date(this.valueOf())},n.toJSON=function(){return this.isValid()?this.toISOString():null},n.toISOString=function(){return this.$d.toISOString()},n.toString=function(){return this.$d.toUTCString()},u}(),re=X.prototype;return d.prototype=re,[["$ms",I],["$s",j],["$m",C],["$H",B],["$W",T],["$M",y],["$y",E],["$D",K]].forEach(function(u){re[u[1]]=function(n){return this.$g(n,u[0],u[1])}}),d.extend=function(u,n){return u.$i||(u(n,X,d),u.$i=!0),d},d.locale=R,d.isDayjs=z,d.unix=function(u){return d(1e3*u)},d.en=v[f],d.Ls=v,d.p={},d})}}]);
