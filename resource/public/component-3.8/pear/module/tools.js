layui.define(['jquery', 'element'],
    function (exports) {

        var $ = layui.jquery;
        var tools = new function () {

            /**
             * @since 防抖算法 
             * 
             * @param fn 要执行的方法
             * @param time 防抖时间参数
             */
            this.debounce = function (fn, time) {
                var timer = null
                return function () {
                    var arguments = arguments[0]
                    if (timer) {
                        clearTimeout(timer)
                    }
                    timer = setTimeout(function () {
                        fn(arguments)
                    }, time)
                }
            }

            // image 转 base64
            this.imageToBase64 = function (img) {
                var canvas = document.createElement("canvas");
                canvas.width = img.width;
                canvas.height = img.height;
                var ctx = canvas.getContext("2d");
                ctx.drawImage(img, 0, 0, img.width, img.height);
                var ext = img.src.substring(img.src.lastIndexOf(".") + 1).toLowerCase();
                var dataURL = canvas.toDataURL("image/" + ext);
                return dataURL;
            }

            this.ajax = (url,type, data, async_)=>{
                return new Promise((resolve) => {
                    $.ajax({
                        type: type?type:"get",
                        url: url,
                        dataType: 'json',
                        data: data,
                        async: async_,
                        success: (res) => {
                            if (res.code === 0) {
                                resolve(res)
                            } else {
                                layer.msg('获取失败：' + res.message, {icon: 2});
                            }
                        },
                        error: (res) => {
                            layer.msg('网络错误：' + res, {icon: 2});

                        }
                    })
                })
            }




        }



        exports('tools', tools);
    })
