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
    ProInfo: "",
    shareUserId: "",
  },

  getProInfo: function () {
    var that = this;
    wx.request({
      url: "https://rggy.godwork.cn/api/getKanjiaPro",
      success: function (res) {
        var product = res.data.data.product;
        // console.log(product);
        product.image.url = config.static.imageDomain + product.image.url;
        for (var i = 0; i < product.images.length; i++) {
          var url = product.images[i].url;
          product.images[i].url = config.static.imageDomain + url;
        }
        // console.log(product);
        that.setData({
          ProInfo: res.data.data.product
        })

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
    this.setData({
      shareUserId: options.shareUserId
    })
    if (app.globalData.userInfo) {
      this.setData({
        userInfo: app.globalData.userInfo,
        hasUserInfo: true
      })
    } else if (this.data.canIUse) {
      // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
      // 所以此处加入 callback 以防止这种情况
      app.userInfoReadyCallback = res => {
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
        console.log(this.data.userInfo)
      }
    }
    //  获取商品信息
    this.getProInfo();
  },

  //  分享按钮
  onShareAppMessage: function (res) {
    if (res.from === 'button') {
      // 来自页面内转发按钮
      console.log(res.target)
    }
    return {
      title: '水果砍价',
      path: '/pages/kanjia/kanjia?shareUserId=' + this.data.userInfo.userid,
      success: function (res) {
        // 转发成功
      },
      fail: function (res) {
        // 转发失败
      }
    }
  },
  //图片点击事件
  imgYu: function (event) {
    var src = event.currentTarget.dataset.src;//获取data-src
    console.log(src);
    // var imgList = event.currentTarget.dataset.list;//获取data-list
    //图片预览
    wx.previewImage({
      current: src, // 当前显示图片的http链接
      urls: [src] // 需要预览的图片http链接列表
    })
  },
})
