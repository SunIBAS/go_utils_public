function readFile(file) {
    return new Promise(function (s) {
        var reader = new FileReader();//new一个FileReader实例
        reader.onload = function () {
            s(this.result);
        };
        reader.readAsText(file);
    });
}