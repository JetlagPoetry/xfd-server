(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[862],{86796:function(re,W,o){"use strict";o.r(W),o.d(W,{default:function(){return a}});var w=o(8963),O=o(35008),A=o(88983),y=o(47933),j=o(90636),S=o(11849),_=o(34792),P=o(48086),z=o(3182),T=o(49111),Y=o(19650),L=o(2824),ne=o(71194),Q=o(48889),G=o(67294),N=o(95916),X=o(69083),I=o(27484),b=o.n(I),J=o(22122),c=o(83707),g=o(65734),d=function(u,i){return G.createElement(g.Z,(0,J.Z)({},u,{ref:i,icon:c.Z}))},t=G.forwardRef(d),q=o(81910),f=o(85893),s={0:"\u5168\u90E8",1:"\u5728\u552E\u4E2D",2:"\u5DF2\u4E0B\u67B6",3:"\u5DF2\u552E\u7F44"},r=Q.Z.confirm,e=function(){var u=(0,G.useState)([]),i=(0,L.Z)(u,2),h=i[0],m=i[1],$=(0,G.useState)({pageNum:1,pageSize:10,queryGoodsListStatus:0}),M=(0,L.Z)($,2),D=M[0],F=M[1],U=(0,G.useState)({showSizeChanger:!0,showQuickJumper:!0,showTotal:function(v){return"\u603B\u5171 ".concat(v," \u6761")}}),K=(0,L.Z)(U,2),R=K[0],C=K[1],H=(0,G.useState)(!1),V=(0,L.Z)(H,2),ee=V[0],ae=V[1],ue=[{title:"\u5546\u54C1\u4FE1\u606F",dataIndex:"info",key:"info",width:340,render:function(v,p){var l=p.key,x=p.info;return(0,f.jsxs)("div",{style:{display:"flex",alignItems:"center"},children:[(0,f.jsx)("img",{src:x.goodsFrontImage,alt:"picture",style:{width:"60px",height:"60px",marginRight:"8px"}}),(0,f.jsxs)("div",{style:{display:"flex",flexDirection:"column",justifyContent:"space-between",padding:"4px"},children:[(0,f.jsx)("span",{onClick:function(){return se(l)},style:{color:"#1890ff",wordBreak:"break-word",wordWrap:"break-word",whiteSpace:"pre-wrap",cursor:"pointer"},children:x.name}),(0,f.jsxs)("div",{children:[(0,f.jsx)("span",{children:"\u5546\u54C1ID\uFF1A"}),(0,f.jsx)("span",{children:x.spuCode})]})]})]})}},{title:"\u72B6\u6001",dataIndex:"status",key:"status",width:100,render:function(v,p){var l=p.status;return(0,f.jsx)("span",{children:s==null?void 0:s[l]})}},{title:"\u91C7\u8D2D\u4EF7\u683C",dataIndex:"buyPrice",key:"buyPrice",render:function(v,p){var l=p.buyPrice;return(0,f.jsx)("span",{children:"\uFFE5".concat(l.minPrice,"~\uFFE5").concat(l.maxPrice)})}},{title:"\u96F6\u552E\u4EF7\u683C",dataIndex:"retailPrice",key:"retailPrice",render:function(v,p){var l=p.retailPrice;return(0,f.jsx)("span",{children:"\uFFE5".concat(l.minPrice,"~\uFFE5").concat(l.maxPrice)})}},{title:"\u96F6\u552E\u9500\u91CF",dataIndex:"retailNum",key:"retailNum",width:100},{title:"\u96F6\u552E\u5E93\u5B58",dataIndex:"stock",key:"stock",width:100},{title:"\u521B\u5EFA\u65F6\u95F4",dataIndex:"createTime",key:"createTime",render:function(v,p){var l=p.createTime;return(0,f.jsx)("span",{children:b()(l).format("YYYY-MM-DD HH:mm:ss")})}},{title:"\u64CD\u4F5C",dataIndex:"action",key:"action",width:100,render:function(v,p){var l=p.key,x=p.status;return(0,f.jsx)(Y.Z,{children:(0,f.jsxs)("div",{style:{display:"flex",flexDirection:"column"},children:[(0,f.jsx)("a",{onClick:function(){return se(l)},children:"\u67E5\u770B"}),x===2&&(0,f.jsx)("a",{onClick:function(){return ie(l,1)},children:"\u4E0A\u67B6"}),x===1&&(0,f.jsx)("a",{onClick:function(){return ie(l,2)},children:"\u4E0B\u67B6"}),(x===1||x===2||x===3)&&(0,f.jsx)("a",{onClick:function(){return oe(l,x)},children:"\u5220\u9664"})]})})}}],se=function(v){q.m8.push("/product/list/detail/".concat(v))},ie=function(){var E=(0,z.Z)((0,j.Z)().mark(function v(p,l){return(0,j.Z)().wrap(function(Z){for(;;)switch(Z.prev=Z.next){case 0:return Z.prev=0,Z.next=3,(0,X.pm)({goodsID:p,goodsStatus:l});case 3:P.ZP.success("\u64CD\u4F5C\u6210\u529F"),F((0,S.Z)({},D)),Z.next=10;break;case 7:Z.prev=7,Z.t0=Z.catch(0),console.log(Z.t0);case 10:case"end":return Z.stop()}},v,null,[[0,7]])}));return function(p,l){return E.apply(this,arguments)}}(),oe=function(v,p){r({title:"\u5220\u9664\u786E\u8BA4",icon:(0,f.jsx)(t,{}),content:"\u786E\u8BA4\u5220\u9664\u8BE5\u5546\u54C1\u5417\uFF1F",onOk:function(){return(0,z.Z)((0,j.Z)().mark(function x(){return(0,j.Z)().wrap(function(B){for(;;)switch(B.prev=B.next){case 0:return B.prev=0,B.next=3,(0,X.ys)({goodsID:v,goodsStatus:p});case 3:P.ZP.success("\u5220\u9664\u6210\u529F"),F((0,S.Z)({},D)),B.next=10;break;case 7:B.prev=7,B.t0=B.catch(0),console.log(B.t0);case 10:case"end":return B.stop()}},x,null,[[0,7]])}))()},onCancel:function(){}})},ce=function(){var E=(0,z.Z)((0,j.Z)().mark(function v(p){var l,x,Z;return(0,j.Z)().wrap(function(te){for(;;)switch(te.prev=te.next){case 0:return ae(!0),te.next=3,(0,X.k1)(p);case 3:l=te.sent,ae(!1),x=(0,S.Z)((0,S.Z)({},R),{},{current:l.pageNum,pageSize:l.pageSize,total:l.totalNum}),C(x),Z=l.goodsList.map(function(k){return{key:k.id,info:{goodsFrontImage:k.goodsFrontImage,name:k.name,spuCode:k.spuCode},status:k==null?void 0:k.status,buyPrice:{minPrice:k.wholesalePriceMin,maxPrice:k.wholesalePriceMax},retailPrice:{minPrice:k.retailPriceMin,maxPrice:k.retailPriceMax},retailNum:k.soldNum,createTime:k.createdAt,updateTime:k.updatedAt,stock:k.stock}}),m(Z);case 9:case"end":return te.stop()}},v)}));return function(p){return E.apply(this,arguments)}}();(0,G.useEffect)(function(){ce(D)},[D]);var le=function(v){var p=v.current,l=v.pageSize,x=(0,S.Z)((0,S.Z)({},D),{},{pageNum:p,pageSize:l});F(x)},de=function(v){var p={pageNum:1,pageSize:D.pageSize,queryGoodsListStatus:v.target.value};F(p)};return(0,f.jsxs)(N.ZP,{children:[(0,f.jsxs)(y.ZP.Group,{defaultValue:0,buttonStyle:"solid",style:{marginBottom:"24px"},size:"large",onChange:de,children:[(0,f.jsx)(y.ZP.Button,{value:0,children:"\u5168\u90E8"}),(0,f.jsx)(y.ZP.Button,{value:1,children:"\u5728\u552E\u4E2D"}),(0,f.jsx)(y.ZP.Button,{value:2,children:"\u5DF2\u4E0B\u67B6"}),(0,f.jsx)(y.ZP.Button,{value:3,children:"\u5DF2\u552E\u7F44"})]}),(0,f.jsx)(O.Z,{columns:ue,dataSource:h,onChange:le,pagination:R,loading:ee,scroll:{x:"max-content"}})]})},a=e},69083:function(re,W,o){"use strict";o.d(W,{g2:function(){return j},k1:function(){return z},mZ:function(){return Y},BH:function(){return ne},JJ:function(){return G},pm:function(){return X},ys:function(){return b}});var w=o(90636),O=o(3182),A=o(99871),y=o(636);function j(c){return S.apply(this,arguments)}function S(){return S=(0,O.Z)((0,w.Z)().mark(function c(g){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/getMyGoodsDetail?".concat((0,A.R)(g))));case 1:case"end":return t.stop()}},c)})),S.apply(this,arguments)}function _(c){return P.apply(this,arguments)}function P(){return P=_asyncToGenerator(_regeneratorRuntime().mark(function c(g){return _regeneratorRuntime().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",request("/api/v1/goods/getGoodsList?".concat(objectToUrlParams(g))));case 1:case"end":return t.stop()}},c)})),P.apply(this,arguments)}function z(c){return T.apply(this,arguments)}function T(){return T=(0,O.Z)((0,w.Z)().mark(function c(g){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/getMyGoodsList?".concat((0,A.R)(g))));case 1:case"end":return t.stop()}},c)})),T.apply(this,arguments)}function Y(c){return L.apply(this,arguments)}function L(){return L=(0,O.Z)((0,w.Z)().mark(function c(g){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/common/area?".concat((0,A.R)(g))));case 1:case"end":return t.stop()}},c)})),L.apply(this,arguments)}function ne(c){return Q.apply(this,arguments)}function Q(){return Q=(0,O.Z)((0,w.Z)().mark(function c(g){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/mall/categories?".concat((0,A.R)(g))));case 1:case"end":return t.stop()}},c)})),Q.apply(this,arguments)}function G(c){return N.apply(this,arguments)}function N(){return N=(0,O.Z)((0,w.Z)().mark(function c(g){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/addGoods",{method:"POST",data:g}));case 1:case"end":return t.stop()}},c)})),N.apply(this,arguments)}function X(c){return I.apply(this,arguments)}function I(){return I=(0,O.Z)((0,w.Z)().mark(function c(g){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/modifyMyGoodsStatus",{method:"POST",data:g}));case 1:case"end":return t.stop()}},c)})),I.apply(this,arguments)}function b(c){return J.apply(this,arguments)}function J(){return J=(0,O.Z)((0,w.Z)().mark(function c(g){return(0,w.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/deleteMyGoods",{method:"DELETE",data:g}));case 1:case"end":return t.stop()}},c)})),J.apply(this,arguments)}},27484:function(re){(function(W,o){re.exports=o()})(this,function(){"use strict";var W=1e3,o=6e4,w=36e5,O="millisecond",A="second",y="minute",j="hour",S="day",_="week",P="month",z="quarter",T="year",Y="date",L="Invalid Date",ne=/^(\d{4})[-/]?(\d{1,2})?[-/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$/,Q=/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g,G={name:"en",weekdays:"Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),months:"January_February_March_April_May_June_July_August_September_October_November_December".split("_"),ordinal:function(s){var r=["th","st","nd","rd"],e=s%100;return"["+s+(r[(e-20)%10]||r[e]||r[0])+"]"}},N=function(s,r,e){var a=String(s);return!a||a.length>=r?s:""+Array(r+1-a.length).join(e)+s},X={s:N,z:function(s){var r=-s.utcOffset(),e=Math.abs(r),a=Math.floor(e/60),n=e%60;return(r<=0?"+":"-")+N(a,2,"0")+":"+N(n,2,"0")},m:function s(r,e){if(r.date()<e.date())return-s(e,r);var a=12*(e.year()-r.year())+(e.month()-r.month()),n=r.clone().add(a,P),u=e-n<0,i=r.clone().add(a+(u?-1:1),P);return+(-(a+(e-n)/(u?n-i:i-n))||0)},a:function(s){return s<0?Math.ceil(s)||0:Math.floor(s)},p:function(s){return{M:P,y:T,w:_,d:S,D:Y,h:j,m:y,s:A,ms:O,Q:z}[s]||String(s||"").toLowerCase().replace(/s$/,"")},u:function(s){return s===void 0}},I="en",b={};b[I]=G;var J="$isDayjsObject",c=function(s){return s instanceof q||!(!s||!s[J])},g=function s(r,e,a){var n;if(!r)return I;if(typeof r=="string"){var u=r.toLowerCase();b[u]&&(n=u),e&&(b[u]=e,n=u);var i=r.split("-");if(!n&&i.length>1)return s(i[0])}else{var h=r.name;b[h]=r,n=h}return!a&&n&&(I=n),n||!a&&I},d=function(s,r){if(c(s))return s.clone();var e=typeof r=="object"?r:{};return e.date=s,e.args=arguments,new q(e)},t=X;t.l=g,t.i=c,t.w=function(s,r){return d(s,{locale:r.$L,utc:r.$u,x:r.$x,$offset:r.$offset})};var q=function(){function s(e){this.$L=g(e.locale,null,!0),this.parse(e),this.$x=this.$x||e.x||{},this[J]=!0}var r=s.prototype;return r.parse=function(e){this.$d=function(a){var n=a.date,u=a.utc;if(n===null)return new Date(NaN);if(t.u(n))return new Date;if(n instanceof Date)return new Date(n);if(typeof n=="string"&&!/Z$/i.test(n)){var i=n.match(ne);if(i){var h=i[2]-1||0,m=(i[7]||"0").substring(0,3);return u?new Date(Date.UTC(i[1],h,i[3]||1,i[4]||0,i[5]||0,i[6]||0,m)):new Date(i[1],h,i[3]||1,i[4]||0,i[5]||0,i[6]||0,m)}}return new Date(n)}(e),this.init()},r.init=function(){var e=this.$d;this.$y=e.getFullYear(),this.$M=e.getMonth(),this.$D=e.getDate(),this.$W=e.getDay(),this.$H=e.getHours(),this.$m=e.getMinutes(),this.$s=e.getSeconds(),this.$ms=e.getMilliseconds()},r.$utils=function(){return t},r.isValid=function(){return this.$d.toString()!==L},r.isSame=function(e,a){var n=d(e);return this.startOf(a)<=n&&n<=this.endOf(a)},r.isAfter=function(e,a){return d(e)<this.startOf(a)},r.isBefore=function(e,a){return this.endOf(a)<d(e)},r.$g=function(e,a,n){return t.u(e)?this[a]:this.set(n,e)},r.unix=function(){return Math.floor(this.valueOf()/1e3)},r.valueOf=function(){return this.$d.getTime()},r.startOf=function(e,a){var n=this,u=!!t.u(a)||a,i=t.p(e),h=function(R,C){var H=t.w(n.$u?Date.UTC(n.$y,C,R):new Date(n.$y,C,R),n);return u?H:H.endOf(S)},m=function(R,C){return t.w(n.toDate()[R].apply(n.toDate("s"),(u?[0,0,0,0]:[23,59,59,999]).slice(C)),n)},$=this.$W,M=this.$M,D=this.$D,F="set"+(this.$u?"UTC":"");switch(i){case T:return u?h(1,0):h(31,11);case P:return u?h(1,M):h(0,M+1);case _:var U=this.$locale().weekStart||0,K=($<U?$+7:$)-U;return h(u?D-K:D+(6-K),M);case S:case Y:return m(F+"Hours",0);case j:return m(F+"Minutes",1);case y:return m(F+"Seconds",2);case A:return m(F+"Milliseconds",3);default:return this.clone()}},r.endOf=function(e){return this.startOf(e,!1)},r.$set=function(e,a){var n,u=t.p(e),i="set"+(this.$u?"UTC":""),h=(n={},n[S]=i+"Date",n[Y]=i+"Date",n[P]=i+"Month",n[T]=i+"FullYear",n[j]=i+"Hours",n[y]=i+"Minutes",n[A]=i+"Seconds",n[O]=i+"Milliseconds",n)[u],m=u===S?this.$D+(a-this.$W):a;if(u===P||u===T){var $=this.clone().set(Y,1);$.$d[h](m),$.init(),this.$d=$.set(Y,Math.min(this.$D,$.daysInMonth())).$d}else h&&this.$d[h](m);return this.init(),this},r.set=function(e,a){return this.clone().$set(e,a)},r.get=function(e){return this[t.p(e)]()},r.add=function(e,a){var n,u=this;e=Number(e);var i=t.p(a),h=function(M){var D=d(u);return t.w(D.date(D.date()+Math.round(M*e)),u)};if(i===P)return this.set(P,this.$M+e);if(i===T)return this.set(T,this.$y+e);if(i===S)return h(1);if(i===_)return h(7);var m=(n={},n[y]=o,n[j]=w,n[A]=W,n)[i]||1,$=this.$d.getTime()+e*m;return t.w($,this)},r.subtract=function(e,a){return this.add(-1*e,a)},r.format=function(e){var a=this,n=this.$locale();if(!this.isValid())return n.invalidDate||L;var u=e||"YYYY-MM-DDTHH:mm:ssZ",i=t.z(this),h=this.$H,m=this.$m,$=this.$M,M=n.weekdays,D=n.months,F=n.meridiem,U=function(C,H,V,ee){return C&&(C[H]||C(a,u))||V[H].slice(0,ee)},K=function(C){return t.s(h%12||12,C,"0")},R=F||function(C,H,V){var ee=C<12?"AM":"PM";return V?ee.toLowerCase():ee};return u.replace(Q,function(C,H){return H||function(V){switch(V){case"YY":return String(a.$y).slice(-2);case"YYYY":return t.s(a.$y,4,"0");case"M":return $+1;case"MM":return t.s($+1,2,"0");case"MMM":return U(n.monthsShort,$,D,3);case"MMMM":return U(D,$);case"D":return a.$D;case"DD":return t.s(a.$D,2,"0");case"d":return String(a.$W);case"dd":return U(n.weekdaysMin,a.$W,M,2);case"ddd":return U(n.weekdaysShort,a.$W,M,3);case"dddd":return M[a.$W];case"H":return String(h);case"HH":return t.s(h,2,"0");case"h":return K(1);case"hh":return K(2);case"a":return R(h,m,!0);case"A":return R(h,m,!1);case"m":return String(m);case"mm":return t.s(m,2,"0");case"s":return String(a.$s);case"ss":return t.s(a.$s,2,"0");case"SSS":return t.s(a.$ms,3,"0");case"Z":return i}return null}(C)||i.replace(":","")})},r.utcOffset=function(){return 15*-Math.round(this.$d.getTimezoneOffset()/15)},r.diff=function(e,a,n){var u,i=this,h=t.p(a),m=d(e),$=(m.utcOffset()-this.utcOffset())*o,M=this-m,D=function(){return t.m(i,m)};switch(h){case T:u=D()/12;break;case P:u=D();break;case z:u=D()/3;break;case _:u=(M-$)/6048e5;break;case S:u=(M-$)/864e5;break;case j:u=M/w;break;case y:u=M/o;break;case A:u=M/W;break;default:u=M}return n?u:t.a(u)},r.daysInMonth=function(){return this.endOf(P).$D},r.$locale=function(){return b[this.$L]},r.locale=function(e,a){if(!e)return this.$L;var n=this.clone(),u=g(e,a,!0);return u&&(n.$L=u),n},r.clone=function(){return t.w(this.$d,this)},r.toDate=function(){return new Date(this.valueOf())},r.toJSON=function(){return this.isValid()?this.toISOString():null},r.toISOString=function(){return this.$d.toISOString()},r.toString=function(){return this.$d.toUTCString()},s}(),f=q.prototype;return d.prototype=f,[["$ms",O],["$s",A],["$m",y],["$H",j],["$W",S],["$M",P],["$y",T],["$D",Y]].forEach(function(s){f[s[1]]=function(r){return this.$g(r,s[0],s[1])}}),d.extend=function(s,r){return s.$i||(s(r,q,d),s.$i=!0),d},d.locale=g,d.isDayjs=c,d.unix=function(s){return d(1e3*s)},d.en=b[I],d.Ls=b,d.p={},d})}}]);
