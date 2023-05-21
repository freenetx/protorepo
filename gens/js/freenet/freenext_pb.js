// source: proto/freenet/freenext.proto
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

goog.exportSymbol('proto.freenet.api.ClientMessage', null, global);
goog.exportSymbol('proto.freenet.api.ClientMessage.RequestCase', null, global);
goog.exportSymbol('proto.freenet.api.ClientMessage.ResponseCase', null, global);
goog.exportSymbol('proto.freenet.api.Data', null, global);
goog.exportSymbol('proto.freenet.api.MessageAck', null, global);
goog.exportSymbol('proto.freenet.api.ServerMessage', null, global);
goog.exportSymbol('proto.freenet.api.ServerMessage.RequestCase', null, global);
goog.exportSymbol('proto.freenet.api.ServerMessage.ResponseCase', null, global);
goog.exportSymbol('proto.freenet.api.SyncMessage', null, global);
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
proto.freenet.api.SyncMessage = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.freenet.api.SyncMessage.repeatedFields_, null);
};
goog.inherits(proto.freenet.api.SyncMessage, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.freenet.api.SyncMessage.displayName = 'proto.freenet.api.SyncMessage';
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
proto.freenet.api.MessageAck = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.freenet.api.MessageAck, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.freenet.api.MessageAck.displayName = 'proto.freenet.api.MessageAck';
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
proto.freenet.api.Data = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.freenet.api.Data, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.freenet.api.Data.displayName = 'proto.freenet.api.Data';
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
proto.freenet.api.ClientMessage = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.freenet.api.ClientMessage.oneofGroups_);
};
goog.inherits(proto.freenet.api.ClientMessage, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.freenet.api.ClientMessage.displayName = 'proto.freenet.api.ClientMessage';
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
proto.freenet.api.ServerMessage = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.freenet.api.ServerMessage.oneofGroups_);
};
goog.inherits(proto.freenet.api.ServerMessage, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.freenet.api.ServerMessage.displayName = 'proto.freenet.api.ServerMessage';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.freenet.api.SyncMessage.repeatedFields_ = [2,3];



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
proto.freenet.api.SyncMessage.prototype.toObject = function(opt_includeInstance) {
  return proto.freenet.api.SyncMessage.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.freenet.api.SyncMessage} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.SyncMessage.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    lanIpList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    vpnIpList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f
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
 * @return {!proto.freenet.api.SyncMessage}
 */
proto.freenet.api.SyncMessage.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.freenet.api.SyncMessage;
  return proto.freenet.api.SyncMessage.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.freenet.api.SyncMessage} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.freenet.api.SyncMessage}
 */
proto.freenet.api.SyncMessage.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addLanIp(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.addVpnIp(value);
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
proto.freenet.api.SyncMessage.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.freenet.api.SyncMessage.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.freenet.api.SyncMessage} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.SyncMessage.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getLanIpList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getVpnIpList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      3,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.freenet.api.SyncMessage.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.freenet.api.SyncMessage} returns this
 */
proto.freenet.api.SyncMessage.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string lan_ip = 2;
 * @return {!Array<string>}
 */
proto.freenet.api.SyncMessage.prototype.getLanIpList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.freenet.api.SyncMessage} returns this
 */
proto.freenet.api.SyncMessage.prototype.setLanIpList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.freenet.api.SyncMessage} returns this
 */
proto.freenet.api.SyncMessage.prototype.addLanIp = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.freenet.api.SyncMessage} returns this
 */
proto.freenet.api.SyncMessage.prototype.clearLanIpList = function() {
  return this.setLanIpList([]);
};


/**
 * repeated string vpn_ip = 3;
 * @return {!Array<string>}
 */
proto.freenet.api.SyncMessage.prototype.getVpnIpList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.freenet.api.SyncMessage} returns this
 */
proto.freenet.api.SyncMessage.prototype.setVpnIpList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.freenet.api.SyncMessage} returns this
 */
proto.freenet.api.SyncMessage.prototype.addVpnIp = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.freenet.api.SyncMessage} returns this
 */
proto.freenet.api.SyncMessage.prototype.clearVpnIpList = function() {
  return this.setVpnIpList([]);
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
proto.freenet.api.MessageAck.prototype.toObject = function(opt_includeInstance) {
  return proto.freenet.api.MessageAck.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.freenet.api.MessageAck} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.MessageAck.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.freenet.api.MessageAck}
 */
proto.freenet.api.MessageAck.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.freenet.api.MessageAck;
  return proto.freenet.api.MessageAck.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.freenet.api.MessageAck} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.freenet.api.MessageAck}
 */
proto.freenet.api.MessageAck.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
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
proto.freenet.api.MessageAck.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.freenet.api.MessageAck.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.freenet.api.MessageAck} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.MessageAck.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.freenet.api.MessageAck.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.freenet.api.MessageAck} returns this
 */
proto.freenet.api.MessageAck.prototype.setId = function(value) {
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
proto.freenet.api.Data.prototype.toObject = function(opt_includeInstance) {
  return proto.freenet.api.Data.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.freenet.api.Data} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.Data.toObject = function(includeInstance, msg) {
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
 * @return {!proto.freenet.api.Data}
 */
proto.freenet.api.Data.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.freenet.api.Data;
  return proto.freenet.api.Data.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.freenet.api.Data} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.freenet.api.Data}
 */
proto.freenet.api.Data.deserializeBinaryFromReader = function(msg, reader) {
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
proto.freenet.api.Data.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.freenet.api.Data.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.freenet.api.Data} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.Data.serializeBinaryToWriter = function(message, writer) {
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
proto.freenet.api.Data.prototype.getData = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes data = 1;
 * This is a type-conversion wrapper around `getData()`
 * @return {string}
 */
proto.freenet.api.Data.prototype.getData_asB64 = function() {
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
proto.freenet.api.Data.prototype.getData_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getData()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.freenet.api.Data} returns this
 */
proto.freenet.api.Data.prototype.setData = function(value) {
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
proto.freenet.api.ClientMessage.oneofGroups_ = [[1,2],[100]];

/**
 * @enum {number}
 */
proto.freenet.api.ClientMessage.RequestCase = {
  REQUEST_NOT_SET: 0,
  DATA: 1,
  SYNC_MESSAGE: 2
};

/**
 * @return {proto.freenet.api.ClientMessage.RequestCase}
 */
proto.freenet.api.ClientMessage.prototype.getRequestCase = function() {
  return /** @type {proto.freenet.api.ClientMessage.RequestCase} */(jspb.Message.computeOneofCase(this, proto.freenet.api.ClientMessage.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.freenet.api.ClientMessage.ResponseCase = {
  RESPONSE_NOT_SET: 0,
  MESSAGE_ACK: 100
};

/**
 * @return {proto.freenet.api.ClientMessage.ResponseCase}
 */
proto.freenet.api.ClientMessage.prototype.getResponseCase = function() {
  return /** @type {proto.freenet.api.ClientMessage.ResponseCase} */(jspb.Message.computeOneofCase(this, proto.freenet.api.ClientMessage.oneofGroups_[1]));
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
proto.freenet.api.ClientMessage.prototype.toObject = function(opt_includeInstance) {
  return proto.freenet.api.ClientMessage.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.freenet.api.ClientMessage} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.ClientMessage.toObject = function(includeInstance, msg) {
  var f, obj = {
    data: (f = msg.getData()) && proto.freenet.api.Data.toObject(includeInstance, f),
    syncMessage: (f = msg.getSyncMessage()) && proto.freenet.api.SyncMessage.toObject(includeInstance, f),
    messageAck: (f = msg.getMessageAck()) && proto.freenet.api.MessageAck.toObject(includeInstance, f)
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
 * @return {!proto.freenet.api.ClientMessage}
 */
proto.freenet.api.ClientMessage.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.freenet.api.ClientMessage;
  return proto.freenet.api.ClientMessage.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.freenet.api.ClientMessage} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.freenet.api.ClientMessage}
 */
proto.freenet.api.ClientMessage.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.freenet.api.Data;
      reader.readMessage(value,proto.freenet.api.Data.deserializeBinaryFromReader);
      msg.setData(value);
      break;
    case 2:
      var value = new proto.freenet.api.SyncMessage;
      reader.readMessage(value,proto.freenet.api.SyncMessage.deserializeBinaryFromReader);
      msg.setSyncMessage(value);
      break;
    case 100:
      var value = new proto.freenet.api.MessageAck;
      reader.readMessage(value,proto.freenet.api.MessageAck.deserializeBinaryFromReader);
      msg.setMessageAck(value);
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
proto.freenet.api.ClientMessage.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.freenet.api.ClientMessage.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.freenet.api.ClientMessage} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.ClientMessage.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getData();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.freenet.api.Data.serializeBinaryToWriter
    );
  }
  f = message.getSyncMessage();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.freenet.api.SyncMessage.serializeBinaryToWriter
    );
  }
  f = message.getMessageAck();
  if (f != null) {
    writer.writeMessage(
      100,
      f,
      proto.freenet.api.MessageAck.serializeBinaryToWriter
    );
  }
};


/**
 * optional Data data = 1;
 * @return {?proto.freenet.api.Data}
 */
proto.freenet.api.ClientMessage.prototype.getData = function() {
  return /** @type{?proto.freenet.api.Data} */ (
    jspb.Message.getWrapperField(this, proto.freenet.api.Data, 1));
};


/**
 * @param {?proto.freenet.api.Data|undefined} value
 * @return {!proto.freenet.api.ClientMessage} returns this
*/
proto.freenet.api.ClientMessage.prototype.setData = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.freenet.api.ClientMessage.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.freenet.api.ClientMessage} returns this
 */
proto.freenet.api.ClientMessage.prototype.clearData = function() {
  return this.setData(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.freenet.api.ClientMessage.prototype.hasData = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional SyncMessage sync_message = 2;
 * @return {?proto.freenet.api.SyncMessage}
 */
proto.freenet.api.ClientMessage.prototype.getSyncMessage = function() {
  return /** @type{?proto.freenet.api.SyncMessage} */ (
    jspb.Message.getWrapperField(this, proto.freenet.api.SyncMessage, 2));
};


/**
 * @param {?proto.freenet.api.SyncMessage|undefined} value
 * @return {!proto.freenet.api.ClientMessage} returns this
*/
proto.freenet.api.ClientMessage.prototype.setSyncMessage = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.freenet.api.ClientMessage.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.freenet.api.ClientMessage} returns this
 */
proto.freenet.api.ClientMessage.prototype.clearSyncMessage = function() {
  return this.setSyncMessage(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.freenet.api.ClientMessage.prototype.hasSyncMessage = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional MessageAck message_ack = 100;
 * @return {?proto.freenet.api.MessageAck}
 */
proto.freenet.api.ClientMessage.prototype.getMessageAck = function() {
  return /** @type{?proto.freenet.api.MessageAck} */ (
    jspb.Message.getWrapperField(this, proto.freenet.api.MessageAck, 100));
};


/**
 * @param {?proto.freenet.api.MessageAck|undefined} value
 * @return {!proto.freenet.api.ClientMessage} returns this
*/
proto.freenet.api.ClientMessage.prototype.setMessageAck = function(value) {
  return jspb.Message.setOneofWrapperField(this, 100, proto.freenet.api.ClientMessage.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.freenet.api.ClientMessage} returns this
 */
proto.freenet.api.ClientMessage.prototype.clearMessageAck = function() {
  return this.setMessageAck(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.freenet.api.ClientMessage.prototype.hasMessageAck = function() {
  return jspb.Message.getField(this, 100) != null;
};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.freenet.api.ServerMessage.oneofGroups_ = [[1,2],[100]];

/**
 * @enum {number}
 */
proto.freenet.api.ServerMessage.RequestCase = {
  REQUEST_NOT_SET: 0,
  DATA: 1,
  SYNC_MESSAGE: 2
};

/**
 * @return {proto.freenet.api.ServerMessage.RequestCase}
 */
proto.freenet.api.ServerMessage.prototype.getRequestCase = function() {
  return /** @type {proto.freenet.api.ServerMessage.RequestCase} */(jspb.Message.computeOneofCase(this, proto.freenet.api.ServerMessage.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.freenet.api.ServerMessage.ResponseCase = {
  RESPONSE_NOT_SET: 0,
  MESSAGE_ACK: 100
};

/**
 * @return {proto.freenet.api.ServerMessage.ResponseCase}
 */
proto.freenet.api.ServerMessage.prototype.getResponseCase = function() {
  return /** @type {proto.freenet.api.ServerMessage.ResponseCase} */(jspb.Message.computeOneofCase(this, proto.freenet.api.ServerMessage.oneofGroups_[1]));
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
proto.freenet.api.ServerMessage.prototype.toObject = function(opt_includeInstance) {
  return proto.freenet.api.ServerMessage.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.freenet.api.ServerMessage} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.ServerMessage.toObject = function(includeInstance, msg) {
  var f, obj = {
    data: (f = msg.getData()) && proto.freenet.api.Data.toObject(includeInstance, f),
    syncMessage: (f = msg.getSyncMessage()) && proto.freenet.api.SyncMessage.toObject(includeInstance, f),
    messageAck: (f = msg.getMessageAck()) && proto.freenet.api.MessageAck.toObject(includeInstance, f)
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
 * @return {!proto.freenet.api.ServerMessage}
 */
proto.freenet.api.ServerMessage.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.freenet.api.ServerMessage;
  return proto.freenet.api.ServerMessage.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.freenet.api.ServerMessage} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.freenet.api.ServerMessage}
 */
proto.freenet.api.ServerMessage.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.freenet.api.Data;
      reader.readMessage(value,proto.freenet.api.Data.deserializeBinaryFromReader);
      msg.setData(value);
      break;
    case 2:
      var value = new proto.freenet.api.SyncMessage;
      reader.readMessage(value,proto.freenet.api.SyncMessage.deserializeBinaryFromReader);
      msg.setSyncMessage(value);
      break;
    case 100:
      var value = new proto.freenet.api.MessageAck;
      reader.readMessage(value,proto.freenet.api.MessageAck.deserializeBinaryFromReader);
      msg.setMessageAck(value);
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
proto.freenet.api.ServerMessage.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.freenet.api.ServerMessage.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.freenet.api.ServerMessage} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.freenet.api.ServerMessage.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getData();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.freenet.api.Data.serializeBinaryToWriter
    );
  }
  f = message.getSyncMessage();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.freenet.api.SyncMessage.serializeBinaryToWriter
    );
  }
  f = message.getMessageAck();
  if (f != null) {
    writer.writeMessage(
      100,
      f,
      proto.freenet.api.MessageAck.serializeBinaryToWriter
    );
  }
};


/**
 * optional Data data = 1;
 * @return {?proto.freenet.api.Data}
 */
proto.freenet.api.ServerMessage.prototype.getData = function() {
  return /** @type{?proto.freenet.api.Data} */ (
    jspb.Message.getWrapperField(this, proto.freenet.api.Data, 1));
};


/**
 * @param {?proto.freenet.api.Data|undefined} value
 * @return {!proto.freenet.api.ServerMessage} returns this
*/
proto.freenet.api.ServerMessage.prototype.setData = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.freenet.api.ServerMessage.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.freenet.api.ServerMessage} returns this
 */
proto.freenet.api.ServerMessage.prototype.clearData = function() {
  return this.setData(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.freenet.api.ServerMessage.prototype.hasData = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional SyncMessage sync_message = 2;
 * @return {?proto.freenet.api.SyncMessage}
 */
proto.freenet.api.ServerMessage.prototype.getSyncMessage = function() {
  return /** @type{?proto.freenet.api.SyncMessage} */ (
    jspb.Message.getWrapperField(this, proto.freenet.api.SyncMessage, 2));
};


/**
 * @param {?proto.freenet.api.SyncMessage|undefined} value
 * @return {!proto.freenet.api.ServerMessage} returns this
*/
proto.freenet.api.ServerMessage.prototype.setSyncMessage = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.freenet.api.ServerMessage.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.freenet.api.ServerMessage} returns this
 */
proto.freenet.api.ServerMessage.prototype.clearSyncMessage = function() {
  return this.setSyncMessage(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.freenet.api.ServerMessage.prototype.hasSyncMessage = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional MessageAck message_ack = 100;
 * @return {?proto.freenet.api.MessageAck}
 */
proto.freenet.api.ServerMessage.prototype.getMessageAck = function() {
  return /** @type{?proto.freenet.api.MessageAck} */ (
    jspb.Message.getWrapperField(this, proto.freenet.api.MessageAck, 100));
};


/**
 * @param {?proto.freenet.api.MessageAck|undefined} value
 * @return {!proto.freenet.api.ServerMessage} returns this
*/
proto.freenet.api.ServerMessage.prototype.setMessageAck = function(value) {
  return jspb.Message.setOneofWrapperField(this, 100, proto.freenet.api.ServerMessage.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.freenet.api.ServerMessage} returns this
 */
proto.freenet.api.ServerMessage.prototype.clearMessageAck = function() {
  return this.setMessageAck(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.freenet.api.ServerMessage.prototype.hasMessageAck = function() {
  return jspb.Message.getField(this, 100) != null;
};


goog.object.extend(exports, proto.freenet.api);