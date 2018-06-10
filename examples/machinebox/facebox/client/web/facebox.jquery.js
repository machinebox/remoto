(function($){
	$.facebox = $.facebox || {}
	
	// $.facebox.FaceboxClient is a client capable of
	// interacting with Facebox services.
	// Usage:
	// 	var client = new $.facebox.FaceboxClient({
	// 		endpoint: "http://localhost:8080"
	// 	})
	$.facebox.FaceboxClient = function(options){
		this.options = this.options || {}
		this.options.endpoint = this.options.endpoint || "http://localhost:8080"
	}
	
	// Must pass in an instance of $.facebox.CheckFaceprintRequest.
	$.facebox.FaceboxClient.prototype.CheckFaceprint = function(checkFaceprintRequest) {
		return this.CheckFaceprintMulti([checkFaceprintRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.CheckFaceprintMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.CheckFaceprint.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.CheckFaceprintMulti = function(checkFaceprintRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.CheckFaceprint',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.CheckFaceprintResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.CheckFileRequest.
	$.facebox.FaceboxClient.prototype.CheckFile = function(checkFileRequest) {
		return this.CheckFileMulti([checkFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.CheckFileMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.CheckFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.CheckFileMulti = function(checkFileRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.CheckFile',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.CheckFileResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.CheckURLRequest.
	$.facebox.FaceboxClient.prototype.CheckURL = function(checkURLRequest) {
		return this.CheckURLMulti([checkURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.CheckURLMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.CheckURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.CheckURLMulti = function(checkURLRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.CheckURL',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.CheckURLResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.FaceprintCompareRequest.
	$.facebox.FaceboxClient.prototype.FaceprintCompare = function(faceprintCompareRequest) {
		return this.FaceprintCompareMulti([faceprintCompareRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.FaceprintCompareMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.FaceprintCompare.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.FaceprintCompareMulti = function(faceprintCompareRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.FaceprintCompare',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.FaceprintCompareResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.GetStateRequest.
	$.facebox.FaceboxClient.prototype.GetState = function(getStateRequest) {
		return this.GetStateMulti([getStateRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.GetStateMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.GetState.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.GetStateMulti = function(getStateRequests) {
		
		if (getStateRequests.length > 1) {
			throw '$.facebox.FaceboxClient.GetStateMulti: batch requests are not supported for file responses, use $.facebox.FaceboxClient.GetState instead.'
		}var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.GetState',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.remototypes.FileResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.PutStateRequest.
	$.facebox.FaceboxClient.prototype.PutState = function(putStateRequest) {
		return this.PutStateMulti([putStateRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.PutStateMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.PutState.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.PutStateMulti = function(putStateRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.PutState',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.PutStateResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.RemoveIDRequest.
	$.facebox.FaceboxClient.prototype.RemoveID = function(removeIDRequest) {
		return this.RemoveIDMulti([removeIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.RemoveIDMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.RemoveID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.RemoveIDMulti = function(removeIDRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.RemoveID',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.RemoveIDResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.RenameRequest.
	$.facebox.FaceboxClient.prototype.Rename = function(renameRequest) {
		return this.RenameMulti([renameRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.RenameMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.Rename.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.RenameMulti = function(renameRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.Rename',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.RenameResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.RenameIDRequest.
	$.facebox.FaceboxClient.prototype.RenameID = function(renameIDRequest) {
		return this.RenameIDMulti([renameIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.RenameIDMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.RenameID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.RenameIDMulti = function(renameIDRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.RenameID',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.RenameIDResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.SimilarFileRequest.
	$.facebox.FaceboxClient.prototype.SimilarFile = function(similarFileRequest) {
		return this.SimilarFileMulti([similarFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.SimilarFileMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.SimilarFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.SimilarFileMulti = function(similarFileRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.SimilarFile',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.SimilarFileResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.SimilarIDRequest.
	$.facebox.FaceboxClient.prototype.SimilarID = function(similarIDRequest) {
		return this.SimilarIDMulti([similarIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.SimilarIDMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.SimilarID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.SimilarIDMulti = function(similarIDRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.SimilarID',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.SimilarIDResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.SimilarURLRequest.
	$.facebox.FaceboxClient.prototype.SimilarURL = function(similarURLRequest) {
		return this.SimilarURLMulti([similarURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.SimilarURLMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.SimilarURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.SimilarURLMulti = function(similarURLRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.SimilarURL',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.SimilarURLResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.TeachFaceprintRequest.
	$.facebox.FaceboxClient.prototype.TeachFaceprint = function(teachFaceprintRequest) {
		return this.TeachFaceprintMulti([teachFaceprintRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.TeachFaceprintMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.TeachFaceprint.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.TeachFaceprintMulti = function(teachFaceprintRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.TeachFaceprint',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.TeachFaceprintResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.TeachFileRequest.
	$.facebox.FaceboxClient.prototype.TeachFile = function(teachFileRequest) {
		return this.TeachFileMulti([teachFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.TeachFileMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.TeachFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.TeachFileMulti = function(teachFileRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.TeachFile',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.TeachFileResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.TeachURLRequest.
	$.facebox.FaceboxClient.prototype.TeachURL = function(teachURLRequest) {
		return this.TeachURLMulti([teachURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.FaceboxClient.prototype.TeachURLMulti is a batch
	// version of $.facebox.FaceboxClient.prototype.TeachURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.FaceboxClient.prototype.TeachURLMulti = function(teachURLRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Facebox.TeachURL',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.TeachURLResponse(response))
			})
			return responseObjects
		})
	}
	
	
	// CheckFaceprintRequest is the request object for CheckFaceprint calls.
$.facebox.CheckFaceprintRequest = function(data) {
		this.data = data
	}
	$.facebox.CheckFaceprintRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.CheckFaceprintRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFaceprints gets the faceprints from this object.
	$.facebox.CheckFaceprintRequest.prototype.getFaceprints = function() {
		return this.data.faceprints
	}
	
	
	// setFaceprints sets the faceprints on this object.
	$.facebox.CheckFaceprintRequest.prototype.setFaceprints = function(faceprints) {
		this.data.faceprints = faceprints
	}
	
	// CheckFaceprintResponse is the response object for CheckFaceprint calls.
$.facebox.CheckFaceprintResponse = function(data) {
		this.data = data
	}
	$.facebox.CheckFaceprintResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.CheckFaceprintResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.CheckFaceprintResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// CheckFileRequest is the request object for CheckFile calls.
$.facebox.CheckFileRequest = function(data) {
		this.data = data
	}
	$.facebox.CheckFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.CheckFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFile gets the file from this object.
	$.facebox.CheckFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.CheckFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// CheckFileResponse is the response object for CheckFile calls.
$.facebox.CheckFileResponse = function(data) {
		this.data = data
	}
	$.facebox.CheckFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.CheckFileResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.CheckFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// CheckURLRequest is the request object for CheckURL calls.
$.facebox.CheckURLRequest = function(data) {
		this.data = data
	}
	$.facebox.CheckURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.CheckURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFile gets the file from this object.
	$.facebox.CheckURLRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.CheckURLRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// CheckURLResponse is the response object for CheckURL calls.
$.facebox.CheckURLResponse = function(data) {
		this.data = data
	}
	$.facebox.CheckURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.CheckURLResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.CheckURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// Face describes a face.
$.facebox.Face = function(data) {
		this.data = data
	}
	$.facebox.Face.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getID gets the iD from this object.
	$.facebox.Face.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.Face.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.Face.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.Face.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getMatched gets the matched from this object.
	$.facebox.Face.prototype.getMatched = function() {
		return this.data.matched
	}
	
	
	// setMatched sets the matched on this object.
	$.facebox.Face.prototype.setMatched = function(matched) {
		this.data.matched = matched
	}
	
	// getFaceprint gets the faceprint from this object.
	$.facebox.Face.prototype.getFaceprint = function() {
		return this.data.faceprint
	}
	
	
	// setFaceprint sets the faceprint on this object.
	$.facebox.Face.prototype.setFaceprint = function(faceprint) {
		this.data.faceprint = faceprint
	}
	
	// getRect gets the rect from this object.
	$.facebox.Face.prototype.getRect = function() {
		return this.data.rect
	}
	
	
	// setRect sets the rect on this object.
	$.facebox.Face.prototype.setRect = function(rect) {
		this.data.rect = rect
	}
	
	// FaceprintCompareRequest is the request object for FaceprintCompare calls.
$.facebox.FaceprintCompareRequest = function(data) {
		this.data = data
	}
	$.facebox.FaceprintCompareRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.FaceprintCompareRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getTarget gets the target from this object.
	$.facebox.FaceprintCompareRequest.prototype.getTarget = function() {
		return this.data.target
	}
	
	
	// setTarget sets the target on this object.
	$.facebox.FaceprintCompareRequest.prototype.setTarget = function(target) {
		this.data.target = target
	}
	
	// getFaceprints gets the faceprints from this object.
	$.facebox.FaceprintCompareRequest.prototype.getFaceprints = function() {
		return this.data.faceprints
	}
	
	
	// setFaceprints sets the faceprints on this object.
	$.facebox.FaceprintCompareRequest.prototype.setFaceprints = function(faceprints) {
		this.data.faceprints = faceprints
	}
	
	// FaceprintCompareResponse is the response object for FaceprintCompare calls.
$.facebox.FaceprintCompareResponse = function(data) {
		this.data = data
	}
	$.facebox.FaceprintCompareResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getConfidences gets the confidences from this object.
	$.facebox.FaceprintCompareResponse.prototype.getConfidences = function() {
		return this.data.confidences
	}
	
	
	// getError gets the error from this object.
	$.facebox.FaceprintCompareResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// FaceprintFace is a face.
$.facebox.FaceprintFace = function(data) {
		this.data = data
	}
	$.facebox.FaceprintFace.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getMatched gets the matched from this object.
	$.facebox.FaceprintFace.prototype.getMatched = function() {
		return this.data.matched
	}
	
	
	// setMatched sets the matched on this object.
	$.facebox.FaceprintFace.prototype.setMatched = function(matched) {
		this.data.matched = matched
	}
	
	// getConfidence gets the confidence from this object.
	$.facebox.FaceprintFace.prototype.getConfidence = function() {
		return this.data.confidence
	}
	
	
	// setConfidence sets the confidence on this object.
	$.facebox.FaceprintFace.prototype.setConfidence = function(confidence) {
		this.data.confidence = confidence
	}
	
	// getID gets the iD from this object.
	$.facebox.FaceprintFace.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.FaceprintFace.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.FaceprintFace.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.FaceprintFace.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// GetStateRequest is the request object for GetState calls.
$.facebox.GetStateRequest = function(data) {
		this.data = data
	}
	$.facebox.GetStateRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.GetStateRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// PutStateRequest is the request object for PutState calls.
$.facebox.PutStateRequest = function(data) {
		this.data = data
	}
	$.facebox.PutStateRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.PutStateRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getStateFile gets the stateFile from this object.
	$.facebox.PutStateRequest.prototype.getStateFile = function() {
		return this.data.stateFile
	}
	
	
	// setStateFile sets the stateFile on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.PutStateRequest.prototype.setStateFile = function(request, stateFile) {
		this.data.stateFile = request._addFile(stateFile)
	}
	
	
	// PutStateResponse is the response object for PutState calls.
$.facebox.PutStateResponse = function(data) {
		this.data = data
	}
	$.facebox.PutStateResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.PutStateResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// Rect is a bounding box describing a rectangle of an image.
$.facebox.Rect = function(data) {
		this.data = data
	}
	$.facebox.Rect.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getTop gets the top from this object.
	$.facebox.Rect.prototype.getTop = function() {
		return this.data.top
	}
	
	
	// setTop sets the top on this object.
	$.facebox.Rect.prototype.setTop = function(top) {
		this.data.top = top
	}
	
	// getLeft gets the left from this object.
	$.facebox.Rect.prototype.getLeft = function() {
		return this.data.left
	}
	
	
	// setLeft sets the left on this object.
	$.facebox.Rect.prototype.setLeft = function(left) {
		this.data.left = left
	}
	
	// getWidth gets the width from this object.
	$.facebox.Rect.prototype.getWidth = function() {
		return this.data.width
	}
	
	
	// setWidth sets the width on this object.
	$.facebox.Rect.prototype.setWidth = function(width) {
		this.data.width = width
	}
	
	// getHeight gets the height from this object.
	$.facebox.Rect.prototype.getHeight = function() {
		return this.data.height
	}
	
	
	// setHeight sets the height on this object.
	$.facebox.Rect.prototype.setHeight = function(height) {
		this.data.height = height
	}
	
	// RemoveIDRequest is the request object for RemoveID calls.
$.facebox.RemoveIDRequest = function(data) {
		this.data = data
	}
	$.facebox.RemoveIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.RemoveIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.RemoveIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.RemoveIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// RemoveIDResponse is the response object for RemoveID calls.
$.facebox.RemoveIDResponse = function(data) {
		this.data = data
	}
	$.facebox.RemoveIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.RemoveIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// RenameIDRequest is the request object for RenameID calls.
$.facebox.RenameIDRequest = function(data) {
		this.data = data
	}
	$.facebox.RenameIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.RenameIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.RenameIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.RenameIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.RenameIDRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.RenameIDRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// RenameIDResponse is the response object for RenameID calls.
$.facebox.RenameIDResponse = function(data) {
		this.data = data
	}
	$.facebox.RenameIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.RenameIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// RenameRequest is the request object for Rename calls.
$.facebox.RenameRequest = function(data) {
		this.data = data
	}
	$.facebox.RenameRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.RenameRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFrom gets the from from this object.
	$.facebox.RenameRequest.prototype.getFrom = function() {
		return this.data.from
	}
	
	
	// setFrom sets the from on this object.
	$.facebox.RenameRequest.prototype.setFrom = function(from) {
		this.data.from = from
	}
	
	// getTo gets the to from this object.
	$.facebox.RenameRequest.prototype.getTo = function() {
		return this.data.to
	}
	
	
	// setTo sets the to on this object.
	$.facebox.RenameRequest.prototype.setTo = function(to) {
		this.data.to = to
	}
	
	// RenameResponse is the response object for Rename calls.
$.facebox.RenameResponse = function(data) {
		this.data = data
	}
	$.facebox.RenameResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.RenameResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarFace is a detected face with similar matching faces.
$.facebox.SimilarFace = function(data) {
		this.data = data
	}
	$.facebox.SimilarFace.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getRect gets the rect from this object.
	$.facebox.SimilarFace.prototype.getRect = function() {
		return this.data.rect
	}
	
	
	// setRect sets the rect on this object.
	$.facebox.SimilarFace.prototype.setRect = function(rect) {
		this.data.rect = rect
	}
	
	// getSimilarFaces gets the similarFaces from this object.
	$.facebox.SimilarFace.prototype.getSimilarFaces = function() {
		return this.data.similarFaces
	}
	
	
	// setSimilarFaces sets the similarFaces on this object.
	$.facebox.SimilarFace.prototype.setSimilarFaces = function(similarFaces) {
		this.data.similarFaces = similarFaces
	}
	
	// SimilarFileRequest is the request object for SimilarFile calls.
$.facebox.SimilarFileRequest = function(data) {
		this.data = data
	}
	$.facebox.SimilarFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.SimilarFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFile gets the file from this object.
	$.facebox.SimilarFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.SimilarFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// SimilarFileResponse is the response object for SimilarFile calls.
$.facebox.SimilarFileResponse = function(data) {
		this.data = data
	}
	$.facebox.SimilarFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.SimilarFileResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.SimilarFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarIDRequest is the request object for SimilarID calls.
$.facebox.SimilarIDRequest = function(data) {
		this.data = data
	}
	$.facebox.SimilarIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.SimilarIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.SimilarIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.SimilarIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// SimilarIDResponse is the response object for SimilarID calls.
$.facebox.SimilarIDResponse = function(data) {
		this.data = data
	}
	$.facebox.SimilarIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.SimilarIDResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.SimilarIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarURLRequest is the request object for SimilarURL calls.
$.facebox.SimilarURLRequest = function(data) {
		this.data = data
	}
	$.facebox.SimilarURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.SimilarURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getURL gets the uRL from this object.
	$.facebox.SimilarURLRequest.prototype.getURL = function() {
		return this.data.uRL
	}
	
	
	// setURL sets the uRL on this object.
	$.facebox.SimilarURLRequest.prototype.setURL = function(uRL) {
		this.data.uRL = uRL
	}
	
	// SimilarURLResponse is the response object for SimilarURL calls.
$.facebox.SimilarURLResponse = function(data) {
		this.data = data
	}
	$.facebox.SimilarURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.SimilarURLResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.SimilarURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachFaceprintRequest is the request object for TeachFaceprint calls.
$.facebox.TeachFaceprintRequest = function(data) {
		this.data = data
	}
	$.facebox.TeachFaceprintRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.TeachFaceprintRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.TeachFaceprintRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.TeachFaceprintRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.TeachFaceprintRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.TeachFaceprintRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getFaceprint gets the faceprint from this object.
	$.facebox.TeachFaceprintRequest.prototype.getFaceprint = function() {
		return this.data.faceprint
	}
	
	
	// setFaceprint sets the faceprint on this object.
	$.facebox.TeachFaceprintRequest.prototype.setFaceprint = function(faceprint) {
		this.data.faceprint = faceprint
	}
	
	// TeachFaceprintResponse is the response object for TeachFaceprint calls.
$.facebox.TeachFaceprintResponse = function(data) {
		this.data = data
	}
	$.facebox.TeachFaceprintResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.TeachFaceprintResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachFileRequest is the request object for TeachFile calls.
$.facebox.TeachFileRequest = function(data) {
		this.data = data
	}
	$.facebox.TeachFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.TeachFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.TeachFileRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.TeachFileRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.TeachFileRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.TeachFileRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getFile gets the file from this object.
	$.facebox.TeachFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.TeachFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// TeachFileResponse is the response object for TeachFile calls.
$.facebox.TeachFileResponse = function(data) {
		this.data = data
	}
	$.facebox.TeachFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.TeachFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachURLRequest is the request object for TeachURL calls.
$.facebox.TeachURLRequest = function(data) {
		this.data = data
	}
	$.facebox.TeachURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.TeachURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.TeachURLRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.TeachURLRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.TeachURLRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.TeachURLRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getURL gets the uRL from this object.
	$.facebox.TeachURLRequest.prototype.getURL = function() {
		return this.data.uRL
	}
	
	
	// setURL sets the uRL on this object.
	$.facebox.TeachURLRequest.prototype.setURL = function(uRL) {
		this.data.uRL = uRL
	}
	
	// TeachURLResponse is the response object for TeachURL calls.
$.facebox.TeachURLResponse = function(data) {
		this.data = data
	}
	$.facebox.TeachURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.TeachURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// $.facebox.SuggestionboxClient is a client capable of
	// interacting with Suggestionbox services.
	// Usage:
	// 	var client = new $.facebox.SuggestionboxClient({
	// 		endpoint: "http://localhost:8080"
	// 	})
	$.facebox.SuggestionboxClient = function(options){
		this.options = this.options || {}
		this.options.endpoint = this.options.endpoint || "http://localhost:8080"
	}
	
	// Must pass in an instance of $.facebox.CheckFaceprintRequest.
	$.facebox.SuggestionboxClient.prototype.CheckFaceprint = function(checkFaceprintRequest) {
		return this.CheckFaceprintMulti([checkFaceprintRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.CheckFaceprintMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.CheckFaceprint.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.CheckFaceprintMulti = function(checkFaceprintRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.CheckFaceprint',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.CheckFaceprintResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.CheckFileRequest.
	$.facebox.SuggestionboxClient.prototype.CheckFile = function(checkFileRequest) {
		return this.CheckFileMulti([checkFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.CheckFileMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.CheckFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.CheckFileMulti = function(checkFileRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.CheckFile',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.CheckFileResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.CheckURLRequest.
	$.facebox.SuggestionboxClient.prototype.CheckURL = function(checkURLRequest) {
		return this.CheckURLMulti([checkURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.CheckURLMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.CheckURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.CheckURLMulti = function(checkURLRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.CheckURL',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.CheckURLResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.FaceprintCompareRequest.
	$.facebox.SuggestionboxClient.prototype.FaceprintCompare = function(faceprintCompareRequest) {
		return this.FaceprintCompareMulti([faceprintCompareRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.FaceprintCompareMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.FaceprintCompare.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.FaceprintCompareMulti = function(faceprintCompareRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.FaceprintCompare',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.FaceprintCompareResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.GetStateRequest.
	$.facebox.SuggestionboxClient.prototype.GetState = function(getStateRequest) {
		return this.GetStateMulti([getStateRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.GetStateMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.GetState.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.GetStateMulti = function(getStateRequests) {
		
		if (getStateRequests.length > 1) {
			throw '$.facebox.SuggestionboxClient.GetStateMulti: batch requests are not supported for file responses, use $.facebox.SuggestionboxClient.GetState instead.'
		}var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.GetState',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.remototypes.FileResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.PutStateRequest.
	$.facebox.SuggestionboxClient.prototype.PutState = function(putStateRequest) {
		return this.PutStateMulti([putStateRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.PutStateMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.PutState.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.PutStateMulti = function(putStateRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.PutState',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.PutStateResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.RemoveIDRequest.
	$.facebox.SuggestionboxClient.prototype.RemoveID = function(removeIDRequest) {
		return this.RemoveIDMulti([removeIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.RemoveIDMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.RemoveID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.RemoveIDMulti = function(removeIDRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.RemoveID',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.RemoveIDResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.RenameRequest.
	$.facebox.SuggestionboxClient.prototype.Rename = function(renameRequest) {
		return this.RenameMulti([renameRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.RenameMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.Rename.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.RenameMulti = function(renameRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.Rename',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.RenameResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.RenameIDRequest.
	$.facebox.SuggestionboxClient.prototype.RenameID = function(renameIDRequest) {
		return this.RenameIDMulti([renameIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.RenameIDMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.RenameID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.RenameIDMulti = function(renameIDRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.RenameID',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.RenameIDResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.SimilarFileRequest.
	$.facebox.SuggestionboxClient.prototype.SimilarFile = function(similarFileRequest) {
		return this.SimilarFileMulti([similarFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.SimilarFileMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.SimilarFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.SimilarFileMulti = function(similarFileRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.SimilarFile',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.SimilarFileResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.SimilarIDRequest.
	$.facebox.SuggestionboxClient.prototype.SimilarID = function(similarIDRequest) {
		return this.SimilarIDMulti([similarIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.SimilarIDMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.SimilarID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.SimilarIDMulti = function(similarIDRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.SimilarID',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.SimilarIDResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.SimilarURLRequest.
	$.facebox.SuggestionboxClient.prototype.SimilarURL = function(similarURLRequest) {
		return this.SimilarURLMulti([similarURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.SimilarURLMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.SimilarURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.SimilarURLMulti = function(similarURLRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.SimilarURL',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.SimilarURLResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.TeachFaceprintRequest.
	$.facebox.SuggestionboxClient.prototype.TeachFaceprint = function(teachFaceprintRequest) {
		return this.TeachFaceprintMulti([teachFaceprintRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.TeachFaceprintMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.TeachFaceprint.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.TeachFaceprintMulti = function(teachFaceprintRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.TeachFaceprint',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.TeachFaceprintResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.TeachFileRequest.
	$.facebox.SuggestionboxClient.prototype.TeachFile = function(teachFileRequest) {
		return this.TeachFileMulti([teachFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.TeachFileMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.TeachFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.TeachFileMulti = function(teachFileRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.TeachFile',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.TeachFileResponse(response))
			})
			return responseObjects
		})
	}
	
	// Must pass in an instance of $.facebox.TeachURLRequest.
	$.facebox.SuggestionboxClient.prototype.TeachURL = function(teachURLRequest) {
		return this.TeachURLMulti([teachURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.facebox.SuggestionboxClient.prototype.TeachURLMulti is a batch
	// version of $.facebox.SuggestionboxClient.prototype.TeachURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.facebox.SuggestionboxClient.prototype.TeachURLMulti = function(teachURLRequests) {
		var filesCount = 0
		var formData = new FormData()
		requests.forEach(function(request){
			request._files.forEach(function(file) {
				formData.append(file, request._files[file])
			})
		})
		formData.append('json', JSON.stringify(requests))
		return $.ajax({
			method: 'post', url: this.options.endpoint + '/remoto/Suggestionbox.TeachURL',
			data: formData,
			contentType: false,
			processData: false
		}).then(function(responses) {
			var responseObjects = []
			responses.forEach(function(response){
				responseObjects.push(new $.facebox.TeachURLResponse(response))
			})
			return responseObjects
		})
	}
	
	
	// CheckFaceprintRequest is the request object for CheckFaceprint calls.
$.facebox.CheckFaceprintRequest = function(data) {
		this.data = data
	}
	$.facebox.CheckFaceprintRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.CheckFaceprintRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFaceprints gets the faceprints from this object.
	$.facebox.CheckFaceprintRequest.prototype.getFaceprints = function() {
		return this.data.faceprints
	}
	
	
	// setFaceprints sets the faceprints on this object.
	$.facebox.CheckFaceprintRequest.prototype.setFaceprints = function(faceprints) {
		this.data.faceprints = faceprints
	}
	
	// CheckFaceprintResponse is the response object for CheckFaceprint calls.
$.facebox.CheckFaceprintResponse = function(data) {
		this.data = data
	}
	$.facebox.CheckFaceprintResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.CheckFaceprintResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.CheckFaceprintResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// CheckFileRequest is the request object for CheckFile calls.
$.facebox.CheckFileRequest = function(data) {
		this.data = data
	}
	$.facebox.CheckFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.CheckFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFile gets the file from this object.
	$.facebox.CheckFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.CheckFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// CheckFileResponse is the response object for CheckFile calls.
$.facebox.CheckFileResponse = function(data) {
		this.data = data
	}
	$.facebox.CheckFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.CheckFileResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.CheckFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// CheckURLRequest is the request object for CheckURL calls.
$.facebox.CheckURLRequest = function(data) {
		this.data = data
	}
	$.facebox.CheckURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.CheckURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFile gets the file from this object.
	$.facebox.CheckURLRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.CheckURLRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// CheckURLResponse is the response object for CheckURL calls.
$.facebox.CheckURLResponse = function(data) {
		this.data = data
	}
	$.facebox.CheckURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.CheckURLResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.CheckURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// Face describes a face.
$.facebox.Face = function(data) {
		this.data = data
	}
	$.facebox.Face.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getID gets the iD from this object.
	$.facebox.Face.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.Face.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.Face.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.Face.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getMatched gets the matched from this object.
	$.facebox.Face.prototype.getMatched = function() {
		return this.data.matched
	}
	
	
	// setMatched sets the matched on this object.
	$.facebox.Face.prototype.setMatched = function(matched) {
		this.data.matched = matched
	}
	
	// getFaceprint gets the faceprint from this object.
	$.facebox.Face.prototype.getFaceprint = function() {
		return this.data.faceprint
	}
	
	
	// setFaceprint sets the faceprint on this object.
	$.facebox.Face.prototype.setFaceprint = function(faceprint) {
		this.data.faceprint = faceprint
	}
	
	// getRect gets the rect from this object.
	$.facebox.Face.prototype.getRect = function() {
		return this.data.rect
	}
	
	
	// setRect sets the rect on this object.
	$.facebox.Face.prototype.setRect = function(rect) {
		this.data.rect = rect
	}
	
	// FaceprintCompareRequest is the request object for FaceprintCompare calls.
$.facebox.FaceprintCompareRequest = function(data) {
		this.data = data
	}
	$.facebox.FaceprintCompareRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.FaceprintCompareRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getTarget gets the target from this object.
	$.facebox.FaceprintCompareRequest.prototype.getTarget = function() {
		return this.data.target
	}
	
	
	// setTarget sets the target on this object.
	$.facebox.FaceprintCompareRequest.prototype.setTarget = function(target) {
		this.data.target = target
	}
	
	// getFaceprints gets the faceprints from this object.
	$.facebox.FaceprintCompareRequest.prototype.getFaceprints = function() {
		return this.data.faceprints
	}
	
	
	// setFaceprints sets the faceprints on this object.
	$.facebox.FaceprintCompareRequest.prototype.setFaceprints = function(faceprints) {
		this.data.faceprints = faceprints
	}
	
	// FaceprintCompareResponse is the response object for FaceprintCompare calls.
$.facebox.FaceprintCompareResponse = function(data) {
		this.data = data
	}
	$.facebox.FaceprintCompareResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getConfidences gets the confidences from this object.
	$.facebox.FaceprintCompareResponse.prototype.getConfidences = function() {
		return this.data.confidences
	}
	
	
	// getError gets the error from this object.
	$.facebox.FaceprintCompareResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// FaceprintFace is a face.
$.facebox.FaceprintFace = function(data) {
		this.data = data
	}
	$.facebox.FaceprintFace.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getMatched gets the matched from this object.
	$.facebox.FaceprintFace.prototype.getMatched = function() {
		return this.data.matched
	}
	
	
	// setMatched sets the matched on this object.
	$.facebox.FaceprintFace.prototype.setMatched = function(matched) {
		this.data.matched = matched
	}
	
	// getConfidence gets the confidence from this object.
	$.facebox.FaceprintFace.prototype.getConfidence = function() {
		return this.data.confidence
	}
	
	
	// setConfidence sets the confidence on this object.
	$.facebox.FaceprintFace.prototype.setConfidence = function(confidence) {
		this.data.confidence = confidence
	}
	
	// getID gets the iD from this object.
	$.facebox.FaceprintFace.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.FaceprintFace.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.FaceprintFace.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.FaceprintFace.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// GetStateRequest is the request object for GetState calls.
$.facebox.GetStateRequest = function(data) {
		this.data = data
	}
	$.facebox.GetStateRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.GetStateRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// PutStateRequest is the request object for PutState calls.
$.facebox.PutStateRequest = function(data) {
		this.data = data
	}
	$.facebox.PutStateRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.PutStateRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getStateFile gets the stateFile from this object.
	$.facebox.PutStateRequest.prototype.getStateFile = function() {
		return this.data.stateFile
	}
	
	
	// setStateFile sets the stateFile on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.PutStateRequest.prototype.setStateFile = function(request, stateFile) {
		this.data.stateFile = request._addFile(stateFile)
	}
	
	
	// PutStateResponse is the response object for PutState calls.
$.facebox.PutStateResponse = function(data) {
		this.data = data
	}
	$.facebox.PutStateResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.PutStateResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// Rect is a bounding box describing a rectangle of an image.
$.facebox.Rect = function(data) {
		this.data = data
	}
	$.facebox.Rect.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getTop gets the top from this object.
	$.facebox.Rect.prototype.getTop = function() {
		return this.data.top
	}
	
	
	// setTop sets the top on this object.
	$.facebox.Rect.prototype.setTop = function(top) {
		this.data.top = top
	}
	
	// getLeft gets the left from this object.
	$.facebox.Rect.prototype.getLeft = function() {
		return this.data.left
	}
	
	
	// setLeft sets the left on this object.
	$.facebox.Rect.prototype.setLeft = function(left) {
		this.data.left = left
	}
	
	// getWidth gets the width from this object.
	$.facebox.Rect.prototype.getWidth = function() {
		return this.data.width
	}
	
	
	// setWidth sets the width on this object.
	$.facebox.Rect.prototype.setWidth = function(width) {
		this.data.width = width
	}
	
	// getHeight gets the height from this object.
	$.facebox.Rect.prototype.getHeight = function() {
		return this.data.height
	}
	
	
	// setHeight sets the height on this object.
	$.facebox.Rect.prototype.setHeight = function(height) {
		this.data.height = height
	}
	
	// RemoveIDRequest is the request object for RemoveID calls.
$.facebox.RemoveIDRequest = function(data) {
		this.data = data
	}
	$.facebox.RemoveIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.RemoveIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.RemoveIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.RemoveIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// RemoveIDResponse is the response object for RemoveID calls.
$.facebox.RemoveIDResponse = function(data) {
		this.data = data
	}
	$.facebox.RemoveIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.RemoveIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// RenameIDRequest is the request object for RenameID calls.
$.facebox.RenameIDRequest = function(data) {
		this.data = data
	}
	$.facebox.RenameIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.RenameIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.RenameIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.RenameIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.RenameIDRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.RenameIDRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// RenameIDResponse is the response object for RenameID calls.
$.facebox.RenameIDResponse = function(data) {
		this.data = data
	}
	$.facebox.RenameIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.RenameIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// RenameRequest is the request object for Rename calls.
$.facebox.RenameRequest = function(data) {
		this.data = data
	}
	$.facebox.RenameRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.RenameRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFrom gets the from from this object.
	$.facebox.RenameRequest.prototype.getFrom = function() {
		return this.data.from
	}
	
	
	// setFrom sets the from on this object.
	$.facebox.RenameRequest.prototype.setFrom = function(from) {
		this.data.from = from
	}
	
	// getTo gets the to from this object.
	$.facebox.RenameRequest.prototype.getTo = function() {
		return this.data.to
	}
	
	
	// setTo sets the to on this object.
	$.facebox.RenameRequest.prototype.setTo = function(to) {
		this.data.to = to
	}
	
	// RenameResponse is the response object for Rename calls.
$.facebox.RenameResponse = function(data) {
		this.data = data
	}
	$.facebox.RenameResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.RenameResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarFace is a detected face with similar matching faces.
$.facebox.SimilarFace = function(data) {
		this.data = data
	}
	$.facebox.SimilarFace.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getRect gets the rect from this object.
	$.facebox.SimilarFace.prototype.getRect = function() {
		return this.data.rect
	}
	
	
	// setRect sets the rect on this object.
	$.facebox.SimilarFace.prototype.setRect = function(rect) {
		this.data.rect = rect
	}
	
	// getSimilarFaces gets the similarFaces from this object.
	$.facebox.SimilarFace.prototype.getSimilarFaces = function() {
		return this.data.similarFaces
	}
	
	
	// setSimilarFaces sets the similarFaces on this object.
	$.facebox.SimilarFace.prototype.setSimilarFaces = function(similarFaces) {
		this.data.similarFaces = similarFaces
	}
	
	// SimilarFileRequest is the request object for SimilarFile calls.
$.facebox.SimilarFileRequest = function(data) {
		this.data = data
	}
	$.facebox.SimilarFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.SimilarFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getFile gets the file from this object.
	$.facebox.SimilarFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.SimilarFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// SimilarFileResponse is the response object for SimilarFile calls.
$.facebox.SimilarFileResponse = function(data) {
		this.data = data
	}
	$.facebox.SimilarFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.SimilarFileResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.SimilarFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarIDRequest is the request object for SimilarID calls.
$.facebox.SimilarIDRequest = function(data) {
		this.data = data
	}
	$.facebox.SimilarIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.SimilarIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.SimilarIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.SimilarIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// SimilarIDResponse is the response object for SimilarID calls.
$.facebox.SimilarIDResponse = function(data) {
		this.data = data
	}
	$.facebox.SimilarIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.SimilarIDResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.SimilarIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarURLRequest is the request object for SimilarURL calls.
$.facebox.SimilarURLRequest = function(data) {
		this.data = data
	}
	$.facebox.SimilarURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.SimilarURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getURL gets the uRL from this object.
	$.facebox.SimilarURLRequest.prototype.getURL = function() {
		return this.data.uRL
	}
	
	
	// setURL sets the uRL on this object.
	$.facebox.SimilarURLRequest.prototype.setURL = function(uRL) {
		this.data.uRL = uRL
	}
	
	// SimilarURLResponse is the response object for SimilarURL calls.
$.facebox.SimilarURLResponse = function(data) {
		this.data = data
	}
	$.facebox.SimilarURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.facebox.SimilarURLResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.facebox.SimilarURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachFaceprintRequest is the request object for TeachFaceprint calls.
$.facebox.TeachFaceprintRequest = function(data) {
		this.data = data
	}
	$.facebox.TeachFaceprintRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.TeachFaceprintRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.TeachFaceprintRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.TeachFaceprintRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.TeachFaceprintRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.TeachFaceprintRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getFaceprint gets the faceprint from this object.
	$.facebox.TeachFaceprintRequest.prototype.getFaceprint = function() {
		return this.data.faceprint
	}
	
	
	// setFaceprint sets the faceprint on this object.
	$.facebox.TeachFaceprintRequest.prototype.setFaceprint = function(faceprint) {
		this.data.faceprint = faceprint
	}
	
	// TeachFaceprintResponse is the response object for TeachFaceprint calls.
$.facebox.TeachFaceprintResponse = function(data) {
		this.data = data
	}
	$.facebox.TeachFaceprintResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.TeachFaceprintResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachFileRequest is the request object for TeachFile calls.
$.facebox.TeachFileRequest = function(data) {
		this.data = data
	}
	$.facebox.TeachFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.TeachFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.TeachFileRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.TeachFileRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.TeachFileRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.TeachFileRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getFile gets the file from this object.
	$.facebox.TeachFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.facebox.TeachFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// TeachFileResponse is the response object for TeachFile calls.
$.facebox.TeachFileResponse = function(data) {
		this.data = data
	}
	$.facebox.TeachFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.TeachFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachURLRequest is the request object for TeachURL calls.
$.facebox.TeachURLRequest = function(data) {
		this.data = data
	}
	$.facebox.TeachURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.facebox.TeachURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.facebox._filesCount+']'
		this._files[fieldname] = file
		$.facebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.facebox.TeachURLRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.facebox.TeachURLRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.facebox.TeachURLRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.facebox.TeachURLRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getURL gets the uRL from this object.
	$.facebox.TeachURLRequest.prototype.getURL = function() {
		return this.data.uRL
	}
	
	
	// setURL sets the uRL on this object.
	$.facebox.TeachURLRequest.prototype.setURL = function(uRL) {
		this.data.uRL = uRL
	}
	
	// TeachURLResponse is the response object for TeachURL calls.
$.facebox.TeachURLResponse = function(data) {
		this.data = data
	}
	$.facebox.TeachURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.facebox.TeachURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// _filesCount keeps track of the number of files being added, and is used
	// to generate unique field names.
	$.facebox._filesCount = 0

})(jQuery)
