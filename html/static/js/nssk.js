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
                    break;
                case 2:
                    var $state = $('<a href="javascript:;" title="接受连接请求"><i class="fa fa-check-square-o" aria-hidden="true"></i> </a>');
                    $state.on("click", function () {
                        base.post("/api/b/connect/verify", { "address": user });
                    });
                    break;
                case undefined:
                    var $state = $('<a href="javascript:;" title="发送连接请求"><i class="fa fa-user-plus" aria-hidden="true"></i> </a>');
                    $state.on("click", function () {
                        base.post("/api/a/connect/create", { "address": user });
                    });
                    break;
                default:
                    break;
            }
            $user.prepend($state);
            $connectGrids.append($user);
        });
    }

    // 消息中心
    var $sendList = $("#send").find("ul");
    triggers.send = function () {
    }

    // 用户列表
    var $userList = $("#users").find("ul");
    triggers.users = function () {
        base.get(master + "/api/s/user/list", {}, function (data) {
            $userList.empty();
            $(data["body"]).each(function (index, user) {
                $userList.append("<ol>" + user + "</ol>");
            });
            return true;
        }, true);
    };

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
        base.get("/api/all/pull", {}, function (data) {
            data = data["body"];
            refreshLogs(data["logs"]);
            refreshConnects(data["connects"]);
            console.log(data)
            return true;
        }, true);
        setTimeout(refreshData, 1000);
    }
    $(document).ready(function () {
        refreshData();
    });


})();