(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[818],{76504:function(R){R.exports={actions:"actions___VoCkO","search-wrapper":"search-wrapper___2CWWw","table-wrapper":"table-wrapper___I8xyZ"}},45842:function(R,g,a){"use strict";a.r(g);var v=a(8963),T=a(38291),y=a(88983),D=a(47933),o=a(13062),$=a(71230),I=a(49111),O=a(19650),G=a(89032),E=a(15746),S=a(47673),U=a(77808),j=a(90636),m=a(3182),ae=a(9715),w=a(71481),W=a(11849),Q=a(57663),A=a(71577),B=a(2824),b=a(67294),X=a(12666),p=a(27484),_=a.n(p),q=a(36773),te=a(99871),s=a(1987),n=a(76504),t=a.n(n),e=a(85893),r=function(){var u=[{label:"\u5168\u90E8",value:""},{label:"\u5F85\u53D1\u8D27",value:"3"},{label:"\u5F85\u6536\u8D27",value:"4"},{label:"\u5DF2\u7B7E\u6536",value:"5"},{label:"\u552E\u540E/\u7ED3\u675F",value:"6"}],l=(0,b.useState)(""),d=(0,B.Z)(l,2),h=d[0],P=d[1],L=function(c){var f=c.target.value;P(f)},F=[{title:"\u64CD\u4F5C",dataIndex:"queryOrderID",key:"queryOrderID",render:function(c,f){return(0,e.jsx)(A.Z,{type:"link",children:"\u67E5\u770B"})}},{title:"\u8BA2\u5355\u72B6\u6001",dataIndex:"status",key:"status",render:function(c){var f=u.find(function(x){return x.value===c.toString()});return f?f.label:""}},{title:"\u8BA2\u5355\u7F16\u53F7",dataIndex:"orderSn",key:"orderSn"},{title:"\u5546\u54C1\u540D\u79F0",dataIndex:"name",key:"name"},{title:"\u5546\u6237\u540D\u79F0",dataIndex:"supplierOrganizationName",key:"supplierOrganizationName"},{title:"\u652F\u4ED8\u91D1\u989D",dataIndex:"verifyTime",key:"verifyTime",render:function(c,f){var x=parseFloat(f.totalPrice)+parseFloat(f.postPrice);return"\xA5".concat(x.toFixed(2))}},{title:"\u5458\u5DE5\u7535\u8BDD",dataIndex:"consumerUserPhone",key:"consumerUserPhone"},{title:"\u5458\u5DE5\u7EC4\u7EC7",dataIndex:"consumerOrganizationName",key:"consumerOrganizationName"},{title:"\u652F\u4ED8\u65F6\u95F4",dataIndex:"payedAt",key:"payedAt",render:function(c){return _()(c).format("YYYY-MM-DD HH:mm")}}],z=(0,b.useState)([]),Y=(0,B.Z)(z,2),k=Y[0],C=Y[1],K=(0,b.useState)(!1),H=(0,B.Z)(K,2),ee=H[0],re=H[1],he=(0,b.useState)({current:1,pageSize:10,showSizeChanger:!0,showQuickJumper:!0,showTotal:function(c){return"\u603B\u5171 ".concat(c," \u6761")}}),se=(0,B.Z)(he,2),ue=se[0],ie=se[1],me=function(c){ie({current:c.current||1,pageSize:c.pageSize||10}),ce((0,W.Z)((0,W.Z)({},J),{},{pageNum:c.current||1,pageSize:c.pageSize||10}))},pe=(0,b.useState)([]),oe=(0,B.Z)(pe,2),_e=oe[0],ve=oe[1],De=function(c){ve(c)},Oe={selectedRowKeys:_e,onChange:De},Ee=w.Z.useForm(),Me=(0,B.Z)(Ee,1),ne=Me[0],ge=(0,b.useState)({orderSn:"",goodName:"",pageNum:1,pageSize:10,supplierOrganizationName:"",userPhone:"",userOrganizationName:""}),le=(0,B.Z)(ge,2),J=le[0],ce=le[1],de=function(){var Z=(0,m.Z)((0,j.Z)().mark(function c(){var f;return(0,j.Z)().wrap(function(M){for(;;)switch(M.prev=M.next){case 0:return M.next=2,ne.getFieldsValue();case 2:f=M.sent,ce((0,W.Z)((0,W.Z)({},J),{},{pageNum:1,orderSn:f.orderSn||"",goodName:f.goodName||"",supplierOrganizationName:f.supplierOrganizationName||"",userPhone:f.userPhone||"",userOrganizationName:f.userOrganizationName||""}));case 4:case"end":return M.stop()}},c)}));return function(){return Z.apply(this,arguments)}}(),ye=function(){var Z=(0,m.Z)((0,j.Z)().mark(function c(){return(0,j.Z)().wrap(function(x){for(;;)switch(x.prev=x.next){case 0:return x.next=2,ne.resetFields();case 2:de();case 3:case"end":return x.stop()}},c)}));return function(){return Z.apply(this,arguments)}}(),Pe=function(){var Z=(0,m.Z)((0,j.Z)().mark(function c(){var f;return(0,j.Z)().wrap(function(M){for(;;)switch(M.prev=M.next){case 0:return M.prev=0,re(!0),M.next=4,s.Fw((0,W.Z)((0,W.Z)({},J),{},{status:h,pageNum:J.pageNum,pageSize:J.pageSize}));case 4:f=M.sent,f&&(C(f.list),ie((0,W.Z)((0,W.Z)({},ue),{},{current:J.pageNum,total:f.totalNum}))),re(!1),M.next=12;break;case 9:M.prev=9,M.t0=M.catch(0),re(!1);case 12:case"end":return M.stop()}},c,null,[[0,9]])}));return function(){return Z.apply(this,arguments)}}();(0,b.useEffect)(function(){Pe()},[h,JSON.stringify(J)]);var $e=function(){var Z=(0,m.Z)((0,j.Z)().mark(function c(){var f,x,M,fe,V;return(0,j.Z)().wrap(function(N){for(;;)switch(N.prev=N.next){case 0:return N.prev=0,f=X.Z.create({baseURL:"/",timeout:3e4,withCredentials:!1}),x=localStorage.getItem("token"),N.next=5,f("/api/v1/order/exportOrder?".concat((0,te.R)({ids:_e})),{responseType:"arraybuffer",headers:{Authorization:"".concat(x)}});case 5:M=N.sent,fe=window.URL.createObjectURL(new Blob([M.data],{type:"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"})),V=document.createElement("a"),V.style.display="none",V.href=fe,V.setAttribute("download","excel.xlsx"),document.body.appendChild(V),V.click(),document.body.removeChild(V),N.next=18;break;case 16:N.prev=16,N.t0=N.catch(0);case 18:case"end":return N.stop()}},c,null,[[0,16]])}));return function(){return Z.apply(this,arguments)}}();return(0,e.jsxs)(q.ZP,{children:[(0,e.jsx)("div",{className:t()["search-wrapper"],children:(0,e.jsx)(w.Z,{form:ne,labelCol:{span:8},wrapperCol:{span:16},children:(0,e.jsxs)($.Z,{children:[(0,e.jsx)(E.Z,{span:6,children:(0,e.jsx)(w.Z.Item,{label:"\u8BA2\u5355\u7F16\u53F7",name:"orderSn",rules:[{required:!1}],children:(0,e.jsx)(U.Z,{})})}),(0,e.jsx)(E.Z,{span:6,children:(0,e.jsx)(w.Z.Item,{label:"\u5546\u54C1\u540D\u79F0",name:"goodName",rules:[{required:!1}],children:(0,e.jsx)(U.Z,{})})}),(0,e.jsx)(E.Z,{span:6,children:(0,e.jsx)(w.Z.Item,{label:"\u5546\u6237\u540D\u79F0",name:"supplierOrganizationName",rules:[{required:!1}],children:(0,e.jsx)(U.Z,{})})}),(0,e.jsx)(E.Z,{span:6,children:(0,e.jsx)(w.Z.Item,{label:"\u5458\u5DE5\u7535\u8BDD",name:"userPhone",rules:[{required:!1}],children:(0,e.jsx)(U.Z,{})})}),(0,e.jsx)(E.Z,{span:6,children:(0,e.jsx)(w.Z.Item,{label:"\u5458\u5DE5\u7EC4\u7EC7",name:"userOrganizationName",rules:[{required:!1}],children:(0,e.jsx)(U.Z,{})})}),(0,e.jsx)(E.Z,{offset:12,span:6,style:{textAlign:"right"},children:(0,e.jsxs)(O.Z,{children:[(0,e.jsx)(A.Z,{type:"primary",onClick:de,children:"\u67E5\u8BE2"}),(0,e.jsx)(A.Z,{onClick:ye,children:"\u91CD\u7F6E"})]})})]})})}),(0,e.jsxs)("div",{className:t()["table-wrapper"],children:[(0,e.jsxs)("div",{className:t().actions,children:[(0,e.jsx)(D.ZP.Group,{options:u,optionType:"button",buttonStyle:"solid",onChange:L,value:h}),(0,e.jsx)(A.Z,{type:"primary",onClick:$e,children:"\u5BFC\u51FA"})]}),(0,e.jsx)(T.Z,{columns:F,dataSource:k,loading:ee,onChange:me,pagination:ue,rowKey:"queryOrderID",rowSelection:Oe,scroll:{x:"max-content"}})]})]})};g.default=r},1987:function(R,g,a){"use strict";a.d(g,{dz:function(){return o},Fw:function(){return I}});var v=a(90636),T=a(3182),y=a(99871),D=a(636);function o(S){return $.apply(this,arguments)}function $(){return $=(0,T.Z)((0,v.Z)().mark(function S(U){return(0,v.Z)().wrap(function(m){for(;;)switch(m.prev=m.next){case 0:return m.abrupt("return",(0,D.Z)("/api/v1/order/getOrderDetail?queryOrderID=".concat(U)));case 1:case"end":return m.stop()}},S)})),$.apply(this,arguments)}function I(S){return O.apply(this,arguments)}function O(){return O=(0,T.Z)((0,v.Z)().mark(function S(U){return(0,v.Z)().wrap(function(m){for(;;)switch(m.prev=m.next){case 0:return m.abrupt("return",(0,D.Z)("/api/v1/order/getOrderList?".concat((0,y.R)(U))));case 1:case"end":return m.stop()}},S)})),O.apply(this,arguments)}function G(S){return E.apply(this,arguments)}function E(){return E=_asyncToGenerator(_regeneratorRuntime().mark(function S(U){return _regeneratorRuntime().wrap(function(m){for(;;)switch(m.prev=m.next){case 0:return m.abrupt("return",request("/api/v1/order/exportOrder?".concat(objectToUrlParams(U)),{responseType:"arraybuffer"}));case 1:case"end":return m.stop()}},S)})),E.apply(this,arguments)}},99871:function(R,g,a){"use strict";a.d(g,{R:function(){return v},D:function(){return T}});function v(y){var D=Object.keys(y).map(function(o){return"".concat(o,"=").concat(y[o])});return D.join("&")}function T(y){var D=new RegExp("(^|&)"+y+"=([^&]*)(&|$)"),o=window.location.search.substr(1).match(D);return o!=null?decodeURIComponent(o[2]):null}},636:function(R,g,a){"use strict";var v=a(34792),T=a(48086),y=a(12666),D=y.Z.create({baseURL:"/",timeout:3e4,withCredentials:!1});D.interceptors.request.use(function(o){o&&o.headers&&(o.headers["Content-Type"]||(o.headers["Content-Type"]="application/json"));var $=localStorage.getItem("token");return(o==null?void 0:o.url)!=="/api/v1/user/login"&&(o.headers.Authorization="".concat($)),o},function(o){return Promise.reject(o)}),D.interceptors.response.use(function(o){var $=o.data,I=o.data,O=o.status,G=o.statusText;if(O!==200)return T.ZP.error(G),I;if(I.status===10010)return I;if(I.status!==200)throw T.ZP.error(I.data||o.msg),new Error(o.msg);return I.data},function(o){return console.log("err"+o),Promise.reject(o)}),g.Z=D},15746:function(R,g,a){"use strict";var v=a(21584);g.Z=v.Z},89032:function(R,g,a){"use strict";var v=a(38663),T=a.n(v),y=a(6999)},71230:function(R,g,a){"use strict";var v=a(92820);g.Z=v.Z},13062:function(R,g,a){"use strict";var v=a(38663),T=a.n(v),y=a(6999)},27484:function(R){(function(g,a){R.exports=a()})(this,function(){"use strict";var g=1e3,a=6e4,v=36e5,T="millisecond",y="second",D="minute",o="hour",$="day",I="week",O="month",G="quarter",E="year",S="date",U="Invalid Date",j=/^(\d{4})[-/]?(\d{1,2})?[-/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$/,m=/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g,ae={name:"en",weekdays:"Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),months:"January_February_March_April_May_June_July_August_September_October_November_December".split("_"),ordinal:function(s){var n=["th","st","nd","rd"],t=s%100;return"["+s+(n[(t-20)%10]||n[t]||n[0])+"]"}},w=function(s,n,t){var e=String(s);return!e||e.length>=n?s:""+Array(n+1-e.length).join(t)+s},W={s:w,z:function(s){var n=-s.utcOffset(),t=Math.abs(n),e=Math.floor(t/60),r=t%60;return(n<=0?"+":"-")+w(e,2,"0")+":"+w(r,2,"0")},m:function s(n,t){if(n.date()<t.date())return-s(t,n);var e=12*(t.year()-n.year())+(t.month()-n.month()),r=n.clone().add(e,O),i=t-r<0,u=n.clone().add(e+(i?-1:1),O);return+(-(e+(t-r)/(i?r-u:u-r))||0)},a:function(s){return s<0?Math.ceil(s)||0:Math.floor(s)},p:function(s){return{M:O,y:E,w:I,d:$,D:S,h:o,m:D,s:y,ms:T,Q:G}[s]||String(s||"").toLowerCase().replace(/s$/,"")},u:function(s){return s===void 0}},Q="en",A={};A[Q]=ae;var B="$isDayjsObject",b=function(s){return s instanceof q||!(!s||!s[B])},X=function s(n,t,e){var r;if(!n)return Q;if(typeof n=="string"){var i=n.toLowerCase();A[i]&&(r=i),t&&(A[i]=t,r=i);var u=n.split("-");if(!r&&u.length>1)return s(u[0])}else{var l=n.name;A[l]=n,r=l}return!e&&r&&(Q=r),r||!e&&Q},p=function(s,n){if(b(s))return s.clone();var t=typeof n=="object"?n:{};return t.date=s,t.args=arguments,new q(t)},_=W;_.l=X,_.i=b,_.w=function(s,n){return p(s,{locale:n.$L,utc:n.$u,x:n.$x,$offset:n.$offset})};var q=function(){function s(t){this.$L=X(t.locale,null,!0),this.parse(t),this.$x=this.$x||t.x||{},this[B]=!0}var n=s.prototype;return n.parse=function(t){this.$d=function(e){var r=e.date,i=e.utc;if(r===null)return new Date(NaN);if(_.u(r))return new Date;if(r instanceof Date)return new Date(r);if(typeof r=="string"&&!/Z$/i.test(r)){var u=r.match(j);if(u){var l=u[2]-1||0,d=(u[7]||"0").substring(0,3);return i?new Date(Date.UTC(u[1],l,u[3]||1,u[4]||0,u[5]||0,u[6]||0,d)):new Date(u[1],l,u[3]||1,u[4]||0,u[5]||0,u[6]||0,d)}}return new Date(r)}(t),this.init()},n.init=function(){var t=this.$d;this.$y=t.getFullYear(),this.$M=t.getMonth(),this.$D=t.getDate(),this.$W=t.getDay(),this.$H=t.getHours(),this.$m=t.getMinutes(),this.$s=t.getSeconds(),this.$ms=t.getMilliseconds()},n.$utils=function(){return _},n.isValid=function(){return this.$d.toString()!==U},n.isSame=function(t,e){var r=p(t);return this.startOf(e)<=r&&r<=this.endOf(e)},n.isAfter=function(t,e){return p(t)<this.startOf(e)},n.isBefore=function(t,e){return this.endOf(e)<p(t)},n.$g=function(t,e,r){return _.u(t)?this[e]:this.set(r,t)},n.unix=function(){return Math.floor(this.valueOf()/1e3)},n.valueOf=function(){return this.$d.getTime()},n.startOf=function(t,e){var r=this,i=!!_.u(e)||e,u=_.p(t),l=function(k,C){var K=_.w(r.$u?Date.UTC(r.$y,C,k):new Date(r.$y,C,k),r);return i?K:K.endOf($)},d=function(k,C){return _.w(r.toDate()[k].apply(r.toDate("s"),(i?[0,0,0,0]:[23,59,59,999]).slice(C)),r)},h=this.$W,P=this.$M,L=this.$D,F="set"+(this.$u?"UTC":"");switch(u){case E:return i?l(1,0):l(31,11);case O:return i?l(1,P):l(0,P+1);case I:var z=this.$locale().weekStart||0,Y=(h<z?h+7:h)-z;return l(i?L-Y:L+(6-Y),P);case $:case S:return d(F+"Hours",0);case o:return d(F+"Minutes",1);case D:return d(F+"Seconds",2);case y:return d(F+"Milliseconds",3);default:return this.clone()}},n.endOf=function(t){return this.startOf(t,!1)},n.$set=function(t,e){var r,i=_.p(t),u="set"+(this.$u?"UTC":""),l=(r={},r[$]=u+"Date",r[S]=u+"Date",r[O]=u+"Month",r[E]=u+"FullYear",r[o]=u+"Hours",r[D]=u+"Minutes",r[y]=u+"Seconds",r[T]=u+"Milliseconds",r)[i],d=i===$?this.$D+(e-this.$W):e;if(i===O||i===E){var h=this.clone().set(S,1);h.$d[l](d),h.init(),this.$d=h.set(S,Math.min(this.$D,h.daysInMonth())).$d}else l&&this.$d[l](d);return this.init(),this},n.set=function(t,e){return this.clone().$set(t,e)},n.get=function(t){return this[_.p(t)]()},n.add=function(t,e){var r,i=this;t=Number(t);var u=_.p(e),l=function(P){var L=p(i);return _.w(L.date(L.date()+Math.round(P*t)),i)};if(u===O)return this.set(O,this.$M+t);if(u===E)return this.set(E,this.$y+t);if(u===$)return l(1);if(u===I)return l(7);var d=(r={},r[D]=a,r[o]=v,r[y]=g,r)[u]||1,h=this.$d.getTime()+t*d;return _.w(h,this)},n.subtract=function(t,e){return this.add(-1*t,e)},n.format=function(t){var e=this,r=this.$locale();if(!this.isValid())return r.invalidDate||U;var i=t||"YYYY-MM-DDTHH:mm:ssZ",u=_.z(this),l=this.$H,d=this.$m,h=this.$M,P=r.weekdays,L=r.months,F=r.meridiem,z=function(C,K,H,ee){return C&&(C[K]||C(e,i))||H[K].slice(0,ee)},Y=function(C){return _.s(l%12||12,C,"0")},k=F||function(C,K,H){var ee=C<12?"AM":"PM";return H?ee.toLowerCase():ee};return i.replace(m,function(C,K){return K||function(H){switch(H){case"YY":return String(e.$y).slice(-2);case"YYYY":return _.s(e.$y,4,"0");case"M":return h+1;case"MM":return _.s(h+1,2,"0");case"MMM":return z(r.monthsShort,h,L,3);case"MMMM":return z(L,h);case"D":return e.$D;case"DD":return _.s(e.$D,2,"0");case"d":return String(e.$W);case"dd":return z(r.weekdaysMin,e.$W,P,2);case"ddd":return z(r.weekdaysShort,e.$W,P,3);case"dddd":return P[e.$W];case"H":return String(l);case"HH":return _.s(l,2,"0");case"h":return Y(1);case"hh":return Y(2);case"a":return k(l,d,!0);case"A":return k(l,d,!1);case"m":return String(d);case"mm":return _.s(d,2,"0");case"s":return String(e.$s);case"ss":return _.s(e.$s,2,"0");case"SSS":return _.s(e.$ms,3,"0");case"Z":return u}return null}(C)||u.replace(":","")})},n.utcOffset=function(){return 15*-Math.round(this.$d.getTimezoneOffset()/15)},n.diff=function(t,e,r){var i,u=this,l=_.p(e),d=p(t),h=(d.utcOffset()-this.utcOffset())*a,P=this-d,L=function(){return _.m(u,d)};switch(l){case E:i=L()/12;break;case O:i=L();break;case G:i=L()/3;break;case I:i=(P-h)/6048e5;break;case $:i=(P-h)/864e5;break;case o:i=P/v;break;case D:i=P/a;break;case y:i=P/g;break;default:i=P}return r?i:_.a(i)},n.daysInMonth=function(){return this.endOf(O).$D},n.$locale=function(){return A[this.$L]},n.locale=function(t,e){if(!t)return this.$L;var r=this.clone(),i=X(t,e,!0);return i&&(r.$L=i),r},n.clone=function(){return _.w(this.$d,this)},n.toDate=function(){return new Date(this.valueOf())},n.toJSON=function(){return this.isValid()?this.toISOString():null},n.toISOString=function(){return this.$d.toISOString()},n.toString=function(){return this.$d.toUTCString()},s}(),te=q.prototype;return p.prototype=te,[["$ms",T],["$s",y],["$m",D],["$H",o],["$W",$],["$M",O],["$y",E],["$D",S]].forEach(function(s){te[s[1]]=function(n){return this.$g(n,s[0],s[1])}}),p.extend=function(s,n){return s.$i||(s(n,q,p),s.$i=!0),p},p.locale=X,p.isDayjs=b,p.unix=function(s){return p(1e3*s)},p.en=A[Q],p.Ls=A,p.p={},p})}}]);
