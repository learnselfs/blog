layui.define(['jquery', 'element'], function(exports) {
	"use strict";

	var MOD_NAME = 'context',
		$ = layui.jquery,
		element = layui.element;

	var context = new function() {

		this.get = (key)=>{
			var result = localStorage.getItem(key);
			if (JSON.isRawJSON(result)){}
		}

		this.set = (key, obj) =>{
			if (obj instanceof Object) {
				const o = JSON.stringify(obj);
				localStorage.setItem(key,o);
			}else{
				localStorage.setItem(key,obj);
			}
		}

		this.getObject = (key) => {
			return JSON.parse(localStorage.getItem(key))
		}
	}
	exports(MOD_NAME, context);
});
