(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[862],{86796:function(ee,k,o){"use strict";o.r(k),o.d(k,{default:function(){return a}});var S=o(8963),j=o(38291),P=o(88983),d=o(47933),c=o(90636),$=o(11849),I=o(34792),Z=o(48086),B=o(3182),A=o(49111),U=o(19650),G=o(2824),re=o(71194),K=o(50146),b=o(67294),H=o(36773),V=o(69083),F=o(27484),L=o.n(F),z=o(22122),l=o(83707),D=o(65734),f=function(u,i){return b.createElement(D.Z,(0,z.Z)({},u,{ref:i,icon:l.Z}))},t=b.forwardRef(f),X=o(81910),h=o(85893),s={0:"\u5168\u90E8",1:"\u5728\u552E\u4E2D",2:"\u5DF2\u4E0B\u67B6",3:"\u5DF2\u552E\u7F44"},r=K.Z.confirm,e=function(){var u=(0,b.useState)([]),i=(0,G.Z)(u,2),m=i[0],g=i[1],w=(0,b.useState)({pageNum:1,pageSize:10,queryGoodsListStatus:0}),x=(0,G.Z)(w,2),C=x[0],R=x[1],W=(0,b.useState)({showSizeChanger:!0,showQuickJumper:!0,showTotal:function(y){return"\u603B\u5171 ".concat(y," \u6761")}}),N=(0,G.Z)(W,2),Y=N[0],T=N[1],_=(0,b.useState)(!1),J=(0,G.Z)(_,2),q=J[0],ae=J[1],ie=[{title:"\u5546\u54C1\u4FE1\u606F",dataIndex:"info",key:"info",width:340,render:function(y,v){var p=v.key,M=v.info;return(0,h.jsxs)("div",{style:{display:"flex",cursor:"pointer",alignItems:"center"},onClick:function(){return ue(p)},children:[(0,h.jsx)("img",{src:M.goodsFrontImage,alt:"picture",style:{width:"60px",height:"60px",marginRight:"8px"}}),(0,h.jsxs)("div",{style:{display:"flex",flexDirection:"column",justifyContent:"space-between",padding:"4px"},children:[(0,h.jsx)("span",{style:{color:"#1890ff",wordBreak:"break-word",wordWrap:"break-word",whiteSpace:"pre-wrap"},children:M.name}),(0,h.jsxs)("div",{children:[(0,h.jsx)("span",{children:"\u5546\u54C1ID\uFF1A"}),(0,h.jsx)("span",{children:M.spuCode})]})]})]})}},{title:"\u72B6\u6001",dataIndex:"status",key:"status",width:100,render:function(y,v){var p=v.status;return(0,h.jsx)("span",{children:s==null?void 0:s[p]})}},{title:"\u91C7\u8D2D\u4EF7\u683C",dataIndex:"buyPrice",key:"buyPrice",render:function(y,v){var p=v.buyPrice;return(0,h.jsx)("span",{children:"\uFFE5".concat(p.minPrice,"~\uFFE5").concat(p.maxPrice)})}},{title:"\u96F6\u552E\u4EF7\u683C",dataIndex:"retailPrice",key:"retailPrice",render:function(y,v){var p=v.retailPrice;return(0,h.jsx)("span",{children:"\uFFE5".concat(p.minPrice,"~\uFFE5").concat(p.maxPrice)})}},{title:"\u96F6\u552E\u6570\u91CF",dataIndex:"retailNum",key:"retailNum",width:100},{title:"\u521B\u5EFA\u65F6\u95F4",dataIndex:"createTime",key:"createTime",render:function(y,v){var p=v.createTime;return(0,h.jsx)("span",{children:L()(p).format("YYYY-MM-DD HH:mm:ss")})}},{title:"\u66F4\u65B0\u65F6\u95F4",dataIndex:"updateTime",key:"updateTime",render:function(y,v){var p=v.updateTime;return(0,h.jsx)("span",{children:L()(p).format("YYYY-MM-DD HH:mm:ss")})}},{title:"\u64CD\u4F5C",dataIndex:"action",key:"action",width:100,render:function(y,v){var p=v.key,M=v.status;return(0,h.jsx)(U.Z,{children:(0,h.jsxs)("div",{style:{display:"flex",flexDirection:"column"},children:[M===2&&(0,h.jsx)("a",{onClick:function(){return se(p)},children:"\u4E0A\u67B6"}),M===1&&(0,h.jsx)("a",{onClick:function(){return se(p)},children:"\u4E0B\u67B6"}),(M===1||M===2||M===3)&&(0,h.jsx)("a",{onClick:function(){return oe(p,M)},children:"\u5220\u9664"})]})})}}],ue=function(y){X.m8.push("/product/detail/".concat(y))},se=function(){var E=(0,B.Z)((0,c.Z)().mark(function y(v){return(0,c.Z)().wrap(function(M){for(;;)switch(M.prev=M.next){case 0:return M.next=2,(0,V.pm)({goodsID:v});case 2:Z.ZP.success("\u64CD\u4F5C\u6210\u529F"),R((0,$.Z)({},C));case 4:case"end":return M.stop()}},y)}));return function(v){return E.apply(this,arguments)}}(),oe=function(y,v){r({title:"\u5220\u9664\u786E\u8BA4",icon:(0,h.jsx)(t,{}),content:"\u786E\u8BA4\u5220\u9664\u8BE5\u5546\u54C1\u5417\uFF1F",onOk:function(){return(0,B.Z)((0,c.Z)().mark(function M(){return(0,c.Z)().wrap(function(te){for(;;)switch(te.prev=te.next){case 0:return te.next=2,(0,V.ys)({goodsID:y,goodsStatus:v});case 2:Z.ZP.success("\u5220\u9664\u6210\u529F"),R((0,$.Z)({},C));case 4:case"end":return te.stop()}},M)}))()},onCancel:function(){}})},ce=function(){var E=(0,B.Z)((0,c.Z)().mark(function y(v){var p,M,Q;return(0,c.Z)().wrap(function(ne){for(;;)switch(ne.prev=ne.next){case 0:return ae(!0),ne.next=3,(0,V.k1)(v);case 3:p=ne.sent,ae(!1),M=(0,$.Z)((0,$.Z)({},Y),{},{current:p.pageNum,pageSize:p.pageSize,total:p.totalNum}),T(M),Q=p.goodsList.map(function(O){return{key:O.id,info:{goodsFrontImage:O.goodsFrontImage,name:O.name,spuCode:O.spuCode},status:O==null?void 0:O.status,buyPrice:{minPrice:O.wholesalePriceMin,maxPrice:O.wholesalePriceMax},retailPrice:{minPrice:O.retailPriceMin,maxPrice:O.retailPriceMax},retailNum:O.soldNum,createTime:O.createdAt,updateTime:O.updatedAt}}),g(Q);case 9:case"end":return ne.stop()}},y)}));return function(v){return E.apply(this,arguments)}}();(0,b.useEffect)(function(){ce(C)},[C]);var le=function(y){var v=y.current,p=y.pageSize,M=(0,$.Z)((0,$.Z)({},C),{},{pageNum:v,pageSize:p});R(M)},de=function(y){var v={pageNum:1,pageSize:C.pageSize,queryGoodsListStatus:y.target.value};R(v)};return(0,h.jsxs)(H.ZP,{children:[(0,h.jsxs)(d.ZP.Group,{defaultValue:0,buttonStyle:"solid",style:{marginBottom:"24px"},size:"large",onChange:de,children:[(0,h.jsx)(d.ZP.Button,{value:0,children:"\u5168\u90E8"}),(0,h.jsx)(d.ZP.Button,{value:1,children:"\u5728\u552E\u4E2D"}),(0,h.jsx)(d.ZP.Button,{value:2,children:"\u5DF2\u4E0B\u67B6"}),(0,h.jsx)(d.ZP.Button,{value:3,children:"\u5DF2\u552E\u7F44"})]}),(0,h.jsx)(j.Z,{columns:ie,dataSource:m,onChange:le,pagination:Y,loading:q,scroll:{x:"max-content"}})]})},a=e},69083:function(ee,k,o){"use strict";o.d(k,{zE:function(){return c},k1:function(){return B},mZ:function(){return U},BH:function(){return re},JJ:function(){return b},pm:function(){return V},ys:function(){return L}});var S=o(90636),j=o(3182),P=o(99871),d=o(636);function c(l){return $.apply(this,arguments)}function $(){return $=(0,j.Z)((0,S.Z)().mark(function l(D){return(0,S.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,d.Z)("/api/v1/goods/getGoodsDetail?".concat((0,P.R)(D))));case 1:case"end":return t.stop()}},l)})),$.apply(this,arguments)}function I(l){return Z.apply(this,arguments)}function Z(){return Z=_asyncToGenerator(_regeneratorRuntime().mark(function l(D){return _regeneratorRuntime().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",request("/api/v1/goods/getGoodsList?".concat(objectToUrlParams(D))));case 1:case"end":return t.stop()}},l)})),Z.apply(this,arguments)}function B(l){return A.apply(this,arguments)}function A(){return A=(0,j.Z)((0,S.Z)().mark(function l(D){return(0,S.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,d.Z)("/api/v1/goods/getMyGoodsList?".concat((0,P.R)(D))));case 1:case"end":return t.stop()}},l)})),A.apply(this,arguments)}function U(l){return G.apply(this,arguments)}function G(){return G=(0,j.Z)((0,S.Z)().mark(function l(D){return(0,S.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,d.Z)("/api/v1/common/area?".concat((0,P.R)(D))));case 1:case"end":return t.stop()}},l)})),G.apply(this,arguments)}function re(l){return K.apply(this,arguments)}function K(){return K=(0,j.Z)((0,S.Z)().mark(function l(D){return(0,S.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,d.Z)("/api/v1/mall/categories?".concat((0,P.R)(D))));case 1:case"end":return t.stop()}},l)})),K.apply(this,arguments)}function b(l){return H.apply(this,arguments)}function H(){return H=(0,j.Z)((0,S.Z)().mark(function l(D){return(0,S.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,d.Z)("/api/v1/goods/addGoods",{method:"POST",data:D}));case 1:case"end":return t.stop()}},l)})),H.apply(this,arguments)}function V(l){return F.apply(this,arguments)}function F(){return F=(0,j.Z)((0,S.Z)().mark(function l(D){return(0,S.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,d.Z)("/api/v1/goods/modifyMyGoodsStatus",{method:"POST",data:D}));case 1:case"end":return t.stop()}},l)})),F.apply(this,arguments)}function L(l){return z.apply(this,arguments)}function z(){return z=(0,j.Z)((0,S.Z)().mark(function l(D){return(0,S.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,d.Z)("/api/v1/goods/deleteMyGoods",{method:"DELETE",data:D}));case 1:case"end":return t.stop()}},l)})),z.apply(this,arguments)}},99871:function(ee,k,o){"use strict";o.d(k,{R:function(){return S},D:function(){return j}});function S(P){var d=Object.keys(P).map(function(c){return"".concat(c,"=").concat(P[c])});return d.join("&")}function j(P){var d=new RegExp("(^|&)"+P+"=([^&]*)(&|$)"),c=window.location.search.substr(1).match(d);return c!=null?decodeURIComponent(c[2]):null}},636:function(ee,k,o){"use strict";var S=o(34792),j=o(48086),P=o(12666),d=P.Z.create({baseURL:"/",timeout:3e4,withCredentials:!1});d.interceptors.request.use(function(c){c&&c.headers&&(c.headers["Content-Type"]||(c.headers["Content-Type"]="application/json"));var $=localStorage.getItem("token");return(c==null?void 0:c.url)!=="/api/v1/user/login"&&(c.headers.Authorization="".concat($)),c},function(c){return Promise.reject(c)}),d.interceptors.response.use(function(c){var $=c.data;console.log(Object.prototype.toString.call(c));var I=c.data,Z=c.status,B=c.statusText;if(Z!==200)return j.ZP.error(B),null;if(I.status===10010)return I;if(I.status!==200)throw j.ZP.error(c.msg),new Error(c.msg);return I.data},function(c){return console.log("err"+c),Promise.reject(c)}),k.Z=d},27484:function(ee){(function(k,o){ee.exports=o()})(this,function(){"use strict";var k=1e3,o=6e4,S=36e5,j="millisecond",P="second",d="minute",c="hour",$="day",I="week",Z="month",B="quarter",A="year",U="date",G="Invalid Date",re=/^(\d{4})[-/]?(\d{1,2})?[-/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$/,K=/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g,b={name:"en",weekdays:"Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),months:"January_February_March_April_May_June_July_August_September_October_November_December".split("_"),ordinal:function(s){var r=["th","st","nd","rd"],e=s%100;return"["+s+(r[(e-20)%10]||r[e]||r[0])+"]"}},H=function(s,r,e){var a=String(s);return!a||a.length>=r?s:""+Array(r+1-a.length).join(e)+s},V={s:H,z:function(s){var r=-s.utcOffset(),e=Math.abs(r),a=Math.floor(e/60),n=e%60;return(r<=0?"+":"-")+H(a,2,"0")+":"+H(n,2,"0")},m:function s(r,e){if(r.date()<e.date())return-s(e,r);var a=12*(e.year()-r.year())+(e.month()-r.month()),n=r.clone().add(a,Z),u=e-n<0,i=r.clone().add(a+(u?-1:1),Z);return+(-(a+(e-n)/(u?n-i:i-n))||0)},a:function(s){return s<0?Math.ceil(s)||0:Math.floor(s)},p:function(s){return{M:Z,y:A,w:I,d:$,D:U,h:c,m:d,s:P,ms:j,Q:B}[s]||String(s||"").toLowerCase().replace(/s$/,"")},u:function(s){return s===void 0}},F="en",L={};L[F]=b;var z="$isDayjsObject",l=function(s){return s instanceof X||!(!s||!s[z])},D=function s(r,e,a){var n;if(!r)return F;if(typeof r=="string"){var u=r.toLowerCase();L[u]&&(n=u),e&&(L[u]=e,n=u);var i=r.split("-");if(!n&&i.length>1)return s(i[0])}else{var m=r.name;L[m]=r,n=m}return!a&&n&&(F=n),n||!a&&F},f=function(s,r){if(l(s))return s.clone();var e=typeof r=="object"?r:{};return e.date=s,e.args=arguments,new X(e)},t=V;t.l=D,t.i=l,t.w=function(s,r){return f(s,{locale:r.$L,utc:r.$u,x:r.$x,$offset:r.$offset})};var X=function(){function s(e){this.$L=D(e.locale,null,!0),this.parse(e),this.$x=this.$x||e.x||{},this[z]=!0}var r=s.prototype;return r.parse=function(e){this.$d=function(a){var n=a.date,u=a.utc;if(n===null)return new Date(NaN);if(t.u(n))return new Date;if(n instanceof Date)return new Date(n);if(typeof n=="string"&&!/Z$/i.test(n)){var i=n.match(re);if(i){var m=i[2]-1||0,g=(i[7]||"0").substring(0,3);return u?new Date(Date.UTC(i[1],m,i[3]||1,i[4]||0,i[5]||0,i[6]||0,g)):new Date(i[1],m,i[3]||1,i[4]||0,i[5]||0,i[6]||0,g)}}return new Date(n)}(e),this.init()},r.init=function(){var e=this.$d;this.$y=e.getFullYear(),this.$M=e.getMonth(),this.$D=e.getDate(),this.$W=e.getDay(),this.$H=e.getHours(),this.$m=e.getMinutes(),this.$s=e.getSeconds(),this.$ms=e.getMilliseconds()},r.$utils=function(){return t},r.isValid=function(){return this.$d.toString()!==G},r.isSame=function(e,a){var n=f(e);return this.startOf(a)<=n&&n<=this.endOf(a)},r.isAfter=function(e,a){return f(e)<this.startOf(a)},r.isBefore=function(e,a){return this.endOf(a)<f(e)},r.$g=function(e,a,n){return t.u(e)?this[a]:this.set(n,e)},r.unix=function(){return Math.floor(this.valueOf()/1e3)},r.valueOf=function(){return this.$d.getTime()},r.startOf=function(e,a){var n=this,u=!!t.u(a)||a,i=t.p(e),m=function(Y,T){var _=t.w(n.$u?Date.UTC(n.$y,T,Y):new Date(n.$y,T,Y),n);return u?_:_.endOf($)},g=function(Y,T){return t.w(n.toDate()[Y].apply(n.toDate("s"),(u?[0,0,0,0]:[23,59,59,999]).slice(T)),n)},w=this.$W,x=this.$M,C=this.$D,R="set"+(this.$u?"UTC":"");switch(i){case A:return u?m(1,0):m(31,11);case Z:return u?m(1,x):m(0,x+1);case I:var W=this.$locale().weekStart||0,N=(w<W?w+7:w)-W;return m(u?C-N:C+(6-N),x);case $:case U:return g(R+"Hours",0);case c:return g(R+"Minutes",1);case d:return g(R+"Seconds",2);case P:return g(R+"Milliseconds",3);default:return this.clone()}},r.endOf=function(e){return this.startOf(e,!1)},r.$set=function(e,a){var n,u=t.p(e),i="set"+(this.$u?"UTC":""),m=(n={},n[$]=i+"Date",n[U]=i+"Date",n[Z]=i+"Month",n[A]=i+"FullYear",n[c]=i+"Hours",n[d]=i+"Minutes",n[P]=i+"Seconds",n[j]=i+"Milliseconds",n)[u],g=u===$?this.$D+(a-this.$W):a;if(u===Z||u===A){var w=this.clone().set(U,1);w.$d[m](g),w.init(),this.$d=w.set(U,Math.min(this.$D,w.daysInMonth())).$d}else m&&this.$d[m](g);return this.init(),this},r.set=function(e,a){return this.clone().$set(e,a)},r.get=function(e){return this[t.p(e)]()},r.add=function(e,a){var n,u=this;e=Number(e);var i=t.p(a),m=function(x){var C=f(u);return t.w(C.date(C.date()+Math.round(x*e)),u)};if(i===Z)return this.set(Z,this.$M+e);if(i===A)return this.set(A,this.$y+e);if(i===$)return m(1);if(i===I)return m(7);var g=(n={},n[d]=o,n[c]=S,n[P]=k,n)[i]||1,w=this.$d.getTime()+e*g;return t.w(w,this)},r.subtract=function(e,a){return this.add(-1*e,a)},r.format=function(e){var a=this,n=this.$locale();if(!this.isValid())return n.invalidDate||G;var u=e||"YYYY-MM-DDTHH:mm:ssZ",i=t.z(this),m=this.$H,g=this.$m,w=this.$M,x=n.weekdays,C=n.months,R=n.meridiem,W=function(T,_,J,q){return T&&(T[_]||T(a,u))||J[_].slice(0,q)},N=function(T){return t.s(m%12||12,T,"0")},Y=R||function(T,_,J){var q=T<12?"AM":"PM";return J?q.toLowerCase():q};return u.replace(K,function(T,_){return _||function(J){switch(J){case"YY":return String(a.$y).slice(-2);case"YYYY":return t.s(a.$y,4,"0");case"M":return w+1;case"MM":return t.s(w+1,2,"0");case"MMM":return W(n.monthsShort,w,C,3);case"MMMM":return W(C,w);case"D":return a.$D;case"DD":return t.s(a.$D,2,"0");case"d":return String(a.$W);case"dd":return W(n.weekdaysMin,a.$W,x,2);case"ddd":return W(n.weekdaysShort,a.$W,x,3);case"dddd":return x[a.$W];case"H":return String(m);case"HH":return t.s(m,2,"0");case"h":return N(1);case"hh":return N(2);case"a":return Y(m,g,!0);case"A":return Y(m,g,!1);case"m":return String(g);case"mm":return t.s(g,2,"0");case"s":return String(a.$s);case"ss":return t.s(a.$s,2,"0");case"SSS":return t.s(a.$ms,3,"0");case"Z":return i}return null}(T)||i.replace(":","")})},r.utcOffset=function(){return 15*-Math.round(this.$d.getTimezoneOffset()/15)},r.diff=function(e,a,n){var u,i=this,m=t.p(a),g=f(e),w=(g.utcOffset()-this.utcOffset())*o,x=this-g,C=function(){return t.m(i,g)};switch(m){case A:u=C()/12;break;case Z:u=C();break;case B:u=C()/3;break;case I:u=(x-w)/6048e5;break;case $:u=(x-w)/864e5;break;case c:u=x/S;break;case d:u=x/o;break;case P:u=x/k;break;default:u=x}return n?u:t.a(u)},r.daysInMonth=function(){return this.endOf(Z).$D},r.$locale=function(){return L[this.$L]},r.locale=function(e,a){if(!e)return this.$L;var n=this.clone(),u=D(e,a,!0);return u&&(n.$L=u),n},r.clone=function(){return t.w(this.$d,this)},r.toDate=function(){return new Date(this.valueOf())},r.toJSON=function(){return this.isValid()?this.toISOString():null},r.toISOString=function(){return this.$d.toISOString()},r.toString=function(){return this.$d.toUTCString()},s}(),h=X.prototype;return f.prototype=h,[["$ms",j],["$s",P],["$m",d],["$H",c],["$W",$],["$M",Z],["$y",A],["$D",U]].forEach(function(s){h[s[1]]=function(r){return this.$g(r,s[0],s[1])}}),f.extend=function(s,r){return s.$i||(s(r,X,f),s.$i=!0),f},f.locale=D,f.isDayjs=l,f.unix=function(s){return f(1e3*s)},f.en=L[F],f.Ls=L,f.p={},f})}}]);
