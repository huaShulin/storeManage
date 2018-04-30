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
        $("#guruDatagrid").datagrid({
            fit:true,
            fitColumns:true,
            pagination:true,
            pageList:[2,3,5,10],
            pageNumber:1,
            pageSize:3,
            url:'/guru/queryAll',
            columns:[[
                {title:'名称',field:'name',width:100,align:'center'},
                {title:'头像',field:'src',width:140,align:'center',
                    formatter:function(value){
                        var src ="${pageContext.request.contextPath}/"+value;
                        return  '<img src="'+src+'" style="height:30px">';
                    }},
                {title:'操作',field:'options',width:120,align:'center'
                    ,formatter:function (value,row,index) {
                        return "<a class='option' data-options=\"plain:true,iconCls:'icon-photo'\" onclick=\"guruInline('"+ row.src +"')\">预览</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_edit'\" onclick=\"openGuruEdit('"+ row.id +"')\">修改</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_delete'\" onclick=\"guruDelete('"+ row.id +"')\">删除</a>"
                    }
                }
            ]],
            onLoadSuccess:function () {
                $(".option").linkbutton();
            },
            toolbar:'#guruTools'
        });
    });

    //add
    function guruSave() {
        $('#guruSaveForm').form('submit', {
            url:'/guru/save',
            onSubmit: function(user){
                //return $('#guruSaveForm').form('validate');
                return true;
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    $('#openGuruSave').dialog('close',true);
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#guruDatagrid").datagrid('reload');
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Guru Save
    function openGuruSave() {
        $("#openGuruSave").dialog({
            top:200,
            width:680,
            title:"保存轮播图",
            iconCls:'icon-photo_edit',
            href:'/back/guru/saveGuru.jsp',
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:guruSave
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openGuruSave").dialog('close',true);
                }
            }]
        });
    }

    function guruInline(src) {
        window.open('/user/inline?src='+src);
    }

    //update
    function guruEdit() {
        $('#guruEditForm').form('submit', {
            url:'/guru/edit',
            onSubmit: function(user){
                return $('#guruEditForm').form('validate');
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#guruDatagrid").datagrid('reload');
                    $('#openGuruEdit').dialog('close',true);
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Update dialog
    function openGuruEdit(id) {
        $("#openGuruEdit").dialog({
            top:200,
            width:700,
            title:"修改用户信息",
            iconCls:'icon-man',
            href:'/back/guru/editGuru.jsp?id='+id,
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:guruEdit
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openGuruEdit").dialog('close',true);
                }
            }]
        });
    }

    //delete
    function guruDelete(id) {
        console.log(id);
        //友情提醒
        $.messager.confirm('提示','确定要删除吗?',function(r){
            if(r){
                //发送ajax请求删除用户
                $.post('/guru/remove',{"id":id},function(){
                    var tab = $tabs.tabs('getSelected');
                    tab.panel('refresh');
                    $.messager.show({
                        title:'提示',
                        msg:'上师删除成功...',
                    });
                    //刷新datagrid
                    $("#dg").datagrid('reload');//始终保持在当前页展示
                });
            }

        });
    }

    
</script>

<%--datagrid--%>
<table id="guruDatagrid"></table>

<%--工具栏--%>
<div id="guruTools">
    <a class="easyui-linkbutton" onclick="openGuruSave();" data-options="iconCls:'icon-photo_add',plain:true">添加</a>
</div>

<%--用户保存对话框--%>
<div id="openGuruSave"></div>

<%--用户修改对话框--%>
<div id="openGuruEdit"></div>