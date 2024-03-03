package service

// BasicServiceImpl Web static service implement
type BasicServiceImpl struct{}

// AfterInject do inject
func (bs *BasicServiceImpl) AfterInject() error {
	return nil
}

// ServeHome return the `home`
func (bs *BasicServiceImpl) ServeHome() string {
	return homeBody
}

// ServeToys return the `toys`
func (bs *BasicServiceImpl) ServeToys() string {
	return toysBody
}

// ServeCrypto return the `toys/crypto`
func (bs *BasicServiceImpl) ServeCrypto() string {
	return cryptoBody
}

// ServeTinyURL return the `toys/tinyurl`
func (bs *BasicServiceImpl) ServeTinyURL() string {
	return tinyurlBody
}

// ServePastebin return the `toys/pastebin`
func (bs *BasicServiceImpl) ServePastebin() string {
	return pastebinBody
}

// ServeStorage return the `toys/storage`
func (bs *BasicServiceImpl) ServeStorage() string {
	return storageBody
}

// ServeAbout return the `about`
func (bs *BasicServiceImpl) ServeAbout() string {
	return aboutBody
}

const (
	homeBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<div class="section center valign" style="width: 100%;">
				<h6>
					<a style="margin: 16px;" href="toys" target="_blank">TOYS</a>
					<a style="margin: 16px;" href="blogs" target="_blank">BLOGS</a>
					<a style="margin: 16px;" href="https://github.com/binacs" target="_blank">GITHUB</a>
					<a style="margin: 16px;" href="mailto:binacs055@vip.qq.om">EMAIL</a>
					<a style="margin: 16px;" href="about" target="_blank">ABOUT</a>
				</h6>
			</div>
		</div>
	</div>
`
	aboutBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h1>ABOUT</h1>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">Blockchain and Cloud Native</div>
				</div>
			</div>
			<div style="font-size: 14px; height: 40vh;">
				<div class="card grey lighten-5">
					<div class="card-content">
						<span class="card-title"> </span>
						<a style="color:black">You can download my resume on</a>
						<a href="https://github.com/binacs/resume/releases/tag/latest">GitHub Release</a>
						<br>
						<a style="color:black">OR</a>
						<br>
						<a style="color:black">preview</a>
						<a href="https://docs.google.com/viewer?url=https://github.com/binacs/resume/releases/download/latest/resume.pdf">resume.pdf</a>
						<a href="https://docs.google.com/viewer?url=https://github.com/binacs/resume/releases/download/latest/resume_zh.pdf">resume_zh.pdf</a>
						<a style="color:black">online by Google Docs</a>
						<br><br><br><br>
						<a style="color:black">路漫漫其修远兮</a>
					</div>
				</div>
			</div>
		</div>
	</div>
`
	toysBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h1>TOYS</h1>
			<div class="section" style="height: 10vh;">
				<div class="row">
				</div>
			</div>
			<div class="container">
				<div class="row">
					<a class="col s8 offset-s2 waves-effect waves-teal btn-large grey lighten-5" style="height: 5vh;" href="toys/crypto" target="_blank">Crypto</a>
				</div>
				<div class="row">
					<a class="col s8 offset-s2 waves-effect waves-teal btn-large grey lighten-5" style="height: 5vh;" href="toys/pastebin" target="_blank">Pastebin</a>
				</div>
				<div class="row">
					<a class="col s8 offset-s2 waves-effect waves-teal btn-large grey lighten-5" style="height: 5vh;" href="toys/tinyurl" target="_blank">TinyURL</a>
				</div>
				<div class="row">
					<a class="col s8 offset-s2 waves-effect waves-teal btn-large grey lighten-5" style="height: 5vh;" href="toys/storage" target="_blank">Storage</a>
				</div>
				<div class="row">
					<a class="col s8 offset-s2 waves-effect waves-teal btn-large grey lighten-5" style="height: 5vh;" href="https://docs.binacs.space" target="_blank">Docs</a>
				</div>
			</div>
		</div>
	</div>
`
	cryptoBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h1>Crypto</h1>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">
						The type of key input is too complicated, so use the server to reserve the key temporarily and use Base64/AES/DES<br>
						 for encryption and decryption on the web side.
						<a href="https://github.com/binacs/cli">
						<u>For more functions, please use the dedicated client</u>
						</a>
					</div>
				</div>
			</div>
			<div class="container">
				<div style="font-size: 12px;">
					<form id="form" class="row">
						<div class="input-field col s12">
							<textarea id="enc_source" class="materialize-textarea"></textarea>
							<label for="enc_source">Plaintext</label>
						</div>

						<div class="col s12" style="font-size: 20px;">
							<button class="btn waves-effect waves-light" type="button" onclick="enc()" value="enc" id="encid">Encrypto
								<i class="material-icons right">arrow_downward</i>
							</button>
							<button class="btn waves-effect waves-light" type="button" onclick="dec()" value="dec" id="decid">Decrypto
								<i class="material-icons right">arrow_upward</i>
							</button>
						</div>

						<div class="input-field col s12">
							<textarea id="dec_source" class="materialize-textarea"></textarea>
							<label for="dec_source">Ciphertext</label>
						</div>
					</form>

					<script type="text/javascript">
						function enc() {
							$.ajax({
								'url': '/api/crypto/encrypto',
								'data': {
									"text": $('#enc_source').val(),
									"type": "AES"
								},
								'type': 'post',
								'dataType': "text",
								'success': function (data) {
									console.log("enc success")
									$('#dec_source').val(data);
								},
								error: function () {
									alert("error!");
								}
							})
						}

						function dec() {
							$.ajax({
								'url': '/api/crypto/decrypto',
								'data': {
									"text": $('#dec_source').val(),
									"type": "AES"
								},
								'type': 'post',
								'dataType': "text",
								'success': function (data) {
									console.log("dec success")
									$('#enc_source').val(data);
								},
								error: function () {
									alert("error!");p_
								}
							})
						}
					</script>
				</div>
			</div>
		</div>
	</div>
`
	tinyurlBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h1>TinyURL</h1>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">The md5 algorithm is used by default</div>
				</div>
			</div>
			<div class="container">
				<div style="font-size: 12px;">
					<div class="row">
						<form class="col s12">
							<div class="row">
								<div class="input-field col s12">
									<input placeholder="https / http ..." id="enc_source" type="text" class="validate">
									<label for="enc_source">OriginURL</label>
								</div>
							</div>
							<div class="row" style="font-size: 20px;">
								<div class="col s12">
									<button class="btn waves-effect waves-light" type="button" onclick="enc()" value="enc" id="encid">Encode
										<i class="material-icons right">arrow_downward</i>
									</button>
									<button class="btn waves-effect waves-light" type="button" onclick="dec()" value="dec" id="decid">Decode
										<i class="material-icons right">arrow_upward</i>
									</button>
								</div>
							</div>
							<div class="row">
								<div class="input-field col s12">
									<input placeholder="binacs.space/r/ ..." id="dec_source" type="text" class="validate">
									<label for="dec_source">TinyURL</label>
								</div>
							</div>
						</form>
					</div>

					<script type="text/javascript">
						function enc() {
							$.ajax({
								'url': '/api/tinyurl/encode',
								'data': {
									"text": $('#enc_source').val(),
								},
								'type': 'post',
								'dataType': "text",
								'success': function(data) {
									console.log("enc success")
									$('#dec_source').val(data);
								},
								error: function() {
									alert("error!");
								}
							})
						}
						
						function dec() {
							$.ajax({
								'url': '/api/tinyurl/decode',
								'data': {
									"text": $('#dec_source').val(),
								},
								'type': 'post',
								'dataType': "text",
								'success': function(data) {
									console.log("dec success")
									$('#enc_source').val(data);
								},
								error: function() {
									alert("error!");
								}
							})
						}
					</script>
				</div>
			</div>
		</div>
	</div>
`
	pastebinBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h1>Pastebin</h1>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">Render using Markdown syntax by default</div>
				</div>
			</div>
			<div class="container">
				<div style="font-size: 12px;">
					<div class="row">
						<form class="col s12" name="pasteForm">
							<div class="row">
								<div class="input-field col s6">
									<input id="sub_poster" type="text" id="sub_poster">
									<label for="sub_poster">Poster</label>
								</div>
								<div class="input-field col s6">
									<input id="sub_syntax" type="text">
									<label for="sub_syntax">Syntax</label>
								</div>
							</div>
							<div class="row">
								<div class="input-field col s12">
									<textarea id="sub_content" class="materialize-textarea"></textarea>
									<label for="sub_content">Content</label>
								</div>
							</div>
							<div class="row" style="font-size: 20px;">
								<div class="col s12">
									<button class="btn waves-effect waves-light" type="button" onclick="paste_submit()" value="submit" id="submitid">Submit
										<i class="material-icons right">send</i>
									</button>
								</div>
							</div>
						</form>
					</div>

					<script type="text/javascript">
						M.textareaAutoResize($('#sub_content'));
						function paste_submit() {
							$.ajax({
								'url': '/api/pastebin/submit',
								'data': {
									"poster": $('#sub_poster').val(),
									"syntax": $('#sub_syntax').val(),
									"content": $('#sub_content').val(),
								},
								'type': 'post',
								'dataType': "text",
								'success': function(data) {
									console.log("submit success")
									console.log(data)
									window.location.href=data;
								},
								error: function() {
									alert("error!");
								}
							})
						}
					</script>
				</div>
			</div>
		</div>
	</div>
`
	storageBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h1>Storage</h1>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">COS Storage</div>
				</div>
			</div>
			<div class="container">
				<div style="font-size: 12px;">
					<div class="row">
						<form class="col s12" name="cosForm" enctype="multipart/form-data">
							<div class="file-field input-field">
								<div class="btn">
									<span>File</span>
									<input type="file" id="file" multiple="">
								</div>
								<div class="file-path-wrapper">
									<input class="file-path validate" type="text" placeholder="Upload one or more files">
								</div>
							</div>
							<div class="input-field">
								<input id="pass_key" type="text">
								<label for="pass_key">PassKey</label>
							</div>
							<div class="row" style="font-size: 20px;">
								<button class="btn waves-effect waves-light" type="button" onclick="cos_put()" value="submit" id="submitid">Upload
									<i class="material-icons right">cloud_upload</i>
								</button>
							</div>
						</form>
					</div>

					<script type="text/javascript">
						function cos_put() {
							var formData = new FormData();
							console.log($('#file')[0].files.length);
							for (var i = 0; i < $('#file')[0].files.length; i++) {
								formData.append('file', $("#file")[0].files[i]);
							}
							formData.append('key', $('#pass_key').val())

							$.ajax({
								'url': '/api/cos/put',
								'async': false,
								'data': formData,
								'type': 'post',
								'dataType': "text",
								'cache': false,
								'processData': false,
								'contentType': false,
								'success': function(data) {
									console.log("put success")
									console.log(data)
									alert(data);
								},
								error: function() {
									alert("error!");
								}
							})
						}
					</script>
				</div>
			</div>
		</div>
	</div>
`
)
