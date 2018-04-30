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
        $("#musicEditForm").form('load','/music/queryOne?id=${param.id}');
    });
</script>
<div style="text-align: center;">
    <form id="musicEditForm" method="post">
        <input class="easyui-textbox" name="id" type="hidden">
        <input class="easyui-textbox" name="src" type="hidden">
        <input class="easyui-textbox" name="createDate" type="hidden">
        <div style="margin-top: 2%;">
            名称
            <input name="name" class="easyui-textbox" data-options="required:true"/>
        </div>
        <div style="margin-top: 1%">
            作者
            <input name="writer" class="easyui-textbox" data-options="required:true"/>
        </div>
        <div style="margin-top: 1%">
            播音
            <input name="announcer" class="easyui-textbox" data-options="required:true"/>
        </div>
        <div style="margin-top: 1%">
            集数
            <input name="sum" type="text" class="easyui-numberbox" data-options="min:0,precision:0"></input>
        </div>
        <div style="margin-top: 1%">
            评分
            <input name="score" type="text" class="easyui-numberbox" data-options="min:0,precision:0,max:5"></input>
        </div>
        <div style="margin-top: 1%">
            描述
            <input name="content" class="easyui-textbox" data-options="required:true"/>
        </div>
    </form>
</div>
