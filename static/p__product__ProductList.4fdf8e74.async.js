(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[862],{86796:function(re,W,o){"use strict";o.r(W),o.d(W,{default:function(){return a}});var Z=o(8963),T=o(38291),A=o(88983),y=o(47933),j=o(90636),S=o(11849),_=o(34792),w=o(48086),z=o(3182),E=o(49111),Y=o(19650),L=o(2824),ne=o(71194),Q=o(50146),G=o(67294),N=o(36773),X=o(69083),I=o(27484),b=o.n(I),J=o(22122),c=o(83707),g=o(65734),l=function(u,i){return G.createElement(g.Z,(0,J.Z)({},u,{ref:i,icon:c.Z}))},t=G.forwardRef(l),q=o(81910),h=o(85893),s={0:"\u5168\u90E8",1:"\u5728\u552E\u4E2D",2:"\u5DF2\u4E0B\u67B6",3:"\u5DF2\u552E\u7F44"},r=Q.Z.confirm,e=function(){var u=(0,G.useState)([]),i=(0,L.Z)(u,2),d=i[0],m=i[1],$=(0,G.useState)({pageNum:1,pageSize:10,queryGoodsListStatus:0}),M=(0,L.Z)($,2),D=M[0],F=M[1],U=(0,G.useState)({showSizeChanger:!0,showQuickJumper:!0,showTotal:function(v){return"\u603B\u5171 ".concat(v," \u6761")}}),K=(0,L.Z)(U,2),R=K[0],C=K[1],H=(0,G.useState)(!1),V=(0,L.Z)(H,2),ee=V[0],ae=V[1],ie=[{title:"\u5546\u54C1\u4FE1\u606F",dataIndex:"info",key:"info",width:340,render:function(v,p){var f=p.key,P=p.info;return(0,h.jsxs)("div",{style:{display:"flex",cursor:"pointer",alignItems:"center"},onClick:function(){return ue(f)},children:[(0,h.jsx)("img",{src:P.goodsFrontImage,alt:"picture",style:{width:"60px",height:"60px",marginRight:"8px"}}),(0,h.jsxs)("div",{style:{display:"flex",flexDirection:"column",justifyContent:"space-between",padding:"4px"},children:[(0,h.jsx)("span",{style:{color:"#1890ff",wordBreak:"break-word",wordWrap:"break-word",whiteSpace:"pre-wrap"},children:P.name}),(0,h.jsxs)("div",{children:[(0,h.jsx)("span",{children:"\u5546\u54C1ID\uFF1A"}),(0,h.jsx)("span",{children:P.spuCode})]})]})]})}},{title:"\u72B6\u6001",dataIndex:"status",key:"status",width:100,render:function(v,p){var f=p.status;return(0,h.jsx)("span",{children:s==null?void 0:s[f]})}},{title:"\u91C7\u8D2D\u4EF7\u683C",dataIndex:"buyPrice",key:"buyPrice",render:function(v,p){var f=p.buyPrice;return(0,h.jsx)("span",{children:"\uFFE5".concat(f.minPrice,"~\uFFE5").concat(f.maxPrice)})}},{title:"\u96F6\u552E\u4EF7\u683C",dataIndex:"retailPrice",key:"retailPrice",render:function(v,p){var f=p.retailPrice;return(0,h.jsx)("span",{children:"\uFFE5".concat(f.minPrice,"~\uFFE5").concat(f.maxPrice)})}},{title:"\u96F6\u552E\u9500\u91CF",dataIndex:"retailNum",key:"retailNum",width:100},{title:"\u96F6\u552E\u5E93\u5B58",dataIndex:"stock",key:"stock",width:100},{title:"\u521B\u5EFA\u65F6\u95F4",dataIndex:"createTime",key:"createTime",render:function(v,p){var f=p.createTime;return(0,h.jsx)("span",{children:b()(f).format("YYYY-MM-DD HH:mm:ss")})}},{title:"\u64CD\u4F5C",dataIndex:"action",key:"action",width:100,render:function(v,p){var f=p.key,P=p.status;return(0,h.jsx)(Y.Z,{children:(0,h.jsxs)("div",{style:{display:"flex",flexDirection:"column"},children:[P===2&&(0,h.jsx)("a",{onClick:function(){return se(f,1)},children:"\u4E0A\u67B6"}),P===1&&(0,h.jsx)("a",{onClick:function(){return se(f,2)},children:"\u4E0B\u67B6"}),(P===1||P===2||P===3)&&(0,h.jsx)("a",{onClick:function(){return oe(f,P)},children:"\u5220\u9664"})]})})}}],ue=function(v){q.m8.push("/product/detail/".concat(v))},se=function(){var k=(0,z.Z)((0,j.Z)().mark(function v(p,f){return(0,j.Z)().wrap(function(x){for(;;)switch(x.prev=x.next){case 0:return x.prev=0,x.next=3,(0,X.pm)({goodsID:p,goodsStatus:f});case 3:w.ZP.success("\u64CD\u4F5C\u6210\u529F"),F((0,S.Z)({},D)),x.next=10;break;case 7:x.prev=7,x.t0=x.catch(0),console.log(x.t0);case 10:case"end":return x.stop()}},v,null,[[0,7]])}));return function(p,f){return k.apply(this,arguments)}}(),oe=function(v,p){r({title:"\u5220\u9664\u786E\u8BA4",icon:(0,h.jsx)(t,{}),content:"\u786E\u8BA4\u5220\u9664\u8BE5\u5546\u54C1\u5417\uFF1F",onOk:function(){return(0,z.Z)((0,j.Z)().mark(function P(){return(0,j.Z)().wrap(function(B){for(;;)switch(B.prev=B.next){case 0:return B.prev=0,B.next=3,(0,X.ys)({goodsID:v,goodsStatus:p});case 3:w.ZP.success("\u5220\u9664\u6210\u529F"),F((0,S.Z)({},D)),B.next=10;break;case 7:B.prev=7,B.t0=B.catch(0),console.log(B.t0);case 10:case"end":return B.stop()}},P,null,[[0,7]])}))()},onCancel:function(){}})},ce=function(){var k=(0,z.Z)((0,j.Z)().mark(function v(p){var f,P,x;return(0,j.Z)().wrap(function(te){for(;;)switch(te.prev=te.next){case 0:return ae(!0),te.next=3,(0,X.k1)(p);case 3:f=te.sent,ae(!1),P=(0,S.Z)((0,S.Z)({},R),{},{current:f.pageNum,pageSize:f.pageSize,total:f.totalNum}),C(P),x=f.goodsList.map(function(O){return{key:O.id,info:{goodsFrontImage:O.goodsFrontImage,name:O.name,spuCode:O.spuCode},status:O==null?void 0:O.status,buyPrice:{minPrice:O.wholesalePriceMin,maxPrice:O.wholesalePriceMax},retailPrice:{minPrice:O.retailPriceMin,maxPrice:O.retailPriceMax},retailNum:O.soldNum,createTime:O.createdAt,updateTime:O.updatedAt}}),m(x);case 9:case"end":return te.stop()}},v)}));return function(p){return k.apply(this,arguments)}}();(0,G.useEffect)(function(){ce(D)},[D]);var le=function(v){var p=v.current,f=v.pageSize,P=(0,S.Z)((0,S.Z)({},D),{},{pageNum:p,pageSize:f});F(P)},de=function(v){var p={pageNum:1,pageSize:D.pageSize,queryGoodsListStatus:v.target.value};F(p)};return(0,h.jsxs)(N.ZP,{children:[(0,h.jsxs)(y.ZP.Group,{defaultValue:0,buttonStyle:"solid",style:{marginBottom:"24px"},size:"large",onChange:de,children:[(0,h.jsx)(y.ZP.Button,{value:0,children:"\u5168\u90E8"}),(0,h.jsx)(y.ZP.Button,{value:1,children:"\u5728\u552E\u4E2D"}),(0,h.jsx)(y.ZP.Button,{value:2,children:"\u5DF2\u4E0B\u67B6"}),(0,h.jsx)(y.ZP.Button,{value:3,children:"\u5DF2\u552E\u7F44"})]}),(0,h.jsx)(T.Z,{columns:ie,dataSource:d,onChange:le,pagination:R,loading:ee,scroll:{x:"max-content"}})]})},a=e},69083:function(re,W,o){"use strict";o.d(W,{g2:function(){return j},k1:function(){return z},mZ:function(){return Y},BH:function(){return ne},JJ:function(){return G},pm:function(){return X},ys:function(){return b}});var Z=o(90636),T=o(3182),A=o(99871),y=o(636);function j(c){return S.apply(this,arguments)}function S(){return S=(0,T.Z)((0,Z.Z)().mark(function c(g){return(0,Z.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/getMyGoodsDetail?".concat((0,A.R)(g))));case 1:case"end":return t.stop()}},c)})),S.apply(this,arguments)}function _(c){return w.apply(this,arguments)}function w(){return w=_asyncToGenerator(_regeneratorRuntime().mark(function c(g){return _regeneratorRuntime().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",request("/api/v1/goods/getGoodsList?".concat(objectToUrlParams(g))));case 1:case"end":return t.stop()}},c)})),w.apply(this,arguments)}function z(c){return E.apply(this,arguments)}function E(){return E=(0,T.Z)((0,Z.Z)().mark(function c(g){return(0,Z.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/getMyGoodsList?".concat((0,A.R)(g))));case 1:case"end":return t.stop()}},c)})),E.apply(this,arguments)}function Y(c){return L.apply(this,arguments)}function L(){return L=(0,T.Z)((0,Z.Z)().mark(function c(g){return(0,Z.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/common/area?".concat((0,A.R)(g))));case 1:case"end":return t.stop()}},c)})),L.apply(this,arguments)}function ne(c){return Q.apply(this,arguments)}function Q(){return Q=(0,T.Z)((0,Z.Z)().mark(function c(g){return(0,Z.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/mall/categories?".concat((0,A.R)(g))));case 1:case"end":return t.stop()}},c)})),Q.apply(this,arguments)}function G(c){return N.apply(this,arguments)}function N(){return N=(0,T.Z)((0,Z.Z)().mark(function c(g){return(0,Z.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/addGoods",{method:"POST",data:g}));case 1:case"end":return t.stop()}},c)})),N.apply(this,arguments)}function X(c){return I.apply(this,arguments)}function I(){return I=(0,T.Z)((0,Z.Z)().mark(function c(g){return(0,Z.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/modifyMyGoodsStatus",{method:"POST",data:g}));case 1:case"end":return t.stop()}},c)})),I.apply(this,arguments)}function b(c){return J.apply(this,arguments)}function J(){return J=(0,T.Z)((0,Z.Z)().mark(function c(g){return(0,Z.Z)().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,y.Z)("/api/v1/goods/deleteMyGoods",{method:"DELETE",data:g}));case 1:case"end":return t.stop()}},c)})),J.apply(this,arguments)}},27484:function(re){(function(W,o){re.exports=o()})(this,function(){"use strict";var W=1e3,o=6e4,Z=36e5,T="millisecond",A="second",y="minute",j="hour",S="day",_="week",w="month",z="quarter",E="year",Y="date",L="Invalid Date",ne=/^(\d{4})[-/]?(\d{1,2})?[-/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$/,Q=/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g,G={name:"en",weekdays:"Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),months:"January_February_March_April_May_June_July_August_September_October_November_December".split("_"),ordinal:function(s){var r=["th","st","nd","rd"],e=s%100;return"["+s+(r[(e-20)%10]||r[e]||r[0])+"]"}},N=function(s,r,e){var a=String(s);return!a||a.length>=r?s:""+Array(r+1-a.length).join(e)+s},X={s:N,z:function(s){var r=-s.utcOffset(),e=Math.abs(r),a=Math.floor(e/60),n=e%60;return(r<=0?"+":"-")+N(a,2,"0")+":"+N(n,2,"0")},m:function s(r,e){if(r.date()<e.date())return-s(e,r);var a=12*(e.year()-r.year())+(e.month()-r.month()),n=r.clone().add(a,w),u=e-n<0,i=r.clone().add(a+(u?-1:1),w);return+(-(a+(e-n)/(u?n-i:i-n))||0)},a:function(s){return s<0?Math.ceil(s)||0:Math.floor(s)},p:function(s){return{M:w,y:E,w:_,d:S,D:Y,h:j,m:y,s:A,ms:T,Q:z}[s]||String(s||"").toLowerCase().replace(/s$/,"")},u:function(s){return s===void 0}},I="en",b={};b[I]=G;var J="$isDayjsObject",c=function(s){return s instanceof q||!(!s||!s[J])},g=function s(r,e,a){var n;if(!r)return I;if(typeof r=="string"){var u=r.toLowerCase();b[u]&&(n=u),e&&(b[u]=e,n=u);var i=r.split("-");if(!n&&i.length>1)return s(i[0])}else{var d=r.name;b[d]=r,n=d}return!a&&n&&(I=n),n||!a&&I},l=function(s,r){if(c(s))return s.clone();var e=typeof r=="object"?r:{};return e.date=s,e.args=arguments,new q(e)},t=X;t.l=g,t.i=c,t.w=function(s,r){return l(s,{locale:r.$L,utc:r.$u,x:r.$x,$offset:r.$offset})};var q=function(){function s(e){this.$L=g(e.locale,null,!0),this.parse(e),this.$x=this.$x||e.x||{},this[J]=!0}var r=s.prototype;return r.parse=function(e){this.$d=function(a){var n=a.date,u=a.utc;if(n===null)return new Date(NaN);if(t.u(n))return new Date;if(n instanceof Date)return new Date(n);if(typeof n=="string"&&!/Z$/i.test(n)){var i=n.match(ne);if(i){var d=i[2]-1||0,m=(i[7]||"0").substring(0,3);return u?new Date(Date.UTC(i[1],d,i[3]||1,i[4]||0,i[5]||0,i[6]||0,m)):new Date(i[1],d,i[3]||1,i[4]||0,i[5]||0,i[6]||0,m)}}return new Date(n)}(e),this.init()},r.init=function(){var e=this.$d;this.$y=e.getFullYear(),this.$M=e.getMonth(),this.$D=e.getDate(),this.$W=e.getDay(),this.$H=e.getHours(),this.$m=e.getMinutes(),this.$s=e.getSeconds(),this.$ms=e.getMilliseconds()},r.$utils=function(){return t},r.isValid=function(){return this.$d.toString()!==L},r.isSame=function(e,a){var n=l(e);return this.startOf(a)<=n&&n<=this.endOf(a)},r.isAfter=function(e,a){return l(e)<this.startOf(a)},r.isBefore=function(e,a){return this.endOf(a)<l(e)},r.$g=function(e,a,n){return t.u(e)?this[a]:this.set(n,e)},r.unix=function(){return Math.floor(this.valueOf()/1e3)},r.valueOf=function(){return this.$d.getTime()},r.startOf=function(e,a){var n=this,u=!!t.u(a)||a,i=t.p(e),d=function(R,C){var H=t.w(n.$u?Date.UTC(n.$y,C,R):new Date(n.$y,C,R),n);return u?H:H.endOf(S)},m=function(R,C){return t.w(n.toDate()[R].apply(n.toDate("s"),(u?[0,0,0,0]:[23,59,59,999]).slice(C)),n)},$=this.$W,M=this.$M,D=this.$D,F="set"+(this.$u?"UTC":"");switch(i){case E:return u?d(1,0):d(31,11);case w:return u?d(1,M):d(0,M+1);case _:var U=this.$locale().weekStart||0,K=($<U?$+7:$)-U;return d(u?D-K:D+(6-K),M);case S:case Y:return m(F+"Hours",0);case j:return m(F+"Minutes",1);case y:return m(F+"Seconds",2);case A:return m(F+"Milliseconds",3);default:return this.clone()}},r.endOf=function(e){return this.startOf(e,!1)},r.$set=function(e,a){var n,u=t.p(e),i="set"+(this.$u?"UTC":""),d=(n={},n[S]=i+"Date",n[Y]=i+"Date",n[w]=i+"Month",n[E]=i+"FullYear",n[j]=i+"Hours",n[y]=i+"Minutes",n[A]=i+"Seconds",n[T]=i+"Milliseconds",n)[u],m=u===S?this.$D+(a-this.$W):a;if(u===w||u===E){var $=this.clone().set(Y,1);$.$d[d](m),$.init(),this.$d=$.set(Y,Math.min(this.$D,$.daysInMonth())).$d}else d&&this.$d[d](m);return this.init(),this},r.set=function(e,a){return this.clone().$set(e,a)},r.get=function(e){return this[t.p(e)]()},r.add=function(e,a){var n,u=this;e=Number(e);var i=t.p(a),d=function(M){var D=l(u);return t.w(D.date(D.date()+Math.round(M*e)),u)};if(i===w)return this.set(w,this.$M+e);if(i===E)return this.set(E,this.$y+e);if(i===S)return d(1);if(i===_)return d(7);var m=(n={},n[y]=o,n[j]=Z,n[A]=W,n)[i]||1,$=this.$d.getTime()+e*m;return t.w($,this)},r.subtract=function(e,a){return this.add(-1*e,a)},r.format=function(e){var a=this,n=this.$locale();if(!this.isValid())return n.invalidDate||L;var u=e||"YYYY-MM-DDTHH:mm:ssZ",i=t.z(this),d=this.$H,m=this.$m,$=this.$M,M=n.weekdays,D=n.months,F=n.meridiem,U=function(C,H,V,ee){return C&&(C[H]||C(a,u))||V[H].slice(0,ee)},K=function(C){return t.s(d%12||12,C,"0")},R=F||function(C,H,V){var ee=C<12?"AM":"PM";return V?ee.toLowerCase():ee};return u.replace(Q,function(C,H){return H||function(V){switch(V){case"YY":return String(a.$y).slice(-2);case"YYYY":return t.s(a.$y,4,"0");case"M":return $+1;case"MM":return t.s($+1,2,"0");case"MMM":return U(n.monthsShort,$,D,3);case"MMMM":return U(D,$);case"D":return a.$D;case"DD":return t.s(a.$D,2,"0");case"d":return String(a.$W);case"dd":return U(n.weekdaysMin,a.$W,M,2);case"ddd":return U(n.weekdaysShort,a.$W,M,3);case"dddd":return M[a.$W];case"H":return String(d);case"HH":return t.s(d,2,"0");case"h":return K(1);case"hh":return K(2);case"a":return R(d,m,!0);case"A":return R(d,m,!1);case"m":return String(m);case"mm":return t.s(m,2,"0");case"s":return String(a.$s);case"ss":return t.s(a.$s,2,"0");case"SSS":return t.s(a.$ms,3,"0");case"Z":return i}return null}(C)||i.replace(":","")})},r.utcOffset=function(){return 15*-Math.round(this.$d.getTimezoneOffset()/15)},r.diff=function(e,a,n){var u,i=this,d=t.p(a),m=l(e),$=(m.utcOffset()-this.utcOffset())*o,M=this-m,D=function(){return t.m(i,m)};switch(d){case E:u=D()/12;break;case w:u=D();break;case z:u=D()/3;break;case _:u=(M-$)/6048e5;break;case S:u=(M-$)/864e5;break;case j:u=M/Z;break;case y:u=M/o;break;case A:u=M/W;break;default:u=M}return n?u:t.a(u)},r.daysInMonth=function(){return this.endOf(w).$D},r.$locale=function(){return b[this.$L]},r.locale=function(e,a){if(!e)return this.$L;var n=this.clone(),u=g(e,a,!0);return u&&(n.$L=u),n},r.clone=function(){return t.w(this.$d,this)},r.toDate=function(){return new Date(this.valueOf())},r.toJSON=function(){return this.isValid()?this.toISOString():null},r.toISOString=function(){return this.$d.toISOString()},r.toString=function(){return this.$d.toUTCString()},s}(),h=q.prototype;return l.prototype=h,[["$ms",T],["$s",A],["$m",y],["$H",j],["$W",S],["$M",w],["$y",E],["$D",Y]].forEach(function(s){h[s[1]]=function(r){return this.$g(r,s[0],s[1])}}),l.extend=function(s,r){return s.$i||(s(r,q,l),s.$i=!0),l},l.locale=g,l.isDayjs=c,l.unix=function(s){return l(1e3*s)},l.en=b[I],l.Ls=b,l.p={},l})}}]);
