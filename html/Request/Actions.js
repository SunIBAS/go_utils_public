const Actions = (function() {
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
                        demo.Content = this.b64str.decode(demo.Content);
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

    class RecordsAction {
        constructor(toFetch,b64str,tag) {
            this.toFetch = toFetch;
            this.b64str = b64str;
            this.date = new Date();
            this.tag = tag;
        }

        // new Promise(s => s(id))
        InsertRecords(content,field1,field2,field3) {
            return this.toFetch
                .setMethod("InsertRecords")
                .setContent(JSON.stringify({tag: this.tag,content: this.b64str.encode(content),field1,field2,field3,}))
                .Post()
                .then(_ => {
                    return _.Content;
                });
        }

        UpdateRecords(id,content,field1,field2,field3) {
            if (!id) {  window.showMessage("id 为空","error"); return; }
            return this.toFetch
                .setMethod("UpdateRecords")
                .setContent(JSON.stringify({id,content: this.b64str.encode(content),field1,field2,field3,tag: this.tag}))
                .Post();
        }

        ListRecords(page,count,field1,field2,field3) {
            let params = {
                page: typeof page === "number" ? page : 0,
                count: typeof count === "number" ? count : 10,
                other: JSON.stringify({
                    tag: this.tag,
                    ...(field1 ? {field1} : {}),
                    ...(field2 ? {field2} : {}),
                    ...(field3 ? {field3} : {}),
                })
            };
            let $this = this;
            return this.toFetch
                .setMethod("ListRecords")
                .setContent(JSON.stringify(params))
                .PostParse()
                .then(arr => {
                    return arr.map(record => {
                        $this.date.setTime(record.CreateTime.substring(0,13));
                        record.Content = this.b64str.decode(record.Content);
                        record.CreateTime = `${$this.date.getFullYear()}-${$this.date.getMonth() + 1}-${$this.date.getDate()}#${$this.date.getHours()}:${$this.date.getMinutes()}:${$this.date.getSeconds()}`;
                        return record;
                    });
                });
        }

        GetRecordsCount() {
            return this.toFetch
                .setMethod("GetRecordsCount")
                .setContent(JSON.stringify({}))
                .Post();
        }

        DeleteRecords(id) {
            if (!id) {  window.showMessage("id 为空","error"); return; }
            return this.toFetch
                .setMethod("DeleteRecords")
                .setContent(JSON.stringify({id}))
                .Post();
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
            this.FileActionRecords = new RecordsAction(this.toFetch,this.b64str,"fileAction");
            this.FilesRecords = new RecordsAction(this.toFetch,this.b64str,"files");
            this.InitSourcesFromDB = () =>
                this.toFetch
                .setMethod("InitSourcesFromDB").Post();
        }
    }
})();