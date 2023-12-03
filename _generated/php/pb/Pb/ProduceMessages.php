<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: producer.proto

namespace Pb;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>pb.ProduceMessages</code>
 */
class ProduceMessages extends \Google\Protobuf\Internal\Message
{
    protected $data;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Pb\CreateProducerRequest $start
     *     @type \Pb\Msg $msg
     *     @type \Pb\Stop $stop
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Producer::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.pb.CreateProducerRequest start = 1;</code>
     * @return \Pb\CreateProducerRequest|null
     */
    public function getStart()
    {
        return $this->readOneof(1);
    }

    public function hasStart()
    {
        return $this->hasOneof(1);
    }

    /**
     * Generated from protobuf field <code>.pb.CreateProducerRequest start = 1;</code>
     * @param \Pb\CreateProducerRequest $var
     * @return $this
     */
    public function setStart($var)
    {
        GPBUtil::checkMessage($var, \Pb\CreateProducerRequest::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.pb.Msg msg = 2;</code>
     * @return \Pb\Msg|null
     */
    public function getMsg()
    {
        return $this->readOneof(2);
    }

    public function hasMsg()
    {
        return $this->hasOneof(2);
    }

    /**
     * Generated from protobuf field <code>.pb.Msg msg = 2;</code>
     * @param \Pb\Msg $var
     * @return $this
     */
    public function setMsg($var)
    {
        GPBUtil::checkMessage($var, \Pb\Msg::class);
        $this->writeOneof(2, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.pb.Stop stop = 3;</code>
     * @return \Pb\Stop|null
     */
    public function getStop()
    {
        return $this->readOneof(3);
    }

    public function hasStop()
    {
        return $this->hasOneof(3);
    }

    /**
     * Generated from protobuf field <code>.pb.Stop stop = 3;</code>
     * @param \Pb\Stop $var
     * @return $this
     */
    public function setStop($var)
    {
        GPBUtil::checkMessage($var, \Pb\Stop::class);
        $this->writeOneof(3, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getData()
    {
        return $this->whichOneof("data");
    }

}

