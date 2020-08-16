const insertCssOrJs = (function () {
    const insertDomToHead = (domType,doFn,onload) => {
        return new Promise(function (s) {
            let dom = document.createElement(domType);
            doFn(dom);
            document.head.append(dom);
            if (onload) {
                dom.onload = function () {
                    s();
                }
            } else {
                s();
            }
        })
    };
    const insertJSFile = (link) => {
        return insertDomToHead('script',function (script) {
            script.src = link;
        },true);
    };
    const insertCSSFile = (link) => {
        return insertDomToHead('link',function (dom) {
            dom.setAttribute('rel','stylesheet');
            dom.href = link;
        },false);
    };
    return (cb,files) => {
        console.log(files)
        let task = [];
        files.forEach(file => {
            if (file.split('?')[0].endsWith('.js')) {
                task.push(insertJSFile.bind(null,file));
            } else if (file.split('?')[0].endsWith('.css')) {
                task.push(insertCSSFile.bind(null,file));
            }
        });
        function doIt() {
            if (task.length) {
                task.shift()().then(doIt);
            } else {
                setTimeout(cb,200);
            }
        }
        setTimeout(function () {
            doIt();
        });
    };
})();

window.sourceIP = "";
window.baseIP = "";

window.loading = ()=>{};
window.closeLoading = ()=>{};
window.showMessage = (msg,type)=>{};

const injectIviewCore = function (cb,ofiles) {
    let sourcePath = "";
    let child = !(window.parent === window);
    window.sourceIP = location.origin;
    if (location.pathname.substring(0,7) !== "/public") {
        window.baseIP = `http://localhost:3000/`;
        window.sourceIP += "/go_utils_public/html"
    } else if (location.host === "www.sunibas.cn") {
        window.sourceIP += "/public/utils";
    } else {
        window.baseIP = window.sourceIP + "/";
        window.sourceIP += "/public"
    }
    cb = cb || (() => {});

    ofiles = ofiles || [];
    const canGetFromParentSource = [
        ...[
            "toFetch.js",
            "pubweb.js",
            "Actions.js",
            "index.js",
            "FileSaver.min.js",
        ].map(_ => window.sourceIP + "/Request/" + _)
    ];
    const baseSource = [
        ...[
            'iview.css',
            'vue.min.js',
            'iview.min.js',
        ].map(_ => window.sourceIP + "/iview/" + _),
        ...ofiles.map(_ => window.sourceIP + "/" + _)
    ];
    insertCssOrJs(function () {
        if (child) {
            window.RequestInstance = window.parent.RequestInstance;
            window.loading = window.parent.loading;
            window.closeLoading = window.parent.closeLoading;
            window.showMessage = window.parent.showMessage;
            window.showMessageType = window.parent.showMessageType;
            window.saveFile = window.parent.saveFile;
            window.inputDialog = window.parent.inputDialog;
        } else {
            window.RequestInstance(baseIP);
            window.saveFile = function (str,filename,type) {
                const blob = new Blob([str], { type: type || "" });
                if (!filename) {
                    filename = (new Date().getTime()) + ".txt";
                }
                saveAs(blob,filename);
            };
            window.inputDialog = function ($this,placeholder,cb,defaultValue,title) {
                let inpVal = "";
                let change = false;
                defaultValue = defaultValue || "";
                $this.$Modal.confirm({
                    render: (h) => {
                        return h('div',[
                            h('div',{},title || "请输入"),
                            h('Input', {
                                props: {
                                    value: defaultValue,
                                    autofocus: true,
                                    placeholder: placeholder || 'Please enter something'
                                },
                                on: {
                                    input: (val) => {
                                        inpVal = val;
                                        change = true;
                                    }
                                }
                            })
                        ]);
                    },
                    onOk() {
                        //$this.$Modal.remove();
                        cb(change ? inpVal : defaultValue);
                    }
                })
            }
        }
        cb();
    },[
        ...(child ? []:canGetFromParentSource),
        ...baseSource,
    ]);
};