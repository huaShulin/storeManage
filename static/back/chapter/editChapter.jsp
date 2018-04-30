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
        $("#chapterEditForm").form('load','/chapter/queryOne?id=${param.id}');
    });
</script>
<div style="text-align: center;">
    <form id="chapterEditForm" method="post">
        <input class="easyui-textbox" name="id" type="hidden">
        <input class="easyui-textbox" name="src" type="hidden">
        <input class="easyui-textbox" name="size" type="hidden">
        <div>
            name
            <input name="name" class="easyui-textbox" data-options="required:true"/>
        </div>
        <div style="margin-top: 1%">
            主题
            <input name="musicId"
                   class="easyui-combobox"
                   style="width: 15%"
                   data-options="
                                url:'/music/queryOptions',
                                method:'get',
                                valueField:'id',
                                textField:'name',
                                panelHeight:'auto',
                                onSelect: function(rec){
                                    }
			        ">
        </div>
    </form>
</div>
