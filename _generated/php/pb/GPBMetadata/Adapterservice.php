<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: adapterservice.proto

namespace GPBMetadata;

class Adapterservice
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Status::initOnce();
        \GPBMetadata\Station::initOnce();
        \GPBMetadata\Producer::initOnce();
        \GPBMetadata\Consumer::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
adapterservice.protopbstation.protoproducer.protoconsumer.proto2�
AdapterService7
CreateStation.pb.CreateStationRequest
.pb.Status" 9
DestroyStation.pb.DestroyStationRequest
.pb.Status" .
Produce.pb.ProduceMessages
.pb.Status" (9
Consume.pb.ConsumeMessages.pb.ConsumeResponse" (0B"Z github.com/g41797/memphisgrpc;pbbproto3'
        , true);

        static::$is_initialized = true;
    }
}

