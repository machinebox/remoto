
import * as facebox from '../facebox.es6.js';

let options = new facebox.FaceboxClientOptions()
options.endpoint = "http://localhost:8080"
let client = new facebox.FaceboxClient(options)

let req = new facebox.CheckURLRequest()
req.setFile(req, "filename", file)
client.CheckURL(req).then(function(response){
	console.info(response)
})
