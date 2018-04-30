<%--
  Created by IntelliJ IDEA.
  User: Shulin
  Date: 14/11/2017
  Time: 下午 5:19
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<div style="text-align: center; margin-top: 3%">
    <form id="bannerSaveForm" method="post" enctype="multipart/form-data">
        <div>
            名称
            <input name="name" class="easyui-textbox" data-options="required:true"/>
        </div>
        <div style="margin-top: 1%">
            状态
            <select class="easyui-combobox" name="status" style="width: 15%;">
                <option value="on">激活</option>
                <option value="off">未激活</option>
            </select>
        </div>
        <div style="margin-top: 1%">
            选择文件
            <%--<input name="multipartFile" type="file"/>--%>
            <input class="easyui-filebox" name="multipartFile" style="width: 23%">
        </div>

    </form>
</div>
<script>
</script>