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
        $("#wisdomEditForm").form('load','/wisdom/queryOne?id=${param.id}');
    });
</script>
<div style="text-align: center;">
    <form id="wisdomEditForm" method="post">
        <input class="easyui-textbox" name="id" type="hidden">
        <input class="easyui-textbox" name="src" type="hidden">
        <input class="easyui-textbox" name="createDate" type="hidden">
        <div style="margin-top:40px;">
            标题:<input class="easyui-textbox" name="title" data-options="required:true"><br/>
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
                                onSelect: editeditorr
			        ">
        </div>
        <%--<div style="margin-top: 1%">
            内容
            <input name="content" class="easyui-textbox" data-options="required:true"/>
        </div>--%>
        <div>
            <div style="margin-top: 10px">内容</div>
            <div>
                <textarea id="edittxt" name="content"></textarea>
            </div>
        </div>
    </form>
</div>
<script>
    //创建kindeditor
    var editeditor = KindEditor.create('#edittxt',{
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

    function editeditorr(rec){
        editeditor.html($('#edittxt').val());
    }



</script>
