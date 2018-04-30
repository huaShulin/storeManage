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
        $("#wisdomLuceneDatagrid").datagrid({
            fit:true,
            fitColumns:false,
            url:'/wisdom/lucene?field=${param.field}&value=${param.value}',
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
                {title:'内容',field:'content',width:250,align:'center'}
            ]],
            onLoadSuccess:function () {
                $(".option").linkbutton();
            }
        });
    });
    
</script>

<%--datagrid--%>
<table id="wisdomLuceneDatagrid"></table>