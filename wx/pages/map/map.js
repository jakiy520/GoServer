Page({
  data: {
    latitude:  28.642557,
    longitude: 121.471452,
    markers: [{
      id: 1,
      latitude: 28.642557,
      longitude: 121.471452,
      name: '如果果业'
    }]
    
  },
  onReady: function (e) {
    this.mapCtx = wx.createMapContext('myMap')
  }
})
