function save() {
    var webpage = '<!DOCTYPE html> <html><head>' + document.head.innerHTML.replace(/\.\/static/g, "http://zhengxiaoyao0716.github.io/MarkedWithToc/static") + '</head><body>' + document.body.innerHTML.replace(/\.\/index/g, "http://zhengxiaoyao0716.github.io/MarkedWithToc/index") +'</body></html>';
    var blob = new Blob([webpage], {type: "text/html;charset=utf-8"});
    var aLink = document.createElement('a');
    var evt = document.createEvent("HTMLEvents");
    evt.initEvent("click", false, false);//initEvent 不加后两个参数在FF下会报错, 感谢 Barret Lee 的反馈
    aLink.download = "index.html";
    aLink.href = URL.createObjectURL(blob);
    aLink.dispatchEvent(evt);
}