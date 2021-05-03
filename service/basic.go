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
					<a style="margin: 10px;" href="toys" target="_blank">TOYS</a>
					<a style="margin: 10px;" href="https://github.com/BinacsLee" target="_blank">GITHUB</a>
					<a style="margin: 10px;" href="mailto:binacs055@vip.qq.om">EMAIL</a>
					<a style="margin: 10px;" href="about" target="_blank">ABOUT</a>
				</h6>
			</div>
		</div>
	</div>
`
	aboutBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h5>关于我</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">whu - cyber security - algorithm - go - cpp - cloud native - blockchain</div>
				</div>
			</div>
			<div style="font-size: 14px; height: 40vh;">
				<a style="color:black">武汉大学 - 国家网络安全学院  2015 - 2019</a>
				<br>
				<a style="color:black">懂点安全   学过算法</a>
				<br>
				<a style="color:black"> Go && C++</a>
				<br>
				<a style="color:black">云原生 && 区块链</a>
				<br>
				<a style="color:black">路漫漫其修远兮</a>
			</div>
		</div>
	</div>
`
	toysBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h5>TOYS</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">solutions - cryptology - pastebin - tinyurl - prometheus - grafana - jenkins - doscify</div>
				</div>
			</div>
			<div style="font-size: 14px; height: 40vh;">
				<a href="toys/crypto" target="_blank">Crypto</a>
				<br>
				<a href="toys/pastebin" target="_blank">Pastebin</a>
				<br>
				<a href="toys/tinyurl" target="_blank">TinyURL</a>
				<br>
				<a href="toys/storage" target="_blank">Storage</a>
				<br>
				<a href="https://prometheus.binacs.cn" target="_blank">Prometheus</a>
				<br>
				<a href="https://grafana.binacs.cn" target="_blank">Grafana</a>
				<br>
				<a href="https://jenkins.binacs.cn" target="_blank">Jenkins</a>
				<br>
				<a href="https://docs.binacs.cn" target="_blank">Docs</a>
				<br>
			</div>
		</div>
	</div>
`
	cryptoBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h5>Online encryption and decryption based on cryptography technology</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">
						<a>
						The type of key input is too complicated, so use the server to reserve the key temporarily and<br>
						 use Base64/AES/DES for encryption and decryption on the web side
						</a>
						<a href="https://github.com/BinacsLee/cli">
						For more functions, please use the dedicated client
						</a>
					</div>
				</div>
			</div>
			<div style="font-size: 12px;">
				<form id="form" class="row">
					<div class="input-field col s12">
						<textarea id="enc_source" class="materialize-textarea"></textarea>
						<label for="enc_source">Plaintext</label>
					</div>

					<div class="col s12" style="font-size: 20px;">
						<button class="white" type="button" onclick="enc()" value="enc" id="encid">Encrypto</button>
						<button class="white" type="button" onclick="dec()" value="dec" id="decid">Decrypto</button>
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
`
	tinyurlBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h5>Tiny link generation service</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">The md5 algorithm is used by default</div>
				</div>
			</div>
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
								<button class="white" type="button" onclick="enc()" value="enc" id="encid">Encode</button>
								<button class="white" type="button" onclick="dec()" value="dec" id="decid">Decode</button>
							</div>
						</div>
						<div class="row">
							<div class="input-field col s12">
								<input placeholder="binacs.cn/r/ ..." id="dec_source" type="text" class="validate">
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
`
	pastebinBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h5>Pastebin</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">Render using Markdown syntax by default</div>
				</div>
			</div>
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
								<button class="white" type="button" onclick="paste_submit()" value="submit" id="submitid">Submit</button>
							</div>
						</div>
					</form>
				</div>

				<script type="text/javascript">
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
`
	storageBody = `
	<div class="container valign-wrapper" style="height: 90vh;">
		<div class="section center valign" style="width: 100%;">
			<h5>Storage</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">COS Storage</div>
				</div>
			</div>
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
							<button class="white" type="button" onclick="cos_put()" value="submit" id="submitid">Submit</button>
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
`
)
