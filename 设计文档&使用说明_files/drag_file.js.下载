function sendFile(files) {
    if (!files || files.length < 1) {
        return;
    }
    var dropbox = document.getElementById('dropbox');
      
    var percent = document.createElement('div' );
    dropbox.appendChild(percent);

    var formData = new FormData();             // 创建一个表单对象FormData
    formData.append( 'submit', '中文' );  // 往表单对象添加文本字段
    
    var fileNames = '' ;
    
    for ( var i = 0; i < files.length; i++) {
        var file = files[i];    // file 对象有 name, size 属性
        
        if (file.name.indexOf(".md") == -1) continue;
        
        var reader = new FileReader();  
        reader.onload = function()  
        {
            dropbox.remove();
            markedWithToc(this.result);
        };
        reader.readAsText(file); 
    }
}

document.addEventListener("dragover", function(e) {
      e.stopPropagation();
      e.preventDefault();            // 必须调用。否则浏览器会进行默认处理，比如文本类型的文件直接打开，非文本的可能弹出一个下载文件框。
}, false);

document.addEventListener("drop", function(e) {
      e.stopPropagation();
      e.preventDefault();            // 必须调用。否则浏览器会进行默认处理，比如文本类型的文件直接打开，非文本的可能弹出一个下载文件框。

      sendFile(e.dataTransfer.files);
}, false);