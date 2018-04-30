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
        $("#chapterDatagrid").datagrid({
            fit:true,
            fitColumns:true,
            pagination:true,
            pageList:[2,3,5,10],
            pageNumber:1,
            pageSize:3,
            url:'/chapter/queryAll',
            columns:[[
                {title:'名称',field:'name',width:150,align:'center'},
                {title:'大小',field:'size',width:200,align:'center'},
                {title:'路径',field:'src',width:150,align:'center',
                    formatter:function(value){
                        var src ="${pageContext.request.contextPath}/"+value;
                        return  '<audio src="'+src+'" preload="auto" controls="controls"></audio>';
                    }},
                {title:'妙音编号',field:'musicId',width:300,align:'center'},
                {title:'操作',field:'options',width:200,align:'center'
                    ,formatter:function (value,row,index) {
                        return "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_edit'\" onclick=\"openChapterEdit('"+ row.id +"')\">修改</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_delete'\" onclick=\"chapterDelete('"+ row.id +"')\">删除</a>"
                    }
                }
            ]],
            onLoadSuccess:function () {
                $(".option").linkbutton();
            },
            toolbar:'#chapterTools'
        });
    });

    //add
    function chapterSave() {
        $('#chapterSaveForm').form('submit', {
            url:'/chapter/save',
            onSubmit: function(user){
                return $('#chapterSaveForm').form('validate');
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    $('#openChapterSave').dialog('close',true);
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#chapterDatagrid").datagrid('reload');
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Wisdom Save
    function openChapterSave() {
        $("#openChapterSave").dialog({
            top:200,
            width:680,
            title:"保存轮播图",
            iconCls:'icon-photo_edit',
            href:'/back/chapter/saveChapter.jsp',
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:chapterSave
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openChapterSave").dialog('close',true);
                }
            }]
        });
    }

    //update
    function chapterEdit() {
        $('#chapterEditForm').form('submit', {
            url:'/chapter/edit',
            onSubmit: function(user){
                return $('#chapterEditForm').form('validate');
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#chapterDatagrid").datagrid('reload');
                    $('#openChapterEdit').dialog('close',true);
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Update dialog
    function openChapterEdit(id) {
        $("#openChapterEdit").dialog({
            top:200,
            width:700,
            title:"修改用户信息",
            iconCls:'icon-man',
            href:'/back/chapter/editChapter.jsp?id='+id,
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:chapterEdit
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openChapterEdit").dialog('close',true);
                }
            }]
        });
    }

    //delete
    function chapterDelete(id) {
        console.log(id);
        //友情提醒
        $.messager.confirm('提示','确定要删除吗?',function(r){
            if(r){
                //发送ajax请求删除章节
                $.post('/chapter/remove',{"id":id},function(){
                    var tab = $tabs.tabs('getSelected');
                    tab.panel('refresh');
                    $.messager.show({
                        title:'提示',
                        msg:'章节删除成功...',
                    });
                    //刷新datagrid
                    $("#dg").datagrid('reload');//始终保持在当前页展示
                });
            }

        });
    }

    
</script>

<%--datagrid--%>
<table id="chapterDatagrid"></table>

<%--工具栏--%>
<div id="chapterTools">
    <a class="easyui-linkbutton" onclick="openChapterSave();" data-options="iconCls:'icon-photo_add',plain:true">添加</a>
</div>

<%--用户保存对话框--%>
<div id="openChapterSave"></div>

<%--用户修改对话框--%>
<div id="openChapterEdit"></div>