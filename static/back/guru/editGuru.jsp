<%--
  Created by IntelliJ IDEA.
  User: Shulin
  Date: 14/11/2017
  Time: 下午 11:51
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" isELIgnored="false" %>
<script>
    $(function(){
        $("#guruEditForm").form('load','/guru/queryOne?id=${param.id}');
    });
</script>
<div style="text-align: center;">
    <form id="guruEditForm" method="post">
        <input class="easyui-textbox" name="id" type="hidden">
        <input class="easyui-textbox" name="src" type="hidden">
        <div style="margin-top: 2%;">
            名称:<input class="easyui-textbox" name="name" data-options="required:true"><br/>
        </div>
    </form>
</div>
