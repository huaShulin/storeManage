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
        $("#countDatagrid").datagrid({
            fit:true,
            fitColumns:true,
            pagination:true,
            pageList:[2,3,5,10],
            pageNumber:1,
            pageSize:3,
            url:'/count/queryAll',
            columns:[[
                {title:'计数器',field:'name',width:100,align:'center'},
                {title:'当前计数',field:'count',width:200,align:'center'},
                {title:'创建时间',field:'createDate',width:100,align:'center'},
                {title:'所属功课',field:'lessonId',width:100,align:'center'}
            ]],
            onLoadSuccess:function () {
                $(".option").linkbutton();
            }
        });
    });
    
</script>

<%--datagrid--%>
<table id="countDatagrid"></table>
