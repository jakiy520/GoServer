//logs.js
const util = require('../../utils/util.js')

Page({
  data: {
    kanjiaRecords: "",
  },
  onLoad: function (options) {
    //  获取砍价ID参数
    if (options.kanjiaID > 0) {
      this.setData({
        kanjiaID: options.kanjiaID
      })
      // console.log(options.kanjiaID)
      this.getKanjiaRecords(options.kanjiaID)

    }
  },
  //  获取砍价记录
  getKanjiaRecords: function (kanjiaID) {
    var that = this;
    wx.request({
      url: "https://rggy.godwork.cn/api/GetKanjiaRecords/" + kanjiaID,
      header: {
        'content-type': 'application/json',
      },
      method: "Get",
      success: function (res) {
        var kanjiaRecords = res.data.data.KanjiaRecords;
        console.log(kanjiaRecords);
        for (var i = 0; i < kanjiaRecords.length; i++) {
          console.log(kanjiaRecords[i].createdAt);
          kanjiaRecords[i].createdAt = util.formatTimeISO(kanjiaRecords[i].createdAt);
        }

        that.setData({
          kanjiaRecords: kanjiaRecords,
        })

        console.log(kanjiaRecords);
      }
    });
  }
})
