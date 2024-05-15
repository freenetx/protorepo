// package: lannet.api
// file: lannet/channel.proto

var lannet_channel_pb = require("../lannet/channel_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var CenterService = (function () {
  function CenterService() {}
  CenterService.serviceName = "lannet.api.CenterService";
  return CenterService;
}());

CenterService.Channel = {
  methodName: "Channel",
  service: CenterService,
  requestStream: true,
  responseStream: true,
  requestType: lannet_channel_pb.CenterRequest,
  responseType: lannet_channel_pb.CenterResponse
};

exports.CenterService = CenterService;

function CenterServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CenterServiceClient.prototype.channel = function channel(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(CenterService.Channel, {
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

exports.CenterServiceClient = CenterServiceClient;

var DirectNetService = (function () {
  function DirectNetService() {}
  DirectNetService.serviceName = "lannet.api.DirectNetService";
  return DirectNetService;
}());

DirectNetService.Channel = {
  methodName: "Channel",
  service: DirectNetService,
  requestStream: true,
  responseStream: true,
  requestType: lannet_channel_pb.DirectNetRequest,
  responseType: lannet_channel_pb.DirectNetResponse
};

exports.DirectNetService = DirectNetService;

function DirectNetServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

DirectNetServiceClient.prototype.channel = function channel(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(DirectNetService.Channel, {
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

exports.DirectNetServiceClient = DirectNetServiceClient;

