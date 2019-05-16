export function formatDate(timestamp) {
    var date = new Date(timestamp * 1000)
    var Y = date.getFullYear() + '-'
    var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-'
    var D = date.getDate() + ' '
    var h = (date.getHours() + 1 < 10 ? '0' + (date.getHours() + 1) : date.getHours() + 1) + ':'
    var m = (date.getMinutes() + 1 < 10 ? '0' + (date.getMinutes() + 1) : date.getMinutes() + 1) + ':'
    var s = (date.getSeconds() + 1 < 10 ? '0' + (date.getSeconds() + 1) : date.getSeconds() + 1)
    return Y + M + D + h + m + s
  }
  
  export function formatDuring(seconds_ts) {
    var days = Math.floor(seconds_ts / (60 * 60 * 24))
    var hours = Math.floor((seconds_ts % (60 * 60 * 24)) / (60 * 60))
    var minutes = Math.floor((seconds_ts % (60 * 60)) / (60))
    var seconds = Math.round(seconds_ts % 60)
    return days + '天' + hours + '小时' + minutes + '分钟' + seconds + '秒'
  }
  
  