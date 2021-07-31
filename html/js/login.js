$("#login-button").on("click", function () {
    var username = $("#username")
    var password = $("#password")
    if (username.val().trim().length === 0) {
        username.attr("placeholder","账号不能为空")
        showToast("账号不能为空")
        return
    } else if (password.val().trim().length === 0) {
        password.attr("placeholder","密码不能为空")
        showToast("密码不能为空")
        return
    }

    var params=new URLSearchParams()
    params.append("username",username.val())
    params.append("password",sha256(password.val()))
    params.append("gw_address",GatewayAddress)
    params.append("gw_port",GatewayPort)
    params.append("mac",MAC)

    $.ajax({
        type: "POST",
        url: "/fas/login",
        data: params.toString(),
        success:function(data,textStatus,jqXHR){
            showToast(data.msg);
            if (data.code==200){
                // window.location.replace(data.data)
                window.location.href=data.data
            }
        },
        error:function(xhr,error){
            showToast('出错了，错误码：'+xhr.status);
        },
        complete:function(){
            console.log('结束');
        }
    });
})

function showToast(msg,duration){
    duration=isNaN(duration)?1000:duration;
    var m = document.createElement('div');
    m.innerHTML = msg;
    m.style.cssText="width:20%; min-width:20%; background:#fff; opacity:0.6; height:auto;min-height: 30px; color:#000; line-height:30px; text-align:center; border-radius:4px; position:fixed; top:80%; left:40%; z-index:999999;";
    document.body.appendChild(m);
    setTimeout(function() {
        var d = 0.5;
        m.style.webkitTransition = '-webkit-transform ' + d + 's ease-in, opacity ' + d + 's ease-in';
        m.style.opacity = '0';
        setTimeout(function() { document.body.removeChild(m) }, d * 1000);
    }, duration);
}