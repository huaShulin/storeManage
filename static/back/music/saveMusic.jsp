<%--
  Created by IntelliJ IDEA.
  User: Shulin
  Date: 14/11/2017
  Time: 下午 5:19
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<div style="text-align: center; margin-top: 3%">
    <form id="musicSaveForm" method="post" enctype="multipart/form-data">
        <div>
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
            评分
            <%--<input name="score" class="easyui-textbox" data-options="required:true"/>--%>
            <input name="score" type="text" class="easyui-numberbox" value="5" data-options="min:0,precision:0,max:5"></input>
        </div>
        <div style="margin-top: 1%">
            选择封面
            <%--<input name="multipartFile" type="file"/>--%>
            <input class="easyui-filebox" name="multipartFile" style="width: 23%">
        </div>
        <div style="margin-top: 1%">
            描述
            <input name="content" class="easyui-textbox" data-options="required:true"/>
        </div>

    </form>
</div>
<script>
</script>