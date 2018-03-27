var config = require("../config/config.js");

var login = {
  login: function (app) {
    var self = this;
    var resData = {};
    var jsCodeDone, userInfoDone;

    function setUserInfo() {
      wx.request({
        url: config.api.setWeAppUser,
        data: {
          encryptedData: resData.encryptedData,
          iv: resData.iv
        },
        header: {
          'content-type': 'application/json',
          'Cookie': "sid=" + resData.sid
        },
        method: "POST",
        success: function (res) {
          resData.userInfo.userid = res.data.data.userid;
          // console.log(resData.userInfo)
          app.globalData.userInfo = resData.userInfo;
          app.globalData.encryptedData = resData.encryptedData;
          app.globalData.iv = resData.iv;
          app.globalData.sid = resData.sid;

          // console.log(app.globalData.userInfo);
          // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
          // 所以此处加入 callback 以防止这种情况
          if (app.userInfoReadyCallback) {
            app.userInfoReadyCallback(resData)
          }
        }
      });
    }

    wx.login({
      success: function (res) {
        if (res.code) {
          wx.request({
            url: config.api.weAppLogin,
            data: {
              code: res.code
            },
            success: function (res) {
              // console.log(res)
              resData.sid = res.data.data.sid;
              jsCodeDone = true;
              jsCodeDone && userInfoDone && setUserInfo();
            }
          });

          wx.getUserInfo({
            success: function (res) {
              resData.userInfo = res.userInfo;
              resData.encryptedData = res.encryptedData;
              resData.iv = res.iv;
              userInfoDone = true;
              jsCodeDone && userInfoDone && setUserInfo();

              
            },
            fail: function (data) {
              console.log(data);
            }
          });
        }
      }
    });
  },

  logout: function () {

  }
}

module.exports = login;