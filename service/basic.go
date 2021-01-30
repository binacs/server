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
			<h5>一些代替了BLOG存在的小玩意儿</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">solutions - cryptology - pastebin - tinyurl - prometheus - grafana - jenkins - doscify</div>
				</div>
			</div>
			<div style="font-size: 14px; height: 40vh;">
				<a href="https://github.com/BinacsLee/BinacsLee.github.io/tree/master/article" target="_blank">早期题解</a>
				<br>
				<a href="toys/crypto" target="_blank">基于密码学加解密(k8s service)</a>
				<br>
				<a href="toys/pastebin" target="_blank">语法高亮的粘贴板 Pastebin</a>
				<br>
				<a href="toys/tinyurl" target="_blank">短链接服务 TinyURL</a>
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
			<h5>基于密码学技术的在线加解密</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">密钥输入种类过于复杂 故暂使用服务器保留密钥 并随机使用Base64/AES/DES加解密</div>
				</div>
			</div>
			<div style="font-size: 12px;">
				<form id="form" class="row">
					<div class="input-field col s12">
						<textarea id="enc_source" class="materialize-textarea"></textarea>
						<label for="enc_source">明文</label>
					</div>	

					<div class="col s12" style="font-size: 20px;">
						<button class="white" type="button" onclick="enc()" value="enc" id="encid">加密</button>
						<button class="white" type="button" onclick="dec()" value="dec" id="decid">解密</button>
					</div>

					<div class="input-field col s12">
						<textarea id="dec_source" class="materialize-textarea"></textarea>
						<label for="dec_source">密文</label>
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
								alert("异常！");
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
								alert("异常！");p_
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
			<h5>短链接生成服务</h5>
			<div class="section" style="height: 10vh;">
				<div class="row">
					<div class="col s12 grey-text text-darken-1">默认使用md5算法</div>
				</div>
			</div>
			<div style="font-size: 12px;">
				<div class="row">
					<form class="col s12">
						<div class="row">
							<div class="input-field col s12">
								<input placeholder="https / http ..." id="enc_source" type="text" class="validate">
								<label for="enc_source">原始URL</label>
							</div>
						</div>
						<div class="row" style="font-size: 20px;">
							<div class="col s12">
								<button class="white" type="button" onclick="enc()" value="enc" id="encid">编码</button>
								<button class="white" type="button" onclick="dec()" value="dec" id="decid">解码</button>
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
								alert("异常！");
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
								alert("异常！");
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
					<div class="col s12 grey-text text-darken-1">默认使用Markdown语法渲染</div>
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
								<input id="sub_syntax" type="text" id="sub_syntax">
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
								<button class="white" type="button" onclick="paste_submit()" value="submit" id="submitid">提交</button>
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
								alert("异常！");
							}
						})
					}
				</script>
			</div>
		</div>
	</div>
`
)
