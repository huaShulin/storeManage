<%--
  Created by IntelliJ IDEA.
  User: Shulin
  Date: 14/11/2017
  Time: 下午 3:58
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" pageEncoding="UTF-8" language="java" %>

<!-- 为ECharts准备一个具备大小（宽高）的Dom -->
<div id="main" style="width: 600px;height:400px;"></div>
<script type="text/javascript">

    var myChart = echarts.init(document.getElementById('main'));
    // 显示标题，图例和空的坐标轴
    myChart.setOption({
        title: {
            text: '异步数据加载示例'
        },
        tooltip: {},
        legend: {
            data:['文件数']
        },
        color:[],
        xAxis: {
            data: []
        },
        yAxis: {},
        series: [{
            name: '文件数',
            type: 'bar',
            data: [],
            itemStyle: {
                normal:{
                    color: function (params){
                        var colorList = ['rgb(164,205,238)','rgb(42,170,227)','rgb(25,46,94)'];
                        return colorList[params.dataIndex]}
                }
            }
        }]
    });

    // 异步加载数据
    /*$(function(){
        var jsonCategories = eval("("+ $("#jsonCategories").val() +")");
        var jsonData = eval("("+ $("#jsonData").val() +")");

        myChart.setOption({
            xAxis: {
                data: jsonCategories
            },
            series: [{
                // 根据名字对应到相应的系列
                name: '文件数',
                data: jsonData
            }],
            //color: ['red', 'green','yellow']
        });
    });*/
     $.get('/banner/files').done(function (data) {
        // 填入数据
        myChart.setOption({
            xAxis: {
                data: data.categories
            },
            series: [{
                // 根据名字对应到相应的系列
                name: '销量',
                data: data.data
            }]
        });
    });
    // 使用刚指定的配置项和数据显示图表。
    //myChart.setOption(option);
</script>