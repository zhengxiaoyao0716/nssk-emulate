function setCatalogFold() {
    var catalogOl = document.getElementById("catalog");
    var contentDiv = document.getElementById("markdownContent");
    if (!catalogOl || !contentDiv) return false;
    
    catalogOl.onmouseover = function () { fold(20, 1); }
    catalogOl.onmouseout = function () { fold(1, -1); }
    
    var timer = null;
    function fold(target, speed) {
        clearInterval(timer);
        var width = parseInt(catalogOl.style.width.slice(0, -1));
        timer = setInterval(function (){
            if (width == target) {
                clearInterval(timer);
            }
            else {
                width += speed;         
                catalogOl.style.width = width + "%";
                contentDiv.style.width = 90 - width + "%";
            }
        }, 10);
    }
}
var oldOnload = window.onload;
window.onload = function () {
    if (oldOnload) oldOnload();
    
    document.body.style.maxWidth = document.documentElement.clientWidth + "px";
    document.body.style.width = document.documentElement.clientWidth - 60 + "px";
    document.body.style.paddingTop = "0px";
    document.body.style.marginTop = "10px";
    
    setCatalogFold();
};
function markedWithToc(content) {
    var catalogOl = document.createElement("ol");
    catalogOl.id = "catalog";
    catalogOl.style.position = "fixed";
    catalogOl.style.left = "0px";
    catalogOl.style.top = "0px";
    catalogOl.style.lineHeight = "30px";
    catalogOl.style.fontSize = "24px";
    catalogOl.style.border = "3px solid";
    catalogOl.style.padding = "30px";
    catalogOl.style.width = "1%";
    catalogOl.style.height = document.documentElement.clientHeight - 60 + "px";
    catalogOl.style.overflowY = "auto";
    document.body.appendChild(catalogOl);
    var contentDiv = document.createElement("div");
    contentDiv.id = "markdownContent";
    contentDiv.innerHTML = marked(content);
    contentDiv.style.cssFloat = "right";
    contentDiv.style.marginRight = "6%";
    contentDiv.style.width = "89%";
    contentDiv.style.height = document.documentElement.clientHeight - 60 + "px";
    contentDiv.style.overflowY = "auto";
    document.body.appendChild(contentDiv);
    
    var item = contentDiv.firstElementChild;
    var h1Count = 0;
    var h2Count = 0;
    var secondCatalogOl;
    while(true) {
        item = item.nextElementSibling;
        if (!item) break;
        
        if (item.tagName == 'H1') {
            h1Count++;
            h2Count = 0;
            var id = h1Count;
            
            var catalogA = document.createElement("a");
            catalogA.textContent = item.textContent;
            catalogA.href = '#' + id;
            secondCatalogOl = document.createElement("ol");
            var catalogLi = document.createElement("li");
            catalogLi.style.marginBottom = "16px";
            catalogLi.appendChild(catalogA);
            catalogLi.appendChild(secondCatalogOl);
            catalogOl.appendChild(catalogLi);
            
            item.innerHTML = '<a name = "' + id + '"></a>' + id + '. ' + item.textContent;
        }
        else if (item.tagName == 'H2') {
            if (!secondCatalogOl) continue;
            
            h2Count++;
            var id = h1Count + '.' + h2Count;
            
            var catalogA = document.createElement("a");
            catalogA.textContent = item.textContent;
            catalogA.href = '#' + id;
            var catalogLi = document.createElement("li");
            catalogLi.appendChild(catalogA);
            secondCatalogOl.appendChild(catalogLi);
            
            item.innerHTML = '<a name = "' + id + '"></a>' + id + ' ' + item.textContent;
        }
    };
    setCatalogFold();
};