// base
// @author: zhengxiaoyao0716
; (function () {
    var base = {};

    base.query = function (key) {
        var reg = new RegExp("(^|&)" + key + "=([^&]*)(&|$)");
        var value = window.location.search.substr(1).match(reg);
        return value && unescape(value[2]);
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
