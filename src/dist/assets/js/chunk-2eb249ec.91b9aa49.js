(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2eb249ec"],{4528:function(t,e,a){"use strict";var i=a("d286"),l=a.n(i);l.a},a2c6:function(t,e,a){"use strict";a.r(e);var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"edit"},[a("Spin",{directives:[{name:"show",rawName:"v-show",value:t.loading,expression:"loading"}],attrs:{fix:"",size:"large"}},[a("loading")],1),a("div",{staticClass:"labelTitle"},[t._v("\n    文章信息编辑"),a("Button",{attrs:{type:"success"},on:{click:t.goback}},[t._v("返回文章页")])],1),a("div",{staticClass:"title"},[a("label",{attrs:{for:""}},[t._v("文章名称：")]),a("Input",{staticStyle:{width:"300px"},attrs:{type:"text",clearable:"",placeholder:"请输入文章名称",size:"large"},model:{value:t.title,callback:function(e){t.title=e},expression:"title"}})],1),a("div",{staticClass:"title"},[a("label",{attrs:{for:""}},[t._v("文章分类：")]),a("Select",{staticStyle:{width:"300px"},attrs:{"label-in-value":!0,size:"large"},on:{"on-change":t.change},model:{value:t.label.id,callback:function(e){t.$set(t.label,"id",e)},expression:"label.id"}},t._l(t.labelArr,function(e,i){return a("Option",{key:i,attrs:{value:e.id}},[t._v("\n                "+t._s(e.labelname)+"\n            ")])}))],1),a("div",{staticClass:"article"},[a("label",{attrs:{for:""}},[t._v("文章内容：")]),a("div",{staticClass:"wangedit"},[a("i-editor",{attrs:{autosize:{minRows:20}},model:{value:t.content,callback:function(e){t.content=e},expression:"content"}})],1)]),a("div",{staticClass:"submit"},[a("Button",{staticStyle:{width:"10%"},attrs:{type:"info"},on:{click:t.submit}},[t._v("提交")])],1)],1)},l=[],n={data:function(){return{title:"",label:{labelname:"",id:""},content:"",labelArr:[],loading:!1,article:{},articleId:""}},created:function(){var t=this;this.$route.params.article?(localStorage.setItem("mid_blog_cbim",JSON.stringify(this.$route.params.article)),this.article=this.$route.params.article):this.article=JSON.parse(localStorage.getItem("mid_blog_cbim")),this.title=this.article.articleTitle,this.label.labelname=this.article.articleLabel,this.label.id=this.article.labelId,this.content=this.article.articleInfo,this.articleId=this.article.articleId,this.loading=!0,this.remote({url:"/labelInfo",method:"get"}).then(function(e){t.labelArr=e.data,t.loading=!1},function(e){t.loading=!1})},methods:{change:function(t){this.label.labelname=t.label},goback:function(){this.$router.back(-1)},submit:function(){var t=this,e="";e=this.articleId?"修改":"新增",this.title&&this.label&&this.content?this.$Modal.confirm({render:function(t){return t("div",{props:{style:{textAlign:"center"}}},"确认"+e+"吗？")},onOk:function(){t.loading=!0,t.remote({url:"/addArticle",method:"POST",data:{articleName:t.title,articleLabel:t.label.labelname,articleInfo:t.content,labelId:t.label.id,articleId:t.articleId}}).then(function(e){t.loading=!1,200==e.code?(t.$Message.success("操作成功"),t.$router.push({path:"/main/home"})):t.$Message.success("操作失败")},function(e){t.loading=!1})}}):this.$Message.warning("文章名称，文章分类，文章内容不能为空")}}},s=n,c=(a("4528"),a("2877")),r=Object(c["a"])(s,i,l,!1,null,"2ee40884",null);r.options.__file="edit.vue";e["default"]=r.exports},d286:function(t,e,a){}}]);
//# sourceMappingURL=chunk-2eb249ec.91b9aa49.js.map