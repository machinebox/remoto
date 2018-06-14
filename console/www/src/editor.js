import $ from 'jquery'
require('../node_modules/ace-builds/src-noconflict/ace.js')

$(function(){

    var textarea = $('textarea[data-editor]')
    var editor = ace.edit("editor")
    ace.config.set("modePath", "/static")
    ace.config.set("workerPath", "/static") 
    ace.config.set("themePath", "/static") 
    editor.setTheme("ace/theme/monokai")
    editor.session.setMode("ace/mode/golang")
    editor.getSession().on("change", function () {
        textarea.val(editor.getSession().getValue());
    });

})
