(function($){

    // $.remoto is the default namespace for Remoto services.
    $.remoto = $.remoto || {}
    
    // new $.GreetFormatter creates a jQuery client.
    // GreetFormatter provides formattable greeting services.
    // Usage:
    //     var remoteGreetFormatter = new $.GreetFormatter({endpoint: "https://example.machinebox.io/remoto"})
    $.remoto.GreetFormatter = function(options) {
        options = options || {}
        this.endpoint = options.endpoint || "http://localhost:8080"
    }
    
    // 
    // Usage:
    //     remoteGreetFormatter.Greet([{
    //         format: null,
    //         name: null,
    //     }])
    //       .done(function(responses) { /* called on success */ })
    //       .fail(function(){ /* called on failure */ })
    //       .always(function(){ /* always gets called */ })
    $.remoto.GreetFormatter.prototype.Greet = function(requests) {
        if (!Array.isArray(requests)) {
            throw 'GreetFormatter.Greet: first argument must be Array'
        }
        return $.ajax({
            method: 'post', url: this.endpoint + '/remoto/GreetFormatter.Greet',
            data: JSON.stringify(requests),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
    }
    
    // new $.Greeter creates a jQuery client.
    // Greeter provides greeting services.
    // Usage:
    //     var remoteGreeter = new $.Greeter({endpoint: "https://example.machinebox.io/remoto"})
    $.remoto.Greeter = function(options) {
        options = options || {}
        this.endpoint = options.endpoint || "http://localhost:8080"
    }
    
    // 
    // Usage:
    //     remoteGreeter.Greet([{
    //         name: null,
    //     }])
    //       .done(function(responses) { /* called on success */ })
    //       .fail(function(){ /* called on failure */ })
    //       .always(function(){ /* always gets called */ })
    $.remoto.Greeter.prototype.Greet = function(requests) {
        if (!Array.isArray(requests)) {
            throw 'Greeter.Greet: first argument must be Array'
        }
        return $.ajax({
            method: 'post', url: this.endpoint + '/remoto/Greeter.Greet',
            data: JSON.stringify(requests),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
    }
    
    // 
    // Usage:
    //     remoteGreeter.GreetPhoto([{
    //         photo: null,
    //         name: null,
    //     }])
    //       .done(function(responses) { /* called on success */ })
    //       .fail(function(){ /* called on failure */ })
    //       .always(function(){ /* always gets called */ })
    $.remoto.Greeter.prototype.GreetPhoto = function(requests) {
        if (!Array.isArray(requests)) {
            throw 'Greeter.GreetPhoto: first argument must be Array'
        }
        return $.ajax({
            method: 'post', url: this.endpoint + '/remoto/Greeter.GreetPhoto',
            data: JSON.stringify(requests),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
    }
    
})(jQuery)
