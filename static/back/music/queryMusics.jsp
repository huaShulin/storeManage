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
        $("#musicDatagrid").datagrid({
            fit:true,
            fitColumns:true,
            pagination:true,
            pageList:[2,3,5,10],
            pageNumber:1,
            pageSize:3,
            url:'/music/queryAll',
            columns:[[
                {title:'名称',field:'name',width:100,align:'center'},
                {title:'封面',field:'src',width:140,align:'center',
                    formatter:function(value){
                        var src ="${pageContext.request.contextPath}/"+value;
                        return  '<img src="'+src+'" style="height:30px">';
                    }},
                {title:'作者',field:'writer',width:90,align:'center'},
                {title:'播音',field:'announcer',width:90,align:'center'},
                {title:'总集数',field:'sum',width:50,align:'center'},
                {title:'创建日期',field:'createDate',width:140,align:'center'},
                {title:'评分',field:'score',width:50,align:'center'},
                {title:'操作',field:'options',width:200,align:'center'
                    ,formatter:function (value,row,index) {
                        return "<a class='option' data-options=\"plain:true,iconCls:'icon-photo'\" onclick=\"musicInline('"+ row.src +"')\">预览</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_edit'\" onclick=\"openMusicEdit('"+ row.id +"')\">修改</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_delete'\" onclick=\"musicDelete('"+ row.id +"')\">删除</a>"
                    }
                },
                {title:'描述',field:'content',width:140,align:'center'}
            ]],
            onLoadSuccess:function () {
                $(".option").linkbutton();
            },
            toolbar:'#musicTools'
        });
    });

    //add
    function musicSave() {
        $('#musicSaveForm').form('submit', {
            url:'/music/save',
            onSubmit: function(user){
                //return $('#guruSaveForm').form('validate');
                return true;
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    $('#openMusicSave').dialog('close',true);
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#musicDatagrid").datagrid('reload');
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Music Save
    function openMusicSave() {
        $("#openMusicSave").dialog({
            top:200,
            width:680,
            title:"保存轮播图",
            iconCls:'icon-photo_edit',
            href:'/back/music/saveMusic.jsp',
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:musicSave
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openMusicSave").dialog('close',true);
                }
            }]
        });
    }

    function musicInline(src) {
        window.open('/music/inline?src='+src);
    }

    //update
    function musicEdit() {
        $('#musicEditForm').form('submit', {
            url:'/music/edit',
            onSubmit: function(user){
                return $('#musicEditForm').form('validate');
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#musicDatagrid").datagrid('reload');
                    $('#openMusicEdit').dialog('close',true);
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Update dialog
    function openMusicEdit(id) {
        $("#openMusicEdit").dialog({
            top:200,
            width:700,
            title:"修改用户信息",
            iconCls:'icon-man',
            href:'/back/music/editMusic.jsp?id='+id,
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:musicEdit
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openMusicEdit").dialog('close',true);
                }
            }]
        });
    }

    //delete
    function musicDelete(id) {
        console.log(id);
        //友情提醒
        $.messager.confirm('提示','确定要删除吗?',function(r){
            if(r){
                //发送ajax请求删除用户
                $.post('/music/remove',{"id":id},function(){
                    var tab = $tabs.tabs('getSelected');
                    tab.panel('refresh');
                    $.messager.show({
                        title:'提示',
                        msg:'吉祥妙音删除成功...',
                    });
                    //刷新datagrid
                    $("#dg").datagrid('reload');//始终保持在当前页展示
                });
            }

        });
    }

    
</script>

<%--datagrid--%>
<table id="musicDatagrid"></table>

<%--工具栏--%>
<div id="musicTools">
    <a class="easyui-linkbutton" onclick="openMusicSave();" data-options="iconCls:'icon-photo_add',plain:true">添加</a>
</div>

<%--用户保存对话框--%>
<div id="openMusicSave"></div>

<%--用户修改对话框--%>
<div id="openMusicEdit"></div>