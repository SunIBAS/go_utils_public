let pubWeb = class {
    constructor(baseIp) {
        this.url = baseIp + "pubWebAction";
        this.toFetch = new toFetch(this.url);
    }
    list(page,count) {
        return this.toFetch.setContent(JSON.stringify({page,count}))
            .setMethod("list")
            .Post();
    }
    // dir 电脑的路径
    // perpath 访问前缀
    add(dir,perpath) {
        return this.toFetch.setContent(JSON.stringify({dir,perpath}))
            .setMethod("add")
            .Post();
    }
    checkUrl(perpath) {
        return this.toFetch.setContent(JSON.stringify({perpath}))
            .setMethod("checkUrl")
            .Post();
    }
};