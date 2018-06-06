(function($){

    // $.remoto is the default namespace for Remoto services.
    $.remoto = $.remoto || {}
    
    // new $.Facebox creates a jQuery client.
    // Facebox provides facial detection and recognition capabilities.
    // Usage:
    //     var remoteFacebox = new $.Facebox({endpoint: "https://example.machinebox.io/remoto"})
    $.remoto.Facebox = function(options) {
        options = options || {}
        this.endpoint = options.endpoint || "http://localhost:8080"
    }
    
    // Usage:
    //     remoteFacebox.Check([{
    //         image: null,
    //     }])
    //       .done(function(responses) { /* called on success */ })
    //       .fail(function(){ /* called on failure */ })
    //       .always(function(){ /* always gets called */ })
    $.remoto.Facebox.prototype.Check = function(requests) {
        if (!Array.isArray(requests)) {
            throw 'Facebox.Check: first argument must be Array'
        }
        return $.ajax({
            method: 'post', url: this.endpoint + '/remoto/Facebox.Check',
            data: JSON.stringify(requests),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
    }
    
    // Usage:
    //     remoteFacebox.Teach([{
    //         name: null,
    //         teachFiles: null,
    //     }])
    //       .done(function(responses) { /* called on success */ })
    //       .fail(function(){ /* called on failure */ })
    //       .always(function(){ /* always gets called */ })
    $.remoto.Facebox.prototype.Teach = function(requests) {
        if (!Array.isArray(requests)) {
            throw 'Facebox.Teach: first argument must be Array'
        }
        return $.ajax({
            method: 'post', url: this.endpoint + '/remoto/Facebox.Teach',
            data: JSON.stringify(requests),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
    }
    
})(jQuery)
