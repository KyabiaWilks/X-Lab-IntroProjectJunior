$(document).ready(function() {
    var commentsPerPage = 10; // 每页显示的评论数
    var currentPage = 1; // 当前页码
    var totalComments = 0; // 总评论数
    var comments = []; // 用于存储所有评论
    var host = 'http://localhost:8080'; // 服务器地址
    var orderBool = 1;
    var orderStr = ["Forum Order", "Time Order"];
    var colorOptions = ['#78C2AD', '#1B998B', '#62A8AC', '#46B1C9', '#7B68EE', '#B49FCC', '#C19AB7', '#D0E1F9', '#79ADDC', '#40A798'];


    


    // 定时更新评论区内容
    function updateComments() {
        setInterval(function() {
            getComments();
        }, 30000);
    }

    // 获取评论
    function getComments() {
        var request = {
            page: currentPage,
            size: commentsPerPage
        };
        $.ajax({
            url: host + '/comment/get',
            method: 'GET',
            dataType: 'json',
            data: request,
            success: function(response) {
                //alert(response.code + response.msg);
                if (response.code === 0) {
                    comments = response.data.comments;
                    // currentPage = page,
                    // commentsPerPage = size,
                    totalComments = response.data.total;
                    renderComments();
                } else {
                    alert('Failed to get comments: ' + response.msg);
                }
            },
            error: function() {
                alert('Failed to get comments due to network error.');
            }
        });
    }

    // 渲染评论
    function renderComments() {
        $(".comment-section").empty(); 

        // var start = (currentPage - 1) * commentsPerPage; // 计算当前页的起始索引
        // var end = start + commentsPerPage; // 计算当前页的结束索引
        // var pageComments = comments.slice(start, end); // 获取当前页的评论

        comments.forEach(function(comment) {
            var commentElement = $('<div class="comment"></div>'); // 创建评论元素
            var userElement = $('<div class="comment-user"></div>').text(comment.name); // 创建用户名元素
            var textElement = $('<div class="comment-text"></div>').text(comment.content); // 创建评论内容元素
            var deleteButton = $('<button class="delete-button">Delete</button>'); // 创建删除按钮

             // 设置评论元素的边框颜色
            var color = getColorForUsername(comment.name);
            commentElement.css('border-left', '5px solid ' + color);

            // 绑定删除事件
            deleteButton.on('click', function() {
                //var commentIndex = start + index; // 计算要删除的评论在数组中的索引
                deleteComment(comment.id);
            });

            commentElement.append(userElement); // 将用户名元素添加到评论元素中
            commentElement.append(textElement); // 将评论内容元素添加到评论元素中
            commentElement.append(deleteButton); // 将删除按钮添加到评论元素中
            $(".comment-section").append(commentElement); // 将评论元素添加到评论展示区
        });

        updatePagination(); // 更新分页按钮状态
    }

    // 添加评论
    $('.submit-button').on('click', function() {
        var username = $(".username-input").val(); // 获取用户名输入框的值
        var commentText = $(".comment-input").val(); // 获取评论内容输入框的值
        
        // var notice = $('<div class="comment-section"></div>'); // 创建评论元素
        // var usernameMiss = $('<div class="missing-notice"></div>').text("Username missing.");
        // var contentMiss = $('<div class="missing-notice"></div>').text("Comment content missing.");
        // var allMiss = $('<div class="missing-notice"></div>').text("Multiple error.");

        if (username && commentText) {
            var comment = {
                Name: username,
                Content: commentText
            };
            $.ajax({
                url: host + '/comment/add',
                method: 'POST',
                contentType: 'application/json',
                data: JSON.stringify(comment),
                success: function(response) {
                    //alert(response.code);
                    if (response.code === 0) {
                        getComments();
                        $(".username-input").val("");
                        $(".comment-input").val("");
                    } 
                    // else {
                    //     alert('Failed to add comment: ' + response.msg);
                    // }
                },
                error: function() {
                    alert('Failed to add comment due to network error.');
                }
            });
        } else {
            alert('Username and comment content are required.');
        }
    });

    // 删除评论
    function deleteComment(id) {
        $.ajax({
            url: host + '/comment/delete',
            method: 'POST',
            contentType: 'application/json',
            dataType: 'json',
            data: JSON.stringify({ id: id }),
            success: function(response) {
                if (response.code === 0) {
                    getComments();
                    // renderComments(); // 重新渲染评论
                } else {
                    alert('Failed to delete comment: ' + response.msg);
                }
            }
        });
    }

    // 更新分页按钮状态
    function updatePagination() {
        var totalPages = Math.ceil(totalComments / commentsPerPage); // 计算总页数
        if (totalPages <= 0) totalPages = 1;
        $(".page-num").text("Page " + currentPage); // 更新页码显示
        $(".prev-button").prop("disabled", currentPage === 1); // 如果是第一页，禁用“上一页”按钮
        $(".next-button").prop("disabled", currentPage === totalPages); // 如果是最后一页，禁用“下一页”按钮
    }

    // 绑定分页按钮事件
    $(".prev-button").on('click', function() {
        if (currentPage > 1) {
            currentPage--; // 当前页码减1
            getComments(); // 重新渲染评论
        }
    });

    $(".next-button").on('click', function() {
        var totalPages = Math.ceil(totalComments / commentsPerPage); // 计算总页数
        if (currentPage < totalPages) {
            currentPage++; // 当前页码加1
            getComments(); // 重新渲染评论
        }
    });

    $(".order-button").on('click', function() {
        $.ajax({
            url: host + '/order/switch',
            method: 'GET',
            contentType: 'application/json',
            dataType: 'json',
            
            success: function(response) {
                if (response.code === 0) {
                    orderBool = orderBool===0?1:0;
                    getComments();
                    document.querySelector('.order-button').textContent = orderStr[orderBool];
                    
                    // renderComments(); // 重新渲染评论
                } else {
                    alert('Failed to delete comment: ' + response.msg);
                }
            }
        });
    });

    function hashString(str) {
        var hash = 0;
        for (var i = 0; i < str.length; i++) {
            var char = str.charCodeAt(i);
            hash = ((hash << 5) - hash) + char;
            hash = hash & hash; // Convert to 32bit integer
        }
        return Math.abs(hash);
    }
    
    function getColorForUsername(username) {
        var hash = hashString(username);
        var colorIndex = hash % colorOptions.length;
        return colorOptions[colorIndex];
    }
    

    // 初始渲染评论
    getComments();
    updateComments();
});
