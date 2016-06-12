function getFormJson(frm) {
    var o={};
    var a=$(frm).serializeArray()
    $.each(a,function () {
        o[this.name]=this.value || '';
    })
    return o;
}
function ajaxSubmit(frm,fn) {
    var dat=getFormJson(frm)
     return $.ajax({
        url  : frm.action,
        type : frm.method,
        data : dat,
        success:fn
    })
}
function  markDowmToThml(markDown) {
    var text=$.trim(markDown);
    var converter = new showdown.Converter();
    //支持显示如同github的勾选框，默认false
    //ep: - [x] This task is done
    converter.setOption("tasklists", true);
    //支持显示table，默认false
    converter.setOption("tables", true);
    //支持图片大小设置，默认为false
    converter.setOption("parseImgDimensions", true);
    /**
     **更多地请看https://github.com/showdownjs/showdown文档
     **/
    var html = converter.makeHtml(text);
    return html;
}