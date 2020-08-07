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

const injectIviewCore = function (cb) {
    let sourcePath = "";
    let child = !(window.parent === window);
    window.sourceIP = location.origin;
    if (location.pathname.substring(0,7) !== "/public") {
        window.baseIP = `http://localhost:3000/`;
        window.sourceIP += "/go_utils_public/html"
    } else {
        window.baseIP = window.sourceIP + "/";
        window.sourceIP += "/public"
    }
    cb = cb || (() => {});

    const canGetFromParentSource = [
        ...[
            "toFetch.js",
            "pubweb.js",
            "Actions.js",
            "index.js",
        ].map(_ => window.sourceIP + "/Request/" + _)
    ];
    const baseSource = [
        ...[
            'iview.css',
            'vue.min.js',
            'iview.min.js',
        ].map(_ => window.sourceIP + "/iview/" + _)
    ];
    insertCssOrJs(function () {
        if (child) {
            window.RequestInstance = window.parent.RequestInstance;
            window.loading = window.parent.loading;
            window.closeLoading = window.parent.closeLoading;
            window.showMessage = window.parent.showMessage;
        } else {
            window.RequestInstance(baseIP);
        }
        cb();
    },[
        ...(child ? []:canGetFromParentSource),
        ...baseSource,
    ]);
};