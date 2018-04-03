var config = require("../../config/config.js");
//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    motto: 'Hello World',
    userInfo: app.globalData.userInfo,
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    proInfo: "",
    kanjiaID: 0,//砍价的id 如果该id为0，说明来到该页面的还未参与砍价活动
    kanjiaInfo: "",//被砍价的信息
    kanjiaMoney: 0.00
  },

  getProInfo: function () {
    console.log(this.data.userInfo)
    var that = this;
    wx.request({
      url: "https://rggy.godwork.cn/api/getKanjiaPro/" + this.data.kanjiaID + "/" + that.data.userInfo.userid,
      success: function (res) {
        // console.log(res)
        var product = res.data.data.product;
        var kanjiainfo = res.data.data.kanjiaInfo;
        // console.log(product);
        product.image.url = config.static.imageDomain + product.image.url;
        for (var i = 0; i < product.images.length; i++) {
          var url = product.images[i].url;
          product.images[i].url = config.static.imageDomain + url;
        }
        // console.log(kanjiainfo);
        that.setData({
          proInfo: product,
          kanjiaInfo: kanjiainfo,
          kanjiaMoney: res.data.data.kanjiaMoney,
        })
        if (kanjiainfo.id > 0) {
          that.setData({
            kanjiaID: kanjiainfo.id,
          })
        }
      }
    });
  },

  //事件处理函数
  bindViewTap: function () {
    wx.navigateTo({
      url: '/pages/logs/logs'
    })
  },
  onLoad: function (options) {
    //  获取砍价ID参数
    if (options.kanjiaID > 0) {
      this.setData({
        kanjiaID: options.kanjiaID
      })
    }
    if (app.globalData.userInfo) {
      this.setData({
        userInfo: app.globalData.userInfo,
        hasUserInfo: true
      })
      //  获取商品信息
      this.getProInfo();
    } else if (this.data.canIUse) {
      // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
      // 所以此处加入 callback 以防止这种情况
      app.userInfoReadyCallback = res => {
        // console.log(res)
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
        // console.log(this.data.userInfo)
        //  获取商品信息
        this.getProInfo();
      }
    }
  },

  //  分享按钮
  onShareAppMessage: function (res) {
    if (res.from === 'button') {
      // 来自页面内转发按钮
      // console.log(res.target)
    }
    return {
      title: '水果砍价',
      path: '/pages/kanjia/kanjia?kanjiaID=' + this.data.kanjiaID,
      success: function (res) {
        // 转发成功
      },
      fail: function (res) {
        // 转发失败
      }
    }
  },

  onPullDownRefresh: function () {
    this.getProInfo();
    // this.setData({
    //   kanjiaID:0
    // })
    wx.stopPullDownRefresh()
  },

  //  参与砍价活动
  joinKanjia: function () {
    // console.log(this.data.userInfo)
    var that = this;
    wx.request({
      url: "https://rggy.godwork.cn/api/JoinKanjia",
      data: {
        userID: this.data.userInfo.userid,
        userNickName: this.data.userInfo.nickName,
        userAvatarUrl: this.data.userInfo.avatarUrl,
        productID: this.data.proInfo.id,
      },
      header: {
        'content-type': 'application/json',
      },
      method: "POST",
      success: function (res) {
        // console.log(res)
        that.setData({
          kanjiaID: res.data.data.kanjiaID
        })
        that.getProInfo();
      }
    });
  },

  //  帮他砍
  bangtakan: function () {
    // console.log(this.data.userInfo)
    var that = this;
    wx.request({
      url: "https://rggy.godwork.cn/api/Bangtakan",
      data: {
        userID: this.data.userInfo.userid,
        userNickName: this.data.userInfo.nickName,
        userAvatarUrl: this.data.userInfo.avatarUrl,
        kanjiaID: this.data.kanjiaID,
        productID: this.data.proInfo.id,
      },
      header: {
        'content-type': 'application/json',
      },
      method: "POST",
      success: function (res) {
        console.log(res)
        if (res.data.errNo == "1") {
          wx.showToast({
            title: res.data.msg,
            icon: "none",
            duration: 3000
          })
        }
        if (res.data.errNo == "0") {
          wx.showToast({
            title: "砍价成功",
            icon: "success",
            duration: 2000
          })
          that.setData({
            kanjiaMoney: res.data.data.AllKanjiaMoney
          })
        }
      }
    });
  },
})
