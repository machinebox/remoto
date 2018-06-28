(function($){
	$.machinebox = $.machinebox || {}
	
	// $.machinebox.FaceboxClient is a client capable of
	// interacting with Facebox services.
	// Usage:
	// 	var client = new $.machinebox.FaceboxClient({
	// 		endpoint: "http://localhost:8080"
	// 	})
	$.machinebox.FaceboxClient = function(options){
		this.options = this.options || {}
		this.options.endpoint = this.options.endpoint || "http://localhost:8080"
	}
	// CheckFaceprint checks to see if a Faceprint matches any known
// faces.

	// Must pass in an instance of $.machinebox.CheckFaceprintRequest.
	$.machinebox.FaceboxClient.prototype.CheckFaceprint = function(checkFaceprintRequest) {
		return this.CheckFaceprintMulti([checkFaceprintRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.CheckFaceprintMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.CheckFaceprint.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.CheckFaceprintMulti = function(checkFaceprintRequests) {
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
				responseObjects.push(new $.machinebox.CheckFaceprintResponse(response))
			})
			return responseObjects
		})
	}
	// CheckFile checks an image file for faces.

	// Must pass in an instance of $.machinebox.CheckFileRequest.
	$.machinebox.FaceboxClient.prototype.CheckFile = function(checkFileRequest) {
		return this.CheckFileMulti([checkFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.CheckFileMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.CheckFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.CheckFileMulti = function(checkFileRequests) {
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
				responseObjects.push(new $.machinebox.CheckFileResponse(response))
			})
			return responseObjects
		})
	}
	// CheckURL checks a hosted image file for faces.

	// Must pass in an instance of $.machinebox.CheckURLRequest.
	$.machinebox.FaceboxClient.prototype.CheckURL = function(checkURLRequest) {
		return this.CheckURLMulti([checkURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.CheckURLMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.CheckURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.CheckURLMulti = function(checkURLRequests) {
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
				responseObjects.push(new $.machinebox.CheckURLResponse(response))
			})
			return responseObjects
		})
	}
	// FaceprintCompare compares faceprints to a specified target describing
// similarity.

	// Must pass in an instance of $.machinebox.FaceprintCompareRequest.
	$.machinebox.FaceboxClient.prototype.FaceprintCompare = function(faceprintCompareRequest) {
		return this.FaceprintCompareMulti([faceprintCompareRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.FaceprintCompareMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.FaceprintCompare.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.FaceprintCompareMulti = function(faceprintCompareRequests) {
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
				responseObjects.push(new $.machinebox.FaceprintCompareResponse(response))
			})
			return responseObjects
		})
	}
	// GetState gets the Facebox state file.

	// Must pass in an instance of $.machinebox.GetStateRequest.
	$.machinebox.FaceboxClient.prototype.GetState = function(getStateRequest) {
		return this.GetStateMulti([getStateRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.GetStateMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.GetState.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.GetStateMulti = function(getStateRequests) {
		
		if (getStateRequests.length > 1) {
			throw '$.machinebox.FaceboxClient.GetStateMulti: batch requests are not supported for file responses, use $.machinebox.FaceboxClient.GetState instead.'
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
				responseObjects.push(new $.machinebox.remototypes.FileResponse(response))
			})
			return responseObjects
		})
	}
	// PutState sets the Facebox state file.

	// Must pass in an instance of $.machinebox.PutStateRequest.
	$.machinebox.FaceboxClient.prototype.PutState = function(putStateRequest) {
		return this.PutStateMulti([putStateRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.PutStateMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.PutState.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.PutStateMulti = function(putStateRequests) {
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
				responseObjects.push(new $.machinebox.PutStateResponse(response))
			})
			return responseObjects
		})
	}
	// RemoveID removes a face with the specified ID.

	// Must pass in an instance of $.machinebox.RemoveIDRequest.
	$.machinebox.FaceboxClient.prototype.RemoveID = function(removeIDRequest) {
		return this.RemoveIDMulti([removeIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.RemoveIDMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.RemoveID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.RemoveIDMulti = function(removeIDRequests) {
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
				responseObjects.push(new $.machinebox.RemoveIDResponse(response))
			})
			return responseObjects
		})
	}
	// Rename changes a person&#39;s name.

	// Must pass in an instance of $.machinebox.RenameRequest.
	$.machinebox.FaceboxClient.prototype.Rename = function(renameRequest) {
		return this.RenameMulti([renameRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.RenameMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.Rename.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.RenameMulti = function(renameRequests) {
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
				responseObjects.push(new $.machinebox.RenameResponse(response))
			})
			return responseObjects
		})
	}
	// RenameID changes the name of a previously taught face, by ID.

	// Must pass in an instance of $.machinebox.RenameIDRequest.
	$.machinebox.FaceboxClient.prototype.RenameID = function(renameIDRequest) {
		return this.RenameIDMulti([renameIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.RenameIDMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.RenameID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.RenameIDMulti = function(renameIDRequests) {
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
				responseObjects.push(new $.machinebox.RenameIDResponse(response))
			})
			return responseObjects
		})
	}
	// SimilarFile checks for similar faces from the face in an image file.

	// Must pass in an instance of $.machinebox.SimilarFileRequest.
	$.machinebox.FaceboxClient.prototype.SimilarFile = function(similarFileRequest) {
		return this.SimilarFileMulti([similarFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.SimilarFileMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.SimilarFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.SimilarFileMulti = function(similarFileRequests) {
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
				responseObjects.push(new $.machinebox.SimilarFileResponse(response))
			})
			return responseObjects
		})
	}
	// SimilarID checks for similar faces by ID.

	// Must pass in an instance of $.machinebox.SimilarIDRequest.
	$.machinebox.FaceboxClient.prototype.SimilarID = function(similarIDRequest) {
		return this.SimilarIDMulti([similarIDRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.SimilarIDMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.SimilarID.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.SimilarIDMulti = function(similarIDRequests) {
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
				responseObjects.push(new $.machinebox.SimilarIDResponse(response))
			})
			return responseObjects
		})
	}
	// SimilarURL checks for similar faces in a hosted image file.

	// Must pass in an instance of $.machinebox.SimilarURLRequest.
	$.machinebox.FaceboxClient.prototype.SimilarURL = function(similarURLRequest) {
		return this.SimilarURLMulti([similarURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.SimilarURLMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.SimilarURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.SimilarURLMulti = function(similarURLRequests) {
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
				responseObjects.push(new $.machinebox.SimilarURLResponse(response))
			})
			return responseObjects
		})
	}
	// TeachFaceprint teaches Facebox about a face from a Faceprint.

	// Must pass in an instance of $.machinebox.TeachFaceprintRequest.
	$.machinebox.FaceboxClient.prototype.TeachFaceprint = function(teachFaceprintRequest) {
		return this.TeachFaceprintMulti([teachFaceprintRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.TeachFaceprintMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.TeachFaceprint.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.TeachFaceprintMulti = function(teachFaceprintRequests) {
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
				responseObjects.push(new $.machinebox.TeachFaceprintResponse(response))
			})
			return responseObjects
		})
	}
	// TeachFile teaches Facebox a new face from an image file.

	// Must pass in an instance of $.machinebox.TeachFileRequest.
	$.machinebox.FaceboxClient.prototype.TeachFile = function(teachFileRequest) {
		return this.TeachFileMulti([teachFileRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.TeachFileMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.TeachFile.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.TeachFileMulti = function(teachFileRequests) {
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
				responseObjects.push(new $.machinebox.TeachFileResponse(response))
			})
			return responseObjects
		})
	}
	// TeachURL teaches Facebox a new face from an image on the web.

	// Must pass in an instance of $.machinebox.TeachURLRequest.
	$.machinebox.FaceboxClient.prototype.TeachURL = function(teachURLRequest) {
		return this.TeachURLMulti([teachURLRequest])
			.then(function(responses){
				return responses[0]
			})
	}
	
	// $.machinebox.FaceboxClient.prototype.TeachURLMulti is a batch
	// version of $.machinebox.FaceboxClient.prototype.TeachURL.
	// Pass in an array of request objects, and get back an array of response objects.
	$.machinebox.FaceboxClient.prototype.TeachURLMulti = function(teachURLRequests) {
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
				responseObjects.push(new $.machinebox.TeachURLResponse(response))
			})
			return responseObjects
		})
	}
	
	
	// CheckFaceprintRequest is the request object for CheckFaceprint calls.
$.machinebox.CheckFaceprintRequest = function(data) {
		this.data = data
	}
	$.machinebox.CheckFaceprintRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.CheckFaceprintRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getFaceprints gets the faceprints from this object.
	$.machinebox.CheckFaceprintRequest.prototype.getFaceprints = function() {
		return this.data.faceprints
	}
	
	
	// setFaceprints sets the faceprints on this object.
	$.machinebox.CheckFaceprintRequest.prototype.setFaceprints = function(faceprints) {
		this.data.faceprints = faceprints
	}
	
	// CheckFaceprintResponse is the response object for CheckFaceprint calls.
$.machinebox.CheckFaceprintResponse = function(data) {
		this.data = data
	}
	$.machinebox.CheckFaceprintResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.machinebox.CheckFaceprintResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.machinebox.CheckFaceprintResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// CheckFileRequest is the request object for CheckFile calls.
$.machinebox.CheckFileRequest = function(data) {
		this.data = data
	}
	$.machinebox.CheckFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.CheckFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getFile gets the file from this object.
	$.machinebox.CheckFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.machinebox.CheckFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// CheckFileResponse is the response object for CheckFile calls.
$.machinebox.CheckFileResponse = function(data) {
		this.data = data
	}
	$.machinebox.CheckFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.machinebox.CheckFileResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.machinebox.CheckFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// CheckURLRequest is the request object for CheckURL calls.
$.machinebox.CheckURLRequest = function(data) {
		this.data = data
	}
	$.machinebox.CheckURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.CheckURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getURL gets the uRL from this object.
	$.machinebox.CheckURLRequest.prototype.getURL = function() {
		return this.data.uRL
	}
	
	
	// setURL sets the uRL on this object.
	$.machinebox.CheckURLRequest.prototype.setURL = function(uRL) {
		this.data.uRL = uRL
	}
	
	// CheckURLResponse is the response object for CheckURL calls.
$.machinebox.CheckURLResponse = function(data) {
		this.data = data
	}
	$.machinebox.CheckURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.machinebox.CheckURLResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.machinebox.CheckURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// Face describes a face.
$.machinebox.Face = function(data) {
		this.data = data
	}
	$.machinebox.Face.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getID gets the iD from this object.
	$.machinebox.Face.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.machinebox.Face.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.machinebox.Face.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.machinebox.Face.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getMatched gets the matched from this object.
	$.machinebox.Face.prototype.getMatched = function() {
		return this.data.matched
	}
	
	
	// setMatched sets the matched on this object.
	$.machinebox.Face.prototype.setMatched = function(matched) {
		this.data.matched = matched
	}
	
	// getFaceprint gets the faceprint from this object.
	$.machinebox.Face.prototype.getFaceprint = function() {
		return this.data.faceprint
	}
	
	
	// setFaceprint sets the faceprint on this object.
	$.machinebox.Face.prototype.setFaceprint = function(faceprint) {
		this.data.faceprint = faceprint
	}
	
	// getRect gets the rect from this object.
	$.machinebox.Face.prototype.getRect = function() {
		return this.data.rect
	}
	
	
	// setRect sets the rect on this object.
	$.machinebox.Face.prototype.setRect = function(rect) {
		this.data.rect = rect
	}
	
	// FaceprintCompareRequest is the request object for FaceprintCompare calls.
$.machinebox.FaceprintCompareRequest = function(data) {
		this.data = data
	}
	$.machinebox.FaceprintCompareRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.FaceprintCompareRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getTarget gets the target from this object.
	$.machinebox.FaceprintCompareRequest.prototype.getTarget = function() {
		return this.data.target
	}
	
	
	// setTarget sets the target on this object.
	$.machinebox.FaceprintCompareRequest.prototype.setTarget = function(target) {
		this.data.target = target
	}
	
	// getFaceprints gets the faceprints from this object.
	$.machinebox.FaceprintCompareRequest.prototype.getFaceprints = function() {
		return this.data.faceprints
	}
	
	
	// setFaceprints sets the faceprints on this object.
	$.machinebox.FaceprintCompareRequest.prototype.setFaceprints = function(faceprints) {
		this.data.faceprints = faceprints
	}
	
	// FaceprintCompareResponse is the response object for FaceprintCompare calls.
$.machinebox.FaceprintCompareResponse = function(data) {
		this.data = data
	}
	$.machinebox.FaceprintCompareResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getConfidences gets the confidences from this object.
	$.machinebox.FaceprintCompareResponse.prototype.getConfidences = function() {
		return this.data.confidences
	}
	
	
	// getError gets the error from this object.
	$.machinebox.FaceprintCompareResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// FaceprintFace is a face.
$.machinebox.FaceprintFace = function(data) {
		this.data = data
	}
	$.machinebox.FaceprintFace.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getMatched gets the matched from this object.
	$.machinebox.FaceprintFace.prototype.getMatched = function() {
		return this.data.matched
	}
	
	
	// setMatched sets the matched on this object.
	$.machinebox.FaceprintFace.prototype.setMatched = function(matched) {
		this.data.matched = matched
	}
	
	// getConfidence gets the confidence from this object.
	$.machinebox.FaceprintFace.prototype.getConfidence = function() {
		return this.data.confidence
	}
	
	
	// setConfidence sets the confidence on this object.
	$.machinebox.FaceprintFace.prototype.setConfidence = function(confidence) {
		this.data.confidence = confidence
	}
	
	// getID gets the iD from this object.
	$.machinebox.FaceprintFace.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.machinebox.FaceprintFace.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.machinebox.FaceprintFace.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.machinebox.FaceprintFace.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// GetStateRequest is the request object for GetState calls.
$.machinebox.GetStateRequest = function(data) {
		this.data = data
	}
	$.machinebox.GetStateRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.GetStateRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// PutStateRequest is the request object for PutState calls.
$.machinebox.PutStateRequest = function(data) {
		this.data = data
	}
	$.machinebox.PutStateRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.PutStateRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getStateFile gets the stateFile from this object.
	$.machinebox.PutStateRequest.prototype.getStateFile = function() {
		return this.data.stateFile
	}
	
	
	// setStateFile sets the stateFile on this object.
	// The root request must also be provided so it can be informed of the file.
	$.machinebox.PutStateRequest.prototype.setStateFile = function(request, stateFile) {
		this.data.stateFile = request._addFile(stateFile)
	}
	
	
	// PutStateResponse is the response object for PutState calls.
$.machinebox.PutStateResponse = function(data) {
		this.data = data
	}
	$.machinebox.PutStateResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.machinebox.PutStateResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// Rect is a bounding box describing a rectangle of an image.
$.machinebox.Rect = function(data) {
		this.data = data
	}
	$.machinebox.Rect.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getTop gets the top from this object.
	$.machinebox.Rect.prototype.getTop = function() {
		return this.data.top
	}
	
	
	// setTop sets the top on this object.
	$.machinebox.Rect.prototype.setTop = function(top) {
		this.data.top = top
	}
	
	// getLeft gets the left from this object.
	$.machinebox.Rect.prototype.getLeft = function() {
		return this.data.left
	}
	
	
	// setLeft sets the left on this object.
	$.machinebox.Rect.prototype.setLeft = function(left) {
		this.data.left = left
	}
	
	// getWidth gets the width from this object.
	$.machinebox.Rect.prototype.getWidth = function() {
		return this.data.width
	}
	
	
	// setWidth sets the width on this object.
	$.machinebox.Rect.prototype.setWidth = function(width) {
		this.data.width = width
	}
	
	// getHeight gets the height from this object.
	$.machinebox.Rect.prototype.getHeight = function() {
		return this.data.height
	}
	
	
	// setHeight sets the height on this object.
	$.machinebox.Rect.prototype.setHeight = function(height) {
		this.data.height = height
	}
	
	// RemoveIDRequest is the request object for RemoveID calls.
$.machinebox.RemoveIDRequest = function(data) {
		this.data = data
	}
	$.machinebox.RemoveIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.RemoveIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.machinebox.RemoveIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.machinebox.RemoveIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// RemoveIDResponse is the response object for RemoveID calls.
$.machinebox.RemoveIDResponse = function(data) {
		this.data = data
	}
	$.machinebox.RemoveIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.machinebox.RemoveIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// RenameIDRequest is the request object for RenameID calls.
$.machinebox.RenameIDRequest = function(data) {
		this.data = data
	}
	$.machinebox.RenameIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.RenameIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.machinebox.RenameIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.machinebox.RenameIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.machinebox.RenameIDRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.machinebox.RenameIDRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// RenameIDResponse is the response object for RenameID calls.
$.machinebox.RenameIDResponse = function(data) {
		this.data = data
	}
	$.machinebox.RenameIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.machinebox.RenameIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// RenameRequest is the request object for Rename calls.
$.machinebox.RenameRequest = function(data) {
		this.data = data
	}
	$.machinebox.RenameRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.RenameRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getFrom gets the from from this object.
	$.machinebox.RenameRequest.prototype.getFrom = function() {
		return this.data.from
	}
	
	
	// setFrom sets the from on this object.
	$.machinebox.RenameRequest.prototype.setFrom = function(from) {
		this.data.from = from
	}
	
	// getTo gets the to from this object.
	$.machinebox.RenameRequest.prototype.getTo = function() {
		return this.data.to
	}
	
	
	// setTo sets the to on this object.
	$.machinebox.RenameRequest.prototype.setTo = function(to) {
		this.data.to = to
	}
	
	// RenameResponse is the response object for Rename calls.
$.machinebox.RenameResponse = function(data) {
		this.data = data
	}
	$.machinebox.RenameResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.machinebox.RenameResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarFace is a detected face with similar matching faces.
$.machinebox.SimilarFace = function(data) {
		this.data = data
	}
	$.machinebox.SimilarFace.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getRect gets the rect from this object.
	$.machinebox.SimilarFace.prototype.getRect = function() {
		return this.data.rect
	}
	
	
	// setRect sets the rect on this object.
	$.machinebox.SimilarFace.prototype.setRect = function(rect) {
		this.data.rect = rect
	}
	
	// getSimilarFaces gets the similarFaces from this object.
	$.machinebox.SimilarFace.prototype.getSimilarFaces = function() {
		return this.data.similarFaces
	}
	
	
	// setSimilarFaces sets the similarFaces on this object.
	$.machinebox.SimilarFace.prototype.setSimilarFaces = function(similarFaces) {
		this.data.similarFaces = similarFaces
	}
	
	// SimilarFileRequest is the request object for SimilarFile calls.
$.machinebox.SimilarFileRequest = function(data) {
		this.data = data
	}
	$.machinebox.SimilarFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.SimilarFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getFile gets the file from this object.
	$.machinebox.SimilarFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.machinebox.SimilarFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// SimilarFileResponse is the response object for SimilarFile calls.
$.machinebox.SimilarFileResponse = function(data) {
		this.data = data
	}
	$.machinebox.SimilarFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.machinebox.SimilarFileResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.machinebox.SimilarFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarIDRequest is the request object for SimilarID calls.
$.machinebox.SimilarIDRequest = function(data) {
		this.data = data
	}
	$.machinebox.SimilarIDRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.SimilarIDRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.machinebox.SimilarIDRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.machinebox.SimilarIDRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// SimilarIDResponse is the response object for SimilarID calls.
$.machinebox.SimilarIDResponse = function(data) {
		this.data = data
	}
	$.machinebox.SimilarIDResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.machinebox.SimilarIDResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.machinebox.SimilarIDResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// SimilarURLRequest is the request object for SimilarURL calls.
$.machinebox.SimilarURLRequest = function(data) {
		this.data = data
	}
	$.machinebox.SimilarURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.SimilarURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getURL gets the uRL from this object.
	$.machinebox.SimilarURLRequest.prototype.getURL = function() {
		return this.data.uRL
	}
	
	
	// setURL sets the uRL on this object.
	$.machinebox.SimilarURLRequest.prototype.setURL = function(uRL) {
		this.data.uRL = uRL
	}
	
	// SimilarURLResponse is the response object for SimilarURL calls.
$.machinebox.SimilarURLResponse = function(data) {
		this.data = data
	}
	$.machinebox.SimilarURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getFaces gets the faces from this object.
	$.machinebox.SimilarURLResponse.prototype.getFaces = function() {
		return this.data.faces
	}
	
	
	// getError gets the error from this object.
	$.machinebox.SimilarURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachFaceprintRequest is the request object for TeachFaceprint calls.
$.machinebox.TeachFaceprintRequest = function(data) {
		this.data = data
	}
	$.machinebox.TeachFaceprintRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.TeachFaceprintRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.machinebox.TeachFaceprintRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.machinebox.TeachFaceprintRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.machinebox.TeachFaceprintRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.machinebox.TeachFaceprintRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getFaceprint gets the faceprint from this object.
	$.machinebox.TeachFaceprintRequest.prototype.getFaceprint = function() {
		return this.data.faceprint
	}
	
	
	// setFaceprint sets the faceprint on this object.
	$.machinebox.TeachFaceprintRequest.prototype.setFaceprint = function(faceprint) {
		this.data.faceprint = faceprint
	}
	
	// TeachFaceprintResponse is the response object for TeachFaceprint calls.
$.machinebox.TeachFaceprintResponse = function(data) {
		this.data = data
	}
	$.machinebox.TeachFaceprintResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.machinebox.TeachFaceprintResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachFileRequest is the request object for TeachFile calls.
$.machinebox.TeachFileRequest = function(data) {
		this.data = data
	}
	$.machinebox.TeachFileRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.TeachFileRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.machinebox.TeachFileRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.machinebox.TeachFileRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.machinebox.TeachFileRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.machinebox.TeachFileRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getFile gets the file from this object.
	$.machinebox.TeachFileRequest.prototype.getFile = function() {
		return this.data.file
	}
	
	
	// setFile sets the file on this object.
	// The root request must also be provided so it can be informed of the file.
	$.machinebox.TeachFileRequest.prototype.setFile = function(request, file) {
		this.data.file = request._addFile(file)
	}
	
	
	// TeachFileResponse is the response object for TeachFile calls.
$.machinebox.TeachFileResponse = function(data) {
		this.data = data
	}
	$.machinebox.TeachFileResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.machinebox.TeachFileResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// TeachURLRequest is the request object for TeachURL calls.
$.machinebox.TeachURLRequest = function(data) {
		this.data = data
	}
	$.machinebox.TeachURLRequest.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	// _addFile adds a file to the request.
	$.machinebox.TeachURLRequest.prototype._addFile = function(file) {
		this._files = this._files || {}
		var fieldname = 'files['+$.machinebox._filesCount+']'
		this._files[fieldname] = file
		$.machinebox._filesCount++
		return fieldname
	}
	
	
	// getID gets the iD from this object.
	$.machinebox.TeachURLRequest.prototype.getID = function() {
		return this.data.iD
	}
	
	
	// setID sets the iD on this object.
	$.machinebox.TeachURLRequest.prototype.setID = function(iD) {
		this.data.iD = iD
	}
	
	// getName gets the name from this object.
	$.machinebox.TeachURLRequest.prototype.getName = function() {
		return this.data.name
	}
	
	
	// setName sets the name on this object.
	$.machinebox.TeachURLRequest.prototype.setName = function(name) {
		this.data.name = name
	}
	
	// getURL gets the uRL from this object.
	$.machinebox.TeachURLRequest.prototype.getURL = function() {
		return this.data.uRL
	}
	
	
	// setURL sets the uRL on this object.
	$.machinebox.TeachURLRequest.prototype.setURL = function(uRL) {
		this.data.uRL = uRL
	}
	
	// TeachURLResponse is the response object for TeachURL calls.
$.machinebox.TeachURLResponse = function(data) {
		this.data = data
	}
	$.machinebox.TeachURLResponse.prototype.toJSON = function() {
		return JSON.stringify(this.data)
	}
	
	
	// getError gets the error from this object.
	$.machinebox.TeachURLResponse.prototype.getError = function() {
		return this.data.error
	}
	
	
	// _filesCount keeps track of the number of files being added, and is used
	// to generate unique field names.
	$.machinebox._filesCount = 0

})(jQuery)
