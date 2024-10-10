layui.define(['jquery', 'element', 'form'],
    function (exports) {

        var form = layui.form,
            $ = layui.jquery;
        var formInput = new function () {

            /**
             * @since form
             * 
             * @param config 配置项
             */
            this.render = config => {
                var f = $(`<form class="layui-form" style="margin-top: 9px"></form>`)
                var d = $(`<div class="layui-row layui-col-space8"> </div>`)
                config.data.forEach((e,i)=>{
                    var input = ''
                    if (!e.colClass){e.colClass='layui-col-md1'}
                    if (e.type === 'select'){
                        input = $(`<div class="${e.colClass}"></div>`)
                        /*
                        在 <select> 元素上设置 lay-search 可开启选择框的搜索功能，如：lay-search="{caseSensitive:false, fuzzy: false}"，支持的可选项如下：
                            caseSensitive：是否区分大小写，默认值 false
                            fuzzy：是否开启模糊匹配，设置 true 开启后将忽略匹配字符出现在字符串中的位置。默认值 false
                        在 <select> 元素上设置 lay-creatable="" 可允许创建新的 option，需开启 lay-search 后生效。
                        */
                        var select = $(` <select name="${e.name}" lay-search="{caseSensitive:false, fuzzy: false}" lay-creatable=""><option value="">请选择</option></select>`)
                        if (e.children){
                            e.children.forEach((item,index)=>{
                                var option = new Option(item.text, item.value)
                                select.append(option)
                            })
                            input.append(select)
                        }
                    } else if (e.type==='button') {
                        input = $(`<div class="layui-btn-container ${e.colClass}"></div>`)
                        if (e.layFilter){
                            var btn = $(`<button type="${e.buttonType}" class="${e.class}" lay-submit lay-filter="${e.layFilter}">${e.placeholder}</button>`)
                        }

                        var btn = $(`<button type="${e.buttonType}" class="${e.class}" >${e.placeholder}</button>`)
                        if (e.icon){
                            btn.prepend($(`<i class="layui-icon ${e.icon}"></i>`))
                        }
                        input.append(btn)
                    } else {
                        if (e.icon) {
                            input = $(`
                    <div class="${e.colClass}">
                        <div class="layui-input-wrap">
                            <div class="layui-input-prefix">
                                <i class="layui-icon ${e.icon}"></i>
                            </div>
                            <input type="${e.type}" name="${e.name}" placeholder="${e.placeholder}" lay-affix="${e.layAffix}" class="${e.class}" value="${e.value}">
                        </div>
                    </div>
                    `)
                        }else{
                            input = $(`
                    <div class="${e.colClass}">
                        <div class="layui-input-wrap">
                            <input type="${e.type}" name="${e.name}" placeholder="${e.placeholder}" lay-affix="${e.layAffix}" class="${e.class}" value="${e.value}">
                        </div>
                    </div>
                    `)

                        }}
                    d.append(input)

                })
                f.append(d)
                $(config.elem).append(f)
                form.render()
            }

        }



        exports('formInput', formInput);
    })
