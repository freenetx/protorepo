// package: channel.api
// file: channel/channel.proto

var channel_channel_pb = require("../channel/channel_pb");
var share_share_pb = require("../share/share_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ChannelService = (function () {
  function ChannelService() {}
  ChannelService.serviceName = "channel.api.ChannelService";
  return ChannelService;
}());

ChannelService.Test = {
  methodName: "Test",
  service: ChannelService,
  requestStream: false,
  responseStream: false,
  requestType: share_share_pb.Empty,
  responseType: share_share_pb.Empty
};

ChannelService.Channel = {
  methodName: "Channel",
  service: ChannelService,
  requestStream: true,
  responseStream: true,
  requestType: channel_channel_pb.ChannelRequest,
  responseType: channel_channel_pb.ReceiveResponse
};

exports.ChannelService = ChannelService;

function ChannelServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ChannelServiceClient.prototype.test = function test(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ChannelService.Test, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ChannelServiceClient.prototype.channel = function channel(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(ChannelService.Channel, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  client.onMessage(function (message) {
    listeners.data.forEach(function (handler) {
      handler(message);
    })
  });
  client.start(metadata);
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.ChannelServiceClient = ChannelServiceClient;

