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

            /**
             * @since ajax 请求
             *
             * @param url 请求地址
             * @param type 请求方式
             * @param data 请求数据
             * @param async_ 请求异步
             */
            this.ajax = (url,type, data, asy)=>{
                return new Promise((resolve) => {
                    $.ajax({
                        type: type?type:"get",
                        url: url,
                        dataType: 'json',
                        data: JSON.stringify(data),
                        async: asy,
                        contentType: "application/json",
                        success: (res) => {
                            if (res.code === 0|| res.openapi) {
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


            /**
             * @since 解析jwt
             *
             * @param token jwt token
             */
            this.jwtDecode = token =>{
                var t = token.split('.')
                var claim = t[1]
                var claimLength = claim.length%4
                if (t.length===3) {
                    claim.replace(/-/g, '+').replace(/_/g,'/')
                    if (claimLength){
                        claim += '='.repeat(4-claimLength)
                    }
                 return  window.atob(claim)
                   // return  window.atob(claim.replace(/-/g, '+').replace(/_/g,'/'))
                }
            }

        }



        exports('tools', tools);
    })
