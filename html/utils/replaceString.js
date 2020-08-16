function replaceString(s,form) {
    let fs = [s];
    form.forEach(f => {
        let fs_ = [];
        let reg = new RegExp(`{${f.name}}`,'g');
        f.vals.forEach(fv => {
            fs.forEach(fsCopy => {
                fs_.push(fsCopy.replace(reg,fv))
            });
        });
        fs = fs_;
    });
    return fs;
}
/*
var str = "{year}/{dd}/GLDAS_NOAH025_3H.A{year}{month}{day}.{param0}.021.{param1}";
var form = [
    {
        name: 'year',
        vals: ['2018']
    },
    {
        name: 'month',
        vals: ['01']
    },
    {
        name: 'day',
        vals: ['01']
    },
    {
        name: 'dd',
        vals: ['001']
    },
    {
        name: 'param0',
        vals: ['0300','0900']
    },
    {
        name: 'param1',
        vals: ['.nc4','.nc4.xml']
    }
];
var ret = replaceString(str,form);
ret = [
    "2018/001/GLDAS_NOAH025_3H.A20180101.0300.021.nc4",
    "2018/001/GLDAS_NOAH025_3H.A20180101.0300.021.nc4.xml",
    "2018/001/GLDAS_NOAH025_3H.A20180101.0900.021.nc4",
    "2018/001/GLDAS_NOAH025_3H.A20180101.0900.021.nc4.xml",
];
*/