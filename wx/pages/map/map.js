Page({
  data: {
    latitude: 28.637681,
    longitude: 121.462591,
    markers: [{
      id: 1,
      latitude: 28.637681,
      longitude: 121.462591,
      name: '如果果业'
    }]
    
  },
  onReady: function (e) {
    this.mapCtx = wx.createMapContext('myMap')
  },
  calling1:function(){
    wx.makePhoneCall({
      phoneNumber: '0576-88123056', //此号码并非真实电话号码，仅用于测试
      success: function () {
        // console.log("拨打电话成功！")
      },
      fail: function () {
        // console.log("拨打电话失败！")
      }
    })
  },
  calling2: function () {
    wx.makePhoneCall({
      phoneNumber: '13666804352', //此号码并非真实电话号码，仅用于测试
      success: function () {
        // console.log("拨打电话成功！")
      },
      fail: function () {
        // console.log("拨打电话失败！")
      }
    })
  }
})
