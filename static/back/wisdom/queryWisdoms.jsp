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
        $("#wisdomDatagrid").datagrid({
            fit:true,
            fitColumns:false,
            pagination:true,
            pageList:[2,3,5,10],
            pageNumber:1,
            pageSize:3,
            url:'/wisdom/queryAll',
            columns:[[
                {title:'标题',field:'title',width:150,align:'center'},
                {title:'配图',field:'src',width:200,align:'center',
                    formatter:function(value){
                        var src ="${pageContext.request.contextPath}/"+value;
                        return  '<img src="'+src+'" style="height:30px">';
                    }},
                {title:'创建日期',field:'createDate',width:200,align:'center'},
                {title:'类型',field:'type',width:150,align:'center'},
                {title:'上师编号',field:'guruId',width:200,align:'center'},
                {title:'操作',field:'options',width:300,align:'center'
                    ,formatter:function (value,row,index) {
                        return "<a class='option' data-options=\"plain:true,iconCls:'icon-photo'\" onclick=\"wisdomInline('"+ row.src +"')\">预览</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_edit'\" onclick=\"openWisdomEdit('"+ row.id +"')\">修改</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+
                                "<a class='option' data-options=\"plain:true,iconCls:'icon-photo_delete'\" onclick=\"wisdomDelete('"+ row.id +"')\">删除</a>"
                    }
                },
                {title:'内容',field:'content',width:250,align:'center'}
            ]],
            onLoadSuccess:function () {
                $(".option").linkbutton();
            },
            toolbar:'#wisdomTools'
        });
    });

    //add
    function wisdomSave() {
        $('#wisdomSaveForm').form('submit', {
            url:'/wisdom/save',
            onSubmit: function(wisdom){
                //return $('#wisdomSaveForm').form('validate');
                saveeditor.sync();
                return true;
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    $('#openWisdomSave').dialog('close',true);
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#wisdomDatagrid").datagrid('reload');
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Wisdom Save
    function openWisdomSave() {
        $("#openWisdomSave").dialog({
            top:200,
            width:680,
            title:"保存轮播图",
            iconCls:'icon-photo_edit',
            href:'/back/wisdom/saveWisdom.jsp',
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:wisdomSave
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openWisdomSave").dialog('close',true);
                }
            }]
        });
    }

    function wisdomInline(src) {
        window.open('/user/inline?src='+src);
    }

    //update
    function wisdomEdit() {
        $('#wisdomEditForm').form('submit', {
            url:'/wisdom/edit',
            onSubmit: function(wisdom){
                editeditor.sync();
                return $('#wisdomEditForm').form('validate');
            },
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.success){
                    //var tab = $tabs.tabs('getSelected');
                    //tab.panel('refresh');
                    $("#wisdomDatagrid").datagrid('reload');
                    $('#openWisdomEdit').dialog('close',true);
                } else {
                    alert(data.message)
                }
            }
        });
    }
    //Open Update dialog
    function openWisdomEdit(id) {
        $("#openWisdomEdit").dialog({
            top:200,
            width:700,
            title:"修改用户信息",
            iconCls:'icon-man',
            href:'/back/wisdom/editWisdom.jsp?id='+id,
            buttons:[{
                text:'保存',
                iconCls:'icon-save',
                handler:wisdomEdit
            },{
                text:'关闭',
                iconCls:'icon-cancel',
                handler:function () {
                    $("#openWisdomEdit").dialog('close',true);
                }
            }]
        });
    }

    //delete
    function wisdomDelete(id) {
        console.log(id);
        //友情提醒
        $.messager.confirm('提示','确定要删除吗?',function(r){
            if(r){
                //发送ajax请求删除用户
                $.post('/wisdom/remove',{"id":id},function(){
                    var tab = $tabs.tabs('getSelected');
                    tab.panel('refresh');
                    $.messager.show({
                        title:'提示',
                        msg:'言教或法要删除成功...',
                    });
                    //刷新datagrid
                    $("#dg").datagrid('reload');//始终保持在当前页展示
                });
            }

        });
    }

    //处理搜索
    function search(value,name) {

        if (!$tabs.tabs('exists','lucene')) {
            $tabs.tabs('add',{
                title:"lucene",
                iconCls:"icon-user",
                closable:true,
                href:"/back/wisdom/queryWisdomsLucene.jsp?field="+name+"&value="+value,
                ///back/wisdom/queryWisdomsLucene?field="+name+"&value"+value
                tools:[{
                    //刷新按钮
                    iconCls:'icon-reload',
                    handler:function () {
                        var tab = $tabs.tabs('getSelected');
                        tab.panel('refresh');
                    }
                }]

            });
        }else{

            //$('#wisdomLuceneDatagrid').datagrid('load',"/back/wisdom/queryWisdomsLucene.jsp?field="+name+"&value="+value);
            $("#wisdomLuceneDatagrid").datagrid({
                url:"/back/wisdom/queryWisdomsLucene.jsp?field="+name+"&value="+value,
            });

            $tabs.tabs('select','lucene');

            /*$('#dg').datagrid('load',{
                code: '01',
                name: 'name01'
            });*/


        }
    }

    
</script>

<%--datagrid--%>
<table id="wisdomDatagrid"></table>

<%--工具栏--%>
<div id="wisdomTools">
    <a class="easyui-linkbutton" onclick="openWisdomSave();" data-options="iconCls:'icon-photo_add',plain:true">添加</a>
    <%--搜索框--%>
    <input id="search" class="easyui-searchbox" data-options="searcher:search,prompt:'请输入查询条件',menu:'#condition',width:400"></div>
    <%--搜索框条件--%>
    <div id="condition" style="height: auto">
        <%--<div data-options="name:'id'">id</div>--%>
        <div data-options="name:'title'">title</div>
        <div data-options="name:'content'">content</div>
        <div data-options="name:'src'">src</div>
        <div data-options="name:'createDate'">createDate</div>
        <div data-options="name:'type'">type</div>
        <div data-options="name:'groupId'">groupId</div>
    </div>
</div>

<%--用户保存对话框--%>
<div id="openWisdomSave"></div>

<%--用户修改对话框--%>
<div id="openWisdomEdit"></div>