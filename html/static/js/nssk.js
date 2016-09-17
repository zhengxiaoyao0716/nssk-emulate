(function () {
    // 卡片按钮
    var triggers = {};
    var $dialogs = $(".dialog");
    $(".card").on("click", function () {
        var name = $(this).attr("name");
        $dialogs.hide();
        triggers[name] && triggers[name]();
        $("#" + name).fadeIn(300);
    });

    // 通讯录
    var $connectGrids = $("#connect").find(".grids");
    triggers.connect = function () {
        base.get(master + "/api/s/user/list", {}, function (data) {
            refreshConnects(undefined, data["body"]);
            return true;
        }, true);
    };
    function refreshConnects(connects, users) {
        if (connects) {
            $connectGrids.connects = connects;
        } else {
            connects = $connectGrids.connects;
        }
        if (users) {
            $connectGrids.users = users;
        } else {
            users = $connectGrids.users;
        }
        $connectGrids.empty();
        $(users).each(function (index, user) {
            var $user = $("<p>" + user + "</p>");
            var $state;
            switch (connects[user]) {
                case 0:
                    $state = $("<span>[已连接] </span>");
                    break;
                case 1:
                    $state = $("<span>[请求中] </span>");
                    $state.on("click", function () {
                        base.post("./api/a/connect/create", { "address": user });
                    });
                    break;
                case 2:
                    var $state = $('<a href="javascript:;" title="接受连接请求"><i class="fa fa-check-square-o" aria-hidden="true"></i> </a>');
                    $state.on("click", function () {
                        base.post("./api/b/connect/verify", { "address": user });
                    });
                    break;
                case undefined:
                    var $state = $('<a href="javascript:;" title="发送连接请求"><i class="fa fa-user-plus" aria-hidden="true"></i> </a>');
                    $state.on("click", function () {
                        base.post("./api/a/connect/create", { "address": user });
                    });
                    break;
                default:
                    break;
            }
            $user.prepend($state);
            $connectGrids.append($user);
        });
    }

    // 消息
    var $messageGrids = $("#message").find(".grids");
    $messageGrids.cache = {};
    function refreshMessages(messageMap) {
        // $messageGrids.empty();
        for (var user in messageMap) {
            var size = messageMap[user].length;
            var cachedUser = $messageGrids.cache[user];
            var $messages;
            if (cachedUser) {
                if (size == cachedUser["size"]) {
                    continue;
                } else {
                    $messages = cachedUser["$messages"];
                    $messages.empty();
                }
            } else {
                var $user = $('<div class="brief"></div>');
                $user.append("<p>" + user + "</p>");
                $messages = $('<ul class="detail"></ul>');
                $user.append($messages);
                $messageGrids.append($user);
            }
            $(messageMap[user]).each(function (index, message) {
                $messages.append("<ol>" + message + "</ol>");
            });
            $messageGrids.cache[user] = {"$messages": $messages, "size": size};
        }
    }

    // 发信
    (function () {
        var $sendDialog = $('#send');
        $sendDialog.close = function () {$sendDialog.hide();};
        var $close = $sendDialog.find("#close");
        $close.on("click", $sendDialog.close);
        var $input = $sendDialog.find("input");
        var $ul = $sendDialog.find("ul");
        $ul.open = function () {
            $ul.filter();
            $ul.fadeIn();
        };
        $ul.close = function () {
            $ul.fadeOut();
        };
        $ul.filter = function () {
            $ul.empty();
            var filter = $input.val();
            for (var user in $connectGrids.connects) {
                if ($connectGrids.connects[user] == 0 && user.indexOf(filter) != -1) {
                    $user = $("<ol>" + user + "</ol>");
                    $user.on("click", (function (user) {
                        return function () {$input.val(user);}
                    })(user));
                    $ul.prepend($user);
                }
            }
        }
        $input.on("focus", $ul.open);
        $input.on("blur", $ul.close);
        $input.on("input", $ul.filter);
        var $textarea = $sendDialog.find("textarea");
        var $submit = $sendDialog.find("#submit");
        $submit.on("click", function () {
            var user = $input.val();
            if (user == "") {
                base.dialog.toast.show("警告", "地址不可为空")
                return;
            }
            var message = $textarea.val();
            if (message == "") {
                base.dialog.toast.show("警告", "正文不可为空")
                return;
            }
            base.post("./api/message/send", { "address": user, "message": message}, $sendDialog.close);
        });
    })();

    // 用户列表
    var $userGrids = $("#users").find(".grids");
    triggers.users = function () {
        base.get("./api/s/user/list", {}, function (data) {
            $userGrids.empty();
            $(data["body"]).each(function (index, user) {
                $userGrids.append("<p>" + user + "</p>");
            });
            return true;
        }, true);
    };

    // 获取应用
    triggers.download = function () {
        open("./resource/app/download");
    }

    // 日志
    var $consoleList = $("#console").find("ul");
    function refreshLogs(logs) {
        $consoleList.empty();
        $(logs).each(function (index, log) {
            $consoleList.append("<ol>" + log + "</ol>");
        });
    }

    // 刷新数据
    function refreshData() {
        base.get("./api/all/pull", {}, function (data) {
            data = data["body"];
            refreshLogs(data["logs"]);
            if (!window.isMaster) {
                refreshConnects(data["connects"]);
                refreshMessages(data["messages"]);
            }
            return true;
        }, true);
        setTimeout(refreshData, 1000);
    }
    $(document).ready(function () {
        refreshData();
        $(".titlebar").find("i").on("click", function () {
            $dialogs.hide();
            if (window.isMaster) {
                $(".if-master").hide();
                $(".ifn-master").show();
                window.isMaster = false;
            } else {
                $(".if-master").show();
                $(".ifn-master").hide();
                window.isMaster = true;
            }
        });
    });

    // 配置部署服务器
    window.nssk = {};
    nssk.bindServer = function (address) {base.post("./api/server/bind", {"address": address});};
})();