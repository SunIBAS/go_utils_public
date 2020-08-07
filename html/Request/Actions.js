Actions = (function() {
    class DemoAction{
        constructor(toFetch,b64str) {
            this.toFetch = toFetch;
            this.b64str = b64str;
            this.date = new Date();
        }

        InsertDemo(lang,description,content) {
            return this.toFetch
                .setMethod("InsertDemo")
                .setContent(JSON.stringify({lang,description,content: this.b64str.encode(content)}))
                .Post();
        }

        UpdateDemo(id,lang,description,content) {
            if (!id) {  window.showMessage("id 为空","error"); return; }
            return this.toFetch
                .setMethod("UpdateDemo")
                .setContent(JSON.stringify({id,lang,description,content: this.b64str.encode(content)}))
                .Post();
        }

        ListDemo(page,count,lang,description) {
            let params = {page,count};
            if (lang || description) {
                params = {
                    ...params,
                    other: JSON.stringify({
                        ...(lang ? {lang} : {}),
                        ...(description ? {description} : {})
                    })
                };
            }
            let $this = this;
            return this.toFetch
                .setMethod("ListDemo")
                .setContent(JSON.stringify(params))
                .PostParse()
                .then(arr => {
                    return arr.map(demo => {
                        $this.date.setTime(demo.CreateTime.substring(0,13));
                        demo.content = this.b64str.decode(demo.Content);
                        demo.CreateTime = `${$this.date.getFullYear()}-${$this.date.getMonth() + 1}-${$this.date.getDate()}#${$this.date.getHours()}:${$this.date.getMinutes()}:${$this.date.getSeconds()}`;
                        return demo;
                    });
                });
        }

        // GetByDemoId(id) {
        //     if (!id) {  window.showMessage("id 为空","error"); return; }
        //     return this.toFetch
        //         .setMethod("GetByDemoId")
        //         .setContent(JSON.stringify({id}))
        //         .Post();
        // }

        GetDemoCount() {
            return this.toFetch
                .setMethod("GetDemoCount")
                .setContent(JSON.stringify({}))
                .Post();
        }

        DeleteDemo(id) {
            if (!id) {  window.showMessage("id 为空","error"); return; }
            return this.toFetch
                .setMethod("DeleteDemo")
                .setContent(JSON.stringify({id}))
                .Post();
        }

        RunDeom(id) {
            if (!id) {  window.showMessage("id 为空","error"); return; }
            return this.toFetch
                .setMethod("RunDeom")
                .setContent(JSON.stringify({id}))
                .Post()
                .then(_ => {
                    _.Message = this.b64str.decodeNoDe(_.Message);
                    return _;
                });
        }
    }

    class FileOpAction {
        constructor(toFetch,b64str) {
            this.toFetch = toFetch;
            this.b64str = b64str;
        }

        GetSubFiles(path,child,zip) {
            if (!path) {  window.showMessage("path 为空","error"); return; }
            child = !!child;
            zip = zip ? zip : (typeof zip === "boolean" ? zip : true);
            return this.toFetch
                .setMethod("GetSubFiles")
                .setContent(JSON.stringify({path,child}))
                .PostParse();
        }

    }

    return class {
        constructor(baseIp) {
            this.url = baseIp + "api";
            this.toFetch = new toFetch(this.url);
            this.b64str = new Base64String();
            this.DemoAction = new DemoAction(this.toFetch,this.b64str);
            this.FileOpAction = new FileOpAction(this.toFetch,this.b64str);
        }
    }
})();