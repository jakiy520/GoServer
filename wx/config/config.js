var url       = 'https://rggy.godwork.cn';
var apiPrefix = url + '/api';

var config = {
    name: "爱宝宝微商城",
    wemallSession: "wemallSession",
    static: {
        imageDomain: url
    },
    api: {
        weAppLogin: '/weAppLogin',
        setWeAppUser: '/setWeAppUser',
    }
};

for (var key in config.api) {
    config.api[key] = apiPrefix + config.api[key];
}

module.exports = config;