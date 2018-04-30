<%--
  Created by IntelliJ IDEA.
  User: Shulin
  Date: 14/11/2017
  Time: 下午 5:19
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<div style="text-align: center; margin-top: 3%">
    <form id="chapterSaveForm" method="post" enctype="multipart/form-data">
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
        <div style="margin-top: 1%">
            选择文件
            <%--<input name="multipartFile" type="file"/>--%>
            <input class="easyui-filebox" name="multipartFile" style="width: 23%">
        </div>

    </form>
</div>
<script>
</script>