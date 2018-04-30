<%--
  Created by IntelliJ IDEA.
  User: Shulin
  Date: 14/11/2017
  Time: 下午 5:19
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<div style="text-align: center; margin-top: 3%">
    <form id="userSaveForm" method="post" enctype="multipart/form-data">
        <div>
            用户名
            <input name="title" class="easyui-textbox" data-options="required:true"/>
        </div>
        <div style="margin-top: 1%">
            类型
            <select class="easyui-combobox" name="性别" style="width: 15%;">
                <option value="男">男</option>
                <option value="女">女</option>
            </select>
        </div>
        <div style="margin-top: 1%">
            上师
            <input id="guru"
                   name="guruId"
                   class="easyui-combobox"
                   style="width: 15%"
                   data-options="
                                url:'/guru/queryOptions',
                                method:'get',
                                valueField:'id',
                                textField:'name',
                                panelHeight:'auto',
                                onSelect: function(rec){
                                    }
			        ">
        </div>
        <div style="margin-top: 1%">
            选择头像
            <%--<input name="multipartFile" type="file"/>--%>
            <input class="easyui-filebox" name="multipartFile" style="width: 23%">
        </div>
        <div style="margin-top: 1%">
            签名
            <input name="content" class="easyui-textbox" data-options="required:true"/>
        </div>

    </form>
</div>
<script>
</script>