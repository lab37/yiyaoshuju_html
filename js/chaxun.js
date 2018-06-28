function checkInput(elementId) {
  var el = document.getElementById(elementId);   // Username input
  var isNull = false;
  var elReal = el.value.replace(/\s+/g,"");
  if (elReal.length < 1) {            // If username too short
   isNull = true;                             // Clear msg
  }
  return isNull;
}


function clickZbButton(){
	if (checkInput("ml_yaopin_name")&&checkInput("ml_qiye_name")) {
		alert('你未输入任何内容，请检查后重新输入');
		return;
	}
    var ypName = document.getElementById("ml_yaopin_name");
    var cjName = document.getElementById("ml_qiye_name");
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		if (xhr.status == 200) {
			  responseObject = JSON.parse(xhr.responseText);
			  var newContent = '';
           	  newContent += '<tr class="table_title">';
              newContent += '<td>通用名</td>';
              newContent += '<td>剂型</td>';
              newContent += '<td>规格</td>';
              newContent += '<td>商品名</td>';
              newContent += '<td>价格</td>';
              newContent += '<td>价格说明</td>';
              newContent += '<td>零售限价</td>';
			  newContent += '<td>包装</td>';
              newContent += '<td>采购类型</td>';
              newContent += '<td>是否基药</td>';
              newContent += '<td>生产企业</td>';
              newContent += '<td>适用</td>';
              newContent += '</tr>';
              for (var i = 0; i < responseObject.length; i++) {    // Loop through object
                newContent += '<tr>';
                newContent += '<td>'+responseObject[i].Tym+'</td>';
                newContent += '<td>'+responseObject[i].Jx+'</td>';
                newContent += '<td>'+responseObject[i].Gg+'</td>';
                newContent += '<td>'+responseObject[i].Spm+'</td>';
                newContent += '<td>'+responseObject[i].Jg+'</td>';
                newContent += '<td>'+responseObject[i].Jgsm+'</td>';
				newContent += '<td>'+responseObject[i].Lsxj+'</td>';
                newContent += '<td>'+responseObject[i].Bz+'</td>';
                newContent += '<td>'+responseObject[i].Cglx+'</td>';
                newContent += '<td>'+responseObject[i].Sfjy+'</td>';
                newContent += '<td>'+responseObject[i].Scqy+'</td>';
                newContent += '<td>'+responseObject[i].Sy+'</td>';
                newContent += '</tr>';
              }
        // Update the page with the new content
              document.getElementById('results').innerHTML = newContent;
		}
	};
	xhr.open('GET', 'chaxun?neirong=zhongbiao&yaopinname=' + encodeURIComponent(ypName.value.replace(/\s+/g,"")) + '&shengchanqiye=' + encodeURIComponent(cjName.value.replace(/\s+/g,"")), true);
	xhr.send(null);
	
}

function clickYbButton(){
	if (checkInput("ml_yaopin_name")) {
		alert('你未输入任何内容，请检查后重新输入');
		return;
	}
    var ypName = document.getElementById("ml_yaopin_name");
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		if (xhr.status == 200) {
			  responseObject = JSON.parse(xhr.responseText);
			  var newContent = '';
              newContent += '<tr class="table_title">';
              newContent += '<td>分类代码</td>';
              newContent += '<td>分类I</td>';
              newContent += '<td>分类III</td>';
              newContent += '<td>医保类型</td>';
              newContent += '<td>编号</td>';
              newContent += '<td>药品名称</td>';
              newContent += '<td>剂型</td>';
              newContent += '<td>备注</td>';
              newContent += '</tr>';
              for (var i = 0; i < responseObject.length; i++) {    // Loop through object
                 newContent += '<tr>';
                 newContent += '<td>'+responseObject[i].Fldm+'</td>';
                 newContent += '<td>'+responseObject[i].Fl1+'</td>';
                 newContent += '<td>'+responseObject[i].Fl3+'</td>';
                 newContent += '<td>'+responseObject[i].Yblx+'</td>';
                 newContent += '<td>'+responseObject[i].Bh+'</td>';
                 newContent += '<td>'+responseObject[i].Ypmc+'</td>';
                 newContent += '<td>'+responseObject[i].Jx+'</td>';
                 newContent += '<td>'+responseObject[i].Bz+'</td>';
                 newContent += '</tr>';
              }
              document.getElementById('results').innerHTML = newContent;
		}
	};
	xhr.open('GET', 'chaxun?neirong=yibao&yaopinname=' + encodeURIComponent(ypName.value.replace(/\s+/g,"")), true);
	xhr.send(null);
	
}

function clickJcbcButton(){
	if (checkInput("ml_yaopin_name")&&checkInput("ml_qiye_name")) {
		alert('你未输入任何内容，请检查后重新输入');
		return;
	}
    var ypName = document.getElementById("ml_yaopin_name");
    var cjName = document.getElementById("ml_qiye_name");
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		if (xhr.status == 200) {
			  responseObject = JSON.parse(xhr.responseText);
			  var newContent = '';
           	  newContent += '<tr class="table_title">';
              newContent += '<td>通用名</td>';
              newContent += '<td>剂型</td>';
              newContent += '<td>规格</td>';
              newContent += '<td>包装</td>';
              newContent += '<td>包装材料</td>';
              newContent += '<td>生产厂家</td>';
              newContent += '<td>价格</td>';
			  newContent += '<td>价格说明</td>';
              newContent += '<td>最低交易价</td>';
              newContent += '<td>平均交易价</td>';
              newContent += '<td>采购类型</td>';
              newContent += '<td>项目名称</td>';
              newContent += '</tr>';
              for (var i = 0; i < responseObject.length; i++) {    // Loop through object
                newContent += '<tr>';
                newContent += '<td>'+responseObject[i].Tym+'</td>';
                newContent += '<td>'+responseObject[i].Jx+'</td>';
                newContent += '<td>'+responseObject[i].Gg+'</td>';
                newContent += '<td>'+responseObject[i].Bz+'</td>';
                newContent += '<td>'+responseObject[i].Bzcl+'</td>';
                newContent += '<td>'+responseObject[i].Sccj+'</td>';
				newContent += '<td>'+responseObject[i].Jg+'</td>';
                newContent += '<td>'+responseObject[i].Jgsm+'</td>';
                newContent += '<td>'+responseObject[i].Zdjyjg+'</td>';
                newContent += '<td>'+responseObject[i].Pjjyjg+'</td>';
                newContent += '<td>'+responseObject[i].Cglx+'</td>';
                newContent += '<td>'+responseObject[i].Xmmc+'</td>';
                newContent += '</tr>';
              }
        // Update the page with the new content
              document.getElementById('results').innerHTML = newContent;
		}
	};
	xhr.open('GET', 'chaxun?neirong=jcbc&yaopinname=' + encodeURIComponent(ypName.value.replace(/\s+/g,"")) + '&shengchanqiye=' + encodeURIComponent(cjName.value.replace(/\s+/g,"")), true);
	xhr.send(null);
	
}

function clickRlButton(){
	if (checkInput("rl_yaopin_name")&&checkInput("rl_qiye_name")&&checkInput("rl_dishi_name")) {
		alert('你未输入任何内容，请检查后重新输入');
		return;
	}
    var ypName = document.getElementById("rl_yaopin_name");
    var cjName = document.getElementById("rl_qiye_name");
	var dsName = document.getElementById("rl_dishi_name");
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		if (xhr.status == 200) {
			  responseObject = JSON.parse(xhr.responseText);
			  var newContent = '';
           	  newContent += '<tr class="table_title">';
              newContent += '<td>通用名</td>';
              newContent += '<td>剂型</td>';
              newContent += '<td>规格</td>';
              newContent += '<td>生产企业</td>';
              newContent += '<td>地市</td>';
			  newContent += '<td>医院</td>';
              newContent += '<td>负责人</td>';
              newContent += '<td>联系方式</td>';
			  newContent += '<td>负责人类型</td>';
			  newContent += '<td>备注</td>';
              newContent += '</tr>';
              for (var i = 0; i < responseObject.length; i++) {    // Loop through object
                newContent += '<tr>';
                newContent += '<td>'+responseObject[i].Tym+'</td>';
                newContent += '<td>'+responseObject[i].Jx+'</td>';
                newContent += '<td>'+responseObject[i].Gg+'</td>';
                newContent += '<td>'+responseObject[i].Scqy+'</td>';
                newContent += '<td>'+responseObject[i].Ds+'</td>';
                newContent += '<td>'+responseObject[i].Yy+'</td>';
				newContent += '<td>'+responseObject[i].Fzr+'</td>';
                newContent += '<td>'+responseObject[i].Lxfs+'</td>';
                newContent += '<td>'+responseObject[i].Rylx+'</td>';
                newContent += '<td>'+responseObject[i].Bz+'</td>';
                newContent += '</tr>';
              }
        // Update the page with the new content
             document.getElementById('results').innerHTML = newContent;
		}
	};
	xhr.open('GET', 'chaxun?neirong=rlzy&yaopinname=' + encodeURIComponent(ypName.value.replace(/\s+/g,"")) + '&shengchanqiye=' + encodeURIComponent(cjName.value.replace(/\s+/g,"")) + '&dishi=' + encodeURIComponent(dsName.value.replace(/\s+/g,"")), true);
	xhr.send(null);
	
}


function queryDone() {
  var listType1 =  document.getElementById("r1");
  var listType2 =  document.getElementById("r2");
  var listType3 =  document.getElementById("r3");
  if (listType1.checked) clickZbButton();
  if (listType2.checked) clickJcbcButton();
  if (listType3.checked) clickYbButton();
  
}

var el = document.getElementById('query_button');
if (el.addEventListener) {                       
  el.addEventListener('click', queryDone, false);                                     
} else {                                        
  el.attachEvent('onclick',queryDone);
}

var rlzy = document.getElementById('rl_button');
rlzy.addEventListener('click', clickRlButton, false);
