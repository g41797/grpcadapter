<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: msg.proto

namespace g4197\memphisphp\pb;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>pb.Headers</code>
 */
class Headers extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>map<string, string> headers = 1;</code>
     */
    private $headers;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array|\Google\Protobuf\Internal\MapField $headers
     * }
     */
    public function __construct($data = NULL) {
        \g4197\memphisphp\metadata\Msg::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>map<string, string> headers = 1;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getHeaders()
    {
        return $this->headers;
    }

    /**
     * Generated from protobuf field <code>map<string, string> headers = 1;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setHeaders($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->headers = $arr;

        return $this;
    }

}

