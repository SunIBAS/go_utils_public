class toFetch {
    // Success = 0
    // Error = 100
    // Fail = 200
    // End = 300
    constructor(url,method,content,ContentToString,parseMixin) {
        this.method = method || "";
        this.content = content || "";
        this.ContentToString = ContentToString || (_ => _);
        this.parseMixin = parseMixin || (_ => _);
        this.url = url;
        this.ErrorHandle = ()=>{};
        this.FailHandle = ()=>{};
        this.ErrorThrow = true;
        this.FailThrow = true;
    }
    setContent(content) {
        this.content = content;
        return this;
    }
    setMethod(method) {
        this.method = method;
        return this;
    }
    // 理解为无需处理返回值的请求
    Post() {
        let that = this;
        window.loading();
        return fetch(this.url, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            headers: {
                'Content-Type': 'application/json'
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify({
                content: this.ContentToString(this.content),
                method: this.method
            }) // body data type must match "Content-Type" header
        }).then(_ => _.text()).then(_ => {
            return that.parseMixin(_);
        }).then(JSON.parse)
            .then(_ => {
                // window.closeLoading();
                return new Promise(function (s,f) {
                    setTimeout(function () {
                        window.closeLoading();
                        if (_.Code === 0) {
                            s(_);
                        } else if (_.Code === 100) {
                            that.ErrorHandle(_);
                            f(new Error(_.Message));
                        } else if (_.Code === 200) {
                            that.FailHandle(_);
                            console.warn(_.Message);
                            f(new Error(_.Message));
                        }
                    },500);
                });
            })
            .catch(e => {
                window.closeLoading();
                alert(`请求[${that.method}]发生错误，${e.message}`);
                throw e;
            });
    }
    // 理解为需要将返回值做一定处理的请求
    PostParse() {
        return this.Post()
            .then(_ => {
                return JSON.parse(_.Content)
            })
    }

    DownloadFile(filename) {
        return fetch(host, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            headers: {
                'Content-Type': 'application/json'
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify({
                content: this.ContentToString(this.content),
                method: this.method
            }) // body data type must match "Content-Type" header
        }).then(_ => _.blob()).then(blob => {
            let url = window.URL.createObjectURL(blob);
            let a = document.createElement('a');
            a.href = url;
            a.download = filename;
            a.click();
            window.URL.revokeObjectURL(url);
        })
    }
}

class Base64String {
    constructor() {
    }
    encode(str){
        // 对字符串进行编码
        var encode = encodeURI(str);
        // 对编码的字符串转化base64
        var base64 = btoa(encode);
        return base64;
    }
    decode(base64){
        // 对base64转编码
        var decode = atob(base64);
        // 编码转字符串
        var str = decodeURI(decode);
        return str;
    }
    decodeNoDe(base64){
        // 对base64转编码
        var decode = atob(base64);
        return decode;
    }
}

// response = {
//     "Code": 0,
//     "Message": "success",
//     "Content": "[\n{\n\"Id\": \"3fac0f10-9bd3-4e24-9b56-583aaa22bc29\",\n\"Lang\": \"js\",\n\"Description\": \"测试\",\n\"Content\": \"console.log('123')\",\n\"CreateTime\": \"\",\n\"Tags\": \"\",\n\"Bak1\": \"\",\n\"Bak2\": \"\"\n}\n]"
// };