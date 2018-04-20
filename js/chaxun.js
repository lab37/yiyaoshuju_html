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
	if (checkInput("zb_yaopin_name")&&checkInput("qiye_name")) {
		alert('你未输入任何内容，请检查后重新输入');
		return;
	}
    var ypName = document.getElementById("zb_yaopin_name");
    var cjName = document.getElementById("qiye_name");
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		if (xhr.status == 200) {
			  responseObject = JSON.parse(xhr.responseText);
			  var newContent = '';
           	  newContent += '<tr>';
              newContent += '<td>通用名</td>';
              newContent += '<td>剂型</td>';
              newContent += '<td>规格</td>';
              newContent += '<td>商品名</td>';
              newContent += '<td>价格</td>';
              newContent += '<td>价格说明</td>';
              newContent += '<td>包装</td>';
              newContent += '<td>药品类别</td>';
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
                newContent += '<td>'+responseObject[i].Bz+'</td>';
                newContent += '<td>'+responseObject[i].Yplb+'</td>';
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
	if (checkInput("yb_yaopin_name")) {
		alert('你未输入任何内容，请检查后重新输入');
		return;
	}
    var ypName = document.getElementById("yb_yaopin_name");
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		if (xhr.status == 200) {
			  responseObject = JSON.parse(xhr.responseText);
                          alert(xhr.responseText);
			  var newContent = '';
              newContent += '<tr>';
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

var zb = document.getElementById('zhongbiao_button');
var yb = document.getElementById('yibao_button');
zb.addEventListener('click', clickZbButton, false);
yb.addEventListener('click', clickYbButton, false);