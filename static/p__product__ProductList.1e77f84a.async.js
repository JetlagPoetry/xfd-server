(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[862],{86796:function(re,N,o){"use strict";o.r(N),o.d(N,{default:function(){return a}});var w=o(8963),E=o(38291),k=o(88983),$=o(47933),C=o(90636),M=o(11849),X=o(34792),x=o(48086),z=o(3182),O=o(49111),F=o(19650),L=o(2824),ne=o(71194),K=o(50146),G=o(67294),H=o(36773),V=o(69083),I=o(27484),A=o.n(I),U=o(22122),c=o(83707),v=o(65734),l=function(u,i){return G.createElement(v.Z,(0,U.Z)({},u,{ref:i,icon:c.Z}))},t=G.forwardRef(l),_=o(81910),d=o(85893),s={0:"\u5168\u90E8",1:"\u5728\u552E\u4E2D",2:"\u5DF2\u4E0B\u67B6",3:"\u5DF2\u552E\u7F44"},r=K.Z.confirm,e=function(){var u=(0,G.useState)([]),i=(0,L.Z)(u,2),f=i[0],m=i[1],S=(0,G.useState)({pageNum:1,pageSize:10,queryGoodsListStatus:0}),D=(0,L.Z)(S,2),Z=D[0],b=D[1],R=(0,G.useState)({showSizeChanger:!0,showQuickJumper:!0,showTotal:function(g){return"\u603B\u5171 ".concat(g," \u6761")}}),W=(0,L.Z)(R,2),Y=W[0],P=W[1],B=(0,G.useState)(!1),J=(0,L.Z)(B,2),q=J[0],ae=J[1],ie=[{title:"\u5546\u54C1\u4FE1\u606F",dataIndex:"info",key:"info",width:340,render:function(g,p){var h=p.key,y=p.info;return(0,d.jsxs)("div",{style:{display:"flex",cursor:"pointer",alignItems:"center"},onClick:function(){return ue(h)},children:[(0,d.jsx)("img",{src:y.goodsFrontImage,alt:"picture",style:{width:"60px",height:"60px",marginRight:"8px"}}),(0,d.jsxs)("div",{style:{display:"flex",flexDirection:"column",justifyContent:"space-between",padding:"4px"},children:[(0,d.jsx)("span",{style:{color:"#1890ff",wordBreak:"break-word",wordWrap:"break-word",whiteSpace:"pre-wrap"},children:y.name}),(0,d.jsxs)("div",{children:[(0,d.jsx)("span",{children:"\u5546\u54C1ID\uFF1A"}),(0,d.jsx)("span",{children:y.spuCode})]})]})]})}},{title:"\u72B6\u6001",dataIndex:"status",key:"status",width:100,render:function(g,p){var h=p.status;return(0,d.jsx)("span",{children:s==null?void 0:s[h]})}},{title:"\u91C7\u8D2D\u4EF7\u683C",dataIndex:"buyPrice",key:"buyPrice",render:function(g,p){var h=p.buyPrice;return(0,d.jsx)("span",{children:"\uFFE5".concat(h.minPrice,"~\uFFE5").concat(h.maxPrice)})}},{title:"\u96F6\u552E\u4EF7\u683C",dataIndex:"retailPrice",key:"retailPrice",render:function(g,p){var h=p.retailPrice;return(0,d.jsx)("span",{children:"\uFFE5".concat(h.minPrice,"~\uFFE5").concat(h.maxPrice)})}},{title:"\u96F6\u552E\u6570\u91CF",dataIndex:"retailNum",key:"retailNum",width:100},{title:"\u521B\u5EFA\u65F6\u95F4",dataIndex:"createTime",key:"createTime",render:function(g,p){var h=p.createTime;return(0,d.jsx)("span",{children:A()(h).format("YYYY-MM-DD HH:mm:ss")})}},{title:"\u66F4\u65B0\u65F6\u95F4",dataIndex:"updateTime",key:"updateTime",render:function(g,p){var h=p.updateTime;return(0,d.jsx)("span",{children:A()(h).format("YYYY-MM-DD HH:mm:ss")})}},{title:"\u64CD\u4F5C",dataIndex:"action",key:"action",width:100,render:function(g,p){var h=p.key,y=p.status;return(0,d.jsx)(F.Z,{children:(0,d.jsxs)("div",{style:{display:"flex",flexDirection:"column"},children:[y===2&&(0,d.jsx)("a",{onClick:function(){return se(h)},children:"\u4E0A\u67B6"}),y===1&&(0,d.jsx)("a",{onClick:function(){return se(h)},children:"\u4E0B\u67B6"}),(y===1||y===2||y===3)&&(0,d.jsx)("a",{onClick:function(){return oe(h,y)},children:"\u5220\u9664"})]})})}}],ue=function(g){_.m8.push("/product/detail/".concat(g))},se=function(){var j=(0,z.Z)((0,C.Z)().mark(function g(p){return(0,C.Z)().wrap(function(y){for(;;)switch(y.prev=y.next){case 0:return y.next=2,(0,V.pm)({goodsID:p});case 2:x.ZP.success("\u64CD\u4F5C\u6210\u529F"),b((0,M.Z)({},Z));case 4:case"end":return y.stop()}},g)}));return function(p){return j.apply(this,arguments)}}(),oe=function(g,p){r({title:"\u5220\u9664\u786E\u8BA4",icon:(0,d.jsx)(t,{}),content:"\u786E\u8BA4\u5220\u9664\u8BE5\u5546\u54C1\u5417\uFF1F",onOk:function(){return(0,z.Z)((0,C.Z)().mark(function y(){return(0,C.Z)().wrap(function(ee){for(;;)switch(ee.prev=ee.next){case 0:return ee.next=2,(0,V.ys)({goodsID:g,goodsStatus:p});case 2:x.ZP.success("\u5220\u9664\u6210\u529F"),b((0,M.Z)({},Z));case 4:case"end":return ee.stop()}},y)}))()},onCancel:function(){}})},ce=function(){var j=(0,z.Z)((0,C.Z)().mark(function g(p){var h,y,Q;return(0,C.Z)().wrap(function(te){for(;;)switch(te.prev=te.next){case 0:return ae(!0),te.next=3,(0,V.k1)(p);case 3:h=te.sent,ae(!1),y=(0,M.Z)((0,M.Z)({},Y),{},{current:h.pageNum,pageSize:h.pageSize,total:h.totalNum}),P(y),Q=h.goodsList.map(function(T){return{key:T.id,info:{goodsFrontImage:T.goodsFrontImage,name:T.name,spuCode:T.spuCode},status:T==null?void 0:T.status,buyPrice:{minPrice:T.wholesalePriceMin,maxPrice:T.wholesalePriceMax},retailPrice:{minPrice:T.retailPriceMin,maxPrice:T.retailPriceMax},retailNum:T.soldNum,createTime:T.createdAt,updateTime:T.updatedAt}}),m(Q);case 9:case"end":return te.stop()}},g)}));return function(p){return j.apply(this,arguments)}}();(0,G.useEffect)(function(){ce(Z)},[Z]);var le=function(g){var p=g.current,h=g.pageSize,y=(0,M.Z)((0,M.Z)({},Z),{},{pageNum:p,pageSize:h});b(y)},de=function(g){var p={pageNum:1,pageSize:Z.pageSize,queryGoodsListStatus:g.target.value};b(p)};return(0,d.jsxs)(H.ZP,{children:[(0,d.jsxs)($.ZP.Group,{defaultValue:0,buttonStyle:"solid",style:{marginBottom:"24px"},size:"large",onChange:de,children:[(0,d.jsx)($.ZP.Button,{value:0,children:"\u5168\u90E8"}),(0,d.jsx)($.ZP.Button,{value:1,children:"\u5728\u552E\u4E2D"}),(0,d.jsx)($.ZP.Button,{value:2,children:"\u5DF2\u4E0B\u67B6"}),(0,d.jsx)($.ZP.Button,{value:3,children:"\u5DF2\u552E\u7F44"})]}),(0,d.jsx)(E.Z,{columns:ie,dataSource:f,onChange:le,pagination:Y,loading:q,scroll:{x:"max-content"}})]})},a=e},69083:function(re,N,o){"use strict";o.d(N,{zE:function(){return C},k1:function(){return z},mZ:function(){return F},BH:function(){return ne},JJ:function(){return G},pm:function(){return V},ys:function(){return A}});var w=o(90636),E=o(3182),k=o(99871),$=o(636);function C(c){return M.apply(this,arguments)}function M(){return M=(0,E.Z)((0,w.Z)().mark(function c(v){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,$.Z)("/api/v1/goods/getGoodsDetail?".concat((0,k.R)(v))));case 1:case"end":return t.stop()}},c)})),M.apply(this,arguments)}function X(c){return x.apply(this,arguments)}function x(){return x=_asyncToGenerator(_regeneratorRuntime().mark(function c(v){return _regeneratorRuntime().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",request("/api/v1/goods/getGoodsList?".concat(objectToUrlParams(v))));case 1:case"end":return t.stop()}},c)})),x.apply(this,arguments)}function z(c){return O.apply(this,arguments)}function O(){return O=(0,E.Z)((0,w.Z)().mark(function c(v){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,$.Z)("/api/v1/goods/getMyGoodsList?".concat((0,k.R)(v))));case 1:case"end":return t.stop()}},c)})),O.apply(this,arguments)}function F(c){return L.apply(this,arguments)}function L(){return L=(0,E.Z)((0,w.Z)().mark(function c(v){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,$.Z)("/api/v1/common/area?".concat((0,k.R)(v))));case 1:case"end":return t.stop()}},c)})),L.apply(this,arguments)}function ne(c){return K.apply(this,arguments)}function K(){return K=(0,E.Z)((0,w.Z)().mark(function c(v){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,$.Z)("/api/v1/mall/categories?".concat((0,k.R)(v))));case 1:case"end":return t.stop()}},c)})),K.apply(this,arguments)}function G(c){return H.apply(this,arguments)}function H(){return H=(0,E.Z)((0,w.Z)().mark(function c(v){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,$.Z)("/api/v1/goods/addGoods",{method:"POST",data:v}));case 1:case"end":return t.stop()}},c)})),H.apply(this,arguments)}function V(c){return I.apply(this,arguments)}function I(){return I=(0,E.Z)((0,w.Z)().mark(function c(v){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,$.Z)("/api/v1/goods/modifyMyGoodsStatus",{method:"POST",data:v}));case 1:case"end":return t.stop()}},c)})),I.apply(this,arguments)}function A(c){return U.apply(this,arguments)}function U(){return U=(0,E.Z)((0,w.Z)().mark(function c(v){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,$.Z)("/api/v1/goods/deleteMyGoods",{method:"DELETE",data:v}));case 1:case"end":return t.stop()}},c)})),U.apply(this,arguments)}},27484:function(re){(function(N,o){re.exports=o()})(this,function(){"use strict";var N=1e3,o=6e4,w=36e5,E="millisecond",k="second",$="minute",C="hour",M="day",X="week",x="month",z="quarter",O="year",F="date",L="Invalid Date",ne=/^(\d{4})[-/]?(\d{1,2})?[-/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$/,K=/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g,G={name:"en",weekdays:"Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),months:"January_February_March_April_May_June_July_August_September_October_November_December".split("_"),ordinal:function(s){var r=["th","st","nd","rd"],e=s%100;return"["+s+(r[(e-20)%10]||r[e]||r[0])+"]"}},H=function(s,r,e){var a=String(s);return!a||a.length>=r?s:""+Array(r+1-a.length).join(e)+s},V={s:H,z:function(s){var r=-s.utcOffset(),e=Math.abs(r),a=Math.floor(e/60),n=e%60;return(r<=0?"+":"-")+H(a,2,"0")+":"+H(n,2,"0")},m:function s(r,e){if(r.date()<e.date())return-s(e,r);var a=12*(e.year()-r.year())+(e.month()-r.month()),n=r.clone().add(a,x),u=e-n<0,i=r.clone().add(a+(u?-1:1),x);return+(-(a+(e-n)/(u?n-i:i-n))||0)},a:function(s){return s<0?Math.ceil(s)||0:Math.floor(s)},p:function(s){return{M:x,y:O,w:X,d:M,D:F,h:C,m:$,s:k,ms:E,Q:z}[s]||String(s||"").toLowerCase().replace(/s$/,"")},u:function(s){return s===void 0}},I="en",A={};A[I]=G;var U="$isDayjsObject",c=function(s){return s instanceof _||!(!s||!s[U])},v=function s(r,e,a){var n;if(!r)return I;if(typeof r=="string"){var u=r.toLowerCase();A[u]&&(n=u),e&&(A[u]=e,n=u);var i=r.split("-");if(!n&&i.length>1)return s(i[0])}else{var f=r.name;A[f]=r,n=f}return!a&&n&&(I=n),n||!a&&I},l=function(s,r){if(c(s))return s.clone();var e=typeof r=="object"?r:{};return e.date=s,e.args=arguments,new _(e)},t=V;t.l=v,t.i=c,t.w=function(s,r){return l(s,{locale:r.$L,utc:r.$u,x:r.$x,$offset:r.$offset})};var _=function(){function s(e){this.$L=v(e.locale,null,!0),this.parse(e),this.$x=this.$x||e.x||{},this[U]=!0}var r=s.prototype;return r.parse=function(e){this.$d=function(a){var n=a.date,u=a.utc;if(n===null)return new Date(NaN);if(t.u(n))return new Date;if(n instanceof Date)return new Date(n);if(typeof n=="string"&&!/Z$/i.test(n)){var i=n.match(ne);if(i){var f=i[2]-1||0,m=(i[7]||"0").substring(0,3);return u?new Date(Date.UTC(i[1],f,i[3]||1,i[4]||0,i[5]||0,i[6]||0,m)):new Date(i[1],f,i[3]||1,i[4]||0,i[5]||0,i[6]||0,m)}}return new Date(n)}(e),this.init()},r.init=function(){var e=this.$d;this.$y=e.getFullYear(),this.$M=e.getMonth(),this.$D=e.getDate(),this.$W=e.getDay(),this.$H=e.getHours(),this.$m=e.getMinutes(),this.$s=e.getSeconds(),this.$ms=e.getMilliseconds()},r.$utils=function(){return t},r.isValid=function(){return this.$d.toString()!==L},r.isSame=function(e,a){var n=l(e);return this.startOf(a)<=n&&n<=this.endOf(a)},r.isAfter=function(e,a){return l(e)<this.startOf(a)},r.isBefore=function(e,a){return this.endOf(a)<l(e)},r.$g=function(e,a,n){return t.u(e)?this[a]:this.set(n,e)},r.unix=function(){return Math.floor(this.valueOf()/1e3)},r.valueOf=function(){return this.$d.getTime()},r.startOf=function(e,a){var n=this,u=!!t.u(a)||a,i=t.p(e),f=function(Y,P){var B=t.w(n.$u?Date.UTC(n.$y,P,Y):new Date(n.$y,P,Y),n);return u?B:B.endOf(M)},m=function(Y,P){return t.w(n.toDate()[Y].apply(n.toDate("s"),(u?[0,0,0,0]:[23,59,59,999]).slice(P)),n)},S=this.$W,D=this.$M,Z=this.$D,b="set"+(this.$u?"UTC":"");switch(i){case O:return u?f(1,0):f(31,11);case x:return u?f(1,D):f(0,D+1);case X:var R=this.$locale().weekStart||0,W=(S<R?S+7:S)-R;return f(u?Z-W:Z+(6-W),D);case M:case F:return m(b+"Hours",0);case C:return m(b+"Minutes",1);case $:return m(b+"Seconds",2);case k:return m(b+"Milliseconds",3);default:return this.clone()}},r.endOf=function(e){return this.startOf(e,!1)},r.$set=function(e,a){var n,u=t.p(e),i="set"+(this.$u?"UTC":""),f=(n={},n[M]=i+"Date",n[F]=i+"Date",n[x]=i+"Month",n[O]=i+"FullYear",n[C]=i+"Hours",n[$]=i+"Minutes",n[k]=i+"Seconds",n[E]=i+"Milliseconds",n)[u],m=u===M?this.$D+(a-this.$W):a;if(u===x||u===O){var S=this.clone().set(F,1);S.$d[f](m),S.init(),this.$d=S.set(F,Math.min(this.$D,S.daysInMonth())).$d}else f&&this.$d[f](m);return this.init(),this},r.set=function(e,a){return this.clone().$set(e,a)},r.get=function(e){return this[t.p(e)]()},r.add=function(e,a){var n,u=this;e=Number(e);var i=t.p(a),f=function(D){var Z=l(u);return t.w(Z.date(Z.date()+Math.round(D*e)),u)};if(i===x)return this.set(x,this.$M+e);if(i===O)return this.set(O,this.$y+e);if(i===M)return f(1);if(i===X)return f(7);var m=(n={},n[$]=o,n[C]=w,n[k]=N,n)[i]||1,S=this.$d.getTime()+e*m;return t.w(S,this)},r.subtract=function(e,a){return this.add(-1*e,a)},r.format=function(e){var a=this,n=this.$locale();if(!this.isValid())return n.invalidDate||L;var u=e||"YYYY-MM-DDTHH:mm:ssZ",i=t.z(this),f=this.$H,m=this.$m,S=this.$M,D=n.weekdays,Z=n.months,b=n.meridiem,R=function(P,B,J,q){return P&&(P[B]||P(a,u))||J[B].slice(0,q)},W=function(P){return t.s(f%12||12,P,"0")},Y=b||function(P,B,J){var q=P<12?"AM":"PM";return J?q.toLowerCase():q};return u.replace(K,function(P,B){return B||function(J){switch(J){case"YY":return String(a.$y).slice(-2);case"YYYY":return t.s(a.$y,4,"0");case"M":return S+1;case"MM":return t.s(S+1,2,"0");case"MMM":return R(n.monthsShort,S,Z,3);case"MMMM":return R(Z,S);case"D":return a.$D;case"DD":return t.s(a.$D,2,"0");case"d":return String(a.$W);case"dd":return R(n.weekdaysMin,a.$W,D,2);case"ddd":return R(n.weekdaysShort,a.$W,D,3);case"dddd":return D[a.$W];case"H":return String(f);case"HH":return t.s(f,2,"0");case"h":return W(1);case"hh":return W(2);case"a":return Y(f,m,!0);case"A":return Y(f,m,!1);case"m":return String(m);case"mm":return t.s(m,2,"0");case"s":return String(a.$s);case"ss":return t.s(a.$s,2,"0");case"SSS":return t.s(a.$ms,3,"0");case"Z":return i}return null}(P)||i.replace(":","")})},r.utcOffset=function(){return 15*-Math.round(this.$d.getTimezoneOffset()/15)},r.diff=function(e,a,n){var u,i=this,f=t.p(a),m=l(e),S=(m.utcOffset()-this.utcOffset())*o,D=this-m,Z=function(){return t.m(i,m)};switch(f){case O:u=Z()/12;break;case x:u=Z();break;case z:u=Z()/3;break;case X:u=(D-S)/6048e5;break;case M:u=(D-S)/864e5;break;case C:u=D/w;break;case $:u=D/o;break;case k:u=D/N;break;default:u=D}return n?u:t.a(u)},r.daysInMonth=function(){return this.endOf(x).$D},r.$locale=function(){return A[this.$L]},r.locale=function(e,a){if(!e)return this.$L;var n=this.clone(),u=v(e,a,!0);return u&&(n.$L=u),n},r.clone=function(){return t.w(this.$d,this)},r.toDate=function(){return new Date(this.valueOf())},r.toJSON=function(){return this.isValid()?this.toISOString():null},r.toISOString=function(){return this.$d.toISOString()},r.toString=function(){return this.$d.toUTCString()},s}(),d=_.prototype;return l.prototype=d,[["$ms",E],["$s",k],["$m",$],["$H",C],["$W",M],["$M",x],["$y",O],["$D",F]].forEach(function(s){d[s[1]]=function(r){return this.$g(r,s[0],s[1])}}),l.extend=function(s,r){return s.$i||(s(r,_,l),s.$i=!0),l},l.locale=v,l.isDayjs=c,l.unix=function(s){return l(1e3*s)},l.en=A[I],l.Ls=A,l.p={},l})}}]);