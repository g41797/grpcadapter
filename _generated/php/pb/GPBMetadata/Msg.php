<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: msg.proto

namespace GPBMetadata;

class Msg
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
	msg.protopb"d
Headers)
headers (2.pb.Headers.HeadersEntry.
HeadersEntry
key (	
value (	:8"P
Msg!
headers (2.pb.HeadersH �
body (H�B

_headersB
_bodyB"Z github.com/g41797/memphisgrpc;pbbproto3'
        , true);

        static::$is_initialized = true;
    }
}

