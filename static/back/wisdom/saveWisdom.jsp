<%--
  Created by IntelliJ IDEA.
  User: Shulin
  Date: 14/11/2017
  Time: 下午 5:19
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<div style="text-align: center; margin-top: 3%">
    <form id="wisdomSaveForm" method="post" enctype="multipart/form-data">
        <div>
            标题
            <input name="title" class="easyui-textbox" data-options="required:true"/>
        </div>
        <div style="margin-top: 1%">
            类型
            <select class="easyui-combobox" name="type" style="width: 15%;">
                <option value="上师言教">上师言教</option>
                <option value="显密法要">显密法要</option>
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
            选择文件
            <%--<input name="multipartFile" type="file"/>--%>
            <input class="easyui-filebox" name="multipartFile" style="width: 23%">
        </div>
        <%--<div style="margin-top: 1%">
            内容
            <input name="content" class="easyui-textbox" data-options="required:true"/>
        </div>--%>
        <div>
            <div style="margin-top: 10px">内容</div>
            <div>
                <textarea id="savetxt" name="content"></textarea>
            </div>
        </div>

    </form>
</div>
<script>
    //创建kindeditor
    var saveeditor = KindEditor.create('#savetxt',{
        width:670,
        height:200,
        items:[
            'source', 'undo', 'redo', '|', 'preview', 'print', 'template', 'code', 'cut', 'copy', 'paste',
            'plainpaste', 'wordpaste', '|', 'justifyleft', 'justifycenter', 'justifyright',
            'justifyfull', 'insertorderedlist', 'insertunorderedlist', 'indent', 'outdent', 'subscript',
            'superscript', 'clearhtml', 'quickformat', 'selectall', '|', 'fullscreen',
            'formatblock', 'fontname', 'fontsize', '|', 'forecolor', 'hilitecolor', 'bold',
            'italic', 'underline', 'strikethrough', 'lineheight', 'removeformat', '|', 'image', 'multiimage',
            'flash', 'media', 'insertfile', 'table', 'hr', 'emoticons', 'baidumap', 'pagebreak',
            'anchor', 'link', 'unlink', '|', 'about'
        ],
        uploadJson:'${pageContext.request.contextPath}/wisdom/upload',
    });
</script>