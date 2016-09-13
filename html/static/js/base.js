// base
// @author: zhengxiaoyao0716
; (function () {
    var base = {
        "config": {
            "baseUrl": "",
            "urlMap": {},
            "simulatAjax": window.location.host == "",
            "expectRespMap": {}
        }
    };
    base.dialog = {
        "waiting": {
            "show": function () {console.log("waiting");},
            "hide": function () {}
        },
        "toast": {
            "show": function (title, message, duration) {
                console.log("toast?title=" + title + "&message=" + JSON.stringify(message) + "&duration=" + duration);
            }
        }
    };

    base.query = function (key) {
        var reg = new RegExp("(^|&)" + key + "=([^&]*)(&|$)");
        var value = window.location.search.substr(1).match(reg);
        return value && unescape(value[2]);
    };

    function getUrl(url) {
        return base.config.baseUrl + (base.config.urlMap[url] || url);
    }
    /** Ajax方法 */
    base.ajax = function(url, type, contentType, data, success, silence, extend) {
        silence || base.dialog.waiting.show();
        var config = {
            "url": getUrl(url),
            "type": type,
            "contentType": contentType,
            "dataType": "json",
            "data": data,
            "success": function(data) {
                success && success(data) || silence || base.dialog.toast.show(undefined, undefined, 1000);
            },
            "error": function(resp) {
                var reas;
                // var then;
                try {
                    respJson = JSON.parse(resp.responseText);
                    reas = respJson["reas"];
                    // then = respJson["then"];
                } catch (e) {
                    reas = resp.statusText;
                }
                base.dialog.toast.show("出错了", reas);
            },
            "complete": base.dialog.waiting.hide
        };
        $.extend(config, extend)
        $.ajax(config);
    };
    // 模拟Ajax调试
    if (base.config.simulatAjax) {
        base.ajax = function(url, type, contentType, data, success, silence, extend) {
            data = contentType == "application/json" ? JSON.parse(data) : data;
            console.group("模拟Ajax请求:")
            console.group("=request=")
            console.debug("url: " + getUrl(url));
            console.debug("type: " + type);
            console.debug("contentType: " + contentType);
            console.debug("data: " + data);
            console.debug("success: " + success);
            console.debug("silence: " + silence);
            console.debug("extend: " + extend);
            console.groupEnd();

            silence || base.dialog.waiting.show();
            var resp = base.config.expectRespMap[url] || {};
            if (typeof resp == "function") {
                data = resp(data) || {};
            }
            if (resp.flag === undefined) {
                data = {
                    "flag": true,
                    "body": data,
                    "reas": undefined,
                    "then": undefined
                };
            }
            var time = Math.random() * 3000;
            setTimeout(function() {
                success && success(data) || silence || base.dialog.toast.show(undefined, undefined, 1000);
                base.dialog.waiting.hide();
            }, time);

            console.group("=response=");
            console.debug("data: " + JSON.stringify(data));
            console.groupEnd();
            console.log("随机耗时: " + time);
            console.groupEnd();
        };
    }
    base.get = function(url, data, func, silence) {
        base.ajax(url, "GET", "application/x-www-form-urlencoded", data, func, silence);
    };
    base.post = function(url, data, func, silence) {
        base.ajax(url, "POST", "application/json", JSON.stringify(data), func, silence);
    };

    // Module defined.
    if (typeof define === 'function' && define.amd) {
        define(function () {
            return base;
        });
    } else if (typeof module !== 'undefined' && module.exports) {
        module.exports = base;
    } else {
        window.base = base;
    }
})();
