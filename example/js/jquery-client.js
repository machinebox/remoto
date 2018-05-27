(function($){
    
    // new $.GreetFormatter creates a jQuery client.
    // GreetFormatter provides formattable greeting services.
    // Usage:
    //     var remoteGreetFormatter = new $.GreetFormatter("https://example.machinebox.io/remoto")
    $.GreetFormatter = function(endpoint) {
        this.endpoint = endpoint
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
    $.GreetFormatter.prototype.Greet = function(requests) {
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
    //     var remoteGreeter = new $.Greeter("https://example.machinebox.io/remoto")
    $.Greeter = function(endpoint) {
        this.endpoint = endpoint
    }
    
    // 
    // Usage:
    //     remoteGreeter.Greet([{
    //         name: null,
    //     }])
    //       .done(function(responses) { /* called on success */ })
    //       .fail(function(){ /* called on failure */ })
    //       .always(function(){ /* always gets called */ })
    $.Greeter.prototype.Greet = function(requests) {
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
    
})(jQuery)
