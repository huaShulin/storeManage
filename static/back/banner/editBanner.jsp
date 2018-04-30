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
        $("#bannerEditForm").form('load','/banner/queryOne?id=${param.id}');
    });
</script>
<div style="text-align: center;">
    <form id="bannerEditForm" method="post">
        <div style="margin-top:40px;">
            <input class="easyui-textbox" name="id" type="hidden">
            <input class="easyui-textbox" name="src" type="hidden">
            名称:<input class="easyui-textbox" name="name" data-options="required:true"><br/>
        </div>
        <div style="margin-top:20px;">
            状态:<input class="easyui-textbox" name="status" data-options="required:true"><br/>
        </div>
    </form>
</div>
