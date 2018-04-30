<%--
  Created by IntelliJ IDEA.
  User: Shulin
  Date: 14/11/2017
  Time: 下午 3:58
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" pageEncoding="UTF-8" language="java" %>

<script type="text/javascript">

    $(function () {
        $("#bannerDatagrid").datagrid({
            fit:true,
            fitColumns:true,
            pagination:true,
            pageList:[2,3,5,10],
            pageNumber:1,
            pageSize:3,
            url:'/banner/queryAll',
            columns:[[
                {title:'名称',field:'name',width:100,align:'center'},
                {title:'路径',field:'src',width:140,align:'center',
                    formatter:function(value){
                        var src ="${pageContext.request.contextPath}/"+value;
                        return  '<img src="'+src+'" style="height:30px">';
                }},

                {title:'状态',field:'status',width:100,align:'center'},
                {title:'操作',field:'options',width:120,align:'center'
                    ,formatter:function (value,row,index) {
                        return "<a class='option' data-options=\"plain:true,iconCls:'icon-photo'\" onclick=\"inline('"+ row.src +"')\">预览</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_edit'\" onclick=\"openBannerEdit('"+ row.id +"')\">修改</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_delete'\" onclick=\"userDelete('"+ row.id +"')\">删除</a>"
                    }
                }
            ]],
            onLoadSuccess:function () {
                $(".option").linkbutton();
            },
            toolbar:'#bannerTools'
        });
    });

    //add
    function bannerSave() {
        $('#bannerSaveForm').form('submit', {
            url:'/banner/save',
            onSubmit: function(){
                return $('#bannerSaveForm').form('validate');
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    $('#openBannerSave').dialog('close',true);
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#bannerDatagrid").datagrid('reload');
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Banner Save
    function openBannerSave() {
        $("#openBannerSave").dialog({
            top:200,
            width:680,
            title:"保存轮播图",
            iconCls:'icon-photo_edit',
            href:'/back/banner/saveBanner.jsp',
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:bannerSave
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openBannerSave").dialog('close',true);
                }
            }]
        });
    }

    function inline(src) {
        window.open('/user/inline?src='+src);
    }

    //update
    function bannerEdit() {
        $('#bannerEditForm').form('submit', {
            url:'/banner/edit',
            onSubmit: function(user){
                return $('#bannerEditForm').form('validate');
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#bannerDatagrid").datagrid('reload');
                    $('#openBannerEdit').dialog('close',true);
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Update dialog
    function openBannerEdit(id) {
        $("#openBannerEdit").dialog({
            top:200,
            width:700,
            title:"修改用户信息",
            iconCls:'icon-man',
            href:'/back/banner/editBanner.jsp?id='+id,
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:bannerEdit
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openBannerEdit").dialog('close',true);
                }
            }]
        });
    }

    //delete
    function userDelete(id) {
        console.log(id);
        //友情提醒
        $.messager.confirm('提示','确定要删除吗?',function(r){
            if(r){
                //发送ajax请求删除用户
                $.post('/banner/remove',{"id":id},function(){
                    var tab = $tabs.tabs('getSelected');
                    tab.panel('refresh');
                    $.messager.show({
                        title:'提示',
                        msg:'轮播图删除成功...',
                    });
                    //刷新datagrid
                    $("#dg").datagrid('reload');//始终保持在当前页展示
                });
            }

        });
    }

    
</script>

<%--datagrid--%>
<table id="bannerDatagrid"></table>

<%--工具栏--%>
<div id="bannerTools">
    <a class="easyui-linkbutton" onclick="openBannerSave();" data-options="iconCls:'icon-photo_add',plain:true">添加</a>
</div>

<%--用户保存对话框--%>
<div id="openBannerSave"></div>

<%--用户修改对话框--%>
<div id="openBannerEdit"></div>