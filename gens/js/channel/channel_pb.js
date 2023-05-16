// source: proto/channel/channel.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var share_share_pb = require('../../share/share_pb.js');
goog.object.extend(proto, share_share_pb);
goog.exportSymbol('proto.file.api.ChannelRequest', null, global);
goog.exportSymbol('proto.file.api.ChannelRequest.ReqCase', null, global);
goog.exportSymbol('proto.file.api.Data', null, global);
goog.exportSymbol('proto.file.api.KeyAccept', null, global);
goog.exportSymbol('proto.file.api.ReceiveResponse', null, global);
goog.exportSymbol('proto.file.api.ReceiveResponse.RespCase', null, global);
goog.exportSymbol('proto.file.api.SetKey', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.file.api.SetKey = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.file.api.SetKey, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.file.api.SetKey.displayName = 'proto.file.api.SetKey';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.file.api.Data = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.file.api.Data, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.file.api.Data.displayName = 'proto.file.api.Data';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.file.api.ChannelRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.file.api.ChannelRequest.oneofGroups_);
};
goog.inherits(proto.file.api.ChannelRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.file.api.ChannelRequest.displayName = 'proto.file.api.ChannelRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.file.api.KeyAccept = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.file.api.KeyAccept, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.file.api.KeyAccept.displayName = 'proto.file.api.KeyAccept';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.file.api.ReceiveResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.file.api.ReceiveResponse.oneofGroups_);
};
goog.inherits(proto.file.api.ReceiveResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.file.api.ReceiveResponse.displayName = 'proto.file.api.ReceiveResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.file.api.SetKey.prototype.toObject = function(opt_includeInstance) {
  return proto.file.api.SetKey.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.file.api.SetKey} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.SetKey.toObject = function(includeInstance, msg) {
  var f, obj = {
    key: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.file.api.SetKey}
 */
proto.file.api.SetKey.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.file.api.SetKey;
  return proto.file.api.SetKey.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.file.api.SetKey} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.file.api.SetKey}
 */
proto.file.api.SetKey.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setKey(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.file.api.SetKey.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.file.api.SetKey.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.file.api.SetKey} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.SetKey.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKey();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string key = 1;
 * @return {string}
 */
proto.file.api.SetKey.prototype.getKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.file.api.SetKey} returns this
 */
proto.file.api.SetKey.prototype.setKey = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.file.api.Data.prototype.toObject = function(opt_includeInstance) {
  return proto.file.api.Data.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.file.api.Data} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.Data.toObject = function(includeInstance, msg) {
  var f, obj = {
    data: msg.getData_asB64()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.file.api.Data}
 */
proto.file.api.Data.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.file.api.Data;
  return proto.file.api.Data.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.file.api.Data} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.file.api.Data}
 */
proto.file.api.Data.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setData(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.file.api.Data.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.file.api.Data.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.file.api.Data} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.Data.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getData_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
};


/**
 * optional bytes data = 1;
 * @return {!(string|Uint8Array)}
 */
proto.file.api.Data.prototype.getData = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes data = 1;
 * This is a type-conversion wrapper around `getData()`
 * @return {string}
 */
proto.file.api.Data.prototype.getData_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getData()));
};


/**
 * optional bytes data = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getData()`
 * @return {!Uint8Array}
 */
proto.file.api.Data.prototype.getData_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getData()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.file.api.Data} returns this
 */
proto.file.api.Data.prototype.setData = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.file.api.ChannelRequest.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.file.api.ChannelRequest.ReqCase = {
  REQ_NOT_SET: 0,
  SET_KEY: 1,
  DATA: 2
};

/**
 * @return {proto.file.api.ChannelRequest.ReqCase}
 */
proto.file.api.ChannelRequest.prototype.getReqCase = function() {
  return /** @type {proto.file.api.ChannelRequest.ReqCase} */(jspb.Message.computeOneofCase(this, proto.file.api.ChannelRequest.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.file.api.ChannelRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.file.api.ChannelRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.file.api.ChannelRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.ChannelRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    setKey: (f = msg.getSetKey()) && proto.file.api.SetKey.toObject(includeInstance, f),
    data: (f = msg.getData()) && proto.file.api.Data.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.file.api.ChannelRequest}
 */
proto.file.api.ChannelRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.file.api.ChannelRequest;
  return proto.file.api.ChannelRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.file.api.ChannelRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.file.api.ChannelRequest}
 */
proto.file.api.ChannelRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.file.api.SetKey;
      reader.readMessage(value,proto.file.api.SetKey.deserializeBinaryFromReader);
      msg.setSetKey(value);
      break;
    case 2:
      var value = new proto.file.api.Data;
      reader.readMessage(value,proto.file.api.Data.deserializeBinaryFromReader);
      msg.setData(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.file.api.ChannelRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.file.api.ChannelRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.file.api.ChannelRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.ChannelRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSetKey();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.file.api.SetKey.serializeBinaryToWriter
    );
  }
  f = message.getData();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.file.api.Data.serializeBinaryToWriter
    );
  }
};


/**
 * optional SetKey set_key = 1;
 * @return {?proto.file.api.SetKey}
 */
proto.file.api.ChannelRequest.prototype.getSetKey = function() {
  return /** @type{?proto.file.api.SetKey} */ (
    jspb.Message.getWrapperField(this, proto.file.api.SetKey, 1));
};


/**
 * @param {?proto.file.api.SetKey|undefined} value
 * @return {!proto.file.api.ChannelRequest} returns this
*/
proto.file.api.ChannelRequest.prototype.setSetKey = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.file.api.ChannelRequest.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.file.api.ChannelRequest} returns this
 */
proto.file.api.ChannelRequest.prototype.clearSetKey = function() {
  return this.setSetKey(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.file.api.ChannelRequest.prototype.hasSetKey = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Data data = 2;
 * @return {?proto.file.api.Data}
 */
proto.file.api.ChannelRequest.prototype.getData = function() {
  return /** @type{?proto.file.api.Data} */ (
    jspb.Message.getWrapperField(this, proto.file.api.Data, 2));
};


/**
 * @param {?proto.file.api.Data|undefined} value
 * @return {!proto.file.api.ChannelRequest} returns this
*/
proto.file.api.ChannelRequest.prototype.setData = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.file.api.ChannelRequest.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.file.api.ChannelRequest} returns this
 */
proto.file.api.ChannelRequest.prototype.clearData = function() {
  return this.setData(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.file.api.ChannelRequest.prototype.hasData = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.file.api.KeyAccept.prototype.toObject = function(opt_includeInstance) {
  return proto.file.api.KeyAccept.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.file.api.KeyAccept} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.KeyAccept.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.file.api.KeyAccept}
 */
proto.file.api.KeyAccept.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.file.api.KeyAccept;
  return proto.file.api.KeyAccept.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.file.api.KeyAccept} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.file.api.KeyAccept}
 */
proto.file.api.KeyAccept.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.file.api.KeyAccept.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.file.api.KeyAccept.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.file.api.KeyAccept} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.KeyAccept.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.file.api.ReceiveResponse.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.file.api.ReceiveResponse.RespCase = {
  RESP_NOT_SET: 0,
  KEY_ACCEPT: 1,
  DATA: 2
};

/**
 * @return {proto.file.api.ReceiveResponse.RespCase}
 */
proto.file.api.ReceiveResponse.prototype.getRespCase = function() {
  return /** @type {proto.file.api.ReceiveResponse.RespCase} */(jspb.Message.computeOneofCase(this, proto.file.api.ReceiveResponse.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.file.api.ReceiveResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.file.api.ReceiveResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.file.api.ReceiveResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.ReceiveResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    keyAccept: (f = msg.getKeyAccept()) && proto.file.api.KeyAccept.toObject(includeInstance, f),
    data: (f = msg.getData()) && proto.file.api.Data.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.file.api.ReceiveResponse}
 */
proto.file.api.ReceiveResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.file.api.ReceiveResponse;
  return proto.file.api.ReceiveResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.file.api.ReceiveResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.file.api.ReceiveResponse}
 */
proto.file.api.ReceiveResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.file.api.KeyAccept;
      reader.readMessage(value,proto.file.api.KeyAccept.deserializeBinaryFromReader);
      msg.setKeyAccept(value);
      break;
    case 2:
      var value = new proto.file.api.Data;
      reader.readMessage(value,proto.file.api.Data.deserializeBinaryFromReader);
      msg.setData(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.file.api.ReceiveResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.file.api.ReceiveResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.file.api.ReceiveResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.file.api.ReceiveResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKeyAccept();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.file.api.KeyAccept.serializeBinaryToWriter
    );
  }
  f = message.getData();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.file.api.Data.serializeBinaryToWriter
    );
  }
};


/**
 * optional KeyAccept key_accept = 1;
 * @return {?proto.file.api.KeyAccept}
 */
proto.file.api.ReceiveResponse.prototype.getKeyAccept = function() {
  return /** @type{?proto.file.api.KeyAccept} */ (
    jspb.Message.getWrapperField(this, proto.file.api.KeyAccept, 1));
};


/**
 * @param {?proto.file.api.KeyAccept|undefined} value
 * @return {!proto.file.api.ReceiveResponse} returns this
*/
proto.file.api.ReceiveResponse.prototype.setKeyAccept = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.file.api.ReceiveResponse.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.file.api.ReceiveResponse} returns this
 */
proto.file.api.ReceiveResponse.prototype.clearKeyAccept = function() {
  return this.setKeyAccept(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.file.api.ReceiveResponse.prototype.hasKeyAccept = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Data data = 2;
 * @return {?proto.file.api.Data}
 */
proto.file.api.ReceiveResponse.prototype.getData = function() {
  return /** @type{?proto.file.api.Data} */ (
    jspb.Message.getWrapperField(this, proto.file.api.Data, 2));
};


/**
 * @param {?proto.file.api.Data|undefined} value
 * @return {!proto.file.api.ReceiveResponse} returns this
*/
proto.file.api.ReceiveResponse.prototype.setData = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.file.api.ReceiveResponse.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.file.api.ReceiveResponse} returns this
 */
proto.file.api.ReceiveResponse.prototype.clearData = function() {
  return this.setData(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.file.api.ReceiveResponse.prototype.hasData = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.file.api);
