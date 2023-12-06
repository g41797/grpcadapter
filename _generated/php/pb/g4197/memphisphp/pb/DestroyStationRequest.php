<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: station.proto

namespace g4197\memphisphp\pb;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>pb.DestroyStationRequest</code>
 */
class DestroyStationRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.pb.Station station = 1;</code>
     */
    protected $station = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \g4197\memphisphp\pb\Station $station
     * }
     */
    public function __construct($data = NULL) {
        \g4197\memphisphp\metadata\Station::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.pb.Station station = 1;</code>
     * @return \g4197\memphisphp\pb\Station|null
     */
    public function getStation()
    {
        return $this->station;
    }

    public function hasStation()
    {
        return isset($this->station);
    }

    public function clearStation()
    {
        unset($this->station);
    }

    /**
     * Generated from protobuf field <code>.pb.Station station = 1;</code>
     * @param \g4197\memphisphp\pb\Station $var
     * @return $this
     */
    public function setStation($var)
    {
        GPBUtil::checkMessage($var, \g4197\memphisphp\pb\Station::class);
        $this->station = $var;

        return $this;
    }

}

