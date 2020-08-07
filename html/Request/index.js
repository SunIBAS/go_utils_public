window.RequestInstance = (function (baseIp) {
    window.RequestInstance = {
        pubweb: new pubWeb(baseIp),
        action: new Actions(baseIp),
        flowChartWeb: baseIp + "/public/flowchart/index.html"
    }
});