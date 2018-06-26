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
    var form = textarea.closest('form')
    var validate = debounce(function(src){
        $.ajax({
            method: 'post', url: '/api/define',
            data: {"definition": src}
        }).then(function(response){
            if (response.ok) {
                $('[data-definition-valid="true"]').show()
                $('[data-definition-valid="false"]').hide()
                $('[data-definition-invalid]').prop('disabled', false)
                $('[data-definition-valid-error]').text('')
                console.info(response)
                return
            }
            // clean up the error message a bit, the filename for example isn't
            // actually useful.
            var pos = response.error.indexOf('io.Reader.go:')
            if (pos > -1) {
                response.error = response.error.substr(pos+('io.Reader.go:'.length))
            }
            response.error = response.error.replace(/^io.Reader.go:/, '')
            pos = response.error.indexOf(':')
            var line = parseInt(response.error.substr(0, pos))
            pos = response.error.indexOf(':', pos+1)
            var col = parseInt(response.error.substr(0, pos))
            $('[data-definition-valid="true"]').hide()
            $('[data-definition-valid="false"]').show()
            $('[data-definition-invalid]').prop('disabled', true)
            $('[data-definition-valid-error]')
                .text(response.error || "An unknown error occurred")
                .attr('href', '#')
                .click(function(e){
                    e.preventDefault()
                    editor.gotoLine(line)
                })
        }).catch(function(error){
            $('[data-definition-valid="true"]').hide()
            $('[data-definition-valid="false"]').show()
            console.info(error)
            $('[data-definition-valid-error]').text(error.responseText || "An unknown error occurred").attr('href', '')
        })
    }, 500)
    editor.getSession().on("change", function () {
        var src = editor.getSession().getValue()
        textarea.val(src)
        $('[data-definition-valid="true"]').hide()
        $('[data-definition-valid="false"]').hide()
        validate(src)
    });

    // validate whatever's rendered into the editor
    var src = editor.getSession().getValue()
    textarea.val(src)
    $('[data-definition-valid="true"]').hide()
    $('[data-definition-valid="false"]').hide()
    validate(src)

    // Returns a function, that, as long as it continues to be invoked, will not
    // be triggered. The function will be called after it stops being called for
    // N milliseconds. If `immediate` is passed, trigger the function on the
    // leading edge, instead of the trailing.
    function debounce(func, wait, immediate) {
        var timeout;
        return function() {
            var context = this, args = arguments;
            var later = function() {
                timeout = null;
                if (!immediate) func.apply(context, args);
            };
            var callNow = immediate && !timeout;
            clearTimeout(timeout);
            timeout = setTimeout(later, wait);
            if (callNow) func.apply(context, args);
        };
    }; // from https://davidwalsh.name/javascript-debounce-function 

})
