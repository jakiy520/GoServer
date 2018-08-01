Page({
  data: {
    latitude: 28.643813,
    longitude: 121.468869,
    markers: [{
      id: 1,
      latitude: 28.643813,
      longitude: 121.468869,
      name: '如果果业'
    }]
    
  },
  onReady: function (e) {
    this.mapCtx = wx.createMapContext('myMap')
  }
})
