# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: topic.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0btopic.proto\x12\x0ftopic_selection\";\n\x0b\x46\x61\x63tRequest\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\x12\x0e\n\x06second\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\"\x1d\n\x0c\x46\x61\x63tResponse\x12\r\n\x05\x66\x61\x63ts\x18\x01 \x03(\t\";\n\x0c\x43olorRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06second\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\"6\n\rColorResponse\x12\x0b\n\x03hex\x18\x01 \x01(\t\x12\x0b\n\x03rgb\x18\x02 \x01(\t\x12\x0b\n\x03hsl\x18\x03 \x01(\t2\xb6\x01\n\x0eTopicSelection\x12P\n\x0fStreamMeowFacts\x12\x1c.topic_selection.FactRequest\x1a\x1d.topic_selection.FactResponse0\x01\x12R\n\x0fStreamColorInfo\x12\x1d.topic_selection.ColorRequest\x1a\x1e.topic_selection.ColorResponse0\x01\x42\x15Z\x13go_server/topic_pb2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'topic_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\023go_server/topic_pb2'
  _globals['_FACTREQUEST']._serialized_start=32
  _globals['_FACTREQUEST']._serialized_end=91
  _globals['_FACTRESPONSE']._serialized_start=93
  _globals['_FACTRESPONSE']._serialized_end=122
  _globals['_COLORREQUEST']._serialized_start=124
  _globals['_COLORREQUEST']._serialized_end=183
  _globals['_COLORRESPONSE']._serialized_start=185
  _globals['_COLORRESPONSE']._serialized_end=239
  _globals['_TOPICSELECTION']._serialized_start=242
  _globals['_TOPICSELECTION']._serialized_end=424
# @@protoc_insertion_point(module_scope)
